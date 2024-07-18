package fwprovider

import (
	"context"
	"fmt"

	"github.com/okta/terraform-provider-oktapam/oktapam/client"

	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource              = &serverCheckoutSettingsResource{}
	_ resource.ResourceWithConfigure = &serverCheckoutSettingsResource{}
)

func NewServerCheckoutSettingsResource() resource.Resource {
	return &serverCheckoutSettingsResource{}
}

type serverCheckoutSettingsResource struct {
	client *client.SDKClientWrapper
}

func (r *serverCheckoutSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve the server checkout settings values from the plan
	var plan serverCheckoutSettingsResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Call the SDK client to creat the server checkout settings from plan
	serverCheckoutSettings := &pam.ResourceCheckoutSettings{
		CheckoutRequired:          plan.CheckoutRequired,
		CheckoutDurationInSeconds: plan.CheckoutDurationInSeconds,
		IncludeList:               plan.IncludeList,
		ExcludeList:               plan.ExcludeList,
	}

	_, err := r.client.SDKClient.ProjectsAPI.UpdateResourceGroupServerBasedProjectCheckoutSettings(ctx, r.client.Team, plan.ResourceGroupId, plan.ProjectId).ResourceCheckoutSettings(*serverCheckoutSettings).Execute()

	if err != nil {
		resp.Diagnostics.AddError("Error creating server checkout settings", err.Error())
		return
	}
	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

type serverCheckoutSettingsResourceModel struct {
	ResourceGroupId           string   `json:"resource_group_id"`
	ProjectId                 string   `json:"project_id"`
	CheckoutRequired          bool     `json:"checkout_required"`
	CheckoutDurationInSeconds *int32   `json:"checkout_duration_in_seconds"`
	IncludeList               []string `json:"include_list"`
	ExcludeList               []string `json:"exclude_list"`
}

// Metadata implements resource.Resource.
func (r *serverCheckoutSettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_server_checkout_settings"
}

// Read implements resource.Resource.
func (r *serverCheckoutSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state serverCheckoutSettingsResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed order value from HashiCups
	serverCheckoutSettings, _, err := r.client.SDKClient.ProjectsAPI.FetchResourceGroupServerBasedProjectCheckoutSettings(ctx, r.client.Team, state.ResourceGroupId, state.ProjectId).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading server checkout settings",
			"Could not read server checkout settings for team:"+r.client.Team+"resource_group:"+state.ResourceGroupId+"project_id:"+state.ProjectId+": "+err.Error(),
		)
		return
	}

	// Overwrite items with refreshed state
	state.CheckoutRequired = serverCheckoutSettings.CheckoutRequired
	state.CheckoutDurationInSeconds = serverCheckoutSettings.CheckoutDurationInSeconds
	state.IncludeList = serverCheckoutSettings.IncludeList
	state.ExcludeList = serverCheckoutSettings.ExcludeList

	// Set state to fully populated data
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Schema implements resource.Resource.
func (r *serverCheckoutSettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"checkout_required": schema.BoolAttribute{
				Required: true,
			},
			"checkout_duration_in_seconds": schema.Int64Attribute{
				Optional: true,
			},
			"include_list": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},
			"exclude_list": schema.ListAttribute{

				Optional:    true,
				ElementType: types.StringType,
			},
		},
	}
}

// Update implements resource.Resource.
func (r *serverCheckoutSettingsResource) Update(context.Context, resource.UpdateRequest, *resource.UpdateResponse) {

}

// Delete implements resource.Resource.
func (r *serverCheckoutSettingsResource) Delete(context.Context, resource.DeleteRequest, *resource.DeleteResponse) {
	panic("unimplemented")
}

// Configure adds the provider configured client to the resource.
func (r *serverCheckoutSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.SDKClientWrapper)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *client.SDKClientWrapper, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func formatServerCheckoutSettingsID(resourceGroupID string, projectID string) string {
	// password settings don't have an identifier in itself and is really an attribute of a project.
	// we manage it as a separate resource since it's lifecycle is somewhat separate from a project
	return fmt.Sprintf("%s/%s", resourceGroupID, projectID)
}
