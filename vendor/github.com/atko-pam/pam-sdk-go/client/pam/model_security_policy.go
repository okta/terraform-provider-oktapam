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

// checks if the SecurityPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityPolicy{}

// SecurityPolicy struct for SecurityPolicy
type SecurityPolicy struct {
	// The UUID of the Security Policy
	Id *string `json:"id,omitempty"`
	// The name of the Security Policy
	Name string `json:"name"`
	// The description of the Security Policy
	Description string `json:"description"`
	// If `true`, indicates that the Security Policy is active
	Active     bool                     `json:"active"`
	Principals SecurityPolicyPrincipals `json:"principals"`
	// The rules associated with the Security Policy. A Security Policy can set multiple rules that define privileges available for matching resources.
	Rules []SecurityPolicyRule `json:"rules"`
}

// NewSecurityPolicy instantiates a new SecurityPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityPolicy(name string, description string, active bool, principals SecurityPolicyPrincipals, rules []SecurityPolicyRule) *SecurityPolicy {
	this := SecurityPolicy{}
	this.Name = name
	this.Description = description
	this.Active = active
	this.Principals = principals
	this.Rules = rules
	return &this
}

// NewSecurityPolicyWithDefaults instantiates a new SecurityPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityPolicyWithDefaults() *SecurityPolicy {
	this := SecurityPolicy{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SecurityPolicy) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityPolicy) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SecurityPolicy) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SecurityPolicy) SetId(v string) *SecurityPolicy {
	o.Id = &v
	return o
}

// GetName returns the Name field value
func (o *SecurityPolicy) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SecurityPolicy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SecurityPolicy) SetName(v string) *SecurityPolicy {
	o.Name = v
	return o
}

// GetDescription returns the Description field value
func (o *SecurityPolicy) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SecurityPolicy) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *SecurityPolicy) SetDescription(v string) *SecurityPolicy {
	o.Description = v
	return o
}

// GetActive returns the Active field value
func (o *SecurityPolicy) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *SecurityPolicy) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *SecurityPolicy) SetActive(v bool) *SecurityPolicy {
	o.Active = v
	return o
}

// GetPrincipals returns the Principals field value
func (o *SecurityPolicy) GetPrincipals() SecurityPolicyPrincipals {
	if o == nil {
		var ret SecurityPolicyPrincipals
		return ret
	}

	return o.Principals
}

// GetPrincipalsOk returns a tuple with the Principals field value
// and a boolean to check if the value has been set.
func (o *SecurityPolicy) GetPrincipalsOk() (*SecurityPolicyPrincipals, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Principals, true
}

// SetPrincipals sets field value
func (o *SecurityPolicy) SetPrincipals(v SecurityPolicyPrincipals) *SecurityPolicy {
	o.Principals = v
	return o
}

// GetRules returns the Rules field value
func (o *SecurityPolicy) GetRules() []SecurityPolicyRule {
	if o == nil {
		var ret []SecurityPolicyRule
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *SecurityPolicy) GetRulesOk() ([]SecurityPolicyRule, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rules, true
}

// SetRules sets field value
func (o *SecurityPolicy) SetRules(v []SecurityPolicyRule) *SecurityPolicy {
	o.Rules = v
	return o
}

func (o SecurityPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["active"] = o.Active
	toSerialize["principals"] = o.Principals
	toSerialize["rules"] = o.Rules
	return toSerialize, nil
}

type NullableSecurityPolicy struct {
	value *SecurityPolicy
	isSet bool
}

func (v NullableSecurityPolicy) Get() *SecurityPolicy {
	return v.value
}

func (v *NullableSecurityPolicy) Set(val *SecurityPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityPolicy(val *SecurityPolicy) *NullableSecurityPolicy {
	return &NullableSecurityPolicy{value: val, isSet: true}
}

func (v NullableSecurityPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
