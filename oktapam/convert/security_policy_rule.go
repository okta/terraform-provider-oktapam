package convert

import (
	"context"

	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SecurityPolicyRuleResourceTypeModel types.String
type SecurityPolicyRuleConditionTypeModel types.String
type SecurityPolicyRuleResourceSelectorTypeModel types.String

type SecurityPolicyRuleModel struct {
	Name                              types.String                                            `tfsdk:"name"`
	ResourceType                      SecurityPolicyRuleResourceTypeModel                     `tfsdk:"resource_type"`
	ResourceSelector                  SecurityPolicyRuleResourceSelectorModel                 `tfsdk:"resources"` // openapi field: resource_selector
	Privileges                        SecurityPolicyRulePrivilegeContainerPrivilegeValueModel `tfsdk:"privileges"`
	Conditions                        SecurityPolicyRuleConditionContainerModel               `tfsdk:"conditions"`
	OverrideCheckoutDurationInSeconds types.Int64                                             `tfsdk:"override_checkout_duration_in_seconds"`
	SecurityPolicyID                  types.String                                            `tfsdk:"security_policy_id"` // openapi readOnly
}

func SecurityPolicyRulesSchema() schema.Attribute {
	return schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"name":                       schema.StringAttribute{Required: true},
				"resource_type":              schema.StringAttribute{Optional: true},
				"override_checkout_duration": schema.Int64Attribute{Optional: true},
				"security_policy_id":         schema.StringAttribute{Optional: true}, //TODO(ja) do I even need this?
				"resources":                  SecurityPolicyRuleResourceSelectorSchema(),
				"privileges":                 SecurityPolicyRulePrivilegesSchema(),
				"conditions":                 SecurityPolicyRuleConditionsSchema(),
			},
		},
		Required: true,
	}
}

func SecurityPolicyRuleFromModelToSDK(ctx context.Context, in *SecurityPolicyRuleModel, out *pam.SecurityPolicyRule) diag.Diagnostics {
	out.Name = in.Name.ValueString()
	// ResourceType
	// ResourceSelector
	// Privileges
	// Conditions
	out.OverrideCheckoutDurationInSeconds.Set(in.OverrideCheckoutDurationInSeconds.ValueInt64Pointer())
	return nil
}

func SecurityPolicyRuleFromSDKToModel(ctx context.Context, in *pam.SecurityPolicyRule, out *SecurityPolicyRuleModel) diag.Diagnostics {
	out.Name = types.StringValue(in.Name)
	out.OverrideCheckoutDurationInSeconds = types.Int64PointerValue(in.OverrideCheckoutDurationInSeconds.Get())

	return nil
}
