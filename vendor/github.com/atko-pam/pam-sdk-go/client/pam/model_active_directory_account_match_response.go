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

// checks if the ActiveDirectoryAccountMatchResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActiveDirectoryAccountMatchResponse{}

// ActiveDirectoryAccountMatchResponse Details about the mapping between an Active Directory account and an Okta Privileged Access user
type ActiveDirectoryAccountMatchResponse struct {
	AdAccountId *string `json:"ad_account_id,omitempty"`
	// -| The Okta Privileged Access user ID of the Okta Privileged Access user automatically matched to this AD account based on the Active Directory account rules configured on the connection. If null, the AD user did not match any Okta user based on the rule configuration the last time the rules were evaluated
	RuleMatchUserId *string `json:"rule_match_user_id,omitempty"`
	// -| The Okta Privileged Access user ID of the Okta Privileged Access user specified as matching this AD account by a Resource Administrator, overriding the automatic match. If set, this takes precedence over the rule_match_user_name. Only one of match_override_user_id or force_no_match may be set
	MatchOverrideUserId *string `json:"match_override_user_id,omitempty"`
	// -| If true, a resource administrator has configured this Active Directory account to not be matched with any Okta Privileged Access user. This overrides any automatic match from an Active Directory account rule. Only one of match_override_user_id or force_no_match may be set
	ForceNoMatch *bool `json:"force_no_match,omitempty"`
}

// NewActiveDirectoryAccountMatchResponse instantiates a new ActiveDirectoryAccountMatchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActiveDirectoryAccountMatchResponse() *ActiveDirectoryAccountMatchResponse {
	this := ActiveDirectoryAccountMatchResponse{}
	return &this
}

// NewActiveDirectoryAccountMatchResponseWithDefaults instantiates a new ActiveDirectoryAccountMatchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActiveDirectoryAccountMatchResponseWithDefaults() *ActiveDirectoryAccountMatchResponse {
	this := ActiveDirectoryAccountMatchResponse{}
	return &this
}

// GetAdAccountId returns the AdAccountId field value if set, zero value otherwise.
func (o *ActiveDirectoryAccountMatchResponse) GetAdAccountId() string {
	if o == nil || IsNil(o.AdAccountId) {
		var ret string
		return ret
	}
	return *o.AdAccountId
}

// GetAdAccountIdOk returns a tuple with the AdAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveDirectoryAccountMatchResponse) GetAdAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AdAccountId) {
		return nil, false
	}
	return o.AdAccountId, true
}

// HasAdAccountId returns a boolean if a field has been set.
func (o *ActiveDirectoryAccountMatchResponse) HasAdAccountId() bool {
	if o != nil && !IsNil(o.AdAccountId) {
		return true
	}

	return false
}

// SetAdAccountId gets a reference to the given string and assigns it to the AdAccountId field.
func (o *ActiveDirectoryAccountMatchResponse) SetAdAccountId(v string) *ActiveDirectoryAccountMatchResponse {
	o.AdAccountId = &v
	return o
}

// GetRuleMatchUserId returns the RuleMatchUserId field value if set, zero value otherwise.
func (o *ActiveDirectoryAccountMatchResponse) GetRuleMatchUserId() string {
	if o == nil || IsNil(o.RuleMatchUserId) {
		var ret string
		return ret
	}
	return *o.RuleMatchUserId
}

// GetRuleMatchUserIdOk returns a tuple with the RuleMatchUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveDirectoryAccountMatchResponse) GetRuleMatchUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.RuleMatchUserId) {
		return nil, false
	}
	return o.RuleMatchUserId, true
}

// HasRuleMatchUserId returns a boolean if a field has been set.
func (o *ActiveDirectoryAccountMatchResponse) HasRuleMatchUserId() bool {
	if o != nil && !IsNil(o.RuleMatchUserId) {
		return true
	}

	return false
}

