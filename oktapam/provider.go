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
	providerADCertificateRequestKey        = "oktapam_ad_certificate_request"
	providerADConnectionKey                = "oktapam_ad_connection"
	providerADConnectionsKey               = "oktapam_ad_connections"
	providerADTaskSettingsKey              = "oktapam_ad_task_settings"
	providerGatewaysKey                    = "oktapam_gateways"
	providerGatewaySetupTokenKey           = "oktapam_gateway_setup_token"
	providerGatewaySetupTokensKey          = "oktapam_gateway_setup_tokens"
	providerGroupKey                       = "oktapam_group"
	providerGroupsKey                      = "oktapam_groups"
	providerKubernetesClusterKey           = "oktapam_kubernetes_cluster"
	providerKubernetesClusterConnectionKey = "oktapam_kubernetes_cluster_connection"
	providerKubernetesClusterGroupKey      = "oktapam_kubernetes_cluster_group"
	providerProjectKey                     = "oktapam_project"
	providerProjectsKey                    = "oktapam_projects"
	providerProjectGroupKey                = "oktapam_project_group"
	providerProjectGroupsKey               = "oktapam_project_groups"
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
			providerADCertificateRequestKey:        resourceADCertificateRequest(),
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
			providerADConnectionsKey:          dataSourceADConnections(),
			providerGatewaysKey:               dataSourceGateways(),
			providerGatewaySetupTokenKey:      dataSourceGatewaySetupToken(),
			providerGatewaySetupTokensKey:     dataSourceGatewaySetupTokens(),
			providerGroupKey:                  dataSourceGroup(),
			providerGroupsKey:                 dataSourceGroups(),
			providerProjectKey:                dataSourceProject(),
			providerProjectsKey:               dataSourceProjects(),
			providerProjectGroupKey:           dataSourceProjectGroup(),
			providerProjectGroupsKey:          dataSourceProjectGroups(),
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
