package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"
)

type CloudConnectionDetails struct {
	AccountId  *string `json:"account_id"`
	ExternalId *string `json:"external_id"`
	RoleArn    *string `json:"role_arn"`
}

type CloudConnection struct {
	DeletedAt              *string                 `json:"deleted_at,omitempty"`
	ID                     *string                 `json:"id,omitempty"`
	Name                   *string                 `json:"name"`
	Provider               *string                 `json:"provider"`
	CloudConnectionDetails *CloudConnectionDetails `json:"cloud_connection_details"`
}

func (c CloudConnection) Exists() bool {
	return utils.IsNonEmpty(c.ID) && utils.IsBlank(c.DeletedAt)
}

func (c CloudConnection) ToResourceMap() map[string]any {
	m := make(map[string]any, 10) // TODO: check that size again

	if c.Name != nil {
		m[attributes.Name] = *c.Name
	}
	if c.ID != nil {
		m[attributes.ID] = *c.ID
	}
	if c.DeletedAt != nil {
		m[attributes.DeletedAt] = *c.DeletedAt
	}
	m[attributes.CloudConnectionDetails] = c.CloudConnectionDetails

	return m
}

func (c OktaPAMClient) GetCloudConnection(ctx context.Context, id string, allowDeleted bool) (*CloudConnection, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/cloud_connections/%s", url.PathEscape(c.Team), url.PathEscape(id))
	logging.Tracef("making GET request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetResult(&CloudConnection{}).Get(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}
	statusCode := resp.StatusCode()

	if statusCode == http.StatusOK {
		cloudConnection := resp.Result().(*CloudConnection)
		if cloudConnection.Exists() || allowDeleted {
			return cloudConnection, nil
		}
		return nil, nil
	} else if statusCode == http.StatusNotFound {
		return nil, nil
	}

	return nil, createErrorForInvalidCode(resp, http.StatusOK, http.StatusNotFound)
}

func (c OktaPAMClient) CreateCloudConnection(ctx context.Context, cloudConnection CloudConnection) (*CloudConnection, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/cloud_connections", url.PathEscape(c.Team))
	logging.Tracef("making POST request to %s", requestURL)
	resultingCloudConnection := &CloudConnection{}
	resp, err := c.CreateBaseRequest(ctx).SetBody(cloudConnection).SetResult(resultingCloudConnection).Post(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}
	if _, err = checkStatusCode(resp, http.StatusCreated); err != nil {
		return nil, err
	}
	return resultingCloudConnection, nil
}

func (c OktaPAMClient) UpdateCloudConnection(ctx context.Context, cloudConnection CloudConnection) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/cloud_connections/%s", url.PathEscape(c.Team), url.PathEscape(*cloudConnection.ID))
	logging.Tracef("making PUT request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetBody(cloudConnection).Put(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}
	if _, err = checkStatusCode(resp, http.StatusNoContent); err != nil {
		return err
	}
	return nil
}

func (c OktaPAMClient) DeleteCloudConnection(ctx context.Context, cloudConnectionID string) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/cloud_connections/%s", url.PathEscape(c.Team), url.PathEscape(cloudConnectionID))
	logging.Tracef("making DELETE request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).Delete(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}

	if _, err = checkStatusCode(resp, http.StatusOK, http.StatusNoContent, http.StatusNotFound); err != nil {
		return err
	}

	return nil
}
