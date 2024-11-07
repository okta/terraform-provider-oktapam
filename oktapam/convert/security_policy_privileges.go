package convert

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SecurityPolicyRulePrivilegeTypeModel types.String

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

func SecurityPolicyRulePrivilegesBlock() schema.Block {
	return schema.SingleNestedBlock{
		Blocks: map[string]schema.Block{
			"password_checkout_database": SecurityPolicyPasswordCheckoutDatabasePrivilegeBlock(),
			"principal_account_ssh":      SecurityPolicyPrincipalAccountSSHPrivilegeBlock(),
			// "password_checkout_rdp":
			// "password_checkout_ssh":
			// "principal_account_rdp":
			// "reveal_password":
			// "secret":
			// "update_password":
		},
	}
}
