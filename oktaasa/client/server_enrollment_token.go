package client

import (
	"context"
	"fmt"
	"net/url"

	"github.com/terraform-providers/terraform-provider-oktaasa/oktaasa/logging"
	"github.com/tomnomnom/linkheader"
)

type ServerEnrollmentToken struct {
	Project       string `json:"_"`
	Token         string `json:"token,omitempty"`
	ID            string `json:"id,omitempty"`
	Description   string `json:"description,omitempty"`
	CreatedByUser string `json:"created_by_user,omitempty"`
	IssuedAt      string `json:"issued_at,omitempty"`
}

func (t ServerEnrollmentToken) ToMap() map[string]interface{} {
	m := make(map[string]interface{})

	if t.Project != "" {
		m["project_name"] = t.Project
	}
	if t.Token != "" {
		m["token"] = t.Token
	}
	if t.ID != "" {
		m["id"] = t.ID
	}
	if t.Description != "" {
		m["description"] = t.Description
	}
	if t.CreatedByUser != "" {
		m["created_by_user"] = t.CreatedByUser
	}
	if t.IssuedAt != "" {
		m["issued_at"] = t.IssuedAt
	}

	return m
}

type ServerEnrollmentTokensListResponse struct {
	Tokens []ServerEnrollmentToken `json:"list"`
}

func (c OktaASAClient) ListServerEnrollmentTokens(ctx context.Context, project string) ([]ServerEnrollmentToken, error) {
	if project == "" {
		return []ServerEnrollmentToken{}, fmt.Errorf("supplied blank project name")
	}
	requestURL := fmt.Sprintf("/v1/teams/%s/projects/%s/server_enrollment_tokens", url.PathEscape(c.Team), url.PathEscape(project))
	tokens := make([]ServerEnrollmentToken, 0)

	for {
		// List will paginate, so we make a request, add results to array to return, check if we get a next page, and if so loop again
		logging.Tracef("making GET request to %s", requestURL)
		resp, err := c.CreateBaseRequest(ctx).
			SetResult(&ServerEnrollmentTokensListResponse{}).
			Get(requestURL)
		if err != nil {
			logging.Errorf("received error while making request to %s", requestURL)
			return nil, err
		}
		if _, err := checkStatusCode(resp, 200); err != nil {
			return nil, err
		}

		tokensListResponse := resp.Result().(*ServerEnrollmentTokensListResponse)
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
				break
			}
		}
	}

	for idx, token := range tokens {
		token.Project = project
		tokens[idx] = token
	}

	return tokens, nil
}

func (c OktaASAClient) GetServerEnrollmentToken(ctx context.Context, project, id string) (*ServerEnrollmentToken, error) {
	if project == "" {
		return nil, fmt.Errorf("supplied blank project name")
	}
	if id == "" {
		return nil, fmt.Errorf("supplied blank enrollment token id")
	}
	requestURL := fmt.Sprintf("/v1/teams/%s/projects/%s/server_enrollment_tokens/%s", url.PathEscape(c.Team), url.PathEscape(project), url.PathEscape(id))

	logging.Tracef("making GET request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetResult(&ServerEnrollmentToken{}).Get(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}
	if _, err := checkStatusCode(resp, 200); err != nil {
		return nil, err
	}
	token := resp.Result().(*ServerEnrollmentToken)
	token.Project = project

	return token, nil
}

func (c OktaASAClient) CreateServerEnrollmentToken(ctx context.Context, project, description string) (ServerEnrollmentToken, error) {
	if project == "" {
		return ServerEnrollmentToken{}, fmt.Errorf("supplied blank project name")
	}
	requestURL := fmt.Sprintf("/v1/teams/%s/projects/%s/server_enrollment_tokens", url.PathEscape(c.Team), url.PathEscape(project))

	request := map[string]interface{}{
		"description": description,
	}
	token := ServerEnrollmentToken{}
	logging.Tracef("making POST request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetBody(request).SetResult(&token).Post(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return ServerEnrollmentToken{}, err
	}
	if _, err := checkStatusCode(resp, 201); err != nil {
		logging.Tracef("unexpected status code: %d", resp.StatusCode())
		return ServerEnrollmentToken{}, err
	}
	token.Project = project

	return token, nil
}

func (c OktaASAClient) DeleteServerEnrollmentToken(ctx context.Context, project, id string) error {
	if project == "" {
		return fmt.Errorf("supplied blank project name")
	}
	if id == "" {
		return fmt.Errorf("supplied blank enrollment token id")
	}

	requestURL := fmt.Sprintf("/v1/teams/%s/projects/%s/server_enrollment_tokens/%s", url.PathEscape(c.Team), url.PathEscape(project), url.PathEscape(id))

	logging.Tracef("making DELETE request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).Delete(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}

	_, err = checkStatusCode(resp, 204, 404)
	return err
}
