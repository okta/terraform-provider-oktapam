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

func SecurityPolicyPrincipalsBlock() schema.Block {
	return schema.SingleNestedBlock{
		Attributes: map[string]schema.Attribute{
			"groups": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
		},
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
	if diags := in.UserGroups.ElementsAs(ctx, &out.UserGroups, false); diags.HasError() {
		return diags
	}
	return nil
}
