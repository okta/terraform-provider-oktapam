package pam

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	opentracing "github.com/opentracing/opentracing-go"

	"github.com/go-resty/resty/v2"
	"gopkg.in/square/go-jose.v2"
)

type ErrNonDefaultResponse struct {
	StatusCode int
	Result     any
}

func (endr ErrNonDefaultResponse) Error() string {
	return fmt.Sprintf("got http status code %d, providing non-default result", endr.StatusCode)
}

// APIClient manages communication with the Okta Privileged Access API v1.0.0
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg         *configuration
	restyClient *resty.Client
	tokenCache  *authTokenCache

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services
	AttributesAPI *AttributesAPIService

	ClientsAPI *ClientsAPIService

	GatewaysAPI *GatewaysAPIService

	GroupsAPI *GroupsAPIService

	ProjectsAPI *ProjectsAPIService

	ResourceGroupsAPI *ResourceGroupsAPIService

	SecurityPolicyAPI *SecurityPolicyAPIService

	ServersAPI *ServersAPIService

	ServiceUsersAPI *ServiceUsersAPIService

	TeamsAPI *TeamsAPIService

	UsersAPI *UsersAPIService

	SecretsAPI *SecretsAPIService

	DatabaseResourcesAPI *DatabaseResourcesAPIService

	AccessReportsAPI *ReportsAPIService

	CloudEntiltlementsAPI *CloudEntitlementsAPIService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(opts ...ConfigOption) (*APIClient, error) {
	//Create configuration
	config, err := newConfiguration(opts...)

	if err != nil {
		return nil, err
	}

	apiClient := &APIClient{
		cfg: config,
	}

	//Setup http client
	apiClient.restyClient = newRestyClient(apiClient)
	apiClient.tokenCache = NewAuthTokenCache(apiClient.cfg)

	apiClient.common.client = apiClient

	// API Services
	apiClient.AttributesAPI = (*AttributesAPIService)(&apiClient.common)
	apiClient.ClientsAPI = (*ClientsAPIService)(&apiClient.common)
	apiClient.GatewaysAPI = (*GatewaysAPIService)(&apiClient.common)
	apiClient.GroupsAPI = (*GroupsAPIService)(&apiClient.common)
	apiClient.ProjectsAPI = (*ProjectsAPIService)(&apiClient.common)
	apiClient.ResourceGroupsAPI = (*ResourceGroupsAPIService)(&apiClient.common)
	apiClient.SecurityPolicyAPI = (*SecurityPolicyAPIService)(&apiClient.common)
	apiClient.ServersAPI = (*ServersAPIService)(&apiClient.common)
	apiClient.ServiceUsersAPI = (*ServiceUsersAPIService)(&apiClient.common)
	apiClient.TeamsAPI = (*TeamsAPIService)(&apiClient.common)
	apiClient.UsersAPI = (*UsersAPIService)(&apiClient.common)
	apiClient.SecretsAPI = (*SecretsAPIService)(&apiClient.common)
	apiClient.DatabaseResourcesAPI = (*DatabaseResourcesAPIService)(&apiClient.common)
	apiClient.AccessReportsAPI = (*ReportsAPIService)(&apiClient.common)
	apiClient.CloudEntiltlementsAPI = (*CloudEntitlementsAPIService)(&apiClient.common)
	return apiClient, nil
}

