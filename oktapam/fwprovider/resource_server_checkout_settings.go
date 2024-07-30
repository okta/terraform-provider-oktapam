package fwprovider

import (
	"context"
	"fmt"

	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"

	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
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

type serverCheckoutSettingsResourceModel struct {
	Id                        types.String `tfsdk:"id"`
	ResourceGroup             string       `tfsdk:"resource_group"`
	Project                   string       `tfsdk:"project"`
	CheckoutRequired          bool         `tfsdk:"checkout_required"`
	CheckoutDurationInSeconds *int32       `tfsdk:"checkout_duration_in_seconds"`
	IncludeList               []string     `tfsdk:"include_list"`
	ExcludeList               []string     `tfsdk:"exclude_list"`
}

// Metadata implements resource.Resource.
func (r *serverCheckoutSettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_server_checkout_settings"
}

// Schema implements resource.Resource.
func (r *serverCheckoutSettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: descriptions.ResourceServerCheckoutSettings,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"checkout_duration_in_seconds": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         descriptions.ServerCheckoutDurationInSeconds,
				MarkdownDescription: descriptions.ServerCheckoutDurationInSeconds,
				Validators: []validator.Int64{
					int64validator.Between(900, 86400),
				},
			},
			"checkout_required": schema.BoolAttribute{
				Required:            true,
				Description:         descriptions.ServerCheckoutRequired,
				MarkdownDescription: descriptions.ServerCheckoutRequired,
			},
			"exclude_list": schema.ListAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Description:         descriptions.ServerCheckoutExcludeList,
				MarkdownDescription: descriptions.ServerCheckoutExcludeList,
			},
			"include_list": schema.ListAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Description:         descriptions.ServerCheckoutIncludeList,
				MarkdownDescription: descriptions.ServerCheckoutIncludeList,
			},
			"project": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description:         descriptions.ProjectID,
				MarkdownDescription: descriptions.ProjectID,
			},
			"resource_group": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description:         descriptions.ResourceGroupID,
				MarkdownDescription: descriptions.ResourceGroupID,
			},
		},
	}
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

	err := client.UpdateServerCheckoutSettings(ctx, *r.client, plan.ResourceGroup, plan.Project, *serverCheckoutSettings)

	if err != nil {
		resp.Diagnostics.AddError("Error creating server checkout settings", err.Error())
		return
	}
	// Fetch the created server checkout settings from the API host
	createdServerCheckoutSettings, err := client.GetServerCheckoutSettings(ctx, *r.client, plan.ResourceGroup, plan.Project)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading server checkout settings",
			"Could not read server checkout settings for team:"+r.client.Team+"resource_group:"+plan.ResourceGroup+"project_id:"+plan.Project+": "+err.Error(),
		)
		return
	}
	// Set state to fully populated data
	plan.Id = types.StringValue(FormatServerCheckoutSettingsID(plan.ResourceGroup, plan.Project))
	plan.CheckoutRequired = createdServerCheckoutSettings.CheckoutRequired
	plan.CheckoutDurationInSeconds = createdServerCheckoutSettings.CheckoutDurationInSeconds
	plan.IncludeList = createdServerCheckoutSettings.IncludeList
	plan.ExcludeList = createdServerCheckoutSettings.ExcludeList
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
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

	// Get refreshed server checkout settings from API host
	serverCheckoutSettings, err := client.GetServerCheckoutSettings(ctx, *r.client, state.ResourceGroup, state.Project)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading server checkout settings",
			"Could not read server checkout settings for team:"+r.client.Team+"resource_group:"+state.ResourceGroup+"project_id:"+state.Project+": "+err.Error(),
		)
		return
	}

	// Overwrite server checkout settings with refreshed state
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

// Update implements resource.Resource.
func (r *serverCheckoutSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve the server checkout settings values from the plan
	var plan serverCheckoutSettingsResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Update the existing server checkout settings with the new values
	serverCheckoutSettings := &pam.ResourceCheckoutSettings{
		CheckoutRequired:          plan.CheckoutRequired,
		CheckoutDurationInSeconds: plan.CheckoutDurationInSeconds,
		IncludeList:               plan.IncludeList,
		ExcludeList:               plan.ExcludeList,
	}

	err := client.UpdateServerCheckoutSettings(ctx, *r.client, plan.ResourceGroup, plan.Project, *serverCheckoutSettings)

	if err != nil {
		resp.Diagnostics.AddError("Error updating server checkout settings", err.Error())
		return
	}

	// Fetch the updated server checkout settings from the API host
	updatedServerCheckoutSettings, err := client.GetServerCheckoutSettings(ctx, *r.client, plan.ResourceGroup, plan.Project)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading server checkout settings",
			"Could not read server checkout settings for team:"+r.client.Team+"resource_group:"+plan.ResourceGroup+"project_id:"+plan.Project+": "+err.Error(),
		)
		return
	}

	// update the state with the updated server checkout settings
	plan.CheckoutRequired = updatedServerCheckoutSettings.CheckoutRequired
	plan.CheckoutDurationInSeconds = updatedServerCheckoutSettings.CheckoutDurationInSeconds
	plan.IncludeList = updatedServerCheckoutSettings.IncludeList
	plan.ExcludeList = updatedServerCheckoutSettings.ExcludeList

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

// Delete implements resource.Resource.
func (r *serverCheckoutSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve the server checkout settings values from the plan
	var state serverCheckoutSettingsResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	defaultCheckoutDurationInSeconds := int32(900)
	// Call the SDK client to reset the server checkout settings to default values
	serverCheckoutSettings := &pam.ResourceCheckoutSettings{
		CheckoutRequired:          false,
		CheckoutDurationInSeconds: &defaultCheckoutDurationInSeconds,
		IncludeList:               []string{},
		ExcludeList:               []string{},
	}

	err := client.UpdateServerCheckoutSettings(ctx, *r.client, state.ResourceGroup, state.Project, *serverCheckoutSettings)

	if err != nil {
		resp.Diagnostics.AddError("Error updating server checkout settings", err.Error())
		return
	}
	// Set state to empty data
	state.Id = types.StringValue("")
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

// Configure adds the provider configured client to the resource.
func (r *serverCheckoutSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}

	provider, ok := req.ProviderData.(*OktapamFrameworkProvider)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *OktapamFrameworkProvider, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}
	r.client = provider.SDKClientWrapper
}

func FormatServerCheckoutSettingsID(resourceGroupID string, projectID string) string {
	// server checkout settings don't have an identifier in itself and is really an attribute of a project.
	// we manage it as a separate resource since it's lifecycle is somewhat separate from a project.
	return fmt.Sprintf("%s/%s", resourceGroupID, projectID)
}
