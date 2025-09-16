package pam

import (
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"

	"github.com/kelseyhightower/envconfig"
	"golang.org/x/oauth2"
)

var (
	ErrAuthzModeAPIKeyOrSecretMissing = errors.New("when authorization Mode is set to 'APIKey', both 'APIKey' and 'APISecret' must be provided")
	ErrAuthzModeBearerToken           = errors.New("when authorization Mode is set to 'BearerToken', 'BearerToken' must be provided")
	ErrAuthzModeInvalid               = errors.New("authorization mode config option must be set to one of [APIKey, BearerToken]")
	ErrTLSRootCAsConflict             = errors.New("TLSRootCAs and TLSInsecureSkipVerify cannot be set at the same time")
	ErrOauth2TokenSourceMissing       = errors.New("oauth2 token source missing")
)

const OktaPAMTrustedDomains = "scaleft.com,scaleft.io,okta.com,oktapreview.com,okta-emea.com"

// RequestTimeoutDuration is a custom type that will be able to take a string and convert
// it to a time.Duration
type RequestTimeoutDuration time.Duration

// Decode implements the envconfig.Decoder interface.
func (d *RequestTimeoutDuration) Decode(value string) error {
	duration, err := time.ParseDuration(value)
	if err != nil {
		return err
	}
	*d = RequestTimeoutDuration(duration)
	return nil
}

type AuthorizationMode string

const (
	BearerToken       AuthorizationMode = "BearerToken"
	APIKey            AuthorizationMode = "APIKey"
	OAuth2TokenSource AuthorizationMode = "OAuth2TokenSource"
)

func (a *AuthorizationMode) Decode(value string) error {
	switch value {
	case string(BearerToken):
		*a = BearerToken
	case string(APIKey):
		*a = APIKey
	case string(OAuth2TokenSource):
		*a = OAuth2TokenSource
	default:
		return fmt.Errorf("unsupported auth mode: %s", value)
	}
	return nil
}

// configuration stores the configuration of the API client
type configuration struct {
	// YAML tags can also be added if needed
	Host                  string            `envconfig:"HOST"`
	TeamName              string            `envconfig:"TEAM_NAME"`
	DefaultHeader         map[string]string `envconfig:"DEFAULT_HEADER"`
	UserAgent             string            `envconfig:"USER_AGENT"`
	TrustedDomainOverride string            `envconfig:"TRUSTED_DOMAIN_OVERRIDE"`

	// Request timeout and retries settings
	RequestTimeout RequestTimeoutDuration `envconfig:"REQUEST_TIMEOUT"`
	RateLimit      struct {
		MaxRetries int `envconfig:"MAX_RETRIES"`
		MaxBackoff int `envconfig:"MAX_BACKOFF"`
	} `envconfig:"RATE_LIMIT"`

	//Authentication Config
	AuthorizationMode AuthorizationMode  `envconfig:"AUTHORIZATION_MODE"`
	BearerToken       string             `envconfig:"BEARER_TOKEN"`
	APIKey            string             `envconfig:"API_KEY"`
	APISecret         string             `envconfig:"API_SECRET"`
	OAuth2TokenSource oauth2.TokenSource `ignored:"true"`

	// HTTP Transport Config
	TLSUseBundledCAs      bool           `envconfig:"TLS_USE_BUNDLED_CA"`
	TLSInsecureSkipVerify bool           `envconfig:"TLS_INSECURE_SKIP_VERIFY"`
	TLSRootCAs            *x509.CertPool `ignored:"true"`

	// You can also set it via http environment variable. Refer to godoc `http.ProxyFromEnvironment`
	ProxyURL     string            `envconfig:"PROXY_URL"`
	RoundTripper http.RoundTripper `ignored:"true"`

	//Logging and debugging
	//TODO Provide logger for configuration if needed
	EnableHTTPDebug bool `envconfig:"DEBUG"`
	EnableHTTPTrace bool `envconfig:"ENABLE_HTTP_TRACE"`
	EnablePrintCurl bool `envconfig:"ENABLE_PRINT_CURL"`

	EnableTelemetry bool `envconfig:"ENABLE_OTEL"`

	UserAgentHook func() string `ignored:"true"`

	// ClientFeatures should be set using the WithClientFeatures fn
	ClientFeatures map[string]bool `ignored:"true"`
}

type TestEnv struct {
	Host string
}
type ConfigOption func(*configuration)

// WithRoundTripper Sets custom `*http.Transport` or any `http.RoundTripper` compatible interface implementation.
// It Overwrites the http client transport instance.
func WithRoundTripper(rt http.RoundTripper) ConfigOption {
	return func(c *configuration) {
		c.RoundTripper = rt
	}
}

