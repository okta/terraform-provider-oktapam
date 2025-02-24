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

// SecurityPolicyRuleActiveDirectoryBasedResourceSubSelectorType The type of selector used to target Active Directory resources
type SecurityPolicyRuleActiveDirectoryBasedResourceSubSelectorType string

// List of SecurityPolicyRuleActiveDirectoryBasedResourceSubSelectorType
const (
	SecurityPolicyRuleActiveDirectoryBasedResourceSubSelectorType_ACTIVE_DIRECTORY SecurityPolicyRuleActiveDirectoryBasedResourceSubSelectorType = "active_directory"
)

// All allowed values of SecurityPolicyRuleActiveDirectoryBasedResourceSubSelectorType enum
var AllowedSecurityPolicyRuleActiveDirectoryBasedResourceSubSelectorTypeEnumValues = []SecurityPolicyRuleActiveDirectoryBasedResourceSubSelectorType{
	"active_directory",
}

func (v *SecurityPolicyRuleActiveDirectoryBasedResourceSubSelectorType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SecurityPolicyRuleActiveDirectoryBasedResourceSubSelectorType(value)

	*v = enumTypeValue

	return nil
}

// NewSecurityPolicyRuleActiveDirectoryBasedResourceSubSelectorTypeFromValue returns a pointer to a valid SecurityPolicyRuleActiveDirectoryBasedResourceSubSelectorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSecurityPolicyRuleActiveDirectoryBasedResourceSubSelectorTypeFromValue(v string) (*SecurityPolicyRuleActiveDirectoryBasedResourceSubSelectorType, error) {
	ev := SecurityPolicyRuleActiveDirectoryBasedResourceSubSelectorType(v)

	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SecurityPolicyRuleActiveDirectoryBasedResourceSubSelectorType) IsValid() bool {
	for _, existing := range AllowedSecurityPolicyRuleActiveDirectoryBasedResourceSubSelectorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityPolicyRuleActiveDirectoryBasedResourceSubSelectorType value
func (v SecurityPolicyRuleActiveDirectoryBasedResourceSubSelectorType) Ptr() *SecurityPolicyRuleActiveDirectoryBasedResourceSubSelectorType {
	return &v
}

type NullableSecurityPolicyRuleActiveDirectoryBasedResourceSubSelectorType struct {
	value *SecurityPolicyRuleActiveDirectoryBasedResourceSubSelectorType
	isSet bool
}

func (v NullableSecurityPolicyRuleActiveDirectoryBasedResourceSubSelectorType) Get() *SecurityPolicyRuleActiveDirectoryBasedResourceSubSelectorType {
	return v.value
}

func (v *NullableSecurityPolicyRuleActiveDirectoryBasedResourceSubSelectorType) Set(val *SecurityPolicyRuleActiveDirectoryBasedResourceSubSelectorType) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityPolicyRuleActiveDirectoryBasedResourceSubSelectorType) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityPolicyRuleActiveDirectoryBasedResourceSubSelectorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityPolicyRuleActiveDirectoryBasedResourceSubSelectorType(val *SecurityPolicyRuleActiveDirectoryBasedResourceSubSelectorType) *NullableSecurityPolicyRuleActiveDirectoryBasedResourceSubSelectorType {
	return &NullableSecurityPolicyRuleActiveDirectoryBasedResourceSubSelectorType{value: val, isSet: true}
}

func (v NullableSecurityPolicyRuleActiveDirectoryBasedResourceSubSelectorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityPolicyRuleActiveDirectoryBasedResourceSubSelectorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
