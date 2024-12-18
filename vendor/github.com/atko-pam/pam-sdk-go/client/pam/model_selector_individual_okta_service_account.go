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

// checks if the SelectorIndividualOktaServiceAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SelectorIndividualOktaServiceAccount{}

// SelectorIndividualOktaServiceAccount This resource selector identifies a Service Account that exists on a Okta Universal Directory
type SelectorIndividualOktaServiceAccount struct {
	Type           string      `json:"_type"`
	ServiceAccount NamedObject `json:"service_account"`
}

// NewSelectorIndividualOktaServiceAccount instantiates a new SelectorIndividualOktaServiceAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelectorIndividualOktaServiceAccount(type_ string, serviceAccount NamedObject) *SelectorIndividualOktaServiceAccount {
	this := SelectorIndividualOktaServiceAccount{}
	this.Type = type_
	this.ServiceAccount = serviceAccount
	return &this
}

// NewSelectorIndividualOktaServiceAccountWithDefaults instantiates a new SelectorIndividualOktaServiceAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelectorIndividualOktaServiceAccountWithDefaults() *SelectorIndividualOktaServiceAccount {
	this := SelectorIndividualOktaServiceAccount{}
	return &this
}

// GetType returns the Type field value
func (o *SelectorIndividualOktaServiceAccount) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SelectorIndividualOktaServiceAccount) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SelectorIndividualOktaServiceAccount) SetType(v string) *SelectorIndividualOktaServiceAccount {
	o.Type = v
	return o
}

// GetServiceAccount returns the ServiceAccount field value
func (o *SelectorIndividualOktaServiceAccount) GetServiceAccount() NamedObject {
	if o == nil {
		var ret NamedObject
		return ret
	}

	return o.ServiceAccount
}

// GetServiceAccountOk returns a tuple with the ServiceAccount field value
// and a boolean to check if the value has been set.
func (o *SelectorIndividualOktaServiceAccount) GetServiceAccountOk() (*NamedObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceAccount, true
}

// SetServiceAccount sets field value
func (o *SelectorIndividualOktaServiceAccount) SetServiceAccount(v NamedObject) *SelectorIndividualOktaServiceAccount {
	o.ServiceAccount = v
	return o
}

func (o SelectorIndividualOktaServiceAccount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SelectorIndividualOktaServiceAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_type"] = o.Type
	toSerialize["service_account"] = o.ServiceAccount
	return toSerialize, nil
}

type NullableSelectorIndividualOktaServiceAccount struct {
	value *SelectorIndividualOktaServiceAccount
	isSet bool
}

func (v NullableSelectorIndividualOktaServiceAccount) Get() *SelectorIndividualOktaServiceAccount {
	return v.value
}

func (v *NullableSelectorIndividualOktaServiceAccount) Set(val *SelectorIndividualOktaServiceAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableSelectorIndividualOktaServiceAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableSelectorIndividualOktaServiceAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelectorIndividualOktaServiceAccount(val *SelectorIndividualOktaServiceAccount) *NullableSelectorIndividualOktaServiceAccount {
	return &NullableSelectorIndividualOktaServiceAccount{value: val, isSet: true}
}

func (v NullableSelectorIndividualOktaServiceAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelectorIndividualOktaServiceAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
