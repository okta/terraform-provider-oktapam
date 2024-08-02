package oktapam

import (
	"context"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/config"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			config.ApiHostKey: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Okta PAM API Host",
			},
			config.ApiKeyKey: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Okta PAM API Key",
			},
			config.ApiKeySecretKey: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Okta PAM API Secret",
			},
			config.TeamKey: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Okta PAM Team",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			config.ProviderADCertificateObjectKey:                resourceADCertificateObject(),
			config.ProviderADCertificateRequestKey:               resourceADCertificateRequest(),
			config.ProviderADConnectionKey:                       resourceADConnection(),
			config.ProviderADTaskSettingsKey:                     resourceADServerSyncTaskSettings(),
			config.ProviderADUserSyncTaskSettingsKey:             resourceADUserSyncTaskSettings(),
			config.ProviderCloudConnectionKey:                    resourceCloudConnection(),
			config.ProviderSudoCommandBundleKey:                  resourceSudoCommandBundle(),
			config.ProviderDatabaseKey:                           resourceDatabase(),
			config.ProviderDatabasePasswordSettings:              resourceDatabasePasswordSettings(),
			config.ProviderGatewaySetupTokenKey:                  resourceGatewaySetupToken(),
			config.ProviderGroupKey:                              resourceGroup(),
			config.ProviderKubernetesClusterKey:                  resourceKubernetesCluster(),
			config.ProviderKubernetesClusterConnectionKey:        resourceKubernetesClusterConnection(),
			config.ProviderKubernetesClusterGroupKey:             resourceKubernetesClusterGroup(),
			config.ProviderPasswordSettingsKey:                   resourcePasswordSettings(),
			config.ProviderProjectGroupKey:                       resourceProjectGroup(),
			config.ProviderProjectKey:                            resourceProject(),
			config.ProviderResourceGroupKey:                      resourceResourceGroup(),
			config.ProviderResourceGroupProjectKey:               resourceResourceGroupProject(),
			config.ProviderResourceGroupServerEnrollmentTokenKey: resourceResourceGroupServerEnrollmentToken(),
			config.ProviderSecretKey:                             resourceSecret(),
			config.ProviderSecretFolderKey:                       resourceSecretFolder(),
			config.ProviderSecurityPolicyKey:                     resourceSecurityPolicy(),
			config.ProviderServerEnrollmentTokenKey:              resourceServerEnrollmentToken(),
			config.ProviderTeamSettingsKey:                       resourceTeamSettings(),
			config.ProviderUserGroupAttachmentKey:                resourceUserGroupAttachment(),
			config.ProviderUserKey:                               resourceUser(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			config.ProviderADConnectionsKey:                       dataSourceADConnections(),
			config.ProviderCloudConnectionKey:                     dataSourceCloudConnection(),
			config.ProviderCloudConnectionsKey:                    dataSourceCloudConnections(),
			config.ProviderSudoCommandBundleKey:                   dataSourceSudoCommandBundle(),
			config.ProviderSudoCommandsBundlesKey:                 dataSourceSudoCommandBundles(),
			config.ProviderCurrentUser:                            dataSourceCurrentUser(),
			config.ProviderDatabaseKey:                            dataSourceDatabase(),
			config.ProviderDatabasePasswordSettings:               dataSourceDatabasePasswordSettings(),
			config.ProviderGatewaysKey:                            dataSourceGateways(),
			config.ProviderGatewaySetupTokenKey:                   dataSourceGatewaySetupToken(),
			config.ProviderGatewaySetupTokensKey:                  dataSourceGatewaySetupTokens(),
			config.ProviderGroupKey:                               dataSourceGroup(),
			config.ProviderGroupsKey:                              dataSourceGroups(),
			config.ProviderPasswordSettingsKey:                    dataSourcePasswordSettings(),
			config.ProviderProjectKey:                             dataSourceProject(),
			config.ProviderProjectsKey:                            dataSourceProjects(),
			config.ProviderProjectGroupKey:                        dataSourceProjectGroup(),
			config.ProviderProjectGroupsKey:                       dataSourceProjectGroups(),
			config.ProviderResourceGroupsKey:                      dataSourceResourceGroups(),
			config.ProviderResourceGroupKey:                       dataSourceResourceGroup(),
			config.ProviderResourceGroupProjectKey:                dataSourceResourceGroupProject(),
			config.ProviderResourceGroupProjectsKey:               dataSourceResourceGroupProjects(),
			config.ProviderResourceGroupServerEnrollmentTokenKey:  dataSourceResourceGroupServerEnrollmentToken(),
			config.ProviderResourceGroupServerEnrollmentTokensKey: dataSourceResourceGroupServerEnrollmentTokens(),
			config.ProviderSecretFoldersKey:                       dataSourceSecretFolders(),
			config.ProviderSecretKey:                              dataSourceSecret(),
			config.ProviderSecretsKey:                             dataSourceSecrets(),
			config.ProviderSecurityPoliciesKey:                    dataSourceSecurityPolicies(),
			config.ProviderSecurityPolicyKey:                      dataSourceSecurityPolicy(),
			config.ProviderServerEnrollmentTokenKey:               dataSourceServerEnrollmentToken(),
			config.ProviderServerEnrollmentTokensKey:              dataSourceServerEnrollmentTokens(),
			config.ProviderTeamSettingsKey:                        dataSourceTeamSettings(),
			config.ProviderADUserSyncTaskSettingsKey:              dataSourceADUserSyncTaskSettings(),
			config.ProviderADUserSyncTaskSettingsIDListKey:        dataSourceADUserSyncTaskSettingsIDList(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
	if d.Get(config.ApiKeyKey).(string) == "" {
		if apiKey := os.Getenv(config.ApiKeySchemaEnvVar); apiKey != "" {
			d.Set(config.ApiKeyKey, apiKey)
		}
	}

	if d.Get(config.ApiHostKey).(string) == "" {
		if apiKey := os.Getenv(config.ApiHostSchemaEnvVar); apiKey != "" {
			d.Set(config.ApiHostKey, apiKey)
		}
	}

	if d.Get(config.ApiKeySecretKey).(string) == "" {
		if apiKey := os.Getenv(config.ApiKeySecretSchemaEnvVar); apiKey != "" {
			d.Set(config.ApiKeySecretKey, apiKey)
		}
	}

	if d.Get(config.TeamKey).(string) == "" {
		if apiKey := os.Getenv(config.TeamSchemaEnvVar); apiKey != "" {
			d.Set(config.TeamKey, apiKey)
		}
	}

	team := d.Get(config.TeamKey).(string)
	config := &client.OktaPAMProviderConfig{
		APIKey:       d.Get(config.ApiKeyKey).(string),
		APIKeySecret: d.Get(config.ApiKeySecretKey).(string),
		Team:         team,
		APIHost:      d.Get(config.ApiHostKey).(string),
	}

	if apiClients, err := NewAPIClients(config); err == nil {
		return nil, diag.Errorf("failed to get api clients: %v", err)
	} else {
		return apiClients, nil
	}
}
