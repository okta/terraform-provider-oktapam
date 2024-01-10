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

// checks if the ListUserGroupsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListUserGroupsResponse{}

// ListUserGroupsResponse struct for ListUserGroupsResponse
type ListUserGroupsResponse struct {
	List []Group `json:"list,omitempty"`
}

// NewListUserGroupsResponse instantiates a new ListUserGroupsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListUserGroupsResponse() *ListUserGroupsResponse {
	this := ListUserGroupsResponse{}
	return &this
}

// NewListUserGroupsResponseWithDefaults instantiates a new ListUserGroupsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListUserGroupsResponseWithDefaults() *ListUserGroupsResponse {
	this := ListUserGroupsResponse{}
	return &this
}

// GetList returns the List field value if set, zero value otherwise.
func (o *ListUserGroupsResponse) GetList() []Group {
	if o == nil || IsNil(o.List) {
		var ret []Group
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUserGroupsResponse) GetListOk() ([]Group, bool) {
	if o == nil || IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *ListUserGroupsResponse) HasList() bool {
	if o != nil && !IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given []Group and assigns it to the List field.
func (o *ListUserGroupsResponse) SetList(v []Group) *ListUserGroupsResponse {
	o.List = v
	return o
}

func (o ListUserGroupsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListUserGroupsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.List) {
		toSerialize["list"] = o.List
	}
	return toSerialize, nil
}

type NullableListUserGroupsResponse struct {
	value *ListUserGroupsResponse
	isSet bool
}

func (v NullableListUserGroupsResponse) Get() *ListUserGroupsResponse {
	return v.value
}

func (v *NullableListUserGroupsResponse) Set(val *ListUserGroupsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListUserGroupsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListUserGroupsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListUserGroupsResponse(val *ListUserGroupsResponse) *NullableListUserGroupsResponse {
	return &NullableListUserGroupsResponse{value: val, isSet: true}
}

func (v NullableListUserGroupsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListUserGroupsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
