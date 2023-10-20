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
)

// checks if the EntitlementSudoStructuredCommandsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementSudoStructuredCommandsInner{}

// EntitlementSudoStructuredCommandsInner struct for EntitlementSudoStructuredCommandsInner
type EntitlementSudoStructuredCommandsInner struct {
	Args            NullableEntitlementSudoStructuredCommandsInnerArgs     `json:"args,omitempty"`
	ArgsType        NullableEntitlementSudoStructuredCommandsInnerArgsType `json:"args_type,omitempty"`
	Command         *EntitlementSudoStructuredCommandsInnerCommand         `json:"command,omitempty"`
	CommandType     *EntitlementSudoStructuredCommandType                  `json:"command_type,omitempty"`
	RenderedCommand *string                                                `json:"rendered_command,omitempty"`
}

// NewEntitlementSudoStructuredCommandsInner instantiates a new EntitlementSudoStructuredCommandsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementSudoStructuredCommandsInner() *EntitlementSudoStructuredCommandsInner {
	this := EntitlementSudoStructuredCommandsInner{}
	return &this
}

// NewEntitlementSudoStructuredCommandsInnerWithDefaults instantiates a new EntitlementSudoStructuredCommandsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementSudoStructuredCommandsInnerWithDefaults() *EntitlementSudoStructuredCommandsInner {
	this := EntitlementSudoStructuredCommandsInner{}
	return &this
}

// GetArgs returns the Args field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntitlementSudoStructuredCommandsInner) GetArgs() EntitlementSudoStructuredCommandsInnerArgs {
	if o == nil || IsNil(o.Args.Get()) {
		var ret EntitlementSudoStructuredCommandsInnerArgs
		return ret
	}
	return *o.Args.Get()
}

// GetArgsOk returns a tuple with the Args field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntitlementSudoStructuredCommandsInner) GetArgsOk() (*EntitlementSudoStructuredCommandsInnerArgs, bool) {
	if o == nil {
		return nil, false
	}
	return o.Args.Get(), o.Args.IsSet()
}

// HasArgs returns a boolean if a field has been set.
func (o *EntitlementSudoStructuredCommandsInner) HasArgs() bool {
	if o != nil && o.Args.IsSet() {
		return true
	}

	return false
}

// SetArgs gets a reference to the given NullableEntitlementSudoStructuredCommandsInnerArgs and assigns it to the Args field.
func (o *EntitlementSudoStructuredCommandsInner) SetArgs(v EntitlementSudoStructuredCommandsInnerArgs) *EntitlementSudoStructuredCommandsInner {
	o.Args.Set(&v)
	return o
}

// SetArgsNil sets the value for Args to be an explicit nil
func (o *EntitlementSudoStructuredCommandsInner) SetArgsNil() *EntitlementSudoStructuredCommandsInner {
	o.Args.Set(nil)
	return o
}

// UnsetArgs ensures that no value is present for Args, not even an explicit nil
func (o *EntitlementSudoStructuredCommandsInner) UnsetArgs() *EntitlementSudoStructuredCommandsInner {
	o.Args.Unset()
	return o
}

// GetArgsType returns the ArgsType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntitlementSudoStructuredCommandsInner) GetArgsType() EntitlementSudoStructuredCommandsInnerArgsType {
	if o == nil || IsNil(o.ArgsType.Get()) {
		var ret EntitlementSudoStructuredCommandsInnerArgsType
		return ret
	}
	return *o.ArgsType.Get()
}

// GetArgsTypeOk returns a tuple with the ArgsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntitlementSudoStructuredCommandsInner) GetArgsTypeOk() (*EntitlementSudoStructuredCommandsInnerArgsType, bool) {
	if o == nil {
		return nil, false
	}
	return o.ArgsType.Get(), o.ArgsType.IsSet()
}

// HasArgsType returns a boolean if a field has been set.
func (o *EntitlementSudoStructuredCommandsInner) HasArgsType() bool {
	if o != nil && o.ArgsType.IsSet() {
		return true
	}

	return false
}

// SetArgsType gets a reference to the given NullableEntitlementSudoStructuredCommandsInnerArgsType and assigns it to the ArgsType field.
func (o *EntitlementSudoStructuredCommandsInner) SetArgsType(v EntitlementSudoStructuredCommandsInnerArgsType) *EntitlementSudoStructuredCommandsInner {
	o.ArgsType.Set(&v)
	return o
}

// SetArgsTypeNil sets the value for ArgsType to be an explicit nil
func (o *EntitlementSudoStructuredCommandsInner) SetArgsTypeNil() *EntitlementSudoStructuredCommandsInner {
	o.ArgsType.Set(nil)
	return o
}

// UnsetArgsType ensures that no value is present for ArgsType, not even an explicit nil
func (o *EntitlementSudoStructuredCommandsInner) UnsetArgsType() *EntitlementSudoStructuredCommandsInner {
	o.ArgsType.Unset()
	return o
}

