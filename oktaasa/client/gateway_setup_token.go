package client

import (
	"context"
	"fmt"
	"net/url"

	"github.com/terraform-providers/terraform-provider-oktaasa/oktaasa/logging"
	"github.com/tomnomnom/linkheader"
)

// defines the type of gateway that we'll be registering, currently the only type allowed
var gatewayAgentRegistrationType = "gateway-agent"

type GatewaySetupToken struct {
	ID               *string                        `json:"id,omitempty"`
	Description      *string                        `json:"description"`
	RegistrationType *string                        `json:"registration_type"`
	CreatedAt        *string                        `json:"created_at,omitempty"`
	Details          *GatewaySetupTokenLabelDetails `json:"details,omitempty"`
}

type GatewaySetupTokenLabelDetails struct {
	Labels map[string]string `json:"labels,omitempty"`
}

func (t GatewaySetupToken) ToResourceMap() map[string]interface{} {
	m := make(map[string]interface{})

	m["description"] = t.Description

	if t.ID != nil {
		m["id"] = *t.ID
	}
	if t.CreatedAt != nil {
		m["created_at"] = *t.CreatedAt
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
	requestURL := fmt.Sprintf("/v1/teams/%s/gateway_setup_tokens", url.PathEscape(c.Team))
	tokens := make([]GatewaySetupToken, 0)

	for {
		// List will paginate, so we make a request, add results to array to return, check if we get a next page, and if so loop again
		logging.Tracef("making GET request to %s", requestURL)
		resp, err := c.CreateBaseRequest(ctx).
			SetResult(&GatewaySetupTokensListResponse{}).
			Get(requestURL)
		if err != nil {
			logging.Errorf("received error while making request to %s", requestURL)
			return nil, err
		}
		if _, err := checkStatusCode(resp, 200); err != nil {
			return nil, err
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

func (c OktaASAClient) GetGatewaySetupToken(ctx context.Context, id string) (*GatewaySetupToken, error) {
	if id == "" {
		return nil, fmt.Errorf("supplied blank gateway setup token id")
	}
	requestURL := fmt.Sprintf("/v1/teams/%s/gateway_setup_tokens/%s", url.PathEscape(c.Team), url.PathEscape(id))

	logging.Tracef("making GET request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetResult(&GatewaySetupToken{}).Get(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}
	if statusCode, err := checkStatusCode(resp, 200, 404); err != nil {
		return nil, err
	} else if statusCode == 404 {
		return nil, nil
	}

	token := resp.Result().(*GatewaySetupToken)
	return token, nil
}

func (c OktaASAClient) CreateGatewaySetupToken(ctx context.Context, token GatewaySetupToken) (*GatewaySetupToken, error) {
	// create a token with the values specified within token
	requestURL := fmt.Sprintf("/v1/teams/%s/gateway_setup_tokens", url.PathEscape(c.Team))
	logging.Tracef("making POST request to %s", requestURL)
	token.RegistrationType = &gatewayAgentRegistrationType
	resp, err := c.CreateBaseRequest(ctx).SetBody(token).SetResult(&GatewaySetupToken{}).Post(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}
	if _, err := checkStatusCode(resp, 201); err != nil {
		logging.Tracef("unexpected status code: %d", resp.StatusCode())
		return nil, err
	}
	createdToken := resp.Result().(*GatewaySetupToken)

	return createdToken, nil
}

func (c OktaASAClient) DeleteGatewaySetupToken(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("supplied blank gateway setup token id")
	}
	requestURL := fmt.Sprintf("/v1/teams/%s/gateway_setup_tokens/%s", url.PathEscape(c.Team), url.PathEscape(id))
	logging.Tracef("making DELETE request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).Delete(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}
	if _, err := checkStatusCode(resp, 204); err != nil {
		logging.Tracef("unexpected status code: %d", resp.StatusCode())
		return err
	}

	return nil
}
