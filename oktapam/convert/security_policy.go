package convert

import (
	"context"

	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

//type SecurityPolicyTypeModel types.String

type SecurityPolicyResourceModel struct {
	ID          types.String                   `tfsdk:"id"`
	Name        types.String                   `tfsdk:"name"`
	Type        types.String                   `tfsdk:"type"`
	Description types.String                   `tfsdk:"description"`
	Active      types.Bool                     `tfsdk:"active"`
	Principals  *SecurityPolicyPrincipalsModel `tfsdk:"principals"`
	Rules       types.List/*SecurityPolicyRuleModel*/ `tfsdk:"rules"`
}

func SecurityPolicySchema() map[string]schema.Attribute {
	myAttributes := map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:      true,
			PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"name": schema.StringAttribute{
			Required: true,
		},
		"type": schema.StringAttribute{
			Optional: true,
		},
		"description": schema.StringAttribute{
			Optional: true,
		},
		"active": schema.BoolAttribute{
			Required: true,
		},
		"principals": SecurityPolicyPrincipalsSchema(),
		"rules": schema.ListNestedAttribute{
			NestedObject: SecurityPolicyRuleSchema(),
			Required:     true,
		},
	}

	return myAttributes
}

func SecurityPolicyFromModelToSDK(ctx context.Context, in *SecurityPolicyResourceModel) (*pam.SecurityPolicy, diag.Diagnostics) {
	var out pam.SecurityPolicy
	var diags diag.Diagnostics

	if !in.ID.IsNull() && !in.ID.IsUnknown() {
		out.Id = in.ID.ValueStringPointer()
	}

	if !in.Name.IsNull() && !in.Name.IsUnknown() {
		out.Name = in.Name.ValueString()
	}

	if !in.Type.IsNull() && !in.Type.IsUnknown() {
		out.SetType(pam.SecurityPolicyType(in.Type.ValueString()))
	}

	if !in.Description.IsNull() && !in.Description.IsUnknown() {
		out.Description = in.Description.ValueStringPointer()
	}

	if !in.Active.IsNull() && !in.Active.IsUnknown() {
		out.Active = in.Active.ValueBool()
	}

	var principals *pam.SecurityPolicyPrincipals
	var d diag.Diagnostics
	if in.Principals != nil {
		principals, d = SecurityPolicyPrincipalFromModelToSDK(ctx, in.Principals)
		diags.Append(d...)
	} else {
		// Create empty principals if nil
		principals = &pam.SecurityPolicyPrincipals{}
	}

	if diags.HasError() {
		return nil, diags
	}

	out.Principals = *principals

	if !in.Rules.IsNull() && len(in.Rules.Elements()) > 0 {
		ruleModels := make([]SecurityPolicyRuleModel, 0, len(in.Rules.Elements()))
		diags.Append(in.Rules.ElementsAs(ctx, &ruleModels, false)...)
		if diags.HasError() {
			return nil, diags
		}

		for _, ruleModel := range ruleModels {
			outPolicyRule, d := SecurityPolicyRuleFromModelToSDK(ctx, &ruleModel)
			diags.Append(d...)
			if diags.HasError() {
				return nil, diags
			}
			out.Rules = append(out.Rules, *outPolicyRule)
		}
	}
	return &out, diags
}

func SecurityPolicyFromSDKToModel(ctx context.Context, in *pam.SecurityPolicy) (*SecurityPolicyResourceModel, diag.Diagnostics) {
	var out SecurityPolicyResourceModel
	var diags diag.Diagnostics

	out.ID = types.StringPointerValue(in.Id)
	if val, ok := in.GetNameOk(); ok {
		out.Name = types.StringPointerValue(val)
	}

	if val, ok := in.GetTypeOk(); ok {
		valStr := string(*val)
		out.Type = types.StringValue(valStr)
	}

	if val, ok := in.GetDescriptionOk(); ok {
		out.Description = types.StringPointerValue(val)
	}

	if val, ok := in.GetActiveOk(); ok {
		out.Active = types.BoolPointerValue(val)
	}

	principals, d := SecurityPolicyPrincipalFromSDKToModel(ctx, &in.Principals)
	diags.Append(d...)

	if diags.HasError() {
		return nil, diags
	}
	out.Principals = principals

	if len(in.Rules) > 0 {
		var outRules []SecurityPolicyRuleModel
		for _, rule := range in.Rules {
			outRule, d := SecurityPolicyRuleFromSDKToModel(ctx, &rule)
			diags.Append(d...)
			if diags.HasError() {
				return nil, diags
			}
			outRules = append(outRules, *outRule)

		}

		listValue, d := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: SecurityPolicyRuleAttrTypes()}, outRules)
		diags.Append(d...)
		if diags.HasError() {
			return nil, diags
		}

		out.Rules = listValue

	} else {
		out.Rules = types.ListNull(types.ObjectType{AttrTypes: SecurityPolicyRuleAttrTypes()})
	}

	return &out, diags
}
