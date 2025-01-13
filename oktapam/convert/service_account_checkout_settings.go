package convert

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int32validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"

	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ServiceAccountCheckoutSettingsModel represents the Terraform model for service account checkout settings
type ServiceAccountCheckoutSettingsModel struct {
	CheckoutRequired          types.Bool  `tfsdk:"checkout_required"`
	CheckoutDurationInSeconds types.Int32 `tfsdk:"checkout_duration_in_seconds"`
	IncludeList               types.List  `tfsdk:"include_list"`
	ExcludeList               types.List  `tfsdk:"exclude_list"`
}

// ServiceAccountSettingNameObjectModel represents the Terraform model for service account setting name object
type ServiceAccountSettingNameObjectModel struct {
	Id                     string `tfsdk:"id"`
	ServiceAccountUserName string `tfsdk:"service_account_user_name"`
	SaasAppInstanceName    string `tfsdk:"saas_app_instance_name"`
}

// ServiceAccountCheckoutSettingsSchemaAttributes returns the schema attributes for service account checkout settings
func ServiceAccountCheckoutSettingsSchemaAttributes(mergeIntoMap map[string]schema.Attribute) map[string]schema.Attribute {
	myMap := map[string]schema.Attribute{
		"checkout_duration_in_seconds": schema.Int32Attribute{
			Required:    true,
			Description: descriptions.CheckoutDurationInSeconds,
			Validators: []validator.Int32{
				int32validator.Between(900, 86400),
			},
		},
		"checkout_required": schema.BoolAttribute{
			Required:    true,
			Description: descriptions.CheckoutRequired,
		},
		"exclude_list": schema.ListAttribute{
			ElementType: types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"id":                        types.StringType,
					"service_account_user_name": types.StringType,
					"saas_app_instance_name":    types.StringType,
				},
			},
			Optional:    true,
			Description: descriptions.ExcludeList,
		},
		"include_list": schema.ListAttribute{
			ElementType: types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"id":                        types.StringType,
					"service_account_user_name": types.StringType,
					"saas_app_instance_name":    types.StringType,
				},
			},
			Optional:    true,
			Description: descriptions.IncludeList,
		},
	}

	for key, value := range myMap {
		mergeIntoMap[key] = value
	}
	return mergeIntoMap
}

// ServiceAccountCheckoutSettingsFromModelToSDK converts from the Terraform model to the SDK type
func ServiceAccountCheckoutSettingsFromModelToSDK(ctx context.Context, in *ServiceAccountCheckoutSettingsModel) (*pam.APIServiceAccountCheckoutSettings, diag.Diagnostics) {
	var out pam.APIServiceAccountCheckoutSettings
	var diags diag.Diagnostics

	if !in.CheckoutRequired.IsNull() && !in.CheckoutRequired.IsUnknown() {
		out.CheckoutRequired = in.CheckoutRequired.ValueBool()
	}
	if !in.CheckoutDurationInSeconds.IsNull() && !in.CheckoutDurationInSeconds.IsUnknown() {
		out.CheckoutDurationInSeconds = in.CheckoutDurationInSeconds.ValueInt32()
	}

	if !in.IncludeList.IsNull() && !in.IncludeList.IsUnknown() {
		var modelList []ServiceAccountSettingNameObjectModel
		diags.Append(in.IncludeList.ElementsAs(ctx, &modelList, false)...)
		if diags.HasError() {
			return nil, diags
		}

		includeList := make([]pam.ServiceAccountSettingNameObject, len(modelList))
		for i, item := range modelList {
			includeList[i] = pam.ServiceAccountSettingNameObject{
				Id:                     item.Id,
				ServiceAccountUserName: &item.ServiceAccountUserName,
				SaasAppInstanceName:    &item.SaasAppInstanceName,
			}
		}
		out.IncludeList = includeList
	}

	if !in.ExcludeList.IsNull() && !in.ExcludeList.IsUnknown() {
		var modelList []ServiceAccountSettingNameObjectModel
		diags.Append(in.ExcludeList.ElementsAs(ctx, &modelList, false)...)
		if diags.HasError() {
			return nil, diags
		}

		excludeList := make([]pam.ServiceAccountSettingNameObject, len(modelList))
		for i, item := range modelList {
			excludeList[i] = pam.ServiceAccountSettingNameObject{
				Id:                     item.Id,
				ServiceAccountUserName: &item.ServiceAccountUserName,
				SaasAppInstanceName:    &item.SaasAppInstanceName,
			}
		}
		out.ExcludeList = excludeList
	}

	return &out, diags
}

// ServiceAccountCheckoutSettingsFromSDKToModel converts from the SDK type to the Terraform model
func ServiceAccountCheckoutSettingsFromSDKToModel(ctx context.Context, in *pam.APIServiceAccountCheckoutSettings) (*ServiceAccountCheckoutSettingsModel, diag.Diagnostics) {
	var out ServiceAccountCheckoutSettingsModel
	var diags diag.Diagnostics

	if val, ok := in.GetCheckoutRequiredOk(); ok {
		out.CheckoutRequired = types.BoolValue(*val)
	}

	if val, ok := in.GetCheckoutDurationInSecondsOk(); ok {
		out.CheckoutDurationInSeconds = types.Int32Value(*val)
	}

	includeList, d := types.ListValueFrom(ctx, types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id":                        types.StringType,
			"service_account_user_name": types.StringType,
			"saas_app_instance_name":    types.StringType,
		},
	}, in.IncludeList)
	diags.Append(d...)
	if diags.HasError() {
		return nil, diags
	}
	out.IncludeList = includeList

	excludeList, d := types.ListValueFrom(ctx, types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id":                        types.StringType,
			"service_account_user_name": types.StringType,
			"saas_app_instance_name":    types.StringType,
		},
	}, in.ExcludeList)
	diags.Append(d...)
	if diags.HasError() {
		return nil, diags
	}
	out.ExcludeList = excludeList

	return &out, diags
}
