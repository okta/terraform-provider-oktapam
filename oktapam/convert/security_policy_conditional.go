package convert

import (
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

type ConditionsAccessRequestsModel struct {
	RequestTypeId       types.String `tfsdk:"request_type_id"`
	RequestTypeName     types.String `tfsdk:"request_type_name"`
	ExpiresAfterSeconds types.Int32  `tfsdk:"expires_after_seconds"`
}

type ConditionsGatewayModel struct {
	TrafficForwarding types.Bool `tfsdk:"traffic_forwarding"`
	SessionRecording  types.Bool `tfsdk:"session_recording"`
}

type ConditionsMFAModel struct {
	ReAuthFrequencyInSeconds types.Int32                 `tfsdk:"re_auth_frequency_in_seconds"`
	AcrValues                ConditionsMFAACRValuesModel `tfsdk:"acr_values"`
}
