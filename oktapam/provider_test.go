package oktapam

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/config"
)

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

var testAccSDKV2Providers map[string]func() (*schema.Provider, error)
var testAccSDKV2Provider *schema.Provider

var (
	// testAccV6ProviderFactories are used to instantiate a provider during acceptance testing.
	// The factory function is invoked for every TF CLI command executed to create a provider server to which the CLI can reattach.
	testAccV6ProviderFactories map[string]func() (tfprotov6.ProviderServer, error)
	testAccAPIClients          *client.APIClients
	clientOnce                 sync.Once
)

func init() {
	testAccSDKV2Provider = Provider()
	testAccSDKV2Providers = map[string]func() (*schema.Provider, error){}
	testAccSDKV2Providers["oktapam"] = func() (*schema.Provider, error) {
		return testAccSDKV2Provider, nil
	}

	// this is used in acceptance test
	serverFactory, err := V6ProviderServerFactory(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	testAccV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"oktapam": func() (tfprotov6.ProviderServer, error) {
			return serverFactory(), nil
		},
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	requiredEnvVars := []string{config.ApiKeySchemaEnvVar, config.ApiKeySecretSchemaEnvVar, config.TeamSchemaEnvVar}

	for _, envVar := range requiredEnvVars {
		if err := os.Getenv(envVar); err == "" {
			t.Fatalf("%s must be set for acceptance tests", envVar)
		}
	}
}

func newTestAccAPIClients() (*client.APIClients, error) {
	apiKey := os.Getenv(config.ApiKeySchemaEnvVar)
	apiHost := os.Getenv(config.ApiHostSchemaEnvVar)
	apiSecret := os.Getenv(config.ApiKeySecretSchemaEnvVar)
	team := os.Getenv(config.TeamSchemaEnvVar)

	cfg := &client.OktaPAMProviderConfig{
		APIKey:       apiKey,
		APIKeySecret: apiSecret,
		Team:         team,
		APIHost:      apiHost,
	}

	sdkClient, err := client.CreateSDKClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("error while creating sdk client: %w", err)
	}

	sdkClientWrapper := client.SDKClientWrapper{
		SDKClient: sdkClient,
		Team:      team,
	}

	localClient, err := client.CreateLocalPAMClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("error while creating sdk client: %w", err)
	}

	return &client.APIClients{
		SDKClient:   sdkClientWrapper,
		LocalClient: localClient,
	}, nil
}

func getTestAccAPIClients() *client.APIClients {
	clientOnce.Do(func() {
		testAccAPIClients, _ = newTestAccAPIClients()
	})

	return testAccAPIClients
}
