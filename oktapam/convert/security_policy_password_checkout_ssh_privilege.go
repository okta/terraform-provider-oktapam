package convert

import (
	"context"

	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Begin SecurityPolicyPasswordCheckoutSSH

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

func PasswordCheckoutSSHPrivilegeFromModelToSDK(_ context.Context, in *PasswordCheckoutSSHPrivilegeModel, out *pam.SecurityPolicyPasswordCheckoutSSHPrivilege) diag.Diagnostics {
	out.Type = pam.SecurityPolicyRulePrivilegeType_PASSWORD_CHECKOUT_SSH
	out.PasswordCheckoutSsh = in.PasswordCheckoutSSH.ValueBool()
	return nil
}

func PasswordCheckoutSSHPrivilegeFromSDKToModel(_ context.Context, in *pam.SecurityPolicyPasswordCheckoutSSHPrivilege, out *PasswordCheckoutSSHPrivilegeModel) diag.Diagnostics {
	out.PasswordCheckoutSSH = types.BoolValue(in.PasswordCheckoutSsh)
	return nil
}

// End SecurityPolicyPasswordCheckoutSSH
