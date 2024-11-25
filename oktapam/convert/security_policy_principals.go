package convert

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"

	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SecurityPolicyPrincipalsModel struct {
	UserGroups types.List /*NamedObject*/ `tfsdk:"user_groups"`
}

func SecurityPolicyPrincipalsSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"user_groups": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
		},
		Required: true,
	}
}

func SecurityPolicyPrincipalFromSDKToModel(_ context.Context, in *pam.SecurityPolicyPrincipals) (*SecurityPolicyPrincipalsModel, diag.Diagnostics) {
	var out SecurityPolicyPrincipalsModel
	if len(in.UserGroups) > 0 {
		var outUserGroups []attr.Value
		for _, userGroup := range in.UserGroups {
			if id, ok := userGroup.GetIdOk(); ok {
				outUserGroups = append(outUserGroups, types.StringValue(*id))
			}
		}

		if userGroups, diags := types.ListValue(types.StringType, outUserGroups); diags.HasError() {
			return nil, diags
		} else {
			out.UserGroups = userGroups
		}
	}
	return &out, nil
}

func SecurityPolicyPrincipalFromModelToSDK(ctx context.Context, in *SecurityPolicyPrincipalsModel) (*pam.SecurityPolicyPrincipals, diag.Diagnostics) {
	var out pam.SecurityPolicyPrincipals

	if !in.UserGroups.IsNull() && len(in.UserGroups.Elements()) > 0 {
		userGroups := make([]types.String, 0, len(in.UserGroups.Elements()))

		if diags := in.UserGroups.ElementsAs(ctx, &userGroups, false); diags.HasError() {
			return nil, diags
		}
		for _, userGroup := range userGroups {
			out.UserGroups = append(out.UserGroups, *pam.NewNamedObject().SetId(userGroup.ValueString()))
		}
	}
	return &out, nil
}
