package oktapam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/client"
)

const DefaultAPIBaseURL = "https://app.scaleft.com"

const apiHostSchemaEnvVar = "OKTAPAM_API_HOST"
const apiKeySchemaEnvVar = "OKTAPAM_KEY"
const apiKeySecretSchemaEnvVar = "OKTAPAM_SECRET"
const teamSchemaEnvVar = "OKTAPAM_TEAM"

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"oktapam_api_host": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(apiHostSchemaEnvVar, DefaultAPIBaseURL),
				Description: "Okta PAM API Host",
			},
			"oktapam_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc(apiKeySchemaEnvVar, nil),
				Description: "Okta PAM API Key",
			},
			"oktapam_secret": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc(apiKeySecretSchemaEnvVar, nil),
				Description: "Okta PAM API Secret",
			},
			"oktapam_team": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc(teamSchemaEnvVar, nil),
				Description: "Okta PAM Team",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"oktapam_project":                 resourceProject(),
			"oktapam_group":                   resourceGroup(),
			"oktapam_server_enrollment_token": resourceServerEnrollmentToken(),
			"oktapam_project_group":           resourceProjectGroup(),
			"oktapam_gateway_setup_token":     resourceGatewaySetupToken(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"oktapam_projects":                 dataSourceProjects(),
			"oktapam_groups":                   dataSourceGroups(),
			"oktapam_server_enrollment_tokens": dataSourceServerEnrollmentTokens(),
			"oktapam_project_groups":           dataSourceProjectGroups(),
			"oktapam_gateway_setup_tokens":     dataSourceGatewaySetupTokens(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	apiKey := d.Get("oktapam_key").(string)
	apiKeySecret := d.Get("oktapam_secret").(string)
	team := d.Get("oktapam_team").(string)
	apiHost := d.Get("oktapam_api_host").(string)

	client, err := client.CreateOktaPAMClient(apiKey, apiKeySecret, team, apiHost)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return *client, nil
}
