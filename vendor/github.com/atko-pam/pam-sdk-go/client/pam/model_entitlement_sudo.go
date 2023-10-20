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

// checks if the EntitlementSudo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementSudo{}

// EntitlementSudo struct for EntitlementSudo
type EntitlementSudo struct {
	// A list of environment variables to include when running Entitlement commands. See [the sudo documentation](https://www.sudo.ws/man/1.8.13/sudoers.man.html#Command_environment).
	AddEnv []string `json:"add_env,omitempty"`
	// A description of the Entitlement
	Description NullableString `json:"description,omitempty"`
	// The UUID of the Entitlement
	Id string `json:"id"`
	// The name of the Entitlement
	Name string `json:"name"`
	// Whether to allow commands to execute child processes
	OptNoExec NullableBool `json:"opt_no_exec,omitempty"`
	// Whether to require a password when sudo is run. This should generally not be used as Users don't require a password.
	OptNoPasswd NullableBool `json:"opt_no_passwd,omitempty"`
	// A non-root user account used to run the command
	OptRunAs NullableString `json:"opt_run_as,omitempty"`
	// Whether to allow overriding environment variables to commands
	OptSetEnv NullableBool `json:"opt_set_env,omitempty"`
	// A list of commands to allow
	StructuredCommands []EntitlementSudoStructuredCommandsInner `json:"structured_commands,omitempty"`
	// A list of environment variables to ignore when running Entitlement commands. See [the sudo documentation](https://www.sudo.ws/man/1.8.13/sudoers.man.html#Command_environment).
	SubEnv []string `json:"sub_env,omitempty"`
}

// NewEntitlementSudo instantiates a new EntitlementSudo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementSudo(id string, name string) *EntitlementSudo {
	this := EntitlementSudo{}
	this.Id = id
	this.Name = name
	return &this
}

// NewEntitlementSudoWithDefaults instantiates a new EntitlementSudo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementSudoWithDefaults() *EntitlementSudo {
	this := EntitlementSudo{}
	return &this
}

// GetAddEnv returns the AddEnv field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntitlementSudo) GetAddEnv() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AddEnv
}

// GetAddEnvOk returns a tuple with the AddEnv field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntitlementSudo) GetAddEnvOk() ([]string, bool) {
	if o == nil || IsNil(o.AddEnv) {
		return nil, false
	}
	return o.AddEnv, true
}

// HasAddEnv returns a boolean if a field has been set.
func (o *EntitlementSudo) HasAddEnv() bool {
	if o != nil && IsNil(o.AddEnv) {
		return true
	}

	return false
}

// SetAddEnv gets a reference to the given []string and assigns it to the AddEnv field.
func (o *EntitlementSudo) SetAddEnv(v []string) *EntitlementSudo {
	o.AddEnv = v
	return o
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntitlementSudo) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntitlementSudo) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *EntitlementSudo) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *EntitlementSudo) SetDescription(v string) *EntitlementSudo {
	o.Description.Set(&v)
	return o
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *EntitlementSudo) SetDescriptionNil() *EntitlementSudo {
	o.Description.Set(nil)
	return o
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *EntitlementSudo) UnsetDescription() *EntitlementSudo {
	o.Description.Unset()
	return o
}

// GetId returns the Id field value
func (o *EntitlementSudo) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EntitlementSudo) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EntitlementSudo) SetId(v string) *EntitlementSudo {
	o.Id = v
	return o
}

// GetName returns the Name field value
func (o *EntitlementSudo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EntitlementSudo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EntitlementSudo) SetName(v string) *EntitlementSudo {
	o.Name = v
	return o
}

// GetOptNoExec returns the OptNoExec field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntitlementSudo) GetOptNoExec() bool {
	if o == nil || IsNil(o.OptNoExec.Get()) {
		var ret bool
		return ret
	}
	return *o.OptNoExec.Get()
}

// GetOptNoExecOk returns a tuple with the OptNoExec field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntitlementSudo) GetOptNoExecOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.OptNoExec.Get(), o.OptNoExec.IsSet()
}

// HasOptNoExec returns a boolean if a field has been set.
func (o *EntitlementSudo) HasOptNoExec() bool {
	if o != nil && o.OptNoExec.IsSet() {
		return true
	}

	return false
}