func newRestyClient(apiClient *APIClient) *resty.Client {
	restyClient := resty.New()

	//Set Base Settings
	restyClient.SetBaseURL(apiClient.cfg.Host).
		SetHeader("User-Agent", apiClient.cfg.UserAgent).
		SetTimeout(time.Duration(apiClient.cfg.RequestTimeout)).
		SetDebug(apiClient.cfg.EnableHTTPDebug).
		SetError(&APIError{}) //Automatic unmarshalling if response status code is greater than 399

	if apiClient.cfg.EnableHTTPTrace {
		restyClient.EnableTrace()
	}

	if apiClient.cfg.TLSRootCAs != nil {
		restyClient.SetTLSClientConfig(&tls.Config{RootCAs: apiClient.cfg.TLSRootCAs})
	}

	if apiClient.cfg.TLSInsecureSkipVerify {
		restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}

	if apiClient.cfg.ProxyURL != "" {
		restyClient.SetProxy(apiClient.cfg.ProxyURL)
	}

	if apiClient.cfg.RoundTripper != nil {
		restyClient.SetTransport(apiClient.cfg.RoundTripper)
	}

	//Set Rate Limit and retry condition
	restyClient.
		OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
			// authentication logic
			switch authMode := apiClient.cfg.AuthorizationMode; authMode {
			case APIKey:
				// Get the token from the cache
				token, err := apiClient.tokenCache.GetToken()
				if err != nil {
					return err
				}
				// Add the token to the request header
				request.SetHeader("Authorization", "Bearer "+token)
			case BearerToken:
				request.SetHeader("Authorization", "Bearer "+apiClient.cfg.BearerToken)
			case OAuth2TokenSource:
				t, err := apiClient.cfg.OAuth2TokenSource.Token()
				if err != nil {
					return err
				}
				request.SetHeader("Authorization", t.TokenType+" "+t.AccessToken)
			default:
				return ErrAuthzModeInvalid
			}

			if apiClient.cfg.EnablePrintCurl {
				// Print cURL Request
				fullURL, err := url.JoinPath(client.BaseURL, request.URL)
				if err != nil {
					return err
				}
				if err := printcURL(fullURL, request); err != nil {
					return err
				}
			}

			if apiClient.cfg.UserAgentHook != nil {
				userAgent := apiClient.cfg.UserAgentHook()
				request.SetHeader("User-Agent", userAgent)
			}
			return nil
		}).
		// Rate limit condition
		AddRetryCondition(
			func(r *resty.Response, e error) bool {
				return r != nil && r.StatusCode() == http.StatusTooManyRequests
			}).
		SetRetryMaxWaitTime(time.Duration(apiClient.cfg.RateLimit.MaxBackoff) * time.Second).
		SetRetryCount(apiClient.cfg.RateLimit.MaxRetries).
		SetRetryAfter(func(c *resty.Client, r *resty.Response) (time.Duration, error) {
			headers := r.Header()
			retryAtHeader := headers.Get("X-RateLimit-Retry-At")
			if retryAtHeader == "" {
				// if we don't get a header, retry between 5-10 seconds from now
				waitTime := time.Millisecond * time.Duration(rand.Intn(5000)+5000)
				log.Printf("Request was rate limited, waiting %s to retry again\n", waitTime)
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
			log.Printf("Request was rate limited, waiting %s to retry again\n", waitTime)

			return waitTime, nil
		})

	return restyClient
}

type authTokenCache struct {
	mux          sync.Mutex
	serviceToken *AuthTokenResponse
	config       *configuration
}

func (a *authTokenCache) SetToken(token AuthTokenResponse) {
	a.mux.Lock()
	defer a.mux.Unlock()
	a.serviceToken = &token
}

func NewAuthTokenCache(config *configuration) *authTokenCache {
	return &authTokenCache{
		config: config,
	}
}

func (a *authTokenCache) GetToken() (string, error) {
	a.mux.Lock()
	defer a.mux.Unlock()

	jitter := time.Duration(rand.Int63n(int64(2 * time.Minute)))
	if a.serviceToken == nil || time.Now().After(a.serviceToken.ExpiresAt.Add(-1*jitter)) {
		// Token is expired or close to expire refresh it
		newToken, err := createServiceToken(a.config)
		if err != nil {
			return "", err
		}

		a.serviceToken = newToken
	}

	return a.serviceToken.BearerToken, nil
}

func createServiceToken(config *configuration) (*AuthTokenResponse, error) {
	authorizationURL := fmt.Sprintf("%s/v1/teams/%s/service_token", config.Host, url.PathEscape(config.TeamName))
	client := resty.New().SetBaseURL(config.Host).SetHeader("User-Agent", config.UserAgent)

	resp, err := client.R().
		SetBody(map[string]any{"key_id": config.APIKey, "key_secret": config.APISecret}).
		SetHeaders(map[string]string{"Accept": "application/json", "Content-Type": "application/json"}).
		SetResult(&AuthTokenResponse{}).
		Post(authorizationURL)

	if err != nil {
		return nil, err
	}
	if resp.StatusCode() == http.StatusUnauthorized {
		return nil, fmt.Errorf("received a 401 from URL %s when requesting service token.  check credentials and try again", authorizationURL)
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("call to %s resulted in status of %d, expected a %d.\nResponse Body: %s", resp.Request.URL, resp.StatusCode(), http.StatusOK, string(resp.Body()))
	}

	authToken := resp.Result().(*AuthTokenResponse)
	if !strings.EqualFold(authToken.TeamName, config.TeamName) {
		return nil, fmt.Errorf("service token team name does not match supplied team name, expected: %s, got: %s", config.TeamName, authToken.TeamName)
	}

	return authToken, nil
}

func (apiClient *APIClient) getTeamVaultJWKS() (*jose.JSONWebKeySet, error) {
	jwksResponse, resp, err := apiClient.TeamsAPI.GetVaultJWKS(context.Background(), apiClient.cfg.TeamName).Execute()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("call to %s resulted in status of %d, expected a %d", resp.Request.URL, resp.StatusCode, http.StatusOK)
	}

	jwks := &jose.JSONWebKeySet{}
	// marshal jwkResponse to bytes to convert type from *GetVaultJWKSResponse to *jose.JSONWebKeySet
	jwksBytes, err := json.Marshal(jwksResponse)
	if err != nil {
		return nil, fmt.Errorf("could not marshal JWKS response: %w", err)
	}

	if err := json.Unmarshal(jwksBytes, &jwks); err != nil {
		return nil, fmt.Errorf("could not unmarshal JWKS response to *jose.JSONWebKeySet: %w", err)
	}
	return jwks, nil
}

