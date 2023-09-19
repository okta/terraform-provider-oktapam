package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
)

type SecretFolder struct {
	Name            *string       `json:"name"`
	ID              *string       `json:"id,omitempty"`
	Description     *string       `json:"description,omitempty"`
	ParentFolderID  *string       `json:"parent_folder_id,omitempty"`
	Path            []NamedObject `json:"path,omitempty"`
	ResourceGroupID *string       `json:"-"`
	ProjectID       *string       `json:"-"`
}

func (s SecretFolder) ToResourceMap() map[string]any {
	m := make(map[string]any, 8)

	if s.Name != nil {
		m[attributes.Name] = *s.Name
	}
	if s.ID != nil {
		m[attributes.ID] = *s.ID
	}
	if s.Description != nil {
		m[attributes.Description] = *s.Description
	}
	if s.ParentFolderID != nil {
		m[attributes.ParentFolder] = *s.ParentFolderID
	}

	return m
}

func (c OktaPAMClient) GetSecretFolder(ctx context.Context, resourceGroupID string, projectID string, id string) (*SecretFolder, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/resource_groups/%s/projects/%s/secret_folders/%s", c.Team, resourceGroupID, projectID, id)

	logging.Tracef("making GET request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetResult(&SecretFolder{}).Get(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}
	statusCode := resp.StatusCode()

	if statusCode == http.StatusOK {
		secretFolder := resp.Result().(*SecretFolder)
		secretFolder.ResourceGroupID = &resourceGroupID
		secretFolder.ProjectID = &projectID
		return secretFolder, nil
	} else if statusCode == http.StatusNotFound {
		return nil, nil
	}

	return nil, createErrorForInvalidCode(resp, http.StatusOK, http.StatusNotFound)
}

func (c OktaPAMClient) CreateSecretFolder(ctx context.Context, secretFolder SecretFolder) (*SecretFolder, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/resource_groups/%s/projects/%s/secret_folders", url.PathEscape(c.Team), url.PathEscape(*secretFolder.ResourceGroupID), url.PathEscape(*secretFolder.ProjectID))
	logging.Tracef("making POST request to %s", requestURL)
	resultingSecretFolder := &SecretFolder{}
	resp, err := c.CreateBaseRequest(ctx).SetBody(secretFolder).SetResult(resultingSecretFolder).Post(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}
	if _, err := checkStatusCode(resp, 201); err == nil {
		return resultingSecretFolder, nil
	} else {
		return nil, err
	}
}

func (c OktaPAMClient) UpdateSecretFolder(ctx context.Context, secretFolder SecretFolder) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/resource_groups/%s/projects/%s/secret_folders/%s", url.PathEscape(c.Team), url.PathEscape(*secretFolder.ResourceGroupID), url.PathEscape(*secretFolder.ProjectID), url.PathEscape(*secretFolder.ID))
	logging.Tracef("making PUT request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetBody(secretFolder).Put(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}
	_, err = checkStatusCode(resp, http.StatusNoContent, http.StatusOK)
	return err
}

func (c OktaPAMClient) DeleteSecretFolder(ctx context.Context, resourceGroupID string, projectID string, id string) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/resource_groups/%s/projects/%s/secret_folders/%s", url.PathEscape(c.Team), url.PathEscape(resourceGroupID), url.PathEscape(projectID), url.PathEscape(id))
	logging.Tracef("making DELETE request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).Delete(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}

	_, err = checkStatusCode(resp, http.StatusNoContent, http.StatusNotFound)
	return err
}
