package oktapam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
)

const (
	DefaultAPIBaseURL = "https://app.scaleft.com"

	apiHostSchemaEnvVar      = "OKTAPAM_API_HOST"
	apiKeySchemaEnvVar       = "OKTAPAM_KEY"
	apiKeySecretSchemaEnvVar = "OKTAPAM_SECRET"
	teamSchemaEnvVar         = "OKTAPAM_TEAM"

	apiHostKey      = "oktapam_api_host"
	apiKeyKey       = "oktapam_key"
	apiKeySecretKey = "oktapam_secret"
	teamKey         = "oktapam_team"

	providerProjectKey               = "oktapam_project"
	providerGroupKey                 = "oktapam_group"
	providerServerEnrollmentTokenKey = "oktapam_server_enrollment_token"
	providerProjectGroupKey          = "oktapam_project_group"
	providerGatewaySetupTokenKey     = "oktapam_gateway_setup_token"
	providerADConnectionKey          = "oktapam_ad_connection"
	providerADTaskSettingsKey        = "oktapam_ad_task_settings"
	providerGatewayKey               = "oktapam_gateway"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			apiHostKey: {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(apiHostSchemaEnvVar, DefaultAPIBaseURL),
				Description: "Okta PAM API Host",
			},
			apiKeyKey: {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc(apiKeySchemaEnvVar, nil),
				Description: "Okta PAM API Key",
			},
			apiKeySecretKey: {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc(apiKeySecretSchemaEnvVar, nil),
				Description: "Okta PAM API Secret",
			},
			teamKey: {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc(teamSchemaEnvVar, nil),
				Description: "Okta PAM Team",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			providerProjectKey:               resourceProject(),
			providerGroupKey:                 resourceGroup(),
			providerServerEnrollmentTokenKey: resourceServerEnrollmentToken(),
			providerProjectGroupKey:          resourceProjectGroup(),
			providerGatewaySetupTokenKey:     resourceGatewaySetupToken(),
			providerADConnectionKey:          resourceADConnection(),
			providerADTaskSettingsKey:        resourceADTaskSettings(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			providerProjectKey:               dataSourceProjects(),
			providerGroupKey:                 dataSourceGroups(),
			providerServerEnrollmentTokenKey: dataSourceServerEnrollmentTokens(),
			providerProjectGroupKey:          dataSourceProjectGroups(),
			providerGatewaySetupTokenKey:     dataSourceGatewaySetupTokens(),
			providerADConnectionKey:          dataSourceADConnections(),
			providerGatewayKey:               dataSourceGateways(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	apiKey := d.Get(apiKeyKey).(string)
	apiKeySecret := d.Get(apiKeySecretKey).(string)
	team := d.Get(teamKey).(string)
	apiHost := d.Get(apiHostKey).(string)

	client, err := client.CreateOktaPAMClient(apiKey, apiKeySecret, team, apiHost)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return *client, nil
}
