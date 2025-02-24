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

// checks if the ServiceAccountPrivilegeContainer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceAccountPrivilegeContainer{}

// ServiceAccountPrivilegeContainer Privileges granted to a user for a service account
type ServiceAccountPrivilegeContainer struct {
	PrivilegeType  *SecurityPolicyRulePrivilegeTypeForServiceAccounts `json:"privilege_type,omitempty"`
	PrivilegeValue *ServiceAccountPrivilegeContainerPrivilegeValue    `json:"privilege_value,omitempty"`
}

// NewServiceAccountPrivilegeContainer instantiates a new ServiceAccountPrivilegeContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountPrivilegeContainer() *ServiceAccountPrivilegeContainer {
	this := ServiceAccountPrivilegeContainer{}
	return &this
}

// NewServiceAccountPrivilegeContainerWithDefaults instantiates a new ServiceAccountPrivilegeContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountPrivilegeContainerWithDefaults() *ServiceAccountPrivilegeContainer {
	this := ServiceAccountPrivilegeContainer{}
	return &this
}

// GetPrivilegeType returns the PrivilegeType field value if set, zero value otherwise.
func (o *ServiceAccountPrivilegeContainer) GetPrivilegeType() SecurityPolicyRulePrivilegeTypeForServiceAccounts {
	if o == nil || IsNil(o.PrivilegeType) {
		var ret SecurityPolicyRulePrivilegeTypeForServiceAccounts
		return ret
	}
	return *o.PrivilegeType
}

// GetPrivilegeTypeOk returns a tuple with the PrivilegeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountPrivilegeContainer) GetPrivilegeTypeOk() (*SecurityPolicyRulePrivilegeTypeForServiceAccounts, bool) {
	if o == nil || IsNil(o.PrivilegeType) {
		return nil, false
	}
	return o.PrivilegeType, true
}

// HasPrivilegeType returns a boolean if a field has been set.
func (o *ServiceAccountPrivilegeContainer) HasPrivilegeType() bool {
	if o != nil && !IsNil(o.PrivilegeType) {
		return true
	}

	return false
}

// SetPrivilegeType gets a reference to the given SecurityPolicyRulePrivilegeTypeForServiceAccounts and assigns it to the PrivilegeType field.
func (o *ServiceAccountPrivilegeContainer) SetPrivilegeType(v SecurityPolicyRulePrivilegeTypeForServiceAccounts) *ServiceAccountPrivilegeContainer {
	o.PrivilegeType = &v
	return o
}

// GetPrivilegeValue returns the PrivilegeValue field value if set, zero value otherwise.
func (o *ServiceAccountPrivilegeContainer) GetPrivilegeValue() ServiceAccountPrivilegeContainerPrivilegeValue {
	if o == nil || IsNil(o.PrivilegeValue) {
		var ret ServiceAccountPrivilegeContainerPrivilegeValue
		return ret
	}
	return *o.PrivilegeValue
}

// GetPrivilegeValueOk returns a tuple with the PrivilegeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountPrivilegeContainer) GetPrivilegeValueOk() (*ServiceAccountPrivilegeContainerPrivilegeValue, bool) {
	if o == nil || IsNil(o.PrivilegeValue) {
		return nil, false
	}
	return o.PrivilegeValue, true
}

// HasPrivilegeValue returns a boolean if a field has been set.
func (o *ServiceAccountPrivilegeContainer) HasPrivilegeValue() bool {
	if o != nil && !IsNil(o.PrivilegeValue) {
		return true
	}

	return false
}

// SetPrivilegeValue gets a reference to the given ServiceAccountPrivilegeContainerPrivilegeValue and assigns it to the PrivilegeValue field.
func (o *ServiceAccountPrivilegeContainer) SetPrivilegeValue(v ServiceAccountPrivilegeContainerPrivilegeValue) *ServiceAccountPrivilegeContainer {
	o.PrivilegeValue = &v
	return o
}

func (o ServiceAccountPrivilegeContainer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceAccountPrivilegeContainer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PrivilegeType) {
		toSerialize["privilege_type"] = o.PrivilegeType
	}
	if !IsNil(o.PrivilegeValue) {
		toSerialize["privilege_value"] = o.PrivilegeValue
	}
	return toSerialize, nil
}

type NullableServiceAccountPrivilegeContainer struct {
	value *ServiceAccountPrivilegeContainer
	isSet bool
}

func (v NullableServiceAccountPrivilegeContainer) Get() *ServiceAccountPrivilegeContainer {
	return v.value
}

func (v *NullableServiceAccountPrivilegeContainer) Set(val *ServiceAccountPrivilegeContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccountPrivilegeContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccountPrivilegeContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccountPrivilegeContainer(val *ServiceAccountPrivilegeContainer) *NullableServiceAccountPrivilegeContainer {
	return &NullableServiceAccountPrivilegeContainer{value: val, isSet: true}
}

func (v NullableServiceAccountPrivilegeContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccountPrivilegeContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
