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

// checks if the SecurityPolicySecretBasedResourceSelector type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityPolicySecretBasedResourceSelector{}

// SecurityPolicySecretBasedResourceSelector The selector that defines resources targeted by this Security policy for secret based resources
type SecurityPolicySecretBasedResourceSelector struct {
	Type      string                                               `json:"_type"`
	Selectors []SecurityPolicySecretBasedResourceSelectorContainer `json:"selectors"`
}

// NewSecurityPolicySecretBasedResourceSelector instantiates a new SecurityPolicySecretBasedResourceSelector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityPolicySecretBasedResourceSelector(type_ string, selectors []SecurityPolicySecretBasedResourceSelectorContainer) *SecurityPolicySecretBasedResourceSelector {
	this := SecurityPolicySecretBasedResourceSelector{}
	this.Type = type_
	this.Selectors = selectors
	return &this
}

// NewSecurityPolicySecretBasedResourceSelectorWithDefaults instantiates a new SecurityPolicySecretBasedResourceSelector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityPolicySecretBasedResourceSelectorWithDefaults() *SecurityPolicySecretBasedResourceSelector {
	this := SecurityPolicySecretBasedResourceSelector{}
	return &this
}

// GetType returns the Type field value
func (o *SecurityPolicySecretBasedResourceSelector) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SecurityPolicySecretBasedResourceSelector) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SecurityPolicySecretBasedResourceSelector) SetType(v string) *SecurityPolicySecretBasedResourceSelector {
	o.Type = v
	return o
}

// GetSelectors returns the Selectors field value
func (o *SecurityPolicySecretBasedResourceSelector) GetSelectors() []SecurityPolicySecretBasedResourceSelectorContainer {
	if o == nil {
		var ret []SecurityPolicySecretBasedResourceSelectorContainer
		return ret
	}

	return o.Selectors
}

// GetSelectorsOk returns a tuple with the Selectors field value
// and a boolean to check if the value has been set.
func (o *SecurityPolicySecretBasedResourceSelector) GetSelectorsOk() ([]SecurityPolicySecretBasedResourceSelectorContainer, bool) {
	if o == nil {
		return nil, false
	}
	return o.Selectors, true
}

// SetSelectors sets field value
func (o *SecurityPolicySecretBasedResourceSelector) SetSelectors(v []SecurityPolicySecretBasedResourceSelectorContainer) *SecurityPolicySecretBasedResourceSelector {
	o.Selectors = v
	return o
}

func (o SecurityPolicySecretBasedResourceSelector) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityPolicySecretBasedResourceSelector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_type"] = o.Type
	toSerialize["selectors"] = o.Selectors
	return toSerialize, nil
}

type NullableSecurityPolicySecretBasedResourceSelector struct {
	value *SecurityPolicySecretBasedResourceSelector
	isSet bool
}

func (v NullableSecurityPolicySecretBasedResourceSelector) Get() *SecurityPolicySecretBasedResourceSelector {
	return v.value
}

func (v *NullableSecurityPolicySecretBasedResourceSelector) Set(val *SecurityPolicySecretBasedResourceSelector) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityPolicySecretBasedResourceSelector) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityPolicySecretBasedResourceSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityPolicySecretBasedResourceSelector(val *SecurityPolicySecretBasedResourceSelector) *NullableSecurityPolicySecretBasedResourceSelector {
	return &NullableSecurityPolicySecretBasedResourceSelector{value: val, isSet: true}
}

func (v NullableSecurityPolicySecretBasedResourceSelector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityPolicySecretBasedResourceSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}