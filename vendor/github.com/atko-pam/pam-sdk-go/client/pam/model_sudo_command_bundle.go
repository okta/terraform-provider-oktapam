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
	"time"
)

// checks if the SudoCommandBundle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SudoCommandBundle{}

// SudoCommandBundle struct for SudoCommandBundle
type SudoCommandBundle struct {
	// A list of environment variables to include when running sudo commands. See [the sudo documentation](https://www.sudo.ws/man/1.8.13/sudoers.man.html#Command_environment).
	AddEnv []string `json:"add_env,omitempty"`
	// A description of the Sudo Command bundle
	Description *string `json:"description,omitempty"`
	// The UUID of the Sudo Command bundle
	Id *string `json:"id,omitempty"`
	// The name of the Sudo Command bundle. This controls the ordering of all bundles within your Team.
	Name string `json:"name"`
	// Whether to allow commands to execute child processes
	NoExec NullableBool `json:"no_exec,omitempty"`
	// Whether to require a password when sudo is run. This should generally not be used as Users don't require a password.
	NoPasswd NullableBool `json:"no_passwd,omitempty"`
	// A non-root user account used to run the command
	RunAs NullableString `json:"run_as,omitempty"`
	// Whether to allow overriding environment variables to commands
	SetEnv NullableBool `json:"set_env,omitempty"`
	// A list of commands to allow
	StructuredCommands []SudoCommandBundleStructuredCommandsInner `json:"structured_commands,omitempty"`
	// A list of environment variables to ignore when running the commands. See [the sudo documentation](https://www.sudo.ws/man/1.8.13/sudoers.man.html#Command_environment).
	SubEnv []string `json:"sub_env,omitempty"`
	// A timestamp indicating when the Sudo Command bundle was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The username of the User who created the Sudo Command bundle
	CreatedBy *string `json:"created_by,omitempty"`
	// A timestamp indicating when the Sudo Command bundle was last updated
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The username of the User who last updated the Sudo Command bundle
	UpdatedBy *string `json:"updated_by,omitempty"`
}

// NewSudoCommandBundle instantiates a new SudoCommandBundle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSudoCommandBundle(name string) *SudoCommandBundle {
	this := SudoCommandBundle{}
	this.Name = name
	var noPasswd bool = true
	this.NoPasswd = *NewNullableBool(&noPasswd)
	return &this
}

// NewSudoCommandBundleWithDefaults instantiates a new SudoCommandBundle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSudoCommandBundleWithDefaults() *SudoCommandBundle {
	this := SudoCommandBundle{}
	var noPasswd bool = true
	this.NoPasswd = *NewNullableBool(&noPasswd)
	return &this
}

// GetAddEnv returns the AddEnv field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SudoCommandBundle) GetAddEnv() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AddEnv
}

// GetAddEnvOk returns a tuple with the AddEnv field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SudoCommandBundle) GetAddEnvOk() ([]string, bool) {
	if o == nil || IsNil(o.AddEnv) {
		return nil, false
	}
	return o.AddEnv, true
}

// HasAddEnv returns a boolean if a field has been set.
func (o *SudoCommandBundle) HasAddEnv() bool {
	if o != nil && IsNil(o.AddEnv) {
		return true
	}

	return false
}

// SetAddEnv gets a reference to the given []string and assigns it to the AddEnv field.
func (o *SudoCommandBundle) SetAddEnv(v []string) *SudoCommandBundle {
	o.AddEnv = v
	return o
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SudoCommandBundle) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SudoCommandBundle) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SudoCommandBundle) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SudoCommandBundle) SetDescription(v string) *SudoCommandBundle {
	o.Description = &v
	return o
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SudoCommandBundle) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SudoCommandBundle) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SudoCommandBundle) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SudoCommandBundle) SetId(v string) *SudoCommandBundle {
	o.Id = &v
	return o
}

// GetName returns the Name field value
func (o *SudoCommandBundle) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SudoCommandBundle) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SudoCommandBundle) SetName(v string) *SudoCommandBundle {
	o.Name = v
	return o
}

// GetNoExec returns the NoExec field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SudoCommandBundle) GetNoExec() bool {
	if o == nil || IsNil(o.NoExec.Get()) {
		var ret bool
		return ret
	}
	return *o.NoExec.Get()
}

// GetNoExecOk returns a tuple with the NoExec field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SudoCommandBundle) GetNoExecOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NoExec.Get(), o.NoExec.IsSet()
}

// HasNoExec returns a boolean if a field has been set.
func (o *SudoCommandBundle) HasNoExec() bool {
	if o != nil && o.NoExec.IsSet() {
		return true
	}

	return false
}

