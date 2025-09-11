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
	var diags diag.Diagnostics

	if len(in.UserGroups) > 0 {
		var outUserGroups []attr.Value
		for _, userGroup := range in.UserGroups {
			if id, ok := userGroup.GetIdOk(); ok {
				outUserGroups = append(outUserGroups, types.StringValue(*id))
			}
		}

		userGroups, d := types.ListValue(types.StringType, outUserGroups)
		diags.Append(d...)

		if diags.HasError() {
			return nil, diags
		}
		out.UserGroups = userGroups
	}
	return &out, diags
}

func SecurityPolicyPrincipalFromModelToSDK(ctx context.Context, in *SecurityPolicyPrincipalsModel) (*pam.SecurityPolicyPrincipals, diag.Diagnostics) {
	var out pam.SecurityPolicyPrincipals
	var diags diag.Diagnostics

	if !in.UserGroups.IsNull() && len(in.UserGroups.Elements()) > 0 {
		userGroups := make([]types.String, 0, len(in.UserGroups.Elements()))

		diags.Append(in.UserGroups.ElementsAs(ctx, &userGroups, false)...)

		if diags.HasError() {
			return nil, diags
		}

		for _, userGroup := range userGroups {
			out.UserGroups = append(out.UserGroups, *pam.NewNamedObject().SetId(userGroup.ValueString()))
		}
	}
	return &out, diags
}
