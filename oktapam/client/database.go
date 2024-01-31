package client

import (
	"context"
	"github.com/atko-pam/pam-sdk-go/client/pam"
	"net/http"
)

func GetDatabase(ctx context.Context, sdkClient SDKClientWrapper, resourceGroupID, projectID, databaseID string) (*pam.DatabaseResourceResponse, error) {
	request := sdkClient.SDKClient.DatabaseResourcesAPI.GetDatabaseResource(ctx, sdkClient.Team, resourceGroupID, projectID, databaseID)
	resp, httpResp, callErr := request.Execute()
	if httpResp != nil {
		if _, err := checkStatusCodeFromSDK(httpResp, http.StatusOK); err != nil {
			return nil, err
		}
	} else if callErr != nil {
		return nil, callErr
	}

	return resp, nil
}

func CreateDatabase(ctx context.Context, sdkClient SDKClientWrapper, resourceGroupID, projectID, canonicalName, databaseType, managementConnectionDetailsType string, managementConnectionDetails pam.ManagementConnectionDetails, selector *map[string]string) (*pam.DatabaseResourceResponse, error) {
	payload := pam.NewDatabaseResourceCreateOrUpdateRequest(canonicalName, pam.DatabaseType(databaseType), pam.ManagementConnectionDetailsType(managementConnectionDetailsType), managementConnectionDetails)
	payload.ManagementGatewaySelector = selector
	request := sdkClient.SDKClient.DatabaseResourcesAPI.
		CreateDatabaseResource(ctx, sdkClient.Team, resourceGroupID, projectID).
		DatabaseResourceCreateOrUpdateRequest(*payload)

	resp, httpResp, callErr := request.Execute()
	if httpResp != nil {
		if _, err := checkStatusCodeFromSDK(httpResp, http.StatusCreated); err != nil {
			return nil, err
		}
	} else if callErr != nil {
		return nil, callErr
	}

	return resp, nil
}

func UpdateDatabase(ctx context.Context, sdkClient SDKClientWrapper, dbID, resourceGroupID, projectID, canonicalName, databaseType, managementConnectionDetailsType string, managementConnectionDetails pam.ManagementConnectionDetails, selector *map[string]string) error {
	payload := pam.NewDatabaseResourceCreateOrUpdateRequest(canonicalName, pam.DatabaseType(databaseType), pam.ManagementConnectionDetailsType(managementConnectionDetailsType), managementConnectionDetails)
	payload.ManagementGatewaySelector = selector
	request := sdkClient.SDKClient.DatabaseResourcesAPI.
		UpdateDatabaseResource(ctx, sdkClient.Team, resourceGroupID, projectID, dbID).
		DatabaseResourceCreateOrUpdateRequest(*payload)

	httpResp, callErr := request.Execute()
	if httpResp != nil {
		if _, err := checkStatusCodeFromSDK(httpResp, http.StatusNoContent); err != nil {
			return err
		}
	} else if callErr != nil {
		return callErr
	}

	return nil
}

func DeleteDatabase(ctx context.Context, sdkClient SDKClientWrapper, resourceGroupID, projectID, databaseID string) error {
	request := sdkClient.SDKClient.DatabaseResourcesAPI.DeleteDatabaseResource(ctx, sdkClient.Team, resourceGroupID, projectID, databaseID)
	httpResp, callErr := request.Execute()
	if httpResp != nil {
		if _, err := checkStatusCodeFromSDK(httpResp, http.StatusNoContent, http.StatusNotFound); err != nil {
			return err
		}
	} else if callErr != nil {
		return callErr
	}

	return nil
}
