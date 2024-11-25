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
	if in.PrivilegeValue.SecurityPolicyPasswordCheckoutDatabasePrivilege != nil {
		if privilege, diags := PasswordCheckoutDatabasePrivilegeFromSDKToModel(ctx, in.PrivilegeValue.SecurityPolicyPasswordCheckoutDatabasePrivilege); diags.HasError() {
			return nil, diags
		} else {
			out.PasswordCheckoutDatabasePrivilege = privilege
		}
	} else if in.PrivilegeValue.SecurityPolicyPrincipalAccountSSHPrivilege != nil {
		if privilege, diags := PrincipalAccountSSHPrivilegeFromSDKToModel(ctx, in.PrivilegeValue.SecurityPolicyPrincipalAccountSSHPrivilege); diags.HasError() {
			return nil, diags
		} else {
			out.PrincipalAccountSSHPrivilege = privilege
		}
	} else if in.PrivilegeValue.SecurityPolicyPasswordCheckoutSSHPrivilege != nil {
		if privilege, diags := PasswordCheckoutSSHPrivilegeFromSDKToModel(ctx, in.PrivilegeValue.SecurityPolicyPasswordCheckoutSSHPrivilege); diags.HasError() {
			return nil, diags
		} else {
			out.PasswordCheckoutSSHPrivilege = privilege
		}
	} else {
		panic("missing stanza in SecurityPolicyRulePrivilegeContainerFromSDKToModel")
	}
	return &out, nil
}

func SecurityPolicyRulePrivilegeContainerFromModelToSDK(ctx context.Context, in *SecurityPolicyRulePrivilegeContainerModel) (*pam.SecurityPolicyRulePrivilegeContainer, diag.Diagnostics) {
	var out pam.SecurityPolicyRulePrivilegeContainer
	var privilegeValue pam.SecurityPolicyRulePrivilegeContainerPrivilegeValue
	var privilegeType pam.SecurityPolicyRulePrivilegeType

	if in.PasswordCheckoutDatabasePrivilege != nil {
		if outVal, diags := PasswordCheckoutDatabasePrivilegeFromModelToSDK(ctx, in.PasswordCheckoutDatabasePrivilege); diags.HasError() {
			return nil, diags
		} else {
			privilegeValue = pam.SecurityPolicyPasswordCheckoutDatabasePrivilegeAsSecurityPolicyRulePrivilegeContainerPrivilegeValue(outVal)
			privilegeType = pam.SecurityPolicyRulePrivilegeType_PASSWORD_CHECKOUT_DATABASE
		}
	} else if in.PrincipalAccountSSHPrivilege != nil {
		if outVal, diags := PrincipalAccountSSHPrivilegeFromModelToSDK(ctx, in.PrincipalAccountSSHPrivilege); diags.HasError() {
			return nil, diags
		} else {
			privilegeValue = pam.SecurityPolicyPrincipalAccountSSHPrivilegeAsSecurityPolicyRulePrivilegeContainerPrivilegeValue(outVal)
			privilegeType = pam.SecurityPolicyRulePrivilegeType_PRINCIPAL_ACCOUNT_SSH
		}
	} else if in.PasswordCheckoutSSHPrivilege != nil {
		if outVal, diags := PasswordCheckoutSSHPrivilegeFromModelToSDK(ctx, in.PasswordCheckoutSSHPrivilege); diags.HasError() {
			return nil, diags
		} else {
			privilegeValue = pam.SecurityPolicyPasswordCheckoutSSHPrivilegeAsSecurityPolicyRulePrivilegeContainerPrivilegeValue(outVal)
			privilegeType = pam.SecurityPolicyRulePrivilegeType_PASSWORD_CHECKOUT_SSH
		}
	} else {
		panic("missing stanza in SecurityPolicyRulePrivilegeContainerFromModelToSDK")
	}
	out.PrivilegeValue = &privilegeValue
	out.PrivilegeType = &privilegeType
	return &out, nil
}
