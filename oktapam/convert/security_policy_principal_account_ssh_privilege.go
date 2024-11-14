package convert

import (
	"context"
	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
)

// Begin SecurityPolicyPrincipalAccountSSH

type SecurityPolicyPrincipalAccountSSHPrivilegeModel struct {
	PrincipalAccountSSH   types.Bool   `tfsdk:"principal_account_ssh"`
	AdminLevelPermissions types.Bool   `tfsdk:"admin_level_permissions"`
	SudoDisplayName       types.String `tfsdk:"sudo_display_name"`
	SudoCommandBundles    []types.String/*NamedObject*/ `tfsdk:"sudo_command_bundles"`
}

func SecurityPolicyPrincipalAccountSSHPrivilegeFromSDKToModel(_ context.Context, in *pam.SecurityPolicyPrincipalAccountSSHPrivilege, out *SecurityPolicyPrincipalAccountSSHPrivilegeModel) diag.Diagnostics {
	out.PrincipalAccountSSH = types.BoolValue(in.PrincipalAccountSsh)
	out.AdminLevelPermissions = types.BoolPointerValue(in.AdminLevelPermissions)
	out.SudoDisplayName = types.StringPointerValue(in.SudoDisplayName)
	//TODO(ja) SudoCommandBundles
	return nil
}

func SecurityPolicyPrincipalAccountSSHPrivilegeFromModelToSDK(_ context.Context, in *SecurityPolicyPrincipalAccountSSHPrivilegeModel, out *pam.SecurityPolicyPrincipalAccountSSHPrivilege) diag.Diagnostics {
	out.Type = pam.SecurityPolicyRulePrivilegeType_PRINCIPAL_ACCOUNT_SSH
	out.PrincipalAccountSsh = in.PrincipalAccountSSH.ValueBool()
	out.AdminLevelPermissions = in.AdminLevelPermissions.ValueBoolPointer()
	out.SudoDisplayName = in.SudoDisplayName.ValueStringPointer()
	//TODO(ja) SudoCommandBundles
	return nil
}

func SecurityPolicyPrincipalAccountSSHPrivilegeSchema() schema.Attribute {
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
