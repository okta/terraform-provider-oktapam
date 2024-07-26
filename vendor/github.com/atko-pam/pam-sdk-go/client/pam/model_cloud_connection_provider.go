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

// CloudConnectionProvider The cloud provider associated with the Cloud Connection. Currently, only accepts `aws`.
type CloudConnectionProvider string

// List of CloudConnectionProvider
const (
	CloudConnectionProvider_AWS CloudConnectionProvider = "aws"
)

// All allowed values of CloudConnectionProvider enum
var AllowedCloudConnectionProviderEnumValues = []CloudConnectionProvider{
	"aws",
}

func (v *CloudConnectionProvider) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CloudConnectionProvider(value)

	*v = enumTypeValue

	return nil
}

// NewCloudConnectionProviderFromValue returns a pointer to a valid CloudConnectionProvider
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCloudConnectionProviderFromValue(v string) (*CloudConnectionProvider, error) {
	ev := CloudConnectionProvider(v)

	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CloudConnectionProvider) IsValid() bool {
	for _, existing := range AllowedCloudConnectionProviderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CloudConnectionProvider value
func (v CloudConnectionProvider) Ptr() *CloudConnectionProvider {
	return &v
}

type NullableCloudConnectionProvider struct {
	value *CloudConnectionProvider
	isSet bool
}

func (v NullableCloudConnectionProvider) Get() *CloudConnectionProvider {
	return v.value
}

func (v *NullableCloudConnectionProvider) Set(val *CloudConnectionProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudConnectionProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudConnectionProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudConnectionProvider(val *CloudConnectionProvider) *NullableCloudConnectionProvider {
	return &NullableCloudConnectionProvider{value: val, isSet: true}
}

func (v NullableCloudConnectionProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudConnectionProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
