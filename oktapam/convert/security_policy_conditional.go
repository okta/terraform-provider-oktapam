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
func SecurityPolicyRuleConditionContainerFromModelToSDK(ctx context.Context, in *SecurityPolicyRuleConditionContainerModel) (*pam.SecurityPolicyRuleConditionContainer, diag.Diagnostics) {
	var outType pam.SecurityPolicyRuleConditionType
	var outCondition pam.SecurityPolicyRuleCondition
	var diags diag.Diagnostics

	if in.ConditionsMFA != nil {
		outVal, d := ConditionsMFAFromModelToSDK(ctx, in.ConditionsMFA)
		diags.Append(d...)
		if diags.HasError() {
			return nil, diags
		}
		outType = pam.SecurityPolicyRuleConditionType_MFA
		outCondition.ConditionsMFA = outVal

	} else if in.ConditionsAccessRequests != nil {
		outVal, d := ConditionsAccessRequestFromModelToSDK(ctx, in.ConditionsAccessRequests)
		diags.Append(d...)
		if diags.HasError() {
			return nil, diags
		}
		outType = pam.SecurityPolicyRuleConditionType_ACCESS_REQUEST
		outCondition.ConditionsAccessRequests = outVal

	} else if in.ConditionsGateway != nil {
		outVal, d := ConditionsGatewayFromModelToSDK(ctx, in.ConditionsGateway)
		diags.Append(d...)
		if diags.HasError() {
			return nil, diags
		}
		outType = pam.SecurityPolicyRuleConditionType_GATEWAY
		outCondition.ConditionsGateway = outVal
	} else {
		diags.AddError("unknown or missing condition listed in policy rule",
			"One of the conditions listed in this policy is either incorrect "+
				"or unknown to this version of the OktaPAM Terraform Provider. Please make "+
				"sure each of your conditions are correct, and you're using the latest available version of "+
				"the OktaPAM Terraform provider. If you've done these things, it could be that the "+
				"privilege you're using is not yet supported and you are encouraged to file an issue in "+
				"our GitHub repository.")
		return nil, diags
	}

	out := pam.NewSecurityPolicyRuleConditionContainer().
		SetConditionValue(outCondition).
		SetConditionType(outType)

	return out, diags
}