// SetNoExec gets a reference to the given NullableBool and assigns it to the NoExec field.
func (o *SudoCommandBundle) SetNoExec(v bool) *SudoCommandBundle {
	o.NoExec.Set(&v)
	return o
}

// SetNoExecNil sets the value for NoExec to be an explicit nil
func (o *SudoCommandBundle) SetNoExecNil() *SudoCommandBundle {
	o.NoExec.Set(nil)
	return o
}

// UnsetNoExec ensures that no value is present for NoExec, not even an explicit nil
func (o *SudoCommandBundle) UnsetNoExec() *SudoCommandBundle {
	o.NoExec.Unset()
	return o
}

// GetNoPasswd returns the NoPasswd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SudoCommandBundle) GetNoPasswd() bool {
	if o == nil || IsNil(o.NoPasswd.Get()) {
		var ret bool
		return ret
	}
	return *o.NoPasswd.Get()
}

// GetNoPasswdOk returns a tuple with the NoPasswd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SudoCommandBundle) GetNoPasswdOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NoPasswd.Get(), o.NoPasswd.IsSet()
}

// HasNoPasswd returns a boolean if a field has been set.
func (o *SudoCommandBundle) HasNoPasswd() bool {
	if o != nil && o.NoPasswd.IsSet() {
		return true
	}

	return false
}

// SetNoPasswd gets a reference to the given NullableBool and assigns it to the NoPasswd field.
func (o *SudoCommandBundle) SetNoPasswd(v bool) *SudoCommandBundle {
	o.NoPasswd.Set(&v)
	return o
}

// SetNoPasswdNil sets the value for NoPasswd to be an explicit nil
func (o *SudoCommandBundle) SetNoPasswdNil() *SudoCommandBundle {
	o.NoPasswd.Set(nil)
	return o
}

// UnsetNoPasswd ensures that no value is present for NoPasswd, not even an explicit nil
func (o *SudoCommandBundle) UnsetNoPasswd() *SudoCommandBundle {
	o.NoPasswd.Unset()
	return o
}

// GetRunAs returns the RunAs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SudoCommandBundle) GetRunAs() string {
	if o == nil || IsNil(o.RunAs.Get()) {
		var ret string
		return ret
	}
	return *o.RunAs.Get()
}

// GetRunAsOk returns a tuple with the RunAs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SudoCommandBundle) GetRunAsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RunAs.Get(), o.RunAs.IsSet()
}

// HasRunAs returns a boolean if a field has been set.
func (o *SudoCommandBundle) HasRunAs() bool {
	if o != nil && o.RunAs.IsSet() {
		return true
	}

	return false
}

// SetRunAs gets a reference to the given NullableString and assigns it to the RunAs field.
func (o *SudoCommandBundle) SetRunAs(v string) *SudoCommandBundle {
	o.RunAs.Set(&v)
	return o
}

// SetRunAsNil sets the value for RunAs to be an explicit nil
func (o *SudoCommandBundle) SetRunAsNil() *SudoCommandBundle {
	o.RunAs.Set(nil)
	return o
}

// UnsetRunAs ensures that no value is present for RunAs, not even an explicit nil
func (o *SudoCommandBundle) UnsetRunAs() *SudoCommandBundle {
	o.RunAs.Unset()
	return o
}

// GetSetEnv returns the SetEnv field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SudoCommandBundle) GetSetEnv() bool {
	if o == nil || IsNil(o.SetEnv.Get()) {
		var ret bool
		return ret
	}
	return *o.SetEnv.Get()
}

// GetSetEnvOk returns a tuple with the SetEnv field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SudoCommandBundle) GetSetEnvOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SetEnv.Get(), o.SetEnv.IsSet()
}

// HasSetEnv returns a boolean if a field has been set.
func (o *SudoCommandBundle) HasSetEnv() bool {
	if o != nil && o.SetEnv.IsSet() {
		return true
	}

	return false
}

// SetSetEnv gets a reference to the given NullableBool and assigns it to the SetEnv field.
func (o *SudoCommandBundle) SetSetEnv(v bool) *SudoCommandBundle {
	o.SetEnv.Set(&v)
	return o
}

// SetSetEnvNil sets the value for SetEnv to be an explicit nil
func (o *SudoCommandBundle) SetSetEnvNil() *SudoCommandBundle {
	o.SetEnv.Set(nil)
	return o
}

// UnsetSetEnv ensures that no value is present for SetEnv, not even an explicit nil
func (o *SudoCommandBundle) UnsetSetEnv() *SudoCommandBundle {
	o.SetEnv.Unset()
	return o
}

