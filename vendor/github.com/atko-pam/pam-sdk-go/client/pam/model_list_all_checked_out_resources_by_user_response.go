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

// checks if the ListAllCheckedOutResourcesByUserResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListAllCheckedOutResourcesByUserResponse{}

// ListAllCheckedOutResourcesByUserResponse struct for ListAllCheckedOutResourcesByUserResponse
type ListAllCheckedOutResourcesByUserResponse struct {
	List []CheckedOutResourceByUserDetails `json:"list,omitempty"`
}

// NewListAllCheckedOutResourcesByUserResponse instantiates a new ListAllCheckedOutResourcesByUserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAllCheckedOutResourcesByUserResponse() *ListAllCheckedOutResourcesByUserResponse {
	this := ListAllCheckedOutResourcesByUserResponse{}
	return &this
}

// NewListAllCheckedOutResourcesByUserResponseWithDefaults instantiates a new ListAllCheckedOutResourcesByUserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAllCheckedOutResourcesByUserResponseWithDefaults() *ListAllCheckedOutResourcesByUserResponse {
	this := ListAllCheckedOutResourcesByUserResponse{}
	return &this
}

// GetList returns the List field value if set, zero value otherwise.
func (o *ListAllCheckedOutResourcesByUserResponse) GetList() []CheckedOutResourceByUserDetails {
	if o == nil || IsNil(o.List) {
		var ret []CheckedOutResourceByUserDetails
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAllCheckedOutResourcesByUserResponse) GetListOk() ([]CheckedOutResourceByUserDetails, bool) {
	if o == nil || IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *ListAllCheckedOutResourcesByUserResponse) HasList() bool {
	if o != nil && !IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given []CheckedOutResourceByUserDetails and assigns it to the List field.
func (o *ListAllCheckedOutResourcesByUserResponse) SetList(v []CheckedOutResourceByUserDetails) *ListAllCheckedOutResourcesByUserResponse {
	o.List = v
	return o
}

func (o ListAllCheckedOutResourcesByUserResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListAllCheckedOutResourcesByUserResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.List) {
		toSerialize["list"] = o.List
	}
	return toSerialize, nil
}

type NullableListAllCheckedOutResourcesByUserResponse struct {
	value *ListAllCheckedOutResourcesByUserResponse
	isSet bool
}

func (v NullableListAllCheckedOutResourcesByUserResponse) Get() *ListAllCheckedOutResourcesByUserResponse {
	return v.value
}

func (v *NullableListAllCheckedOutResourcesByUserResponse) Set(val *ListAllCheckedOutResourcesByUserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListAllCheckedOutResourcesByUserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListAllCheckedOutResourcesByUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAllCheckedOutResourcesByUserResponse(val *ListAllCheckedOutResourcesByUserResponse) *NullableListAllCheckedOutResourcesByUserResponse {
	return &NullableListAllCheckedOutResourcesByUserResponse{value: val, isSet: true}
}

func (v NullableListAllCheckedOutResourcesByUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAllCheckedOutResourcesByUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
