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

// CheckedOutResourceByUserDetailsResourceDetails - The additional details of the resource. Only returned if a `resource_type` query was specified.
type CheckedOutResourceByUserDetailsResourceDetails struct {
	OktaUniversalDirectoryAccountCheckedOutResourceDetails *OktaUniversalDirectoryAccountCheckedOutResourceDetails
	SaasAppAccountCheckedOutResourceDetails                *SaasAppAccountCheckedOutResourceDetails
	ServerAccountCheckedOutResourceDetails                 *ServerAccountCheckedOutResourceDetails
}

// OktaUniversalDirectoryAccountCheckedOutResourceDetailsAsCheckedOutResourceByUserDetailsResourceDetails is a convenience function that returns OktaUniversalDirectoryAccountCheckedOutResourceDetails wrapped in CheckedOutResourceByUserDetailsResourceDetails
func OktaUniversalDirectoryAccountCheckedOutResourceDetailsAsCheckedOutResourceByUserDetailsResourceDetails(v *OktaUniversalDirectoryAccountCheckedOutResourceDetails) CheckedOutResourceByUserDetailsResourceDetails {
	return CheckedOutResourceByUserDetailsResourceDetails{
		OktaUniversalDirectoryAccountCheckedOutResourceDetails: v,
	}
}

// SaasAppAccountCheckedOutResourceDetailsAsCheckedOutResourceByUserDetailsResourceDetails is a convenience function that returns SaasAppAccountCheckedOutResourceDetails wrapped in CheckedOutResourceByUserDetailsResourceDetails
func SaasAppAccountCheckedOutResourceDetailsAsCheckedOutResourceByUserDetailsResourceDetails(v *SaasAppAccountCheckedOutResourceDetails) CheckedOutResourceByUserDetailsResourceDetails {
	return CheckedOutResourceByUserDetailsResourceDetails{
		SaasAppAccountCheckedOutResourceDetails: v,
	}
}

