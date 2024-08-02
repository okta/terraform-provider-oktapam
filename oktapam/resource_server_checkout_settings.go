package oktapam

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
				Description:         "The duration in seconds for the checkout. If the checkout is enabled, the duration is the maximum time a user can access the resource before the checkout expires.",
				MarkdownDescription: "The duration in seconds for the checkout. If the checkout is enabled, the duration is the maximum time a user can access the resource before the checkout expires.",
				Validators: []validator.Int64{
					int64validator.Between(900, 86400),
				},
			},
			"checkout_required": schema.BoolAttribute{
				Required:            true,
				Description:         "Indicates whether a checkout is mandatory for accessing resources within the project. If `true`, checkout is enforced for all applicable resources by default. If `false`, checkout is not required, and resources are accessible without it.",
				MarkdownDescription: "Indicates whether a checkout is mandatory for accessing resources within the project. If `true`, checkout is enforced for all applicable resources by default. If `false`, checkout is not required, and resources are accessible without it.",
			},
			"exclude_list": schema.ListAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Description:         "If provided, only the account identifiers listed are excluded from the checkout requirement. This list is only considered if `checkout_required` is set to `true`. Only one of `include_list` and `exclude_list` can be specified in a request since they are mutually exclusive.",
				MarkdownDescription: "If provided, only the account identifiers listed are excluded from the checkout requirement. This list is only considered if `checkout_required` is set to `true`. Only one of `include_list` and `exclude_list` can be specified in a request since they are mutually exclusive.",
			},
			"include_list": schema.ListAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Description:         "If provided, only the account identifiers listed are required to perform a checkout to access the resource. This list is only considered if `checkout_required` is set to `true`. Only one of `include_list` and `exclude_list` can be specified in a request since they are mutually exclusive.",
				MarkdownDescription: "If provided, only the account identifiers listed are required to perform a checkout to access the resource. This list is only considered if `checkout_required` is set to `true`. Only one of `include_list` and `exclude_list` can be specified in a request since they are mutually exclusive.",
			},
			"project": schema.StringAttribute{
				Required:            true,
				Description:         "The UUID of a Project",
				MarkdownDescription: "The UUID of a Project",
			},
			"resource_group": schema.StringAttribute{
				Required:            true,
				Description:         "The UUID of a Resource Group",
				MarkdownDescription: "The UUID of a Resource Group",
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

	_, err := r.client.SDKClient.ProjectsAPI.UpdateResourceGroupServerBasedProjectCheckoutSettings(ctx, r.client.Team, plan.ResourceGroup, plan.Project).ResourceCheckoutSettings(*serverCheckoutSettings).Execute()

	if err != nil {
		resp.Diagnostics.AddError("Error creating server checkout settings", err.Error())
		return
	}
	// Set state to fully populated data
	plan.Id = types.StringValue(formatServerCheckoutSettingsID(plan.ResourceGroup, plan.Project))
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
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
	serverCheckoutSettings, _, err := r.client.SDKClient.ProjectsAPI.FetchResourceGroupServerBasedProjectCheckoutSettings(ctx, r.client.Team, state.ResourceGroup, state.Project).Execute()
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

	_, err := r.client.SDKClient.ProjectsAPI.UpdateResourceGroupServerBasedProjectCheckoutSettings(ctx, r.client.Team, plan.ResourceGroup, plan.Project).ResourceCheckoutSettings(*serverCheckoutSettings).Execute()

	if err != nil {
		resp.Diagnostics.AddError("Error updating server checkout settings", err.Error())
		return
	}

	// Fetch the updated server checkout settings from the API host
	updatedServerCheckoutSettings, _, err := r.client.SDKClient.ProjectsAPI.FetchResourceGroupServerBasedProjectCheckoutSettings(ctx, r.client.Team, plan.ResourceGroup, plan.Project).Execute()
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
	if resp.Diagnostics.HasError() {
		return
	}
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

	_, err := r.client.SDKClient.ProjectsAPI.UpdateResourceGroupServerBasedProjectCheckoutSettings(ctx, r.client.Team, state.ResourceGroup, state.Project).ResourceCheckoutSettings(*serverCheckoutSettings).Execute()

	if err != nil {
		resp.Diagnostics.AddError("Error updating server checkout settings", err.Error())
		return
	}

	// Set state to empty data
	state.Id = types.StringValue("")
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// should we delete the state here?
	// resp.State.RemoveResource(ctx)
}

// Configure adds the provider configured client to the resource.
func (r *serverCheckoutSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}

	sdkClient := getSDKClientFromMetadata(req.ProviderData)
	//provider, ok := req.ProviderData.(*OktapamFrameworkProvider)
	//
	//if !ok {
	//	resp.Diagnostics.AddError(
	//		"Unexpected Data Source Configure Type",
	//		fmt.Sprintf("Expected *OktapamFrameworkProvider, got: %T. Please report this issue to the provider developers.", req.ProviderData),
	//	)
	//
	//	return
	//}
	r.client = &sdkClient
}

func formatServerCheckoutSettingsID(resourceGroupID string, projectID string) string {
	// server checkout settings don't have an identifier in itself and is really an attribute of a project.
	// we manage it as a separate resource since it's lifecycle is somewhat separate from a project.
	// since project password settings are managed as a separate resource with format resourceGroupID/projectID already,
	// we can just append the server_checkout_settings to the end of the ID
	return fmt.Sprintf("%s/%s/server_checkout_settings", resourceGroupID, projectID)
}