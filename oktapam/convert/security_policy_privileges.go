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

func SecurityPolicyRulePrivilegeContainerFromSDKToModel(ctx context.Context, in *pam.SecurityPolicyRulePrivilegeContainer, out *SecurityPolicyRulePrivilegeContainerModel) diag.Diagnostics {
	if in.PrivilegeValue.SecurityPolicyPasswordCheckoutDatabasePrivilege != nil {
		var privilege PasswordCheckoutDatabasePrivilegeModel
		if diags := PasswordCheckoutDatabasePrivilegeFromSDKToModel(ctx, in.PrivilegeValue.SecurityPolicyPasswordCheckoutDatabasePrivilege, out.PasswordCheckoutDatabasePrivilege); diags.HasError() {
			return diags
		} else {
			out.PasswordCheckoutDatabasePrivilege = &privilege
		}
	} else if in.PrivilegeValue.SecurityPolicyPrincipalAccountSSHPrivilege != nil {
		var privilege PrincipalAccountSSHPrivilegeModel
		if diags := PrincipalAccountSSHPrivilegeFromSDKToModel(ctx, in.PrivilegeValue.SecurityPolicyPrincipalAccountSSHPrivilege, &privilege); diags.HasError() {
			return diags
		} else {
			out.PrincipalAccountSSHPrivilege = &privilege
		}
	} else if in.PrivilegeValue.SecurityPolicyPasswordCheckoutSSHPrivilege != nil {
		var privilege PasswordCheckoutSSHPrivilegeModel
		if diags := PasswordCheckoutSSHPrivilegeFromSDKToModel(ctx, in.PrivilegeValue.SecurityPolicyPasswordCheckoutSSHPrivilege, &privilege); diags.HasError() {
			return diags
		} else {
			out.PasswordCheckoutSSHPrivilege = &privilege
		}
	} else {
		panic("missing stanza in SecurityPolicyRulePrivilegeContainerFromSDKToModel")
	}
	return nil
}

func SecurityPolicyRulePrivilegeContainerFromModelToSDK(ctx context.Context, in *SecurityPolicyRulePrivilegeContainerModel, out *pam.SecurityPolicyRulePrivilegeContainer) diag.Diagnostics {
	var privilegeValue pam.SecurityPolicyRulePrivilegeContainerPrivilegeValue
	var privilegeType pam.SecurityPolicyRulePrivilegeType

	if in.PasswordCheckoutDatabasePrivilege != nil {
		var outVal pam.SecurityPolicyPasswordCheckoutDatabasePrivilege
		if diags := PasswordCheckoutDatabasePrivilegeFromModelToSDK(ctx, in.PasswordCheckoutDatabasePrivilege, &outVal); diags.HasError() {
			return diags
		}
		privilegeValue = pam.SecurityPolicyPasswordCheckoutDatabasePrivilegeAsSecurityPolicyRulePrivilegeContainerPrivilegeValue(&outVal)
		privilegeType = pam.SecurityPolicyRulePrivilegeType_PASSWORD_CHECKOUT_DATABASE
	} else if in.PrincipalAccountSSHPrivilege != nil {
		var outVal pam.SecurityPolicyPrincipalAccountSSHPrivilege
		if diags := PrincipalAccountSSHPrivilegeFromModelToSDK(ctx, in.PrincipalAccountSSHPrivilege, &outVal); diags.HasError() {
			return diags
		}
		privilegeValue = pam.SecurityPolicyPrincipalAccountSSHPrivilegeAsSecurityPolicyRulePrivilegeContainerPrivilegeValue(&outVal)
		privilegeType = pam.SecurityPolicyRulePrivilegeType_PRINCIPAL_ACCOUNT_SSH
	} else if in.PasswordCheckoutSSHPrivilege != nil {
		var outVal pam.SecurityPolicyPasswordCheckoutSSHPrivilege
		if diags := PasswordCheckoutSSHPrivilegeFromModelToSDK(ctx, in.PasswordCheckoutSSHPrivilege, &outVal); diags.HasError() {
			return diags
		}
		privilegeValue = pam.SecurityPolicyPasswordCheckoutSSHPrivilegeAsSecurityPolicyRulePrivilegeContainerPrivilegeValue(&outVal)
		privilegeType = pam.SecurityPolicyRulePrivilegeType_PASSWORD_CHECKOUT_SSH
	} else {
		panic("missing stanza in SecurityPolicyRulePrivilegeContainerFromModelToSDK")
	}
	out.PrivilegeValue = &privilegeValue
	out.PrivilegeType = &privilegeType
	return nil
}
