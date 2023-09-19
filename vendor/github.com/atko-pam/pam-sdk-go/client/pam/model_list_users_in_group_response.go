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

// checks if the ListUsersInGroupResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListUsersInGroupResponse{}

// ListUsersInGroupResponse struct for ListUsersInGroupResponse
type ListUsersInGroupResponse struct {
	List []User `json:"list,omitempty"`
}

// NewListUsersInGroupResponse instantiates a new ListUsersInGroupResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListUsersInGroupResponse() *ListUsersInGroupResponse {
	this := ListUsersInGroupResponse{}
	return &this
}

// NewListUsersInGroupResponseWithDefaults instantiates a new ListUsersInGroupResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListUsersInGroupResponseWithDefaults() *ListUsersInGroupResponse {
	this := ListUsersInGroupResponse{}
	return &this
}

// GetList returns the List field value if set, zero value otherwise.
func (o *ListUsersInGroupResponse) GetList() []User {
	if o == nil || IsNil(o.List) {
		var ret []User
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUsersInGroupResponse) GetListOk() ([]User, bool) {
	if o == nil || IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *ListUsersInGroupResponse) HasList() bool {
	if o != nil && !IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given []User and assigns it to the List field.
func (o *ListUsersInGroupResponse) SetList(v []User) *ListUsersInGroupResponse {
	o.List = v
	return o
}

func (o ListUsersInGroupResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListUsersInGroupResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.List) {
		toSerialize["list"] = o.List
	}
	return toSerialize, nil
}

type NullableListUsersInGroupResponse struct {
	value *ListUsersInGroupResponse
	isSet bool
}

func (v NullableListUsersInGroupResponse) Get() *ListUsersInGroupResponse {
	return v.value
}

func (v *NullableListUsersInGroupResponse) Set(val *ListUsersInGroupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListUsersInGroupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListUsersInGroupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListUsersInGroupResponse(val *ListUsersInGroupResponse) *NullableListUsersInGroupResponse {
	return &NullableListUsersInGroupResponse{value: val, isSet: true}
}

func (v NullableListUsersInGroupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListUsersInGroupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
