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

// checks if the ServiceAccountDetailsOktaUserAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceAccountDetailsOktaUserAccount{}

// ServiceAccountDetailsOktaUserAccount Details for managing an Okta user as a service account
type ServiceAccountDetailsOktaUserAccount struct {
	ServiceAccount
	Details ServiceAccountDetailsOktaUserAccountSub `json:"details"`
}

// NewServiceAccountDetailsOktaUserAccount instantiates a new ServiceAccountDetailsOktaUserAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountDetailsOktaUserAccount(details ServiceAccountDetailsOktaUserAccountSub, name string, accountType ServiceAccountType) *ServiceAccountDetailsOktaUserAccount {
	this := ServiceAccountDetailsOktaUserAccount{}
	this.Name = name
	this.AccountType = accountType
	this.Details = details
	return &this
}

// NewServiceAccountDetailsOktaUserAccountWithDefaults instantiates a new ServiceAccountDetailsOktaUserAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountDetailsOktaUserAccountWithDefaults() *ServiceAccountDetailsOktaUserAccount {
	this := ServiceAccountDetailsOktaUserAccount{}
	return &this
}

// GetDetails returns the Details field value
func (o *ServiceAccountDetailsOktaUserAccount) GetDetails() ServiceAccountDetailsOktaUserAccountSub {
	if o == nil {
		var ret ServiceAccountDetailsOktaUserAccountSub
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountDetailsOktaUserAccount) GetDetailsOk() (*ServiceAccountDetailsOktaUserAccountSub, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *ServiceAccountDetailsOktaUserAccount) SetDetails(v ServiceAccountDetailsOktaUserAccountSub) *ServiceAccountDetailsOktaUserAccount {
	o.Details = v
	return o
}

func (o ServiceAccountDetailsOktaUserAccount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceAccountDetailsOktaUserAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedServiceAccount, errServiceAccount := json.Marshal(o.ServiceAccount)
	if errServiceAccount != nil {
		return map[string]interface{}{}, errServiceAccount
	}
	errServiceAccount = json.Unmarshal([]byte(serializedServiceAccount), &toSerialize)
	if errServiceAccount != nil {
		return map[string]interface{}{}, errServiceAccount
	}
	toSerialize["details"] = o.Details
	return toSerialize, nil
}

type NullableServiceAccountDetailsOktaUserAccount struct {
	value *ServiceAccountDetailsOktaUserAccount
	isSet bool
}

func (v NullableServiceAccountDetailsOktaUserAccount) Get() *ServiceAccountDetailsOktaUserAccount {
	return v.value
}

func (v *NullableServiceAccountDetailsOktaUserAccount) Set(val *ServiceAccountDetailsOktaUserAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccountDetailsOktaUserAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccountDetailsOktaUserAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccountDetailsOktaUserAccount(val *ServiceAccountDetailsOktaUserAccount) *NullableServiceAccountDetailsOktaUserAccount {
	return &NullableServiceAccountDetailsOktaUserAccount{value: val, isSet: true}
}

func (v NullableServiceAccountDetailsOktaUserAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccountDetailsOktaUserAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
