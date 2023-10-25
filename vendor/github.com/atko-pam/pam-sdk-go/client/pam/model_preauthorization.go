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
	"time"
)

// checks if the Preauthorization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Preauthorization{}

// Preauthorization struct for Preauthorization
type Preauthorization struct {
	// Indicates the status of the Preauthorization. If `true`, the Preauthorization is disabled.
	Disabled bool `json:"disabled"`
	// A timestamp indicating when the Preauthorization expires
	ExpiresAt time.Time `json:"expires_at"`
	// The UUID of the Preauthorization
	Id string `json:"id"`
	// The Project associated with the Preauthorization
	Projects []string `json:"projects"`
	// The Servers accessed using the Preauthorization
	Servers []string `json:"servers,omitempty"`
	// The User associated with the Preauthorization
	UserName string `json:"user_name"`
}

// NewPreauthorization instantiates a new Preauthorization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreauthorization(disabled bool, expiresAt time.Time, id string, projects []string, userName string) *Preauthorization {
	this := Preauthorization{}
	this.Disabled = disabled
	this.ExpiresAt = expiresAt
	this.Id = id
	this.Projects = projects
	this.UserName = userName
	return &this
}

// NewPreauthorizationWithDefaults instantiates a new Preauthorization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreauthorizationWithDefaults() *Preauthorization {
	this := Preauthorization{}
	return &this
}

// GetDisabled returns the Disabled field value
func (o *Preauthorization) GetDisabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value
// and a boolean to check if the value has been set.
func (o *Preauthorization) GetDisabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Disabled, true
}

// SetDisabled sets field value
func (o *Preauthorization) SetDisabled(v bool) *Preauthorization {
	o.Disabled = v
	return o
}

// GetExpiresAt returns the ExpiresAt field value
func (o *Preauthorization) GetExpiresAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *Preauthorization) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *Preauthorization) SetExpiresAt(v time.Time) *Preauthorization {
	o.ExpiresAt = v
	return o
}

// GetId returns the Id field value
func (o *Preauthorization) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Preauthorization) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Preauthorization) SetId(v string) *Preauthorization {
	o.Id = v
	return o
}

// GetProjects returns the Projects field value
func (o *Preauthorization) GetProjects() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value
// and a boolean to check if the value has been set.
func (o *Preauthorization) GetProjectsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Projects, true
}

// SetProjects sets field value
func (o *Preauthorization) SetProjects(v []string) *Preauthorization {
	o.Projects = v
	return o
}

// GetServers returns the Servers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Preauthorization) GetServers() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Preauthorization) GetServersOk() ([]string, bool) {
	if o == nil || IsNil(o.Servers) {
		return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *Preauthorization) HasServers() bool {
	if o != nil && IsNil(o.Servers) {
		return true
	}

	return false
}

// SetServers gets a reference to the given []string and assigns it to the Servers field.
func (o *Preauthorization) SetServers(v []string) *Preauthorization {
	o.Servers = v
	return o
}

// GetUserName returns the UserName field value
func (o *Preauthorization) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *Preauthorization) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value
func (o *Preauthorization) SetUserName(v string) *Preauthorization {
	o.UserName = v
	return o
}

func (o Preauthorization) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Preauthorization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["disabled"] = o.Disabled
	toSerialize["expires_at"] = o.ExpiresAt
	toSerialize["id"] = o.Id
	toSerialize["projects"] = o.Projects
	if o.Servers != nil {
		toSerialize["servers"] = o.Servers
	}
	toSerialize["user_name"] = o.UserName
	return toSerialize, nil
}

type NullablePreauthorization struct {
	value *Preauthorization
	isSet bool
}

func (v NullablePreauthorization) Get() *Preauthorization {
	return v.value
}

func (v *NullablePreauthorization) Set(val *Preauthorization) {
	v.value = val
	v.isSet = true
}

func (v NullablePreauthorization) IsSet() bool {
	return v.isSet
}

func (v *NullablePreauthorization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreauthorization(val *Preauthorization) *NullablePreauthorization {
	return &NullablePreauthorization{value: val, isSet: true}
}

func (v NullablePreauthorization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreauthorization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
