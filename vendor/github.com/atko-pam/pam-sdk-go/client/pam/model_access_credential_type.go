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

// AccessCredentialType The type of credential used to access the resource
type AccessCredentialType string

// List of AccessCredentialType
const (
	AccessCredentialType_MANAGED                      AccessCredentialType = "managed"
	AccessCredentialType_PASSWORD                     AccessCredentialType = "password"
	AccessCredentialType_SSH_CERTIFICATE              AccessCredentialType = "ssh-certificate"
	AccessCredentialType_SSH_CERTIFICATE_ADMIN        AccessCredentialType = "ssh-certificate-admin"
	AccessCredentialType_RDP_BROKER_CERTIFICATE       AccessCredentialType = "rdp-broker-certificate"
	AccessCredentialType_RDP_BROKER_CERTIFICATE_ADMIN AccessCredentialType = "rdp-broker-certificate-admin"
	AccessCredentialType_ENCRYPTED_SSH_PASSWORD       AccessCredentialType = "encrypted-ssh-password"
	AccessCredentialType_ENCRYPTED_RDP_PASSWORD       AccessCredentialType = "encrypted-rdp-password"
)

// All allowed values of AccessCredentialType enum
var AllowedAccessCredentialTypeEnumValues = []AccessCredentialType{
	"managed",
	"password",
	"ssh-certificate",
	"ssh-certificate-admin",
	"rdp-broker-certificate",
	"rdp-broker-certificate-admin",
	"encrypted-ssh-password",
	"encrypted-rdp-password",
}

func (v *AccessCredentialType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccessCredentialType(value)

	*v = enumTypeValue

	return nil
}

// NewAccessCredentialTypeFromValue returns a pointer to a valid AccessCredentialType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccessCredentialTypeFromValue(v string) (*AccessCredentialType, error) {
	ev := AccessCredentialType(v)

	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccessCredentialType) IsValid() bool {
	for _, existing := range AllowedAccessCredentialTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccessCredentialType value
func (v AccessCredentialType) Ptr() *AccessCredentialType {
	return &v
}

type NullableAccessCredentialType struct {
	value *AccessCredentialType
	isSet bool
}

func (v NullableAccessCredentialType) Get() *AccessCredentialType {
	return v.value
}

func (v *NullableAccessCredentialType) Set(val *AccessCredentialType) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessCredentialType) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessCredentialType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessCredentialType(val *AccessCredentialType) *NullableAccessCredentialType {
	return &NullableAccessCredentialType{value: val, isSet: true}
}

func (v NullableAccessCredentialType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessCredentialType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
