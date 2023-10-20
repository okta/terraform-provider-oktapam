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

// checks if the ListServersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListServersResponse{}

// ListServersResponse struct for ListServersResponse
type ListServersResponse struct {
	List []Server `json:"list,omitempty"`
}

// NewListServersResponse instantiates a new ListServersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListServersResponse() *ListServersResponse {
	this := ListServersResponse{}
	return &this
}

// NewListServersResponseWithDefaults instantiates a new ListServersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListServersResponseWithDefaults() *ListServersResponse {
	this := ListServersResponse{}
	return &this
}

// GetList returns the List field value if set, zero value otherwise.
func (o *ListServersResponse) GetList() []Server {
	if o == nil || IsNil(o.List) {
		var ret []Server
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListServersResponse) GetListOk() ([]Server, bool) {
	if o == nil || IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *ListServersResponse) HasList() bool {
	if o != nil && !IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given []Server and assigns it to the List field.
func (o *ListServersResponse) SetList(v []Server) *ListServersResponse {
	o.List = v
	return o
}

func (o ListServersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListServersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.List) {
		toSerialize["list"] = o.List
	}
	return toSerialize, nil
}

type NullableListServersResponse struct {
	value *ListServersResponse
	isSet bool
}

func (v NullableListServersResponse) Get() *ListServersResponse {
	return v.value
}

func (v *NullableListServersResponse) Set(val *ListServersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListServersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListServersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListServersResponse(val *ListServersResponse) *NullableListServersResponse {
	return &NullableListServersResponse{value: val, isSet: true}
}

func (v NullableListServersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListServersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}