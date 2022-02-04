package oktaasa

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-oktaasa/oktaasa/client"
)

const DefaultAPIBaseURL = "https://app.scaleft.com"

const apiHostSchemaEnvVar = "OKTAASA_API_HOST"
const apiKeySchemaEnvVar = "OKTAASA_KEY"
const apiKeySecretSchemaEnvVar = "OKTAASA_SECRET"
const teamSchemaEnvVar = "OKTAASA_TEAM"

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"oktaasa_api_host": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(apiHostSchemaEnvVar, DefaultAPIBaseURL),
				Description: "Okta ASA API Host",
			},
			"oktaasa_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc(apiKeySchemaEnvVar, nil),
				Description: "Okta ASA API Key",
			},
			"oktaasa_secret": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc(apiKeySecretSchemaEnvVar, nil),
				Description: "Okta ASA API Secret",
			},
			"oktaasa_team": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc(teamSchemaEnvVar, nil),
				Description: "Okta ASA Team",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"oktaasa_project":                 resourceProject(),
			"oktaasa_group":                   resourceGroup(),
			"oktaasa_server_enrollment_token": resourceServerEnrollmentToken(),
			"oktaasa_project_group":           resourceProjectGroup(),
			"oktaasa_gateway_setup_token":     resourceGatewaySetupToken(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"oktaasa_projects":                 dataSourceProjects(),
			"oktaasa_groups":                   dataSourceGroups(),
			"oktaasa_server_enrollment_tokens": dataSourceServerEnrollmentTokens(),
			"oktaasa_project_groups":           dataSourceProjectGroups(),
			"oktaasa_gateway_setup_tokens":     dataSourceGatewaySetupTokens(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	apiKey := d.Get("oktaasa_key").(string)
	apiKeySecret := d.Get("oktaasa_secret").(string)
	team := d.Get("oktaasa_team").(string)
	apiHost := d.Get("oktaasa_api_host").(string)

	client, err := client.CreateOktaASAClient(apiKey, apiKeySecret, team, apiHost)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return client, nil
}
