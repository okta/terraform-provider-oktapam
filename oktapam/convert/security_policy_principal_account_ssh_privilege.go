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

// Begin SecurityPolicyPrincipalAccountSSH

type PrincipalAccountSSHPrivilegeModel struct {
	PrincipalAccountSSH   types.Bool   `tfsdk:"principal_account_ssh"`
	AdminLevelPermissions types.Bool   `tfsdk:"admin_level_permissions"`
	SudoDisplayName       types.String `tfsdk:"sudo_display_name"`
	SudoCommandBundles    []types.String/*NamedObject*/ `tfsdk:"sudo_command_bundles"`
}

func PrincipalAccountSSHPrivilegeFromSDKToModel(_ context.Context, in *pam.SecurityPolicyPrincipalAccountSSHPrivilege) (*PrincipalAccountSSHPrivilegeModel, diag.Diagnostics) {
	var out PrincipalAccountSSHPrivilegeModel
	var diags diag.Diagnostics

	if val, ok := in.GetPrincipalAccountSshOk(); ok {
		out.PrincipalAccountSSH = types.BoolPointerValue(val)
	}

	// NOTE: This is an optional boolean, so if we set it to false, platform just stops returning it as it
	// can't tell the difference between "set to false" and "not set".
	if val, ok := in.GetAdminLevelPermissionsOk(); ok {
		out.AdminLevelPermissions = types.BoolPointerValue(val)
	} else {
		out.AdminLevelPermissions = types.BoolValue(false)
	}

	if val, ok := in.GetSudoDisplayNameOk(); ok {
		out.SudoDisplayName = types.StringPointerValue(val)
	}
	for _, sudoCommandBundle := range in.SudoCommandBundles {
		out.SudoCommandBundles = append(out.SudoCommandBundles, types.StringPointerValue(sudoCommandBundle.Id))
	}
	return &out, diags
}

func PrincipalAccountSSHPrivilegeFromModelToSDK(_ context.Context, in *PrincipalAccountSSHPrivilegeModel) (*pam.SecurityPolicyPrincipalAccountSSHPrivilege, diag.Diagnostics) {
	var out pam.SecurityPolicyPrincipalAccountSSHPrivilege
	var diags diag.Diagnostics

	out.Type = pam.SecurityPolicyRulePrivilegeType_PRINCIPAL_ACCOUNT_SSH

	if !in.PrincipalAccountSSH.IsNull() && !in.PrincipalAccountSSH.IsUnknown() {
		out.PrincipalAccountSsh = in.PrincipalAccountSSH.ValueBool()
	}

	if !in.AdminLevelPermissions.IsNull() && !in.AdminLevelPermissions.IsUnknown() {
		out.AdminLevelPermissions = in.AdminLevelPermissions.ValueBoolPointer()
	}

	if !in.SudoDisplayName.IsNull() && !in.SudoDisplayName.IsUnknown() {
		out.SudoDisplayName = in.SudoDisplayName.ValueStringPointer()
	}

	for _, sudoCommandBundle := range in.SudoCommandBundles {
		if !sudoCommandBundle.IsNull() && !sudoCommandBundle.IsUnknown() {
			out.SudoCommandBundles = append(out.SudoCommandBundles, *pam.NewNamedObject().SetId(sudoCommandBundle.ValueString()))
		}
	}
	return &out, diags
}

func PrincipalAccountSSHPrivilegeSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"principal_account_ssh": schema.BoolAttribute{
				Required:    true,
				Description: descriptions.PrivilegePrincipalAccountSSH,
			},
			"admin_level_permissions": schema.BoolAttribute{
				Optional: true,
				// NOTE: This is an optional boolean, so if we set it to false, platform just stops returning it as it
				// can't tell the difference between "set to false" and "not set".
				Computed:    true,
				Default:     booldefault.StaticBool(false),
				Description: descriptions.AdminLevelPermissions,
			},
			"sudo_display_name": schema.StringAttribute{
				Optional:    true,
				Description: descriptions.SudoDisplayName,
			},
			"sudo_command_bundles": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: descriptions.SudoCommandBundles,
			},
		},
		Optional: true,
	}
}

func PrincipalAccountSSHPrivilegeAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"principal_account_ssh":   types.BoolType,
		"admin_level_permissions": types.BoolType,
		"sudo_display_name":       types.StringType,
		"sudo_command_bundles":    types.ListType{ElemType: types.StringType},
	}
}
