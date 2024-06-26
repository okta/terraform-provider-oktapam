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

// AccessReportTargetResourceType The type of resource associated with the Access Report
type AccessReportTargetResourceType string

// List of AccessReportTargetResourceType
const (
	AccessReportTargetResourceType_SERVER AccessReportTargetResourceType = "server"
)

// All allowed values of AccessReportTargetResourceType enum
var AllowedAccessReportTargetResourceTypeEnumValues = []AccessReportTargetResourceType{
	"server",
}

func (v *AccessReportTargetResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccessReportTargetResourceType(value)

	*v = enumTypeValue

	return nil
}

// NewAccessReportTargetResourceTypeFromValue returns a pointer to a valid AccessReportTargetResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccessReportTargetResourceTypeFromValue(v string) (*AccessReportTargetResourceType, error) {
	ev := AccessReportTargetResourceType(v)

	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccessReportTargetResourceType) IsValid() bool {
	for _, existing := range AllowedAccessReportTargetResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccessReportTargetResourceType value
func (v AccessReportTargetResourceType) Ptr() *AccessReportTargetResourceType {
	return &v
}

type NullableAccessReportTargetResourceType struct {
	value *AccessReportTargetResourceType
	isSet bool
}

func (v NullableAccessReportTargetResourceType) Get() *AccessReportTargetResourceType {
	return v.value
}

func (v *NullableAccessReportTargetResourceType) Set(val *AccessReportTargetResourceType) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessReportTargetResourceType) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessReportTargetResourceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessReportTargetResourceType(val *AccessReportTargetResourceType) *NullableAccessReportTargetResourceType {
	return &NullableAccessReportTargetResourceType{value: val, isSet: true}
}

func (v NullableAccessReportTargetResourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessReportTargetResourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
