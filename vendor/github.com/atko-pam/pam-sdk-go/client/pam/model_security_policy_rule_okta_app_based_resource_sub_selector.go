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

// SecurityPolicyRuleOktaAppBasedResourceSubSelector - The specific parameters used to target resources. The organization of this object depends on the `selector_type`.
type SecurityPolicyRuleOktaAppBasedResourceSubSelector struct {
	SelectorIndividualOktaServiceAccount *SelectorIndividualOktaServiceAccount
}

// SelectorIndividualOktaServiceAccountAsSecurityPolicyRuleOktaAppBasedResourceSubSelector is a convenience function that returns SelectorIndividualOktaServiceAccount wrapped in SecurityPolicyRuleOktaAppBasedResourceSubSelector
func SelectorIndividualOktaServiceAccountAsSecurityPolicyRuleOktaAppBasedResourceSubSelector(v *SelectorIndividualOktaServiceAccount) SecurityPolicyRuleOktaAppBasedResourceSubSelector {
	return SecurityPolicyRuleOktaAppBasedResourceSubSelector{
		SelectorIndividualOktaServiceAccount: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SecurityPolicyRuleOktaAppBasedResourceSubSelector) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'SelectorIndividualOktaServiceAccount'
	if jsonDict["_type"] == "SelectorIndividualOktaServiceAccount" {
		// try to unmarshal JSON data into SelectorIndividualOktaServiceAccount
		err = json.Unmarshal(data, &dst.SelectorIndividualOktaServiceAccount)
		if err == nil {
			return nil // data stored in dst.SelectorIndividualOktaServiceAccount, return on the first match
		} else {
			dst.SelectorIndividualOktaServiceAccount = nil
			return fmt.Errorf("failed to unmarshal SecurityPolicyRuleOktaAppBasedResourceSubSelector as SelectorIndividualOktaServiceAccount: %s", err.Error())
		}
	}

	// check if the discriminator value is 'individual_okta_account'
	if jsonDict["_type"] == "individual_okta_account" {
		// try to unmarshal JSON data into SelectorIndividualOktaServiceAccount
		err = json.Unmarshal(data, &dst.SelectorIndividualOktaServiceAccount)
		if err == nil {
			return nil // data stored in dst.SelectorIndividualOktaServiceAccount, return on the first match
		} else {
			dst.SelectorIndividualOktaServiceAccount = nil
			return fmt.Errorf("failed to unmarshal SecurityPolicyRuleOktaAppBasedResourceSubSelector as SelectorIndividualOktaServiceAccount: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SecurityPolicyRuleOktaAppBasedResourceSubSelector) MarshalJSON() ([]byte, error) {
	if src.SelectorIndividualOktaServiceAccount != nil {
		return json.Marshal(&src.SelectorIndividualOktaServiceAccount)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SecurityPolicyRuleOktaAppBasedResourceSubSelector) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.SelectorIndividualOktaServiceAccount != nil {
		return obj.SelectorIndividualOktaServiceAccount
	}

	// all schemas are nil
	return nil
}

type NullableSecurityPolicyRuleOktaAppBasedResourceSubSelector struct {
	value *SecurityPolicyRuleOktaAppBasedResourceSubSelector
	isSet bool
}

func (v NullableSecurityPolicyRuleOktaAppBasedResourceSubSelector) Get() *SecurityPolicyRuleOktaAppBasedResourceSubSelector {
	return v.value
}

func (v *NullableSecurityPolicyRuleOktaAppBasedResourceSubSelector) Set(val *SecurityPolicyRuleOktaAppBasedResourceSubSelector) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityPolicyRuleOktaAppBasedResourceSubSelector) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityPolicyRuleOktaAppBasedResourceSubSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityPolicyRuleOktaAppBasedResourceSubSelector(val *SecurityPolicyRuleOktaAppBasedResourceSubSelector) *NullableSecurityPolicyRuleOktaAppBasedResourceSubSelector {
	return &NullableSecurityPolicyRuleOktaAppBasedResourceSubSelector{value: val, isSet: true}
}

func (v NullableSecurityPolicyRuleOktaAppBasedResourceSubSelector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityPolicyRuleOktaAppBasedResourceSubSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
