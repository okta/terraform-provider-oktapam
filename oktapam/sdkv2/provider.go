package sdkv2

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/configs"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			configs.ApiHostKey: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Okta PAM API Host",
			},
			configs.ApiKeyKey: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Okta PAM API Key",
			},
			configs.ApiKeySecretKey: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Okta PAM API Secret",
			},
			configs.TeamKey: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Okta PAM Team",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			configs.ProviderADCertificateObjectKey:                resourceADCertificateObject(),
			configs.ProviderADCertificateRequestKey:               resourceADCertificateRequest(),
			configs.ProviderADConnectionKey:                       resourceADConnection(),
			configs.ProviderADTaskSettingsKey:                     resourceADServerSyncTaskSettings(),
			configs.ProviderADUserSyncTaskSettingsKey:             resourceADUserSyncTaskSettings(),
			configs.ProviderDatabaseKey:                           resourceDatabase(),
			configs.ProviderDatabasePasswordSettings:              resourceDatabasePasswordSettings(),
			configs.ProviderGatewaySetupTokenKey:                  resourceGatewaySetupToken(),
			configs.ProviderGroupKey:                              resourceGroup(),
			configs.ProviderKubernetesClusterKey:                  resourceKubernetesCluster(),
			configs.ProviderKubernetesClusterConnectionKey:        resourceKubernetesClusterConnection(),
			configs.ProviderKubernetesClusterGroupKey:             resourceKubernetesClusterGroup(),
			configs.ProviderPasswordSettingsKey:                   resourcePasswordSettings(),
			configs.ProviderProjectGroupKey:                       resourceProjectGroup(),
			configs.ProviderProjectKey:                            resourceProject(),
			configs.ProviderResourceGroupKey:                      resourceResourceGroup(),
			configs.ProviderResourceGroupProjectKey:               resourceResourceGroupProject(),
			configs.ProviderResourceGroupServerEnrollmentTokenKey: resourceResourceGroupServerEnrollmentToken(),
			configs.ProviderSecretFolderKey:                       resourceSecretFolder(),
			configs.ProviderSecurityPolicyKey:                     resourceSecurityPolicy(),
			configs.ProviderServerEnrollmentTokenKey:              resourceServerEnrollmentToken(),
			configs.ProviderTeamSettingsKey:                       resourceTeamSettings(),
			configs.ProviderUserGroupAttachmentKey:                resourceUserGroupAttachment(),
			configs.ProviderUserKey:                               resourceUser(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			configs.ProviderADConnectionsKey:                       dataSourceADConnections(),
			configs.ProviderCurrentUser:                            dataSourceCurrentUser(),
			configs.ProviderDatabaseKey:                            dataSourceDatabase(),
			configs.ProviderDatabasePasswordSettings:               dataSourceDatabasePasswordSettings(),
			configs.ProviderGatewaysKey:                            dataSourceGateways(),
			configs.ProviderGatewaySetupTokenKey:                   dataSourceGatewaySetupToken(),
			configs.ProviderGatewaySetupTokensKey:                  dataSourceGatewaySetupTokens(),
			configs.ProviderGroupKey:                               dataSourceGroup(),
			configs.ProviderGroupsKey:                              dataSourceGroups(),
			configs.ProviderPasswordSettingsKey:                    dataSourcePasswordSettings(),
			configs.ProviderProjectKey:                             dataSourceProject(),
			configs.ProviderProjectsKey:                            dataSourceProjects(),
			configs.ProviderProjectGroupKey:                        dataSourceProjectGroup(),
			configs.ProviderProjectGroupsKey:                       dataSourceProjectGroups(),
			configs.ProviderResourceGroupsKey:                      dataSourceResourceGroups(),
			configs.ProviderResourceGroupKey:                       dataSourceResourceGroup(),
			configs.ProviderResourceGroupProjectKey:                dataSourceResourceGroupProject(),
			configs.ProviderResourceGroupProjectsKey:               dataSourceResourceGroupProjects(),
			configs.ProviderResourceGroupServerEnrollmentTokenKey:  dataSourceResourceGroupServerEnrollmentToken(),
			configs.ProviderResourceGroupServerEnrollmentTokensKey: dataSourceResourceGroupServerEnrollmentTokens(),
			configs.ProviderSecretFoldersKey:                       dataSourceSecretFolders(),
			configs.ProviderSecurityPoliciesKey:                    dataSourceSecurityPolicies(),
			configs.ProviderSecurityPolicyKey:                      dataSourceSecurityPolicy(),
			configs.ProviderServerEnrollmentTokenKey:               dataSourceServerEnrollmentToken(),
			configs.ProviderServerEnrollmentTokensKey:              dataSourceServerEnrollmentTokens(),
			configs.ProviderTeamSettingsKey:                        dataSourceTeamSettings(),
			configs.ProviderADUserSyncTaskSettingsKey:              dataSourceADUserSyncTaskSettings(),
			configs.ProviderADUserSyncTaskSettingsIDListKey:        dataSourceADUserSyncTaskSettingsIDList(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
	if d.Get(configs.ApiKeyKey).(string) == "" {
		if apiKey := os.Getenv(configs.ApiKeySchemaEnvVar); apiKey != "" {
			d.Set(configs.ApiKeyKey, apiKey)
		}
	}

	if d.Get(configs.ApiHostKey).(string) == "" {
		if apiKey := os.Getenv(configs.ApiHostSchemaEnvVar); apiKey != "" {
			d.Set(configs.ApiHostKey, apiKey)
		}
	}

	if d.Get(configs.ApiKeySecretKey).(string) == "" {
		if apiKey := os.Getenv(configs.ApiKeySecretSchemaEnvVar); apiKey != "" {
			d.Set(configs.ApiKeySecretKey, apiKey)
		}
	}

	if d.Get(configs.TeamKey).(string) == "" {
		if apiKey := os.Getenv(configs.TeamSchemaEnvVar); apiKey != "" {
			d.Set(configs.TeamKey, apiKey)
		}
	}

	team := d.Get(configs.TeamKey).(string)
	config := &client.OktaPAMProviderConfig{
		APIKey:       d.Get(configs.ApiKeyKey).(string),
		APIKeySecret: d.Get(configs.ApiKeySecretKey).(string),
		Team:         team,
		APIHost:      d.Get(configs.ApiHostKey).(string),
	}

	sdkClient, err := client.CreateSDKClient(config)
	if err != nil {
		return nil, diag.Errorf("failed to load sdk api client: %v", err)
	}
	sdkClientWrapper := client.SDKClientWrapper{
		SDKClient: sdkClient,
		Team:      team,
	}

	localClient, err := client.CreateLocalPAMClient(config)
	if err != nil {
		return nil, diag.Errorf("failed to load local api client: %v", err)
	}

	return &client.APIClients{
		SDKClient:   sdkClientWrapper,
		LocalClient: localClient,
	}, nil
}
