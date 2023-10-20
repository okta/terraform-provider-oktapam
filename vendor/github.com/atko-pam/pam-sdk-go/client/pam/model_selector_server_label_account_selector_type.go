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

// SelectorServerLabelAccountSelectorType Defines the type of account. Currently only accepts `username`.
type SelectorServerLabelAccountSelectorType string

// List of SelectorServerLabelAccountSelectorType
const (
	SelectorServerLabelAccountSelectorType_USERNAME SelectorServerLabelAccountSelectorType = "username"
	SelectorServerLabelAccountSelectorType_NONE     SelectorServerLabelAccountSelectorType = "none"
)

// All allowed values of SelectorServerLabelAccountSelectorType enum
var AllowedSelectorServerLabelAccountSelectorTypeEnumValues = []SelectorServerLabelAccountSelectorType{
	"username",
	"none",
}

func (v *SelectorServerLabelAccountSelectorType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SelectorServerLabelAccountSelectorType(value)
	for _, existing := range AllowedSelectorServerLabelAccountSelectorTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SelectorServerLabelAccountSelectorType", value)
}

// NewSelectorServerLabelAccountSelectorTypeFromValue returns a pointer to a valid SelectorServerLabelAccountSelectorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSelectorServerLabelAccountSelectorTypeFromValue(v string) (*SelectorServerLabelAccountSelectorType, error) {
	ev := SelectorServerLabelAccountSelectorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SelectorServerLabelAccountSelectorType: valid values are %v", v, AllowedSelectorServerLabelAccountSelectorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SelectorServerLabelAccountSelectorType) IsValid() bool {
	for _, existing := range AllowedSelectorServerLabelAccountSelectorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SelectorServerLabelAccountSelectorType value
func (v SelectorServerLabelAccountSelectorType) Ptr() *SelectorServerLabelAccountSelectorType {
	return &v
}

type NullableSelectorServerLabelAccountSelectorType struct {
	value *SelectorServerLabelAccountSelectorType
	isSet bool
}

func (v NullableSelectorServerLabelAccountSelectorType) Get() *SelectorServerLabelAccountSelectorType {
	return v.value
}

func (v *NullableSelectorServerLabelAccountSelectorType) Set(val *SelectorServerLabelAccountSelectorType) {
	v.value = val
	v.isSet = true
}

func (v NullableSelectorServerLabelAccountSelectorType) IsSet() bool {
	return v.isSet
}

func (v *NullableSelectorServerLabelAccountSelectorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelectorServerLabelAccountSelectorType(val *SelectorServerLabelAccountSelectorType) *NullableSelectorServerLabelAccountSelectorType {
	return &NullableSelectorServerLabelAccountSelectorType{value: val, isSet: true}
}

func (v NullableSelectorServerLabelAccountSelectorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelectorServerLabelAccountSelectorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}