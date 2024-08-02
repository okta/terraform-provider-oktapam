package oktapam

import (
	"context"
	constantsConfig "github.com/okta/terraform-provider-oktapam/oktapam/constants/config"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func New() func() provider.Provider {
	return func() provider.Provider {
		return &FrameworkProvider{}
	}
}

type FrameworkProvider struct{}

type FrameworkProviderModel struct {
	OktapamApiHost types.String `tfsdk:"oktapam_api_host"`
	OktapamApiKey  types.String `tfsdk:"oktapam_key"`
	OktapamSecret  types.String `tfsdk:"oktapam_secret"`
	OktapamTeam    types.String `tfsdk:"oktapam_team"`
}

func (p *FrameworkProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			constantsConfig.ApiHostKey: schema.StringAttribute{
				Optional:    true,
				Description: "Okta PAM API Host",
			},
			constantsConfig.ApiKeyKey: schema.StringAttribute{
				Optional:    true,
				Description: "Okta PAM API Key",
			},
			constantsConfig.ApiKeySecretKey: schema.StringAttribute{
				Optional:    true,
				Description: "Okta PAM API Secret",
			},
			constantsConfig.TeamKey: schema.StringAttribute{
				Optional:    true,
				Description: "Okta PAM Team",
			},
		},
	}
}

func (p *FrameworkProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var d FrameworkProviderModel
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

	apiClients := &client.APIClients{
		SDKClient: client.SDKClientWrapper{
			SDKClient: sdkClient,
			Team:      team,
		},
	}

	resp.DataSourceData = apiClients
	resp.ResourceData = apiClients
}

func (p *FrameworkProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "oktapam"
}

func (p *FrameworkProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (p *FrameworkProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		//Add New Resources here
		NewServerCheckoutSettingsResource,
	}
}

func (p *FrameworkProvider) ConfigureConfigDefaults(config *FrameworkProviderModel) diag.Diagnostics {
	var diags diag.Diagnostics

	if config.OktapamApiKey.IsNull() {
		if apiKey := os.Getenv(constantsConfig.ApiKeySchemaEnvVar); apiKey != "" {
			config.OktapamApiKey = types.StringValue(apiKey)
		}
	}

	if config.OktapamApiHost.IsNull() {
		if apiHost := os.Getenv(constantsConfig.ApiHostSchemaEnvVar); apiHost != "" {
			config.OktapamApiHost = types.StringValue(apiHost)
		} else {
			config.OktapamApiHost = types.StringValue(constantsConfig.DefaultAPIBaseURL)
		}
	}

	if config.OktapamSecret.IsNull() {
		if apiSecret := os.Getenv(constantsConfig.ApiKeySecretSchemaEnvVar); apiSecret != "" {
			config.OktapamSecret = types.StringValue(apiSecret)
		}
	}

	if config.OktapamTeam.IsNull() {
		if apiTeam := os.Getenv(constantsConfig.TeamSchemaEnvVar); apiTeam != "" {
			config.OktapamTeam = types.StringValue(apiTeam)
		}
	}

	return diags
}
