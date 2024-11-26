package convert

import (
	"context"

	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

//type SecurityPolicyRuleResourceTypeModel types.String
//type SecurityPolicyRuleConditionTypeModel types.String
//type SecurityPolicyRuleResourceSelectorTypeModel types.String

type SecurityPolicyRuleModel struct {
	Name                              types.String `tfsdk:"name"`
	ResourceType                      types.String `tfsdk:"resource_type"`
	ResourceSelector                  types.Object/*SecurityPolicyRuleResourceSelectorModel*/ `tfsdk:"resource_selector"`
	Privileges                        types.List `tfsdk:"privileges"`
	Conditions                        types.List/**SecurityPolicyRuleConditionContainerModel*/ `tfsdk:"conditions"`
	OverrideCheckoutDurationInSeconds types.Int64 `tfsdk:"override_checkout_duration_in_seconds"`
	//NOTE: We ignore SecurityPolicyID
}

func SecurityPolicyRuleSchema() schema.NestedAttributeObject {
	return schema.NestedAttributeObject{
		Attributes: map[string]schema.Attribute{
			"name":                                  schema.StringAttribute{Required: true},
			"resource_type":                         schema.StringAttribute{Required: true},
			"override_checkout_duration_in_seconds": schema.Int64Attribute{Optional: true},
			"resource_selector":                     SecurityPolicyRuleResourceSelectorSchema(),
			"privileges": schema.ListNestedAttribute{
				NestedObject: SecurityPolicyRulePrivilegeContainerSchema(),
				Required:     true,
			},
			"conditions": schema.ListNestedAttribute{
				NestedObject: SecurityPolicyRuleConditionContainerSchema(),
				Optional:     true,
			},
		},
	}
}

func SecurityPolicyRuleAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":                                  types.StringType,
		"resource_type":                         types.StringType,
		"override_checkout_duration_in_seconds": types.Int64Type,
		"resource_selector":                     types.ObjectType{AttrTypes: SecurityPolicyRuleResourceSelectorAttrTypes()},
		"privileges":                            types.ListType{ElemType: types.ObjectType{AttrTypes: SecurityPolicyRulePrivilegeContainerAttrTypes()}},
		"conditions":                            types.ListType{ElemType: types.ObjectType{AttrTypes: SecurityPolicyRuleConditionContainerAttrTypes()}},
	}
}

func SecurityPolicyRuleFromModelToSDK(ctx context.Context, in *SecurityPolicyRuleModel) (*pam.SecurityPolicyRule, diag.Diagnostics) {
	var out pam.SecurityPolicyRule
	var diags diag.Diagnostics

	out.ResourceType = pam.SecurityPolicyRuleResourceType(in.ResourceType.ValueString())

	if !in.Name.IsNull() && !in.Name.IsUnknown() {
		out.Name = in.Name.ValueString()
	}

	if !in.ResourceSelector.IsNull() && !in.ResourceSelector.IsUnknown() {
		outSelector, d := SecurityPolicyRuleResourceSelectorFromModelToSDK(ctx, in.ResourceSelector)
		diags.Append(d...)
		if diags.HasError() {
			return nil, diags
		}
		out.ResourceSelector = *outSelector
	}

	if !in.Privileges.IsNull() && len(in.Privileges.Elements()) > 0 {
		privilegesModel := make([]SecurityPolicyRulePrivilegeContainerModel, 0, len(in.Privileges.Elements()))
		diags.Append(in.Privileges.ElementsAs(ctx, &privilegesModel, false)...)
		if diags.HasError() {
			return nil, diags
		}

		for _, privilege := range privilegesModel {
			outPrivilege, d := SecurityPolicyRulePrivilegeContainerFromModelToSDK(ctx, &privilege)
			diags.Append(d...)
			if diags.HasError() {
				return nil, diags
			}
			out.Privileges = append(out.Privileges, *outPrivilege)
		}
	}

	if !in.OverrideCheckoutDurationInSeconds.IsNull() && !in.OverrideCheckoutDurationInSeconds.IsUnknown() {
		out.OverrideCheckoutDurationInSeconds.Set(in.OverrideCheckoutDurationInSeconds.ValueInt64Pointer())
	}

	if !in.Conditions.IsNull() && len(in.Conditions.Elements()) > 0 {
		conditionsModel := make([]SecurityPolicyRuleConditionContainerModel, 0, len(in.Conditions.Elements()))
		diags.Append(in.Conditions.ElementsAs(ctx, &conditionsModel, false)...)
		if diags.HasError() {
			return nil, diags
		}
		for _, condition := range conditionsModel {
			outCondition, d := SecurityPolicyRuleConditionContainerFromModelToSDK(ctx, &condition)
			diags.Append(d...)
			if diags.HasError() {
				return nil, diags
			}
			out.Conditions = append(out.Conditions, *outCondition)

		}
	}
	return &out, diags
}

func SecurityPolicyRuleFromSDKToModel(ctx context.Context, in *pam.SecurityPolicyRule) (*SecurityPolicyRuleModel, diag.Diagnostics) {
	var out SecurityPolicyRuleModel
	var diags diag.Diagnostics

	out.Name = types.StringValue(in.Name)
	out.ResourceType = types.StringValue(string(in.ResourceType))

	outSelector, d := SecurityPolicyRuleResourceSelectorFromSDKToModel(ctx, &in.ResourceSelector)
	diags.Append(d...)
	if diags.HasError() {
		return nil, diags
	}
	out.ResourceSelector = *outSelector

	var outPrivileges []SecurityPolicyRulePrivilegeContainerModel
	for _, privilege := range in.Privileges {
		outPrivilege, d := SecurityPolicyRulePrivilegeContainerFromSDKToModel(ctx, &privilege)
		diags.Append(d...)
		if diags.HasError() {
			return nil, diags
		}
		outPrivileges = append(outPrivileges, *outPrivilege)

	}

	listValue, d := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: SecurityPolicyRulePrivilegeContainerAttrTypes()}, outPrivileges)
	diags.Append(d...)
	if diags.HasError() {
		return nil, diags
	}

	out.Privileges = listValue

	out.OverrideCheckoutDurationInSeconds = types.Int64PointerValue(in.OverrideCheckoutDurationInSeconds.Get())

	if len(in.Conditions) > 0 {
		var conditions []SecurityPolicyRuleConditionContainerModel
		for _, condition := range in.Conditions {
			outCondition, d := SecurityPolicyRuleConditionContainerFromSDKToModel(ctx, &condition)
			diags.Append(d...)
			if diags.HasError() {
				return nil, diags
			}
			conditions = append(conditions, *outCondition)
		}
		outConditions, d := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: SecurityPolicyRuleConditionContainerAttrTypes()}, conditions)
		diags.Append(d...)
		if diags.HasError() {
			return nil, diags
		}
		out.Conditions = outConditions

	} else {
		out.Conditions = types.ListNull(types.ObjectType{AttrTypes: SecurityPolicyRuleConditionContainerAttrTypes()})
	}
	return &out, diags
}
