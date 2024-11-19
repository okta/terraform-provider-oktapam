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
)

// SecurityPolicyRuleOktaAppBasedResourceSubSelectorType The type of selector used to target Okta Application based resources
type SecurityPolicyRuleOktaAppBasedResourceSubSelectorType string

// List of SecurityPolicyRuleOktaAppBasedResourceSubSelectorType
const (
	SecurityPolicyRuleOktaAppBasedResourceSubSelectorType_INDIVIDUAL_OKTA_ACCOUNT SecurityPolicyRuleOktaAppBasedResourceSubSelectorType = "individual_okta_account"
)

// All allowed values of SecurityPolicyRuleOktaAppBasedResourceSubSelectorType enum
var AllowedSecurityPolicyRuleOktaAppBasedResourceSubSelectorTypeEnumValues = []SecurityPolicyRuleOktaAppBasedResourceSubSelectorType{
	"individual_okta_account",
}

func (v *SecurityPolicyRuleOktaAppBasedResourceSubSelectorType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SecurityPolicyRuleOktaAppBasedResourceSubSelectorType(value)

	*v = enumTypeValue

	return nil
}

// NewSecurityPolicyRuleOktaAppBasedResourceSubSelectorTypeFromValue returns a pointer to a valid SecurityPolicyRuleOktaAppBasedResourceSubSelectorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSecurityPolicyRuleOktaAppBasedResourceSubSelectorTypeFromValue(v string) (*SecurityPolicyRuleOktaAppBasedResourceSubSelectorType, error) {
	ev := SecurityPolicyRuleOktaAppBasedResourceSubSelectorType(v)

	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SecurityPolicyRuleOktaAppBasedResourceSubSelectorType) IsValid() bool {
	for _, existing := range AllowedSecurityPolicyRuleOktaAppBasedResourceSubSelectorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityPolicyRuleOktaAppBasedResourceSubSelectorType value
func (v SecurityPolicyRuleOktaAppBasedResourceSubSelectorType) Ptr() *SecurityPolicyRuleOktaAppBasedResourceSubSelectorType {
	return &v
}

type NullableSecurityPolicyRuleOktaAppBasedResourceSubSelectorType struct {
	value *SecurityPolicyRuleOktaAppBasedResourceSubSelectorType
	isSet bool
}

func (v NullableSecurityPolicyRuleOktaAppBasedResourceSubSelectorType) Get() *SecurityPolicyRuleOktaAppBasedResourceSubSelectorType {
	return v.value
}

func (v *NullableSecurityPolicyRuleOktaAppBasedResourceSubSelectorType) Set(val *SecurityPolicyRuleOktaAppBasedResourceSubSelectorType) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityPolicyRuleOktaAppBasedResourceSubSelectorType) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityPolicyRuleOktaAppBasedResourceSubSelectorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityPolicyRuleOktaAppBasedResourceSubSelectorType(val *SecurityPolicyRuleOktaAppBasedResourceSubSelectorType) *NullableSecurityPolicyRuleOktaAppBasedResourceSubSelectorType {
	return &NullableSecurityPolicyRuleOktaAppBasedResourceSubSelectorType{value: val, isSet: true}
}

func (v NullableSecurityPolicyRuleOktaAppBasedResourceSubSelectorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityPolicyRuleOktaAppBasedResourceSubSelectorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}