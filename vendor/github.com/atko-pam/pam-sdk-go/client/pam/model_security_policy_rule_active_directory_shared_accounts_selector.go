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

// checks if the SecurityPolicyRuleActiveDirectorySharedAccountsSelector type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityPolicyRuleActiveDirectorySharedAccountsSelector{}

// SecurityPolicyRuleActiveDirectorySharedAccountsSelector Principals can access Active Directory accounts which are matched to them via discovery rules
type SecurityPolicyRuleActiveDirectorySharedAccountsSelector struct {
	ByMatchingNames []string `json:"by_matching_names,omitempty"`
	// Allow access to a subset of the matched individual accounts they matched to
	ByDomain []string `json:"by_domain,omitempty"`
	// The specific accounts to apply the policy rule to
	SpecificAccounts []NamedObject `json:"specific_accounts,omitempty"`
}

// NewSecurityPolicyRuleActiveDirectorySharedAccountsSelector instantiates a new SecurityPolicyRuleActiveDirectorySharedAccountsSelector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityPolicyRuleActiveDirectorySharedAccountsSelector() *SecurityPolicyRuleActiveDirectorySharedAccountsSelector {
	this := SecurityPolicyRuleActiveDirectorySharedAccountsSelector{}
	return &this
}

// NewSecurityPolicyRuleActiveDirectorySharedAccountsSelectorWithDefaults instantiates a new SecurityPolicyRuleActiveDirectorySharedAccountsSelector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityPolicyRuleActiveDirectorySharedAccountsSelectorWithDefaults() *SecurityPolicyRuleActiveDirectorySharedAccountsSelector {
	this := SecurityPolicyRuleActiveDirectorySharedAccountsSelector{}
	return &this
}

// GetByMatchingNames returns the ByMatchingNames field value if set, zero value otherwise.
func (o *SecurityPolicyRuleActiveDirectorySharedAccountsSelector) GetByMatchingNames() []string {
	if o == nil || IsNil(o.ByMatchingNames) {
		var ret []string
		return ret
	}
	return o.ByMatchingNames
}

// GetByMatchingNamesOk returns a tuple with the ByMatchingNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityPolicyRuleActiveDirectorySharedAccountsSelector) GetByMatchingNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.ByMatchingNames) {
		return nil, false
	}
	return o.ByMatchingNames, true
}

// HasByMatchingNames returns a boolean if a field has been set.
func (o *SecurityPolicyRuleActiveDirectorySharedAccountsSelector) HasByMatchingNames() bool {
	if o != nil && !IsNil(o.ByMatchingNames) {
		return true
	}

	return false
}

// SetByMatchingNames gets a reference to the given []string and assigns it to the ByMatchingNames field.
func (o *SecurityPolicyRuleActiveDirectorySharedAccountsSelector) SetByMatchingNames(v []string) *SecurityPolicyRuleActiveDirectorySharedAccountsSelector {
	o.ByMatchingNames = v
	return o
}

// GetByDomain returns the ByDomain field value if set, zero value otherwise.
func (o *SecurityPolicyRuleActiveDirectorySharedAccountsSelector) GetByDomain() []string {
	if o == nil || IsNil(o.ByDomain) {
		var ret []string
		return ret
	}
	return o.ByDomain
}

// GetByDomainOk returns a tuple with the ByDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityPolicyRuleActiveDirectorySharedAccountsSelector) GetByDomainOk() ([]string, bool) {
	if o == nil || IsNil(o.ByDomain) {
		return nil, false
	}
	return o.ByDomain, true
}

// HasByDomain returns a boolean if a field has been set.
func (o *SecurityPolicyRuleActiveDirectorySharedAccountsSelector) HasByDomain() bool {
	if o != nil && !IsNil(o.ByDomain) {
		return true
	}

	return false
}

// SetByDomain gets a reference to the given []string and assigns it to the ByDomain field.
func (o *SecurityPolicyRuleActiveDirectorySharedAccountsSelector) SetByDomain(v []string) *SecurityPolicyRuleActiveDirectorySharedAccountsSelector {
	o.ByDomain = v
	return o
}

// GetSpecificAccounts returns the SpecificAccounts field value if set, zero value otherwise.
func (o *SecurityPolicyRuleActiveDirectorySharedAccountsSelector) GetSpecificAccounts() []NamedObject {
	if o == nil || IsNil(o.SpecificAccounts) {
		var ret []NamedObject
		return ret
	}
	return o.SpecificAccounts
}

// GetSpecificAccountsOk returns a tuple with the SpecificAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityPolicyRuleActiveDirectorySharedAccountsSelector) GetSpecificAccountsOk() ([]NamedObject, bool) {
	if o == nil || IsNil(o.SpecificAccounts) {
		return nil, false
	}
	return o.SpecificAccounts, true
}

// HasSpecificAccounts returns a boolean if a field has been set.
func (o *SecurityPolicyRuleActiveDirectorySharedAccountsSelector) HasSpecificAccounts() bool {
	if o != nil && !IsNil(o.SpecificAccounts) {
		return true
	}

	return false
}

// SetSpecificAccounts gets a reference to the given []NamedObject and assigns it to the SpecificAccounts field.
func (o *SecurityPolicyRuleActiveDirectorySharedAccountsSelector) SetSpecificAccounts(v []NamedObject) *SecurityPolicyRuleActiveDirectorySharedAccountsSelector {
	o.SpecificAccounts = v
	return o
}

func (o SecurityPolicyRuleActiveDirectorySharedAccountsSelector) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityPolicyRuleActiveDirectorySharedAccountsSelector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ByMatchingNames) {
		toSerialize["by_matching_names"] = o.ByMatchingNames
	}
	if !IsNil(o.ByDomain) {
		toSerialize["by_domain"] = o.ByDomain
	}
	if !IsNil(o.SpecificAccounts) {
		toSerialize["specific_accounts"] = o.SpecificAccounts
	}
	return toSerialize, nil
}

type NullableSecurityPolicyRuleActiveDirectorySharedAccountsSelector struct {
	value *SecurityPolicyRuleActiveDirectorySharedAccountsSelector
	isSet bool
}

func (v NullableSecurityPolicyRuleActiveDirectorySharedAccountsSelector) Get() *SecurityPolicyRuleActiveDirectorySharedAccountsSelector {
	return v.value
}

func (v *NullableSecurityPolicyRuleActiveDirectorySharedAccountsSelector) Set(val *SecurityPolicyRuleActiveDirectorySharedAccountsSelector) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityPolicyRuleActiveDirectorySharedAccountsSelector) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityPolicyRuleActiveDirectorySharedAccountsSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityPolicyRuleActiveDirectorySharedAccountsSelector(val *SecurityPolicyRuleActiveDirectorySharedAccountsSelector) *NullableSecurityPolicyRuleActiveDirectorySharedAccountsSelector {
	return &NullableSecurityPolicyRuleActiveDirectorySharedAccountsSelector{value: val, isSet: true}
}

func (v NullableSecurityPolicyRuleActiveDirectorySharedAccountsSelector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityPolicyRuleActiveDirectorySharedAccountsSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
