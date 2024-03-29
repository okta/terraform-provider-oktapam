/*
Okta Privileged Access

The OPA API is a control plane used to request operations in Okta Privileged Access (formerly ScaleFT/Advanced Server Access)

API version: 1.0.0
Contact: support@okta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pam

import (
	"encoding/json"
	"fmt"
)

// SecurityPolicyRuleCondition - The parameters for the condition. The structure of this object depends on the `condition_type`.
type SecurityPolicyRuleCondition struct {
	ConditionsAccessRequests *ConditionsAccessRequests
	ConditionsGateway        *ConditionsGateway
	ConditionsMFA            *ConditionsMFA
}

// ConditionsAccessRequestsAsSecurityPolicyRuleCondition is a convenience function that returns ConditionsAccessRequests wrapped in SecurityPolicyRuleCondition
func ConditionsAccessRequestsAsSecurityPolicyRuleCondition(v *ConditionsAccessRequests) SecurityPolicyRuleCondition {
	return SecurityPolicyRuleCondition{
		ConditionsAccessRequests: v,
	}
}

// ConditionsGatewayAsSecurityPolicyRuleCondition is a convenience function that returns ConditionsGateway wrapped in SecurityPolicyRuleCondition
func ConditionsGatewayAsSecurityPolicyRuleCondition(v *ConditionsGateway) SecurityPolicyRuleCondition {
	return SecurityPolicyRuleCondition{
		ConditionsGateway: v,
	}
}

// ConditionsMFAAsSecurityPolicyRuleCondition is a convenience function that returns ConditionsMFA wrapped in SecurityPolicyRuleCondition
func ConditionsMFAAsSecurityPolicyRuleCondition(v *ConditionsMFA) SecurityPolicyRuleCondition {
	return SecurityPolicyRuleCondition{
		ConditionsMFA: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SecurityPolicyRuleCondition) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ConditionsAccessRequests
	err = json.Unmarshal(data, &dst.ConditionsAccessRequests)
	if err == nil {
		jsonConditionsAccessRequests, _ := json.Marshal(dst.ConditionsAccessRequests)
		if string(jsonConditionsAccessRequests) == "{}" { // empty struct
			dst.ConditionsAccessRequests = nil
		} else {
			match++
		}
	} else {
		dst.ConditionsAccessRequests = nil
	}

	// try to unmarshal data into ConditionsGateway
	err = json.Unmarshal(data, &dst.ConditionsGateway)
	if err == nil {
		jsonConditionsGateway, _ := json.Marshal(dst.ConditionsGateway)
		if string(jsonConditionsGateway) == "{}" { // empty struct
			dst.ConditionsGateway = nil
		} else {
			match++
		}
	} else {
		dst.ConditionsGateway = nil
	}

	// try to unmarshal data into ConditionsMFA
	err = json.Unmarshal(data, &dst.ConditionsMFA)
	if err == nil {
		jsonConditionsMFA, _ := json.Marshal(dst.ConditionsMFA)
		if string(jsonConditionsMFA) == "{}" { // empty struct
			dst.ConditionsMFA = nil
		} else {
			match++
		}
	} else {
		dst.ConditionsMFA = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ConditionsAccessRequests = nil
		dst.ConditionsGateway = nil
		dst.ConditionsMFA = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SecurityPolicyRuleCondition)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SecurityPolicyRuleCondition)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SecurityPolicyRuleCondition) MarshalJSON() ([]byte, error) {
	if src.ConditionsAccessRequests != nil {
		return json.Marshal(&src.ConditionsAccessRequests)
	}

	if src.ConditionsGateway != nil {
		return json.Marshal(&src.ConditionsGateway)
	}

	if src.ConditionsMFA != nil {
		return json.Marshal(&src.ConditionsMFA)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SecurityPolicyRuleCondition) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ConditionsAccessRequests != nil {
		return obj.ConditionsAccessRequests
	}

	if obj.ConditionsGateway != nil {
		return obj.ConditionsGateway
	}

	if obj.ConditionsMFA != nil {
		return obj.ConditionsMFA
	}

	// all schemas are nil
	return nil
}

type NullableSecurityPolicyRuleCondition struct {
	value *SecurityPolicyRuleCondition
	isSet bool
}

func (v NullableSecurityPolicyRuleCondition) Get() *SecurityPolicyRuleCondition {
	return v.value
}

func (v *NullableSecurityPolicyRuleCondition) Set(val *SecurityPolicyRuleCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityPolicyRuleCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityPolicyRuleCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityPolicyRuleCondition(val *SecurityPolicyRuleCondition) *NullableSecurityPolicyRuleCondition {
	return &NullableSecurityPolicyRuleCondition{value: val, isSet: true}
}

func (v NullableSecurityPolicyRuleCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityPolicyRuleCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
