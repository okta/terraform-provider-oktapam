/*
Okta Privileged Access

The OPA API is a control plane used to request operations in Okta Privileged Access (formerly ScaleFT/Advanced Server Access)

API version: 1.0.0
Contact: support@okta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pam

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

// ServiceUsersAPIService ServiceUsersAPI service
type ServiceUsersAPIService service

type ApiCreateServiceUserRequest struct {
	ctx                   context.Context
	ApiService            *ServiceUsersAPIService
	teamName              string
	createServiceUserBody *CreateServiceUserBody
}

func (r ApiCreateServiceUserRequest) CreateServiceUserBody(createServiceUserBody CreateServiceUserBody) ApiCreateServiceUserRequest {
	r.createServiceUserBody = &createServiceUserBody
	return r
}

func (r ApiCreateServiceUserRequest) Execute() (*User, *http.Response, error) {
	return r.ApiService.CreateServiceUserExecute(r)
}

/*
	CreateServiceUser Create a Service User

	    Creates a Service User that can be used to automate interactions with the OPA API

This endpoint requires the following role: `resource_admin`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    @param teamName The name of your Team
	@return ApiCreateServiceUserRequest
*/
func (a *ServiceUsersAPIService) CreateServiceUser(ctx context.Context, teamName string) ApiCreateServiceUserRequest {
	return ApiCreateServiceUserRequest{
		ApiService: a,
		ctx:        ctx,
		teamName:   teamName,
	}
}

