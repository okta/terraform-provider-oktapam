package client

import (
	"context"
	"fmt"

	"github.com/terraform-providers/terraform-provider-oktaasa/oktaasa/logging"
	"github.com/tomnomnom/linkheader"
)

type ServerEnrollmentToken struct {
	Project       string `json:"_"`
	Token         string `json:"token"`
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
	requestURL := fmt.Sprintf("/v1/teams/%s/projects/%s/server_enrollment_tokens", c.Team, project)
	tokens := make([]ServerEnrollmentToken, 0)

	for {
		logging.Tracef("making GET request to %s", requestURL)
		resp, err := c.CreateBaseRequest(ctx).
			SetResult(&ServerEnrollmentTokensListResponse{}).
			Get(requestURL)
		if err != nil {
			logging.Errorf("received error while making request to %s", requestURL)
			return []ServerEnrollmentToken{}, err
		}
		err = checkStatusCode(resp, 200)
		if err != nil {
			return []ServerEnrollmentToken{}, err
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

func (c OktaASAClient) GetServerEnrollmentToken(ctx context.Context, project, id string) (ServerEnrollmentToken, error) {
	if project == "" {
		return ServerEnrollmentToken{}, fmt.Errorf("supplied blank project name")
	}
	if id == "" {
		return ServerEnrollmentToken{}, fmt.Errorf("supplied blank enrollment token id")
	}
	requestURL := fmt.Sprintf("/v1/teams/%s/projects/%s/server_enrollment_tokens/%s", c.Team, project, id)

	logging.Tracef("making GET request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetResult(&ServerEnrollmentToken{}).Get(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return ServerEnrollmentToken{}, err
	}
	err = checkStatusCode(resp, 200)
	if err != nil {
		return ServerEnrollmentToken{}, err
	}
	token := resp.Result().(*ServerEnrollmentToken)
	token.Project = project

	return *token, nil
}

func (c OktaASAClient) CreateServerEnrollmentToken(ctx context.Context, project, description string) (ServerEnrollmentToken, error) {
	if project == "" {
		return ServerEnrollmentToken{}, fmt.Errorf("supplied blank project name")
	}
	requestURL := fmt.Sprintf("/v1/teams/%s/projects/%s/server_enrollment_tokens", c.Team, project)

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
	err = checkStatusCode(resp, 201)
	if err != nil {
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

	requestURL := fmt.Sprintf("/v1/teams/%s/projects/%s/server_enrollment_tokens/%s", c.Team, project, id)

	logging.Tracef("making DELETE request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).Delete(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}

	return checkStatusCode(resp, 204, 404)
}
