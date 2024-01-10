package fwprovider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

const (
	apiHostSchemaEnvVar      = "OKTAPAM_API_HOST"
	apiKeySchemaEnvVar       = "OKTAPAM_KEY"
	apiKeySecretSchemaEnvVar = "OKTAPAM_SECRET"
	teamSchemaEnvVar         = "OKTAPAM_TEAM"
	DefaultAPIBaseURL        = "https://app.scaleft.com"
)

var _ provider.Provider = (*oktapamFrameworkProvider)(nil)

func New() func() provider.Provider {
	return func() provider.Provider {
		return &oktapamFrameworkProvider{}
	}
}

type oktapamFrameworkProvider struct {
	SDKClientWrapper *client.SDKClientWrapper
}

func OktapamFrameworkProviderSchema(_ context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"oktapam_api_host": schema.StringAttribute{
				Optional:    true,
				Description: "Okta PAM API Host",
			},
			"oktapam_key": schema.StringAttribute{
				Required:    true,
				Description: "Okta PAM API Key",
			},
			"oktapam_secret": schema.StringAttribute{
				Required:    true,
				Description: "Okta PAM API Secret",
			},
			"oktapam_team": schema.StringAttribute{
				Required:    true,
				Description: "Okta PAM Team",
			},
		},
	}
}

type OktapamFrameworkProviderModel struct {
	OktapamApiHost types.String `tfsdk:"oktapam_api_host"`
	OktapamApiKey  types.String `tfsdk:"oktapam_api_key"`
	OktapamSecret  types.String `tfsdk:"oktapam_secret"`
	OktapamTeam    types.String `tfsdk:"oktapam_team"`
}

func (p *oktapamFrameworkProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = OktapamFrameworkProviderSchema(ctx)
}

func (p *oktapamFrameworkProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

	var d OktapamFrameworkProviderModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &d)...)

	if resp.Diagnostics.HasError() {
		return
	}
	p.ConfigureConfigDefaults(&d)
	team := d.OktapamTeam.String()

	config := &client.OktaPAMProviderConfig{
		APIKey:       d.OktapamApiKey.String(),
		APIKeySecret: d.OktapamSecret.String(),
		Team:         team,
		APIHost:      d.OktapamApiHost.String(),
	}

	sdkClient, _ := client.CreateSDKClient(config)

	p.SDKClientWrapper = &client.SDKClientWrapper{
		SDKClient: sdkClient,
		Team:      team,
	}

	resp.DataSourceData = p
	resp.ResourceData = p

}

func (p *oktapamFrameworkProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "oktapam_framework_provider"
}

func (p *oktapamFrameworkProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (p *oktapamFrameworkProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *oktapamFrameworkProvider) ConfigureConfigDefaults(config *OktapamFrameworkProviderModel) diag.Diagnostics {
	var diags diag.Diagnostics

	if config.OktapamApiKey.IsNull() {
		if apiKey := os.Getenv(apiKeySchemaEnvVar); apiKey != "" {
			config.OktapamApiKey = types.StringValue(apiKey)
		}
	}

	if config.OktapamApiHost.IsNull() {
		if apiHost := os.Getenv(apiHostSchemaEnvVar); apiHost != "" {
			config.OktapamApiHost = types.StringValue(apiHost)
		} else {
			config.OktapamApiHost = types.StringValue(DefaultAPIBaseURL)
		}
	}

	if config.OktapamSecret.IsNull() {
		if apiSecret := os.Getenv(apiKeySecretSchemaEnvVar); apiSecret != "" {
			config.OktapamSecret = types.StringValue(apiSecret)
		}
	}

	if config.OktapamTeam.IsNull() {
		if apiTeam := os.Getenv(teamSchemaEnvVar); apiTeam != "" {
			config.OktapamTeam = types.StringValue(apiTeam)
		}
	}

	return diags
}
