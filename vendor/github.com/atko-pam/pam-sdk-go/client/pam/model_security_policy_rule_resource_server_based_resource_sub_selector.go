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

// SecurityPolicyRuleResourceServerBasedResourceSubSelector - The specific parameters used to target resources. The organization of this object depends on the `selector_type`.
type SecurityPolicyRuleResourceServerBasedResourceSubSelector struct {
	SelectorIndividualServer        *SelectorIndividualServer
	SelectorIndividualServerAccount *SelectorIndividualServerAccount
	SelectorServerLabel             *SelectorServerLabel
}

// SelectorIndividualServerAsSecurityPolicyRuleResourceServerBasedResourceSubSelector is a convenience function that returns SelectorIndividualServer wrapped in SecurityPolicyRuleResourceServerBasedResourceSubSelector
func SelectorIndividualServerAsSecurityPolicyRuleResourceServerBasedResourceSubSelector(v *SelectorIndividualServer) SecurityPolicyRuleResourceServerBasedResourceSubSelector {
	return SecurityPolicyRuleResourceServerBasedResourceSubSelector{
		SelectorIndividualServer: v,
	}
}

// SelectorIndividualServerAccountAsSecurityPolicyRuleResourceServerBasedResourceSubSelector is a convenience function that returns SelectorIndividualServerAccount wrapped in SecurityPolicyRuleResourceServerBasedResourceSubSelector
func SelectorIndividualServerAccountAsSecurityPolicyRuleResourceServerBasedResourceSubSelector(v *SelectorIndividualServerAccount) SecurityPolicyRuleResourceServerBasedResourceSubSelector {
	return SecurityPolicyRuleResourceServerBasedResourceSubSelector{
		SelectorIndividualServerAccount: v,
	}
}

