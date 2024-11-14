package convert

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

//type SecurityPolicyRulePrivilegeTypeModel types.String

type SecurityPolicyRulePrivilegeContainerPrivilegeValueModel struct {
	SecurityPolicyPasswordCheckoutDatabasePrivilege *SecurityPolicyPasswordCheckoutDatabasePrivilegeModel `tfsdk:"password_checkout_database"`
	SecurityPolicyPrincipalAccountSSHPrivilege      *SecurityPolicyPrincipalAccountSSHPrivilegeModel      `tfsdk:"principal_account_ssh"`
	//SecurityPolicyPasswordCheckoutRDPPrivilege      *SecurityPolicyPasswordCheckoutRDPPrivilegeModel      `tfsdk:"password_checkout_rdp"`
	//SecurityPolicyPasswordCheckoutSSHPrivilege *SecurityPolicyPasswordCheckoutSSHPrivilegeModel `tfsdk:"password_checkout_ssh"`
	//SecurityPolicyPrincipalAccountRDPPrivilege      *SecurityPolicyPrincipalAccountRDPPrivilegeModel      `tfsdk:"principal_account_rdp"`
	//SecurityPolicyRevealPasswordPrivilege           *SecurityPolicyRevealPasswordPrivilegeModel           `tfsdk:"reveal_password"`
	//SecurityPolicySecretPrivilege                   *SecurityPolicySecretPrivilegeModel                   `tfsdk:"secret"`
	//SecurityPolicyUpdatePasswordPrivilege           *SecurityPolicyUpdatePasswordPrivilegeModel           `tfsdk:"update_password"`
}

type SecurityPolicyRulePrivilegeContainerModel struct {
	SecurityPolicyRulePrivilegeContainerPrivilegeValueModel
	//PrivilegeType  SecurityPolicyRulePrivilegeTypeModel                    `tfsdk:"privilege_type"`
	//PrivilegeValue SecurityPolicyRulePrivilegeContainerPrivilegeValueModel `tfsdk:"privilege_value"`
}

func SecurityPolicyRulePrivilegesSchema() schema.NestedAttributeObject {
	return schema.NestedAttributeObject{
		Attributes: map[string]schema.Attribute{
			"password_checkout_database": SecurityPolicyPasswordCheckoutDatabasePrivilegeSchema(),
			"principal_account_ssh":      SecurityPolicyPrincipalAccountSSHPrivilegeSchema(),
			// "password_checkout_rdp":
			// "password_checkout_ssh":
			// "principal_account_rdp":
			// "reveal_password":
			// "secret":
			// "update_password":
		},
	}
}
