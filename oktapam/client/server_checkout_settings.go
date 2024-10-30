package client

import (
	"context"
	"net/http"

	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
)

func (c *SDKClientWrapper) GetServerCheckoutSettings(ctx context.Context, resourceGroupID string, projectID string) (*pam.ResourceCheckoutSettings, error) {
	settings, httpResp, err := c.SDKClient.ProjectsAPI.FetchResourceGroupServerBasedProjectCheckoutSettings(ctx, c.Team, resourceGroupID, projectID).Execute()
	if err != nil {
		logging.Errorf("received error while fetching server checkout settings: %v", err)
		return nil, err
	}

	statusCode := httpResp.StatusCode
	if statusCode == http.StatusOK {
		return settings, nil
	} else if statusCode == http.StatusNotFound {
		return nil, nil
	}

	return nil, createErrorForInvalidCodeFromSDK(httpResp, http.StatusOK, http.StatusNotFound)
}

func (c *SDKClientWrapper) UpdateServerCheckoutSettings(ctx context.Context, resourceGroupID string, projectID string, serverCheckoutSettings *pam.ResourceCheckoutSettings) error {

	httpResp, err := c.SDKClient.ProjectsAPI.UpdateResourceGroupServerBasedProjectCheckoutSettings(ctx, c.Team, resourceGroupID, projectID).ResourceCheckoutSettings(*serverCheckoutSettings).Execute()
	if err != nil {
		logging.Errorf("received error while updating server checkout settings: %v", err)
		return err
	}

	_, err = checkStatusCodeFromSDK(httpResp, http.StatusNoContent)
	return err
}