// GetStructuredCommands returns the StructuredCommands field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SudoCommandBundle) GetStructuredCommands() []SudoCommandBundleStructuredCommandsInner {
	if o == nil {
		var ret []SudoCommandBundleStructuredCommandsInner
		return ret
	}
	return o.StructuredCommands
}

// GetStructuredCommandsOk returns a tuple with the StructuredCommands field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SudoCommandBundle) GetStructuredCommandsOk() ([]SudoCommandBundleStructuredCommandsInner, bool) {
	if o == nil || IsNil(o.StructuredCommands) {
		return nil, false
	}
	return o.StructuredCommands, true
}

// HasStructuredCommands returns a boolean if a field has been set.
func (o *SudoCommandBundle) HasStructuredCommands() bool {
	if o != nil && IsNil(o.StructuredCommands) {
		return true
	}

	return false
}

// SetStructuredCommands gets a reference to the given []SudoCommandBundleStructuredCommandsInner and assigns it to the StructuredCommands field.
func (o *SudoCommandBundle) SetStructuredCommands(v []SudoCommandBundleStructuredCommandsInner) *SudoCommandBundle {
	o.StructuredCommands = v
	return o
}

// GetSubEnv returns the SubEnv field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SudoCommandBundle) GetSubEnv() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SubEnv
}

// GetSubEnvOk returns a tuple with the SubEnv field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SudoCommandBundle) GetSubEnvOk() ([]string, bool) {
	if o == nil || IsNil(o.SubEnv) {
		return nil, false
	}
	return o.SubEnv, true
}

// HasSubEnv returns a boolean if a field has been set.
func (o *SudoCommandBundle) HasSubEnv() bool {
	if o != nil && IsNil(o.SubEnv) {
		return true
	}

	return false
}

// SetSubEnv gets a reference to the given []string and assigns it to the SubEnv field.
func (o *SudoCommandBundle) SetSubEnv(v []string) *SudoCommandBundle {
	o.SubEnv = v
	return o
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SudoCommandBundle) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SudoCommandBundle) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SudoCommandBundle) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *SudoCommandBundle) SetCreatedAt(v time.Time) *SudoCommandBundle {
	o.CreatedAt = &v
	return o
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *SudoCommandBundle) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SudoCommandBundle) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *SudoCommandBundle) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *SudoCommandBundle) SetCreatedBy(v string) *SudoCommandBundle {
	o.CreatedBy = &v
	return o
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *SudoCommandBundle) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SudoCommandBundle) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *SudoCommandBundle) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *SudoCommandBundle) SetUpdatedAt(v time.Time) *SudoCommandBundle {
	o.UpdatedAt = &v
	return o
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *SudoCommandBundle) GetUpdatedBy() string {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SudoCommandBundle) GetUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *SudoCommandBundle) HasUpdatedBy() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *SudoCommandBundle) SetUpdatedBy(v string) *SudoCommandBundle {
	o.UpdatedBy = &v
	return o
}

func (o SudoCommandBundle) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SudoCommandBundle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AddEnv != nil {
		toSerialize["add_env"] = o.AddEnv
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if o.NoExec.IsSet() {
		toSerialize["no_exec"] = o.NoExec.Get()
	}
	if o.NoPasswd.IsSet() {
		toSerialize["no_passwd"] = o.NoPasswd.Get()
	}
	if o.RunAs.IsSet() {
		toSerialize["run_as"] = o.RunAs.Get()
	}
	if o.SetEnv.IsSet() {
		toSerialize["set_env"] = o.SetEnv.Get()
	}
	if o.StructuredCommands != nil {
		toSerialize["structured_commands"] = o.StructuredCommands
	}
	if o.SubEnv != nil {
		toSerialize["sub_env"] = o.SubEnv
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["created_by"] = o.CreatedBy
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.UpdatedBy) {
		toSerialize["updated_by"] = o.UpdatedBy
	}
	return toSerialize, nil
}

type NullableSudoCommandBundle struct {
	value *SudoCommandBundle
	isSet bool
}

func (v NullableSudoCommandBundle) Get() *SudoCommandBundle {
	return v.value
}

func (v *NullableSudoCommandBundle) Set(val *SudoCommandBundle) {
	v.value = val
	v.isSet = true
}

func (v NullableSudoCommandBundle) IsSet() bool {
	return v.isSet
}

func (v *NullableSudoCommandBundle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSudoCommandBundle(val *SudoCommandBundle) *NullableSudoCommandBundle {
	return &NullableSudoCommandBundle{value: val, isSet: true}
}

func (v NullableSudoCommandBundle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSudoCommandBundle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
