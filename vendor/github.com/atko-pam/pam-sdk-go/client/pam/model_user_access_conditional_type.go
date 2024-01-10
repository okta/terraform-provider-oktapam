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
	"fmt"
)

// UserAccessConditionalType The type of condition
type UserAccessConditionalType string

// List of UserAccessConditionalType
const (
	UserAccessConditionalType_ACCESS_REQUEST                 UserAccessConditionalType = "access_request"
	UserAccessConditionalType_GATEWAY                        UserAccessConditionalType = "gateway"
	UserAccessConditionalType_GATEWAY_WITH_SESSION_RECORDING UserAccessConditionalType = "gateway_with_session_recording"
	UserAccessConditionalType_MFA                            UserAccessConditionalType = "mfa"
)

// All allowed values of UserAccessConditionalType enum
var AllowedUserAccessConditionalTypeEnumValues = []UserAccessConditionalType{
	"access_request",
	"gateway",
	"gateway_with_session_recording",
	"mfa",
}

func (v *UserAccessConditionalType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UserAccessConditionalType(value)
	for _, existing := range AllowedUserAccessConditionalTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UserAccessConditionalType", value)
}

// NewUserAccessConditionalTypeFromValue returns a pointer to a valid UserAccessConditionalType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUserAccessConditionalTypeFromValue(v string) (*UserAccessConditionalType, error) {
	ev := UserAccessConditionalType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UserAccessConditionalType: valid values are %v", v, AllowedUserAccessConditionalTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UserAccessConditionalType) IsValid() bool {
	for _, existing := range AllowedUserAccessConditionalTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserAccessConditionalType value
func (v UserAccessConditionalType) Ptr() *UserAccessConditionalType {
	return &v
}

type NullableUserAccessConditionalType struct {
	value *UserAccessConditionalType
	isSet bool
}

func (v NullableUserAccessConditionalType) Get() *UserAccessConditionalType {
	return v.value
}

func (v *NullableUserAccessConditionalType) Set(val *UserAccessConditionalType) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAccessConditionalType) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAccessConditionalType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAccessConditionalType(val *UserAccessConditionalType) *NullableUserAccessConditionalType {
	return &NullableUserAccessConditionalType{value: val, isSet: true}
}

func (v NullableUserAccessConditionalType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAccessConditionalType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