func (apiClient *APIClient) GetHTTPDefaultTransport() http.RoundTripper {
	return apiClient.restyClient.GetClient().Transport
}

// SetRoundTripper There is a config option for setting round tripper while creating client. We also need one to set after the client is created.
func (apiClient *APIClient) SetRoundTripper(rt http.RoundTripper) {
	apiClient.cfg.RoundTripper = rt
}

type formFile struct {
	fileBytes    []byte
	fileName     string
	formFileName string
}

func (apiClient *APIClient) callAPI(
	ctx context.Context,
	traceKey string,
	path string,
	method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
	formFiles []formFile,
	result interface{},
) (apiResponse *http.Response, err error) {

	if apiClient.cfg.EnableTelemetry && traceKey != "" {
		var sp opentracing.Span
		sp, ctx = opentracing.StartSpanFromContext(ctx, traceKey)
		defer sp.Finish()
	}

	// Using SetDoNotParseResponse tell resty not to parse or close the response in http.Client.Do(requset)
	// This means you need to ensure that the body is read and closed properly to avoid resource leaks.
	req := apiClient.restyClient.R().
		SetContext(ctx).
		SetHeaders(headerParams).
		SetQueryParamsFromValues(queryParams).
		SetFormDataFromValues(formParams).
		SetDoNotParseResponse(true)
	req.Method = method

	for _, ff := range formFiles {
		fileReader := bytes.NewReader(ff.fileBytes)
		req.SetFileReader(ff.fileName, ff.formFileName, fileReader)
	}

	if method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch {
		req.SetBody(postBody)
	}

	response, err := req.Execute(method, path)

	if err != nil {
		return nil, fmt.Errorf("error executing the API request: %w", err)
	}

	// If the response has no content, return the response and nil error directly
	if response.RawResponse.StatusCode == http.StatusNoContent {
		return response.RawResponse, nil
	}

	respBody, err := io.ReadAll(response.RawResponse.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading the raw HTTP response body: %w", err)
	}

	// unmarshal the response body into the result interface
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("error unmarshalling the raw HTTP response body: %w", err)
	}

	_ = response.RawResponse.Body.Close()

	// duplicate the response body so we can return it for the caller to use
	bodyBuffer := bytes.NewBuffer(respBody)
	newRawResponse := &http.Response{
		StatusCode: response.RawResponse.StatusCode,
		Header:     response.RawResponse.Header.Clone(),
		Body:       io.NopCloser(bytes.NewBuffer(bodyBuffer.Bytes())),
	}

	// reset the original response body
	response.RawResponse.Body = io.NopCloser(bytes.NewBuffer(bodyBuffer.Bytes()))

	if response.IsError() {
		responseError := response.Error()
		if responseError != nil {
			if apiError, ok := responseError.(*APIError); ok {
				return newRawResponse, apiError
			}
		}

		return newRawResponse, &APIError{
			Message: response.Status(),
		}
	}

	return response.RawResponse, nil
}

// selectHeaderContentType select a content type from the available list.
func selectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// selectHeaderAccept join all accept types and return
func selectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// contains is a case insensitive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.EqualFold(a, needle) {
			return true
		}
	}
	return false
}

func parameterValueToString(obj interface{}, key string) string {
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		return fmt.Sprintf("%v", obj)
	}
	var param, ok = obj.(MappedNullable)
	if !ok {
		return ""
	}
	dataMap, err := param.ToMap()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%v", dataMap[key])
}

