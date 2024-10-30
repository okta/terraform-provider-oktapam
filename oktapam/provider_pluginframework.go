package oktapam

import (
	"context"
	"fmt"
	"os"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/config"

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
			config.ApiHostKey: schema.StringAttribute{
				Optional:    true,
				Description: "Okta PAM API Host",
			},
			config.ApiKeyKey: schema.StringAttribute{
				Optional:    true,
				Description: "Okta PAM API Key",
			},
			config.ApiKeySecretKey: schema.StringAttribute{
				Optional:    true,
				Description: "Okta PAM API Secret",
			},
			config.TeamKey: schema.StringAttribute{
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

func (p *FrameworkProvider) ConfigureConfigDefaults(providerModel *FrameworkProviderModel) diag.Diagnostics {
	var diags diag.Diagnostics

	if providerModel.OktapamApiKey.IsNull() {
		if apiKey := os.Getenv(config.ApiKeySchemaEnvVar); apiKey != "" {
			providerModel.OktapamApiKey = types.StringValue(apiKey)
		} else {
			diags.Append(diag.NewErrorDiagnostic("error while configuring provider",
				fmt.Sprintf("%s is not set", config.ApiKeySchemaEnvVar)))
		}
	}

	if providerModel.OktapamApiHost.IsNull() {
		if apiHost := os.Getenv(config.ApiHostSchemaEnvVar); apiHost != "" {
			providerModel.OktapamApiHost = types.StringValue(apiHost)
		} else {
			providerModel.OktapamApiHost = types.StringValue(config.DefaultAPIBaseURL)
		}
	}

	if providerModel.OktapamSecret.IsNull() {
		if apiSecret := os.Getenv(config.ApiKeySecretSchemaEnvVar); apiSecret != "" {
			providerModel.OktapamSecret = types.StringValue(apiSecret)
		} else {
			diags.Append(diag.NewErrorDiagnostic("error while configuring provider",
				fmt.Sprintf("%s is not set", config.ApiKeySchemaEnvVar)))
		}
	}

	if providerModel.OktapamTeam.IsNull() {
		if apiTeam := os.Getenv(config.TeamSchemaEnvVar); apiTeam != "" {
			providerModel.OktapamTeam = types.StringValue(apiTeam)
		} else {
			diags.Append(diag.NewErrorDiagnostic("error while configuring provider",
				fmt.Sprintf("%s is not set", config.ApiKeySchemaEnvVar)))
		}
	}

	return diags
}
