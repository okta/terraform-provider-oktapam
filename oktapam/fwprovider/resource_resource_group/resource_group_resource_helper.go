package resource_resource_group

import (
	"context"
	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (v DelegatedResourceAdminGroupsValue) ToPamNamedObject(ctx context.Context) (*pam.NamedObject, diag.Diagnostics) {
	var diags diag.Diagnostics

	if v.IsNull() {
		return nil, diags
	}

	if v.IsUnknown() {
		diags.Append(diag.NewErrorDiagnostic(
			"DelegatedResourceAdminGroupsValue Value Is Unknown",
			`"DelegatedResourceAdminGroupsValue" is unknown.`,
		))

		return nil, diags
	}

	namedObjectType, _ := pam.NewNamedObjectTypeFromValue(v.DelegatedResourceAdminGroupsType.ValueString())

	return &pam.NamedObject{
		Id:   v.Id.ValueStringPointer(),
		Name: v.Name.ValueStringPointer(),
		Type: namedObjectType,
	}, diags
}

func (v DelegatedResourceAdminGroupsValue) FromPamNamedObject(ctx context.Context, apiObject *pam.NamedObject) (DelegatedResourceAdminGroupsValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	if apiObject == nil {
		return NewDelegatedResourceAdminGroupsValueNull(), diags
	}

	var namedObjectStringPointer *string
	if apiObject.Type != nil {
		namedObjectStrVal := string(*apiObject.Type)
		namedObjectStringPointer = &namedObjectStrVal
	}
	return DelegatedResourceAdminGroupsValue{
		Id:                               types.StringPointerValue(apiObject.Id),
		Name:                             types.StringPointerValue(apiObject.Name),
		DelegatedResourceAdminGroupsType: types.StringPointerValue(namedObjectStringPointer),
		state:                            attr.ValueStateKnown,
	}, diags
}