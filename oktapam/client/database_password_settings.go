package client

import (
	"context"
	"github.com/atko-pam/pam-sdk-go/client/pam"
	"net/http"
)

func GetDatabasePasswordSettings(ctx context.Context, sdkClient SDKClientWrapper, resourceGroupID string, projectID string) (*pam.PasswordPolicy, error) {
	request := sdkClient.SDKClient.ProjectsAPI.GetProjectPasswordPolicyForDatabaseResources(ctx, sdkClient.Team, resourceGroupID, projectID)
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

func UpdateDatabasePasswordSettings(ctx context.Context, sdkClient SDKClientWrapper, resourceGroupID string, projectID string, passwordPolicy pam.PasswordPolicy) error {
	request := sdkClient.SDKClient.ProjectsAPI.
		UpdateProjectPasswordPolicyForDatabaseResources(ctx, sdkClient.Team, resourceGroupID, projectID).
		PasswordPolicy(passwordPolicy)
	_, httpResp, callErr := request.Execute()
	if httpResp != nil {
		if _, err := checkStatusCodeFromSDK(httpResp, http.StatusOK); err != nil {
			return err
		}
	} else if callErr != nil {
		return callErr
	}

	return nil
}
