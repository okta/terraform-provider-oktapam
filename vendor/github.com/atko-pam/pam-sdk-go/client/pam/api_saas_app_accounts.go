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

// SaasAppAccountsAPIService SaasAppAccountsAPI service
type SaasAppAccountsAPIService service

type ApiGetResourceGroupProjectSaasAppAccountRequest struct {
	ctx              context.Context
	ApiService       *SaasAppAccountsAPIService
	teamName         string
	resourceGroupId  string
	projectId        string
	saasAppAccountId string
}

func (r ApiGetResourceGroupProjectSaasAppAccountRequest) Execute() (*SaasAppAccount, *http.Response, error) {
	return r.ApiService.GetResourceGroupProjectSaasAppAccountExecute(r)
}

/*
	GetResourceGroupProjectSaasAppAccount Retrieve a SaaS Application Account from a Project

	    Retrieves a SaaS Application Account from a Project in a Resource Group.

This endpoint requires one of the following roles: `resource_admin`, `delegated_resource_admin`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    @param teamName The name of your Team
	    @param resourceGroupId The UUID of a Resource Group
	    @param projectId The UUID of a Project
	    @param saasAppAccountId The UUID of a SaaS Application Account
	@return ApiGetResourceGroupProjectSaasAppAccountRequest
*/
func (a *SaasAppAccountsAPIService) GetResourceGroupProjectSaasAppAccount(ctx context.Context, teamName string, resourceGroupId string, projectId string, saasAppAccountId string) ApiGetResourceGroupProjectSaasAppAccountRequest {
	return ApiGetResourceGroupProjectSaasAppAccountRequest{
		ApiService:       a,
		ctx:              ctx,
		teamName:         teamName,
		resourceGroupId:  resourceGroupId,
		projectId:        projectId,
		saasAppAccountId: saasAppAccountId,
	}
}

// Execute executes the request
//
//	@return SaasAppAccount
func (a *SaasAppAccountsAPIService) GetResourceGroupProjectSaasAppAccountExecute(r ApiGetResourceGroupProjectSaasAppAccountRequest) (*SaasAppAccount, *http.Response, error) {
	var (
		traceKey            = "saasappaccountsapi.getResourceGroupProjectSaasAppAccount"
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SaasAppAccount
	)

	localVarPath := "/v1/teams/{team_name}/resource_groups/{resource_group_id}/projects/{project_id}/saas_app_accounts/{saas_app_account_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"team_name"+"}", url.PathEscape(parameterValueToString(r.teamName, "teamName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resource_group_id"+"}", url.PathEscape(parameterValueToString(r.resourceGroupId, "resourceGroupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project_id"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"saas_app_account_id"+"}", url.PathEscape(parameterValueToString(r.saasAppAccountId, "saasAppAccountId")), -1)

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

type ApiListResourceGroupProjectSaasAppAccountsRequest struct {
	ctx             context.Context
	ApiService      *SaasAppAccountsAPIService
	teamName        string
	resourceGroupId string
	projectId       string
}

func (r ApiListResourceGroupProjectSaasAppAccountsRequest) Execute() (*ListResourceGroupProjectSaasAppAccountsResponse, *http.Response, error) {
	return r.ApiService.ListResourceGroupProjectSaasAppAccountsExecute(r)
}

/*
	ListResourceGroupProjectSaasAppAccounts List all SaaS Application Accounts in a Project

	    Lists all SaaS Application Accounts in a Project in a Resource Group.

This endpoint requires one of the following roles: `resource_admin`, `delegated_resource_admin`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    @param teamName The name of your Team
	    @param resourceGroupId The UUID of a Resource Group
	    @param projectId The UUID of a Project
	@return ApiListResourceGroupProjectSaasAppAccountsRequest
*/
func (a *SaasAppAccountsAPIService) ListResourceGroupProjectSaasAppAccounts(ctx context.Context, teamName string, resourceGroupId string, projectId string) ApiListResourceGroupProjectSaasAppAccountsRequest {
	return ApiListResourceGroupProjectSaasAppAccountsRequest{
		ApiService:      a,
		ctx:             ctx,
		teamName:        teamName,
		resourceGroupId: resourceGroupId,
		projectId:       projectId,
	}
}

// Execute executes the request
//
//	@return ListResourceGroupProjectSaasAppAccountsResponse
func (a *SaasAppAccountsAPIService) ListResourceGroupProjectSaasAppAccountsExecute(r ApiListResourceGroupProjectSaasAppAccountsRequest) (*ListResourceGroupProjectSaasAppAccountsResponse, *http.Response, error) {
	var (
		traceKey            = "saasappaccountsapi.listResourceGroupProjectSaasAppAccounts"
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListResourceGroupProjectSaasAppAccountsResponse
	)

	localVarPath := "/v1/teams/{team_name}/resource_groups/{resource_group_id}/projects/{project_id}/saas_app_accounts"
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

type ApiRotateResourceGroupProjectSaasAppAccountRequest struct {
	ctx              context.Context
	ApiService       *SaasAppAccountsAPIService
	teamName         string
	resourceGroupId  string
	projectId        string
	saasAppAccountId string
}

func (r ApiRotateResourceGroupProjectSaasAppAccountRequest) Execute() (*http.Response, error) {
	return r.ApiService.RotateResourceGroupProjectSaasAppAccountExecute(r)
}

/*
	RotateResourceGroupProjectSaasAppAccount Rotate the password belonging to a SaaS Application Account in a Project

	    Rotates the password belonging to a SaaS Application Account in a Project in a Resource Group.

If the SaaS Application does not support password rotation, returns a `405` Method Not Allowed response.

This endpoint requires one of the following roles: `resource_admin`, `delegated_resource_admin`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    @param teamName The name of your Team
	    @param resourceGroupId The UUID of a Resource Group
	    @param projectId The UUID of a Project
	    @param saasAppAccountId The UUID of a SaaS Application Account
	@return ApiRotateResourceGroupProjectSaasAppAccountRequest
*/
func (a *SaasAppAccountsAPIService) RotateResourceGroupProjectSaasAppAccount(ctx context.Context, teamName string, resourceGroupId string, projectId string, saasAppAccountId string) ApiRotateResourceGroupProjectSaasAppAccountRequest {
	return ApiRotateResourceGroupProjectSaasAppAccountRequest{
		ApiService:       a,
		ctx:              ctx,
		teamName:         teamName,
		resourceGroupId:  resourceGroupId,
		projectId:        projectId,
		saasAppAccountId: saasAppAccountId,
	}
}

// Execute executes the request
func (a *SaasAppAccountsAPIService) RotateResourceGroupProjectSaasAppAccountExecute(r ApiRotateResourceGroupProjectSaasAppAccountRequest) (*http.Response, error) {
	var (
		traceKey           = "saasappaccountsapi.rotateResourceGroupProjectSaasAppAccount"
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localVarPath := "/v1/teams/{team_name}/resource_groups/{resource_group_id}/projects/{project_id}/saas_app_accounts/{saas_app_account_id}/rotate"
	localVarPath = strings.Replace(localVarPath, "{"+"team_name"+"}", url.PathEscape(parameterValueToString(r.teamName, "teamName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resource_group_id"+"}", url.PathEscape(parameterValueToString(r.resourceGroupId, "resourceGroupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project_id"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"saas_app_account_id"+"}", url.PathEscape(parameterValueToString(r.saasAppAccountId, "saasAppAccountId")), -1)

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
