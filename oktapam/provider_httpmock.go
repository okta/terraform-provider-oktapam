package oktapam

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"

	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/go-resty/resty/v2"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	diag2 "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/jarcoal/httpmock"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
)

// httpMockTestV6ProviderFactories is the cousin of testAccV6ProviderFactories, except that
// the http.Client are wired into httpmock for testing.
func httpMockTestV6ProviderFactories() map[string]func() (tfprotov6.ProviderServer, error) {
	serverFactory, err := ProviderServerFactory(context.Background(), &ProviderServerConfig{
		V5ClientCreator: httpMockV5ClientCreator,
		V6ClientCreator: httpMockV6ClientCreator,
	})
	if err != nil {
		log.Fatal(err)
	}
	factories := map[string]func() (tfprotov6.ProviderServer, error){
		"oktapam": func() (tfprotov6.ProviderServer, error) {
			return serverFactory(), nil
		},
	}
	return factories
}

func httpMockV5ClientCreator(_ context.Context, _ diag2.Diagnostics, _ *schema.ResourceData) (*client.APIClients, diag2.Diagnostics) {
	return httpMockClients(), nil
}

func httpMockV6ClientCreator(_ *FrameworkProvider, _ FrameworkProviderModel) (*client.APIClients, diag.Diagnostics) {
	return httpMockClients(), nil
}

func httpMockClients() *client.APIClients {
	httpMockTestConfig := &client.OktaPAMProviderConfig{
		APIKey:       "httpmock-test-api-key-v6",
		APIKeySecret: "httpmock-test-api-key-secret-v6",
		Team:         "httpmock-test-team",
		APIHost:      "https://localhost:8443/",
	}

	apiClientConfigOpts := []pam.ConfigOption{
		pam.WithHost(httpMockTestConfig.APIHost),
		pam.WithTeam(httpMockTestConfig.Team),
		pam.WithBearerToken("ignored"),
		pam.WithTrustedDomainOverride("localhost"),
	}

	pamClient, err := pam.NewAPIClient(apiClientConfigOpts...)
	if err != nil {
		panic("failed to create mock http client " + err.Error())
	}

	httpmock.ActivateNonDefault(pamClient.GetHTTPClient())

	restyClient := resty.New()
	httpmock.ActivateNonDefault(restyClient.GetClient())
	localClient, err := client.NewOktaPAMClient(restyClient, httpMockTestConfig)
	if err != nil {
		panic("failed to create mock deprecated client" + err.Error())
	}

	return &client.APIClients{
		SDKClient: client.SDKClientWrapper{
			SDKClient: pamClient,
			Team:      httpMockTestConfig.Team,
		},
		LocalClient: localClient,
	}
}

