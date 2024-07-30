package client

import (
	"context"
	"net/http"

	"github.com/atko-pam/pam-sdk-go/client/pam"
)

func GetServerCheckoutSettings(ctx context.Context, sdkClient SDKClientWrapper, resourceGroupID string, projectID string) (*pam.ResourceCheckoutSettings, error) {
	request := sdkClient.SDKClient.ProjectsAPI.FetchResourceGroupServerBasedProjectCheckoutSettings(ctx, sdkClient.Team, resourceGroupID, projectID)
	resp, httpResp, callErr := request.Execute()
	if httpResp != nil {
		if status, err := checkStatusCodeFromSDK(httpResp, http.StatusOK, http.StatusNotFound); err != nil {
			return nil, err
		} else if status == http.StatusNotFound {
			return nil, nil
		}
	} else if callErr != nil {
		return nil, callErr
	}

	return resp, nil
}

func UpdateServerCheckoutSettings(ctx context.Context, sdkClient SDKClientWrapper, resourceGroupID string, projectID string, serverCheckoutSettings pam.ResourceCheckoutSettings) error {
	request := sdkClient.SDKClient.ProjectsAPI.
		UpdateResourceGroupServerBasedProjectCheckoutSettings(ctx, sdkClient.Team, resourceGroupID, projectID).
		ResourceCheckoutSettings(serverCheckoutSettings)

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
