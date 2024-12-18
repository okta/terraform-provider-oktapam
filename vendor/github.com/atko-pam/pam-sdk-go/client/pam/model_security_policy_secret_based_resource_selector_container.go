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

// checks if the SecurityPolicySecretBasedResourceSelectorContainer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityPolicySecretBasedResourceSelectorContainer{}

// SecurityPolicySecretBasedResourceSelectorContainer struct for SecurityPolicySecretBasedResourceSelectorContainer
type SecurityPolicySecretBasedResourceSelectorContainer struct {
	// The type of selector used to target secret based resources
	SelectorType string                                                     `json:"selector_type"`
	Selector     SecurityPolicySecretBasedResourceSelectorContainerSelector `json:"selector"`
}

// NewSecurityPolicySecretBasedResourceSelectorContainer instantiates a new SecurityPolicySecretBasedResourceSelectorContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityPolicySecretBasedResourceSelectorContainer(selectorType string, selector SecurityPolicySecretBasedResourceSelectorContainerSelector) *SecurityPolicySecretBasedResourceSelectorContainer {
	this := SecurityPolicySecretBasedResourceSelectorContainer{}
	this.SelectorType = selectorType
	this.Selector = selector
	return &this
}

// NewSecurityPolicySecretBasedResourceSelectorContainerWithDefaults instantiates a new SecurityPolicySecretBasedResourceSelectorContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityPolicySecretBasedResourceSelectorContainerWithDefaults() *SecurityPolicySecretBasedResourceSelectorContainer {
	this := SecurityPolicySecretBasedResourceSelectorContainer{}
	return &this
}

// GetSelectorType returns the SelectorType field value
func (o *SecurityPolicySecretBasedResourceSelectorContainer) GetSelectorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SelectorType
}

// GetSelectorTypeOk returns a tuple with the SelectorType field value
// and a boolean to check if the value has been set.
func (o *SecurityPolicySecretBasedResourceSelectorContainer) GetSelectorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelectorType, true
}

// SetSelectorType sets field value
func (o *SecurityPolicySecretBasedResourceSelectorContainer) SetSelectorType(v string) *SecurityPolicySecretBasedResourceSelectorContainer {
	o.SelectorType = v
	return o
}

// GetSelector returns the Selector field value
func (o *SecurityPolicySecretBasedResourceSelectorContainer) GetSelector() SecurityPolicySecretBasedResourceSelectorContainerSelector {
	if o == nil {
		var ret SecurityPolicySecretBasedResourceSelectorContainerSelector
		return ret
	}

	return o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value
// and a boolean to check if the value has been set.
func (o *SecurityPolicySecretBasedResourceSelectorContainer) GetSelectorOk() (*SecurityPolicySecretBasedResourceSelectorContainerSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Selector, true
}

// SetSelector sets field value
func (o *SecurityPolicySecretBasedResourceSelectorContainer) SetSelector(v SecurityPolicySecretBasedResourceSelectorContainerSelector) *SecurityPolicySecretBasedResourceSelectorContainer {
	o.Selector = v
	return o
}

func (o SecurityPolicySecretBasedResourceSelectorContainer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityPolicySecretBasedResourceSelectorContainer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["selector_type"] = o.SelectorType
	toSerialize["selector"] = o.Selector
	return toSerialize, nil
}

type NullableSecurityPolicySecretBasedResourceSelectorContainer struct {
	value *SecurityPolicySecretBasedResourceSelectorContainer
	isSet bool
}

func (v NullableSecurityPolicySecretBasedResourceSelectorContainer) Get() *SecurityPolicySecretBasedResourceSelectorContainer {
	return v.value
}

func (v *NullableSecurityPolicySecretBasedResourceSelectorContainer) Set(val *SecurityPolicySecretBasedResourceSelectorContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityPolicySecretBasedResourceSelectorContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityPolicySecretBasedResourceSelectorContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityPolicySecretBasedResourceSelectorContainer(val *SecurityPolicySecretBasedResourceSelectorContainer) *NullableSecurityPolicySecretBasedResourceSelectorContainer {
	return &NullableSecurityPolicySecretBasedResourceSelectorContainer{value: val, isSet: true}
}

func (v NullableSecurityPolicySecretBasedResourceSelectorContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityPolicySecretBasedResourceSelectorContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
