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

// checks if the PrivilegedAccountDetailsAppAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrivilegedAccountDetailsAppAccount{}

// PrivilegedAccountDetailsAppAccount Details for a SaaS Application Account which will be managed as a Privileged Account
type PrivilegedAccountDetailsAppAccount struct {
	PrivilegedAccount
	Details PrivilegedAccountDetailsAppAccountAllOfDetails `json:"details"`
}

// NewPrivilegedAccountDetailsAppAccount instantiates a new PrivilegedAccountDetailsAppAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivilegedAccountDetailsAppAccount(details PrivilegedAccountDetailsAppAccountAllOfDetails, name string, description string, accountType PrivilegedAccountType, ownerUserIds []string, ownerGroupIds []string) *PrivilegedAccountDetailsAppAccount {
	this := PrivilegedAccountDetailsAppAccount{}
	this.Name = name
	this.Description = description
	this.AccountType = accountType
	this.OwnerUserIds = ownerUserIds
	this.OwnerGroupIds = ownerGroupIds
	this.Details = details
	return &this
}

// NewPrivilegedAccountDetailsAppAccountWithDefaults instantiates a new PrivilegedAccountDetailsAppAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivilegedAccountDetailsAppAccountWithDefaults() *PrivilegedAccountDetailsAppAccount {
	this := PrivilegedAccountDetailsAppAccount{}
	return &this
}

// GetDetails returns the Details field value
func (o *PrivilegedAccountDetailsAppAccount) GetDetails() PrivilegedAccountDetailsAppAccountAllOfDetails {
	if o == nil {
		var ret PrivilegedAccountDetailsAppAccountAllOfDetails
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *PrivilegedAccountDetailsAppAccount) GetDetailsOk() (*PrivilegedAccountDetailsAppAccountAllOfDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *PrivilegedAccountDetailsAppAccount) SetDetails(v PrivilegedAccountDetailsAppAccountAllOfDetails) *PrivilegedAccountDetailsAppAccount {
	o.Details = v
	return o
}

func (o PrivilegedAccountDetailsAppAccount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrivilegedAccountDetailsAppAccount) ToMap() (map[string]interface{}, error) {
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

type NullablePrivilegedAccountDetailsAppAccount struct {
	value *PrivilegedAccountDetailsAppAccount
	isSet bool
}

func (v NullablePrivilegedAccountDetailsAppAccount) Get() *PrivilegedAccountDetailsAppAccount {
	return v.value
}

func (v *NullablePrivilegedAccountDetailsAppAccount) Set(val *PrivilegedAccountDetailsAppAccount) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivilegedAccountDetailsAppAccount) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivilegedAccountDetailsAppAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivilegedAccountDetailsAppAccount(val *PrivilegedAccountDetailsAppAccount) *NullablePrivilegedAccountDetailsAppAccount {
	return &NullablePrivilegedAccountDetailsAppAccount{value: val, isSet: true}
}

func (v NullablePrivilegedAccountDetailsAppAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivilegedAccountDetailsAppAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}