// SetRuleMatchUserId gets a reference to the given string and assigns it to the RuleMatchUserId field.
func (o *ActiveDirectoryAccountMatchResponse) SetRuleMatchUserId(v string) *ActiveDirectoryAccountMatchResponse {
	o.RuleMatchUserId = &v
	return o
}

// GetMatchOverrideUserId returns the MatchOverrideUserId field value if set, zero value otherwise.
func (o *ActiveDirectoryAccountMatchResponse) GetMatchOverrideUserId() string {
	if o == nil || IsNil(o.MatchOverrideUserId) {
		var ret string
		return ret
	}
	return *o.MatchOverrideUserId
}

// GetMatchOverrideUserIdOk returns a tuple with the MatchOverrideUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveDirectoryAccountMatchResponse) GetMatchOverrideUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.MatchOverrideUserId) {
		return nil, false
	}
	return o.MatchOverrideUserId, true
}

// HasMatchOverrideUserId returns a boolean if a field has been set.
func (o *ActiveDirectoryAccountMatchResponse) HasMatchOverrideUserId() bool {
	if o != nil && !IsNil(o.MatchOverrideUserId) {
		return true
	}

	return false
}

// SetMatchOverrideUserId gets a reference to the given string and assigns it to the MatchOverrideUserId field.
func (o *ActiveDirectoryAccountMatchResponse) SetMatchOverrideUserId(v string) *ActiveDirectoryAccountMatchResponse {
	o.MatchOverrideUserId = &v
	return o
}

// GetForceNoMatch returns the ForceNoMatch field value if set, zero value otherwise.
func (o *ActiveDirectoryAccountMatchResponse) GetForceNoMatch() bool {
	if o == nil || IsNil(o.ForceNoMatch) {
		var ret bool
		return ret
	}
	return *o.ForceNoMatch
}

// GetForceNoMatchOk returns a tuple with the ForceNoMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveDirectoryAccountMatchResponse) GetForceNoMatchOk() (*bool, bool) {
	if o == nil || IsNil(o.ForceNoMatch) {
		return nil, false
	}
	return o.ForceNoMatch, true
}

// HasForceNoMatch returns a boolean if a field has been set.
func (o *ActiveDirectoryAccountMatchResponse) HasForceNoMatch() bool {
	if o != nil && !IsNil(o.ForceNoMatch) {
		return true
	}

	return false
}

// SetForceNoMatch gets a reference to the given bool and assigns it to the ForceNoMatch field.
func (o *ActiveDirectoryAccountMatchResponse) SetForceNoMatch(v bool) *ActiveDirectoryAccountMatchResponse {
	o.ForceNoMatch = &v
	return o
}

func (o ActiveDirectoryAccountMatchResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActiveDirectoryAccountMatchResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdAccountId) {
		toSerialize["ad_account_id"] = o.AdAccountId
	}
	if !IsNil(o.RuleMatchUserId) {
		toSerialize["rule_match_user_id"] = o.RuleMatchUserId
	}
	if !IsNil(o.MatchOverrideUserId) {
		toSerialize["match_override_user_id"] = o.MatchOverrideUserId
	}
	if !IsNil(o.ForceNoMatch) {
		toSerialize["force_no_match"] = o.ForceNoMatch
	}
	return toSerialize, nil
}

type NullableActiveDirectoryAccountMatchResponse struct {
	value *ActiveDirectoryAccountMatchResponse
	isSet bool
}

func (v NullableActiveDirectoryAccountMatchResponse) Get() *ActiveDirectoryAccountMatchResponse {
	return v.value
}

func (v *NullableActiveDirectoryAccountMatchResponse) Set(val *ActiveDirectoryAccountMatchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableActiveDirectoryAccountMatchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableActiveDirectoryAccountMatchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActiveDirectoryAccountMatchResponse(val *ActiveDirectoryAccountMatchResponse) *NullableActiveDirectoryAccountMatchResponse {
	return &NullableActiveDirectoryAccountMatchResponse{value: val, isSet: true}
}

func (v NullableActiveDirectoryAccountMatchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActiveDirectoryAccountMatchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
