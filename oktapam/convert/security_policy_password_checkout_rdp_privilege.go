package convert

import (
	"context"

	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PasswordCheckoutRDPPrivilegeModel struct {
	PasswordCheckoutRDP types.Bool `tfsdk:"password_checkout_rdp"`
}

func PasswordCheckoutRDPPrivilegeSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"password_checkout_rdp": schema.BoolAttribute{
				Required: true,
			},
		},
		Optional: true,
	}
}

func PasswordCheckoutRDPPrivilegeAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"password_checkout_rdp": types.BoolType,
	}
}

func PasswordCheckoutRDPPrivilegeFromModelToSDK(_ context.Context, in *PasswordCheckoutRDPPrivilegeModel) (*pam.SecurityPolicyPasswordCheckoutRDPPrivilege, diag.Diagnostics) {
	var out pam.SecurityPolicyPasswordCheckoutRDPPrivilege
	var diags diag.Diagnostics

	out.Type = pam.SecurityPolicyRulePrivilegeType_PASSWORD_CHECKOUT_RDP

	if !in.PasswordCheckoutRDP.IsNull() && !in.PasswordCheckoutRDP.IsUnknown() {
		out.PasswordCheckoutRdp = in.PasswordCheckoutRDP.ValueBool()
	}

	return &out, diags
}

func PasswordCheckoutRDPPrivilegeFromSDKToModel(_ context.Context, in *pam.SecurityPolicyPasswordCheckoutRDPPrivilege) (*PasswordCheckoutRDPPrivilegeModel, diag.Diagnostics) {
	var out PasswordCheckoutRDPPrivilegeModel
	var diags diag.Diagnostics

	if val, ok := in.GetPasswordCheckoutRdpOk(); ok {
		out.PasswordCheckoutRDP = types.BoolPointerValue(val)
	}
	return &out, diags
}
