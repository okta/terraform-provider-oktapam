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

// OktaUdAccountsAPIService OktaUdAccountsAPI service
type OktaUdAccountsAPIService service

type ApiGetResourceGroupProjectOktaUDAccountRequest struct {
	ctx             context.Context
	ApiService      *OktaUdAccountsAPIService
	teamName        string
	resourceGroupId string
	projectId       string
	oktaUdAccountId string
}

func (r ApiGetResourceGroupProjectOktaUDAccountRequest) Execute() (*OktaUDAccount, *http.Response, error) {
	return r.ApiService.GetResourceGroupProjectOktaUDAccountExecute(r)
}

/*
	GetResourceGroupProjectOktaUDAccount Retrieve an Okta Universal Directory Account from a Project

	    Retrieves an Okta Universal Directory Account from a Project in a Resource Group.

This endpoint requires one of the following roles: `resource_admin`, `delegated_resource_admin`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    @param teamName The name of your Team
	    @param resourceGroupId The UUID of a Resource Group
	    @param projectId The UUID of a Project
	    @param oktaUdAccountId The UUID of an Okta Universal Directory Account
	@return ApiGetResourceGroupProjectOktaUDAccountRequest
*/
func (a *OktaUdAccountsAPIService) GetResourceGroupProjectOktaUDAccount(ctx context.Context, teamName string, resourceGroupId string, projectId string, oktaUdAccountId string) ApiGetResourceGroupProjectOktaUDAccountRequest {
	return ApiGetResourceGroupProjectOktaUDAccountRequest{
		ApiService:      a,
		ctx:             ctx,
		teamName:        teamName,
		resourceGroupId: resourceGroupId,
		projectId:       projectId,
		oktaUdAccountId: oktaUdAccountId,
	}
}

// Execute executes the request
//
//	@return OktaUDAccount
func (a *OktaUdAccountsAPIService) GetResourceGroupProjectOktaUDAccountExecute(r ApiGetResourceGroupProjectOktaUDAccountRequest) (*OktaUDAccount, *http.Response, error) {
	var (
		traceKey            = "oktaudaccountsapi.getResourceGroupProjectOktaUDAccount"
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OktaUDAccount
	)

	localVarPath := "/v1/teams/{team_name}/resource_groups/{resource_group_id}/projects/{project_id}/okta_ud_accounts/{okta_ud_account_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"team_name"+"}", url.PathEscape(parameterValueToString(r.teamName, "teamName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resource_group_id"+"}", url.PathEscape(parameterValueToString(r.resourceGroupId, "resourceGroupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project_id"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"okta_ud_account_id"+"}", url.PathEscape(parameterValueToString(r.oktaUdAccountId, "oktaUdAccountId")), -1)

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

type ApiListResourceGroupProjectOktaUDAccountsRequest struct {
	ctx             context.Context
	ApiService      *OktaUdAccountsAPIService
	teamName        string
	resourceGroupId string
	projectId       string
}

func (r ApiListResourceGroupProjectOktaUDAccountsRequest) Execute() (*ListResourceGroupProjectOktaUDAccountsResponse, *http.Response, error) {
	return r.ApiService.ListResourceGroupProjectOktaUDAccountsExecute(r)
}

/*
	ListResourceGroupProjectOktaUDAccounts List all Okta Universal Directory Accounts in a Project

	    Lists all Okta Universal Directory Accounts in a Project in a Resource Group.

This endpoint requires one of the following roles: `resource_admin`, `delegated_resource_admin`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    @param teamName The name of your Team
	    @param resourceGroupId The UUID of a Resource Group
	    @param projectId The UUID of a Project
	@return ApiListResourceGroupProjectOktaUDAccountsRequest
*/
func (a *OktaUdAccountsAPIService) ListResourceGroupProjectOktaUDAccounts(ctx context.Context, teamName string, resourceGroupId string, projectId string) ApiListResourceGroupProjectOktaUDAccountsRequest {
	return ApiListResourceGroupProjectOktaUDAccountsRequest{
		ApiService:      a,
		ctx:             ctx,
		teamName:        teamName,
		resourceGroupId: resourceGroupId,
		projectId:       projectId,
	}
}

// Execute executes the request
//
//	@return ListResourceGroupProjectOktaUDAccountsResponse
func (a *OktaUdAccountsAPIService) ListResourceGroupProjectOktaUDAccountsExecute(r ApiListResourceGroupProjectOktaUDAccountsRequest) (*ListResourceGroupProjectOktaUDAccountsResponse, *http.Response, error) {
	var (
		traceKey            = "oktaudaccountsapi.listResourceGroupProjectOktaUDAccounts"
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListResourceGroupProjectOktaUDAccountsResponse
	)

	localVarPath := "/v1/teams/{team_name}/resource_groups/{resource_group_id}/projects/{project_id}/okta_ud_accounts"
	localVarPath = strings.Replace(localVarPath, "{"+"team_name"+"}", url.PathEscape(parameterValueToString(r.teamName, "teamName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resource_group_id"+"}", url.PathEscape(parameterValueToString(r.resourceGroupId, "resourceGroupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project_id"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)

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

type ApiRotateResourceGroupProjectOktaUDAccountRequest struct {
	ctx             context.Context
	ApiService      *OktaUdAccountsAPIService
	teamName        string
	resourceGroupId string
	projectId       string
	oktaUdAccountId string
}

func (r ApiRotateResourceGroupProjectOktaUDAccountRequest) Execute() (*http.Response, error) {
	return r.ApiService.RotateResourceGroupProjectOktaUDAccountExecute(r)
}

/*
	RotateResourceGroupProjectOktaUDAccount Rotate the password belonging to an Okta Universal Directory Account in a Project

	    Rotates the password belonging to an Okta Universal Directory Account in a Project in a Resource Group.

This endpoint requires one of the following roles: `resource_admin`, `delegated_resource_admin`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    @param teamName The name of your Team
	    @param resourceGroupId The UUID of a Resource Group
	    @param projectId The UUID of a Project
	    @param oktaUdAccountId The UUID of an Okta Universal Directory Account
	@return ApiRotateResourceGroupProjectOktaUDAccountRequest
*/
func (a *OktaUdAccountsAPIService) RotateResourceGroupProjectOktaUDAccount(ctx context.Context, teamName string, resourceGroupId string, projectId string, oktaUdAccountId string) ApiRotateResourceGroupProjectOktaUDAccountRequest {
	return ApiRotateResourceGroupProjectOktaUDAccountRequest{
		ApiService:      a,
		ctx:             ctx,
		teamName:        teamName,
		resourceGroupId: resourceGroupId,
		projectId:       projectId,
		oktaUdAccountId: oktaUdAccountId,
	}
}

// Execute executes the request
func (a *OktaUdAccountsAPIService) RotateResourceGroupProjectOktaUDAccountExecute(r ApiRotateResourceGroupProjectOktaUDAccountRequest) (*http.Response, error) {
	var (
		traceKey           = "oktaudaccountsapi.rotateResourceGroupProjectOktaUDAccount"
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localVarPath := "/v1/teams/{team_name}/resource_groups/{resource_group_id}/projects/{project_id}/okta_ud_accounts/{okta_ud_account_id}/rotate"
	localVarPath = strings.Replace(localVarPath, "{"+"team_name"+"}", url.PathEscape(parameterValueToString(r.teamName, "teamName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resource_group_id"+"}", url.PathEscape(parameterValueToString(r.resourceGroupId, "resourceGroupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project_id"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"okta_ud_account_id"+"}", url.PathEscape(parameterValueToString(r.oktaUdAccountId, "oktaUdAccountId")), -1)

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
