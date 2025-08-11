package convert

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"

	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NameObjectModel struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Type types.String `tfsdk:"type"`
}

type SecurityPolicyPrincipalsModel struct {
	UserGroups types.List /*NameObjectModel*/ `tfsdk:"user_groups"`
}

func NameObjectAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":   types.StringType,
		"name": types.StringType,
		"type": types.StringType,
	}
}

func SecurityPolicyPrincipalsSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"user_groups": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Required: true,
						},
						"name": schema.StringAttribute{
							Optional: true,
						},
						"type": schema.StringAttribute{
							Optional: true,
						},
					},
				},
				Optional: true,
			},
		},
		Required: true,
	}
}

func SecurityPolicyPrincipalFromSDKToModel(ctx context.Context, in *pam.SecurityPolicyPrincipals) (*SecurityPolicyPrincipalsModel, diag.Diagnostics) {
	var out SecurityPolicyPrincipalsModel
	var diags diag.Diagnostics

	if len(in.UserGroups) > 0 {
		var outUserGroups []NameObjectModel
		for _, userGroup := range in.UserGroups {
			var userGroupModel NameObjectModel

			if id, ok := userGroup.GetIdOk(); ok {
				userGroupModel.ID = types.StringValue(*id)
			}
			if name, ok := userGroup.GetNameOk(); ok {
				userGroupModel.Name = types.StringValue(*name)
			}
			if userGroupType, ok := userGroup.GetTypeOk(); ok {
				userGroupModel.Type = types.StringValue(string(*userGroupType))
			}

			outUserGroups = append(outUserGroups, userGroupModel)
		}

		userGroups, d := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: NameObjectAttrTypes()}, outUserGroups)
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
		userGroups := make([]NameObjectModel, 0, len(in.UserGroups.Elements()))

		diags.Append(in.UserGroups.ElementsAs(ctx, &userGroups, false)...)

		if diags.HasError() {
			return nil, diags
		}

		for _, userGroup := range userGroups {
			namedObject := pam.NewNamedObject()

			if !userGroup.ID.IsNull() {
				namedObject.SetId(userGroup.ID.ValueString())
			}
			if !userGroup.Name.IsNull() {
				namedObject.SetName(userGroup.Name.ValueString())
			}
			if !userGroup.Type.IsNull() {
				namedObject.SetType(pam.NamedObjectType(userGroup.Type.ValueString()))
			}

			out.UserGroups = append(out.UserGroups, *namedObject)
		}
	}
	return &out, diags
}