func WithTelemetry(telemetry bool) ConfigOption {
	return func(c *configuration) {
		c.EnableTelemetry = telemetry
	}
}

func WithHost(host string) ConfigOption {
	return func(c *configuration) {
		c.Host = host
	}
}

func WithRequestTimeout(requestTimeout time.Duration) ConfigOption {
	return func(c *configuration) {
		c.RequestTimeout = RequestTimeoutDuration(requestTimeout)
	}
}

func WithRateLimitMaxRetries(maxRetries int) ConfigOption {
	return func(c *configuration) {
		c.RateLimit.MaxRetries = maxRetries
	}
}

func WithRateLimitMaxBackOff(maxBackoff int) ConfigOption {
	return func(c *configuration) {
		c.RateLimit.MaxBackoff = maxBackoff
	}
}

// WithUserAgent Override default value "pam-sdk-go golang/{go-version} {os-name}/{os-version}
func WithUserAgent(userAgent string) ConfigOption {
	return func(c *configuration) {
		c.UserAgent = userAgent
	}
}

func WithAPIKey(key string) ConfigOption {
	return func(c *configuration) {
		c.APIKey = key
		c.AuthorizationMode = APIKey
	}
}

func WithAPISecret(secret string) ConfigOption {
	return func(c *configuration) {
		c.APISecret = secret
		c.AuthorizationMode = APIKey
	}
}

func WithTeam(teamName string) ConfigOption {
	return func(c *configuration) {
		c.TeamName = teamName
	}
}

func WithEnableHTTPDebug(enableHTTPDebug bool) ConfigOption {
	return func(c *configuration) {
		c.EnableHTTPDebug = enableHTTPDebug
	}
}

func WithEnableHTTPTrace(enableHTTPTrace bool) ConfigOption {
	return func(c *configuration) {
		c.EnableHTTPTrace = enableHTTPTrace
	}
}

func WithEnablePrintCurl(enablePrintCurl bool) ConfigOption {
	return func(c *configuration) {
		c.EnablePrintCurl = enablePrintCurl
	}
}

func WithTrustedDomainOverride(trustedDomainOverride string) ConfigOption {
	return func(c *configuration) {
		c.TrustedDomainOverride = trustedDomainOverride
	}
}

func WithBearerToken(bearerToken string) ConfigOption {
	return func(c *configuration) {
		c.BearerToken = bearerToken
		c.AuthorizationMode = BearerToken
	}
}

func WithTLSUseBundledCAs(tlsUseBundledCAs bool) ConfigOption {
	return func(c *configuration) {
		c.TLSUseBundledCAs = tlsUseBundledCAs
	}
}

// WithTLSRootCAs Set custom root-certificate for TLS. By default, client use system root CAs.
func WithTLSRootCAs(tlsRootCAs x509.CertPool) ConfigOption {
	return func(c *configuration) {
		c.TLSRootCAs = &tlsRootCAs
	}
}

// WithTLSInsecureSkipVerify disable security check, use only for testing purposes
func WithTLSInsecureSkipVerify(tlsInsecureSkipVerify bool) ConfigOption {
	return func(c *configuration) {
		c.TLSInsecureSkipVerify = tlsInsecureSkipVerify
	}
}

// WithProxyURL Set Proxy url for http client transport. Either use this option or overwrite http client transport via
// WithRoundTripper option
func WithProxyURL(proxyURL string) ConfigOption {
	return func(c *configuration) {
		c.ProxyURL = proxyURL
	}
}

// WithOAuth2TokenSource Called before every HTTP request to retrieve auth token
func WithOAuth2TokenSource(tokenSource oauth2.TokenSource) ConfigOption {
	return func(c *configuration) {
		c.OAuth2TokenSource = tokenSource
		c.AuthorizationMode = OAuth2TokenSource
	}
}

// WithUserAgentHook Accepts function that returns User-Agent Header value. Called before every HTTP request
func WithUserAgentHook(userAgentFn func() string) ConfigOption {
	return func(c *configuration) {
		c.UserAgentHook = userAgentFn
	}
}

func WithClientFeatures(clientFeatures ...string) ConfigOption {
	return func(c *configuration) {
		if c.ClientFeatures == nil {
			c.ClientFeatures = map[string]bool{}
		}
		for _, clientFeature := range clientFeatures {
			c.ClientFeatures[clientFeature] = true
		}
	}
}

// AddDefaultHeader adds a new HTTP header to the default header in the request
func (c *configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}