// SetupDefaultMockResponders configures common mock responses for standard resources
func SetupDefaultMockResponders(groupID string, resourceGroupID string, projectID string, groupName string, resourceGroupName string, projectName string) {
	// Group endpoints
	httpmock.RegisterRegexpResponder("POST", regexp.MustCompile(`/v1/teams/httpmock-test-team/groups`),
		func(req *http.Request) (*http.Response, error) {
			var requestBody map[string]interface{}
			if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
				return httpmock.NewStringResponse(400, fmt.Sprintf("Failed to decode: %v", err)), nil
			}
			return httpmock.NewJsonResponse(201, map[string]interface{}{
				"id":   groupID,
				"name": requestBody["name"].(string),
			})
		},
	)

	httpmock.RegisterRegexpResponder("GET", regexp.MustCompile(`/v1/teams/httpmock-test-team/groups/.*`),
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewJsonResponse(200, map[string]interface{}{
				"id":   groupID,
				"name": groupName,
			})
		},
	)

	// Resource Group endpoints
	httpmock.RegisterRegexpResponder("POST", regexp.MustCompile(`^/v1/teams/httpmock-test-team/resource_groups$`),
		func(req *http.Request) (*http.Response, error) {
			bodyBytes, _ := io.ReadAll(req.Body)
			req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			var requestBody struct {
				Name                         string        `json:"name"`
				Description                  string        `json:"description"`
				DelegatedResourceAdminGroups []interface{} `json:"delegated_resource_admin_groups"`
			}
			if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
				return httpmock.NewStringResponse(400, fmt.Sprintf("Failed to decode: %v", err)), nil
			}
			return httpmock.NewJsonResponse(201, map[string]interface{}{
				"id":          resourceGroupID,
				"name":        requestBody.Name,
				"description": requestBody.Description,
				"delegated_resource_admin_groups": []map[string]interface{}{
					{
						"id":   groupID,
						"type": "group",
					},
				},
			})
		},
	)

	httpmock.RegisterRegexpResponder("GET", regexp.MustCompile(`^/v1/teams/httpmock-test-team/resource_groups/[^/]+$`),
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewJsonResponse(200, map[string]interface{}{
				"id":          resourceGroupID,
				"name":        resourceGroupName,
				"description": "test resource group",
				"delegated_resource_admin_groups": []map[string]interface{}{
					{
						"id":   groupID,
						"type": "user_group",
					},
				},
			})
		},
	)

	// Resource Group Project endpoints
	httpmock.RegisterRegexpResponder("POST", regexp.MustCompile(`^/v1/teams/httpmock-test-team/resource_groups/[^/]+/projects$`),
		func(req *http.Request) (*http.Response, error) {
			matches := regexp.MustCompile(`/resource_groups/([^/]+)/projects`).FindStringSubmatch(req.URL.Path)
			if len(matches) != 2 {
				return httpmock.NewStringResponse(400, "Invalid URL format"), nil
			}
			actualResourceGroupID := matches[1]

			var requestBody struct {
				Name               string `json:"name"`
				SSHCertificateType string `json:"ssh_certificate_type"`
				AccountDiscovery   bool   `json:"account_discovery"`
				GatewaySelector    string `json:"gateway_selector,omitempty"`
			}
			if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
				return httpmock.NewStringResponse(400, fmt.Sprintf("Failed to decode: %v", err)), nil
			}
			return httpmock.NewJsonResponse(201, map[string]interface{}{
				"id":                   projectID,
				"name":                 requestBody.Name,
				"resource_group":       actualResourceGroupID,
				"team":                 "httpmock-test-team",
				"ssh_certificate_type": requestBody.SSHCertificateType,
				"account_discovery":    requestBody.AccountDiscovery,
				"gateway_selector":     requestBody.GatewaySelector,
			})
		},
	)

	httpmock.RegisterRegexpResponder("GET", regexp.MustCompile(`/v1/teams/httpmock-test-team/resource_groups/.*/projects/.*`),
		func(req *http.Request) (*http.Response, error) {
			matches := regexp.MustCompile(`/resource_groups/([^/]+)/projects/([^/]+)`).FindStringSubmatch(req.URL.Path)
			if len(matches) != 3 {
				return httpmock.NewStringResponse(404, "Project not found"), nil
			}
			actualResourceGroupID := matches[1]
			actualProjectID := matches[2]

			if actualResourceGroupID != resourceGroupID || actualProjectID != projectID {
				return httpmock.NewStringResponse(404, "Project not found"), nil
			}
			return httpmock.NewJsonResponse(200, map[string]interface{}{
				"id":                   actualProjectID,
				"name":                 projectName,
				"resource_group":       actualResourceGroupID,
				"team":                 "httpmock-test-team",
				"ssh_certificate_type": "CERT_TYPE_ED25519_01",
				"account_discovery":    true,
				"deleted_at":           nil,
			})
		},
	)

	// DELETE responders
	httpmock.RegisterRegexpResponder("DELETE", regexp.MustCompile(`/v1/teams/httpmock-test-team/groups/.*`),
		httpmock.NewStringResponder(204, ""))

	httpmock.RegisterRegexpResponder("DELETE", regexp.MustCompile(`/v1/teams/httpmock-test-team/resource_groups/.*`),
		httpmock.NewStringResponder(204, ""))

	httpmock.RegisterRegexpResponder("DELETE", regexp.MustCompile(`/v1/teams/httpmock-test-team/resource_groups/.*/projects/.*`),
		httpmock.NewStringResponder(204, ""))
}
