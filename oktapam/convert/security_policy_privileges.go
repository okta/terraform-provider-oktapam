package convert

import (
	"context"

	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

//type SecurityPolicyRulePrivilegeTypeModel types.String

//TODO(ja) I think this is modelled wrong

type SecurityPolicyRulePrivilegeContainerModel struct {
	PasswordCheckoutDatabasePrivilege *PasswordCheckoutDatabasePrivilegeModel `tfsdk:"password_checkout_database"`
	PrincipalAccountSSHPrivilege      *PrincipalAccountSSHPrivilegeModel      `tfsdk:"principal_account_ssh"`
	//SecurityPolicyPasswordCheckoutRDPPrivilege      *SecurityPolicyPasswordCheckoutRDPPrivilegeModel      `tfsdk:"password_checkout_rdp"`
	PasswordCheckoutSSHPrivilege *PasswordCheckoutSSHPrivilegeModel `tfsdk:"password_checkout_ssh"`
	//SecurityPolicyPrincipalAccountRDPPrivilege      *SecurityPolicyPrincipalAccountRDPPrivilegeModel      `tfsdk:"principal_account_rdp"`
	//SecurityPolicyRevealPasswordPrivilege           *SecurityPolicyRevealPasswordPrivilegeModel           `tfsdk:"reveal_password"`
	//SecurityPolicySecretPrivilege                   *SecurityPolicySecretPrivilegeModel                   `tfsdk:"secret"`
	//SecurityPolicyUpdatePasswordPrivilege           *SecurityPolicyUpdatePasswordPrivilegeModel           `tfsdk:"update_password"`
}

func SecurityPolicyRulePrivilegeContainerSchema() schema.NestedAttributeObject {
	return schema.NestedAttributeObject{
		Attributes: map[string]schema.Attribute{
			"password_checkout_database": PasswordCheckoutDatabasePrivilegeSchema(),
			"principal_account_ssh":      PrincipalAccountSSHPrivilegeSchema(),
			"password_checkout_ssh":      PasswordCheckoutSSHPrivilegeSchema(),
			// "password_checkout_rdp":
			// "password_checkout_ssh":
			// "principal_account_rdp":
			// "reveal_password":
			// "secret":
			// "update_password":
		},
	}
}

func SecurityPolicyRulePrivilegeContainerAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"password_checkout_database": types.ObjectType{AttrTypes: PasswordCheckoutDatabasePrivilegeAttrTypes()},
		"principal_account_ssh":      types.ObjectType{AttrTypes: PrincipalAccountSSHPrivilegeAttrTypes()},
		"password_checkout_ssh":      types.ObjectType{AttrTypes: PasswordCheckoutSSHPrivilegeAttrTypes()},
	}
}

func SecurityPolicyRulePrivilegeContainerFromSDKToModel(ctx context.Context, in *pam.SecurityPolicyRulePrivilegeContainer) (*SecurityPolicyRulePrivilegeContainerModel, diag.Diagnostics) {
	var out SecurityPolicyRulePrivilegeContainerModel
	var diags diag.Diagnostics

	if in.PrivilegeValue.SecurityPolicyPasswordCheckoutDatabasePrivilege != nil {
		privilege, d := PasswordCheckoutDatabasePrivilegeFromSDKToModel(ctx, in.PrivilegeValue.SecurityPolicyPasswordCheckoutDatabasePrivilege)
		diags.Append(d...)
		if diags.HasError() {
			return nil, diags
		}
		out.PasswordCheckoutDatabasePrivilege = privilege

	} else if in.PrivilegeValue.SecurityPolicyPrincipalAccountSSHPrivilege != nil {
		privilege, d := PrincipalAccountSSHPrivilegeFromSDKToModel(ctx, in.PrivilegeValue.SecurityPolicyPrincipalAccountSSHPrivilege)
		diags.Append(d...)
		if diags.HasError() {
			return nil, diags
		}
		out.PrincipalAccountSSHPrivilege = privilege

	} else if in.PrivilegeValue.SecurityPolicyPasswordCheckoutSSHPrivilege != nil {
		privilege, d := PasswordCheckoutSSHPrivilegeFromSDKToModel(ctx, in.PrivilegeValue.SecurityPolicyPasswordCheckoutSSHPrivilege)
		diags.Append(d...)
		if diags.HasError() {
			return nil, diags
		}
		out.PasswordCheckoutSSHPrivilege = privilege

	} else {
		panic("missing stanza in SecurityPolicyRulePrivilegeContainerFromSDKToModel")
	}
	return &out, diags
}

func SecurityPolicyRulePrivilegeContainerFromModelToSDK(ctx context.Context, in *SecurityPolicyRulePrivilegeContainerModel) (*pam.SecurityPolicyRulePrivilegeContainer, diag.Diagnostics) {
	var out pam.SecurityPolicyRulePrivilegeContainer
	var diags diag.Diagnostics

	var privilegeValue pam.SecurityPolicyRulePrivilegeContainerPrivilegeValue
	var privilegeType pam.SecurityPolicyRulePrivilegeType

	if in.PasswordCheckoutDatabasePrivilege != nil {
		outVal, d := PasswordCheckoutDatabasePrivilegeFromModelToSDK(ctx, in.PasswordCheckoutDatabasePrivilege)
		diags.Append(d...)
		if diags.HasError() {
			return nil, diags
		}
		privilegeValue = pam.SecurityPolicyPasswordCheckoutDatabasePrivilegeAsSecurityPolicyRulePrivilegeContainerPrivilegeValue(outVal)
		privilegeType = pam.SecurityPolicyRulePrivilegeType_PASSWORD_CHECKOUT_DATABASE

	} else if in.PrincipalAccountSSHPrivilege != nil {
		outVal, d := PrincipalAccountSSHPrivilegeFromModelToSDK(ctx, in.PrincipalAccountSSHPrivilege)
		diags.Append(d...)
		if diags.HasError() {
			return nil, diags
		}
		privilegeValue = pam.SecurityPolicyPrincipalAccountSSHPrivilegeAsSecurityPolicyRulePrivilegeContainerPrivilegeValue(outVal)
		privilegeType = pam.SecurityPolicyRulePrivilegeType_PRINCIPAL_ACCOUNT_SSH

	} else if in.PasswordCheckoutSSHPrivilege != nil {
		outVal, d := PasswordCheckoutSSHPrivilegeFromModelToSDK(ctx, in.PasswordCheckoutSSHPrivilege)
		diags.Append(d...)
		if diags.HasError() {
			return nil, diags
		}
		privilegeValue = pam.SecurityPolicyPasswordCheckoutSSHPrivilegeAsSecurityPolicyRulePrivilegeContainerPrivilegeValue(outVal)
		privilegeType = pam.SecurityPolicyRulePrivilegeType_PASSWORD_CHECKOUT_SSH

	} else {
		panic("missing stanza in SecurityPolicyRulePrivilegeContainerFromModelToSDK")
	}
	out.PrivilegeValue = &privilegeValue
	out.PrivilegeType = &privilegeType
	return &out, diags
}
