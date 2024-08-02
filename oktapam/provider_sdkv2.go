package oktapam

import (
	"context"
	"os"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/config"

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

func providerConfigure(_ context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
	var diags diag.Diagnostics

	if d.Get(config.ApiKeyKey).(string) == "" {
		if apiKey := os.Getenv(config.ApiKeySchemaEnvVar); apiKey != "" {
			if err := d.Set(config.ApiKeyKey, apiKey); err != nil {
				diags = append(diags, diag.FromErr(err)...)
			}
		} else {
			diags = append(diags, diag.Errorf("%s is not set", config.ApiKeySchemaEnvVar)...)
		}
	}

	if d.Get(config.ApiKeySecretKey).(string) == "" {
		if apiSecret := os.Getenv(config.ApiKeySecretSchemaEnvVar); apiSecret != "" {
			if err := d.Set(config.ApiKeySecretKey, apiSecret); err != nil {
				diags = append(diags, diag.FromErr(err)...)
			}
		} else {
			diags = append(diags, diag.Errorf("%s is not set", config.ApiKeySecretSchemaEnvVar)...)
		}
	}

	if d.Get(config.ApiHostKey).(string) == "" {
		if apiHost := os.Getenv(config.ApiHostSchemaEnvVar); apiHost != "" {
			if err := d.Set(config.ApiHostKey, apiHost); err != nil {
				diags = append(diags, diag.FromErr(err)...)
			}
		} else {
			if err := d.Set(config.ApiHostKey, config.DefaultAPIBaseURL); err != nil {
				diags = append(diags, diag.FromErr(err)...)
			}
		}
	}

	if d.Get(config.TeamKey).(string) == "" {
		if team := os.Getenv(config.TeamSchemaEnvVar); team != "" {
			if err := d.Set(config.TeamKey, team); err != nil {
				diags = append(diags, diag.FromErr(err)...)
			}
		} else {
			diags = append(diags, diag.Errorf("%s is not set", config.TeamSchemaEnvVar)...)
		}
	}

	if diags.HasError() {
		return nil, diags
	}

	team := d.Get(config.TeamKey).(string)
	cfg := &client.OktaPAMProviderConfig{
		APIKey:       d.Get(config.ApiKeyKey).(string),
		APIKeySecret: d.Get(config.ApiKeySecretKey).(string),
		Team:         team,
		APIHost:      d.Get(config.ApiHostKey).(string),
	}

	if apiClients, err := NewAPIClients(cfg); err == nil {
		return apiClients, nil
	} else {
		return nil, diag.Errorf("failed to get api clients: %v", err)
	}
}
