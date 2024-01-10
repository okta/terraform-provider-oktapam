package fwprovider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ provider.Provider = (*oktapamFrameworkProvider)(nil)

func New() func() provider.Provider {
	return func() provider.Provider {
		return &oktapamFrameworkProvider{}
	}
}

type oktapamFrameworkProvider struct{}

func OktapamFrameworkProviderSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"oktapam_api_host": schema.StringAttribute{
				Optional:    true,
				Description: "Okta PAM API Host",
			},
			"oktapam_key": schema.StringAttribute{
				Optional:    true,
				Description: "Okta PAM API Key",
			},
			"oktapam_secret": schema.StringAttribute{
				Optional:    true,
				Description: "Okta PAM API Secret",
			},
			"oktapam_team": schema.StringAttribute{
				Optional:    true,
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
