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

// ResolveResourceType The type of resource to resolve
type ResolveResourceType string

// List of ResolveResourceType
const (
	ResolveResourceType_SERVER   ResolveResourceType = "server"
	ResolveResourceType_DATABASE ResolveResourceType = "database"
)

// All allowed values of ResolveResourceType enum
var AllowedResolveResourceTypeEnumValues = []ResolveResourceType{
	"server",
	"database",
}

func (v *ResolveResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ResolveResourceType(value)
	for _, existing := range AllowedResolveResourceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ResolveResourceType", value)
}

// NewResolveResourceTypeFromValue returns a pointer to a valid ResolveResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewResolveResourceTypeFromValue(v string) (*ResolveResourceType, error) {
	ev := ResolveResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ResolveResourceType: valid values are %v", v, AllowedResolveResourceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ResolveResourceType) IsValid() bool {
	for _, existing := range AllowedResolveResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ResolveResourceType value
func (v ResolveResourceType) Ptr() *ResolveResourceType {
	return &v
}

type NullableResolveResourceType struct {
	value *ResolveResourceType
	isSet bool
}

func (v NullableResolveResourceType) Get() *ResolveResourceType {
	return v.value
}

func (v *NullableResolveResourceType) Set(val *ResolveResourceType) {
	v.value = val
	v.isSet = true
}

func (v NullableResolveResourceType) IsSet() bool {
	return v.isSet
}

func (v *NullableResolveResourceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResolveResourceType(val *ResolveResourceType) *NullableResolveResourceType {
	return &NullableResolveResourceType{value: val, isSet: true}
}

func (v NullableResolveResourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResolveResourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
