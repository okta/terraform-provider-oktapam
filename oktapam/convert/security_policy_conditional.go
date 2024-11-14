package convert

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SecurityPolicyRuleConditionModel struct {
	ConditionsAccessRequests *ConditionsAccessRequestsModel `tfsdk:"access_request"`
	ConditionsGateway        *ConditionsGatewayModel        `tfsdk:"gateway"`
	ConditionsMFA            *ConditionsMFAModel            `tfsdk:"mfa"`
}

func SecurityPolicyRuleConditionsSchema() schema.NestedAttributeObject {
	return schema.NestedAttributeObject{
		Attributes: map[string]schema.Attribute{
			"access_request": ConditionsAccessRequestsSchema(),
			"gateway":        ConditionsGatewaySchema(),
			"mfa":            ConditionsMFASchema(),
		},
	}
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
