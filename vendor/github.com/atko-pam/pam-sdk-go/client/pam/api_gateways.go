/*
Okta Privileged Access

The OPA API is a control plane used to request operations in Okta Privileged Access (formerly ScaleFT/Advanced Server Access)

API version: 1.0.0
Contact: support@okta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pam

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// GatewaysAPIService GatewaysAPI service
type GatewaysAPIService service

type ApiCreateGatewaySetupTokenRequest struct {
	ctx               context.Context
	ApiService        *GatewaysAPIService
	teamName          string
	gatewaySetupToken *GatewaySetupToken
}

func (r ApiCreateGatewaySetupTokenRequest) GatewaySetupToken(gatewaySetupToken GatewaySetupToken) ApiCreateGatewaySetupTokenRequest {
	r.gatewaySetupToken = &gatewaySetupToken
	return r
}

func (r ApiCreateGatewaySetupTokenRequest) Execute() (*GatewaySetupToken, *http.Response, error) {
	return r.ApiService.CreateGatewaySetupTokenExecute(r)
}

/*
	CreateGatewaySetupToken Create a Gateway Setup Token

	    Creates a Gateway Setup Token for your Team

This endpoint requires the following role: `resource_admin`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    @param teamName The name of your Team
	@return ApiCreateGatewaySetupTokenRequest
*/
func (a *GatewaysAPIService) CreateGatewaySetupToken(ctx context.Context, teamName string) ApiCreateGatewaySetupTokenRequest {
	return ApiCreateGatewaySetupTokenRequest{
		ApiService: a,
		ctx:        ctx,
		teamName:   teamName,
	}
}