// parameterAddToHeaderOrQuery adds the provided object to the request header or url query
// supporting deep object syntax
func parameterAddToHeaderOrQuery(headerOrQueryParams interface{}, keyPrefix string, obj interface{}, collectionType string) {
	var v = reflect.ValueOf(obj)
	var value = ""
	if v == reflect.ValueOf(nil) {
		value = "null"
	} else {
		switch v.Kind() {
		case reflect.Invalid:
			value = "invalid"

		case reflect.Struct:
			if t, ok := obj.(MappedNullable); ok {
				dataMap, err := t.ToMap()
				if err != nil {
					return
				}
				parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefix, dataMap, collectionType)
				return
			}
			if t, ok := obj.(time.Time); ok {
				parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefix, t.Format(time.RFC3339), collectionType)
				return
			}
			value = v.Type().String() + " value"
		case reflect.Slice:
			var indValue = reflect.ValueOf(obj)
			if indValue == reflect.ValueOf(nil) {
				return
			}
			var lenIndValue = indValue.Len()
			for i := 0; i < lenIndValue; i++ {
				var arrayValue = indValue.Index(i)
				parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefix, arrayValue.Interface(), collectionType)
			}
			return

		case reflect.Map:
			var indValue = reflect.ValueOf(obj)
			if indValue == reflect.ValueOf(nil) {
				return
			}
			iter := indValue.MapRange()
			for iter.Next() {
				k, v := iter.Key(), iter.Value()
				parameterAddToHeaderOrQuery(headerOrQueryParams, fmt.Sprintf("%s[%s]", keyPrefix, k.String()), v.Interface(), collectionType)
			}
			return

		case reflect.Interface:
			fallthrough
		case reflect.Ptr:
			parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefix, v.Elem().Interface(), collectionType)
			return

		case reflect.Int, reflect.Int8, reflect.Int16,
			reflect.Int32, reflect.Int64:
			value = strconv.FormatInt(v.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16,
			reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			value = strconv.FormatUint(v.Uint(), 10)
		case reflect.Float32, reflect.Float64:
			value = strconv.FormatFloat(v.Float(), 'g', -1, 32)
		case reflect.Bool:
			value = strconv.FormatBool(v.Bool())
		case reflect.String:
			value = v.String()
		default:
			value = v.Type().String() + " value"
		}
	}

	switch valuesMap := headerOrQueryParams.(type) {
	case url.Values:
		if collectionType == "csv" && valuesMap.Get(keyPrefix) != "" {
			valuesMap.Set(keyPrefix, valuesMap.Get(keyPrefix)+","+value)
		} else {
			valuesMap.Add(keyPrefix, value)
		}
	case map[string]string:
		valuesMap[keyPrefix] = value
	}
}

// helper for converting interface{} parameters to json strings
func parameterToJson(obj interface{}) (string, error) {
	jsonBuf, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(jsonBuf), err
}

// Prevent trying to import "fmt"
func reportError(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}

// A wrapper for strict JSON decoding
func newStrictDecoder(data []byte) *json.Decoder {
	dec := json.NewDecoder(bytes.NewBuffer(data))
	dec.DisallowUnknownFields()
	return dec
}

// Ripped from https://github.com/gregjones/httpcache/blob/master/httpcache.go
type cacheControl map[string]string

func parseCacheControl(headers http.Header) cacheControl {
	cc := cacheControl{}
	ccHeader := headers.Get("Cache-Control")
	for _, part := range strings.Split(ccHeader, ",") {
		part = strings.Trim(part, " ")
		if part == "" {
			continue
		}
		if strings.ContainsRune(part, '=') {
			keyval := strings.Split(part, "=")
			cc[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], ",")
		} else {
			cc[part] = ""
		}
	}
	return cc
}

// CacheExpires helper function to determine remaining time before repeating a request.
func CacheExpires(r *http.Response) time.Time {
	// Figure out when the cache expires.
	var expires time.Time
	now, err := time.Parse(time.RFC1123, r.Header.Get("date"))
	if err != nil {
		return time.Now()
	}
	respCacheControl := parseCacheControl(r.Header)

	if maxAge, ok := respCacheControl["max-age"]; ok {
		lifetime, err := time.ParseDuration(maxAge + "s")
		if err != nil {
			expires = now
		} else {
			expires = now.Add(lifetime)
		}
	} else {
		expiresHeader := r.Header.Get("Expires")
		if expiresHeader != "" {
			expires, err = time.Parse(time.RFC1123, expiresHeader)
			if err != nil {
				expires = now
			}
		}
	}
	return expires
}

// GetHTTPClient returns the underlying http.Client, useful mainly for tests. Please don't use this to circumvent
// the SDK.
func (apiClient *APIClient) GetHTTPClient() *http.Client {
	return apiClient.restyClient.GetClient()
}

// APIError Provides access to the body, error and model on returned errors.
type APIError struct {
	Message string `json:"message"`
}

// Error returns non-empty string if there was an error.
func (e APIError) Error() string {
	return e.Message
}

// Encrypt encrypts the provided string using the team's Vault JWKS and then returns the resulting *EncryptedString for use in API calls
func (apiClient *APIClient) Encrypt(data string) (*EncryptedString, error) {
	jwksResponse, err := apiClient.getTeamVaultJWKS()
	if err != nil {
		return nil, err
	}
	// Currently we only support jose.A256GCM content encryption, will add support for other algorithms in the future
	if jwksResponse != nil && len(jwksResponse.Keys) > 0 {
		encryptor := JWKEncryptor{PublicKeys: jwksResponse.Keys}

		encryptedData, err := encryptor.Encrypt([]byte(data), jose.A256GCM)
		if err != nil {
			return nil, err
		}
		return &EncryptedString{Payload: encryptedData}, nil
	}
	return nil, nil
}
