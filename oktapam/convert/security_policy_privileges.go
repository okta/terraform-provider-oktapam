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

type SecurityPolicyRulePrivilegeContainerModel struct {
	PasswordCheckoutDatabasePrivilege *PasswordCheckoutDatabasePrivilegeModel `tfsdk:"password_checkout_database"`
	PrincipalAccountSSHPrivilege      *PrincipalAccountSSHPrivilegeModel      `tfsdk:"principal_account_ssh"`
	PasswordCheckoutSSHPrivilege      *PasswordCheckoutSSHPrivilegeModel      `tfsdk:"password_checkout_ssh"`
	RevealPasswordPrivilege           *RevealPasswordPrivilegeModel           `tfsdk:"reveal_password"`
	//SecurityPolicyPasswordCheckoutRDPPrivilege      *SecurityPolicyPasswordCheckoutRDPPrivilegeModel      `tfsdk:"password_checkout_rdp"`
	//SecurityPolicyPrincipalAccountRDPPrivilege      *SecurityPolicyPrincipalAccountRDPPrivilegeModel      `tfsdk:"principal_account_rdp"`
	//SecurityPolicySecretPrivilege                   *SecurityPolicySecretPrivilegeModel                   `tfsdk:"secret"`
	//SecurityPolicyUpdatePasswordPrivilege           *SecurityPolicyUpdatePasswordPrivilegeModel           `tfsdk:"update_password"`
}

func SecurityPolicyRulePrivilegeContainerSchema() schema.NestedAttributeObject {
	return schema.NestedAttributeObject{
		Attributes: map[string]schema.Attribute{
			"password_checkout_database": PasswordCheckoutDatabasePrivilegeSchema(),
			"principal_account_ssh":      PrincipalAccountSSHPrivilegeSchema(),
			"password_checkout_ssh":      PasswordCheckoutSSHPrivilegeSchema(),
			"reveal_password":            RevealPasswordPrivilegeSchema(),
			// "password_checkout_rdp":
			// "principal_account_rdp":
			// "secret":
			// "update_password":
			// "rotate_password":
		},
	}
}

func SecurityPolicyRulePrivilegeContainerAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"password_checkout_database": types.ObjectType{AttrTypes: PasswordCheckoutDatabasePrivilegeAttrTypes()},
		"principal_account_ssh":      types.ObjectType{AttrTypes: PrincipalAccountSSHPrivilegeAttrTypes()},
		"password_checkout_ssh":      types.ObjectType{AttrTypes: PasswordCheckoutSSHPrivilegeAttrTypes()},
		"reveal_password":            types.ObjectType{AttrTypes: RevealPasswordPrivilegeAttrTypes()},
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

	} else if in.PrivilegeValue.SecurityPolicyRevealPasswordPrivilege != nil {
		privilege, d := RevealPasswordPrivilegeFromSDKToModel(ctx, in.PrivilegeValue.SecurityPolicyRevealPasswordPrivilege)
		diags.Append(d...)
		if diags.HasError() {
			return nil, diags
		}
		out.RevealPasswordPrivilege = privilege

	} else {
		diags.AddError("missing stanza in OktaPAM provider", "missing stanza in SecurityPolicyRulePrivilegeContainerFromSDKToModel")
		return nil, diags
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

	} else if in.RevealPasswordPrivilege != nil {
		outVal, d := RevealPasswordPrivilegeFromModelToSDK(ctx, in.RevealPasswordPrivilege)
		diags.Append(d...)
		if diags.HasError() {
			return nil, diags
		}
		privilegeValue = pam.SecurityPolicyRevealPasswordPrivilegeAsSecurityPolicyRulePrivilegeContainerPrivilegeValue(outVal)
		privilegeType = pam.SecurityPolicyRulePrivilegeType_REVEAL_PASSWORD

	} else {
		diags.AddError("unknown or missing privilege listed in policy rule",
			"One of the privileges listed in this policy is either incorrect "+
				"or unknown to this version of the OktaPAM Terraform provider. Please make "+
				"sure each of your privileges are correct, and you're using the latest available version of "+
				"the OktaPAM Terraform provider. If you've done these things, it could be that the "+
				"privilege you're using is not yet supported and you are encouraged to file an issue in our GitHub repository.")
		return nil, diags
	}
	out.PrivilegeValue = &privilegeValue
	out.PrivilegeType = &privilegeType
	return &out, diags
}
