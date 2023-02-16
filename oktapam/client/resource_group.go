package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
)

type ResourceGroup struct {
	ID                           *string       `json:"id,omitempty"`
	Name                         *string       `json:"name"`
	TeamID                       *string       `json:"team_id,omitempty"`
	Description                  *string       `json:"description"`
	DelegatedResourceAdminGroups []NamedObject `json:"delegated_resource_admin_groups"`
}

func (rg ResourceGroup) ToResourceMap() map[string]any {
	m := make(map[string]any, 5)

	if rg.ID != nil {
		m[attributes.ID] = *rg.ID
	}
	if rg.Name != nil {
		m[attributes.Name] = *rg.Name
	}
	if rg.Description != nil {
		m[attributes.Description] = *rg.Description
	}
	if rg.DelegatedResourceAdminGroups != nil {
		groupIds := make([]any, len(rg.DelegatedResourceAdminGroups))
		for idx, groupId := range rg.DelegatedResourceAdminGroups {
			g := groupId
			groupIds[idx] = *g.Id
		}
		m[attributes.DelegatedResourceAdminGroups] = groupIds
	}

	return m
}

func (c OktaPAMClient) CreateResourceGroup(ctx context.Context, resourceGroup ResourceGroup) (*ResourceGroup, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/resource_groups", url.PathEscape(c.Team))
	logging.Tracef("making POST request to %s", requestURL)
	resultingResourceGroup := &ResourceGroup{}
	resp, err := c.CreateBaseRequest(ctx).SetBody(resourceGroup).SetResult(resultingResourceGroup).Post(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}
	if _, err = checkStatusCode(resp, http.StatusCreated); err != nil {
		return nil, err
	}
	return resultingResourceGroup, nil
}

func (c OktaPAMClient) GetResourceGroup(ctx context.Context, resourceGroupID string) (*ResourceGroup, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/resource_groups/%s", url.PathEscape(c.Team), url.PathEscape(resourceGroupID))
	logging.Tracef("making GET request to %s", requestURL)
	resourceGroup := &ResourceGroup{}
	resp, err := c.CreateBaseRequest(ctx).SetResult(resourceGroup).Get(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}
	statusCode := resp.StatusCode()
	if statusCode == http.StatusOK {
		return resourceGroup, nil
	} else if statusCode == http.StatusNotFound {
		return nil, nil
	}

	return nil, createErrorForInvalidCode(resp, http.StatusOK, http.StatusNotFound)
}

func (c OktaPAMClient) UpdateResourceGroup(ctx context.Context, resourceGroupID string, resourceGroup ResourceGroup) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/resource_groups/%s", url.PathEscape(c.Team), url.PathEscape(resourceGroupID))
	logging.Tracef("making PUT request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetBody(resourceGroup).Put(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}
	if _, err = checkStatusCode(resp, http.StatusNoContent); err != nil {
		return err
	}
	return nil
}

func (c OktaPAMClient) DeleteResourceGroup(ctx context.Context, resourceGroupID string) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/resource_groups/%s", url.PathEscape(c.Team), url.PathEscape(resourceGroupID))
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