// SelectorServerLabelAsSecurityPolicyRuleResourceServerBasedResourceSubSelector is a convenience function that returns SelectorServerLabel wrapped in SecurityPolicyRuleResourceServerBasedResourceSubSelector
func SelectorServerLabelAsSecurityPolicyRuleResourceServerBasedResourceSubSelector(v *SelectorServerLabel) SecurityPolicyRuleResourceServerBasedResourceSubSelector {
	return SecurityPolicyRuleResourceServerBasedResourceSubSelector{
		SelectorServerLabel: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SecurityPolicyRuleResourceServerBasedResourceSubSelector) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'SelectorIndividualServer'
	if jsonDict["_type"] == "SelectorIndividualServer" {
		// try to unmarshal JSON data into SelectorIndividualServer
		err = json.Unmarshal(data, &dst.SelectorIndividualServer)
		if err == nil {
			return nil // data stored in dst.SelectorIndividualServer, return on the first match
		} else {
			dst.SelectorIndividualServer = nil
			return fmt.Errorf("failed to unmarshal SecurityPolicyRuleResourceServerBasedResourceSubSelector as SelectorIndividualServer: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SelectorIndividualServerAccount'
	if jsonDict["_type"] == "SelectorIndividualServerAccount" {
		// try to unmarshal JSON data into SelectorIndividualServerAccount
		err = json.Unmarshal(data, &dst.SelectorIndividualServerAccount)
		if err == nil {
			return nil // data stored in dst.SelectorIndividualServerAccount, return on the first match
		} else {
			dst.SelectorIndividualServerAccount = nil
			return fmt.Errorf("failed to unmarshal SecurityPolicyRuleResourceServerBasedResourceSubSelector as SelectorIndividualServerAccount: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SelectorServerLabel'
	if jsonDict["_type"] == "SelectorServerLabel" {
		// try to unmarshal JSON data into SelectorServerLabel
		err = json.Unmarshal(data, &dst.SelectorServerLabel)
		if err == nil {
			return nil // data stored in dst.SelectorServerLabel, return on the first match
		} else {
			dst.SelectorServerLabel = nil
			return fmt.Errorf("failed to unmarshal SecurityPolicyRuleResourceServerBasedResourceSubSelector as SelectorServerLabel: %s", err.Error())
		}
	}

	// check if the discriminator value is 'individual_server'
	if jsonDict["_type"] == "individual_server" {
		// try to unmarshal JSON data into SelectorIndividualServer
		err = json.Unmarshal(data, &dst.SelectorIndividualServer)
		if err == nil {
			return nil // data stored in dst.SelectorIndividualServer, return on the first match
		} else {
			dst.SelectorIndividualServer = nil
			return fmt.Errorf("failed to unmarshal SecurityPolicyRuleResourceServerBasedResourceSubSelector as SelectorIndividualServer: %s", err.Error())
		}
	}

	// check if the discriminator value is 'individual_server_account'
	if jsonDict["_type"] == "individual_server_account" {
		// try to unmarshal JSON data into SelectorIndividualServerAccount
		err = json.Unmarshal(data, &dst.SelectorIndividualServerAccount)
		if err == nil {
			return nil // data stored in dst.SelectorIndividualServerAccount, return on the first match
		} else {
			dst.SelectorIndividualServerAccount = nil
			return fmt.Errorf("failed to unmarshal SecurityPolicyRuleResourceServerBasedResourceSubSelector as SelectorIndividualServerAccount: %s", err.Error())
		}
	}

	// check if the discriminator value is 'server_label'
	if jsonDict["_type"] == "server_label" {
		// try to unmarshal JSON data into SelectorServerLabel
		err = json.Unmarshal(data, &dst.SelectorServerLabel)
		if err == nil {
			return nil // data stored in dst.SelectorServerLabel, return on the first match
		} else {
			dst.SelectorServerLabel = nil
			return fmt.Errorf("failed to unmarshal SecurityPolicyRuleResourceServerBasedResourceSubSelector as SelectorServerLabel: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SecurityPolicyRuleResourceServerBasedResourceSubSelector) MarshalJSON() ([]byte, error) {
	if src.SelectorIndividualServer != nil {
		return json.Marshal(&src.SelectorIndividualServer)
	}

	if src.SelectorIndividualServerAccount != nil {
		return json.Marshal(&src.SelectorIndividualServerAccount)
	}

	if src.SelectorServerLabel != nil {
		return json.Marshal(&src.SelectorServerLabel)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SecurityPolicyRuleResourceServerBasedResourceSubSelector) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.SelectorIndividualServer != nil {
		return obj.SelectorIndividualServer
	}

	if obj.SelectorIndividualServerAccount != nil {
		return obj.SelectorIndividualServerAccount
	}

	if obj.SelectorServerLabel != nil {
		return obj.SelectorServerLabel
	}

	// all schemas are nil
	return nil
}

type NullableSecurityPolicyRuleResourceServerBasedResourceSubSelector struct {
	value *SecurityPolicyRuleResourceServerBasedResourceSubSelector
	isSet bool
}

func (v NullableSecurityPolicyRuleResourceServerBasedResourceSubSelector) Get() *SecurityPolicyRuleResourceServerBasedResourceSubSelector {
	return v.value
}

func (v *NullableSecurityPolicyRuleResourceServerBasedResourceSubSelector) Set(val *SecurityPolicyRuleResourceServerBasedResourceSubSelector) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityPolicyRuleResourceServerBasedResourceSubSelector) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityPolicyRuleResourceServerBasedResourceSubSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityPolicyRuleResourceServerBasedResourceSubSelector(val *SecurityPolicyRuleResourceServerBasedResourceSubSelector) *NullableSecurityPolicyRuleResourceServerBasedResourceSubSelector {
	return &NullableSecurityPolicyRuleResourceServerBasedResourceSubSelector{value: val, isSet: true}
}

func (v NullableSecurityPolicyRuleResourceServerBasedResourceSubSelector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityPolicyRuleResourceServerBasedResourceSubSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}