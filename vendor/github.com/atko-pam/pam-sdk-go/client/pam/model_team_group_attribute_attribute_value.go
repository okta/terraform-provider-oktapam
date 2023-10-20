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

// TeamGroupAttributeAttributeValue - struct for TeamGroupAttributeAttributeValue
type TeamGroupAttributeAttributeValue struct {
	Int32  *int32
	String *string
}

// int32AsTeamGroupAttributeAttributeValue is a convenience function that returns int32 wrapped in TeamGroupAttributeAttributeValue
func Int32AsTeamGroupAttributeAttributeValue(v *int32) TeamGroupAttributeAttributeValue {
	return TeamGroupAttributeAttributeValue{
		Int32: v,
	}
}

// stringAsTeamGroupAttributeAttributeValue is a convenience function that returns string wrapped in TeamGroupAttributeAttributeValue
func StringAsTeamGroupAttributeAttributeValue(v *string) TeamGroupAttributeAttributeValue {
	return TeamGroupAttributeAttributeValue{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *TeamGroupAttributeAttributeValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Int32
	err = json.Unmarshal(data, &dst.Int32)
	if err == nil {
		jsonint32, _ := json.Marshal(dst.Int32)
		if string(jsonint32) == "{}" { // empty struct
			dst.Int32 = nil
		} else {
			match++
		}
	} else {
		dst.Int32 = nil
	}

	// try to unmarshal data into String
	err = json.Unmarshal(data, &dst.String)
	if err == nil {
		jsonstring, _ := json.Marshal(dst.String)
		if string(jsonstring) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Int32 = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(TeamGroupAttributeAttributeValue)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(TeamGroupAttributeAttributeValue)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TeamGroupAttributeAttributeValue) MarshalJSON() ([]byte, error) {
	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TeamGroupAttributeAttributeValue) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Int32 != nil {
		return obj.Int32
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableTeamGroupAttributeAttributeValue struct {
	value *TeamGroupAttributeAttributeValue
	isSet bool
}

func (v NullableTeamGroupAttributeAttributeValue) Get() *TeamGroupAttributeAttributeValue {
	return v.value
}

func (v *NullableTeamGroupAttributeAttributeValue) Set(val *TeamGroupAttributeAttributeValue) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamGroupAttributeAttributeValue) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamGroupAttributeAttributeValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamGroupAttributeAttributeValue(val *TeamGroupAttributeAttributeValue) *NullableTeamGroupAttributeAttributeValue {
	return &NullableTeamGroupAttributeAttributeValue{value: val, isSet: true}
}

func (v NullableTeamGroupAttributeAttributeValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamGroupAttributeAttributeValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}