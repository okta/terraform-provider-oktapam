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

	apiHostKey      = "oktapam_api_host"
	apiKeyKey       = "oktapam_key"
	apiKeySecretKey = "oktapam_secret"
	teamKey         = "oktapam_team"

	DefaultAPIBaseURL = "https://app.scaleft.com"
)

func New() func() provider.Provider {
	return func() provider.Provider {
		return &OktapamFrameworkProvider{}
	}
}

type OktapamFrameworkProvider struct {
	SDKClientWrapper *client.SDKClientWrapper
}

type OktapamFrameworkProviderModel struct {
	OktapamApiHost types.String `tfsdk:"oktapam_api_host"`
	OktapamApiKey  types.String `tfsdk:"oktapam_key"`
	OktapamSecret  types.String `tfsdk:"oktapam_secret"`
	OktapamTeam    types.String `tfsdk:"oktapam_team"`
}

func (p *OktapamFrameworkProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			apiHostKey: schema.StringAttribute{
				Optional:    true,
				Description: "Okta PAM API Host",
			},
			apiKeyKey: schema.StringAttribute{
				Optional:    true,
				Description: "Okta PAM API Key",
			},
			apiKeySecretKey: schema.StringAttribute{
				Optional:    true,
				Description: "Okta PAM API Secret",
			},
			teamKey: schema.StringAttribute{
				Optional:    true,
				Description: "Okta PAM Team",
			},
		},
	}
}

func (p *OktapamFrameworkProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var d OktapamFrameworkProviderModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &d)...)

	if resp.Diagnostics.HasError() {
		return
	}
	diags := p.ConfigureConfigDefaults(&d)

	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
	}

	team := d.OktapamTeam.ValueString()

	config := &client.OktaPAMProviderConfig{
		APIKey:       d.OktapamApiKey.ValueString(),
		APIKeySecret: d.OktapamSecret.ValueString(),
		Team:         team,
		APIHost:      d.OktapamApiHost.ValueString(),
	}

	sdkClient, err := client.CreateSDKClient(config)

	if err != nil {
		resp.Diagnostics.Append(diag.NewErrorDiagnostic("error while creating sdk client", err.Error()))
	}

	p.SDKClientWrapper = &client.SDKClientWrapper{
		SDKClient: sdkClient,
		Team:      team,
	}
	
	resp.DataSourceData = p
	resp.ResourceData = p
}

func (p *OktapamFrameworkProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "oktapam"
}

func (p *OktapamFrameworkProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (p *OktapamFrameworkProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewResourceGroupResource,
	}
}

func (p *OktapamFrameworkProvider) ConfigureConfigDefaults(config *OktapamFrameworkProviderModel) diag.Diagnostics {
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
