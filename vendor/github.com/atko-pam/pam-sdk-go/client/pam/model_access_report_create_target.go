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

// AccessReportCreateTarget - struct for AccessReportCreateTarget
type AccessReportCreateTarget struct {
	AccessReportTargetResource *AccessReportTargetResource
	AccessReportTargetUser     *AccessReportTargetUser
}

// AccessReportTargetResourceAsAccessReportCreateTarget is a convenience function that returns AccessReportTargetResource wrapped in AccessReportCreateTarget
func AccessReportTargetResourceAsAccessReportCreateTarget(v *AccessReportTargetResource) AccessReportCreateTarget {
	return AccessReportCreateTarget{
		AccessReportTargetResource: v,
	}
}

// AccessReportTargetUserAsAccessReportCreateTarget is a convenience function that returns AccessReportTargetUser wrapped in AccessReportCreateTarget
func AccessReportTargetUserAsAccessReportCreateTarget(v *AccessReportTargetUser) AccessReportCreateTarget {
	return AccessReportCreateTarget{
		AccessReportTargetUser: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AccessReportCreateTarget) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AccessReportTargetResource
	err = json.Unmarshal(data, &dst.AccessReportTargetResource)
	if err == nil {
		jsonAccessReportTargetResource, _ := json.Marshal(dst.AccessReportTargetResource)
		if string(jsonAccessReportTargetResource) == "{}" { // empty struct
			dst.AccessReportTargetResource = nil
		} else {
			match++
		}
	} else {
		dst.AccessReportTargetResource = nil
	}

	// try to unmarshal data into AccessReportTargetUser
	err = json.Unmarshal(data, &dst.AccessReportTargetUser)
	if err == nil {
		jsonAccessReportTargetUser, _ := json.Marshal(dst.AccessReportTargetUser)
		if string(jsonAccessReportTargetUser) == "{}" { // empty struct
			dst.AccessReportTargetUser = nil
		} else {
			match++
		}
	} else {
		dst.AccessReportTargetUser = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AccessReportTargetResource = nil
		dst.AccessReportTargetUser = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AccessReportCreateTarget)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AccessReportCreateTarget)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AccessReportCreateTarget) MarshalJSON() ([]byte, error) {
	if src.AccessReportTargetResource != nil {
		return json.Marshal(&src.AccessReportTargetResource)
	}

	if src.AccessReportTargetUser != nil {
		return json.Marshal(&src.AccessReportTargetUser)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AccessReportCreateTarget) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AccessReportTargetResource != nil {
		return obj.AccessReportTargetResource
	}

	if obj.AccessReportTargetUser != nil {
		return obj.AccessReportTargetUser
	}

	// all schemas are nil
	return nil
}

type NullableAccessReportCreateTarget struct {
	value *AccessReportCreateTarget
	isSet bool
}

func (v NullableAccessReportCreateTarget) Get() *AccessReportCreateTarget {
	return v.value
}

func (v *NullableAccessReportCreateTarget) Set(val *AccessReportCreateTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessReportCreateTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessReportCreateTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessReportCreateTarget(val *AccessReportCreateTarget) *NullableAccessReportCreateTarget {
	return &NullableAccessReportCreateTarget{value: val, isSet: true}
}

func (v NullableAccessReportCreateTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessReportCreateTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
