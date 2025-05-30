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

// checks if the SaasAppAccountCheckedOutResourceDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SaasAppAccountCheckedOutResourceDetails{}

// SaasAppAccountCheckedOutResourceDetails struct for SaasAppAccountCheckedOutResourceDetails
type SaasAppAccountCheckedOutResourceDetails struct {
	// Type of this Resource object
	Type *string `json:"_type,omitempty"`
	// The name of the SaaS Application associated with the resource
	AppName *string `json:"app_name,omitempty"`
	// The name of the SaaS Application Account associated with the resource
	AccountName *string `json:"account_name,omitempty"`
}

// NewSaasAppAccountCheckedOutResourceDetails instantiates a new SaasAppAccountCheckedOutResourceDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSaasAppAccountCheckedOutResourceDetails() *SaasAppAccountCheckedOutResourceDetails {
	this := SaasAppAccountCheckedOutResourceDetails{}
	return &this
}

// NewSaasAppAccountCheckedOutResourceDetailsWithDefaults instantiates a new SaasAppAccountCheckedOutResourceDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSaasAppAccountCheckedOutResourceDetailsWithDefaults() *SaasAppAccountCheckedOutResourceDetails {
	this := SaasAppAccountCheckedOutResourceDetails{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SaasAppAccountCheckedOutResourceDetails) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaasAppAccountCheckedOutResourceDetails) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SaasAppAccountCheckedOutResourceDetails) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SaasAppAccountCheckedOutResourceDetails) SetType(v string) *SaasAppAccountCheckedOutResourceDetails {
	o.Type = &v
	return o
}

// GetAppName returns the AppName field value if set, zero value otherwise.
func (o *SaasAppAccountCheckedOutResourceDetails) GetAppName() string {
	if o == nil || IsNil(o.AppName) {
		var ret string
		return ret
	}
	return *o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaasAppAccountCheckedOutResourceDetails) GetAppNameOk() (*string, bool) {
	if o == nil || IsNil(o.AppName) {
		return nil, false
	}
	return o.AppName, true
}

// HasAppName returns a boolean if a field has been set.
func (o *SaasAppAccountCheckedOutResourceDetails) HasAppName() bool {
	if o != nil && !IsNil(o.AppName) {
		return true
	}

	return false
}

// SetAppName gets a reference to the given string and assigns it to the AppName field.
func (o *SaasAppAccountCheckedOutResourceDetails) SetAppName(v string) *SaasAppAccountCheckedOutResourceDetails {
	o.AppName = &v
	return o
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *SaasAppAccountCheckedOutResourceDetails) GetAccountName() string {
	if o == nil || IsNil(o.AccountName) {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaasAppAccountCheckedOutResourceDetails) GetAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.AccountName) {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *SaasAppAccountCheckedOutResourceDetails) HasAccountName() bool {
	if o != nil && !IsNil(o.AccountName) {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *SaasAppAccountCheckedOutResourceDetails) SetAccountName(v string) *SaasAppAccountCheckedOutResourceDetails {
	o.AccountName = &v
	return o
}

func (o SaasAppAccountCheckedOutResourceDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SaasAppAccountCheckedOutResourceDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["_type"] = o.Type
	}
	if !IsNil(o.AppName) {
		toSerialize["app_name"] = o.AppName
	}
	if !IsNil(o.AccountName) {
		toSerialize["account_name"] = o.AccountName
	}
	return toSerialize, nil
}

type NullableSaasAppAccountCheckedOutResourceDetails struct {
	value *SaasAppAccountCheckedOutResourceDetails
	isSet bool
}

func (v NullableSaasAppAccountCheckedOutResourceDetails) Get() *SaasAppAccountCheckedOutResourceDetails {
	return v.value
}

func (v *NullableSaasAppAccountCheckedOutResourceDetails) Set(val *SaasAppAccountCheckedOutResourceDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableSaasAppAccountCheckedOutResourceDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableSaasAppAccountCheckedOutResourceDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSaasAppAccountCheckedOutResourceDetails(val *SaasAppAccountCheckedOutResourceDetails) *NullableSaasAppAccountCheckedOutResourceDetails {
	return &NullableSaasAppAccountCheckedOutResourceDetails{value: val, isSet: true}
}

func (v NullableSaasAppAccountCheckedOutResourceDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSaasAppAccountCheckedOutResourceDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
