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

// checks if the SecurityPolicyRuleActiveDirectoryIndividualAccountsSelectorByCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityPolicyRuleActiveDirectoryIndividualAccountsSelectorByCondition{}

// SecurityPolicyRuleActiveDirectoryIndividualAccountsSelectorByCondition Allow access to a subset of the matched individual accounts they matched to
type SecurityPolicyRuleActiveDirectoryIndividualAccountsSelectorByCondition struct {
	// The format of the account name
	AccountNameFormat string `json:"account_name_format"`
	// The value of the account name
	Value string `json:"value"`
}

// NewSecurityPolicyRuleActiveDirectoryIndividualAccountsSelectorByCondition instantiates a new SecurityPolicyRuleActiveDirectoryIndividualAccountsSelectorByCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityPolicyRuleActiveDirectoryIndividualAccountsSelectorByCondition(accountNameFormat string, value string) *SecurityPolicyRuleActiveDirectoryIndividualAccountsSelectorByCondition {
	this := SecurityPolicyRuleActiveDirectoryIndividualAccountsSelectorByCondition{}
	this.AccountNameFormat = accountNameFormat
	this.Value = value
	return &this
}

// NewSecurityPolicyRuleActiveDirectoryIndividualAccountsSelectorByConditionWithDefaults instantiates a new SecurityPolicyRuleActiveDirectoryIndividualAccountsSelectorByCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityPolicyRuleActiveDirectoryIndividualAccountsSelectorByConditionWithDefaults() *SecurityPolicyRuleActiveDirectoryIndividualAccountsSelectorByCondition {
	this := SecurityPolicyRuleActiveDirectoryIndividualAccountsSelectorByCondition{}
	return &this
}

// GetAccountNameFormat returns the AccountNameFormat field value
func (o *SecurityPolicyRuleActiveDirectoryIndividualAccountsSelectorByCondition) GetAccountNameFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountNameFormat
}

// GetAccountNameFormatOk returns a tuple with the AccountNameFormat field value
// and a boolean to check if the value has been set.
func (o *SecurityPolicyRuleActiveDirectoryIndividualAccountsSelectorByCondition) GetAccountNameFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountNameFormat, true
}

// SetAccountNameFormat sets field value
func (o *SecurityPolicyRuleActiveDirectoryIndividualAccountsSelectorByCondition) SetAccountNameFormat(v string) *SecurityPolicyRuleActiveDirectoryIndividualAccountsSelectorByCondition {
	o.AccountNameFormat = v
	return o
}

// GetValue returns the Value field value
func (o *SecurityPolicyRuleActiveDirectoryIndividualAccountsSelectorByCondition) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *SecurityPolicyRuleActiveDirectoryIndividualAccountsSelectorByCondition) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *SecurityPolicyRuleActiveDirectoryIndividualAccountsSelectorByCondition) SetValue(v string) *SecurityPolicyRuleActiveDirectoryIndividualAccountsSelectorByCondition {
	o.Value = v
	return o
}

func (o SecurityPolicyRuleActiveDirectoryIndividualAccountsSelectorByCondition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityPolicyRuleActiveDirectoryIndividualAccountsSelectorByCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_name_format"] = o.AccountNameFormat
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableSecurityPolicyRuleActiveDirectoryIndividualAccountsSelectorByCondition struct {
	value *SecurityPolicyRuleActiveDirectoryIndividualAccountsSelectorByCondition
	isSet bool
}

func (v NullableSecurityPolicyRuleActiveDirectoryIndividualAccountsSelectorByCondition) Get() *SecurityPolicyRuleActiveDirectoryIndividualAccountsSelectorByCondition {
	return v.value
}

func (v *NullableSecurityPolicyRuleActiveDirectoryIndividualAccountsSelectorByCondition) Set(val *SecurityPolicyRuleActiveDirectoryIndividualAccountsSelectorByCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityPolicyRuleActiveDirectoryIndividualAccountsSelectorByCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityPolicyRuleActiveDirectoryIndividualAccountsSelectorByCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityPolicyRuleActiveDirectoryIndividualAccountsSelectorByCondition(val *SecurityPolicyRuleActiveDirectoryIndividualAccountsSelectorByCondition) *NullableSecurityPolicyRuleActiveDirectoryIndividualAccountsSelectorByCondition {
	return &NullableSecurityPolicyRuleActiveDirectoryIndividualAccountsSelectorByCondition{value: val, isSet: true}
}

func (v NullableSecurityPolicyRuleActiveDirectoryIndividualAccountsSelectorByCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityPolicyRuleActiveDirectoryIndividualAccountsSelectorByCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
