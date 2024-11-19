package convert

import (
	"context"
	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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

func PrincipalAccountSSHPrivilegeFromSDKToModel(_ context.Context, in *pam.SecurityPolicyPrincipalAccountSSHPrivilege, out *PrincipalAccountSSHPrivilegeModel) diag.Diagnostics {
	out.PrincipalAccountSSH = types.BoolValue(in.PrincipalAccountSsh)
	out.AdminLevelPermissions = types.BoolPointerValue(in.AdminLevelPermissions)
	out.SudoDisplayName = types.StringPointerValue(in.SudoDisplayName)
	for _, sudoCommandBundle := range in.SudoCommandBundles {
		out.SudoCommandBundles = append(out.SudoCommandBundles, types.StringPointerValue(sudoCommandBundle.Id))
	}
	return nil
}

func PrincipalAccountSSHPrivilegeFromModelToSDK(_ context.Context, in *PrincipalAccountSSHPrivilegeModel, out *pam.SecurityPolicyPrincipalAccountSSHPrivilege) diag.Diagnostics {
	out.Type = pam.SecurityPolicyRulePrivilegeType_PRINCIPAL_ACCOUNT_SSH
	out.PrincipalAccountSsh = in.PrincipalAccountSSH.ValueBool()
	out.AdminLevelPermissions = in.AdminLevelPermissions.ValueBoolPointer()
	out.SudoDisplayName = in.SudoDisplayName.ValueStringPointer()
	for _, sudoCommandBundle := range in.SudoCommandBundles {
		if !sudoCommandBundle.IsNull() && !sudoCommandBundle.IsUnknown() {
			out.SudoCommandBundles = append(out.SudoCommandBundles, *pam.NewNamedObject().SetId(sudoCommandBundle.ValueString()))
		}
	}
	return nil
}

func PrincipalAccountSSHPrivilegeSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"principal_account_ssh": schema.BoolAttribute{
				Required:    true,
				Description: descriptions.PrivilegePrincipalAccountSSH,
			},
			"admin_level_permissions": schema.BoolAttribute{
				Optional:    true,
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
