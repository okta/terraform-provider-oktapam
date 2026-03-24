package oktapam

import (
	"context"

	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
	"github.com/okta/terraform-provider-oktapam/oktapam/convert"
)

var _ resource.Resource = &SecurityPolicyResource{}
var _ resource.ResourceWithImportState = &SecurityPolicyResource{}

//var _ resource.ResourceWithUpgradeState = &SecurityPolicyResource{}

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
		//Version:     1,
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
		if plan.ResourceGroup.IsUnknown() {
			if rg, ok := policyResponse.GetResourceGroupOk(); ok {
				if id, idOk := rg.GetIdOk(); idOk {
					plan.ResourceGroup = types.StringPointerValue(id)
				} else {
					plan.ResourceGroup = types.StringNull()
				}
			} else {
				plan.ResourceGroup = types.StringNull()
			}
		}
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

func (s *SecurityPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Validate that the import ID is a valid UUID
	if !MatchesUUID(req.ID) {
		resp.Diagnostics.AddError(
			"Invalid Import ID",
			"Security policy import requires a valid UUID. "+
				"Please check the ID and try again.",
		)
		return
	}

	// Use the helper function to set the import identifier to the "id" attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// UpgradeState handles migrations when the schema changes in a way that is incompatible with previously stored state.
// The Terraform Plugin Framework stores all schema-defined attributes in state, including optional ones that
// were never configured (as null). When an attribute is removed from the schema, the stored null entry becomes orphaned
// and causes "unsupported attribute" errors. Each upgrader migrates state from one schema version to the next.
//
// When is a state upgrade needed vs not:
// If a state upgrader is needed in the future, the implementation below can be uncommented.
// It re-parses the stored raw state JSON against the current schema with IgnoreUndefinedAttributes
// enabled, which silently drops any attribute present in the old state but absent from the current schema.
//
//   - REMOVING an attribute from the schema  → state upgrade required (old state still has the attribute as null, TF rejects it)
//   - ADDING a new Optional/Computed attribute → state upgrade NOT needed (TF treats missing attributes in old state as null automatically)
//   - RENAMING an attribute  → state upgrade required
//
// How to add a new state upgrader:
//
//  1. Bump the Version in Schema() (e.g., 1 → 2).
//  2. Add a new entry to this map keyed by the OLD version (e.g., 1 → upgradeV1ToV2).
//  3. The upgrader function can reuse upgradeSecurityPolicyState — it is version-agnostic and strips any attributes absent from the current schema.
//func (s *SecurityPolicyResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
//	return map[int64]resource.StateUpgrader{
//		0: {
//			StateUpgrader: upgradeSecurityPolicyState,
//		},
//	}
//}
//
//// upgradeSecurityPolicyState parses the stored raw state JSON against the current schema with IgnoreUndefinedAttributes enabled,
//// which silently drops any attribute present in the old state but absent from the current schema. Terraform already handles new attributes
//// (missing from old state) by treating them as null.
//func upgradeSecurityPolicyState(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
//	if req.RawState == nil {
//		return
//	}
//
//	rawJSON := req.RawState.JSON
//	if len(rawJSON) == 0 {
//		return
//	}
//
//	// Build the current schema type for unmarshalling
//	currentSchema := schema.Schema{
//		Attributes: convert.SecurityPolicySchema(),
//	}
//	resourceSchemaType := currentSchema.Type().TerraformType(ctx)
//
//	// Parse the old state JSON, ignoring attributes that no longer exist in the current schema. This covers any attribute removal at any nesting level
//	// (top-level fields, nested privilege types, condition attributes, etc.)
//	upgradedStateValue, err := tftypes.ValueFromJSONWithOpts(
//		rawJSON,
//		resourceSchemaType,
//		tftypes.ValueFromJSONOpts{IgnoreUndefinedAttributes: true},
//	)
//	if err != nil {
//		resp.Diagnostics.AddError(
//			"Unable to Upgrade Resource State",
//			fmt.Sprintf("An unexpected error occurred reading the prior resource state for upgrade. This needs to be reported to the provider developer: %s", err.Error()),
//		)
//		return
//	}
//
//	resp.State.Raw = upgradedStateValue
//}

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
