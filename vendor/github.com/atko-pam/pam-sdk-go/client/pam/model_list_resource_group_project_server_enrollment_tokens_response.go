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

// checks if the ListResourceGroupProjectServerEnrollmentTokensResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListResourceGroupProjectServerEnrollmentTokensResponse{}

// ListResourceGroupProjectServerEnrollmentTokensResponse struct for ListResourceGroupProjectServerEnrollmentTokensResponse
type ListResourceGroupProjectServerEnrollmentTokensResponse struct {
	List []ServerEnrollmentToken `json:"list,omitempty"`
}

// NewListResourceGroupProjectServerEnrollmentTokensResponse instantiates a new ListResourceGroupProjectServerEnrollmentTokensResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListResourceGroupProjectServerEnrollmentTokensResponse() *ListResourceGroupProjectServerEnrollmentTokensResponse {
	this := ListResourceGroupProjectServerEnrollmentTokensResponse{}
	return &this
}

// NewListResourceGroupProjectServerEnrollmentTokensResponseWithDefaults instantiates a new ListResourceGroupProjectServerEnrollmentTokensResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListResourceGroupProjectServerEnrollmentTokensResponseWithDefaults() *ListResourceGroupProjectServerEnrollmentTokensResponse {
	this := ListResourceGroupProjectServerEnrollmentTokensResponse{}
	return &this
}

// GetList returns the List field value if set, zero value otherwise.
func (o *ListResourceGroupProjectServerEnrollmentTokensResponse) GetList() []ServerEnrollmentToken {
	if o == nil || IsNil(o.List) {
		var ret []ServerEnrollmentToken
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListResourceGroupProjectServerEnrollmentTokensResponse) GetListOk() ([]ServerEnrollmentToken, bool) {
	if o == nil || IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *ListResourceGroupProjectServerEnrollmentTokensResponse) HasList() bool {
	if o != nil && !IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given []ServerEnrollmentToken and assigns it to the List field.
func (o *ListResourceGroupProjectServerEnrollmentTokensResponse) SetList(v []ServerEnrollmentToken) *ListResourceGroupProjectServerEnrollmentTokensResponse {
	o.List = v
	return o
}

func (o ListResourceGroupProjectServerEnrollmentTokensResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListResourceGroupProjectServerEnrollmentTokensResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.List) {
		toSerialize["list"] = o.List
	}
	return toSerialize, nil
}

type NullableListResourceGroupProjectServerEnrollmentTokensResponse struct {
	value *ListResourceGroupProjectServerEnrollmentTokensResponse
	isSet bool
}

func (v NullableListResourceGroupProjectServerEnrollmentTokensResponse) Get() *ListResourceGroupProjectServerEnrollmentTokensResponse {
	return v.value
}

func (v *NullableListResourceGroupProjectServerEnrollmentTokensResponse) Set(val *ListResourceGroupProjectServerEnrollmentTokensResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListResourceGroupProjectServerEnrollmentTokensResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListResourceGroupProjectServerEnrollmentTokensResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListResourceGroupProjectServerEnrollmentTokensResponse(val *ListResourceGroupProjectServerEnrollmentTokensResponse) *NullableListResourceGroupProjectServerEnrollmentTokensResponse {
	return &NullableListResourceGroupProjectServerEnrollmentTokensResponse{value: val, isSet: true}
}

func (v NullableListResourceGroupProjectServerEnrollmentTokensResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListResourceGroupProjectServerEnrollmentTokensResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
