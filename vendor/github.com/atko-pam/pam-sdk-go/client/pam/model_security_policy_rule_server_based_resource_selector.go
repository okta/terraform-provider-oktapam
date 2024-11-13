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

// checks if the SecurityPolicyRuleServerBasedResourceSelector type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityPolicyRuleServerBasedResourceSelector{}

// SecurityPolicyRuleServerBasedResourceSelector The selector that defines resources targeted by this Security Policy for server based resources
type SecurityPolicyRuleServerBasedResourceSelector struct {
	Type      string                                                   `json:"_type"`
	Selectors []SecurityPolicyRuleServerBasedResourceSelectorContainer `json:"selectors"`
}

// NewSecurityPolicyRuleServerBasedResourceSelector instantiates a new SecurityPolicyRuleServerBasedResourceSelector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityPolicyRuleServerBasedResourceSelector(type_ string, selectors []SecurityPolicyRuleServerBasedResourceSelectorContainer) *SecurityPolicyRuleServerBasedResourceSelector {
	this := SecurityPolicyRuleServerBasedResourceSelector{}
	this.Type = type_
	this.Selectors = selectors
	return &this
}

// NewSecurityPolicyRuleServerBasedResourceSelectorWithDefaults instantiates a new SecurityPolicyRuleServerBasedResourceSelector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityPolicyRuleServerBasedResourceSelectorWithDefaults() *SecurityPolicyRuleServerBasedResourceSelector {
	this := SecurityPolicyRuleServerBasedResourceSelector{}
	return &this
}

// GetType returns the Type field value
func (o *SecurityPolicyRuleServerBasedResourceSelector) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SecurityPolicyRuleServerBasedResourceSelector) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SecurityPolicyRuleServerBasedResourceSelector) SetType(v string) *SecurityPolicyRuleServerBasedResourceSelector {
	o.Type = v
	return o
}

// GetSelectors returns the Selectors field value
func (o *SecurityPolicyRuleServerBasedResourceSelector) GetSelectors() []SecurityPolicyRuleServerBasedResourceSelectorContainer {
	if o == nil {
		var ret []SecurityPolicyRuleServerBasedResourceSelectorContainer
		return ret
	}

	return o.Selectors
}

// GetSelectorsOk returns a tuple with the Selectors field value
// and a boolean to check if the value has been set.
func (o *SecurityPolicyRuleServerBasedResourceSelector) GetSelectorsOk() ([]SecurityPolicyRuleServerBasedResourceSelectorContainer, bool) {
	if o == nil {
		return nil, false
	}
	return o.Selectors, true
}

// SetSelectors sets field value
func (o *SecurityPolicyRuleServerBasedResourceSelector) SetSelectors(v []SecurityPolicyRuleServerBasedResourceSelectorContainer) *SecurityPolicyRuleServerBasedResourceSelector {
	o.Selectors = v
	return o
}

func (o SecurityPolicyRuleServerBasedResourceSelector) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityPolicyRuleServerBasedResourceSelector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_type"] = o.Type
	toSerialize["selectors"] = o.Selectors
	return toSerialize, nil
}

type NullableSecurityPolicyRuleServerBasedResourceSelector struct {
	value *SecurityPolicyRuleServerBasedResourceSelector
	isSet bool
}

func (v NullableSecurityPolicyRuleServerBasedResourceSelector) Get() *SecurityPolicyRuleServerBasedResourceSelector {
	return v.value
}

func (v *NullableSecurityPolicyRuleServerBasedResourceSelector) Set(val *SecurityPolicyRuleServerBasedResourceSelector) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityPolicyRuleServerBasedResourceSelector) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityPolicyRuleServerBasedResourceSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityPolicyRuleServerBasedResourceSelector(val *SecurityPolicyRuleServerBasedResourceSelector) *NullableSecurityPolicyRuleServerBasedResourceSelector {
	return &NullableSecurityPolicyRuleServerBasedResourceSelector{value: val, isSet: true}
}

func (v NullableSecurityPolicyRuleServerBasedResourceSelector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityPolicyRuleServerBasedResourceSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
