package fwprovider
package fwprovider

import (
	"context"
	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/fwprovider/resource_resource_group"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
)

var (
	_ resource.ResourceWithConfigure = &resourceGroupResource{}
)

func NewResourceGroupResource() resource.Resource {
	return &resourceGroupResource{}
}

type resourceGroupResource struct {
	clientWrapper *client.SDKClientWrapper
	teamName      string
}

func (r *resourceGroupResource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	providerData, _ := request.ProviderData.(*OktapamFrameworkProvider)

	// Terraform calling the ValidateResourceConfig RPC would not call the ConfigureProvider first so ProviderData can be nil
	// Prevent panic if the provider has not been configured.
	if providerData == nil {
		return
	}
	r.clientWrapper = providerData.SDKClientWrapper
	r.teamName = providerData.SDKClientWrapper.Team
}

func (r *resourceGroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_resource_group_fwk"
	resp.TypeName = req.ProviderTypeName + "_resource_group_fwk"
}

func (r *resourceGroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_resource_group.ResourceGroupResourceSchema(ctx)
}

func (r *resourceGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resource_resource_group.ResourceGroupModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Build API Request for calling SDK
	resourceGroupAPIReq := pam.NewResourceGroupWithDefaults()
	resourceGroupAPIReq.SetDescription(data.Description.ValueString())
	resourceGroupAPIReq.SetName(data.Name.ValueString())

	adminGroupsTF := data.DelegatedResourceAdminGroups.Elements()
	var adminGroupsAPI []pam.NamedObject
	for _, g := range adminGroupsTF {
		adminGroupAPI, _ := g.(resource_resource_group.DelegatedResourceAdminGroupsValue).ToPamNamedObject(ctx)
		adminGroupsAPI = append(adminGroupsAPI, *adminGroupAPI)
	}
	resourceGroupAPIReq.SetDelegatedResourceAdminGroups(adminGroupsAPI)


	// Create API call logic
	resGrp, _, err := r.clientWrapper.SDKClient.ResourceGroupsAPI.CreateResourceGroup(ctx, r.teamName).
		ResourceGroup(*resourceGroupAPIReq).Execute()

	if err != nil {
		logging.Errorf("error while creating resource group: %s", err.Error())
		resp.Diagnostics.Append(diag.NewErrorDiagnostic("error while creating resource group", err.Error()))
		return
	}
	// Example data value setting

	if resGrp != nil {
		r.updateState(ctx, &data, resGrp)
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceGroupResource) updateState(_ context.Context, state *resource_resource_group.ResourceGroupModel, resp *pam.ResourceGroup) {
	state.Id = types.StringValue(*resp.Id)
	state.Name = types.StringValue(*resp.Name)
	state.Description = types.StringValue(*resp.Description)
}

func (r *resourceGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data resource_resource_group.ResourceGroupModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	//Read API call logic
	resGrp, _, err := r.clientWrapper.SDKClient.ResourceGroupsAPI.GetResourceGroup(ctx, r.teamName, data.Id.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.Append(diag.NewErrorDiagnostic("error while getting resource", err.Error()))
	}

	if resGrp != nil {
		data.Id = types.StringValue(*resGrp.Id)
		data.Name = types.StringValue(*resGrp.Name)
		data.Description = types.StringValue(*resGrp.Description)
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data resource_resource_group.ResourceGroupModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update API call logic

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data resource_resource_group.ResourceGroupModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete API call logic
	 _, err := r.clientWrapper.SDKClient.ResourceGroupsAPI.DeleteResourceGroup(ctx, r.teamName, data.Id.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.Append(diag.NewErrorDiagnostic("error while deleting resource", err.Error()))
	}
}
