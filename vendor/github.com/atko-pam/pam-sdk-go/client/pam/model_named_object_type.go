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

// NamedObjectType the model 'NamedObjectType'
type NamedObjectType string

// List of NamedObjectType
const (
	NamedObjectType_USER                     NamedObjectType = "user"
	NamedObjectType_USER_GROUP               NamedObjectType = "user_group"
	NamedObjectType_SERVER                   NamedObjectType = "server"
	NamedObjectType_PROJECT                  NamedObjectType = "project"
	NamedObjectType_SUDO_ENTITLEMENT         NamedObjectType = "sudo_entitlement"
	NamedObjectType_RESOURCE_GROUP           NamedObjectType = "resource_group"
	NamedObjectType_SECRET_FOLDER            NamedObjectType = "secret_folder"
	NamedObjectType_SECRET                   NamedObjectType = "secret"
	NamedObjectType_ACTIVE_DIRECTORY_ACCOUNT NamedObjectType = "active_directory_account"
)

// All allowed values of NamedObjectType enum
var AllowedNamedObjectTypeEnumValues = []NamedObjectType{
	"user",
	"user_group",
	"server",
	"project",
	"sudo_entitlement",
	"resource_group",
	"secret_folder",
	"secret",
	"active_directory_account",
}

func (v *NamedObjectType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NamedObjectType(value)

	*v = enumTypeValue

	return nil
}

// NewNamedObjectTypeFromValue returns a pointer to a valid NamedObjectType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNamedObjectTypeFromValue(v string) (*NamedObjectType, error) {
	ev := NamedObjectType(v)

	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NamedObjectType) IsValid() bool {
	for _, existing := range AllowedNamedObjectTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NamedObjectType value
func (v NamedObjectType) Ptr() *NamedObjectType {
	return &v
}

type NullableNamedObjectType struct {
	value *NamedObjectType
	isSet bool
}

func (v NullableNamedObjectType) Get() *NamedObjectType {
	return v.value
}

func (v *NullableNamedObjectType) Set(val *NamedObjectType) {
	v.value = val
	v.isSet = true
}

func (v NullableNamedObjectType) IsSet() bool {
	return v.isSet
}

func (v *NullableNamedObjectType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNamedObjectType(val *NamedObjectType) *NullableNamedObjectType {
	return &NullableNamedObjectType{value: val, isSet: true}
}

func (v NullableNamedObjectType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNamedObjectType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
