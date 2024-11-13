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

// checks if the ServiceAccountsRevealCredentialsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceAccountsRevealCredentialsResponse{}

// ServiceAccountsRevealCredentialsResponse struct for ServiceAccountsRevealCredentialsResponse
type ServiceAccountsRevealCredentialsResponse struct {
	PasswordJwe string `json:"password_jwe"`
}

// NewServiceAccountsRevealCredentialsResponse instantiates a new ServiceAccountsRevealCredentialsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountsRevealCredentialsResponse(passwordJwe string) *ServiceAccountsRevealCredentialsResponse {
	this := ServiceAccountsRevealCredentialsResponse{}
	this.PasswordJwe = passwordJwe
	return &this
}

// NewServiceAccountsRevealCredentialsResponseWithDefaults instantiates a new ServiceAccountsRevealCredentialsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountsRevealCredentialsResponseWithDefaults() *ServiceAccountsRevealCredentialsResponse {
	this := ServiceAccountsRevealCredentialsResponse{}
	return &this
}

// GetPasswordJwe returns the PasswordJwe field value
func (o *ServiceAccountsRevealCredentialsResponse) GetPasswordJwe() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PasswordJwe
}

// GetPasswordJweOk returns a tuple with the PasswordJwe field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountsRevealCredentialsResponse) GetPasswordJweOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PasswordJwe, true
}

// SetPasswordJwe sets field value
func (o *ServiceAccountsRevealCredentialsResponse) SetPasswordJwe(v string) *ServiceAccountsRevealCredentialsResponse {
	o.PasswordJwe = v
	return o
}

func (o ServiceAccountsRevealCredentialsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceAccountsRevealCredentialsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["password_jwe"] = o.PasswordJwe
	return toSerialize, nil
}

type NullableServiceAccountsRevealCredentialsResponse struct {
	value *ServiceAccountsRevealCredentialsResponse
	isSet bool
}

func (v NullableServiceAccountsRevealCredentialsResponse) Get() *ServiceAccountsRevealCredentialsResponse {
	return v.value
}

func (v *NullableServiceAccountsRevealCredentialsResponse) Set(val *ServiceAccountsRevealCredentialsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccountsRevealCredentialsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccountsRevealCredentialsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccountsRevealCredentialsResponse(val *ServiceAccountsRevealCredentialsResponse) *NullableServiceAccountsRevealCredentialsResponse {
	return &NullableServiceAccountsRevealCredentialsResponse{value: val, isSet: true}
}

func (v NullableServiceAccountsRevealCredentialsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccountsRevealCredentialsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}