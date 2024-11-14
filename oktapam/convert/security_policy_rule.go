package convert

import (
	"context"

	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

//type SecurityPolicyRuleResourceTypeModel types.String
//type SecurityPolicyRuleConditionTypeModel types.String
//type SecurityPolicyRuleResourceSelectorTypeModel types.String

type SecurityPolicyRuleModel struct {
	Name types.String `tfsdk:"name"`
	//ResourceType                      SecurityPolicyRuleResourceTypeModel                     `tfsdk:"resource_type"`
	ResourceSelector                  SecurityPolicyRuleResourceSelectorModel                                                    `tfsdk:"resource_selector"`
	Privileges                        types.List/*SecurityPolicyRulePrivilegeContainerPrivilegeValueModel*/ `tfsdk:"privileges"` //TODO(ja) should this be privilege_value ?!
	Conditions                        types.List/**SecurityPolicyRuleConditionModel*/ `tfsdk:"conditions"`
	OverrideCheckoutDurationInSeconds types.Int64  `tfsdk:"override_checkout_duration_in_seconds"`
	SecurityPolicyID                  types.String `tfsdk:"security_policy_id"` // openapi readOnly
}

func SecurityPolicyRuleSchema() schema.NestedAttributeObject {
	return schema.NestedAttributeObject{
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{Required: true},
			//"resource_type":                         schema.StringAttribute{Optional: true}, //TODO(ja) do I even need this?
			"override_checkout_duration_in_seconds": schema.Int64Attribute{Optional: true},
			"security_policy_id":                    schema.StringAttribute{Optional: true}, //TODO(ja) do I even need this?
			"resource_selector":                     SecurityPolicyRuleResourceSelectorSchema(),
			"privileges": schema.ListNestedAttribute{
				NestedObject: SecurityPolicyRulePrivilegeContainerSchema(),
				Required:     true,
			},
			"conditions": schema.ListNestedAttribute{
				NestedObject: SecurityPolicyRuleConditionsSchema(),
				Optional:     true,
			},
		},
	}
}

func SecurityPolicyRuleFromModelToSDK(ctx context.Context, in *SecurityPolicyRuleModel, out *pam.SecurityPolicyRule) diag.Diagnostics {
	out.Name = in.Name.ValueString()
	if diags := SecurityPolicyRuleResourceSelectorFromModelToSDK(ctx, &in.ResourceSelector, &out.ResourceSelector); diags.HasError() {
		return diags
	}

	privilegesModel := make([]SecurityPolicyRulePrivilegeContainerModel, 0, len(in.Privileges.Elements()))
	if diags := in.Privileges.ElementsAs(ctx, &privilegesModel, false); diags.HasError() {
		return diags
	}
	for _, privilege := range privilegesModel {
		var outPrivilege pam.SecurityPolicyRulePrivilegeContainer
		if diags := SecurityPolicyRulePrivilegeContainerFromModelToSDK(ctx, &privilege, &outPrivilege); diags.HasError() {
			return diags
		}
		out.Privileges = append(out.Privileges, outPrivilege)
	}

	// Conditions
	out.OverrideCheckoutDurationInSeconds.Set(in.OverrideCheckoutDurationInSeconds.ValueInt64Pointer())
	return nil
}

func SecurityPolicyRuleFromSDKToModel(ctx context.Context, in *pam.SecurityPolicyRule, out *SecurityPolicyRuleModel) diag.Diagnostics {
	out.Name = types.StringValue(in.Name)

	if diags := SecurityPolicyRuleResourceSelectorFromSDKToModel(ctx, &in.ResourceSelector, &out.ResourceSelector); diags.HasError() {
		return diags
	}

	out.OverrideCheckoutDurationInSeconds = types.Int64PointerValue(in.OverrideCheckoutDurationInSeconds.Get())

	return nil
}
