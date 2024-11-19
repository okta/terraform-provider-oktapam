package oktapam

import (
	"context"
	"fmt"

	"github.com/okta/terraform-provider-oktapam/oktapam/convert"

	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
)

var (
	_ resource.Resource              = &serverCheckoutSettingsResource{}
	_ resource.ResourceWithConfigure = &serverCheckoutSettingsResource{}
)

func NewServerCheckoutSettingsResource() resource.Resource {
	return &serverCheckoutSettingsResource{}
}

type serverCheckoutSettingsResource struct {
	api      *pam.ProjectsAPIService
	teamName string
}

type serverCheckoutSettingsResourceModel struct {
	Id            types.String `tfsdk:"id"`
	ResourceGroup string       `tfsdk:"resource_group"`
	Project       string       `tfsdk:"project"`
	convert.ResourceCheckoutSettingsModel
}

// Metadata implements resource.Resource.
func (r *serverCheckoutSettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_server_checkout_settings"
}

// Schema implements resource.Resource.
func (r *serverCheckoutSettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: descriptions.ResourceServerCheckoutSettings,
		Attributes: convert.ResourceCheckoutSettingsSchemaAttributes(map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"project": schema.StringAttribute{
				Required:    true,
				Description: descriptions.ProjectID,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"resource_group": schema.StringAttribute{
				Required:    true,
				Description: descriptions.ResourceGroupID,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
		}),
	}
}

func (r *serverCheckoutSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve the server checkout settings values from the plan
	var plan serverCheckoutSettingsResourceModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	var resourceCheckoutSettings pam.ResourceCheckoutSettings

	if inSettings, diags := convert.ResourceCheckoutSettingsFromModelToSDK(ctx, &plan.ResourceCheckoutSettingsModel); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	} else {
		resourceCheckoutSettings = *inSettings
	}

	if _, err := r.api.UpdateResourceGroupServerBasedProjectCheckoutSettings(ctx, r.teamName, plan.ResourceGroup, plan.Project).ResourceCheckoutSettings(resourceCheckoutSettings).Execute(); err != nil {
		resp.Diagnostics.AddError("Error creating server checkout settings", err.Error())
		return
	}
	// Set state to fully populated data
	plan.Id = types.StringValue(formatServerCheckoutSettingsID(plan.ResourceGroup, plan.Project))

	if diags := resp.State.Set(ctx, plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
}

// Read implements resource.Resource.
func (r *serverCheckoutSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state serverCheckoutSettingsResourceModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	// Get refreshed server checkout settings from API host
	if checkoutSettings, _, err := r.api.FetchResourceGroupServerBasedProjectCheckoutSettings(ctx, r.teamName, state.ResourceGroup, state.Project).Execute(); err != nil {
		resp.Diagnostics.AddError(
			"Error reading server checkout settings",
			fmt.Sprintf("Could not read server checkout settings for team: %q resource_group: %q project_id: %q: Error: %s",
				r.teamName,
				state.ResourceGroup,
				state.Project,
				err.Error()))
		return
	} else {
		// Overwrite server checkout settings with refreshed state
		if settingsModel, diags := convert.ResourceCheckoutSettingsFromSDKToModel(ctx, checkoutSettings); diags.HasError() {
			resp.Diagnostics.Append(diags...)
			return
		} else {
			state.ResourceCheckoutSettingsModel = *settingsModel
		}
	}

	// Set state to fully populated data
	if diags := resp.State.Set(ctx, state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
}

// Update implements resource.Resource.
func (r *serverCheckoutSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve the server checkout settings values from the plan
	var plan serverCheckoutSettingsResourceModel

	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	// Update the existing checkout settings with the new values
	var resourceCheckoutSettings pam.ResourceCheckoutSettings
	if inSettings, diags := convert.ResourceCheckoutSettingsFromModelToSDK(ctx, &plan.ResourceCheckoutSettingsModel); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	} else {
		resourceCheckoutSettings = *inSettings
	}

	if _, err := r.api.UpdateResourceGroupServerBasedProjectCheckoutSettings(ctx, r.teamName, plan.ResourceGroup, plan.Project).ResourceCheckoutSettings(resourceCheckoutSettings).Execute(); err != nil {
		resp.Diagnostics.AddError("Error updating server checkout settings", err.Error())
		return
	}

	// Fetch the updated server checkout settings from the API host
	if updatedServerCheckoutSettings, _, err := r.api.FetchResourceGroupServerBasedProjectCheckoutSettings(ctx, r.teamName, plan.ResourceGroup, plan.Project).Execute(); err != nil {
		resp.Diagnostics.AddError(
			"Error reading server checkout settings",
			fmt.Sprintf("Could not read server checkout settings for team: %q resource_group: %q project_id: %q: Error: %s",
				r.teamName,
				plan.ResourceGroup,
				plan.Project,
				err.Error()))
		return
	} else {
		// update the state with the updated checkout settings
		if settingsModel, diags := convert.ResourceCheckoutSettingsFromSDKToModel(ctx, updatedServerCheckoutSettings); diags.HasError() {
			resp.Diagnostics.Append(diags...)
			return
		} else {
			plan.ResourceCheckoutSettingsModel = *settingsModel
		}
	}

	// Set state to fully populated data
	if diags := resp.State.Set(ctx, plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
}

// Delete implements resource.Resource.
func (r *serverCheckoutSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve the server checkout settings values from the plan
	var state serverCheckoutSettingsResourceModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	defaultCheckoutDurationInSeconds := int32(900)
	// Call the SDK client to reset the server checkout settings to default values
	checkoutSettings := &pam.ResourceCheckoutSettings{
		CheckoutRequired:          false,
		CheckoutDurationInSeconds: &defaultCheckoutDurationInSeconds,
		IncludeList:               []string{},
		ExcludeList:               []string{},
	}

	if _, err := r.api.UpdateResourceGroupServerBasedProjectCheckoutSettings(ctx, r.teamName, state.ResourceGroup, state.Project).ResourceCheckoutSettings(*checkoutSettings).Execute(); err != nil {
		resp.Diagnostics.AddError("Error updating server checkout settings", err.Error())
		return
	}

	// Set state to empty data
	state.Id = types.StringValue("")
	if diags := resp.State.Set(ctx, state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
}

// Configure adds the provider configured client to the resource.
func (r *serverCheckoutSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}

	sdkClient := getSDKClientFromMetadata(req.ProviderData)
	r.api = sdkClient.SDKClient.ProjectsAPI
	r.teamName = sdkClient.Team
}

func formatServerCheckoutSettingsID(resourceGroupID string, projectID string) string {
	// server checkout settings don't have an identifier in itself and is really an attribute of a project.
	// we manage it as a separate resource since it's lifecycle is somewhat separate from a project.
	return fmt.Sprintf("%s/%s", resourceGroupID, projectID)
}
