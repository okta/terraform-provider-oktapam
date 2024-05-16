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

// SudoCommandBundleStructuredCommandsInnerCommand - struct for SudoCommandBundleStructuredCommandsInnerCommand
type SudoCommandBundleStructuredCommandsInnerCommand struct {
	Executable *Executable
	String     *string
}

// ExecutableAsSudoCommandBundleStructuredCommandsInnerCommand is a convenience function that returns Executable wrapped in SudoCommandBundleStructuredCommandsInnerCommand
func ExecutableAsSudoCommandBundleStructuredCommandsInnerCommand(v *Executable) SudoCommandBundleStructuredCommandsInnerCommand {
	return SudoCommandBundleStructuredCommandsInnerCommand{
		Executable: v,
	}
}

// stringAsSudoCommandBundleStructuredCommandsInnerCommand is a convenience function that returns string wrapped in SudoCommandBundleStructuredCommandsInnerCommand
func StringAsSudoCommandBundleStructuredCommandsInnerCommand(v *string) SudoCommandBundleStructuredCommandsInnerCommand {
	return SudoCommandBundleStructuredCommandsInnerCommand{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SudoCommandBundleStructuredCommandsInnerCommand) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Executable
	err = json.Unmarshal(data, &dst.Executable)
	if err == nil {
		jsonExecutable, _ := json.Marshal(dst.Executable)
		if string(jsonExecutable) == "{}" { // empty struct
			dst.Executable = nil
		} else {
			match++
		}
	} else {
		dst.Executable = nil
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
		dst.Executable = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SudoCommandBundleStructuredCommandsInnerCommand)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SudoCommandBundleStructuredCommandsInnerCommand)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SudoCommandBundleStructuredCommandsInnerCommand) MarshalJSON() ([]byte, error) {
	if src.Executable != nil {
		return json.Marshal(&src.Executable)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SudoCommandBundleStructuredCommandsInnerCommand) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Executable != nil {
		return obj.Executable
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableSudoCommandBundleStructuredCommandsInnerCommand struct {
	value *SudoCommandBundleStructuredCommandsInnerCommand
	isSet bool
}

func (v NullableSudoCommandBundleStructuredCommandsInnerCommand) Get() *SudoCommandBundleStructuredCommandsInnerCommand {
	return v.value
}

func (v *NullableSudoCommandBundleStructuredCommandsInnerCommand) Set(val *SudoCommandBundleStructuredCommandsInnerCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableSudoCommandBundleStructuredCommandsInnerCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableSudoCommandBundleStructuredCommandsInnerCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSudoCommandBundleStructuredCommandsInnerCommand(val *SudoCommandBundleStructuredCommandsInnerCommand) *NullableSudoCommandBundleStructuredCommandsInnerCommand {
	return &NullableSudoCommandBundleStructuredCommandsInnerCommand{value: val, isSet: true}
}

func (v NullableSudoCommandBundleStructuredCommandsInnerCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSudoCommandBundleStructuredCommandsInnerCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
