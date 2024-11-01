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

// checks if the PrivilegedAccountDetailsOktaUserAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrivilegedAccountDetailsOktaUserAccount{}

// PrivilegedAccountDetailsOktaUserAccount Details for managing an Okta Universal Directory Account as a Privileged Account
type PrivilegedAccountDetailsOktaUserAccount struct {
	PrivilegedAccount
	Details PrivilegedAccountDetailsOktaUserAccountAllOfDetails `json:"details"`
}

// NewPrivilegedAccountDetailsOktaUserAccount instantiates a new PrivilegedAccountDetailsOktaUserAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivilegedAccountDetailsOktaUserAccount(details PrivilegedAccountDetailsOktaUserAccountAllOfDetails, name string, description string, accountType PrivilegedAccountType, ownerUserIds []string, ownerGroupIds []string) *PrivilegedAccountDetailsOktaUserAccount {
	this := PrivilegedAccountDetailsOktaUserAccount{}
	this.Name = name
	this.Description = description
	this.AccountType = accountType
	this.OwnerUserIds = ownerUserIds
	this.OwnerGroupIds = ownerGroupIds
	this.Details = details
	return &this
}

// NewPrivilegedAccountDetailsOktaUserAccountWithDefaults instantiates a new PrivilegedAccountDetailsOktaUserAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivilegedAccountDetailsOktaUserAccountWithDefaults() *PrivilegedAccountDetailsOktaUserAccount {
	this := PrivilegedAccountDetailsOktaUserAccount{}
	return &this
}

// GetDetails returns the Details field value
func (o *PrivilegedAccountDetailsOktaUserAccount) GetDetails() PrivilegedAccountDetailsOktaUserAccountAllOfDetails {
	if o == nil {
		var ret PrivilegedAccountDetailsOktaUserAccountAllOfDetails
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *PrivilegedAccountDetailsOktaUserAccount) GetDetailsOk() (*PrivilegedAccountDetailsOktaUserAccountAllOfDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *PrivilegedAccountDetailsOktaUserAccount) SetDetails(v PrivilegedAccountDetailsOktaUserAccountAllOfDetails) *PrivilegedAccountDetailsOktaUserAccount {
	o.Details = v
	return o
}

func (o PrivilegedAccountDetailsOktaUserAccount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrivilegedAccountDetailsOktaUserAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPrivilegedAccount, errPrivilegedAccount := json.Marshal(o.PrivilegedAccount)
	if errPrivilegedAccount != nil {
		return map[string]interface{}{}, errPrivilegedAccount
	}
	errPrivilegedAccount = json.Unmarshal([]byte(serializedPrivilegedAccount), &toSerialize)
	if errPrivilegedAccount != nil {
		return map[string]interface{}{}, errPrivilegedAccount
	}
	toSerialize["details"] = o.Details
	return toSerialize, nil
}

type NullablePrivilegedAccountDetailsOktaUserAccount struct {
	value *PrivilegedAccountDetailsOktaUserAccount
	isSet bool
}

func (v NullablePrivilegedAccountDetailsOktaUserAccount) Get() *PrivilegedAccountDetailsOktaUserAccount {
	return v.value
}

func (v *NullablePrivilegedAccountDetailsOktaUserAccount) Set(val *PrivilegedAccountDetailsOktaUserAccount) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivilegedAccountDetailsOktaUserAccount) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivilegedAccountDetailsOktaUserAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivilegedAccountDetailsOktaUserAccount(val *PrivilegedAccountDetailsOktaUserAccount) *NullablePrivilegedAccountDetailsOktaUserAccount {
	return &NullablePrivilegedAccountDetailsOktaUserAccount{value: val, isSet: true}
}

func (v NullablePrivilegedAccountDetailsOktaUserAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivilegedAccountDetailsOktaUserAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}