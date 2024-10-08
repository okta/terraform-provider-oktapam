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

// checks if the SecurityPolicyRuleUpdatePasswordPrivilege type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityPolicyRuleUpdatePasswordPrivilege{}

// SecurityPolicyRuleUpdatePasswordPrivilege Privilege to grant access to update password for an account
type SecurityPolicyRuleUpdatePasswordPrivilege struct {
	RevealPassword *bool `json:"reveal_password,omitempty"`
}

// NewSecurityPolicyRuleUpdatePasswordPrivilege instantiates a new SecurityPolicyRuleUpdatePasswordPrivilege object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityPolicyRuleUpdatePasswordPrivilege() *SecurityPolicyRuleUpdatePasswordPrivilege {
	this := SecurityPolicyRuleUpdatePasswordPrivilege{}
	return &this
}

// NewSecurityPolicyRuleUpdatePasswordPrivilegeWithDefaults instantiates a new SecurityPolicyRuleUpdatePasswordPrivilege object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityPolicyRuleUpdatePasswordPrivilegeWithDefaults() *SecurityPolicyRuleUpdatePasswordPrivilege {
	this := SecurityPolicyRuleUpdatePasswordPrivilege{}
	return &this
}

// GetRevealPassword returns the RevealPassword field value if set, zero value otherwise.
func (o *SecurityPolicyRuleUpdatePasswordPrivilege) GetRevealPassword() bool {
	if o == nil || IsNil(o.RevealPassword) {
		var ret bool
		return ret
	}
	return *o.RevealPassword
}

// GetRevealPasswordOk returns a tuple with the RevealPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityPolicyRuleUpdatePasswordPrivilege) GetRevealPasswordOk() (*bool, bool) {
	if o == nil || IsNil(o.RevealPassword) {
		return nil, false
	}
	return o.RevealPassword, true
}

// HasRevealPassword returns a boolean if a field has been set.
func (o *SecurityPolicyRuleUpdatePasswordPrivilege) HasRevealPassword() bool {
	if o != nil && !IsNil(o.RevealPassword) {
		return true
	}

	return false
}

// SetRevealPassword gets a reference to the given bool and assigns it to the RevealPassword field.
func (o *SecurityPolicyRuleUpdatePasswordPrivilege) SetRevealPassword(v bool) *SecurityPolicyRuleUpdatePasswordPrivilege {
	o.RevealPassword = &v
	return o
}

func (o SecurityPolicyRuleUpdatePasswordPrivilege) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityPolicyRuleUpdatePasswordPrivilege) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RevealPassword) {
		toSerialize["reveal_password"] = o.RevealPassword
	}
	return toSerialize, nil
}

type NullableSecurityPolicyRuleUpdatePasswordPrivilege struct {
	value *SecurityPolicyRuleUpdatePasswordPrivilege
	isSet bool
}

func (v NullableSecurityPolicyRuleUpdatePasswordPrivilege) Get() *SecurityPolicyRuleUpdatePasswordPrivilege {
	return v.value
}

func (v *NullableSecurityPolicyRuleUpdatePasswordPrivilege) Set(val *SecurityPolicyRuleUpdatePasswordPrivilege) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityPolicyRuleUpdatePasswordPrivilege) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityPolicyRuleUpdatePasswordPrivilege) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityPolicyRuleUpdatePasswordPrivilege(val *SecurityPolicyRuleUpdatePasswordPrivilege) *NullableSecurityPolicyRuleUpdatePasswordPrivilege {
	return &NullableSecurityPolicyRuleUpdatePasswordPrivilege{value: val, isSet: true}
}

func (v NullableSecurityPolicyRuleUpdatePasswordPrivilege) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityPolicyRuleUpdatePasswordPrivilege) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
