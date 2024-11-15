package convert

import (
	"context"
	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

//type SecurityPolicyRulePrivilegeTypeModel types.String

//TODO(ja) I think this is modelled wrong

type SecurityPolicyRulePrivilegeContainerModel struct {
	SecurityPolicyPasswordCheckoutDatabasePrivilege *SecurityPolicyPasswordCheckoutDatabasePrivilegeModel `tfsdk:"password_checkout_database"`
	SecurityPolicyPrincipalAccountSSHPrivilege      *SecurityPolicyPrincipalAccountSSHPrivilegeModel      `tfsdk:"principal_account_ssh"`
	//SecurityPolicyPasswordCheckoutRDPPrivilege      *SecurityPolicyPasswordCheckoutRDPPrivilegeModel      `tfsdk:"password_checkout_rdp"`
	SecurityPolicyPasswordCheckoutSSHPrivilege *SecurityPolicyPasswordCheckoutSSHPrivilegeModel `tfsdk:"password_checkout_ssh"`
	//SecurityPolicyPrincipalAccountRDPPrivilege      *SecurityPolicyPrincipalAccountRDPPrivilegeModel      `tfsdk:"principal_account_rdp"`
	//SecurityPolicyRevealPasswordPrivilege           *SecurityPolicyRevealPasswordPrivilegeModel           `tfsdk:"reveal_password"`
	//SecurityPolicySecretPrivilege                   *SecurityPolicySecretPrivilegeModel                   `tfsdk:"secret"`
	//SecurityPolicyUpdatePasswordPrivilege           *SecurityPolicyUpdatePasswordPrivilegeModel           `tfsdk:"update_password"`
}

func SecurityPolicyRulePrivilegeContainerSchema() schema.NestedAttributeObject {
	return schema.NestedAttributeObject{
		Attributes: map[string]schema.Attribute{
			"password_checkout_database": SecurityPolicyPasswordCheckoutDatabasePrivilegeSchema(),
			"principal_account_ssh":      SecurityPolicyPrincipalAccountSSHPrivilegeSchema(),
			"password_checkout_ssh":      SecurityPolicyPasswordCheckoutSSHPrivilegeSchema(),
			// "password_checkout_rdp":
			// "password_checkout_ssh":
			// "principal_account_rdp":
			// "reveal_password":
			// "secret":
			// "update_password":
		},
	}
}

func SecurityPolicyRulePrivilegeContainerFromSDKToModel(ctx context.Context, in *pam.SecurityPolicyRulePrivilegeContainer, out *SecurityPolicyRulePrivilegeContainerModel) diag.Diagnostics {
	if in.PrivilegeValue.SecurityPolicyPasswordCheckoutDatabasePrivilege != nil {
		if diags := SecurityPolicyPasswordCheckoutDatabasePrivilegeFromSDKToModel(ctx, in.PrivilegeValue.SecurityPolicyPasswordCheckoutDatabasePrivilege, out.SecurityPolicyPasswordCheckoutDatabasePrivilege); diags.HasError() {
			return diags
		}
	} else if in.PrivilegeValue.SecurityPolicyPrincipalAccountSSHPrivilege != nil {
		if diags := SecurityPolicyPrincipalAccountSSHPrivilegeFromSDKToModel(ctx, in.PrivilegeValue.SecurityPolicyPrincipalAccountSSHPrivilege, out.SecurityPolicyPrincipalAccountSSHPrivilege); diags.HasError() {
			return diags
		}
	} else if in.PrivilegeValue.SecurityPolicyPasswordCheckoutSSHPrivilege != nil {
		if diags := SecurityPolicyPasswordCheckoutSSHPrivilegeFromSDKToModel(ctx, in.PrivilegeValue.SecurityPolicyPasswordCheckoutSSHPrivilege, out.SecurityPolicyPasswordCheckoutSSHPrivilege); diags.HasError() {
			return diags
		}
	} else {
		panic("missing stanza in SecurityPolicyRulePrivilegeContainerFromSDKToModel")
	}
	return nil
}

func SecurityPolicyRulePrivilegeContainerFromModelToSDK(ctx context.Context, in *SecurityPolicyRulePrivilegeContainerModel, out *pam.SecurityPolicyRulePrivilegeContainer) diag.Diagnostics {
	var privilegeValue pam.SecurityPolicyRulePrivilegeContainerPrivilegeValue
	var privilegeType pam.SecurityPolicyRulePrivilegeType
	if in.SecurityPolicyPasswordCheckoutDatabasePrivilege != nil {
		var outVal pam.SecurityPolicyPasswordCheckoutDatabasePrivilege
		if diags := SecurityPolicyPasswordCheckoutDatabasePrivilegeFromModelToSDK(ctx, in.SecurityPolicyPasswordCheckoutDatabasePrivilege, &outVal); diags.HasError() {
			return diags
		}
		privilegeValue = pam.SecurityPolicyPasswordCheckoutDatabasePrivilegeAsSecurityPolicyRulePrivilegeContainerPrivilegeValue(&outVal)
		privilegeType = pam.SecurityPolicyRulePrivilegeType_PASSWORD_CHECKOUT_DATABASE
		out.PrivilegeValue = &privilegeValue
	} else if in.SecurityPolicyPrincipalAccountSSHPrivilege != nil {
		var outVal pam.SecurityPolicyPrincipalAccountSSHPrivilege
		if diags := SecurityPolicyPrincipalAccountSSHPrivilegeFromModelToSDK(ctx, in.SecurityPolicyPrincipalAccountSSHPrivilege, &outVal); diags.HasError() {
			return diags
		}
		privilegeValue = pam.SecurityPolicyPrincipalAccountSSHPrivilegeAsSecurityPolicyRulePrivilegeContainerPrivilegeValue(&outVal)
		privilegeType = pam.SecurityPolicyRulePrivilegeType_PRINCIPAL_ACCOUNT_SSH
		out.PrivilegeValue = &privilegeValue
	} else if in.SecurityPolicyPasswordCheckoutSSHPrivilege != nil {
		var outVal pam.SecurityPolicyPasswordCheckoutSSHPrivilege
		if diags := SecurityPolicyPasswordCheckoutSSHPrivilegeFromModelToSDK(ctx, in.SecurityPolicyPasswordCheckoutSSHPrivilege, &outVal); diags.HasError() {
			return diags
		}
		privilegeValue = pam.SecurityPolicyPasswordCheckoutSSHPrivilegeAsSecurityPolicyRulePrivilegeContainerPrivilegeValue(&outVal)
		privilegeType = pam.SecurityPolicyRulePrivilegeType_PASSWORD_CHECKOUT_SSH
		out.PrivilegeValue = &privilegeValue
	} else {
		panic("missing stanza in SecurityPolicyRulePrivilegeContainerFromModelToSDK")
	}
	out.PrivilegeValue = &privilegeValue
	out.PrivilegeType = &privilegeType
	return nil
}
