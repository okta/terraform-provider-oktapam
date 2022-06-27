package client

import (
	"context"
	"fmt"
	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/logging"
	"github.com/tomnomnom/linkheader"
	"net/url"
)

type Gateway struct {
	Id                string            `json:"id"`
	Name              string            `json:"name"`
	AccessAddress     string            `json:"access_address"`
	DefaultAddress    string            `json:"default_address"`
	Description       string            `json:"description"`
	CloudProvider     string            `json:"cloud_provider"`
	RefuseConnections bool              `json:"refuse_connections"`
	Labels            map[string]string `json:"labels"`
}

type ListGatewayParameters struct {
	Contains string
}

func (p ListGatewayParameters) toQueryParametersMap() map[string]string {
	m := make(map[string]string, 1)
	if p.Contains != "" {
		m[attributes.Contains] = p.Contains
	}
	return m
}

func (gateway Gateway) ToResourceMap() map[string]interface{} {
	m := make(map[string]interface{})
	m[attributes.Name] = gateway.Name
	m[attributes.ID] = gateway.Id
	m[attributes.AccessAddress] = gateway.AccessAddress
	m[attributes.DefaultAddress] = gateway.DefaultAddress
	m[attributes.Description] = gateway.Description
	m[attributes.RefuseConnections] = gateway.RefuseConnections
	m[attributes.Labels] = gateway.Labels
	return m
}

type GatewayListResponse struct {
	Gateways []Gateway `json:"list"`
}

func (c OktaPAMClient) ListGateways(ctx context.Context, parameters ListGatewayParameters) ([]Gateway, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/gateways", url.PathEscape(c.Team))
	gateways := make([]Gateway, 0)

	for {
		// List will paginate, so we make a request, add results to array to return, check if we get a next page, and if so loop again
		logging.Tracef("making GET request to %s", requestURL)

		resp, err := c.CreateBaseRequest(ctx).SetQueryParams(parameters.toQueryParametersMap()).SetResult(&GatewayListResponse{}).Get(requestURL)
		if err != nil {
			logging.Errorf("received error while making request to %s", requestURL)
			return nil, err
		}
		if _, err := checkStatusCode(resp, 200); err != nil {
			return nil, err
		}

		gatewayListResponse := resp.Result().(*GatewayListResponse)
		gateways = append(gateways, gatewayListResponse.Gateways...)

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

	return gateways, nil
}
