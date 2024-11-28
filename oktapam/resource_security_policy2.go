package oktapam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
	"github.com/okta/terraform-provider-oktapam/oktapam/convert"
)

var _ resource.Resource = &SecurityPolicyResource{}

type SecurityPolicyResource struct {
	teamName string
	api      *pam.SecurityPolicyAPIService
}

func NewSecurityPolicyResource() resource.Resource {
	return &SecurityPolicyResource{}
}

func (s *SecurityPolicyResource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_security_policy_v2"
}

func (s *SecurityPolicyResource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description: descriptions.ResourceSecurityPolicy,
		Attributes:  convert.SecurityPolicySchema(),
	}
}

func (s *SecurityPolicyResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var plan convert.SecurityPolicyResourceModel
	if diags := request.Plan.Get(ctx, &plan); diags.HasError() {
		response.Diagnostics.Append(diags...)
		return
	}

	securityPolicy, diags := convert.SecurityPolicyFromModelToSDK(ctx, &plan)
	if diags.HasError() {
		response.Diagnostics.Append(diags...)
		return
	}

	if policyResponse, _, err := s.api.CreateSecurityPolicy(ctx, s.teamName).SecurityPolicy(*securityPolicy).Execute(); err != nil {
		response.Diagnostics.AddError("Error creating security policy", err.Error())
		return
	} else {
		plan.ID = types.StringValue(*policyResponse.Id) //TODO(ja) can this be nil?
		if diags := response.State.Set(ctx, plan); diags.HasError() {
			response.Diagnostics.Append(diags...)
			return
		}
	}
}

func (s *SecurityPolicyResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var state convert.SecurityPolicyResourceModel
	var diags diag.Diagnostics

	diags.Append(request.State.Get(ctx, &state)...)
	if diags.HasError() {
		response.Diagnostics.Append(diags...)
		return
	}

	if responsePolicy, _, err := s.api.GetSecurityPolicy(ctx, s.teamName, state.ID.ValueString()).Execute(); err != nil {
		diags.AddError("Error reading security policy", err.Error())
		response.Diagnostics.Append(diags...)
		return
	} else {
		policyModel, d := convert.SecurityPolicyFromSDKToModel(ctx, responsePolicy)
		diags.Append(d...)
		if diags.HasError() {
			response.Diagnostics.Append(diags...)
			return
		}
		state = *policyModel
	}

	diags.Append(response.State.Set(ctx, state)...)
	if diags.HasError() {
		response.Diagnostics.Append(diags...)
		return
	}

	response.Diagnostics.Append(diags...)
}

func (s *SecurityPolicyResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var plan convert.SecurityPolicyResourceModel
	var diags diag.Diagnostics

	diags.Append(request.Plan.Get(ctx, &plan)...)
	if diags.HasError() {
		response.Diagnostics.Append(diags...)
		return
	}

	securityPolicy, d := convert.SecurityPolicyFromModelToSDK(ctx, &plan)
	diags.Append(d...)

	if diags.HasError() {
		response.Diagnostics.Append(diags...)
		return
	}

	if _, err := s.api.UpdateSecurityPolicy(ctx, s.teamName, plan.ID.ValueString()).SecurityPolicy(*securityPolicy).Execute(); err != nil {
		diags.AddError("Error updating security policy", err.Error())
		response.Diagnostics.Append(diags...)
		return
	}

	if updatedPolicy, _, err := s.api.GetSecurityPolicy(ctx, s.teamName, plan.ID.ValueString()).Execute(); err != nil {
		diags.AddError("Error reading security policy", err.Error())
		response.Diagnostics.Append(diags...)
		return
	} else {
		policyModel, d := convert.SecurityPolicyFromSDKToModel(ctx, updatedPolicy)
		diags.Append(d...)
		if diags.HasError() {
			response.Diagnostics.Append(diags...)
			return
		}
		plan = *policyModel

		diags.Append(response.State.Set(ctx, plan)...)
		if diags.HasError() {
			response.Diagnostics.Append(diags...)
			return
		}
	}

	response.Diagnostics.Append(diags...)
}

func (s *SecurityPolicyResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var state convert.SecurityPolicyResourceModel
	if diags := request.State.Get(ctx, &state); diags.HasError() {
		response.Diagnostics.Append(diags...)
		return
	}

	if _, err := s.api.DeleteSecurityPolicy(ctx, s.teamName, state.ID.ValueString()).Execute(); err != nil {
		response.Diagnostics.AddError("Error deleting security policy", err.Error())
	}
}

func (s *SecurityPolicyResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}

	sdkClient := getSDKClientFromMetadata(req.ProviderData)
	s.teamName = sdkClient.Team
	s.api = sdkClient.SDKClient.SecurityPolicyAPI
}
