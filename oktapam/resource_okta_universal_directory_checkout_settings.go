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
	_ resource.Resource              = &oktaUniversalDirectoryCheckoutSettingsResource{}
	_ resource.ResourceWithConfigure = &oktaUniversalDirectoryCheckoutSettingsResource{}
)

func NewoktaUniversalDirectoryCheckoutSettingsResource() resource.Resource {
	return &oktaUniversalDirectoryCheckoutSettingsResource{}
}

type oktaUniversalDirectoryCheckoutSettingsResource struct {
	api      *pam.ProjectsAPIService
	teamName string
}

type oktaUniversalDirectoryCheckoutSettingsResourceModel struct {
	Id            types.String `tfsdk:"id"`
	ResourceGroup string       `tfsdk:"resource_group"`
	Project       string       `tfsdk:"project"`
	convert.ServiceAccountCheckoutSettingsModel
}

func (r *oktaUniversalDirectoryCheckoutSettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_okta_universal_directory_checkout_settings"
}

func (r *oktaUniversalDirectoryCheckoutSettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages checkout settings for Okta Universal Directory resources in a project",
		Attributes: convert.ServiceAccountCheckoutSettingsSchemaAttributes(map[string]schema.Attribute{
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

func (r *oktaUniversalDirectoryCheckoutSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan oktaUniversalDirectoryCheckoutSettingsResourceModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	var checkoutSettings pam.APIServiceAccountCheckoutSettings
	if settings, diags := convert.ServiceAccountCheckoutSettingsFromModelToSDK(ctx, &plan.ServiceAccountCheckoutSettingsModel); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	} else {
		checkoutSettings = *settings
	}

	if _, err := r.api.UpdateResourceGroupOktaUniversalDirectoryBasedProjectCheckoutSettings(ctx, r.teamName, plan.ResourceGroup, plan.Project).APIServiceAccountCheckoutSettings(checkoutSettings).Execute(); err != nil {
		resp.Diagnostics.AddError("Error creating Okta UD checkout settings", err.Error())
		return
	}

	plan.Id = types.StringValue(formatOktaUDCheckoutSettingsID(plan.ResourceGroup, plan.Project))

	if diags := resp.State.Set(ctx, plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
}

func (r *oktaUniversalDirectoryCheckoutSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state oktaUniversalDirectoryCheckoutSettingsResourceModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	if checkoutSettings, _, err := r.api.FetchResourceGroupOktaUniversalDirectoryBasedProjectCheckoutSettings(ctx, r.teamName, state.ResourceGroup, state.Project).Execute(); err != nil {
		resp.Diagnostics.AddError(
			"Error reading Okta UD checkout settings",
			fmt.Sprintf("Could not read Okta UD checkout settings for team: %q resource_group: %q project_id: %q: Error: %s",
				r.teamName,
				state.ResourceGroup,
				state.Project,
				err.Error()))
		return
	} else {
		if settingsModel, diags := convert.ServiceAccountCheckoutSettingsFromSDKToModel(ctx, checkoutSettings); diags.HasError() {
			resp.Diagnostics.Append(diags...)
			return
		} else {
			state.ServiceAccountCheckoutSettingsModel = *settingsModel
		}
	}

	if diags := resp.State.Set(ctx, state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
}

func (r *oktaUniversalDirectoryCheckoutSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan oktaUniversalDirectoryCheckoutSettingsResourceModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	var checkoutSettings pam.APIServiceAccountCheckoutSettings
	if settings, diags := convert.ServiceAccountCheckoutSettingsFromModelToSDK(ctx, &plan.ServiceAccountCheckoutSettingsModel); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	} else {
		checkoutSettings = *settings
	}

	if _, err := r.api.UpdateResourceGroupOktaUniversalDirectoryBasedProjectCheckoutSettings(ctx, r.teamName, plan.ResourceGroup, plan.Project).APIServiceAccountCheckoutSettings(checkoutSettings).Execute(); err != nil {
		resp.Diagnostics.AddError("Error updating Okta UD checkout settings", err.Error())
		return
	}

	if updatedSettings, _, err := r.api.FetchResourceGroupOktaUniversalDirectoryBasedProjectCheckoutSettings(ctx, r.teamName, plan.ResourceGroup, plan.Project).Execute(); err != nil {
		resp.Diagnostics.AddError(
			"Error reading Okta UD checkout settings",
			fmt.Sprintf("Could not read Okta UD checkout settings for team: %q resource_group: %q project_id: %q: Error: %s",
				r.teamName,
				plan.ResourceGroup,
				plan.Project,
				err.Error()))
		return
	} else {
		if settingsModel, diags := convert.ServiceAccountCheckoutSettingsFromSDKToModel(ctx, updatedSettings); diags.HasError() {
			resp.Diagnostics.Append(diags...)
			return
		} else {
			plan.ServiceAccountCheckoutSettingsModel = *settingsModel
		}
	}

	if diags := resp.State.Set(ctx, plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
}

func (r *oktaUniversalDirectoryCheckoutSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state oktaUniversalDirectoryCheckoutSettingsResourceModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	defaultCheckoutDurationInSeconds := int32(900)
	checkoutSettings := &pam.APIServiceAccountCheckoutSettings{
		CheckoutRequired:          false,
		CheckoutDurationInSeconds: defaultCheckoutDurationInSeconds,
		IncludeList:               []pam.ServiceAccountSettingNameObject{},
		ExcludeList:               []pam.ServiceAccountSettingNameObject{},
	}

	if _, err := r.api.UpdateResourceGroupOktaUniversalDirectoryBasedProjectCheckoutSettings(ctx, r.teamName, state.ResourceGroup, state.Project).APIServiceAccountCheckoutSettings(*checkoutSettings).Execute(); err != nil {
		resp.Diagnostics.AddError("Error resetting Okta UD checkout settings", err.Error())
		return
	}

	state.Id = types.StringValue("")
	if diags := resp.State.Set(ctx, state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
}

func (r *oktaUniversalDirectoryCheckoutSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	sdkClient := getSDKClientFromMetadata(req.ProviderData)
	r.api = sdkClient.SDKClient.ProjectsAPI
	r.teamName = sdkClient.Team
}

func formatOktaUDCheckoutSettingsID(resourceGroupID string, projectID string) string {
	return fmt.Sprintf("%s/%s", resourceGroupID, projectID)
}
