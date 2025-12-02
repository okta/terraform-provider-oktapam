package oktapam

import (
	"context"
	"os"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/config"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
)

type V5ClientCreator func(ctx context.Context, diagnostics diag.Diagnostics, data *schema.ResourceData) (*client.APIClients, diag.Diagnostics)

func V5Provider(v5ClientCreator V5ClientCreator) *schema.Provider {
	provider := &schema.Provider{
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
			config.ProviderSudoCommandBundleKey:                  resourceSudoCommandBundle(),
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
			// Do not add more SDKv2 resources. Use the terraform-plugin-framework, adding to the
			// Resources() or DataSources() methods in provider_pluginframework.go
		},

		DataSourcesMap: map[string]*schema.Resource{
			config.ProviderADConnectionsKey:                       dataSourceADConnections(),
			config.ProviderSudoCommandBundleKey:                   dataSourceSudoCommandBundle(),
			config.ProviderSudoCommandsBundlesKey:                 dataSourceSudoCommandBundles(),
			config.ProviderCurrentUser:                            dataSourceCurrentUser(),
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
			// Do not add more SDKv2 resources. Use the terraform-plugin-framework, adding to the
			// Resources() or DataSources() methods in provider_pluginframework.go
		},
	}

	if v5ClientCreator == nil {
		provider.ConfigureContextFunc = providerConfigureFactory(defaultV5ClientCreator)
	} else {
		provider.ConfigureContextFunc = providerConfigureFactory(v5ClientCreator)
	}

	return provider
}

func defaultV5ClientCreator(_ context.Context, diags diag.Diagnostics, d *schema.ResourceData) (*client.APIClients, diag.Diagnostics) {
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

func providerConfigureFactory(v5ClientCreator V5ClientCreator) func(context.Context, *schema.ResourceData) (any, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
		var diags diag.Diagnostics
		// If there was more setup for the provider, you would do it here. Instead, the V5 provider only needs
		// an SDK Client, so it's the only thing that happens.
		return v5ClientCreator(ctx, diags, d)
	}
}