// Execute executes the request
//
//	@return GatewaySetupToken
func (a *GatewaysAPIService) CreateGatewaySetupTokenExecute(r ApiCreateGatewaySetupTokenRequest) (*GatewaySetupToken, *http.Response, error) {
	var (
		traceKey            = "gatewaysapi.createGatewaySetupToken"
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GatewaySetupToken
	)

	localVarPath := "/v1/teams/{team_name}/gateway_setup_tokens"
	localVarPath = strings.Replace(localVarPath, "{"+"team_name"+"}", url.PathEscape(parameterValueToString(r.teamName, "teamName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.gatewaySetupToken == nil {
		return localVarReturnValue, nil, reportError("gatewaySetupToken is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.gatewaySetupToken
	localVarHTTPResponse, err := a.client.callAPI(r.ctx, traceKey, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles, &localVarReturnValue)

	if err != nil {
		if localVarHTTPResponse == nil {
			return localVarReturnValue, nil, err
		}

		// read and unmarshal error response into right struct
		bodyBytes, err := io.ReadAll(localVarHTTPResponse.Body)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		if err := localVarHTTPResponse.Body.Close(); err != nil {
			return localVarReturnValue, nil, err
		}
		localVarHTTPResponse.Body = io.NopCloser(bytes.NewReader(bodyBytes)) //Reset body for the caller
		var apiError APIError
		if err := json.Unmarshal(bodyBytes, &apiError); err != nil {
			return localVarReturnValue, localVarHTTPResponse, err
		}
		return localVarReturnValue, localVarHTTPResponse, apiError
	}

	return localVarReturnValue, localVarHTTPResponse, err
}

type ApiDeleteGatewayInstanceRequest struct {
	ctx        context.Context
	ApiService *GatewaysAPIService
	teamName   string
	gatewayId  string
}

func (r ApiDeleteGatewayInstanceRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteGatewayInstanceExecute(r)
}

/*
	DeleteGatewayInstance Delete a Gateway

	    Deletes the specified Gateway from your Team

This endpoint requires the following role: `resource_admin`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    @param teamName The name of your Team
	    @param gatewayId The UUID of a Gateway
	@return ApiDeleteGatewayInstanceRequest
*/
func (a *GatewaysAPIService) DeleteGatewayInstance(ctx context.Context, teamName string, gatewayId string) ApiDeleteGatewayInstanceRequest {
	return ApiDeleteGatewayInstanceRequest{
		ApiService: a,
		ctx:        ctx,
		teamName:   teamName,
		gatewayId:  gatewayId,
	}
}

// Execute executes the request
func (a *GatewaysAPIService) DeleteGatewayInstanceExecute(r ApiDeleteGatewayInstanceRequest) (*http.Response, error) {
	var (
		traceKey           = "gatewaysapi.deleteGatewayInstance"
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localVarPath := "/v1/teams/{team_name}/gateways/{gateway_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"team_name"+"}", url.PathEscape(parameterValueToString(r.teamName, "teamName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"gateway_id"+"}", url.PathEscape(parameterValueToString(r.gatewayId, "gatewayId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHTTPResponse, err := a.client.callAPI(r.ctx, traceKey, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles, nil)

	if err != nil {
		if localVarHTTPResponse == nil {
			return nil, err
		}

		// read and unmarshal error response into right struct
		bodyBytes, err := io.ReadAll(localVarHTTPResponse.Body)
		if err != nil {
			return nil, err
		}
		if err := localVarHTTPResponse.Body.Close(); err != nil {
			return nil, err
		}
		localVarHTTPResponse.Body = io.NopCloser(bytes.NewReader(bodyBytes)) //Reset body for the caller
		var apiError APIError
		if err := json.Unmarshal(bodyBytes, &apiError); err != nil {
			return localVarHTTPResponse, err
		}
		return localVarHTTPResponse, apiError
	}

	return localVarHTTPResponse, err
}

type ApiDeleteGatewaySetupTokenRequest struct {
	ctx                 context.Context
	ApiService          *GatewaysAPIService
	teamName            string
	gatewaySetupTokenId string
}

func (r ApiDeleteGatewaySetupTokenRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteGatewaySetupTokenExecute(r)
}

/*
	DeleteGatewaySetupToken Delete a Gateway Setup Token

	    Deletes the specified Gateway Setup Token

This endpoint requires the following role: `resource_admin`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    @param teamName The name of your Team
	    @param gatewaySetupTokenId The UUID of a Gateway Setup Token
	@return ApiDeleteGatewaySetupTokenRequest
*/
func (a *GatewaysAPIService) DeleteGatewaySetupToken(ctx context.Context, teamName string, gatewaySetupTokenId string) ApiDeleteGatewaySetupTokenRequest {
	return ApiDeleteGatewaySetupTokenRequest{
		ApiService:          a,
		ctx:                 ctx,
		teamName:            teamName,
		gatewaySetupTokenId: gatewaySetupTokenId,
	}
}

// Execute executes the request
func (a *GatewaysAPIService) DeleteGatewaySetupTokenExecute(r ApiDeleteGatewaySetupTokenRequest) (*http.Response, error) {
	var (
		traceKey           = "gatewaysapi.deleteGatewaySetupToken"
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localVarPath := "/v1/teams/{team_name}/gateway_setup_tokens/{gateway_setup_token_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"team_name"+"}", url.PathEscape(parameterValueToString(r.teamName, "teamName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"gateway_setup_token_id"+"}", url.PathEscape(parameterValueToString(r.gatewaySetupTokenId, "gatewaySetupTokenId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHTTPResponse, err := a.client.callAPI(r.ctx, traceKey, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles, nil)

	if err != nil {
		if localVarHTTPResponse == nil {
			return nil, err
		}

		// read and unmarshal error response into right struct
		bodyBytes, err := io.ReadAll(localVarHTTPResponse.Body)
		if err != nil {
			return nil, err
		}
		if err := localVarHTTPResponse.Body.Close(); err != nil {
			return nil, err
		}
		localVarHTTPResponse.Body = io.NopCloser(bytes.NewReader(bodyBytes)) //Reset body for the caller
		var apiError APIError
		if err := json.Unmarshal(bodyBytes, &apiError); err != nil {
			return localVarHTTPResponse, err
		}
		return localVarHTTPResponse, apiError
	}

	return localVarHTTPResponse, err
}

type ApiFetchGatewaySetupTokenRequest struct {
	ctx                 context.Context
	ApiService          *GatewaysAPIService
	teamName            string
	gatewaySetupTokenId string
}

func (r ApiFetchGatewaySetupTokenRequest) Execute() (*GatewaySetupToken, *http.Response, error) {
	return r.ApiService.FetchGatewaySetupTokenExecute(r)
}

/*
	FetchGatewaySetupToken Retrieve a Gateway Setup Token

	    Retrieves the specified Gateway Setup Token

This endpoint requires the following role: `resource_admin`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    @param teamName The name of your Team
	    @param gatewaySetupTokenId The UUID of a Gateway Setup Token
	@return ApiFetchGatewaySetupTokenRequest
*/
func (a *GatewaysAPIService) FetchGatewaySetupToken(ctx context.Context, teamName string, gatewaySetupTokenId string) ApiFetchGatewaySetupTokenRequest {
	return ApiFetchGatewaySetupTokenRequest{
		ApiService:          a,
		ctx:                 ctx,
		teamName:            teamName,
		gatewaySetupTokenId: gatewaySetupTokenId,
	}
}

// Execute executes the request
//
//	@return GatewaySetupToken
func (a *GatewaysAPIService) FetchGatewaySetupTokenExecute(r ApiFetchGatewaySetupTokenRequest) (*GatewaySetupToken, *http.Response, error) {
	var (
		traceKey            = "gatewaysapi.fetchGatewaySetupToken"
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GatewaySetupToken
	)

	localVarPath := "/v1/teams/{team_name}/gateway_setup_tokens/{gateway_setup_token_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"team_name"+"}", url.PathEscape(parameterValueToString(r.teamName, "teamName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"gateway_setup_token_id"+"}", url.PathEscape(parameterValueToString(r.gatewaySetupTokenId, "gatewaySetupTokenId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHTTPResponse, err := a.client.callAPI(r.ctx, traceKey, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles, &localVarReturnValue)

	if err != nil {
		if localVarHTTPResponse == nil {
			return localVarReturnValue, nil, err
		}

		// read and unmarshal error response into right struct
		bodyBytes, err := io.ReadAll(localVarHTTPResponse.Body)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		if err := localVarHTTPResponse.Body.Close(); err != nil {
			return localVarReturnValue, nil, err
		}
		localVarHTTPResponse.Body = io.NopCloser(bytes.NewReader(bodyBytes)) //Reset body for the caller
		var apiError APIError
		if err := json.Unmarshal(bodyBytes, &apiError); err != nil {
			return localVarReturnValue, localVarHTTPResponse, err
		}
		return localVarReturnValue, localVarHTTPResponse, apiError
	}

	return localVarReturnValue, localVarHTTPResponse, err
}

type ApiFetchGatewaySetupTokenAsAgentTokenRequest struct {
	ctx                 context.Context
	ApiService          *GatewaysAPIService
	teamName            string
	gatewaySetupTokenId string
}

func (r ApiFetchGatewaySetupTokenAsAgentTokenRequest) Execute() (*GatewaySetupToken, *http.Response, error) {
	return r.ApiService.FetchGatewaySetupTokenAsAgentTokenExecute(r)
}

/*
	FetchGatewaySetupTokenAsAgentToken Retrieve a Gateway Setup Token

	    Retrieves the specified Gateway Setup Token

This endpoint requires the following role: `resource_admin`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    @param teamName The name of your Team
	    @param gatewaySetupTokenId The UUID of a Gateway Setup Token
	@return ApiFetchGatewaySetupTokenAsAgentTokenRequest
*/
func (a *GatewaysAPIService) FetchGatewaySetupTokenAsAgentToken(ctx context.Context, teamName string, gatewaySetupTokenId string) ApiFetchGatewaySetupTokenAsAgentTokenRequest {
	return ApiFetchGatewaySetupTokenAsAgentTokenRequest{
		ApiService:          a,
		ctx:                 ctx,
		teamName:            teamName,
		gatewaySetupTokenId: gatewaySetupTokenId,
	}
}

// Execute executes the request
//
//	@return GatewaySetupToken
func (a *GatewaysAPIService) FetchGatewaySetupTokenAsAgentTokenExecute(r ApiFetchGatewaySetupTokenAsAgentTokenRequest) (*GatewaySetupToken, *http.Response, error) {
	var (
		traceKey            = "gatewaysapi.fetchGatewaySetupTokenAsAgentToken"
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GatewaySetupToken
	)

	localVarPath := "/v1/teams/{team_name}/gateway_setup_tokens/{gateway_setup_token_id}/token"
	localVarPath = strings.Replace(localVarPath, "{"+"team_name"+"}", url.PathEscape(parameterValueToString(r.teamName, "teamName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"gateway_setup_token_id"+"}", url.PathEscape(parameterValueToString(r.gatewaySetupTokenId, "gatewaySetupTokenId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHTTPResponse, err := a.client.callAPI(r.ctx, traceKey, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles, &localVarReturnValue)

	if err != nil {
		if localVarHTTPResponse == nil {
			return localVarReturnValue, nil, err
		}

		// read and unmarshal error response into right struct
		bodyBytes, err := io.ReadAll(localVarHTTPResponse.Body)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		if err := localVarHTTPResponse.Body.Close(); err != nil {
			return localVarReturnValue, nil, err
		}
		localVarHTTPResponse.Body = io.NopCloser(bytes.NewReader(bodyBytes)) //Reset body for the caller
		var apiError APIError
		if err := json.Unmarshal(bodyBytes, &apiError); err != nil {
			return localVarReturnValue, localVarHTTPResponse, err
		}
		return localVarReturnValue, localVarHTTPResponse, apiError
	}

	return localVarReturnValue, localVarHTTPResponse, err
}

type ApiGetGatewayInstanceRequest struct {
	ctx        context.Context
	ApiService *GatewaysAPIService
	teamName   string
	gatewayId  string
}

func (r ApiGetGatewayInstanceRequest) Execute() (*GatewayAgent, *http.Response, error) {
	return r.ApiService.GetGatewayInstanceExecute(r)
}

/*
	GetGatewayInstance Retrieve a Gateway

	    Retrieves the properties of a specified Gateway

This endpoint requires one of the following roles: `resource_admin`, `delegated_resource_admin`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    @param teamName The name of your Team
	    @param gatewayId The UUID of a Gateway
	@return ApiGetGatewayInstanceRequest
*/
func (a *GatewaysAPIService) GetGatewayInstance(ctx context.Context, teamName string, gatewayId string) ApiGetGatewayInstanceRequest {
	return ApiGetGatewayInstanceRequest{
		ApiService: a,
		ctx:        ctx,
		teamName:   teamName,
		gatewayId:  gatewayId,
	}
}

// Execute executes the request
//
//	@return GatewayAgent
func (a *GatewaysAPIService) GetGatewayInstanceExecute(r ApiGetGatewayInstanceRequest) (*GatewayAgent, *http.Response, error) {
	var (
		traceKey            = "gatewaysapi.getGatewayInstance"
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GatewayAgent
	)

	localVarPath := "/v1/teams/{team_name}/gateways/{gateway_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"team_name"+"}", url.PathEscape(parameterValueToString(r.teamName, "teamName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"gateway_id"+"}", url.PathEscape(parameterValueToString(r.gatewayId, "gatewayId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHTTPResponse, err := a.client.callAPI(r.ctx, traceKey, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles, &localVarReturnValue)

	if err != nil {
		if localVarHTTPResponse == nil {
			return localVarReturnValue, nil, err
		}

		// read and unmarshal error response into right struct
		bodyBytes, err := io.ReadAll(localVarHTTPResponse.Body)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		if err := localVarHTTPResponse.Body.Close(); err != nil {
			return localVarReturnValue, nil, err
		}
		localVarHTTPResponse.Body = io.NopCloser(bytes.NewReader(bodyBytes)) //Reset body for the caller
		var apiError APIError
		if err := json.Unmarshal(bodyBytes, &apiError); err != nil {
			return localVarReturnValue, localVarHTTPResponse, err
		}
		return localVarReturnValue, localVarHTTPResponse, apiError
	}

	return localVarReturnValue, localVarHTTPResponse, err
}

type ApiGetGatewayStatusReportRequest struct {
	ctx        context.Context
	ApiService *GatewaysAPIService
	teamName   string
	gatewayId  string
}

func (r ApiGetGatewayStatusReportRequest) Execute() (*GatewayStatusReport, *http.Response, error) {
	return r.ApiService.GetGatewayStatusReportExecute(r)
}

/*
	GetGatewayStatusReport Retrieve a status report for a Gateway

	    Retrieve a status report for the specified Gateway

This endpoint requires one of the following roles: `resource_admin`, `delegated_resource_admin`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    @param teamName The name of your Team
	    @param gatewayId The UUID of a Gateway
	@return ApiGetGatewayStatusReportRequest
*/
func (a *GatewaysAPIService) GetGatewayStatusReport(ctx context.Context, teamName string, gatewayId string) ApiGetGatewayStatusReportRequest {
	return ApiGetGatewayStatusReportRequest{
		ApiService: a,
		ctx:        ctx,
		teamName:   teamName,
		gatewayId:  gatewayId,
	}
}

// Execute executes the request
//
//	@return GatewayStatusReport
func (a *GatewaysAPIService) GetGatewayStatusReportExecute(r ApiGetGatewayStatusReportRequest) (*GatewayStatusReport, *http.Response, error) {
	var (
		traceKey            = "gatewaysapi.getGatewayStatusReport"
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GatewayStatusReport
	)

	localVarPath := "/v1/teams/{team_name}/gateways/{gateway_id}/status"
	localVarPath = strings.Replace(localVarPath, "{"+"team_name"+"}", url.PathEscape(parameterValueToString(r.teamName, "teamName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"gateway_id"+"}", url.PathEscape(parameterValueToString(r.gatewayId, "gatewayId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHTTPResponse, err := a.client.callAPI(r.ctx, traceKey, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles, &localVarReturnValue)

	if err != nil {
		if localVarHTTPResponse == nil {
			return localVarReturnValue, nil, err
		}

		// read and unmarshal error response into right struct
		bodyBytes, err := io.ReadAll(localVarHTTPResponse.Body)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		if err := localVarHTTPResponse.Body.Close(); err != nil {
			return localVarReturnValue, nil, err
		}
		localVarHTTPResponse.Body = io.NopCloser(bytes.NewReader(bodyBytes)) //Reset body for the caller
		var apiError APIError
		if err := json.Unmarshal(bodyBytes, &apiError); err != nil {
			return localVarReturnValue, localVarHTTPResponse, err
		}
		return localVarReturnValue, localVarHTTPResponse, apiError
	}

	return localVarReturnValue, localVarHTTPResponse, err
}

type ApiListGatewaySetupTokensRequest struct {
	ctx        context.Context
	ApiService *GatewaysAPIService
	teamName   string
	count      *int32
	descending *bool
	offset     *string
	prev       *bool
}

// The number of objects per page
func (r ApiListGatewaySetupTokensRequest) Count(count int32) ApiListGatewaySetupTokensRequest {
	r.count = &count
	return r
}

// The object order
func (r ApiListGatewaySetupTokensRequest) Descending(descending bool) ApiListGatewaySetupTokensRequest {
	r.descending = &descending
	return r
}

// The offset value for pagination. The **rel&#x3D;\&quot;next\&quot;** and **rel&#x3D;\&quot;prev\&quot;** &#x60;Link&#x60; headers define the offset for subsequent or previous pages.
func (r ApiListGatewaySetupTokensRequest) Offset(offset string) ApiListGatewaySetupTokensRequest {
	r.offset = &offset
	return r
}

// The direction of paging
func (r ApiListGatewaySetupTokensRequest) Prev(prev bool) ApiListGatewaySetupTokensRequest {
	r.prev = &prev
	return r
}

func (r ApiListGatewaySetupTokensRequest) Execute() (*ListGatewaySetupTokensResponse, *http.Response, error) {
	return r.ApiService.ListGatewaySetupTokensExecute(r)
}

/*
	ListGatewaySetupTokens List all Gateway Setup Tokens

	    Lists all Gateway Setup Tokens for your Team

This endpoint requires the following role: `resource_admin`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    @param teamName The name of your Team
	@return ApiListGatewaySetupTokensRequest
*/
func (a *GatewaysAPIService) ListGatewaySetupTokens(ctx context.Context, teamName string) ApiListGatewaySetupTokensRequest {
	return ApiListGatewaySetupTokensRequest{
		ApiService: a,
		ctx:        ctx,
		teamName:   teamName,
	}
}

// Execute executes the request
//
//	@return ListGatewaySetupTokensResponse
func (a *GatewaysAPIService) ListGatewaySetupTokensExecute(r ApiListGatewaySetupTokensRequest) (*ListGatewaySetupTokensResponse, *http.Response, error) {
	var (
		traceKey            = "gatewaysapi.listGatewaySetupTokens"
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListGatewaySetupTokensResponse
	)

	localVarPath := "/v1/teams/{team_name}/gateway_setup_tokens"
	localVarPath = strings.Replace(localVarPath, "{"+"team_name"+"}", url.PathEscape(parameterValueToString(r.teamName, "teamName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "")
	}
	if r.descending != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "descending", r.descending, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.prev != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "prev", r.prev, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHTTPResponse, err := a.client.callAPI(r.ctx, traceKey, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles, &localVarReturnValue)

	if err != nil {
		if localVarHTTPResponse == nil {
			return localVarReturnValue, nil, err
		}

		// read and unmarshal error response into right struct
		bodyBytes, err := io.ReadAll(localVarHTTPResponse.Body)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		if err := localVarHTTPResponse.Body.Close(); err != nil {
			return localVarReturnValue, nil, err
		}
		localVarHTTPResponse.Body = io.NopCloser(bytes.NewReader(bodyBytes)) //Reset body for the caller
		var apiError APIError
		if err := json.Unmarshal(bodyBytes, &apiError); err != nil {
			return localVarReturnValue, localVarHTTPResponse, err
		}
		return localVarReturnValue, localVarHTTPResponse, apiError
	}

	return localVarReturnValue, localVarHTTPResponse, err
}

type ApiListGatewaysRequest struct {
	ctx        context.Context
	ApiService *GatewaysAPIService
	teamName   string
	contains   *string
	count      *int32
	descending *bool
	offset     *string
	prev       *bool
}

// Only return results that include the specified value
func (r ApiListGatewaysRequest) Contains(contains string) ApiListGatewaysRequest {
	r.contains = &contains
	return r
}

// The number of objects per page
func (r ApiListGatewaysRequest) Count(count int32) ApiListGatewaysRequest {
	r.count = &count
	return r
}

// The object order
func (r ApiListGatewaysRequest) Descending(descending bool) ApiListGatewaysRequest {
	r.descending = &descending
	return r
}

// The offset value for pagination. The **rel&#x3D;\&quot;next\&quot;** and **rel&#x3D;\&quot;prev\&quot;** &#x60;Link&#x60; headers define the offset for subsequent or previous pages.
func (r ApiListGatewaysRequest) Offset(offset string) ApiListGatewaysRequest {
	r.offset = &offset
	return r
}

// The direction of paging
func (r ApiListGatewaysRequest) Prev(prev bool) ApiListGatewaysRequest {
	r.prev = &prev
	return r
}

func (r ApiListGatewaysRequest) Execute() (*ListGatewaysResponse, *http.Response, error) {
	return r.ApiService.ListGatewaysExecute(r)
}

/*
	ListGateways List all Gateways

	    Lists all Gateways for your Team

This endpoint requires one of the following roles: `resource_admin`, `delegated_resource_admin`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    @param teamName The name of your Team
	@return ApiListGatewaysRequest
*/
func (a *GatewaysAPIService) ListGateways(ctx context.Context, teamName string) ApiListGatewaysRequest {
	return ApiListGatewaysRequest{
		ApiService: a,
		ctx:        ctx,
		teamName:   teamName,
	}
}

// Execute executes the request
//
//	@return ListGatewaysResponse
func (a *GatewaysAPIService) ListGatewaysExecute(r ApiListGatewaysRequest) (*ListGatewaysResponse, *http.Response, error) {
	var (
		traceKey            = "gatewaysapi.listGateways"
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListGatewaysResponse
	)

	localVarPath := "/v1/teams/{team_name}/gateways"
	localVarPath = strings.Replace(localVarPath, "{"+"team_name"+"}", url.PathEscape(parameterValueToString(r.teamName, "teamName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.contains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "contains", r.contains, "")
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "")
	}
	if r.descending != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "descending", r.descending, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.prev != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "prev", r.prev, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHTTPResponse, err := a.client.callAPI(r.ctx, traceKey, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles, &localVarReturnValue)

	if err != nil {
		if localVarHTTPResponse == nil {
			return localVarReturnValue, nil, err
		}

		// read and unmarshal error response into right struct
		bodyBytes, err := io.ReadAll(localVarHTTPResponse.Body)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		if err := localVarHTTPResponse.Body.Close(); err != nil {
			return localVarReturnValue, nil, err
		}
		localVarHTTPResponse.Body = io.NopCloser(bytes.NewReader(bodyBytes)) //Reset body for the caller
		var apiError APIError
		if err := json.Unmarshal(bodyBytes, &apiError); err != nil {
			return localVarReturnValue, localVarHTTPResponse, err
		}
		return localVarReturnValue, localVarHTTPResponse, apiError
	}

	return localVarReturnValue, localVarHTTPResponse, err
}

type ApiUpdateGatewayInstanceRequest struct {
	ctx          context.Context
	ApiService   *GatewaysAPIService
	teamName     string
	gatewayId    string
	gatewayAgent *GatewayAgent
}

func (r ApiUpdateGatewayInstanceRequest) GatewayAgent(gatewayAgent GatewayAgent) ApiUpdateGatewayInstanceRequest {
	r.gatewayAgent = &gatewayAgent
	return r
}

func (r ApiUpdateGatewayInstanceRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateGatewayInstanceExecute(r)
}

/*
	UpdateGatewayInstance Update a Gateway

	    Updates a specified Gateway

This endpoint requires the following role: `resource_admin`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    @param teamName The name of your Team
	    @param gatewayId The UUID of a Gateway
	@return ApiUpdateGatewayInstanceRequest
*/
func (a *GatewaysAPIService) UpdateGatewayInstance(ctx context.Context, teamName string, gatewayId string) ApiUpdateGatewayInstanceRequest {
	return ApiUpdateGatewayInstanceRequest{
		ApiService: a,
		ctx:        ctx,
		teamName:   teamName,
		gatewayId:  gatewayId,
	}
}

// Execute executes the request
func (a *GatewaysAPIService) UpdateGatewayInstanceExecute(r ApiUpdateGatewayInstanceRequest) (*http.Response, error) {
	var (
		traceKey           = "gatewaysapi.updateGatewayInstance"
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localVarPath := "/v1/teams/{team_name}/gateways/{gateway_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"team_name"+"}", url.PathEscape(parameterValueToString(r.teamName, "teamName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"gateway_id"+"}", url.PathEscape(parameterValueToString(r.gatewayId, "gatewayId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.gatewayAgent == nil {
		return nil, reportError("gatewayAgent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.gatewayAgent
	localVarHTTPResponse, err := a.client.callAPI(r.ctx, traceKey, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles, nil)

	if err != nil {
		if localVarHTTPResponse == nil {
			return nil, err
		}

		// read and unmarshal error response into right struct
		bodyBytes, err := io.ReadAll(localVarHTTPResponse.Body)
		if err != nil {
			return nil, err
		}
		if err := localVarHTTPResponse.Body.Close(); err != nil {
			return nil, err
		}
		localVarHTTPResponse.Body = io.NopCloser(bytes.NewReader(bodyBytes)) //Reset body for the caller
		var apiError APIError
		if err := json.Unmarshal(bodyBytes, &apiError); err != nil {
			return localVarHTTPResponse, err
		}
		return localVarHTTPResponse, apiError
	}

	return localVarHTTPResponse, err
}
