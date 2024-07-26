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
)

// checks if the PrivilegedAccountCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrivilegedAccountCredentials{}

// PrivilegedAccountCredentials Credentials for a Privileged Account
type PrivilegedAccountCredentials struct {
	// The password associated with the Privileged Account
	Password *string `json:"password,omitempty"`
	// The username associated with the Privileged Account
	Username string `json:"username"`
}

// NewPrivilegedAccountCredentials instantiates a new PrivilegedAccountCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivilegedAccountCredentials(username string) *PrivilegedAccountCredentials {
	this := PrivilegedAccountCredentials{}
	this.Username = username
	return &this
}

// NewPrivilegedAccountCredentialsWithDefaults instantiates a new PrivilegedAccountCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivilegedAccountCredentialsWithDefaults() *PrivilegedAccountCredentials {
	this := PrivilegedAccountCredentials{}
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *PrivilegedAccountCredentials) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivilegedAccountCredentials) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *PrivilegedAccountCredentials) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *PrivilegedAccountCredentials) SetPassword(v string) *PrivilegedAccountCredentials {
	o.Password = &v
	return o
}

// GetUsername returns the Username field value
func (o *PrivilegedAccountCredentials) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *PrivilegedAccountCredentials) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *PrivilegedAccountCredentials) SetUsername(v string) *PrivilegedAccountCredentials {
	o.Username = v
	return o
}

func (o PrivilegedAccountCredentials) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrivilegedAccountCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	toSerialize["username"] = o.Username
	return toSerialize, nil
}

type NullablePrivilegedAccountCredentials struct {
	value *PrivilegedAccountCredentials
	isSet bool
}

func (v NullablePrivilegedAccountCredentials) Get() *PrivilegedAccountCredentials {
	return v.value
}

func (v *NullablePrivilegedAccountCredentials) Set(val *PrivilegedAccountCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivilegedAccountCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivilegedAccountCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivilegedAccountCredentials(val *PrivilegedAccountCredentials) *NullablePrivilegedAccountCredentials {
	return &NullablePrivilegedAccountCredentials{value: val, isSet: true}
}

func (v NullablePrivilegedAccountCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivilegedAccountCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
