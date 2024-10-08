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

// checks if the PrivilegedAccountPrivilegeContainer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrivilegedAccountPrivilegeContainer{}

// PrivilegedAccountPrivilegeContainer Privleges granted to a user for a privileged account
type PrivilegedAccountPrivilegeContainer struct {
	PrivilegeType  *SecurityPolicyRulePrivilegeTypeForPrivilegedAccounts `json:"privilege_type,omitempty"`
	PrivilegeValue *PrivilegedAccountPrivilegeContainerPrivilegeValue    `json:"privilege_value,omitempty"`
}

// NewPrivilegedAccountPrivilegeContainer instantiates a new PrivilegedAccountPrivilegeContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivilegedAccountPrivilegeContainer() *PrivilegedAccountPrivilegeContainer {
	this := PrivilegedAccountPrivilegeContainer{}
	return &this
}

// NewPrivilegedAccountPrivilegeContainerWithDefaults instantiates a new PrivilegedAccountPrivilegeContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivilegedAccountPrivilegeContainerWithDefaults() *PrivilegedAccountPrivilegeContainer {
	this := PrivilegedAccountPrivilegeContainer{}
	return &this
}

// GetPrivilegeType returns the PrivilegeType field value if set, zero value otherwise.
func (o *PrivilegedAccountPrivilegeContainer) GetPrivilegeType() SecurityPolicyRulePrivilegeTypeForPrivilegedAccounts {
	if o == nil || IsNil(o.PrivilegeType) {
		var ret SecurityPolicyRulePrivilegeTypeForPrivilegedAccounts
		return ret
	}
	return *o.PrivilegeType
}

// GetPrivilegeTypeOk returns a tuple with the PrivilegeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivilegedAccountPrivilegeContainer) GetPrivilegeTypeOk() (*SecurityPolicyRulePrivilegeTypeForPrivilegedAccounts, bool) {
	if o == nil || IsNil(o.PrivilegeType) {
		return nil, false
	}
	return o.PrivilegeType, true
}

// HasPrivilegeType returns a boolean if a field has been set.
func (o *PrivilegedAccountPrivilegeContainer) HasPrivilegeType() bool {
	if o != nil && !IsNil(o.PrivilegeType) {
		return true
	}

	return false
}

// SetPrivilegeType gets a reference to the given SecurityPolicyRulePrivilegeTypeForPrivilegedAccounts and assigns it to the PrivilegeType field.
func (o *PrivilegedAccountPrivilegeContainer) SetPrivilegeType(v SecurityPolicyRulePrivilegeTypeForPrivilegedAccounts) *PrivilegedAccountPrivilegeContainer {
	o.PrivilegeType = &v
	return o
}

// GetPrivilegeValue returns the PrivilegeValue field value if set, zero value otherwise.
func (o *PrivilegedAccountPrivilegeContainer) GetPrivilegeValue() PrivilegedAccountPrivilegeContainerPrivilegeValue {
	if o == nil || IsNil(o.PrivilegeValue) {
		var ret PrivilegedAccountPrivilegeContainerPrivilegeValue
		return ret
	}
	return *o.PrivilegeValue
}

// GetPrivilegeValueOk returns a tuple with the PrivilegeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivilegedAccountPrivilegeContainer) GetPrivilegeValueOk() (*PrivilegedAccountPrivilegeContainerPrivilegeValue, bool) {
	if o == nil || IsNil(o.PrivilegeValue) {
		return nil, false
	}
	return o.PrivilegeValue, true
}

// HasPrivilegeValue returns a boolean if a field has been set.
func (o *PrivilegedAccountPrivilegeContainer) HasPrivilegeValue() bool {
	if o != nil && !IsNil(o.PrivilegeValue) {
		return true
	}

	return false
}

// SetPrivilegeValue gets a reference to the given PrivilegedAccountPrivilegeContainerPrivilegeValue and assigns it to the PrivilegeValue field.
func (o *PrivilegedAccountPrivilegeContainer) SetPrivilegeValue(v PrivilegedAccountPrivilegeContainerPrivilegeValue) *PrivilegedAccountPrivilegeContainer {
	o.PrivilegeValue = &v
	return o
}

func (o PrivilegedAccountPrivilegeContainer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrivilegedAccountPrivilegeContainer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PrivilegeType) {
		toSerialize["privilege_type"] = o.PrivilegeType
	}
	if !IsNil(o.PrivilegeValue) {
		toSerialize["privilege_value"] = o.PrivilegeValue
	}
	return toSerialize, nil
}

type NullablePrivilegedAccountPrivilegeContainer struct {
	value *PrivilegedAccountPrivilegeContainer
	isSet bool
}

func (v NullablePrivilegedAccountPrivilegeContainer) Get() *PrivilegedAccountPrivilegeContainer {
	return v.value
}

func (v *NullablePrivilegedAccountPrivilegeContainer) Set(val *PrivilegedAccountPrivilegeContainer) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivilegedAccountPrivilegeContainer) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivilegedAccountPrivilegeContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivilegedAccountPrivilegeContainer(val *PrivilegedAccountPrivilegeContainer) *NullablePrivilegedAccountPrivilegeContainer {
	return &NullablePrivilegedAccountPrivilegeContainer{value: val, isSet: true}
}

func (v NullablePrivilegedAccountPrivilegeContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivilegedAccountPrivilegeContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
