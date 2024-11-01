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

// checks if the ServiceAccountsUpdateCredentialsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceAccountsUpdateCredentialsRequest{}

// ServiceAccountsUpdateCredentialsRequest struct for ServiceAccountsUpdateCredentialsRequest
type ServiceAccountsUpdateCredentialsRequest struct {
	PasswordJwe      string           `json:"password_jwe"`
	UserAccessMethod UserAccessMethod `json:"user_access_method"`
}

// NewServiceAccountsUpdateCredentialsRequest instantiates a new ServiceAccountsUpdateCredentialsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountsUpdateCredentialsRequest(passwordJwe string, userAccessMethod UserAccessMethod) *ServiceAccountsUpdateCredentialsRequest {
	this := ServiceAccountsUpdateCredentialsRequest{}
	this.PasswordJwe = passwordJwe
	this.UserAccessMethod = userAccessMethod
	return &this
}

// NewServiceAccountsUpdateCredentialsRequestWithDefaults instantiates a new ServiceAccountsUpdateCredentialsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountsUpdateCredentialsRequestWithDefaults() *ServiceAccountsUpdateCredentialsRequest {
	this := ServiceAccountsUpdateCredentialsRequest{}
	return &this
}

// GetPasswordJwe returns the PasswordJwe field value
func (o *ServiceAccountsUpdateCredentialsRequest) GetPasswordJwe() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PasswordJwe
}

// GetPasswordJweOk returns a tuple with the PasswordJwe field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountsUpdateCredentialsRequest) GetPasswordJweOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PasswordJwe, true
}

// SetPasswordJwe sets field value
func (o *ServiceAccountsUpdateCredentialsRequest) SetPasswordJwe(v string) *ServiceAccountsUpdateCredentialsRequest {
	o.PasswordJwe = v
	return o
}

// GetUserAccessMethod returns the UserAccessMethod field value
func (o *ServiceAccountsUpdateCredentialsRequest) GetUserAccessMethod() UserAccessMethod {
	if o == nil {
		var ret UserAccessMethod
		return ret
	}

	return o.UserAccessMethod
}

// GetUserAccessMethodOk returns a tuple with the UserAccessMethod field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountsUpdateCredentialsRequest) GetUserAccessMethodOk() (*UserAccessMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserAccessMethod, true
}

// SetUserAccessMethod sets field value
func (o *ServiceAccountsUpdateCredentialsRequest) SetUserAccessMethod(v UserAccessMethod) *ServiceAccountsUpdateCredentialsRequest {
	o.UserAccessMethod = v
	return o
}

func (o ServiceAccountsUpdateCredentialsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceAccountsUpdateCredentialsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["password_jwe"] = o.PasswordJwe
	toSerialize["user_access_method"] = o.UserAccessMethod
	return toSerialize, nil
}

type NullableServiceAccountsUpdateCredentialsRequest struct {
	value *ServiceAccountsUpdateCredentialsRequest
	isSet bool
}

func (v NullableServiceAccountsUpdateCredentialsRequest) Get() *ServiceAccountsUpdateCredentialsRequest {
	return v.value
}

func (v *NullableServiceAccountsUpdateCredentialsRequest) Set(val *ServiceAccountsUpdateCredentialsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccountsUpdateCredentialsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccountsUpdateCredentialsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccountsUpdateCredentialsRequest(val *ServiceAccountsUpdateCredentialsRequest) *NullableServiceAccountsUpdateCredentialsRequest {
	return &NullableServiceAccountsUpdateCredentialsRequest{value: val, isSet: true}
}

func (v NullableServiceAccountsUpdateCredentialsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccountsUpdateCredentialsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
