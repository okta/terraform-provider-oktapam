package oktapam

import (
	"context"
	"fmt"

	"github.com/atko-pam/pam-sdk-go/client/pam"
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

	providerADCertificateObjectKey                 = "oktapam_ad_certificate_object"
	providerADCertificateRequestKey                = "oktapam_ad_certificate_request"
	providerADConnectionKey                        = "oktapam_ad_connection"
	providerADConnectionsKey                       = "oktapam_ad_connections"
	providerADTaskSettingsKey                      = "oktapam_ad_task_settings"
	providerADUserSyncTaskSettingsKey              = "oktapam_ad_user_sync_task_settings"
	providerADUserSyncTaskSettingsIDListKey        = "oktapam_ad_user_sync_task_settings_id_list"
	providerGatewaysKey                            = "oktapam_gateways"
	providerGatewaySetupTokenKey                   = "oktapam_gateway_setup_token"
	providerGatewaySetupTokensKey                  = "oktapam_gateway_setup_tokens"
	providerGroupKey                               = "oktapam_group"
	providerGroupsKey                              = "oktapam_groups"
	providerKubernetesClusterKey                   = "oktapam_kubernetes_cluster"
	providerKubernetesClusterConnectionKey         = "oktapam_kubernetes_cluster_connection"
	providerKubernetesClusterGroupKey              = "oktapam_kubernetes_cluster_group"
	providerPasswordSettingsKey                    = "oktapam_password_settings"
	providerProjectKey                             = "oktapam_project"
	providerProjectsKey                            = "oktapam_projects"
	providerProjectGroupKey                        = "oktapam_project_group"
	providerProjectGroupsKey                       = "oktapam_project_groups"
	providerResourceGroupKey                       = "oktapam_resource_group"
	providerResourceGroupsKey                      = "oktapam_resource_groups"
	providerResourceGroupProjectKey                = "oktapam_resource_group_project"
	providerResourceGroupProjectsKey               = "oktapam_resource_group_projects"
	providerResourceGroupServerEnrollmentTokenKey  = "oktapam_resource_group_server_enrollment_token"
	providerResourceGroupServerEnrollmentTokensKey = "oktapam_resource_group_server_enrollment_tokens"
	providerSecurityPoliciesKey                    = "oktapam_security_policies"
	providerSecurityPolicyKey                      = "oktapam_security_policy"
	providerServerEnrollmentTokenKey               = "oktapam_server_enrollment_token"
	providerServerEnrollmentTokensKey              = "oktapam_server_enrollment_tokens"
	providerTeamSettingsKey                        = "oktapam_team_settings"
	providerUserKey                                = "oktapam_user"
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
			providerADCertificateObjectKey:                resourceADCertificateObject(),
			providerADCertificateRequestKey:               resourceADCertificateRequest(),
			providerADConnectionKey:                       resourceADConnection(),
			providerADTaskSettingsKey:                     resourceADServerSyncTaskSettings(),
			providerADUserSyncTaskSettingsKey:             resourceADUserSyncTaskSettings(),
			providerGatewaySetupTokenKey:                  resourceGatewaySetupToken(),
			providerGroupKey:                              resourceGroup(),
			providerKubernetesClusterKey:                  resourceKubernetesCluster(),
			providerKubernetesClusterConnectionKey:        resourceKubernetesClusterConnection(),
			providerKubernetesClusterGroupKey:             resourceKubernetesClusterGroup(),
			providerPasswordSettingsKey:                   resourcePasswordSettings(),
			providerProjectGroupKey:                       resourceProjectGroup(),
			providerProjectKey:                            resourceProject(),
			providerResourceGroupKey:                      resourceResourceGroup(),
			providerResourceGroupProjectKey:               resourceResourceGroupProject(),
			providerResourceGroupServerEnrollmentTokenKey: resourceResourceGroupServerEnrollmentToken(),
			providerSecurityPolicyKey:                     resourceSecurityPolicy(),
			providerServerEnrollmentTokenKey:              resourceServerEnrollmentToken(),
			providerTeamSettingsKey:                       resourceTeamSettings(),
			providerUserKey:                               resourceUser(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			providerADConnectionsKey:                       dataSourceADConnections(),
			providerGatewaysKey:                            dataSourceGateways(),
			providerGatewaySetupTokenKey:                   dataSourceGatewaySetupToken(),
			providerGatewaySetupTokensKey:                  dataSourceGatewaySetupTokens(),
			providerGroupKey:                               dataSourceGroup(),
			providerGroupsKey:                              dataSourceGroups(),
			providerPasswordSettingsKey:                    dataSourcePasswordSettings(),
			providerProjectKey:                             dataSourceProject(),
			providerProjectsKey:                            dataSourceProjects(),
			providerProjectGroupKey:                        dataSourceProjectGroup(),
			providerProjectGroupsKey:                       dataSourceProjectGroups(),
			providerResourceGroupsKey:                      dataSourceResourceGroups(),
			providerResourceGroupKey:                       dataSourceResourceGroup(),
			providerResourceGroupProjectKey:                dataSourceResourceGroupProject(),
			providerResourceGroupProjectsKey:               dataSourceResourceGroupProjects(),
			providerResourceGroupServerEnrollmentTokenKey:  dataSourceResourceGroupServerEnrollmentToken(),
			providerResourceGroupServerEnrollmentTokensKey: dataSourceResourceGroupServerEnrollmentTokens(),
			providerSecurityPoliciesKey:                    dataSourceSecurityPolicies(),
			providerSecurityPolicyKey:                      dataSourceSecurityPolicy(),
			providerServerEnrollmentTokenKey:               dataSourceServerEnrollmentToken(),
			providerServerEnrollmentTokensKey:              dataSourceServerEnrollmentTokens(),
			providerTeamSettingsKey:                        dataSourceTeamSettings(),
			providerADUserSyncTaskSettingsKey:              dataSourceADUserSyncTaskSettings(),
			providerADUserSyncTaskSettingsIDListKey:        dataSourceADUserSyncTaskSettingsIDList(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
	config := &OktaPAMProviderConfig{
		APIKey:       d.Get(apiKeyKey).(string),
		APIKeySecret: d.Get(apiKeySecretKey).(string),
		Team:         d.Get(teamKey).(string),
		APIHost:      d.Get(apiHostKey).(string),
	}

	sdkClient, err := createSDKClient(config)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	localClient, err := createLocalClient(config)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return &clients{
		sdkClient:   sdkClient,
		localClient: localClient,
	}, nil
}

type OktaPAMProviderConfig struct {
	APIKey       string
	APIKeySecret string
	Team         string
	APIHost      string
}

type clients struct {
	sdkClient   *pam.APIClient
	localClient *client.OktaPAMClient
}

func createSDKClient(providerConfig *OktaPAMProviderConfig) (*pam.APIClient, error) {
	apiClientConfigOpts := []pam.ConfigOption{
		pam.WithHost(providerConfig.APIHost),
		pam.WithTeam(providerConfig.Team),
		pam.WithAPIKey(providerConfig.APIKey),
		pam.WithAPISecret(providerConfig.APIKeySecret),
	}

	pamClient, err := pam.NewAPIClient(apiClientConfigOpts...)
	if err != nil {
		return nil, fmt.Errorf("error while creating sdk client: %w", err)
	}
	return pamClient, nil
}

func createLocalClient(providerConfig *OktaPAMProviderConfig) (*client.OktaPAMClient, error) {
	return client.CreateOktaPAMClient(providerConfig.APIKey, providerConfig.APIKeySecret, providerConfig.Team, providerConfig.APIHost)
}

func getSDKClientFromMetadata(meta interface{}) *pam.APIClient {
	return meta.(*clients).sdkClient
}

func getLocalClientFromMetadata(meta interface{}) *client.OktaPAMClient {
	return meta.(*clients).localClient
}