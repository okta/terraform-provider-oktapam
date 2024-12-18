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

// checks if the ListResourceGroupProjectOktaUniversalDirectoryAccountsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListResourceGroupProjectOktaUniversalDirectoryAccountsResponse{}

// ListResourceGroupProjectOktaUniversalDirectoryAccountsResponse struct for ListResourceGroupProjectOktaUniversalDirectoryAccountsResponse
type ListResourceGroupProjectOktaUniversalDirectoryAccountsResponse struct {
	List []OktaUniversalDirectoryAccountWithSettings `json:"list,omitempty"`
}

// NewListResourceGroupProjectOktaUniversalDirectoryAccountsResponse instantiates a new ListResourceGroupProjectOktaUniversalDirectoryAccountsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListResourceGroupProjectOktaUniversalDirectoryAccountsResponse() *ListResourceGroupProjectOktaUniversalDirectoryAccountsResponse {
	this := ListResourceGroupProjectOktaUniversalDirectoryAccountsResponse{}
	return &this
}

// NewListResourceGroupProjectOktaUniversalDirectoryAccountsResponseWithDefaults instantiates a new ListResourceGroupProjectOktaUniversalDirectoryAccountsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListResourceGroupProjectOktaUniversalDirectoryAccountsResponseWithDefaults() *ListResourceGroupProjectOktaUniversalDirectoryAccountsResponse {
	this := ListResourceGroupProjectOktaUniversalDirectoryAccountsResponse{}
	return &this
}

// GetList returns the List field value if set, zero value otherwise.
func (o *ListResourceGroupProjectOktaUniversalDirectoryAccountsResponse) GetList() []OktaUniversalDirectoryAccountWithSettings {
	if o == nil || IsNil(o.List) {
		var ret []OktaUniversalDirectoryAccountWithSettings
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListResourceGroupProjectOktaUniversalDirectoryAccountsResponse) GetListOk() ([]OktaUniversalDirectoryAccountWithSettings, bool) {
	if o == nil || IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *ListResourceGroupProjectOktaUniversalDirectoryAccountsResponse) HasList() bool {
	if o != nil && !IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given []OktaUniversalDirectoryAccountWithSettings and assigns it to the List field.
func (o *ListResourceGroupProjectOktaUniversalDirectoryAccountsResponse) SetList(v []OktaUniversalDirectoryAccountWithSettings) *ListResourceGroupProjectOktaUniversalDirectoryAccountsResponse {
	o.List = v
	return o
}

func (o ListResourceGroupProjectOktaUniversalDirectoryAccountsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListResourceGroupProjectOktaUniversalDirectoryAccountsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.List) {
		toSerialize["list"] = o.List
	}
	return toSerialize, nil
}

type NullableListResourceGroupProjectOktaUniversalDirectoryAccountsResponse struct {
	value *ListResourceGroupProjectOktaUniversalDirectoryAccountsResponse
	isSet bool
}

func (v NullableListResourceGroupProjectOktaUniversalDirectoryAccountsResponse) Get() *ListResourceGroupProjectOktaUniversalDirectoryAccountsResponse {
	return v.value
}

func (v *NullableListResourceGroupProjectOktaUniversalDirectoryAccountsResponse) Set(val *ListResourceGroupProjectOktaUniversalDirectoryAccountsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListResourceGroupProjectOktaUniversalDirectoryAccountsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListResourceGroupProjectOktaUniversalDirectoryAccountsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListResourceGroupProjectOktaUniversalDirectoryAccountsResponse(val *ListResourceGroupProjectOktaUniversalDirectoryAccountsResponse) *NullableListResourceGroupProjectOktaUniversalDirectoryAccountsResponse {
	return &NullableListResourceGroupProjectOktaUniversalDirectoryAccountsResponse{value: val, isSet: true}
}

func (v NullableListResourceGroupProjectOktaUniversalDirectoryAccountsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListResourceGroupProjectOktaUniversalDirectoryAccountsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
