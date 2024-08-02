package oktapam

import (
	"context"
	"fmt"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/config"
	"log"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// providerFactories are used to instantiate a provider during acceptance testing.
// The factory function will be invoked for every Terraform CLI command executed
// to create a provider server to which the CLI can reattach.

// var ProviderServerFactory map[string]func() (tfprotov6.ProviderServer, error)
var providerFactories map[string]func() (tfprotov6.ProviderServer, error)
var accTestAPIClients *client.APIClients

const DefaultTestTeam = "pam-tf-provider-testing"

func init() {
	providerServerFactory, err := ServerFactoryV6(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	providerFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"oktapam": func() (tfprotov6.ProviderServer, error) {
			return providerServerFactory(), nil
		},
	}

	accTestAPIClients, err = newAcceptanceTestingClient()
	if err != nil {
		log.Fatal(err)
	}
}

func testAccPreCheck(t *testing.T) {
	requiredEnvVars := []string{config.ApiKeySchemaEnvVar, config.ApiKeySecretSchemaEnvVar, config.TeamSchemaEnvVar}

	for _, envVar := range requiredEnvVars {
		if err := os.Getenv(envVar); err == "" {
			t.Fatalf("%s must be set for acceptance tests", envVar)
		}
	}
}

func getTeamName() string {
	teamName := os.Getenv(config.TeamSchemaEnvVar)
	if teamName != "" {
		return teamName
	}
	return DefaultTestTeam
}

func newAcceptanceTestingClient() (*client.APIClients, error) {
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
