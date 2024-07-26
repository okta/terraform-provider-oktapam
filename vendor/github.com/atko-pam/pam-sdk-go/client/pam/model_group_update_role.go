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

// GroupUpdateRole the model 'GroupUpdateRole'
type GroupUpdateRole string

// List of GroupUpdateRole
const (
	GroupUpdateRole_PAM_ADMIN      GroupUpdateRole = "pam_admin"
	GroupUpdateRole_RESOURCE_ADMIN GroupUpdateRole = "resource_admin"
	GroupUpdateRole_SECURITY_ADMIN GroupUpdateRole = "security_admin"
	GroupUpdateRole_END_USER       GroupUpdateRole = "end_user"
)

// All allowed values of GroupUpdateRole enum
var AllowedGroupUpdateRoleEnumValues = []GroupUpdateRole{
	"pam_admin",
	"resource_admin",
	"security_admin",
	"end_user",
}

func (v *GroupUpdateRole) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GroupUpdateRole(value)

	*v = enumTypeValue

	return nil
}

// NewGroupUpdateRoleFromValue returns a pointer to a valid GroupUpdateRole
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGroupUpdateRoleFromValue(v string) (*GroupUpdateRole, error) {
	ev := GroupUpdateRole(v)

	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GroupUpdateRole) IsValid() bool {
	for _, existing := range AllowedGroupUpdateRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GroupUpdateRole value
func (v GroupUpdateRole) Ptr() *GroupUpdateRole {
	return &v
}

type NullableGroupUpdateRole struct {
	value *GroupUpdateRole
	isSet bool
}

func (v NullableGroupUpdateRole) Get() *GroupUpdateRole {
	return v.value
}

func (v *NullableGroupUpdateRole) Set(val *GroupUpdateRole) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupUpdateRole) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupUpdateRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupUpdateRole(val *GroupUpdateRole) *NullableGroupUpdateRole {
	return &NullableGroupUpdateRole{value: val, isSet: true}
}

func (v NullableGroupUpdateRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupUpdateRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
