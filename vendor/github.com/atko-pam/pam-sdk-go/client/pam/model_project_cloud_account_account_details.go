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

// ProjectCloudAccountAccountDetails - The provider-specific account details
type ProjectCloudAccountAccountDetails struct {
	Aws *Aws
	Gce *Gce
}

// AwsAsProjectCloudAccountAccountDetails is a convenience function that returns Aws wrapped in ProjectCloudAccountAccountDetails
func AwsAsProjectCloudAccountAccountDetails(v *Aws) ProjectCloudAccountAccountDetails {
	return ProjectCloudAccountAccountDetails{
		Aws: v,
	}
}

// GceAsProjectCloudAccountAccountDetails is a convenience function that returns Gce wrapped in ProjectCloudAccountAccountDetails
func GceAsProjectCloudAccountAccountDetails(v *Gce) ProjectCloudAccountAccountDetails {
	return ProjectCloudAccountAccountDetails{
		Gce: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ProjectCloudAccountAccountDetails) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Aws
	err = json.Unmarshal(data, &dst.Aws)
	if err == nil {
		jsonAws, _ := json.Marshal(dst.Aws)
		if string(jsonAws) == "{}" { // empty struct
			dst.Aws = nil
		} else {
			match++
		}
	} else {
		dst.Aws = nil
	}

	// try to unmarshal data into Gce
	err = json.Unmarshal(data, &dst.Gce)
	if err == nil {
		jsonGce, _ := json.Marshal(dst.Gce)
		if string(jsonGce) == "{}" { // empty struct
			dst.Gce = nil
		} else {
			match++
		}
	} else {
		dst.Gce = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Aws = nil
		dst.Gce = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ProjectCloudAccountAccountDetails)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ProjectCloudAccountAccountDetails)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ProjectCloudAccountAccountDetails) MarshalJSON() ([]byte, error) {
	if src.Aws != nil {
		return json.Marshal(&src.Aws)
	}

	if src.Gce != nil {
		return json.Marshal(&src.Gce)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ProjectCloudAccountAccountDetails) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Aws != nil {
		return obj.Aws
	}

	if obj.Gce != nil {
		return obj.Gce
	}

	// all schemas are nil
	return nil
}

type NullableProjectCloudAccountAccountDetails struct {
	value *ProjectCloudAccountAccountDetails
	isSet bool
}

func (v NullableProjectCloudAccountAccountDetails) Get() *ProjectCloudAccountAccountDetails {
	return v.value
}

func (v *NullableProjectCloudAccountAccountDetails) Set(val *ProjectCloudAccountAccountDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectCloudAccountAccountDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectCloudAccountAccountDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectCloudAccountAccountDetails(val *ProjectCloudAccountAccountDetails) *NullableProjectCloudAccountAccountDetails {
	return &NullableProjectCloudAccountAccountDetails{value: val, isSet: true}
}

func (v NullableProjectCloudAccountAccountDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectCloudAccountAccountDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
