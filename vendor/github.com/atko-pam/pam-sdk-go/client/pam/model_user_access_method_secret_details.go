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

// checks if the UserAccessMethodSecretDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserAccessMethodSecretDetails{}

// UserAccessMethodSecretDetails struct for UserAccessMethodSecretDetails
type UserAccessMethodSecretDetails struct {
	// The name of the secret used to access the resource
	SecretName *string `json:"secret_name,omitempty"`
	// The ID of the secret used to access the resource
	SecretId   *string          `json:"secret_id,omitempty"`
	Path       []SecretPath     `json:"path,omitempty"`
	Privileges *SecretPrivilege `json:"privileges,omitempty"`
}

// NewUserAccessMethodSecretDetails instantiates a new UserAccessMethodSecretDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAccessMethodSecretDetails() *UserAccessMethodSecretDetails {
	this := UserAccessMethodSecretDetails{}
	return &this
}

// NewUserAccessMethodSecretDetailsWithDefaults instantiates a new UserAccessMethodSecretDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAccessMethodSecretDetailsWithDefaults() *UserAccessMethodSecretDetails {
	this := UserAccessMethodSecretDetails{}
	return &this
}

// GetSecretName returns the SecretName field value if set, zero value otherwise.
func (o *UserAccessMethodSecretDetails) GetSecretName() string {
	if o == nil || IsNil(o.SecretName) {
		var ret string
		return ret
	}
	return *o.SecretName
}

// GetSecretNameOk returns a tuple with the SecretName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAccessMethodSecretDetails) GetSecretNameOk() (*string, bool) {
	if o == nil || IsNil(o.SecretName) {
		return nil, false
	}
	return o.SecretName, true
}

// HasSecretName returns a boolean if a field has been set.
func (o *UserAccessMethodSecretDetails) HasSecretName() bool {
	if o != nil && !IsNil(o.SecretName) {
		return true
	}

	return false
}

// SetSecretName gets a reference to the given string and assigns it to the SecretName field.
func (o *UserAccessMethodSecretDetails) SetSecretName(v string) *UserAccessMethodSecretDetails {
	o.SecretName = &v
	return o
}

// GetSecretId returns the SecretId field value if set, zero value otherwise.
func (o *UserAccessMethodSecretDetails) GetSecretId() string {
	if o == nil || IsNil(o.SecretId) {
		var ret string
		return ret
	}
	return *o.SecretId
}

// GetSecretIdOk returns a tuple with the SecretId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAccessMethodSecretDetails) GetSecretIdOk() (*string, bool) {
	if o == nil || IsNil(o.SecretId) {
		return nil, false
	}
	return o.SecretId, true
}

// HasSecretId returns a boolean if a field has been set.
func (o *UserAccessMethodSecretDetails) HasSecretId() bool {
	if o != nil && !IsNil(o.SecretId) {
		return true
	}

	return false
}

// SetSecretId gets a reference to the given string and assigns it to the SecretId field.
func (o *UserAccessMethodSecretDetails) SetSecretId(v string) *UserAccessMethodSecretDetails {
	o.SecretId = &v
	return o
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *UserAccessMethodSecretDetails) GetPath() []SecretPath {
	if o == nil || IsNil(o.Path) {
		var ret []SecretPath
		return ret
	}
	return o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAccessMethodSecretDetails) GetPathOk() ([]SecretPath, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *UserAccessMethodSecretDetails) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given []SecretPath and assigns it to the Path field.
func (o *UserAccessMethodSecretDetails) SetPath(v []SecretPath) *UserAccessMethodSecretDetails {
	o.Path = v
	return o
}

// GetPrivileges returns the Privileges field value if set, zero value otherwise.
func (o *UserAccessMethodSecretDetails) GetPrivileges() SecretPrivilege {
	if o == nil || IsNil(o.Privileges) {
		var ret SecretPrivilege
		return ret
	}
	return *o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAccessMethodSecretDetails) GetPrivilegesOk() (*SecretPrivilege, bool) {
	if o == nil || IsNil(o.Privileges) {
		return nil, false
	}
	return o.Privileges, true
}

// HasPrivileges returns a boolean if a field has been set.
func (o *UserAccessMethodSecretDetails) HasPrivileges() bool {
	if o != nil && !IsNil(o.Privileges) {
		return true
	}

	return false
}

// SetPrivileges gets a reference to the given SecretPrivilege and assigns it to the Privileges field.
func (o *UserAccessMethodSecretDetails) SetPrivileges(v SecretPrivilege) *UserAccessMethodSecretDetails {
	o.Privileges = &v
	return o
}

func (o UserAccessMethodSecretDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserAccessMethodSecretDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SecretName) {
		toSerialize["secret_name"] = o.SecretName
	}
	if !IsNil(o.SecretId) {
		toSerialize["secret_id"] = o.SecretId
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Privileges) {
		toSerialize["privileges"] = o.Privileges
	}
	return toSerialize, nil
}

type NullableUserAccessMethodSecretDetails struct {
	value *UserAccessMethodSecretDetails
	isSet bool
}

func (v NullableUserAccessMethodSecretDetails) Get() *UserAccessMethodSecretDetails {
	return v.value
}

func (v *NullableUserAccessMethodSecretDetails) Set(val *UserAccessMethodSecretDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAccessMethodSecretDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAccessMethodSecretDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAccessMethodSecretDetails(val *UserAccessMethodSecretDetails) *NullableUserAccessMethodSecretDetails {
	return &NullableUserAccessMethodSecretDetails{value: val, isSet: true}
}

func (v NullableUserAccessMethodSecretDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAccessMethodSecretDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