func SecurityPolicyRuleConditionContainerFromSDKToModel(ctx context.Context, in *pam.SecurityPolicyRuleConditionContainer) (*SecurityPolicyRuleConditionContainerModel, diag.Diagnostics) {
	var out SecurityPolicyRuleConditionContainerModel
	var diags diag.Diagnostics

	if !in.HasConditionType() {
		diags.AddError("No ConditionType sent from platform", "A ConditionType is required to convert a Condition")
		return nil, diags
	}

	switch *in.ConditionType {
	case pam.SecurityPolicyRuleConditionType_MFA:
		outModel, d := ConditionsMFAFromSDKToModel(ctx, in.ConditionValue.ConditionsMFA)
		diags.Append(d...)
		if diags.HasError() {
			return nil, diags
		}
		out.ConditionsMFA = outModel

	case pam.SecurityPolicyRuleConditionType_ACCESS_REQUEST:
		outModel, d := ConditionsAccessRequestFromSDKToModel(ctx, in.ConditionValue.ConditionsAccessRequests)
		diags.Append(d...)
		if diags.HasError() {
			return nil, diags
		}
		out.ConditionsAccessRequests = outModel
	case pam.SecurityPolicyRuleConditionType_GATEWAY:
		outModel, d := ConditionsGatewayFromSDKToModel(ctx, in.ConditionValue.ConditionsGateway)
		diags.Append(d...)
		if diags.HasError() {
			return nil, diags
		}
		out.ConditionsGateway = outModel
	default:
		diags.AddError("missing stanza in OktaPAM provider", "missing condition type stanza in SecurityPolicyRuleConditionContainerFromSDKToModel")
		return nil, diags
	}
	return &out, diags
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

func ConditionsAccessRequestFromSDKToModel(_ context.Context, in *pam.ConditionsAccessRequests) (*ConditionsAccessRequestsModel, diag.Diagnostics) {
	var out ConditionsAccessRequestsModel
	var diags diag.Diagnostics

	if val, ok := in.GetRequestTypeNameOk(); ok {
		out.RequestTypeName = types.StringPointerValue(val)
	}

	if val, ok := in.GetRequestTypeIdOk(); ok {
		out.RequestTypeId = types.StringPointerValue(val)
	}

	if val, ok := in.GetExpiresAfterSecondsOk(); ok {
		out.ExpiresAfterSeconds = types.Int32PointerValue(val)
	}
	return &out, diags
}

func ConditionsAccessRequestFromModelToSDK(_ context.Context, in *ConditionsAccessRequestsModel) (*pam.ConditionsAccessRequests, diag.Diagnostics) {
	var out pam.ConditionsAccessRequests
	var diags diag.Diagnostics

	out.Type = string(pam.SecurityPolicyRuleConditionType_ACCESS_REQUEST)

	if !in.RequestTypeId.IsUnknown() && !in.RequestTypeId.IsNull() {
		out.SetRequestTypeId(in.RequestTypeId.ValueString())
	}

	if !in.RequestTypeName.IsUnknown() && !in.RequestTypeName.IsNull() {
		out.SetRequestTypeName(in.RequestTypeName.ValueString())
	}

	if !in.ExpiresAfterSeconds.IsUnknown() && !in.ExpiresAfterSeconds.IsNull() {
		out.SetExpiresAfterSeconds(in.ExpiresAfterSeconds.ValueInt32())
	}

	return &out, diags
}

type ConditionsGatewayModel struct {
	TrafficForwarding types.Bool `tfsdk:"traffic_forwarding"`
	SessionRecording  types.Bool `tfsdk:"session_recording"`
}

func ConditionsGatewaySchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"traffic_forwarding": schema.BoolAttribute{Required: true},
			"session_recording":  schema.BoolAttribute{Required: true},
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

func ConditionsGatewayFromSDKToModel(_ context.Context, in *pam.ConditionsGateway) (*ConditionsGatewayModel, diag.Diagnostics) {
	var out ConditionsGatewayModel
	var diags diag.Diagnostics

	if val, ok := in.GetTrafficForwardingOk(); ok {
		out.TrafficForwarding = types.BoolPointerValue(val)
	}
	if val, ok := in.GetSessionRecordingOk(); ok {
		out.SessionRecording = types.BoolPointerValue(val)
	}
	return &out, diags
}

func ConditionsGatewayFromModelToSDK(_ context.Context, in *ConditionsGatewayModel) (*pam.ConditionsGateway, diag.Diagnostics) {
	var out pam.ConditionsGateway
	var diags diag.Diagnostics

	out.Type = string(pam.SecurityPolicyRuleConditionType_GATEWAY)
	if !in.TrafficForwarding.IsUnknown() && !in.TrafficForwarding.IsNull() {
		out.TrafficForwarding = in.TrafficForwarding.ValueBoolPointer()
	}

	if !in.SessionRecording.IsUnknown() && !in.SessionRecording.IsNull() {
		out.SessionRecording = in.SessionRecording.ValueBoolPointer()
	}
	return &out, diags
}

type ConditionsMFAModel struct {
	ReAuthFrequencyInSeconds types.Int32 `tfsdk:"re_auth_frequency_in_seconds"`
	AcrValues                types.String/*ConditionsMFAACRValuesModel*/ `tfsdk:"acr_values"`
}

func ConditionsMFASchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"re_auth_frequency_in_seconds": schema.Int32Attribute{Required: true},
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

func ConditionsMFAFromSDKToModel(_ context.Context, in *pam.ConditionsMFA) (*ConditionsMFAModel, diag.Diagnostics) {
	var out ConditionsMFAModel
	var diags diag.Diagnostics

	if val, ok := in.GetReAuthFrequencyInSecondsOk(); ok {
		out.ReAuthFrequencyInSeconds = types.Int32PointerValue(val)
	}

	if val, ok := in.GetAcrValuesOk(); ok {
		valStr := string(*val)
		out.AcrValues = types.StringValue(valStr)
	}

	return &out, diags
}
func ConditionsMFAFromModelToSDK(_ context.Context, in *ConditionsMFAModel) (*pam.ConditionsMFA, diag.Diagnostics) {
	var out pam.ConditionsMFA
	var diags diag.Diagnostics

	out.Type = string(pam.SecurityPolicyRuleConditionType_MFA)
	if !in.AcrValues.IsUnknown() && !in.AcrValues.IsNull() {
		val, err := pam.NewConditionsMFAACRValuesFromValue(in.AcrValues.ValueString())
		if err != nil {
			diags.AddError("could not convert Terraform ConditionsMFAModel to SDK ConditionsMFA", err.Error())
			return nil, diags
		}
		out.SetAcrValues(*val)
	}

	if !in.ReAuthFrequencyInSeconds.IsUnknown() && !in.ReAuthFrequencyInSeconds.IsNull() {
		out.SetReAuthFrequencyInSeconds(in.ReAuthFrequencyInSeconds.ValueInt32())
	}

	return &out, diags
}
