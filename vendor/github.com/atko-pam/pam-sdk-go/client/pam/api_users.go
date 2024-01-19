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

// UsersAPIService UsersAPI service
type UsersAPIService service

type ApiGetCurrentUserInfoRequest struct {
	ctx        context.Context
	ApiService *UsersAPIService
	teamName   string
}

func (r ApiGetCurrentUserInfoRequest) Execute() (*CurrentUserInfo, *http.Response, error) {
	return r.ApiService.GetCurrentUserInfoExecute(r)
}

/*
GetCurrentUserInfo Retrieve User details

	Retrieves details about the current User

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

	@param teamName The name of your Team

@return ApiGetCurrentUserInfoRequest
*/
func (a *UsersAPIService) GetCurrentUserInfo(ctx context.Context, teamName string) ApiGetCurrentUserInfoRequest {
	return ApiGetCurrentUserInfoRequest{
		ApiService: a,
		ctx:        ctx,
		teamName:   teamName,
	}
}

// Execute executes the request
//
//	@return CurrentUserInfo
func (a *UsersAPIService) GetCurrentUserInfoExecute(r ApiGetCurrentUserInfoRequest) (*CurrentUserInfo, *http.Response, error) {
	var (
		traceKey            = "usersapi.getCurrentUserInfo"
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CurrentUserInfo
	)

	localVarPath := "/v1/teams/{team_name}/current_user"
	localVarPath = strings.Replace(localVarPath, "{"+"team_name"+"}", url.PathEscape(parameterValueToString(r.teamName, "teamName")), -1)

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

type ApiGetUserRequest struct {
	ctx        context.Context
	ApiService *UsersAPIService
	teamName   string
	userName   string
}

func (r ApiGetUserRequest) Execute() (*User, *http.Response, error) {
	return r.ApiService.GetUserExecute(r)
}

/*
	GetUser Retrieve a User

	    Retrieves a User from your Team

This endpoint requires one of the following roles: `pam_admin`, `resource_admin`, `security_admin`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    @param teamName The name of your Team
	    @param userName The username for an existing User
	@return ApiGetUserRequest
*/
func (a *UsersAPIService) GetUser(ctx context.Context, teamName string, userName string) ApiGetUserRequest {
	return ApiGetUserRequest{
		ApiService: a,
		ctx:        ctx,
		teamName:   teamName,
		userName:   userName,
	}
}

// Execute executes the request
//
//	@return User
func (a *UsersAPIService) GetUserExecute(r ApiGetUserRequest) (*User, *http.Response, error) {
	var (
		traceKey            = "usersapi.getUser"
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *User
	)

	localVarPath := "/v1/teams/{team_name}/users/{user_name}"
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

type ApiListUserGroupsRequest struct {
	ctx                context.Context
	ApiService         *UsersAPIService
	teamName           string
	userName           string
	contains           *string
	count              *int32
	descending         *bool
	id                 *string
	ignore             *string
	includeDeleted     *bool
	offset             *string
	onlyIncludeDeleted *bool
	prev               *bool
}

// Only return results that include the specified value
func (r ApiListUserGroupsRequest) Contains(contains string) ApiListUserGroupsRequest {
	r.contains = &contains
	return r
}

// The number of objects per page
func (r ApiListUserGroupsRequest) Count(count int32) ApiListUserGroupsRequest {
	r.count = &count
	return r
}

// The object order
func (r ApiListUserGroupsRequest) Descending(descending bool) ApiListUserGroupsRequest {
	r.descending = &descending
	return r
}

// Only return results with the specified IDs
func (r ApiListUserGroupsRequest) Id(id string) ApiListUserGroupsRequest {
	r.id = &id
	return r
}

// Ignore Groups with the specified names. This is case sensitive.
func (r ApiListUserGroupsRequest) Ignore(ignore string) ApiListUserGroupsRequest {
	r.ignore = &ignore
	return r
}

// If &#x60;true&#x60;, include deleted Groups in the results
func (r ApiListUserGroupsRequest) IncludeDeleted(includeDeleted bool) ApiListUserGroupsRequest {
	r.includeDeleted = &includeDeleted
	return r
}

// The offset value for pagination. The **rel&#x3D;\&quot;next\&quot;** and **rel&#x3D;\&quot;prev\&quot;** &#x60;Link&#x60; headers define the offset for subsequent or previous pages.
func (r ApiListUserGroupsRequest) Offset(offset string) ApiListUserGroupsRequest {
	r.offset = &offset
	return r
}

// If &#x60;true&#x60;, only return deleted Groups in the results
func (r ApiListUserGroupsRequest) OnlyIncludeDeleted(onlyIncludeDeleted bool) ApiListUserGroupsRequest {
	r.onlyIncludeDeleted = &onlyIncludeDeleted
	return r
}

// The direction of paging
func (r ApiListUserGroupsRequest) Prev(prev bool) ApiListUserGroupsRequest {
	r.prev = &prev
	return r
}

func (r ApiListUserGroupsRequest) Execute() (*ListUserGroupsResponse, *http.Response, error) {
	return r.ApiService.ListUserGroupsExecute(r)
}

/*
	ListUserGroups List all Groups for a User

	    Lists all Groups for a specified User

This endpoint requires one of the following roles: `pam_admin`, `resource_admin`, `security_admin`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    @param teamName The name of your Team
	    @param userName The username for an existing User
	@return ApiListUserGroupsRequest
*/
func (a *UsersAPIService) ListUserGroups(ctx context.Context, teamName string, userName string) ApiListUserGroupsRequest {
	return ApiListUserGroupsRequest{
		ApiService: a,
		ctx:        ctx,
		teamName:   teamName,
		userName:   userName,
	}
}

// Execute executes the request
//
//	@return ListUserGroupsResponse
func (a *UsersAPIService) ListUserGroupsExecute(r ApiListUserGroupsRequest) (*ListUserGroupsResponse, *http.Response, error) {
	var (
		traceKey            = "usersapi.listUserGroups"
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListUserGroupsResponse
	)

	localVarPath := "/v1/teams/{team_name}/users/{user_name}/groups"
	localVarPath = strings.Replace(localVarPath, "{"+"team_name"+"}", url.PathEscape(parameterValueToString(r.teamName, "teamName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"user_name"+"}", url.PathEscape(parameterValueToString(r.userName, "userName")), -1)

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
	if r.ignore != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ignore", r.ignore, "")
	}
	if r.includeDeleted != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include_deleted", r.includeDeleted, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.onlyIncludeDeleted != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "only_include_deleted", r.onlyIncludeDeleted, "")
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

	if localVarHTTPResponse == nil && err != nil {
		return localVarReturnValue, nil, err
	}

	return localVarReturnValue, localVarHTTPResponse, err
}

type ApiListUsersRequest struct {
	ctx                 context.Context
	ApiService          *UsersAPIService
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
func (r ApiListUsersRequest) Contains(contains string) ApiListUsersRequest {
	r.contains = &contains
	return r
}

// The number of objects per page
func (r ApiListUsersRequest) Count(count int32) ApiListUsersRequest {
	r.count = &count
	return r
}

// The object order
func (r ApiListUsersRequest) Descending(descending bool) ApiListUsersRequest {
	r.descending = &descending
	return r
}

// Only return results with the specified IDs
func (r ApiListUsersRequest) Id(id string) ApiListUsersRequest {
	r.id = &id
	return r
}

// Only return Service Users in the results
func (r ApiListUsersRequest) IncludeServiceUsers(includeServiceUsers string) ApiListUsersRequest {
	r.includeServiceUsers = &includeServiceUsers
	return r
}

// The offset value for pagination. The **rel&#x3D;\&quot;next\&quot;** and **rel&#x3D;\&quot;prev\&quot;** &#x60;Link&#x60; headers define the offset for subsequent or previous pages.
func (r ApiListUsersRequest) Offset(offset string) ApiListUsersRequest {
	r.offset = &offset
	return r
}

// The direction of paging
func (r ApiListUsersRequest) Prev(prev bool) ApiListUsersRequest {
	r.prev = &prev
	return r
}

// Only return Users with a name that begins with the specified value
func (r ApiListUsersRequest) StartsWith(startsWith string) ApiListUsersRequest {
	r.startsWith = &startsWith
	return r
}

// Only return Users with the specified status. Valid statuses: &#x60;ACTIVE&#x60;, &#x60;DISABLED&#x60;, and &#x60;DELETED&#x60;.
func (r ApiListUsersRequest) Status(status string) ApiListUsersRequest {
	r.status = &status
	return r
}

func (r ApiListUsersRequest) Execute() (*ListUsersResponse, *http.Response, error) {
	return r.ApiService.ListUsersExecute(r)
}

/*
	ListUsers List all Users for a Team

	    Lists all Users for your Team

This endpoint requires one of the following roles: `pam_admin`, `resource_admin`, `security_admin`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    @param teamName The name of your Team
	@return ApiListUsersRequest
*/
func (a *UsersAPIService) ListUsers(ctx context.Context, teamName string) ApiListUsersRequest {
	return ApiListUsersRequest{
		ApiService: a,
		ctx:        ctx,
		teamName:   teamName,
	}
}

// Execute executes the request
//
//	@return ListUsersResponse
func (a *UsersAPIService) ListUsersExecute(r ApiListUsersRequest) (*ListUsersResponse, *http.Response, error) {
	var (
		traceKey            = "usersapi.listUsers"
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListUsersResponse
	)

	localVarPath := "/v1/teams/{team_name}/users"
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
