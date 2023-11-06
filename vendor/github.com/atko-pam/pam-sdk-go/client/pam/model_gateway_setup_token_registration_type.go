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

// GatewaySetupTokenRegistrationType The type of registration policy
type GatewaySetupTokenRegistrationType string

// List of GatewaySetupTokenRegistrationType
const (
	GatewaySetupTokenRegistrationType_GATEWAY_AGENT GatewaySetupTokenRegistrationType = "gateway-agent"
)

// All allowed values of GatewaySetupTokenRegistrationType enum
var AllowedGatewaySetupTokenRegistrationTypeEnumValues = []GatewaySetupTokenRegistrationType{
	"gateway-agent",
}

func (v *GatewaySetupTokenRegistrationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GatewaySetupTokenRegistrationType(value)
	for _, existing := range AllowedGatewaySetupTokenRegistrationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GatewaySetupTokenRegistrationType", value)
}

// NewGatewaySetupTokenRegistrationTypeFromValue returns a pointer to a valid GatewaySetupTokenRegistrationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGatewaySetupTokenRegistrationTypeFromValue(v string) (*GatewaySetupTokenRegistrationType, error) {
	ev := GatewaySetupTokenRegistrationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GatewaySetupTokenRegistrationType: valid values are %v", v, AllowedGatewaySetupTokenRegistrationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GatewaySetupTokenRegistrationType) IsValid() bool {
	for _, existing := range AllowedGatewaySetupTokenRegistrationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GatewaySetupTokenRegistrationType value
func (v GatewaySetupTokenRegistrationType) Ptr() *GatewaySetupTokenRegistrationType {
	return &v
}

type NullableGatewaySetupTokenRegistrationType struct {
	value *GatewaySetupTokenRegistrationType
	isSet bool
}

func (v NullableGatewaySetupTokenRegistrationType) Get() *GatewaySetupTokenRegistrationType {
	return v.value
}

func (v *NullableGatewaySetupTokenRegistrationType) Set(val *GatewaySetupTokenRegistrationType) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewaySetupTokenRegistrationType) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewaySetupTokenRegistrationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewaySetupTokenRegistrationType(val *GatewaySetupTokenRegistrationType) *NullableGatewaySetupTokenRegistrationType {
	return &NullableGatewaySetupTokenRegistrationType{value: val, isSet: true}
}

func (v NullableGatewaySetupTokenRegistrationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewaySetupTokenRegistrationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