// ServerAccountCheckedOutResourceDetailsAsCheckedOutResourceByUserDetailsResourceDetails is a convenience function that returns ServerAccountCheckedOutResourceDetails wrapped in CheckedOutResourceByUserDetailsResourceDetails
func ServerAccountCheckedOutResourceDetailsAsCheckedOutResourceByUserDetailsResourceDetails(v *ServerAccountCheckedOutResourceDetails) CheckedOutResourceByUserDetailsResourceDetails {
	return CheckedOutResourceByUserDetailsResourceDetails{
		ServerAccountCheckedOutResourceDetails: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CheckedOutResourceByUserDetailsResourceDetails) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'OktaUniversalDirectoryAccountCheckedOutResourceDetails'
	if jsonDict["_type"] == "OktaUniversalDirectoryAccountCheckedOutResourceDetails" {
		// try to unmarshal JSON data into OktaUniversalDirectoryAccountCheckedOutResourceDetails
		err = json.Unmarshal(data, &dst.OktaUniversalDirectoryAccountCheckedOutResourceDetails)
		if err == nil {
			return nil // data stored in dst.OktaUniversalDirectoryAccountCheckedOutResourceDetails, return on the first match
		} else {
			dst.OktaUniversalDirectoryAccountCheckedOutResourceDetails = nil
			return fmt.Errorf("failed to unmarshal CheckedOutResourceByUserDetailsResourceDetails as OktaUniversalDirectoryAccountCheckedOutResourceDetails: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SaasAppAccountCheckedOutResourceDetails'
	if jsonDict["_type"] == "SaasAppAccountCheckedOutResourceDetails" {
		// try to unmarshal JSON data into SaasAppAccountCheckedOutResourceDetails
		err = json.Unmarshal(data, &dst.SaasAppAccountCheckedOutResourceDetails)
		if err == nil {
			return nil // data stored in dst.SaasAppAccountCheckedOutResourceDetails, return on the first match
		} else {
			dst.SaasAppAccountCheckedOutResourceDetails = nil
			return fmt.Errorf("failed to unmarshal CheckedOutResourceByUserDetailsResourceDetails as SaasAppAccountCheckedOutResourceDetails: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ServerAccountCheckedOutResourceDetails'
	if jsonDict["_type"] == "ServerAccountCheckedOutResourceDetails" {
		// try to unmarshal JSON data into ServerAccountCheckedOutResourceDetails
		err = json.Unmarshal(data, &dst.ServerAccountCheckedOutResourceDetails)
		if err == nil {
			return nil // data stored in dst.ServerAccountCheckedOutResourceDetails, return on the first match
		} else {
			dst.ServerAccountCheckedOutResourceDetails = nil
			return fmt.Errorf("failed to unmarshal CheckedOutResourceByUserDetailsResourceDetails as ServerAccountCheckedOutResourceDetails: %s", err.Error())
		}
	}

	// check if the discriminator value is 'managed_saas_app_account_password_login'
	if jsonDict["_type"] == "managed_saas_app_account_password_login" {
		// try to unmarshal JSON data into SaasAppAccountCheckedOutResourceDetails
		err = json.Unmarshal(data, &dst.SaasAppAccountCheckedOutResourceDetails)
		if err == nil {
			return nil // data stored in dst.SaasAppAccountCheckedOutResourceDetails, return on the first match
		} else {
			dst.SaasAppAccountCheckedOutResourceDetails = nil
			return fmt.Errorf("failed to unmarshal CheckedOutResourceByUserDetailsResourceDetails as SaasAppAccountCheckedOutResourceDetails: %s", err.Error())
		}
	}

	// check if the discriminator value is 'okta_universal_directory_account_password_login'
	if jsonDict["_type"] == "okta_universal_directory_account_password_login" {
		// try to unmarshal JSON data into OktaUniversalDirectoryAccountCheckedOutResourceDetails
		err = json.Unmarshal(data, &dst.OktaUniversalDirectoryAccountCheckedOutResourceDetails)
		if err == nil {
			return nil // data stored in dst.OktaUniversalDirectoryAccountCheckedOutResourceDetails, return on the first match
		} else {
			dst.OktaUniversalDirectoryAccountCheckedOutResourceDetails = nil
			return fmt.Errorf("failed to unmarshal CheckedOutResourceByUserDetailsResourceDetails as OktaUniversalDirectoryAccountCheckedOutResourceDetails: %s", err.Error())
		}
	}

	// check if the discriminator value is 'server_account_password_login'
	if jsonDict["_type"] == "server_account_password_login" {
		// try to unmarshal JSON data into ServerAccountCheckedOutResourceDetails
		err = json.Unmarshal(data, &dst.ServerAccountCheckedOutResourceDetails)
		if err == nil {
			return nil // data stored in dst.ServerAccountCheckedOutResourceDetails, return on the first match
		} else {
			dst.ServerAccountCheckedOutResourceDetails = nil
			return fmt.Errorf("failed to unmarshal CheckedOutResourceByUserDetailsResourceDetails as ServerAccountCheckedOutResourceDetails: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CheckedOutResourceByUserDetailsResourceDetails) MarshalJSON() ([]byte, error) {
	if src.OktaUniversalDirectoryAccountCheckedOutResourceDetails != nil {
		return json.Marshal(&src.OktaUniversalDirectoryAccountCheckedOutResourceDetails)
	}

	if src.SaasAppAccountCheckedOutResourceDetails != nil {
		return json.Marshal(&src.SaasAppAccountCheckedOutResourceDetails)
	}

	if src.ServerAccountCheckedOutResourceDetails != nil {
		return json.Marshal(&src.ServerAccountCheckedOutResourceDetails)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CheckedOutResourceByUserDetailsResourceDetails) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.OktaUniversalDirectoryAccountCheckedOutResourceDetails != nil {
		return obj.OktaUniversalDirectoryAccountCheckedOutResourceDetails
	}

	if obj.SaasAppAccountCheckedOutResourceDetails != nil {
		return obj.SaasAppAccountCheckedOutResourceDetails
	}

	if obj.ServerAccountCheckedOutResourceDetails != nil {
		return obj.ServerAccountCheckedOutResourceDetails
	}

	// all schemas are nil
	return nil
}

type NullableCheckedOutResourceByUserDetailsResourceDetails struct {
	value *CheckedOutResourceByUserDetailsResourceDetails
	isSet bool
}

func (v NullableCheckedOutResourceByUserDetailsResourceDetails) Get() *CheckedOutResourceByUserDetailsResourceDetails {
	return v.value
}

func (v *NullableCheckedOutResourceByUserDetailsResourceDetails) Set(val *CheckedOutResourceByUserDetailsResourceDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckedOutResourceByUserDetailsResourceDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckedOutResourceByUserDetailsResourceDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckedOutResourceByUserDetailsResourceDetails(val *CheckedOutResourceByUserDetailsResourceDetails) *NullableCheckedOutResourceByUserDetailsResourceDetails {
	return &NullableCheckedOutResourceByUserDetailsResourceDetails{value: val, isSet: true}
}

func (v NullableCheckedOutResourceByUserDetailsResourceDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckedOutResourceByUserDetailsResourceDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
