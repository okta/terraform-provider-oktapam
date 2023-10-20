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

// checks if the ListGroupsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListGroupsResponse{}

// ListGroupsResponse struct for ListGroupsResponse
type ListGroupsResponse struct {
	List []Group `json:"list,omitempty"`
}

// NewListGroupsResponse instantiates a new ListGroupsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListGroupsResponse() *ListGroupsResponse {
	this := ListGroupsResponse{}
	return &this
}

// NewListGroupsResponseWithDefaults instantiates a new ListGroupsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListGroupsResponseWithDefaults() *ListGroupsResponse {
	this := ListGroupsResponse{}
	return &this
}

// GetList returns the List field value if set, zero value otherwise.
func (o *ListGroupsResponse) GetList() []Group {
	if o == nil || IsNil(o.List) {
		var ret []Group
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListGroupsResponse) GetListOk() ([]Group, bool) {
	if o == nil || IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *ListGroupsResponse) HasList() bool {
	if o != nil && !IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given []Group and assigns it to the List field.
func (o *ListGroupsResponse) SetList(v []Group) *ListGroupsResponse {
	o.List = v
	return o
}

func (o ListGroupsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListGroupsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.List) {
		toSerialize["list"] = o.List
	}
	return toSerialize, nil
}

type NullableListGroupsResponse struct {
	value *ListGroupsResponse
	isSet bool
}

func (v NullableListGroupsResponse) Get() *ListGroupsResponse {
	return v.value
}

func (v *NullableListGroupsResponse) Set(val *ListGroupsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListGroupsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListGroupsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListGroupsResponse(val *ListGroupsResponse) *NullableListGroupsResponse {
	return &NullableListGroupsResponse{value: val, isSet: true}
}

func (v NullableListGroupsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListGroupsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
