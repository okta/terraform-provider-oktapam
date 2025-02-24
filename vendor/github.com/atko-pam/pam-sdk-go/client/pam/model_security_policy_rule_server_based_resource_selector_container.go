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

// checks if the SecurityPolicyRuleServerBasedResourceSelectorContainer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityPolicyRuleServerBasedResourceSelectorContainer{}

// SecurityPolicyRuleServerBasedResourceSelectorContainer struct for SecurityPolicyRuleServerBasedResourceSelectorContainer
type SecurityPolicyRuleServerBasedResourceSelectorContainer struct {
	SelectorType SecurityPolicyRuleServerBasedResourceSubSelectorType     `json:"selector_type"`
	Selector     SecurityPolicyRuleResourceServerBasedResourceSubSelector `json:"selector"`
}

// NewSecurityPolicyRuleServerBasedResourceSelectorContainer instantiates a new SecurityPolicyRuleServerBasedResourceSelectorContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityPolicyRuleServerBasedResourceSelectorContainer(selectorType SecurityPolicyRuleServerBasedResourceSubSelectorType, selector SecurityPolicyRuleResourceServerBasedResourceSubSelector) *SecurityPolicyRuleServerBasedResourceSelectorContainer {
	this := SecurityPolicyRuleServerBasedResourceSelectorContainer{}
	this.SelectorType = selectorType
	this.Selector = selector
	return &this
}

// NewSecurityPolicyRuleServerBasedResourceSelectorContainerWithDefaults instantiates a new SecurityPolicyRuleServerBasedResourceSelectorContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityPolicyRuleServerBasedResourceSelectorContainerWithDefaults() *SecurityPolicyRuleServerBasedResourceSelectorContainer {
	this := SecurityPolicyRuleServerBasedResourceSelectorContainer{}
	return &this
}

// GetSelectorType returns the SelectorType field value
func (o *SecurityPolicyRuleServerBasedResourceSelectorContainer) GetSelectorType() SecurityPolicyRuleServerBasedResourceSubSelectorType {
	if o == nil {
		var ret SecurityPolicyRuleServerBasedResourceSubSelectorType
		return ret
	}

	return o.SelectorType
}

// GetSelectorTypeOk returns a tuple with the SelectorType field value
// and a boolean to check if the value has been set.
func (o *SecurityPolicyRuleServerBasedResourceSelectorContainer) GetSelectorTypeOk() (*SecurityPolicyRuleServerBasedResourceSubSelectorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelectorType, true
}

// SetSelectorType sets field value
func (o *SecurityPolicyRuleServerBasedResourceSelectorContainer) SetSelectorType(v SecurityPolicyRuleServerBasedResourceSubSelectorType) *SecurityPolicyRuleServerBasedResourceSelectorContainer {
	o.SelectorType = v
	return o
}

// GetSelector returns the Selector field value
func (o *SecurityPolicyRuleServerBasedResourceSelectorContainer) GetSelector() SecurityPolicyRuleResourceServerBasedResourceSubSelector {
	if o == nil {
		var ret SecurityPolicyRuleResourceServerBasedResourceSubSelector
		return ret
	}

	return o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value
// and a boolean to check if the value has been set.
func (o *SecurityPolicyRuleServerBasedResourceSelectorContainer) GetSelectorOk() (*SecurityPolicyRuleResourceServerBasedResourceSubSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Selector, true
}

// SetSelector sets field value
func (o *SecurityPolicyRuleServerBasedResourceSelectorContainer) SetSelector(v SecurityPolicyRuleResourceServerBasedResourceSubSelector) *SecurityPolicyRuleServerBasedResourceSelectorContainer {
	o.Selector = v
	return o
}

func (o SecurityPolicyRuleServerBasedResourceSelectorContainer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityPolicyRuleServerBasedResourceSelectorContainer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["selector_type"] = o.SelectorType
	toSerialize["selector"] = o.Selector
	return toSerialize, nil
}

type NullableSecurityPolicyRuleServerBasedResourceSelectorContainer struct {
	value *SecurityPolicyRuleServerBasedResourceSelectorContainer
	isSet bool
}

func (v NullableSecurityPolicyRuleServerBasedResourceSelectorContainer) Get() *SecurityPolicyRuleServerBasedResourceSelectorContainer {
	return v.value
}

func (v *NullableSecurityPolicyRuleServerBasedResourceSelectorContainer) Set(val *SecurityPolicyRuleServerBasedResourceSelectorContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityPolicyRuleServerBasedResourceSelectorContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityPolicyRuleServerBasedResourceSelectorContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityPolicyRuleServerBasedResourceSelectorContainer(val *SecurityPolicyRuleServerBasedResourceSelectorContainer) *NullableSecurityPolicyRuleServerBasedResourceSelectorContainer {
	return &NullableSecurityPolicyRuleServerBasedResourceSelectorContainer{value: val, isSet: true}
}

func (v NullableSecurityPolicyRuleServerBasedResourceSelectorContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityPolicyRuleServerBasedResourceSelectorContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
