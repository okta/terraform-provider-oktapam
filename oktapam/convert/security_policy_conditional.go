package convert

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ConditionsMFAACRValuesModel types.String

type SecurityPolicyRuleConditionContainerModel struct {
	SecurityPolicyRuleConditionModel
}

type SecurityPolicyRuleConditionModel struct {
	ConditionsAccessRequests *ConditionsAccessRequestsModel `tfsdk:"access_request"`
	ConditionsGateway        *ConditionsGatewayModel        `tfsdk:"gateway"`
	ConditionsMFA            *ConditionsMFAModel            `tfsdk:"mfa"`
}

func SecurityPolicyRuleConditionsBlock() schema.Block {
	return schema.SingleNestedBlock{
		Blocks: map[string]schema.Block{
			"access_request": ConditionsAccessRequestsBlock(),
			"gateway":        ConditionsGatewayBlock(),
			"mfa":            ConditionsMFABlock(),
		},
	}
}

type ConditionsAccessRequestsModel struct {
	RequestTypeId       types.String `tfsdk:"request_type_id"`
	RequestTypeName     types.String `tfsdk:"request_type_name"`
	ExpiresAfterSeconds types.Int32  `tfsdk:"expires_after_seconds"`
}

func ConditionsAccessRequestsBlock() schema.Block {
	return schema.SingleNestedBlock{
		Attributes: map[string]schema.Attribute{
			"request_type_id":       schema.StringAttribute{ /*TODO(ja)*/ },
			"request_type_name":     schema.StringAttribute{Required: true},
			"expires_after_seconds": schema.Int32Attribute{Optional: true},
		},
	}
}

type ConditionsGatewayModel struct {
	TrafficForwarding types.Bool `tfsdk:"traffic_forwarding"`
	SessionRecording  types.Bool `tfsdk:"session_recording"`
}

func ConditionsGatewayBlock() schema.Block {
	return schema.SingleNestedBlock{
		Attributes: map[string]schema.Attribute{
			"traffic_forwarding": schema.BoolAttribute{Optional: true},
			"session_recording":  schema.BoolAttribute{Optional: true},
		},
	}
}

type ConditionsMFAModel struct {
	ReAuthFrequencyInSeconds types.Int32                 `tfsdk:"re_auth_frequency_in_seconds"`
	AcrValues                ConditionsMFAACRValuesModel `tfsdk:"acr_values"`
}

func ConditionsMFABlock() schema.Block {
	return schema.SingleNestedBlock{
		Attributes: map[string]schema.Attribute{
			"re_auth_frequency_in_seconds": schema.Int32Attribute{},
			"acr_values":                   schema.StringAttribute{},
		},
	}
}
