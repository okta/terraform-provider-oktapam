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

// checks if the SecurityPolicyNoneAccountSelector type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityPolicyNoneAccountSelector{}

// SecurityPolicyNoneAccountSelector An empty object, required when no account selector is used (aka the \"none\" selector).
type SecurityPolicyNoneAccountSelector struct {
	Type *string `json:"_type,omitempty"`
}

// NewSecurityPolicyNoneAccountSelector instantiates a new SecurityPolicyNoneAccountSelector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityPolicyNoneAccountSelector() *SecurityPolicyNoneAccountSelector {
	this := SecurityPolicyNoneAccountSelector{}
	return &this
}

// NewSecurityPolicyNoneAccountSelectorWithDefaults instantiates a new SecurityPolicyNoneAccountSelector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityPolicyNoneAccountSelectorWithDefaults() *SecurityPolicyNoneAccountSelector {
	this := SecurityPolicyNoneAccountSelector{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SecurityPolicyNoneAccountSelector) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityPolicyNoneAccountSelector) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SecurityPolicyNoneAccountSelector) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SecurityPolicyNoneAccountSelector) SetType(v string) *SecurityPolicyNoneAccountSelector {
	o.Type = &v
	return o
}

func (o SecurityPolicyNoneAccountSelector) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityPolicyNoneAccountSelector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["_type"] = o.Type
	}
	return toSerialize, nil
}

type NullableSecurityPolicyNoneAccountSelector struct {
	value *SecurityPolicyNoneAccountSelector
	isSet bool
}

func (v NullableSecurityPolicyNoneAccountSelector) Get() *SecurityPolicyNoneAccountSelector {
	return v.value
}

func (v *NullableSecurityPolicyNoneAccountSelector) Set(val *SecurityPolicyNoneAccountSelector) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityPolicyNoneAccountSelector) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityPolicyNoneAccountSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityPolicyNoneAccountSelector(val *SecurityPolicyNoneAccountSelector) *NullableSecurityPolicyNoneAccountSelector {
	return &NullableSecurityPolicyNoneAccountSelector{value: val, isSet: true}
}

func (v NullableSecurityPolicyNoneAccountSelector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityPolicyNoneAccountSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}