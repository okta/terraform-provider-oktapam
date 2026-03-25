package convert

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// PasswordCheckoutDatabasePrivilegeModel is kept in the schema for backward compatibility with existing state.
// The API no longer supports this privilege type, but old state files may still contain it as null.
// Removing it from the schema would cause "unsupported attribute" errors during Terraform plan JSON processing.
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
		Optional:           true,
		DeprecationMessage: "The attribute is deprecated and will be removed in a future major version.",
	}
}

func PasswordCheckoutDatabasePrivilegeAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"password_checkout_database": types.BoolType,
	}
}
