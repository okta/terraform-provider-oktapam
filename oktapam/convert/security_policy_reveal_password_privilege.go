package convert

import (
	"context"

	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RevealPasswordPrivilegeModel struct {
	RevealPassword types.Bool `tfsdk:"reveal_password"`
}

func RevealPasswordPrivilegeSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"reveal_password": schema.BoolAttribute{
				Required: true,
			},
		},
		Optional: true,
	}
}

func RevealPasswordPrivilegeAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"reveal_password": types.BoolType,
	}
}

func RevealPasswordPrivilegeFromModelToSDK(_ context.Context, in *RevealPasswordPrivilegeModel) (*pam.SecurityPolicyRevealPasswordPrivilege, diag.Diagnostics) {
	var out pam.SecurityPolicyRevealPasswordPrivilege
	var diags diag.Diagnostics

	out.Type = pam.SecurityPolicyRulePrivilegeType_REVEAL_PASSWORD

	if !in.RevealPassword.IsNull() && !in.RevealPassword.IsUnknown() {
		out.RevealPassword = in.RevealPassword.ValueBool()
	}

	return &out, diags
}

func RevealPasswordPrivilegeFromSDKToModel(_ context.Context, in *pam.SecurityPolicyRevealPasswordPrivilege) (*RevealPasswordPrivilegeModel, diag.Diagnostics) {
	var out RevealPasswordPrivilegeModel
	var diags diag.Diagnostics

	if val, ok := in.GetRevealPasswordOk(); ok {
		out.RevealPassword = types.BoolPointerValue(val)
	}
	return &out, diags
}
