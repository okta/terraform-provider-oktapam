package convert

import (
	"context"
	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SecurityPolicyRuleConditionContainerModel struct {
	ConditionsAccessRequests *ConditionsAccessRequestsModel `tfsdk:"access_request"`
	ConditionsGateway        *ConditionsGatewayModel        `tfsdk:"gateway"`
	ConditionsMFA            *ConditionsMFAModel            `tfsdk:"mfa"`
}

func SecurityPolicyRuleConditionContainerSchema() schema.NestedAttributeObject {
	return schema.NestedAttributeObject{
		Attributes: map[string]schema.Attribute{
			"access_request": ConditionsAccessRequestsSchema(),
			"gateway":        ConditionsGatewaySchema(),
			"mfa":            ConditionsMFASchema(),
		},
	}
}

func SecurityPolicyRuleConditionContainerAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"access_request": types.ObjectType{AttrTypes: ConditionsAccessRequestsAttrTypes()},
		"gateway":        types.ObjectType{AttrTypes: ConditionsGatewayAttrTypes()},
		"mfa":            types.ObjectType{AttrTypes: ConditionsMFAAttrTypes()},
	}
}
func SecurityPolicyRuleConditionFromModelToSDK(ctx context.Context, in *SecurityPolicyRuleConditionContainerModel) (*pam.SecurityPolicyRuleCondition, diag.Diagnostics) {
	var out pam.SecurityPolicyRuleCondition
	return &out, nil
}

func SecurityPolicyRuleConditionContainerFromSDKToModel(ctx context.Context, in *pam.SecurityPolicyRuleCondition) (*SecurityPolicyRuleConditionContainerModel, diag.Diagnostics) {
	var out SecurityPolicyRuleConditionContainerModel
	return &out, nil
}

type ConditionsAccessRequestsModel struct {
	RequestTypeId       types.String `tfsdk:"request_type_id"`
	RequestTypeName     types.String `tfsdk:"request_type_name"`
	ExpiresAfterSeconds types.Int32  `tfsdk:"expires_after_seconds"`
}

func ConditionsAccessRequestsSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"request_type_id":       schema.StringAttribute{Optional: true},
			"request_type_name":     schema.StringAttribute{Required: true},
			"expires_after_seconds": schema.Int32Attribute{Optional: true},
		},
		Optional: true,
	}
}

func ConditionsAccessRequestsAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"request_type_id":       types.StringType,
		"request_type_name":     types.StringType,
		"expires_after_seconds": types.Int32Type,
	}
}

func ConditionsAccessRequestFromSDKToModel(ctx context.Context, in *pam.ConditionsAccessRequests) (*ConditionsAccessRequestsModel, diag.Diagnostics) {
	var out ConditionsAccessRequestsModel
	return &out, nil
}
func ConditionsAccessRequestFromModelToSDK(ctx context.Context, in *ConditionsAccessRequestsModel) (*pam.ConditionsAccessRequests, diag.Diagnostics) {
	var out pam.ConditionsAccessRequests

	return &out, nil
}

type ConditionsGatewayModel struct {
	TrafficForwarding types.Bool `tfsdk:"traffic_forwarding"`
	SessionRecording  types.Bool `tfsdk:"session_recording"`
}

func ConditionsGatewaySchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"traffic_forwarding": schema.BoolAttribute{Optional: true},
			"session_recording":  schema.BoolAttribute{Optional: true},
		},
		Optional: true,
	}
}

func ConditionsGatewayAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"traffic_forwarding": types.BoolType,
		"session_recording":  types.BoolType,
	}
}

func ConditionsGatewayFromSDKToModel(ctx context.Context, in *pam.ConditionsGateway) (*ConditionsGatewayModel, diag.Diagnostics) {
	var out ConditionsGatewayModel

	return &out, nil
}

func ConditionsGatewayFromModelToSDK(ctx context.Context, in *ConditionsGatewayModel) (*pam.ConditionsGateway, diag.Diagnostics) {
	var out pam.ConditionsGateway

	return &out, nil
}

type ConditionsMFAModel struct {
	ReAuthFrequencyInSeconds types.Int32 `tfsdk:"re_auth_frequency_in_seconds"`
	AcrValues                types.String/*ConditionsMFAACRValuesModel*/ `tfsdk:"acr_values"`
}

func ConditionsMFASchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"re_auth_frequency_in_seconds": schema.Int32Attribute{Optional: true},
			"acr_values":                   schema.StringAttribute{Optional: true},
		},
		Optional: true,
	}
}

func ConditionsMFAAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"re_auth_frequency_in_seconds": types.Int32Type,
		"acr_values":                   types.StringType,
	}
}

func ConditionsMFAFromSDKToModel(ctx context.Context, in *pam.ConditionsMFA) (*ConditionsMFAModel, diag.Diagnostics) {
	var out ConditionsMFAModel

	return &out, nil
}
func ConditionsMFAFromModelToSDK(ctx context.Context, in *ConditionsMFAModel) (*pam.ConditionsMFA, diag.Diagnostics) {
	var out pam.ConditionsMFA

	return &out, nil
}
