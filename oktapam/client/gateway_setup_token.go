package client

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"

	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
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
	Token            *string                        `json:"token,omitempty"`
}

type GatewaySetupTokenValue struct {
	Token *string `json:"token,omitempty"`
}

type GatewaySetupTokenLabelDetails struct {
	Labels map[string]string `json:"labels,omitempty"`
}

func (t GatewaySetupToken) ToResourceMap() map[string]interface{} {
	m := make(map[string]interface{})

	m[attributes.Description] = t.Description

	if t.ID != nil {
		m[attributes.ID] = *t.ID
	}
	if t.CreatedAt != nil {
		m[attributes.CreatedAt] = *t.CreatedAt
	}
	if t.Details != nil {
		m[attributes.Labels] = t.Details.Labels
	}

	if t.Token != nil {
		m[attributes.Token] = t.Token
	}

	return m
}

type GatewaySetupTokensListResponse struct {
	Tokens []GatewaySetupToken `json:"list"`
}

func (c OktaPAMClient) ListGatewaySetupTokens(ctx context.Context, descriptionContains string) ([]GatewaySetupToken, error) {
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
		if descriptionContains == "" {
			tokens = append(tokens, tokensListResponse.Tokens...)
		} else {
			// the API does not currently allow for a filter here, so we do this client-side
			// Note that all tokens returned by the API will still be scoped to the user's access
			for _, token := range tokensListResponse.Tokens {
				if strings.Contains(*token.Description, descriptionContains) {
					tokens = append(tokens, token)
				}
			}
		}

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

	// Retrieve token values for each resource
	for _, token := range tokens {
		value, err := c.GetGatewaySetupTokenValue(ctx, *token.ID)
		if err != nil {
			logging.Errorf("received error while retrieving token value for token id %s", *token.ID)
			return nil, err
		}
		token.Token = value.Token
	}

	return tokens, nil
}

func (c OktaPAMClient) GetGatewaySetupToken(ctx context.Context, id string) (*GatewaySetupToken, error) {
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

	// Retrieve token value
	value, err := c.GetGatewaySetupTokenValue(ctx, *token.ID)
	if err != nil {
		logging.Errorf("received error while retrieving token value for token id %s", *token.ID)
		return nil, err
	}
	token.Token = value.Token

	return token, nil
}

func (c OktaPAMClient) GetGatewaySetupTokenValue(ctx context.Context, id string) (*GatewaySetupTokenValue, error) {
	if id == "" {
		return nil, fmt.Errorf("supplied blank gateway setup token id")
	}
	requestURL := fmt.Sprintf("/v1/teams/%s/gateway_setup_tokens/%s/token", url.PathEscape(c.Team), url.PathEscape(id))

	logging.Tracef("making GET request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetResult(&GatewaySetupTokenValue{}).Get(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}
	if statusCode, err := checkStatusCode(resp, 200, 404); err != nil {
		return nil, err
	} else if statusCode == 404 {
		return nil, nil
	}

	token := resp.Result().(*GatewaySetupTokenValue)
	return token, nil
}

func (c OktaPAMClient) CreateGatewaySetupToken(ctx context.Context, token GatewaySetupToken) (*GatewaySetupToken, error) {
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

func (c OktaPAMClient) DeleteGatewaySetupToken(ctx context.Context, id string) error {
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
