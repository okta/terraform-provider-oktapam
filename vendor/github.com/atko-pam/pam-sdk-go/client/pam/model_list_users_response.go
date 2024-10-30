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

// checks if the ListUsersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListUsersResponse{}

// ListUsersResponse struct for ListUsersResponse
type ListUsersResponse struct {
	List []User `json:"list,omitempty"`
}

// NewListUsersResponse instantiates a new ListUsersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListUsersResponse() *ListUsersResponse {
	this := ListUsersResponse{}
	return &this
}

// NewListUsersResponseWithDefaults instantiates a new ListUsersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListUsersResponseWithDefaults() *ListUsersResponse {
	this := ListUsersResponse{}
	return &this
}

// GetList returns the List field value if set, zero value otherwise.
func (o *ListUsersResponse) GetList() []User {
	if o == nil || IsNil(o.List) {
		var ret []User
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUsersResponse) GetListOk() ([]User, bool) {
	if o == nil || IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *ListUsersResponse) HasList() bool {
	if o != nil && !IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given []User and assigns it to the List field.
func (o *ListUsersResponse) SetList(v []User) *ListUsersResponse {
	o.List = v
	return o
}

func (o ListUsersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListUsersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.List) {
		toSerialize["list"] = o.List
	}
	return toSerialize, nil
}

type NullableListUsersResponse struct {
	value *ListUsersResponse
	isSet bool
}

func (v NullableListUsersResponse) Get() *ListUsersResponse {
	return v.value
}

func (v *NullableListUsersResponse) Set(val *ListUsersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListUsersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListUsersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListUsersResponse(val *ListUsersResponse) *NullableListUsersResponse {
	return &NullableListUsersResponse{value: val, isSet: true}
}

func (v NullableListUsersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListUsersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
