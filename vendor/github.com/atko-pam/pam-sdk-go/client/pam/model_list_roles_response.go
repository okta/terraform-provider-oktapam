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

// checks if the ListRolesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListRolesResponse{}

// ListRolesResponse struct for ListRolesResponse
type ListRolesResponse struct {
	List []Roles `json:"list,omitempty"`
}

// NewListRolesResponse instantiates a new ListRolesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListRolesResponse() *ListRolesResponse {
	this := ListRolesResponse{}
	return &this
}

// NewListRolesResponseWithDefaults instantiates a new ListRolesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListRolesResponseWithDefaults() *ListRolesResponse {
	this := ListRolesResponse{}
	return &this
}

// GetList returns the List field value if set, zero value otherwise.
func (o *ListRolesResponse) GetList() []Roles {
	if o == nil || IsNil(o.List) {
		var ret []Roles
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRolesResponse) GetListOk() ([]Roles, bool) {
	if o == nil || IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *ListRolesResponse) HasList() bool {
	if o != nil && !IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given []Roles and assigns it to the List field.
func (o *ListRolesResponse) SetList(v []Roles) *ListRolesResponse {
	o.List = v
	return o
}

func (o ListRolesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListRolesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.List) {
		toSerialize["list"] = o.List
	}
	return toSerialize, nil
}

type NullableListRolesResponse struct {
	value *ListRolesResponse
	isSet bool
}

func (v NullableListRolesResponse) Get() *ListRolesResponse {
	return v.value
}

func (v *NullableListRolesResponse) Set(val *ListRolesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListRolesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListRolesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListRolesResponse(val *ListRolesResponse) *NullableListRolesResponse {
	return &NullableListRolesResponse{value: val, isSet: true}
}

func (v NullableListRolesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListRolesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