// Execute executes the request
//
//	@return User
func (a *ServiceUsersAPIService) CreateServiceUserExecute(r ApiCreateServiceUserRequest) (*User, *http.Response, error) {
	var (
		traceKey            = "serviceusersapi.createServiceUser"
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *User
	)

	localVarPath := "/v1/teams/{team_name}/service_users"
	localVarPath = strings.Replace(localVarPath, "{"+"team_name"+"}", url.PathEscape(parameterValueToString(r.teamName, "teamName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createServiceUserBody == nil {
		return localVarReturnValue, nil, reportError("createServiceUserBody is required and must be specified")
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
	localVarPostBody = r.createServiceUserBody
	localVarHTTPResponse, err := a.client.callAPI(r.ctx, traceKey, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles, &localVarReturnValue)

	if localVarHTTPResponse == nil && err != nil {
		return localVarReturnValue, nil, err
	}

	return localVarReturnValue, localVarHTTPResponse, err
}

type ApiDeleteServiceUserKeyRequest struct {
	ctx        context.Context
	ApiService *ServiceUsersAPIService
	teamName   string
	userName   string
	keyId      string
}

func (r ApiDeleteServiceUserKeyRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteServiceUserKeyExecute(r)
}

/*
	DeleteServiceUserKey Delete an API key

	    Deletes an API key for a Service User. The Service User can no longer authenticate with this API key.

This endpoint requires the following role: `resource_admin`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    @param teamName The name of your Team
	    @param userName The username for an existing User
	    @param keyId The UUID of a Service User key
	@return ApiDeleteServiceUserKeyRequest
*/
func (a *ServiceUsersAPIService) DeleteServiceUserKey(ctx context.Context, teamName string, userName string, keyId string) ApiDeleteServiceUserKeyRequest {
	return ApiDeleteServiceUserKeyRequest{
		ApiService: a,
		ctx:        ctx,
		teamName:   teamName,
		userName:   userName,
		keyId:      keyId,
	}
}

// Execute executes the request
func (a *ServiceUsersAPIService) DeleteServiceUserKeyExecute(r ApiDeleteServiceUserKeyRequest) (*http.Response, error) {
	var (
		traceKey           = "serviceusersapi.deleteServiceUserKey"
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localVarPath := "/v1/teams/{team_name}/service_users/{user_name}/keys/{key_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"team_name"+"}", url.PathEscape(parameterValueToString(r.teamName, "teamName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"user_name"+"}", url.PathEscape(parameterValueToString(r.userName, "userName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"key_id"+"}", url.PathEscape(parameterValueToString(r.keyId, "keyId")), -1)

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

	if localVarHTTPResponse == nil && err != nil {
		return nil, err
	}

	return localVarHTTPResponse, err
}

type ApiGetServiceUserRequest struct {
	ctx        context.Context
	ApiService *ServiceUsersAPIService
	teamName   string
	userName   string
}

func (r ApiGetServiceUserRequest) Execute() (*User, *http.Response, error) {
	return r.ApiService.GetServiceUserExecute(r)
}

/*
	GetServiceUser Retrieve a Service User

	    Retrieve a specified Service User

This endpoint requires one of the following roles: `pam_admin`, `security_admin`, `resource_admin`, `delegated_resource_admin`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    @param teamName The name of your Team
	    @param userName The username for an existing User
	@return ApiGetServiceUserRequest
*/
func (a *ServiceUsersAPIService) GetServiceUser(ctx context.Context, teamName string, userName string) ApiGetServiceUserRequest {
	return ApiGetServiceUserRequest{
		ApiService: a,
		ctx:        ctx,
		teamName:   teamName,
		userName:   userName,
	}
}

// Execute executes the request
//
//	@return User
func (a *ServiceUsersAPIService) GetServiceUserExecute(r ApiGetServiceUserRequest) (*User, *http.Response, error) {
	var (
		traceKey            = "serviceusersapi.getServiceUser"
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *User
	)

	localVarPath := "/v1/teams/{team_name}/service_users/{user_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"team_name"+"}", url.PathEscape(parameterValueToString(r.teamName, "teamName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"user_name"+"}", url.PathEscape(parameterValueToString(r.userName, "userName")), -1)

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

	if localVarHTTPResponse == nil && err != nil {
		return localVarReturnValue, nil, err
	}

	return localVarReturnValue, localVarHTTPResponse, err
}

type ApiIssueServiceTokenRequest struct {
	ctx                          context.Context
	ApiService                   *ServiceUsersAPIService
	teamName                     string
	issueServiceTokenRequestBody *IssueServiceTokenRequestBody
}

func (r ApiIssueServiceTokenRequest) IssueServiceTokenRequestBody(issueServiceTokenRequestBody IssueServiceTokenRequestBody) ApiIssueServiceTokenRequest {
	r.issueServiceTokenRequestBody = &issueServiceTokenRequestBody
	return r
}

func (r ApiIssueServiceTokenRequest) Execute() (*AuthTokenResponse, *http.Response, error) {
	return r.ApiService.IssueServiceTokenExecute(r)
}

/*
	IssueServiceToken Issue a Service User token

	    Most calls to the OPA API require an HTTP `Authorization` header with a value of `Bearer ${AUTH_TOKEN}`.

To retrieve an auth token, you need to create an API key for a Service User and pass the API key information to this endpoint.
Auth tokens may expire at any time, so code that uses them should be prepared to handle a 401 response code by creating a new auth token.

This endpoint requires the following role: `resource_admin`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    @param teamName The name of your Team
	@return ApiIssueServiceTokenRequest
*/
func (a *ServiceUsersAPIService) IssueServiceToken(ctx context.Context, teamName string) ApiIssueServiceTokenRequest {
	return ApiIssueServiceTokenRequest{
		ApiService: a,
		ctx:        ctx,
		teamName:   teamName,
	}
}

// Execute executes the request
//
//	@return AuthTokenResponse
func (a *ServiceUsersAPIService) IssueServiceTokenExecute(r ApiIssueServiceTokenRequest) (*AuthTokenResponse, *http.Response, error) {
	var (
		traceKey            = "serviceusersapi.issueServiceToken"
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AuthTokenResponse
	)

	localVarPath := "/v1/teams/{team_name}/service_token"
	localVarPath = strings.Replace(localVarPath, "{"+"team_name"+"}", url.PathEscape(parameterValueToString(r.teamName, "teamName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.issueServiceTokenRequestBody == nil {
		return localVarReturnValue, nil, reportError("issueServiceTokenRequestBody is required and must be specified")
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
	localVarPostBody = r.issueServiceTokenRequestBody
	localVarHTTPResponse, err := a.client.callAPI(r.ctx, traceKey, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles, &localVarReturnValue)

	if localVarHTTPResponse == nil && err != nil {
		return localVarReturnValue, nil, err
	}

	return localVarReturnValue, localVarHTTPResponse, err
}

type ApiListServiceUserKeysRequest struct {
	ctx        context.Context
	ApiService *ServiceUsersAPIService
	teamName   string
	userName   string
}

func (r ApiListServiceUserKeysRequest) Execute() (*ListServiceUserKeysResponse, *http.Response, error) {
	return r.ApiService.ListServiceUserKeysExecute(r)
}

/*
	ListServiceUserKeys List all API keys

	    Lists all API keys for a specified Service User. This doesn't return the corresponding secret for each API key.

This endpoint requires the following role: `resource_admin`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    @param teamName The name of your Team
	    @param userName The username for an existing User
	@return ApiListServiceUserKeysRequest
*/
func (a *ServiceUsersAPIService) ListServiceUserKeys(ctx context.Context, teamName string, userName string) ApiListServiceUserKeysRequest {
	return ApiListServiceUserKeysRequest{
		ApiService: a,
		ctx:        ctx,
		teamName:   teamName,
		userName:   userName,
	}
}

// Execute executes the request
//
//	@return ListServiceUserKeysResponse
func (a *ServiceUsersAPIService) ListServiceUserKeysExecute(r ApiListServiceUserKeysRequest) (*ListServiceUserKeysResponse, *http.Response, error) {
	var (
		traceKey            = "serviceusersapi.listServiceUserKeys"
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListServiceUserKeysResponse
	)

	localVarPath := "/v1/teams/{team_name}/service_users/{user_name}/keys"
	localVarPath = strings.Replace(localVarPath, "{"+"team_name"+"}", url.PathEscape(parameterValueToString(r.teamName, "teamName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"user_name"+"}", url.PathEscape(parameterValueToString(r.userName, "userName")), -1)

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

	if localVarHTTPResponse == nil && err != nil {
		return localVarReturnValue, nil, err
	}

	return localVarReturnValue, localVarHTTPResponse, err
}

type ApiListServiceUsersRequest struct {
	ctx                 context.Context
	ApiService          *ServiceUsersAPIService
	teamName            string
	contains            *string
	count               *int32
	descending          *bool
	id                  *string
	includeServiceUsers *string
	offset              *string
	prev                *bool
	startsWith          *string
	status              *string
}

// Only return results that include the specified value
func (r ApiListServiceUsersRequest) Contains(contains string) ApiListServiceUsersRequest {
	r.contains = &contains
	return r
}

// The number of objects per page
func (r ApiListServiceUsersRequest) Count(count int32) ApiListServiceUsersRequest {
	r.count = &count
	return r
}

// The object order
func (r ApiListServiceUsersRequest) Descending(descending bool) ApiListServiceUsersRequest {
	r.descending = &descending
	return r
}

// Only return results with the specified IDs
func (r ApiListServiceUsersRequest) Id(id string) ApiListServiceUsersRequest {
	r.id = &id
	return r
}

// Only return Service Users in the results
func (r ApiListServiceUsersRequest) IncludeServiceUsers(includeServiceUsers string) ApiListServiceUsersRequest {
	r.includeServiceUsers = &includeServiceUsers
	return r
}

// The offset value for pagination. The **rel&#x3D;\&quot;next\&quot;** and **rel&#x3D;\&quot;prev\&quot;** &#x60;Link&#x60; headers define the offset for subsequent or previous pages.
func (r ApiListServiceUsersRequest) Offset(offset string) ApiListServiceUsersRequest {
	r.offset = &offset
	return r
}

// The direction of paging
func (r ApiListServiceUsersRequest) Prev(prev bool) ApiListServiceUsersRequest {
	r.prev = &prev
	return r
}

// Only return Users with a name that begins with the specified value
func (r ApiListServiceUsersRequest) StartsWith(startsWith string) ApiListServiceUsersRequest {
	r.startsWith = &startsWith
	return r
}

// Only return Users with the specified status. Valid statuses: &#x60;ACTIVE&#x60;, &#x60;DISABLED&#x60;, and &#x60;DELETED&#x60;.
func (r ApiListServiceUsersRequest) Status(status string) ApiListServiceUsersRequest {
	r.status = &status
	return r
}

func (r ApiListServiceUsersRequest) Execute() (*ListServiceUsersResponse, *http.Response, error) {
	return r.ApiService.ListServiceUsersExecute(r)
}

/*
	ListServiceUsers List all Service Users for a Team

	    Lists all Service Users for your Team

This endpoint requires one of the following roles: `pam_admin`, `security_admin`, `resource_admin`, `delegated_resource_admin`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    @param teamName The name of your Team
	@return ApiListServiceUsersRequest
*/
func (a *ServiceUsersAPIService) ListServiceUsers(ctx context.Context, teamName string) ApiListServiceUsersRequest {
	return ApiListServiceUsersRequest{
		ApiService: a,
		ctx:        ctx,
		teamName:   teamName,
	}
}

// Execute executes the request
//
//	@return ListServiceUsersResponse
func (a *ServiceUsersAPIService) ListServiceUsersExecute(r ApiListServiceUsersRequest) (*ListServiceUsersResponse, *http.Response, error) {
	var (
		traceKey            = "serviceusersapi.listServiceUsers"
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListServiceUsersResponse
	)

	localVarPath := "/v1/teams/{team_name}/service_users"
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
	if r.id != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "")
	}
	if r.includeServiceUsers != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include_service_users", r.includeServiceUsers, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.prev != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "prev", r.prev, "")
	}
	if r.startsWith != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "starts_with", r.startsWith, "")
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "")
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

	if localVarHTTPResponse == nil && err != nil {
		return localVarReturnValue, nil, err
	}

	return localVarReturnValue, localVarHTTPResponse, err
}

type ApiRotateServiceUserKeyRequest struct {
	ctx        context.Context
	ApiService *ServiceUsersAPIService
	teamName   string
	userName   string
}

func (r ApiRotateServiceUserKeyRequest) Execute() (*ServiceUserKeyWithSecret, *http.Response, error) {
	return r.ApiService.RotateServiceUserKeyExecute(r)
}

/*
	RotateServiceUserKey Rotate all API keys

	    Rotates all API keys for a specified Service User. This also sets an expiration date for the existing API keys.

This endpoint requires the following role: `resource_admin`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    @param teamName The name of your Team
	    @param userName The username for an existing User
	@return ApiRotateServiceUserKeyRequest
*/
func (a *ServiceUsersAPIService) RotateServiceUserKey(ctx context.Context, teamName string, userName string) ApiRotateServiceUserKeyRequest {
	return ApiRotateServiceUserKeyRequest{
		ApiService: a,
		ctx:        ctx,
		teamName:   teamName,
		userName:   userName,
	}
}

// Execute executes the request
//
//	@return ServiceUserKeyWithSecret
func (a *ServiceUsersAPIService) RotateServiceUserKeyExecute(r ApiRotateServiceUserKeyRequest) (*ServiceUserKeyWithSecret, *http.Response, error) {
	var (
		traceKey            = "serviceusersapi.rotateServiceUserKey"
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ServiceUserKeyWithSecret
	)

	localVarPath := "/v1/teams/{team_name}/service_users/{user_name}/keys"
	localVarPath = strings.Replace(localVarPath, "{"+"team_name"+"}", url.PathEscape(parameterValueToString(r.teamName, "teamName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"user_name"+"}", url.PathEscape(parameterValueToString(r.userName, "userName")), -1)

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

	if localVarHTTPResponse == nil && err != nil {
		return localVarReturnValue, nil, err
	}

	return localVarReturnValue, localVarHTTPResponse, err
}

type ApiServiceUserUpdateRequest struct {
	ctx                   context.Context
	ApiService            *ServiceUsersAPIService
	teamName              string
	userName              string
	updateServiceUserBody *UpdateServiceUserBody
}

func (r ApiServiceUserUpdateRequest) UpdateServiceUserBody(updateServiceUserBody UpdateServiceUserBody) ApiServiceUserUpdateRequest {
	r.updateServiceUserBody = &updateServiceUserBody
	return r
}

func (r ApiServiceUserUpdateRequest) Execute() (*User, *http.Response, error) {
	return r.ApiService.ServiceUserUpdateExecute(r)
}

/*
	ServiceUserUpdate Update a Service User

	    Updates a specified Service User

This endpoint requires the following role: `resource_admin`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    @param teamName The name of your Team
	    @param userName The username for an existing User
	@return ApiServiceUserUpdateRequest
*/
func (a *ServiceUsersAPIService) ServiceUserUpdate(ctx context.Context, teamName string, userName string) ApiServiceUserUpdateRequest {
	return ApiServiceUserUpdateRequest{
		ApiService: a,
		ctx:        ctx,
		teamName:   teamName,
		userName:   userName,
	}
}

// Execute executes the request
//
//	@return User
func (a *ServiceUsersAPIService) ServiceUserUpdateExecute(r ApiServiceUserUpdateRequest) (*User, *http.Response, error) {
	var (
		traceKey            = "serviceusersapi.serviceUserUpdate"
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *User
	)

	localVarPath := "/v1/teams/{team_name}/service_users/{user_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"team_name"+"}", url.PathEscape(parameterValueToString(r.teamName, "teamName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"user_name"+"}", url.PathEscape(parameterValueToString(r.userName, "userName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateServiceUserBody == nil {
		return localVarReturnValue, nil, reportError("updateServiceUserBody is required and must be specified")
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
	localVarPostBody = r.updateServiceUserBody
	localVarHTTPResponse, err := a.client.callAPI(r.ctx, traceKey, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles, &localVarReturnValue)

	if localVarHTTPResponse == nil && err != nil {
		return localVarReturnValue, nil, err
	}

	return localVarReturnValue, localVarHTTPResponse, err
}
