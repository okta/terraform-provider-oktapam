package client

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/terraform-providers/terraform-provider-oktaasa/oktaasa/logging"
	"github.com/terraform-providers/terraform-provider-oktaasa/oktaasa/version"
)

var terraformUserAgent = "terraform_provider_oktaasa/" + version.Version

type ServiceToken struct {
	TeamName    string    `json:"team_name"`
	BearerToken string    `json:"bearer_token"`
	ExpiresAt   time.Time `json:"expires_at"`
}

type OktaASAClient struct {
	Team   string
	client *resty.Client
}

func CreateOktaASAClient(apiKey, apiKeySecret, team, apiHost string) (OktaASAClient, error) {
	logging.Infof("Creating ASA Client")
	authorizationURL := fmt.Sprintf("%s/v1/teams/%s/service_token", apiHost, team)
	client := resty.New()

	resp, err := client.R().
		SetHeader("User-Agent", terraformUserAgent).
		SetBody(map[string]interface{}{"key_id": apiKey, "key_secret": apiKeySecret}).
		SetHeaders(map[string]string{"Accept": "application/json", "Content-Type": "application/json"}).
		SetResult(&ServiceToken{}).
		Post(authorizationURL)
	if err != nil {
		return OktaASAClient{}, err
	}
	if resp.StatusCode() == 401 {
		return OktaASAClient{}, fmt.Errorf("received a 401 when requesting service token.  check credentials and try again")
	}

	err = checkStatusCode(resp, 200)
	if err != nil {
		return OktaASAClient{}, err
	}

	serviceToken := resp.Result().(*ServiceToken)
	if !strings.EqualFold(serviceToken.TeamName, team) {
		return OktaASAClient{}, fmt.Errorf("service token team name does not match supplied team name, expected: %s, got: %s", team, serviceToken.TeamName)
	}

	return newOktaASAClient(apiHost, team, *serviceToken), nil
}

func newOktaASAClient(apiHost, team string, serviceToken ServiceToken) OktaASAClient {
	client := setBaseHTTPSettings(resty.New(), apiHost, serviceToken)
	client = setRateLimitRetryLogic(client)

	return OktaASAClient{Team: team, client: client}
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
			now := time.Now()
			laterUnix, err := strconv.ParseInt(retryAtHeader, 10, 64)
			if err != nil {
				return 0, err
			}
			later := time.Unix(laterUnix, 0)
			diff := later.Unix() - now.Unix()

			// duration returned is the diff between now and the time given by
			// the server, plus jitter between 100-3000ms
			waitTime := (time.Second * time.Duration(diff)) + (time.Millisecond * time.Duration(rand.Intn(2900)+100))
			logging.Infof("Request was rate limited, waiting %s to retry again", waitTime)

			return waitTime, nil
		})
}

func (c OktaASAClient) CreateBaseRequest(ctx context.Context) *resty.Request {
	return c.client.R().SetContext(ctx)
}

func checkStatusCode(resp *resty.Response, allowed ...int) error {
	received := resp.StatusCode()
	for _, c := range allowed {
		if received == c {
			return nil
		}
	}
	return createErrorForInvalidCode(resp, allowed...)
}

func createErrorForInvalidCode(resp *resty.Response, allowed ...int) error {
	if len(allowed) == 1 {
		return fmt.Errorf("call to %s resulted in status of %d, expected a %d.\nResponse Body: %s", resp.Request.URL, resp.StatusCode(), allowed[0], string(resp.Body()))
	}

	return fmt.Errorf("call to %s resulted in status of %d, expected one of %v.\nResponse Body: %s", resp.Request.URL, resp.StatusCode(), allowed, string(resp.Body()))
}

func parseBool(i interface{}) (bool, error) {
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