// SetOptNoExec gets a reference to the given NullableBool and assigns it to the OptNoExec field.
func (o *EntitlementSudo) SetOptNoExec(v bool) *EntitlementSudo {
	o.OptNoExec.Set(&v)
	return o
}

// SetOptNoExecNil sets the value for OptNoExec to be an explicit nil
func (o *EntitlementSudo) SetOptNoExecNil() *EntitlementSudo {
	o.OptNoExec.Set(nil)
	return o
}

// UnsetOptNoExec ensures that no value is present for OptNoExec, not even an explicit nil
func (o *EntitlementSudo) UnsetOptNoExec() *EntitlementSudo {
	o.OptNoExec.Unset()
	return o
}

// GetOptNoPasswd returns the OptNoPasswd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntitlementSudo) GetOptNoPasswd() bool {
	if o == nil || IsNil(o.OptNoPasswd.Get()) {
		var ret bool
		return ret
	}
	return *o.OptNoPasswd.Get()
}

// GetOptNoPasswdOk returns a tuple with the OptNoPasswd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntitlementSudo) GetOptNoPasswdOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.OptNoPasswd.Get(), o.OptNoPasswd.IsSet()
}

// HasOptNoPasswd returns a boolean if a field has been set.
func (o *EntitlementSudo) HasOptNoPasswd() bool {
	if o != nil && o.OptNoPasswd.IsSet() {
		return true
	}

	return false
}

// SetOptNoPasswd gets a reference to the given NullableBool and assigns it to the OptNoPasswd field.
func (o *EntitlementSudo) SetOptNoPasswd(v bool) *EntitlementSudo {
	o.OptNoPasswd.Set(&v)
	return o
}

// SetOptNoPasswdNil sets the value for OptNoPasswd to be an explicit nil
func (o *EntitlementSudo) SetOptNoPasswdNil() *EntitlementSudo {
	o.OptNoPasswd.Set(nil)
	return o
}

// UnsetOptNoPasswd ensures that no value is present for OptNoPasswd, not even an explicit nil
func (o *EntitlementSudo) UnsetOptNoPasswd() *EntitlementSudo {
	o.OptNoPasswd.Unset()
	return o
}

// GetOptRunAs returns the OptRunAs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntitlementSudo) GetOptRunAs() string {
	if o == nil || IsNil(o.OptRunAs.Get()) {
		var ret string
		return ret
	}
	return *o.OptRunAs.Get()
}

// GetOptRunAsOk returns a tuple with the OptRunAs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntitlementSudo) GetOptRunAsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OptRunAs.Get(), o.OptRunAs.IsSet()
}

// HasOptRunAs returns a boolean if a field has been set.
func (o *EntitlementSudo) HasOptRunAs() bool {
	if o != nil && o.OptRunAs.IsSet() {
		return true
	}

	return false
}

// SetOptRunAs gets a reference to the given NullableString and assigns it to the OptRunAs field.
func (o *EntitlementSudo) SetOptRunAs(v string) *EntitlementSudo {
	o.OptRunAs.Set(&v)
	return o
}

// SetOptRunAsNil sets the value for OptRunAs to be an explicit nil
func (o *EntitlementSudo) SetOptRunAsNil() *EntitlementSudo {
	o.OptRunAs.Set(nil)
	return o
}

// UnsetOptRunAs ensures that no value is present for OptRunAs, not even an explicit nil
func (o *EntitlementSudo) UnsetOptRunAs() *EntitlementSudo {
	o.OptRunAs.Unset()
	return o
}

// GetOptSetEnv returns the OptSetEnv field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntitlementSudo) GetOptSetEnv() bool {
	if o == nil || IsNil(o.OptSetEnv.Get()) {
		var ret bool
		return ret
	}
	return *o.OptSetEnv.Get()
}

// GetOptSetEnvOk returns a tuple with the OptSetEnv field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntitlementSudo) GetOptSetEnvOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.OptSetEnv.Get(), o.OptSetEnv.IsSet()
}

