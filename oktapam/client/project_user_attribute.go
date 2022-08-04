package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"

	"github.com/okta/terraform-provider-oktapam/oktapam/utils"

	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
)

type ProjectUserAttribute struct {
	ID             *string     `json:"id,omitempty"`
	AttributeName  *string     `json:"attribute_name,omitempty"`
	AttributeValue interface{} `json:"attribute_value,omitempty"`
	Managed        *bool       `json:"managed,omitempty"`
	Inherited      *bool       `json:"inherited,omitempty"`
}

type ProjectUserAttributeMap struct {
	Key   string
	Value ProjectUserAttribute
}

type ProjectUserAttributes struct {
	Attributes map[string]ProjectUserAttribute `json:"attributes"`
}

type UserAttributePatchOperation struct {
	Op    string      `json:"op,omitempty"`
	Path  string      `json:"path,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

type UserAttributePatch struct {
	Operations []*UserAttributePatchOperation `json:"operations"`
}

type UpdateAttribute struct {
	AttributeName  string `json:"attribute_name,omitempty"`
	AttributeValue interface{} `json:"attribute_value,omitempty"`
}

func (attr ProjectUserAttribute) ToResourceMap() map[string]interface{} {
	m := make(map[string]interface{})

	if attr.ID != nil {
		m[attributes.ID] = *attr.ID
	}
	if attr.AttributeName != nil {
		m[attributes.ProjectUserAttributeName] = *attr.AttributeName
	}
	if attr.AttributeValue != nil {
		m[attributes.ProjectUserAttributeValue] = attr.AttributeValue
	}
	if attr.Managed != nil {
		m[attributes.ProjectUserAttributeIsManaged] = *attr.Managed
	}
	if attr.Inherited != nil {
		m[attributes.ProjectUserAttributeIsInherited] = *attr.Inherited
	}

	return m
}

func (attr ProjectUserAttribute) Exists() bool {
	return utils.IsNonEmpty(attr.ID)
}

func (c OktaPAMClient) ListProjectUserAttributes(ctx context.Context, projectName string, userName string) (*ProjectUserAttributes, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/projects/%s/users/%s/attributes", url.PathEscape(c.Team), url.PathEscape(projectName), url.PathEscape(userName))
	logging.Tracef("making GET request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetResult(&ProjectUserAttributes{}).Get(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}
	if _, err := checkStatusCode(resp, http.StatusOK); err != nil {
		logging.Tracef("unexpected status code: %d", resp.StatusCode())
		return nil, err
	}

	attributes := resp.Result().(*ProjectUserAttributes)
	return attributes, nil
}

func (c OktaPAMClient) CreateProjectUserAttributes(ctx context.Context, projectName string, userName string, operations *UserAttributePatch) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/projects/%s/users/%s/attributes", url.PathEscape(c.Team), url.PathEscape(projectName), url.PathEscape(userName))

	logging.Tracef("making PATCH request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetBody(operations).Patch(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}
	if _, err := checkStatusCode(resp, http.StatusNoContent); err != nil {
		logging.Tracef("unexpected status code: %d", resp.StatusCode())
		return err
	}
	return nil
}

func (c OktaPAMClient) GetProjectUserAttribute(ctx context.Context, projectName string, userName string, attributeId string) (*ProjectUserAttribute, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/projects/%s/users/%s/attributes/%s", url.PathEscape(c.Team), url.PathEscape(projectName), url.PathEscape(userName), url.PathEscape(attributeId))

	logging.Tracef("making GET request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetResult(&ProjectUserAttribute{}).Get(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}
	statusCode := resp.StatusCode()

	if statusCode == http.StatusOK {
		attribute := resp.Result().(*ProjectUserAttribute)
		if attribute.Exists() {
			return attribute, nil
		}
	} else if statusCode == http.StatusNotFound {
		return nil, nil
	}

	return nil, createErrorForInvalidCode(resp, http.StatusOK, http.StatusNotFound)
}

func (c OktaPAMClient) UpdateProjectUserAttribute(ctx context.Context, projectName string, userName string, attributeId string, update *UpdateAttribute) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/projects/%s/users/%s/attributes/%s", url.PathEscape(c.Team), url.PathEscape(projectName), url.PathEscape(userName), url.PathEscape(attributeId))

	logging.Tracef("making PUT request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetBody(update).Get(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}

	_, err = checkStatusCode(resp, http.StatusNoContent)
	return err
}

func (c OktaPAMClient) DeleteProjectUserAttribute(ctx context.Context, projectName string, userName string, attributeId string) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/projects/%s/users/%s/attributes/%s", url.PathEscape(c.Team), url.PathEscape(projectName), url.PathEscape(userName), url.PathEscape(attributeId))

	logging.Tracef("making DELETE request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).Delete(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}

	_, err = checkStatusCode(resp, http.StatusNoContent, http.StatusNotFound)
	return err
}
