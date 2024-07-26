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

// DatabaseAccountType The type of the Database Static Account
type DatabaseAccountType string

// List of DatabaseAccountType
const (
	DatabaseAccountType_MYSQL_STATIC_ACCOUNT DatabaseAccountType = "mysql-static-account"
)

// All allowed values of DatabaseAccountType enum
var AllowedDatabaseAccountTypeEnumValues = []DatabaseAccountType{
	"mysql-static-account",
}

func (v *DatabaseAccountType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DatabaseAccountType(value)

	*v = enumTypeValue

	return nil
}

// NewDatabaseAccountTypeFromValue returns a pointer to a valid DatabaseAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDatabaseAccountTypeFromValue(v string) (*DatabaseAccountType, error) {
	ev := DatabaseAccountType(v)

	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DatabaseAccountType) IsValid() bool {
	for _, existing := range AllowedDatabaseAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DatabaseAccountType value
func (v DatabaseAccountType) Ptr() *DatabaseAccountType {
	return &v
}

type NullableDatabaseAccountType struct {
	value *DatabaseAccountType
	isSet bool
}

func (v NullableDatabaseAccountType) Get() *DatabaseAccountType {
	return v.value
}

func (v *NullableDatabaseAccountType) Set(val *DatabaseAccountType) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseAccountType) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseAccountType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseAccountType(val *DatabaseAccountType) *NullableDatabaseAccountType {
	return &NullableDatabaseAccountType{value: val, isSet: true}
}

func (v NullableDatabaseAccountType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseAccountType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