// GetCommand returns the Command field value if set, zero value otherwise.
func (o *EntitlementSudoStructuredCommandsInner) GetCommand() EntitlementSudoStructuredCommandsInnerCommand {
	if o == nil || IsNil(o.Command) {
		var ret EntitlementSudoStructuredCommandsInnerCommand
		return ret
	}
	return *o.Command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementSudoStructuredCommandsInner) GetCommandOk() (*EntitlementSudoStructuredCommandsInnerCommand, bool) {
	if o == nil || IsNil(o.Command) {
		return nil, false
	}
	return o.Command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *EntitlementSudoStructuredCommandsInner) HasCommand() bool {
	if o != nil && !IsNil(o.Command) {
		return true
	}

	return false
}

// SetCommand gets a reference to the given EntitlementSudoStructuredCommandsInnerCommand and assigns it to the Command field.
func (o *EntitlementSudoStructuredCommandsInner) SetCommand(v EntitlementSudoStructuredCommandsInnerCommand) *EntitlementSudoStructuredCommandsInner {
	o.Command = &v
	return o
}

// GetCommandType returns the CommandType field value if set, zero value otherwise.
func (o *EntitlementSudoStructuredCommandsInner) GetCommandType() EntitlementSudoStructuredCommandType {
	if o == nil || IsNil(o.CommandType) {
		var ret EntitlementSudoStructuredCommandType
		return ret
	}
	return *o.CommandType
}

// GetCommandTypeOk returns a tuple with the CommandType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementSudoStructuredCommandsInner) GetCommandTypeOk() (*EntitlementSudoStructuredCommandType, bool) {
	if o == nil || IsNil(o.CommandType) {
		return nil, false
	}
	return o.CommandType, true
}

// HasCommandType returns a boolean if a field has been set.
func (o *EntitlementSudoStructuredCommandsInner) HasCommandType() bool {
	if o != nil && !IsNil(o.CommandType) {
		return true
	}

	return false
}

// SetCommandType gets a reference to the given EntitlementSudoStructuredCommandType and assigns it to the CommandType field.
func (o *EntitlementSudoStructuredCommandsInner) SetCommandType(v EntitlementSudoStructuredCommandType) *EntitlementSudoStructuredCommandsInner {
	o.CommandType = &v
	return o
}

// GetRenderedCommand returns the RenderedCommand field value if set, zero value otherwise.
func (o *EntitlementSudoStructuredCommandsInner) GetRenderedCommand() string {
	if o == nil || IsNil(o.RenderedCommand) {
		var ret string
		return ret
	}
	return *o.RenderedCommand
}

// GetRenderedCommandOk returns a tuple with the RenderedCommand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementSudoStructuredCommandsInner) GetRenderedCommandOk() (*string, bool) {
	if o == nil || IsNil(o.RenderedCommand) {
		return nil, false
	}
	return o.RenderedCommand, true
}

// HasRenderedCommand returns a boolean if a field has been set.
func (o *EntitlementSudoStructuredCommandsInner) HasRenderedCommand() bool {
	if o != nil && !IsNil(o.RenderedCommand) {
		return true
	}

	return false
}

// SetRenderedCommand gets a reference to the given string and assigns it to the RenderedCommand field.
func (o *EntitlementSudoStructuredCommandsInner) SetRenderedCommand(v string) *EntitlementSudoStructuredCommandsInner {
	o.RenderedCommand = &v
	return o
}

func (o EntitlementSudoStructuredCommandsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementSudoStructuredCommandsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Args.IsSet() {
		toSerialize["args"] = o.Args.Get()
	}
	if o.ArgsType.IsSet() {
		toSerialize["args_type"] = o.ArgsType.Get()
	}
	if !IsNil(o.Command) {
		toSerialize["command"] = o.Command
	}
	if !IsNil(o.CommandType) {
		toSerialize["command_type"] = o.CommandType
	}
	if !IsNil(o.RenderedCommand) {
		toSerialize["rendered_command"] = o.RenderedCommand
	}
	return toSerialize, nil
}

type NullableEntitlementSudoStructuredCommandsInner struct {
	value *EntitlementSudoStructuredCommandsInner
	isSet bool
}

func (v NullableEntitlementSudoStructuredCommandsInner) Get() *EntitlementSudoStructuredCommandsInner {
	return v.value
}

func (v *NullableEntitlementSudoStructuredCommandsInner) Set(val *EntitlementSudoStructuredCommandsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementSudoStructuredCommandsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementSudoStructuredCommandsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementSudoStructuredCommandsInner(val *EntitlementSudoStructuredCommandsInner) *NullableEntitlementSudoStructuredCommandsInner {
	return &NullableEntitlementSudoStructuredCommandsInner{value: val, isSet: true}
}

func (v NullableEntitlementSudoStructuredCommandsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementSudoStructuredCommandsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
