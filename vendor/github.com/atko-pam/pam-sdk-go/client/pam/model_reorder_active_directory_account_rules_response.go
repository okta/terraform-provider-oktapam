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

// checks if the ReorderActiveDirectoryAccountRulesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReorderActiveDirectoryAccountRulesResponse{}

// ReorderActiveDirectoryAccountRulesResponse struct for ReorderActiveDirectoryAccountRulesResponse
type ReorderActiveDirectoryAccountRulesResponse struct {
	List []ActiveDirectoryAccountRuleResponse `json:"list,omitempty"`
}

// NewReorderActiveDirectoryAccountRulesResponse instantiates a new ReorderActiveDirectoryAccountRulesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReorderActiveDirectoryAccountRulesResponse() *ReorderActiveDirectoryAccountRulesResponse {
	this := ReorderActiveDirectoryAccountRulesResponse{}
	return &this
}

// NewReorderActiveDirectoryAccountRulesResponseWithDefaults instantiates a new ReorderActiveDirectoryAccountRulesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReorderActiveDirectoryAccountRulesResponseWithDefaults() *ReorderActiveDirectoryAccountRulesResponse {
	this := ReorderActiveDirectoryAccountRulesResponse{}
	return &this
}

// GetList returns the List field value if set, zero value otherwise.
func (o *ReorderActiveDirectoryAccountRulesResponse) GetList() []ActiveDirectoryAccountRuleResponse {
	if o == nil || IsNil(o.List) {
		var ret []ActiveDirectoryAccountRuleResponse
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReorderActiveDirectoryAccountRulesResponse) GetListOk() ([]ActiveDirectoryAccountRuleResponse, bool) {
	if o == nil || IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *ReorderActiveDirectoryAccountRulesResponse) HasList() bool {
	if o != nil && !IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given []ActiveDirectoryAccountRuleResponse and assigns it to the List field.
func (o *ReorderActiveDirectoryAccountRulesResponse) SetList(v []ActiveDirectoryAccountRuleResponse) *ReorderActiveDirectoryAccountRulesResponse {
	o.List = v
	return o
}

func (o ReorderActiveDirectoryAccountRulesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReorderActiveDirectoryAccountRulesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.List) {
		toSerialize["list"] = o.List
	}
	return toSerialize, nil
}

type NullableReorderActiveDirectoryAccountRulesResponse struct {
	value *ReorderActiveDirectoryAccountRulesResponse
	isSet bool
}

func (v NullableReorderActiveDirectoryAccountRulesResponse) Get() *ReorderActiveDirectoryAccountRulesResponse {
	return v.value
}

func (v *NullableReorderActiveDirectoryAccountRulesResponse) Set(val *ReorderActiveDirectoryAccountRulesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReorderActiveDirectoryAccountRulesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReorderActiveDirectoryAccountRulesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReorderActiveDirectoryAccountRulesResponse(val *ReorderActiveDirectoryAccountRulesResponse) *NullableReorderActiveDirectoryAccountRulesResponse {
	return &NullableReorderActiveDirectoryAccountRulesResponse{value: val, isSet: true}
}

func (v NullableReorderActiveDirectoryAccountRulesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReorderActiveDirectoryAccountRulesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
