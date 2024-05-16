package client

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/atko-pam/pam-sdk-go/client/pam"

	"github.com/go-resty/resty/v2"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
	"github.com/okta/terraform-provider-oktapam/oktapam/version"
)

const OKTAPAM_TRUSTED_DOMAINS = "scaleft.com,scaleft.io,okta.com,oktapreview.com,okta-emea.com"
const TRUSTED_DOMAIN_OVERRIDE_ENV_VAR = "OKTAPAM_TRUSTED_DOMAIN_OVERRIDE"

var terraformUserAgent = "terraform_provider_oktapam/" + version.Version

type ServiceToken struct {
	TeamName    string    `json:"team_name"`
	BearerToken string    `json:"bearer_token"`
	ExpiresAt   time.Time `json:"expires_at"`
}

type OktaPAMClient struct {
	Team   string
	client *resty.Client
}

func CreateLocalPAMClient(providerConfig *OktaPAMProviderConfig) (*OktaPAMClient, error) {
	logging.Infof("Creating PAM Client")
	if serviceToken, err := createServiceToken(providerConfig.APIKey, providerConfig.APIKeySecret, providerConfig.APIHost, providerConfig.Team); err != nil {
		return nil, err
	} else {
		client := setBaseHTTPSettings(resty.New(), providerConfig.APIHost, *serviceToken)
		client = setRateLimitRetryLogic(client)

		return &OktaPAMClient{Team: providerConfig.Team, client: client}, nil
	}
}

func checkTrustedDomain(apiHost string) error {
	// this is more of a check for misconfiguration than anything as we allow it to be overriden
	// for testing/debugging purposes
	u, err := url.Parse(apiHost)
	if err != nil {
		logging.Errorf("error while parsing api host value: %s, \n%w", apiHost, err)
		return err
	}
	hostname := u.Hostname()

	for _, domain := range strings.Split(OKTAPAM_TRUSTED_DOMAINS, ",") {
		if strings.HasSuffix(hostname, domain) {
			return nil
		}
	}

	override, present := os.LookupEnv(TRUSTED_DOMAIN_OVERRIDE_ENV_VAR)
	if present {
		if hostname == override {
			return nil
		}
		return fmt.Errorf("configured api host is not within a trusted domain.  host: %q, override: %q", hostname, override)
	}

	return fmt.Errorf("configured api host is not within a trusted domain.  host: %q", hostname)
}

func createServiceToken(apiKey, apiKeySecret, apiHost, team string) (*ServiceToken, error) {
	err := checkTrustedDomain(apiHost)
	if err != nil {
		return nil, err
	}

	authorizationURL := fmt.Sprintf("%s/v1/teams/%s/service_token", apiHost, url.PathEscape(team))
	client := resty.New()

	resp, err := client.R().
		SetHeader("User-Agent", terraformUserAgent).
		SetBody(map[string]any{"key_id": apiKey, "key_secret": apiKeySecret}).
		SetHeaders(map[string]string{"Accept": "application/json", "Content-Type": "application/json"}).
		SetResult(&ServiceToken{}).
		Post(authorizationURL)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() == 401 {
		return nil, fmt.Errorf("received a 401 from URL %s when requesting service token.  check credentials and try again", authorizationURL)
	}

	_, err = checkStatusCode(resp, http.StatusOK)
	if err != nil {
		return nil, err
	}

	serviceToken := resp.Result().(*ServiceToken)
	if !strings.EqualFold(serviceToken.TeamName, team) {
		return nil, fmt.Errorf("service token team name does not match supplied team name, expected: %s, got: %s", team, serviceToken.TeamName)
	}

	return serviceToken, nil
}

func setBaseHTTPSettings(client *resty.Client, apiHost string, serviceToken ServiceToken) *resty.Client {
	return client.
		SetBaseURL(apiHost).
		SetAuthToken(serviceToken.BearerToken).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetHeader("User-Agent", terraformUserAgent)
}

