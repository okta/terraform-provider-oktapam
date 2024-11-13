package convert

import (
	"context"

	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SecurityPolicyPrincipalsModel struct {
	UserGroups types.List `tfsdk:"groups"` // openapi field: user_groups
}

func SecurityPolicyPrincipalsSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"groups": schema.ListAttribute{
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
	//TODO(ja) there must be a better way to do this, also I'm not checking a lot of things.
	userGroups := make([]types.String, 0, len(in.UserGroups.Elements()))

	if diags := in.UserGroups.ElementsAs(ctx, &userGroups, false); diags.HasError() {
		return diags
	}
	for _, userGroup := range userGroups {
		out.UserGroups = append(out.UserGroups, *pam.NewNamedObject().SetName(userGroup.ValueString()))
	}
	return nil
}
