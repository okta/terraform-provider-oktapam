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

	providerADCertificateObjectKey         = "oktapam_ad_certificate_object"
	providerADCertificateSigningRequestKey = "oktapam_ad_certificate_signing_request"
	providerADConnectionKey                = "oktapam_ad_connection"
	providerADTaskSettingsKey              = "oktapam_ad_task_settings"
	providerGatewayKey                     = "oktapam_gateway"
	providerGatewaySetupTokenKey           = "oktapam_gateway_setup_token"
	providerGatewaySetupTokensKey          = "oktapam_gateway_setup_tokens"
	providerGroupKey                       = "oktapam_group"
	providerKubernetesClusterKey           = "oktapam_kubernetes_cluster"
	providerKubernetesClusterConnectionKey = "oktapam_kubernetes_cluster_connection"
	providerKubernetesClusterGroupKey      = "oktapam_kubernetes_cluster_group"
	providerProjectKey                     = "oktapam_project"
	providerProjectGroupKey                = "oktapam_project_group"
	providerServerEnrollmentTokenKey       = "oktapam_server_enrollment_token"
	providerServerEnrollmentTokensKey      = "oktapam_server_enrollment_tokens"
	providerUserKey                        = "oktapam_user"
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
			providerADCertificateObjectKey:         resourceADCertificateObject(),
			providerADCertificateSigningRequestKey: resourceADCertificateSigningRequest(),
			providerADConnectionKey:                resourceADConnection(),
			providerADTaskSettingsKey:              resourceADTaskSettings(),
			providerGatewaySetupTokenKey:           resourceGatewaySetupToken(),
			providerGroupKey:                       resourceGroup(),
			providerKubernetesClusterKey:           resourceKubernetesCluster(),
			providerKubernetesClusterConnectionKey: resourceKubernetesClusterConnection(),
			providerKubernetesClusterGroupKey:      resourceKubernetesClusterGroup(),
			providerProjectGroupKey:                resourceProjectGroup(),
			providerProjectKey:                     resourceProject(),
			providerServerEnrollmentTokenKey:       resourceServerEnrollmentToken(),
			providerUserKey:                        resourceUser(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			providerADConnectionKey:           dataSourceADConnections(),
			providerGatewayKey:                dataSourceGateways(),
			providerGatewaySetupTokenKey:      dataSourceGatewaySetupToken(),
			providerGatewaySetupTokensKey:     dataSourceGatewaySetupTokens(),
			providerGroupKey:                  dataSourceGroups(),
			providerProjectKey:                dataSourceProjects(),
			providerProjectGroupKey:           dataSourceProjectGroups(),
			providerServerEnrollmentTokenKey:  dataSourceServerEnrollmentToken(),
			providerServerEnrollmentTokensKey: dataSourceServerEnrollmentTokens(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	apiKey := d.Get(apiKeyKey).(string)
	apiKeySecret := d.Get(apiKeySecretKey).(string)
	team := d.Get(teamKey).(string)
	apiHost := d.Get(apiHostKey).(string)

	pamClient, err := client.CreateOktaPAMClient(apiKey, apiKeySecret, team, apiHost)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return *pamClient, nil
}
