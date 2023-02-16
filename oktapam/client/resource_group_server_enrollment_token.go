package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"

	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
	"github.com/tomnomnom/linkheader"
)

type ResourceGroupServerEnrollmentToken struct {
	Project       *string `json:"-"`
	ResourceGroup *string `json:"-"`
	Description   *string `json:"description"`
	Token         *string `json:"token,omitempty"`
	ID            *string `json:"id,omitempty"`
	CreatedByUser *string `json:"created_by_user,omitempty"`
	IssuedAt      *string `json:"issued_at,omitempty"`
}

func (t ResourceGroupServerEnrollmentToken) ToResourceMap() map[string]any {
	m := make(map[string]any)

	if t.Project != nil {
		m[attributes.Project] = *t.Project
	}
	if t.ResourceGroup != nil {
		m[attributes.ResourceGroup] = *t.ResourceGroup
	}
	if t.Description != nil {
		m[attributes.Description] = *t.Description
	}
	if t.Token != nil {
		m[attributes.Token] = *t.Token
	}
	if t.ID != nil {
		m[attributes.ID] = *t.ID
	}
	if t.CreatedByUser != nil {
		m[attributes.CreatedByUser] = *t.CreatedByUser
	}
	if t.IssuedAt != nil {
		m[attributes.IssuedAt] = *t.IssuedAt
	}

	return m
}

type ResourceGroupServerEnrollmentTokensListResponse struct {
	Tokens []ResourceGroupServerEnrollmentToken `json:"list"`
}

func (c OktaPAMClient) ListResourceGroupServerEnrollmentTokens(ctx context.Context, resourceGroupID string, projectID string) ([]ResourceGroupServerEnrollmentToken, error) {
	if resourceGroupID == "" {
		return nil, fmt.Errorf("supplied blank resource group id")
	}
	if projectID == "" {
		return nil, fmt.Errorf("supplied blank project id")
	}
	requestURL := fmt.Sprintf("/v1/teams/%s/resource_groups/%s/projects/%s/server_enrollment_tokens", url.PathEscape(c.Team), url.PathEscape(resourceGroupID), url.PathEscape(projectID))
	tokens := make([]ResourceGroupServerEnrollmentToken, 0)

	for {
		// List will paginate, so we make a request, add results to array to return, check if we get a next page, and if so loop again
		logging.Tracef("making GET request to %s", requestURL)
		resp, err := c.CreateBaseRequest(ctx).
			SetResult(&ResourceGroupServerEnrollmentTokensListResponse{}).
			Get(requestURL)
		if err != nil {
			logging.Errorf("received error while making request to %s", requestURL)
			return nil, err
		}
		if _, err := checkStatusCode(resp, http.StatusOK, http.StatusNotFound); err != nil {
			return nil, err
		}

		if resp.StatusCode() == http.StatusNotFound {
			logging.Warnf("received a 404 for %s, could indicate that the referenced project does not exist", requestURL)
			break
		}

		tokensListResponse := resp.Result().(*ResourceGroupServerEnrollmentTokensListResponse)
		tokens = append(tokens, tokensListResponse.Tokens...)

		linkHeader := resp.Header().Get("Link")
		if linkHeader == "" {
			break
		}
		links := linkheader.Parse(linkHeader)
		requestURL = ""
		for _, link := range links {
			if link.Rel == "next" {
				requestURL = link.URL
			}
		}
		if requestURL == "" {
			break
		}
	}

	for idx, token := range tokens {
		token.Project = &projectID
		token.ResourceGroup = &resourceGroupID
		tokens[idx] = token
	}

	return tokens, nil
}

func (c OktaPAMClient) GetResourceGroupServerEnrollmentToken(ctx context.Context, resourceGroupID, projectID, id string) (*ResourceGroupServerEnrollmentToken, error) {
	if resourceGroupID == "" {
		return nil, fmt.Errorf("supplied blank resource group id")
	}
	if projectID == "" {
		return nil, fmt.Errorf("supplied blank project id")
	}
	if id == "" {
		return nil, fmt.Errorf("supplied blank enrollment token id")
	}
	requestURL := fmt.Sprintf("/v1/teams/%s/resource_groups/%s/projects/%s/server_enrollment_tokens/%s", url.PathEscape(c.Team), url.PathEscape(resourceGroupID), url.PathEscape(projectID), url.PathEscape(id))

	logging.Tracef("making GET request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetResult(&ResourceGroupServerEnrollmentToken{}).Get(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}
	if _, err := checkStatusCode(resp, http.StatusOK); err != nil {
		return nil, err
	}
	token := resp.Result().(*ResourceGroupServerEnrollmentToken)
	token.Project = &projectID
	token.ResourceGroup = &resourceGroupID

	return token, nil
}

func (c OktaPAMClient) CreateResourceGroupServerEnrollmentToken(ctx context.Context, resourceGroupID, projectID, description string) (ResourceGroupServerEnrollmentToken, error) {
	if resourceGroupID == "" {
		return ResourceGroupServerEnrollmentToken{}, fmt.Errorf("supplied blank resource group id")
	}
	if projectID == "" {
		return ResourceGroupServerEnrollmentToken{}, fmt.Errorf("supplied blank project id")
	}
	requestURL := fmt.Sprintf("/v1/teams/%s/resource_groups/%s/projects/%s/server_enrollment_tokens", url.PathEscape(c.Team), url.PathEscape(resourceGroupID), url.PathEscape(projectID))

	request := map[string]any{
		attributes.Description: description,
	}
	token := ResourceGroupServerEnrollmentToken{}
	logging.Tracef("making POST request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetBody(request).SetResult(&token).Post(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return ResourceGroupServerEnrollmentToken{}, err
	}
	if _, err := checkStatusCode(resp, 201); err != nil {
		logging.Tracef("unexpected status code: %d", resp.StatusCode())
		return ResourceGroupServerEnrollmentToken{}, err
	}
	token.Project = &projectID
	token.ResourceGroup = &resourceGroupID

	return token, nil
}

func (c OktaPAMClient) DeleteResourceGroupServerEnrollmentToken(ctx context.Context, resourceGroupID, projectID, id string) error {
	if resourceGroupID == "" {
		return fmt.Errorf("supplied blank resource group id")
	}
	if projectID == "" {
		return fmt.Errorf("supplied blank project id")
	}
	if id == "" {
		return fmt.Errorf("supplied blank enrollment token id")
	}

	requestURL := fmt.Sprintf("/v1/teams/%s/resource_groups/%s/projects/%s/server_enrollment_tokens/%s", url.PathEscape(c.Team), url.PathEscape(resourceGroupID), url.PathEscape(projectID), url.PathEscape(id))

	logging.Tracef("making DELETE request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).Delete(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}

	_, err = checkStatusCode(resp, http.StatusNoContent, http.StatusNotFound)
	return err
}
