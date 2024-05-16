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

// CheckinType Identifies the trigger for the check-in process
type CheckinType string

// List of CheckinType
const (
	CheckinType_SELF                    CheckinType = "self"
	CheckinType_ADMIN                   CheckinType = "admin"
	CheckinType_SYSTEM_CHECKOUT_EXPIRED CheckinType = "system_checkout_expired"
	CheckinType_SYSTEM_USER_LOST_ACCESS CheckinType = "system_user_lost_access"
)

// All allowed values of CheckinType enum
var AllowedCheckinTypeEnumValues = []CheckinType{
	"self",
	"admin",
	"system_checkout_expired",
	"system_user_lost_access",
}

func (v *CheckinType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CheckinType(value)

	*v = enumTypeValue

	return nil
}

// NewCheckinTypeFromValue returns a pointer to a valid CheckinType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCheckinTypeFromValue(v string) (*CheckinType, error) {
	ev := CheckinType(v)

	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CheckinType) IsValid() bool {
	for _, existing := range AllowedCheckinTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CheckinType value
func (v CheckinType) Ptr() *CheckinType {
	return &v
}

type NullableCheckinType struct {
	value *CheckinType
	isSet bool
}

func (v NullableCheckinType) Get() *CheckinType {
	return v.value
}

func (v *NullableCheckinType) Set(val *CheckinType) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckinType) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckinType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckinType(val *CheckinType) *NullableCheckinType {
	return &NullableCheckinType{value: val, isSet: true}
}

func (v NullableCheckinType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckinType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