func setRateLimitRetryLogic(client *resty.Client) *resty.Client {
	return client.
		AddRetryCondition(func(r *resty.Response, e error) bool {
			return r.StatusCode() == http.StatusTooManyRequests
		}).
		SetRetryMaxWaitTime(time.Hour * 24). // high duration to ensure we don't cap the retry length given
		SetRetryCount(2147483647).           // just set to a high number to indicate we want to keep retrying
		SetRetryAfter(func(c *resty.Client, r *resty.Response) (time.Duration, error) {
			headers := r.Header()
			retryAtHeader := headers.Get("X-RateLimit-Retry-At")
			if retryAtHeader == "" {
				// if we don't get a header, retry between 5-10 seconds from now
				waitTime := time.Millisecond * time.Duration(rand.Intn(5000)+5000)
				logging.Infof("Request was rate limited, waiting %s to retry again", waitTime)
				return waitTime, nil
			}

			laterUnix, err := strconv.ParseInt(retryAtHeader, 10, 64)
			if err != nil {
				return 0, err
			}
			later := time.Unix(laterUnix, 0)
			diff := time.Until(later)

			// duration returned is the diff between now and the time given by
			// the server, plus jitter between 100-3000ms
			waitTime := diff + (time.Millisecond * time.Duration(rand.Intn(2900)+100))
			logging.Infof("Request was rate limited, waiting %s to retry again", waitTime)

			return waitTime, nil
		})
}

func (c OktaPAMClient) CreateBaseRequest(ctx context.Context) *resty.Request {
	return c.client.R().SetContext(ctx)
}

func checkStatusCode(resp *resty.Response, allowed ...int) (int, error) {
	received := resp.StatusCode()
	for _, c := range allowed {
		if received == c {
			return received, nil
		}
	}
	return received, createErrorForInvalidCode(resp, allowed...)
}

func createErrorForInvalidCode(resp *resty.Response, allowed ...int) error {
	if len(allowed) == 1 {
		return fmt.Errorf("call to %s resulted in status of %d, expected a %d.\nResponse Body: %s", resp.Request.URL, resp.StatusCode(), allowed[0], string(resp.Body()))
	}

	return fmt.Errorf("call to %s resulted in status of %d, expected one of %v.\nResponse Body: %s", resp.Request.URL, resp.StatusCode(), allowed, string(resp.Body()))
}

func parseBool(i any) (bool, error) {
	switch v := i.(type) {
	case bool:
		return v, nil
	case int:
		return v != 0, nil
	case string:
		b, err := strconv.ParseBool(v)
		if err != nil {
			return false, err
		}
		return b, nil
	default:
		return false, fmt.Errorf("cannot convert %T to bool", v)
	}
}

func stringSliceToInterfaceSlice(arr []string) []any {
	arrI := make([]any, len(arr))

	for idx, s := range arr {
		arrI[idx] = s
	}

	return arrI
}

type OktaPAMProviderConfig struct {
	APIKey       string
	APIKeySecret string
	Team         string
	APIHost      string
}

type APIClients struct {
	SDKClient   SDKClientWrapper
	LocalClient *OktaPAMClient
}

func CreateSDKClient(providerConfig *OktaPAMProviderConfig) (*pam.APIClient, error) {
	apiClientConfigOpts := []pam.ConfigOption{
		pam.WithHost(providerConfig.APIHost),
		pam.WithTeam(providerConfig.Team),
		pam.WithAPIKey(providerConfig.APIKey),
		pam.WithAPISecret(providerConfig.APIKeySecret),
		pam.WithUserAgent(terraformUserAgent),
	}

	override, present := os.LookupEnv(TRUSTED_DOMAIN_OVERRIDE_ENV_VAR)
	if present {
		apiClientConfigOpts = append(apiClientConfigOpts, pam.WithTrustedDomainOverride(override))
	}
	pamClient, err := pam.NewAPIClient(apiClientConfigOpts...)
	if err != nil {
		return nil, fmt.Errorf("error while creating sdk client: %w", err)
	}
	return pamClient, nil
}
