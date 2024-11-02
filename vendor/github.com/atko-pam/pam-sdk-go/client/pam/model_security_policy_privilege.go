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

// checks if the SecurityPolicyPrivilege type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityPolicyPrivilege{}

// SecurityPolicyPrivilege struct for SecurityPolicyPrivilege
type SecurityPolicyPrivilege struct {
	Type SecurityPolicyRulePrivilegeType `json:"_type"`
}

// NewSecurityPolicyPrivilege instantiates a new SecurityPolicyPrivilege object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityPolicyPrivilege(type_ SecurityPolicyRulePrivilegeType) *SecurityPolicyPrivilege {
	this := SecurityPolicyPrivilege{}
	this.Type = type_
	return &this
}

// NewSecurityPolicyPrivilegeWithDefaults instantiates a new SecurityPolicyPrivilege object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityPolicyPrivilegeWithDefaults() *SecurityPolicyPrivilege {
	this := SecurityPolicyPrivilege{}
	return &this
}

// GetType returns the Type field value
func (o *SecurityPolicyPrivilege) GetType() SecurityPolicyRulePrivilegeType {
	if o == nil {
		var ret SecurityPolicyRulePrivilegeType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SecurityPolicyPrivilege) GetTypeOk() (*SecurityPolicyRulePrivilegeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SecurityPolicyPrivilege) SetType(v SecurityPolicyRulePrivilegeType) *SecurityPolicyPrivilege {
	o.Type = v
	return o
}

func (o SecurityPolicyPrivilege) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityPolicyPrivilege) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_type"] = o.Type
	return toSerialize, nil
}

type NullableSecurityPolicyPrivilege struct {
	value *SecurityPolicyPrivilege
	isSet bool
}

func (v NullableSecurityPolicyPrivilege) Get() *SecurityPolicyPrivilege {
	return v.value
}

func (v *NullableSecurityPolicyPrivilege) Set(val *SecurityPolicyPrivilege) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityPolicyPrivilege) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityPolicyPrivilege) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityPolicyPrivilege(val *SecurityPolicyPrivilege) *NullableSecurityPolicyPrivilege {
	return &NullableSecurityPolicyPrivilege{value: val, isSet: true}
}

func (v NullableSecurityPolicyPrivilege) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityPolicyPrivilege) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
