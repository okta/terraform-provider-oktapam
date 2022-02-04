package client

import (
	"context"
	"fmt"

	"github.com/terraform-providers/terraform-provider-oktaasa/oktaasa/logging"
	"github.com/tomnomnom/linkheader"
)

const gatewayAgentRegistrationType = "gateway-agent"

type GatewaySetupToken struct {
	ID               string                         `json:"id,omitempty"`
	Description      string                         `json:"description,omitempty"`
	RegistrationType string                         `json:"registration_type,omitempty"`
	CreatedAt        string                         `json:"created_at,omitempty"`
	Details          *GatewaySetupTokenLabelDetails `json:"details,omitempty"`
}

type GatewaySetupTokenLabelDetails struct {
	Labels map[string]string `json:"labels,omitempty"`
}

func (t GatewaySetupToken) ToMap() map[string]interface{} {
	m := make(map[string]interface{})

	if t.ID != "" {
		m["id"] = t.ID
	}
	if t.Description != "" {
		m["description"] = t.Description
	}
	if t.CreatedAt != "" {
		m["created_at"] = t.CreatedAt
	}
	if t.Details != nil {
		m["labels"] = t.Details.Labels
	}

	return m
}

type GatewaySetupTokensListResponse struct {
	Tokens []GatewaySetupToken `json:"list"`
}

func (c OktaASAClient) ListGatewaySetupTokens(ctx context.Context) ([]GatewaySetupToken, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/gateway_setup_tokens", c.Team)
	tokens := make([]GatewaySetupToken, 0)

	for {
		logging.Tracef("making GET request to %s", requestURL)
		resp, err := c.CreateBaseRequest(ctx).
			SetResult(&GatewaySetupTokensListResponse{}).
			Get(requestURL)
		if err != nil {
			logging.Errorf("received error while making request to %s", requestURL)
			return []GatewaySetupToken{}, err
		}
		err = checkStatusCode(resp, 200)
		if err != nil {
			return []GatewaySetupToken{}, err
		}

		tokensListResponse := resp.Result().(*GatewaySetupTokensListResponse)
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

	return tokens, nil
}

func (c *OktaASAClient) GetGatewaySetupToken(ctx context.Context, id string) (GatewaySetupToken, error) {
	if id == "" {
		return GatewaySetupToken{}, fmt.Errorf("supplied blank gateway setup token id")
	}
	requestURL := fmt.Sprintf("/v1/teams/%s/gateway_setup_tokens/%s", c.Team, id)

	logging.Tracef("making GET request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetResult(&GatewaySetupToken{}).Get(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return GatewaySetupToken{}, err
	}
	err = checkStatusCode(resp, 200, 404)
	if err != nil {
		return GatewaySetupToken{}, err
	}
	if resp.StatusCode() == 404 {
		return GatewaySetupToken{}, nil
	}

	token := resp.Result().(*GatewaySetupToken)
	return *token, nil
}

func (c *OktaASAClient) CreateGatewaySetupToken(ctx context.Context, token GatewaySetupToken) (GatewaySetupToken, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/gateway_setup_tokens", c.Team)
	logging.Tracef("making POST request to %s", requestURL)
	token.RegistrationType = gatewayAgentRegistrationType
	resp, err := c.CreateBaseRequest(ctx).SetBody(token).SetResult(&GatewaySetupToken{}).Post(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return GatewaySetupToken{}, err
	}
	err = checkStatusCode(resp, 201)
	if err != nil {
		logging.Tracef("unexpected status code: %d", resp.StatusCode())
		return GatewaySetupToken{}, err
	}
	createdToken := resp.Result().(*GatewaySetupToken)

	return *createdToken, nil
}

func (c *OktaASAClient) DeleteGatewaySetupToken(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("supplied blank gateway setup token id")
	}
	requestURL := fmt.Sprintf("/v1/teams/%s/gateway_setup_tokens/%s", c.Team, id)
	logging.Tracef("making DELETE request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).Delete(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}
	err = checkStatusCode(resp, 204)
	if err != nil {
		logging.Tracef("unexpected status code: %d", resp.StatusCode())
		return err
	}

	return nil
}