// HasOptSetEnv returns a boolean if a field has been set.
func (o *EntitlementSudo) HasOptSetEnv() bool {
	if o != nil && o.OptSetEnv.IsSet() {
		return true
	}

	return false
}

// SetOptSetEnv gets a reference to the given NullableBool and assigns it to the OptSetEnv field.
func (o *EntitlementSudo) SetOptSetEnv(v bool) *EntitlementSudo {
	o.OptSetEnv.Set(&v)
	return o
}

// SetOptSetEnvNil sets the value for OptSetEnv to be an explicit nil
func (o *EntitlementSudo) SetOptSetEnvNil() *EntitlementSudo {
	o.OptSetEnv.Set(nil)
	return o
}

// UnsetOptSetEnv ensures that no value is present for OptSetEnv, not even an explicit nil
func (o *EntitlementSudo) UnsetOptSetEnv() *EntitlementSudo {
	o.OptSetEnv.Unset()
	return o
}

// GetStructuredCommands returns the StructuredCommands field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntitlementSudo) GetStructuredCommands() []EntitlementSudoStructuredCommandsInner {
	if o == nil {
		var ret []EntitlementSudoStructuredCommandsInner
		return ret
	}
	return o.StructuredCommands
}

// GetStructuredCommandsOk returns a tuple with the StructuredCommands field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntitlementSudo) GetStructuredCommandsOk() ([]EntitlementSudoStructuredCommandsInner, bool) {
	if o == nil || IsNil(o.StructuredCommands) {
		return nil, false
	}
	return o.StructuredCommands, true
}

// HasStructuredCommands returns a boolean if a field has been set.
func (o *EntitlementSudo) HasStructuredCommands() bool {
	if o != nil && IsNil(o.StructuredCommands) {
		return true
	}

	return false
}

// SetStructuredCommands gets a reference to the given []EntitlementSudoStructuredCommandsInner and assigns it to the StructuredCommands field.
func (o *EntitlementSudo) SetStructuredCommands(v []EntitlementSudoStructuredCommandsInner) *EntitlementSudo {
	o.StructuredCommands = v
	return o
}

// GetSubEnv returns the SubEnv field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntitlementSudo) GetSubEnv() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SubEnv
}

// GetSubEnvOk returns a tuple with the SubEnv field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntitlementSudo) GetSubEnvOk() ([]string, bool) {
	if o == nil || IsNil(o.SubEnv) {
		return nil, false
	}
	return o.SubEnv, true
}

// HasSubEnv returns a boolean if a field has been set.
func (o *EntitlementSudo) HasSubEnv() bool {
	if o != nil && IsNil(o.SubEnv) {
		return true
	}

	return false
}

// SetSubEnv gets a reference to the given []string and assigns it to the SubEnv field.
func (o *EntitlementSudo) SetSubEnv(v []string) *EntitlementSudo {
	o.SubEnv = v
	return o
}

func (o EntitlementSudo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementSudo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AddEnv != nil {
		toSerialize["add_env"] = o.AddEnv
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if o.OptNoExec.IsSet() {
		toSerialize["opt_no_exec"] = o.OptNoExec.Get()
	}
	if o.OptNoPasswd.IsSet() {
		toSerialize["opt_no_passwd"] = o.OptNoPasswd.Get()
	}
	if o.OptRunAs.IsSet() {
		toSerialize["opt_run_as"] = o.OptRunAs.Get()
	}
	if o.OptSetEnv.IsSet() {
		toSerialize["opt_set_env"] = o.OptSetEnv.Get()
	}
	if o.StructuredCommands != nil {
		toSerialize["structured_commands"] = o.StructuredCommands
	}
	if o.SubEnv != nil {
		toSerialize["sub_env"] = o.SubEnv
	}
	return toSerialize, nil
}

type NullableEntitlementSudo struct {
	value *EntitlementSudo
	isSet bool
}

func (v NullableEntitlementSudo) Get() *EntitlementSudo {
	return v.value
}

func (v *NullableEntitlementSudo) Set(val *EntitlementSudo) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementSudo) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementSudo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementSudo(val *EntitlementSudo) *NullableEntitlementSudo {
	return &NullableEntitlementSudo{value: val, isSet: true}
}

func (v NullableEntitlementSudo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementSudo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}