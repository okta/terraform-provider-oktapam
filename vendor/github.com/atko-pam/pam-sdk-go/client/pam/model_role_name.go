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

// RoleName The name of the available roles
type RoleName string

// List of RoleName
const (
	RoleName_DELEGATED_RESOURCE_ADMIN RoleName = "delegated_resource_admin"
	RoleName_END_USER                 RoleName = "end_user"
	RoleName_PAM_ADMIN                RoleName = "pam_admin"
	RoleName_RESOURCE_ADMIN           RoleName = "resource_admin"
	RoleName_SECURITY_ADMIN           RoleName = "security_admin"
)

// All allowed values of RoleName enum
var AllowedRoleNameEnumValues = []RoleName{
	"delegated_resource_admin",
	"end_user",
	"pam_admin",
	"resource_admin",
	"security_admin",
}

func (v *RoleName) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RoleName(value)

	*v = enumTypeValue

	return nil
}

// NewRoleNameFromValue returns a pointer to a valid RoleName
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRoleNameFromValue(v string) (*RoleName, error) {
	ev := RoleName(v)

	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RoleName) IsValid() bool {
	for _, existing := range AllowedRoleNameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RoleName value
func (v RoleName) Ptr() *RoleName {
	return &v
}

type NullableRoleName struct {
	value *RoleName
	isSet bool
}

func (v NullableRoleName) Get() *RoleName {
	return v.value
}

func (v *NullableRoleName) Set(val *RoleName) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleName) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleName(val *RoleName) *NullableRoleName {
	return &NullableRoleName{value: val, isSet: true}
}

func (v NullableRoleName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
