package convert

import (
	"context"

	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SecurityPolicyPrincipalsModel struct {
	UserGroups types.List `tfsdk:"user_groups"`
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

func SecurityPolicyPrincipalFromSDKToModel(ctx context.Context, in *pam.SecurityPolicyPrincipals, out *SecurityPolicyPrincipalsModel) diag.Diagnostics {
	if userGroups, diags := types.ListValueFrom(ctx, types.StringType, in.UserGroups); diags.HasError() {
		return diags
	} else {
		out.UserGroups = userGroups
	}
	return nil
}

func SecurityPolicyPrincipalFromModelToSDK(ctx context.Context, in *SecurityPolicyPrincipalsModel, out *pam.SecurityPolicyPrincipals) diag.Diagnostics {
	userGroups := make([]types.String, 0, len(in.UserGroups.Elements()))

	if diags := in.UserGroups.ElementsAs(ctx, &userGroups, false); diags.HasError() {
		return diags
	}
	for _, userGroup := range userGroups {
		out.UserGroups = append(out.UserGroups, *pam.NewNamedObject().SetId(userGroup.ValueString()))
	}
	return nil
}
