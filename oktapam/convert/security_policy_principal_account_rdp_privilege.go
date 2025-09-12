package convert

import (
	"context"

	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
)

type PrincipalAccountRDPPrivilegeModel struct {
	PrincipalAccountRDP   types.Bool `tfsdk:"principal_account_rdp"`
	AdminLevelPermissions types.Bool `tfsdk:"admin_level_permissions"`
}

func PrincipalAccountRDPPrivilegeSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"principal_account_rdp": schema.BoolAttribute{
				Required:    true,
				Description: descriptions.PrivilegePrincipalAccountRDP,
			},
			"admin_level_permissions": schema.BoolAttribute{
				Optional: true,
				// NOTE: This is an optional boolean, so if we set it to false, platform just stops returning it as it
				// can't tell the difference between "set to false" and "not set".
				Computed:    true,
				Default:     booldefault.StaticBool(false),
				Description: descriptions.AdminLevelPermissions,
			},
		},
		Optional: true,
	}
}

func PrincipalAccountRDPPrivilegeAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"principal_account_rdp":   types.BoolType,
		"admin_level_permissions": types.BoolType,
	}
}

func PrincipalAccountRDPPrivilegeFromModelToSDK(_ context.Context, in *PrincipalAccountRDPPrivilegeModel) (*pam.SecurityPolicyPrincipalAccountRDPPrivilege, diag.Diagnostics) {
	var out pam.SecurityPolicyPrincipalAccountRDPPrivilege
	var diags diag.Diagnostics

	out.Type = pam.SecurityPolicyRulePrivilegeType_PRINCIPAL_ACCOUNT_RDP

	if !in.PrincipalAccountRDP.IsNull() && !in.PrincipalAccountRDP.IsUnknown() {
		out.PrincipalAccountRdp = in.PrincipalAccountRDP.ValueBool()
	}

	if !in.AdminLevelPermissions.IsNull() && !in.AdminLevelPermissions.IsUnknown() {
		out.AdminLevelPermissions = in.AdminLevelPermissions.ValueBoolPointer()
	}

	return &out, diags
}

func PrincipalAccountRDPPrivilegeFromSDKToModel(_ context.Context, in *pam.SecurityPolicyPrincipalAccountRDPPrivilege) (*PrincipalAccountRDPPrivilegeModel, diag.Diagnostics) {
	var out PrincipalAccountRDPPrivilegeModel
	var diags diag.Diagnostics

	if val, ok := in.GetPrincipalAccountRdpOk(); ok {
		out.PrincipalAccountRDP = types.BoolPointerValue(val)
	}

	// NOTE: This is an optional boolean, so if we set it to false, platform just stops returning it as it
	// can't tell the difference between "set to false" and "not set".
	if val, ok := in.GetAdminLevelPermissionsOk(); ok {
		out.AdminLevelPermissions = types.BoolPointerValue(val)
	} else {
		out.AdminLevelPermissions = types.BoolValue(false)
	}
	return &out, diags
}
