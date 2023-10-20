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

// checks if the NameResolutionResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NameResolutionResult{}

// NameResolutionResult struct for NameResolutionResult
type NameResolutionResult struct {
	// Name of the resource
	Name *string `json:"name,omitempty"`
	// A list of servers and bastions used to reach the target resource
	Servers []Server `json:"servers,omitempty"`
	// A list of user access methods used to access the resource
	UserAccessMethods []UserAccessMethod `json:"user_access_methods,omitempty"`
}

// NewNameResolutionResult instantiates a new NameResolutionResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNameResolutionResult() *NameResolutionResult {
	this := NameResolutionResult{}
	return &this
}

// NewNameResolutionResultWithDefaults instantiates a new NameResolutionResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNameResolutionResultWithDefaults() *NameResolutionResult {
	this := NameResolutionResult{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NameResolutionResult) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NameResolutionResult) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NameResolutionResult) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NameResolutionResult) SetName(v string) *NameResolutionResult {
	o.Name = &v
	return o
}

// GetServers returns the Servers field value if set, zero value otherwise.
func (o *NameResolutionResult) GetServers() []Server {
	if o == nil || IsNil(o.Servers) {
		var ret []Server
		return ret
	}
	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NameResolutionResult) GetServersOk() ([]Server, bool) {
	if o == nil || IsNil(o.Servers) {
		return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *NameResolutionResult) HasServers() bool {
	if o != nil && !IsNil(o.Servers) {
		return true
	}

	return false
}

// SetServers gets a reference to the given []Server and assigns it to the Servers field.
func (o *NameResolutionResult) SetServers(v []Server) *NameResolutionResult {
	o.Servers = v
	return o
}

// GetUserAccessMethods returns the UserAccessMethods field value if set, zero value otherwise.
func (o *NameResolutionResult) GetUserAccessMethods() []UserAccessMethod {
	if o == nil || IsNil(o.UserAccessMethods) {
		var ret []UserAccessMethod
		return ret
	}
	return o.UserAccessMethods
}

// GetUserAccessMethodsOk returns a tuple with the UserAccessMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NameResolutionResult) GetUserAccessMethodsOk() ([]UserAccessMethod, bool) {
	if o == nil || IsNil(o.UserAccessMethods) {
		return nil, false
	}
	return o.UserAccessMethods, true
}

// HasUserAccessMethods returns a boolean if a field has been set.
func (o *NameResolutionResult) HasUserAccessMethods() bool {
	if o != nil && !IsNil(o.UserAccessMethods) {
		return true
	}

	return false
}

// SetUserAccessMethods gets a reference to the given []UserAccessMethod and assigns it to the UserAccessMethods field.
func (o *NameResolutionResult) SetUserAccessMethods(v []UserAccessMethod) *NameResolutionResult {
	o.UserAccessMethods = v
	return o
}

func (o NameResolutionResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NameResolutionResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Servers) {
		toSerialize["servers"] = o.Servers
	}
	if !IsNil(o.UserAccessMethods) {
		toSerialize["user_access_methods"] = o.UserAccessMethods
	}
	return toSerialize, nil
}

type NullableNameResolutionResult struct {
	value *NameResolutionResult
	isSet bool
}

func (v NullableNameResolutionResult) Get() *NameResolutionResult {
	return v.value
}

func (v *NullableNameResolutionResult) Set(val *NameResolutionResult) {
	v.value = val
	v.isSet = true
}

func (v NullableNameResolutionResult) IsSet() bool {
	return v.isSet
}

func (v *NullableNameResolutionResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNameResolutionResult(val *NameResolutionResult) *NullableNameResolutionResult {
	return &NullableNameResolutionResult{value: val, isSet: true}
}

func (v NullableNameResolutionResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNameResolutionResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}