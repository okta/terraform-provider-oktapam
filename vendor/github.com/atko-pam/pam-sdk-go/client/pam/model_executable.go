/*
Okta Privileged Access

The ScaleFT API is a control plane API for operations in Okta Privileged Access (formerly ScaleFT)

API version: 1.0.0
Contact: support@okta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pam

import (
	"encoding/json"
	"fmt"
)

// Executable the model 'Executable'
type Executable string

// List of executable
const (
	Executable_ANY    Executable = "any"
	Executable_CUSTOM Executable = "custom"
	Executable_NONE   Executable = "none"
)

// All allowed values of Executable enum
var AllowedExecutableEnumValues = []Executable{
	"any",
	"custom",
	"none",
}

func (v *Executable) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Executable(value)
	for _, existing := range AllowedExecutableEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Executable", value)
}

// NewExecutableFromValue returns a pointer to a valid Executable
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewExecutableFromValue(v string) (*Executable, error) {
	ev := Executable(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Executable: valid values are %v", v, AllowedExecutableEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Executable) IsValid() bool {
	for _, existing := range AllowedExecutableEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to executable value
func (v Executable) Ptr() *Executable {
	return &v
}

type NullableExecutable struct {
	value *Executable
	isSet bool
}

func (v NullableExecutable) Get() *Executable {
	return v.value
}

func (v *NullableExecutable) Set(val *Executable) {
	v.value = val
	v.isSet = true
}

func (v NullableExecutable) IsSet() bool {
	return v.isSet
}

func (v *NullableExecutable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecutable(val *Executable) *NullableExecutable {
	return &NullableExecutable{value: val, isSet: true}
}

func (v NullableExecutable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecutable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}