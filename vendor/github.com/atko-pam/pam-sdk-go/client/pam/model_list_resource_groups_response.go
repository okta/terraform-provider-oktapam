/*
Okta Privileged Access

The ScaleFT API is a control plane API for operations in Okta Privileged Access (formerly ScaleFT)

API version: 1.0.0
Contact: support@okta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pam

import (
	"encoding/json"
)

// checks if the ListResourceGroupsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListResourceGroupsResponse{}

// ListResourceGroupsResponse struct for ListResourceGroupsResponse
type ListResourceGroupsResponse struct {
	List []ListResourceGroup `json:"list,omitempty"`
}

// NewListResourceGroupsResponse instantiates a new ListResourceGroupsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListResourceGroupsResponse() *ListResourceGroupsResponse {
	this := ListResourceGroupsResponse{}
	return &this
}

// NewListResourceGroupsResponseWithDefaults instantiates a new ListResourceGroupsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListResourceGroupsResponseWithDefaults() *ListResourceGroupsResponse {
	this := ListResourceGroupsResponse{}
	return &this
}

// GetList returns the List field value if set, zero value otherwise.
func (o *ListResourceGroupsResponse) GetList() []ListResourceGroup {
	if o == nil || IsNil(o.List) {
		var ret []ListResourceGroup
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListResourceGroupsResponse) GetListOk() ([]ListResourceGroup, bool) {
	if o == nil || IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *ListResourceGroupsResponse) HasList() bool {
	if o != nil && !IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given []ListResourceGroup and assigns it to the List field.
func (o *ListResourceGroupsResponse) SetList(v []ListResourceGroup) *ListResourceGroupsResponse {
	o.List = v
	return o
}

func (o ListResourceGroupsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListResourceGroupsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.List) {
		toSerialize["list"] = o.List
	}
	return toSerialize, nil
}

type NullableListResourceGroupsResponse struct {
	value *ListResourceGroupsResponse
	isSet bool
}

func (v NullableListResourceGroupsResponse) Get() *ListResourceGroupsResponse {
	return v.value
}

func (v *NullableListResourceGroupsResponse) Set(val *ListResourceGroupsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListResourceGroupsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListResourceGroupsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListResourceGroupsResponse(val *ListResourceGroupsResponse) *NullableListResourceGroupsResponse {
	return &NullableListResourceGroupsResponse{value: val, isSet: true}
}

func (v NullableListResourceGroupsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListResourceGroupsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}