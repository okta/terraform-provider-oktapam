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
	ID          types.String                  `tfsdk:"id"`
	Name        types.String                  `tfsdk:"name"`
	Type        types.String                  `tfsdk:"type"`
	Description types.String                  `tfsdk:"description"`
	Active      types.Bool                    `tfsdk:"active"`
	Principals  SecurityPolicyPrincipalsModel `tfsdk:"principals"`
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

func SecurityPolicyFromModelToSDK(ctx context.Context, in *SecurityPolicyResourceModel, out *pam.SecurityPolicy) diag.Diagnostics {
	//TODO(ja) do this pattern everywhere
	if !in.ID.IsNull() && !in.ID.IsUnknown() {
		out.Id = in.ID.ValueStringPointer()
	}
	out.Name = in.Name.ValueString()
	if !in.Type.IsNull() && !in.Type.IsUnknown() {
		out.SetType(pam.SecurityPolicyType(in.Type.ValueString()))
	}
	out.Description = in.Description.ValueStringPointer()
	out.Active = in.Active.ValueBool()

	if diags := SecurityPolicyPrincipalFromModelToSDK(ctx, &in.Principals, &out.Principals); diags.HasError() {
		return diags
	}

	if !in.Rules.IsNull() && len(in.Rules.Elements()) > 0 {
		ruleModels := make([]SecurityPolicyRuleModel, 0, len(in.Rules.Elements()))
		if diags := in.Rules.ElementsAs(ctx, &ruleModels, false); diags.HasError() {
			return diags
		}
		for _, ruleModel := range ruleModels {
			var outPolicyRule pam.SecurityPolicyRule
			if diags := SecurityPolicyRuleFromModelToSDK(ctx, &ruleModel, &outPolicyRule); diags.HasError() {
				return diags
			}
			out.Rules = append(out.Rules, outPolicyRule)
		}
	}
	return nil
}

func SecurityPolicyFromSDKToModel(ctx context.Context, in *pam.SecurityPolicy, out *SecurityPolicyResourceModel) diag.Diagnostics {
	out.ID = types.StringPointerValue(in.Id)
	out.Name = types.StringValue(in.Name)
	if in.Type != nil {
		out.Type = types.StringValue(string(*in.Type))
	}

	if in.Description != nil {
		out.Description = types.StringPointerValue(in.Description)
	}

	out.Active = types.BoolValue(in.Active)

	if diags := SecurityPolicyPrincipalFromSDKToModel(ctx, &in.Principals, &out.Principals); diags.HasError() {
		return diags
	}

	if len(in.Rules) > 0 {
		var outRules []SecurityPolicyRuleModel
		for _, rule := range in.Rules {
			var outRule SecurityPolicyRuleModel
			if diags := SecurityPolicyRuleFromSDKToModel(ctx, &rule, &outRule); diags.HasError() {
				return diags
			}
			outRules = append(outRules, outRule)
		}
	}
	//TODO(ja) - Rules
	return nil
}
