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

	out.ResourceType = pam.SecurityPolicyRuleResourceType(in.ResourceType.ValueString())

	if !in.Name.IsNull() && !in.Name.IsUnknown() {
		out.Name = in.Name.ValueString()
	}

	if !in.ResourceSelector.IsNull() && !in.ResourceSelector.IsUnknown() {
		if outSelector, diags := SecurityPolicyRuleResourceSelectorFromModelToSDK(ctx, in.ResourceSelector); diags.HasError() {
			return nil, diags
		} else {
			out.ResourceSelector = *outSelector
		}
	}

	if !in.Privileges.IsNull() && len(in.Privileges.Elements()) > 0 {
		privilegesModel := make([]SecurityPolicyRulePrivilegeContainerModel, 0, len(in.Privileges.Elements()))
		if diags := in.Privileges.ElementsAs(ctx, &privilegesModel, false); diags.HasError() {
			return nil, diags
		}
		for _, privilege := range privilegesModel {
			if outPrivilege, diags := SecurityPolicyRulePrivilegeContainerFromModelToSDK(ctx, &privilege); diags.HasError() {
				return nil, diags
			} else {
				out.Privileges = append(out.Privileges, *outPrivilege)
			}
		}
	}

	if !in.OverrideCheckoutDurationInSeconds.IsNull() && !in.OverrideCheckoutDurationInSeconds.IsUnknown() {
		out.OverrideCheckoutDurationInSeconds.Set(in.OverrideCheckoutDurationInSeconds.ValueInt64Pointer())
	}

	if !in.Conditions.IsNull() && len(in.Conditions.Elements()) > 0 {
		conditionsModel := make([]SecurityPolicyRuleConditionContainerModel, 0, len(in.Conditions.Elements()))
		if diags := in.Conditions.ElementsAs(ctx, &conditionsModel, false); diags.HasError() {
			return nil, diags
		}
		for _, condition := range conditionsModel {
			if outCondition, diags := SecurityPolicyRuleConditionContainerFromModelToSDK(ctx, &condition); diags.HasError() {
				return nil, diags
			} else {
				out.Conditions = append(out.Conditions, *outCondition)
			}
		}
	}

	return &out, nil
}

func SecurityPolicyRuleFromSDKToModel(ctx context.Context, in *pam.SecurityPolicyRule) (*SecurityPolicyRuleModel, diag.Diagnostics) {
	var out SecurityPolicyRuleModel

	out.Name = types.StringValue(in.Name)
	out.ResourceType = types.StringValue(string(in.ResourceType))

	if outSelector, diags := SecurityPolicyRuleResourceSelectorFromSDKToModel(ctx, &in.ResourceSelector); diags.HasError() {
		return nil, diags
	} else {
		out.ResourceSelector = *outSelector
	}

	var outPrivileges []SecurityPolicyRulePrivilegeContainerModel
	for _, privilege := range in.Privileges {
		if outPrivilege, diags := SecurityPolicyRulePrivilegeContainerFromSDKToModel(ctx, &privilege); diags.HasError() {
			return nil, diags
		} else {
			outPrivileges = append(outPrivileges, *outPrivilege)
		}
	}

	if listValue, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: SecurityPolicyRulePrivilegeContainerAttrTypes()}, outPrivileges); diags.HasError() {
		return nil, diags
	} else {
		out.Privileges = listValue
	}

	out.OverrideCheckoutDurationInSeconds = types.Int64PointerValue(in.OverrideCheckoutDurationInSeconds.Get())

	if len(in.Conditions) > 0 {
		var conditions []SecurityPolicyRuleConditionContainerModel
		for _, condition := range in.Conditions {
			if outCondition, diags := SecurityPolicyRuleConditionContainerFromSDKToModel(ctx, &condition); diags.HasError() {
				return nil, diags
			} else {
				conditions = append(conditions, *outCondition)
			}
		}
		if outConditions, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: SecurityPolicyRuleConditionContainerAttrTypes()}, conditions); diags.HasError() {
			return nil, diags
		} else {
			out.Conditions = outConditions
		}
	} else {
		out.Conditions = types.ListNull(types.ObjectType{AttrTypes: SecurityPolicyRuleConditionContainerAttrTypes()})
	}
	return &out, nil
}
