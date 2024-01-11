package fwprovider

import (
	"context"
	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/fwprovider/resource_resource_group"
)

//var (
//	_ resource.ResourceWithConfigure   = &resourceGroupResource{}
//)

func NewResourceGroupResource() resource.Resource {
	return &resourceGroupResource{}
}

type resourceGroupResource struct {
	clientWrapper *client.SDKClientWrapper
	teamName      string
}

//func (r *resourceGroupResource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
//	providerData, _ := request.ProviderData.(*OktapamFrameworkProvider)
//
//	r.clientWrapper = providerData.SDKClientWrapper
//	r.teamName = providerData.SDKClientWrapper.Team
//}

func (r *resourceGroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
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

	// Convert
	resourceGroupAPIReq := pam.NewResourceGroupWithDefaults()
	resourceGroupAPIReq.SetDescription(data.Description.String())
	resourceGroupAPIReq.SetName(data.Name.String())

	adminGroups := data.DelegatedResourceAdminGroups.Elements()
	var adminGroupsAPI []pam.NamedObject
	for _, g := range adminGroups {
		adminGroupAPI, _ := g.(resource_resource_group.DelegatedResourceAdminGroupsValue).ToPamNamedObject(ctx)
		adminGroupsAPI =  append(adminGroupsAPI, *adminGroupAPI)
	}
	resourceGroupAPIReq.SetDelegatedResourceAdminGroups(adminGroupsAPI)

	// Create API call logic
	//r.clientWrapper.SDKClient.ResourceGroupsAPI.CreateResourceGroup(ctx, r.teamName).ResourceGroup(*resourceGroupAPIReq).Execute()

	// Example data value setting
	data.Id = types.StringValue("example-id")
	data.Name = types.StringValue("example-name")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data resource_resource_group.ResourceGroupModel

	data.Id = types.StringValue("example-id")
	data.Name = types.StringValue("example-name")
	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	//resGrp, _, err := r.clientWrapper.SDKClient.ResourceGroupsAPI.GetResourceGroup(ctx, r.teamName, data.Id.String()).Execute()
	//if err != nil {
	//	resp.Diagnostics.Append(diag.NewErrorDiagnostic("error while getting resource", err.Error()))
	//}


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

}
