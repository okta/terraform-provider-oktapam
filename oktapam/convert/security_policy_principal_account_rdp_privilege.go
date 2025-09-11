package convert

import (
	"context"

	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PrincipalAccountRDPPrivilegeModel struct {
	PrincipalAccountRDP types.Bool `tfsdk:"principal_account_rdp"`
}

func PrincipalAccountRDPPrivilegeSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"principal_account_rdp": schema.BoolAttribute{
				Required: true,
			},
		},
		Optional: true,
	}
}

func PrincipalAccountRDPPrivilegeAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"principal_account_rdp": types.BoolType,
	}
}

func PrincipalAccountRDPPrivilegeFromModelToSDK(_ context.Context, in *PrincipalAccountRDPPrivilegeModel) (*pam.SecurityPolicyPrincipalAccountRDPPrivilege, diag.Diagnostics) {
	var out pam.SecurityPolicyPrincipalAccountRDPPrivilege
	var diags diag.Diagnostics

	out.Type = pam.SecurityPolicyRulePrivilegeType_PRINCIPAL_ACCOUNT_RDP

	if !in.PrincipalAccountRDP.IsNull() && !in.PrincipalAccountRDP.IsUnknown() {
		out.PrincipalAccountRdp = in.PrincipalAccountRDP.ValueBool()
	}

	return &out, diags
}

func PrincipalAccountRDPPrivilegeFromSDKToModel(_ context.Context, in *pam.SecurityPolicyPrincipalAccountRDPPrivilege) (*PrincipalAccountRDPPrivilegeModel, diag.Diagnostics) {
	var out PrincipalAccountRDPPrivilegeModel
	var diags diag.Diagnostics

	if val, ok := in.GetPrincipalAccountRdpOk(); ok {
		out.PrincipalAccountRDP = types.BoolPointerValue(val)
	}
	return &out, diags
}
