package convert

import (
	"context"

	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PasswordCheckoutSSHPrivilegeModel struct {
	PasswordCheckoutSSH types.Bool `tfsdk:"password_checkout_ssh"`
}

func PasswordCheckoutSSHPrivilegeSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"password_checkout_ssh": schema.BoolAttribute{
				Required: true,
			},
		},
		Optional: true,
	}
}

func PasswordCheckoutSSHPrivilegeAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"password_checkout_ssh": types.BoolType,
	}
}

func PasswordCheckoutSSHPrivilegeFromModelToSDK(_ context.Context, in *PasswordCheckoutSSHPrivilegeModel) (*pam.SecurityPolicyPasswordCheckoutSSHPrivilege, diag.Diagnostics) {
	var out pam.SecurityPolicyPasswordCheckoutSSHPrivilege
	var diags diag.Diagnostics

	out.Type = pam.SecurityPolicyRulePrivilegeType_PASSWORD_CHECKOUT_SSH

	if !in.PasswordCheckoutSSH.IsNull() && !in.PasswordCheckoutSSH.IsUnknown() {
		out.PasswordCheckoutSsh = in.PasswordCheckoutSSH.ValueBool()
	}
	return &out, diags
}

func PasswordCheckoutSSHPrivilegeFromSDKToModel(_ context.Context, in *pam.SecurityPolicyPasswordCheckoutSSHPrivilege) (*PasswordCheckoutSSHPrivilegeModel, diag.Diagnostics) {
	var out PasswordCheckoutSSHPrivilegeModel
	var diags diag.Diagnostics

	out.PasswordCheckoutSSH = types.BoolValue(in.PasswordCheckoutSsh)
	return &out, diags
}
