package oktapam

import (
	"context"
	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/go-resty/resty/v2"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	diag2 "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/jarcoal/httpmock"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"log"
)

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
