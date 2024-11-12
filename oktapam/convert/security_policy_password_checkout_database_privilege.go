package convert

import (
	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Begin SecurityPolicyPasswordCheckoutDatabase

type SecurityPolicyPasswordCheckoutDatabasePrivilegeModel struct {
	PasswordCheckoutDatabase types.Bool `tfsdk:"password_checkout_database"`
}

func SecurityPolicyPasswordCheckoutDatabasePrivilegeSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"password_checkout_database": schema.BoolAttribute{
				Required: true,
			},
		},
		Optional: true,
	}
}

func SecurityPolicyPasswordCheckoutDatabasePrivilegeFromModelToSDK(in *SecurityPolicyPasswordCheckoutDatabasePrivilegeModel, out *pam.SecurityPolicyPasswordCheckoutDatabasePrivilege) diag.Diagnostics {
	out.Type = pam.SecurityPolicyRulePrivilegeType_PASSWORD_CHECKOUT_DATABASE
	out.PasswordCheckoutDatabase = in.PasswordCheckoutDatabase.ValueBool()
	return nil
}

func SecurityPolicyPasswordCheckoutDatabasePrivilegeFromSDKToModel(in *pam.SecurityPolicyPasswordCheckoutDatabasePrivilege, out *SecurityPolicyPasswordCheckoutDatabasePrivilegeModel) diag.Diagnostics {
	out.PasswordCheckoutDatabase = types.BoolValue(in.PasswordCheckoutDatabase)
	return nil
}

// End SecurityPolicyPasswordCheckoutDatabase