func newConfiguration(opts ...ConfigOption) (*configuration, error) {
	config := &configuration{}
	// Set defaults
	setConfigDefaults(config)

	config, err := readConfigFromEnvironment(*config)
	if err != nil {
		return nil, fmt.Errorf("error while parsing environment variables %v", err)
	}
	for _, opt := range opts {
		opt(config)
	}

	config, err = validateConfig(config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func validateConfig(c *configuration) (*configuration, error) {
	if err := checkTrustedDomain(c); err != nil {
		return nil, err
	}

	if err := checkAuthorization(c); err != nil {
		return nil, err
	}

	if err := checkTLSConfig(c); err != nil {
		return nil, err
	}
	return c, nil
}

func checkTrustedDomain(c *configuration) error {
	// this is more of a check for misconfiguration than anything as we allow it to be overriden
	// for testing/debugging purposes
	u, err := url.Parse(c.Host)
	if err != nil {
		return err
	}
	hostname := u.Hostname()

	for _, domain := range strings.Split(OktaPAMTrustedDomains, ",") {
		if strings.HasSuffix(hostname, domain) {
			return nil
		}
	}

	if c.TrustedDomainOverride != "" {
		if hostname == c.TrustedDomainOverride {
			return nil
		}
		return fmt.Errorf("configured api host is not within a trusted domain.  host: %q, override: %q", hostname, c.TrustedDomainOverride)
	}

	return fmt.Errorf("configured api host is not within a trusted domain.  host: %q", hostname)
}

func checkAuthorization(c *configuration) error {
	switch authMode := c.AuthorizationMode; authMode {
	case APIKey:
		if c.APIKey == "" || c.APISecret == "" {
			return ErrAuthzModeAPIKeyOrSecretMissing
		}
	case BearerToken:
		if c.BearerToken == "" {
			return ErrAuthzModeBearerToken
		}
	case OAuth2TokenSource:
		if c.OAuth2TokenSource == nil {
			return ErrOauth2TokenSourceMissing
		}
	default:
		return ErrAuthzModeInvalid
	}

	return nil
}

func checkTLSConfig(c *configuration) error {
	if c.TLSRootCAs != nil && c.TLSInsecureSkipVerify {
		return ErrTLSRootCAsConflict
	}
	return nil
}

func setConfigDefaults(c *configuration) {
	conf := []ConfigOption{
		WithUserAgent(NewUserAgent().String()),
		WithRequestTimeout(60 * time.Second),
		WithRateLimitMaxBackOff(30),
		WithRateLimitMaxRetries(3),
		WithEnableHTTPDebug(false),
		WithEnableHTTPTrace(false),
		WithEnablePrintCurl(false),
		WithTelemetry(true),
		WithTLSUseBundledCAs(false),
	}
	for _, confSetter := range conf {
		confSetter(c)
	}
}

func readConfigFromEnvironment(c configuration) (*configuration, error) {
	if err := envconfig.Process("OKTAPAM", &c); err != nil {
		return nil, err
	}
	// automatically set authorization mode
	if c.APIKey != "" || c.APISecret != "" {
		c.AuthorizationMode = APIKey
	} else if c.BearerToken != "" {
		c.AuthorizationMode = BearerToken
	}

	return &c, nil
}

func printcURL(fullURL string, req *resty.Request) error {
	command := fmt.Sprintf("curl -X %s '%s'", req.Method, fullURL)

	for k, v := range req.Header {
		command += fmt.Sprintf(" -H '%s: %s'", k, strings.Join(v, ", "))
	}
	bodyBytes, err := json.Marshal(req.Body)
	if err != nil {
		return err
	}
	command += fmt.Sprintf(" -d %q", string(bodyBytes))

	_, _ = fmt.Fprintf(os.Stderr, "cURL Command: %s\n", command)
	return nil
}

type UserAgent struct {
	goVersion string
	osName    string
	osVersion string
}

func NewUserAgent() UserAgent {
	ua := UserAgent{}
	ua.goVersion = runtime.Version()
	ua.osName = runtime.GOOS
	ua.osVersion = runtime.GOARCH
	return ua
}

func (ua UserAgent) String() string {
	// TODO Add SDK Version by getting it from git tags
	sdkString := strings.Join([]string{"pam-sdk-go"}, "/")
	golangString := strings.Join([]string{"golang", ua.goVersion}, "/")
	osString := strings.Join([]string{ua.osName, ua.osVersion}, "/")

	return strings.Join([]string{sdkString, golangString, osString}, " ")
}
