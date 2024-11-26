package convert

import (
	"context"

	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PasswordCheckoutDatabasePrivilegeModel struct {
	PasswordCheckoutDatabase types.Bool `tfsdk:"password_checkout_database"`
}

func PasswordCheckoutDatabasePrivilegeSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"password_checkout_database": schema.BoolAttribute{
				Required: true,
			},
		},
		Optional: true,
	}
}

func PasswordCheckoutDatabasePrivilegeAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"password_checkout_database": types.BoolType,
	}
}

func PasswordCheckoutDatabasePrivilegeFromModelToSDK(_ context.Context, in *PasswordCheckoutDatabasePrivilegeModel) (*pam.SecurityPolicyPasswordCheckoutDatabasePrivilege, diag.Diagnostics) {
	var out pam.SecurityPolicyPasswordCheckoutDatabasePrivilege
	var diags diag.Diagnostics

	out.Type = pam.SecurityPolicyRulePrivilegeType_PASSWORD_CHECKOUT_DATABASE

	if !in.PasswordCheckoutDatabase.IsNull() && !in.PasswordCheckoutDatabase.IsUnknown() {
		out.PasswordCheckoutDatabase = in.PasswordCheckoutDatabase.ValueBool()
	}

	return &out, diags
}

func PasswordCheckoutDatabasePrivilegeFromSDKToModel(_ context.Context, in *pam.SecurityPolicyPasswordCheckoutDatabasePrivilege) (*PasswordCheckoutDatabasePrivilegeModel, diag.Diagnostics) {
	var out PasswordCheckoutDatabasePrivilegeModel
	var diags diag.Diagnostics

	if val, ok := in.GetPasswordCheckoutDatabaseOk(); ok {
		out.PasswordCheckoutDatabase = types.BoolPointerValue(val)
	}
	return &out, diags
}
