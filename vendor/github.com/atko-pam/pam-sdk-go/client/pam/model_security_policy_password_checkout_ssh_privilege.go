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

// checks if the SecurityPolicyPasswordCheckoutSSHPrivilege type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityPolicyPasswordCheckoutSSHPrivilege{}

// SecurityPolicyPasswordCheckoutSSHPrivilege SecurityPolicyPasswordCheckoutSSHPrivilege lets the user sign in via SSH by checking out a vaulted password. It  indicates that the principal will be allowed to check out passwords to sign in with SSH to servers for accounts specified via a  ServerBasedResourceSelector.
type SecurityPolicyPasswordCheckoutSSHPrivilege struct {
	SecurityPolicyPrivilege
	PasswordCheckoutSsh bool `json:"password_checkout_ssh"`
}

// NewSecurityPolicyPasswordCheckoutSSHPrivilege instantiates a new SecurityPolicyPasswordCheckoutSSHPrivilege object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityPolicyPasswordCheckoutSSHPrivilege(passwordCheckoutSsh bool, type_ SecurityPolicyRulePrivilegeType) *SecurityPolicyPasswordCheckoutSSHPrivilege {
	this := SecurityPolicyPasswordCheckoutSSHPrivilege{}
	this.Type = type_
	this.PasswordCheckoutSsh = passwordCheckoutSsh
	return &this
}

// NewSecurityPolicyPasswordCheckoutSSHPrivilegeWithDefaults instantiates a new SecurityPolicyPasswordCheckoutSSHPrivilege object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityPolicyPasswordCheckoutSSHPrivilegeWithDefaults() *SecurityPolicyPasswordCheckoutSSHPrivilege {
	this := SecurityPolicyPasswordCheckoutSSHPrivilege{}
	return &this
}

// GetPasswordCheckoutSsh returns the PasswordCheckoutSsh field value
func (o *SecurityPolicyPasswordCheckoutSSHPrivilege) GetPasswordCheckoutSsh() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PasswordCheckoutSsh
}

// GetPasswordCheckoutSshOk returns a tuple with the PasswordCheckoutSsh field value
// and a boolean to check if the value has been set.
func (o *SecurityPolicyPasswordCheckoutSSHPrivilege) GetPasswordCheckoutSshOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PasswordCheckoutSsh, true
}

// SetPasswordCheckoutSsh sets field value
func (o *SecurityPolicyPasswordCheckoutSSHPrivilege) SetPasswordCheckoutSsh(v bool) *SecurityPolicyPasswordCheckoutSSHPrivilege {
	o.PasswordCheckoutSsh = v
	return o
}

func (o SecurityPolicyPasswordCheckoutSSHPrivilege) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityPolicyPasswordCheckoutSSHPrivilege) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedSecurityPolicyPrivilege, errSecurityPolicyPrivilege := json.Marshal(o.SecurityPolicyPrivilege)
	if errSecurityPolicyPrivilege != nil {
		return map[string]interface{}{}, errSecurityPolicyPrivilege
	}
	errSecurityPolicyPrivilege = json.Unmarshal([]byte(serializedSecurityPolicyPrivilege), &toSerialize)
	if errSecurityPolicyPrivilege != nil {
		return map[string]interface{}{}, errSecurityPolicyPrivilege
	}
	toSerialize["password_checkout_ssh"] = o.PasswordCheckoutSsh
	return toSerialize, nil
}

type NullableSecurityPolicyPasswordCheckoutSSHPrivilege struct {
	value *SecurityPolicyPasswordCheckoutSSHPrivilege
	isSet bool
}

func (v NullableSecurityPolicyPasswordCheckoutSSHPrivilege) Get() *SecurityPolicyPasswordCheckoutSSHPrivilege {
	return v.value
}

func (v *NullableSecurityPolicyPasswordCheckoutSSHPrivilege) Set(val *SecurityPolicyPasswordCheckoutSSHPrivilege) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityPolicyPasswordCheckoutSSHPrivilege) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityPolicyPasswordCheckoutSSHPrivilege) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityPolicyPasswordCheckoutSSHPrivilege(val *SecurityPolicyPasswordCheckoutSSHPrivilege) *NullableSecurityPolicyPasswordCheckoutSSHPrivilege {
	return &NullableSecurityPolicyPasswordCheckoutSSHPrivilege{value: val, isSet: true}
}

func (v NullableSecurityPolicyPasswordCheckoutSSHPrivilege) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityPolicyPasswordCheckoutSSHPrivilege) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
