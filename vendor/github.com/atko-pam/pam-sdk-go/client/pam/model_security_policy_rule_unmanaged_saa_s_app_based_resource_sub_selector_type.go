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

// SecurityPolicyRuleUnmanagedSaaSAppBasedResourceSubSelectorType The type of selector used to target unmanaged SaaS Application based resources
type SecurityPolicyRuleUnmanagedSaaSAppBasedResourceSubSelectorType string

// List of SecurityPolicyRuleUnmanagedSaaSAppBasedResourceSubSelectorType
const (
	SecurityPolicyRuleUnmanagedSaaSAppBasedResourceSubSelectorType_INDIVIDUAL_UNMANAGED_SAAS_APP_ACCOUNT SecurityPolicyRuleUnmanagedSaaSAppBasedResourceSubSelectorType = "individual_unmanaged_saas_app_account"
)

// All allowed values of SecurityPolicyRuleUnmanagedSaaSAppBasedResourceSubSelectorType enum
var AllowedSecurityPolicyRuleUnmanagedSaaSAppBasedResourceSubSelectorTypeEnumValues = []SecurityPolicyRuleUnmanagedSaaSAppBasedResourceSubSelectorType{
	"individual_unmanaged_saas_app_account",
}

func (v *SecurityPolicyRuleUnmanagedSaaSAppBasedResourceSubSelectorType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SecurityPolicyRuleUnmanagedSaaSAppBasedResourceSubSelectorType(value)

	*v = enumTypeValue

	return nil
}

// NewSecurityPolicyRuleUnmanagedSaaSAppBasedResourceSubSelectorTypeFromValue returns a pointer to a valid SecurityPolicyRuleUnmanagedSaaSAppBasedResourceSubSelectorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSecurityPolicyRuleUnmanagedSaaSAppBasedResourceSubSelectorTypeFromValue(v string) (*SecurityPolicyRuleUnmanagedSaaSAppBasedResourceSubSelectorType, error) {
	ev := SecurityPolicyRuleUnmanagedSaaSAppBasedResourceSubSelectorType(v)

	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SecurityPolicyRuleUnmanagedSaaSAppBasedResourceSubSelectorType) IsValid() bool {
	for _, existing := range AllowedSecurityPolicyRuleUnmanagedSaaSAppBasedResourceSubSelectorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityPolicyRuleUnmanagedSaaSAppBasedResourceSubSelectorType value
func (v SecurityPolicyRuleUnmanagedSaaSAppBasedResourceSubSelectorType) Ptr() *SecurityPolicyRuleUnmanagedSaaSAppBasedResourceSubSelectorType {
	return &v
}

type NullableSecurityPolicyRuleUnmanagedSaaSAppBasedResourceSubSelectorType struct {
	value *SecurityPolicyRuleUnmanagedSaaSAppBasedResourceSubSelectorType
	isSet bool
}

func (v NullableSecurityPolicyRuleUnmanagedSaaSAppBasedResourceSubSelectorType) Get() *SecurityPolicyRuleUnmanagedSaaSAppBasedResourceSubSelectorType {
	return v.value
}

func (v *NullableSecurityPolicyRuleUnmanagedSaaSAppBasedResourceSubSelectorType) Set(val *SecurityPolicyRuleUnmanagedSaaSAppBasedResourceSubSelectorType) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityPolicyRuleUnmanagedSaaSAppBasedResourceSubSelectorType) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityPolicyRuleUnmanagedSaaSAppBasedResourceSubSelectorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityPolicyRuleUnmanagedSaaSAppBasedResourceSubSelectorType(val *SecurityPolicyRuleUnmanagedSaaSAppBasedResourceSubSelectorType) *NullableSecurityPolicyRuleUnmanagedSaaSAppBasedResourceSubSelectorType {
	return &NullableSecurityPolicyRuleUnmanagedSaaSAppBasedResourceSubSelectorType{value: val, isSet: true}
}

func (v NullableSecurityPolicyRuleUnmanagedSaaSAppBasedResourceSubSelectorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityPolicyRuleUnmanagedSaaSAppBasedResourceSubSelectorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}