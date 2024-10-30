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

// ArgsType The `args_type` is only allowed for the 'executable' command type
type ArgsType string

// List of args_type
const (
	ArgsType_ANY    ArgsType = "any"
	ArgsType_CUSTOM ArgsType = "custom"
	ArgsType_NONE   ArgsType = "none"
)

// All allowed values of ArgsType enum
var AllowedArgsTypeEnumValues = []ArgsType{
	"any",
	"custom",
	"none",
}

func (v *ArgsType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ArgsType(value)

	*v = enumTypeValue

	return nil
}

// NewArgsTypeFromValue returns a pointer to a valid ArgsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewArgsTypeFromValue(v string) (*ArgsType, error) {
	ev := ArgsType(v)

	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ArgsType) IsValid() bool {
	for _, existing := range AllowedArgsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to args_type value
func (v ArgsType) Ptr() *ArgsType {
	return &v
}

type NullableArgsType struct {
	value *ArgsType
	isSet bool
}

func (v NullableArgsType) Get() *ArgsType {
	return v.value
}

func (v *NullableArgsType) Set(val *ArgsType) {
	v.value = val
	v.isSet = true
}

func (v NullableArgsType) IsSet() bool {
	return v.isSet
}

func (v *NullableArgsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArgsType(val *ArgsType) *NullableArgsType {
	return &NullableArgsType{value: val, isSet: true}
}

func (v NullableArgsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArgsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
