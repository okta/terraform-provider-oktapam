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

// checks if the ListUAMForServiceAccountResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListUAMForServiceAccountResponse{}

// ListUAMForServiceAccountResponse List of `User Access Methods` for a service account
type ListUAMForServiceAccountResponse struct {
	List []UserAccessMethod `json:"list"`
}

// NewListUAMForServiceAccountResponse instantiates a new ListUAMForServiceAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListUAMForServiceAccountResponse(list []UserAccessMethod) *ListUAMForServiceAccountResponse {
	this := ListUAMForServiceAccountResponse{}
	this.List = list
	return &this
}

// NewListUAMForServiceAccountResponseWithDefaults instantiates a new ListUAMForServiceAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListUAMForServiceAccountResponseWithDefaults() *ListUAMForServiceAccountResponse {
	this := ListUAMForServiceAccountResponse{}
	return &this
}

// GetList returns the List field value
func (o *ListUAMForServiceAccountResponse) GetList() []UserAccessMethod {
	if o == nil {
		var ret []UserAccessMethod
		return ret
	}

	return o.List
}

// GetListOk returns a tuple with the List field value
// and a boolean to check if the value has been set.
func (o *ListUAMForServiceAccountResponse) GetListOk() ([]UserAccessMethod, bool) {
	if o == nil {
		return nil, false
	}
	return o.List, true
}

// SetList sets field value
func (o *ListUAMForServiceAccountResponse) SetList(v []UserAccessMethod) *ListUAMForServiceAccountResponse {
	o.List = v
	return o
}

func (o ListUAMForServiceAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListUAMForServiceAccountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["list"] = o.List
	return toSerialize, nil
}

type NullableListUAMForServiceAccountResponse struct {
	value *ListUAMForServiceAccountResponse
	isSet bool
}

func (v NullableListUAMForServiceAccountResponse) Get() *ListUAMForServiceAccountResponse {
	return v.value
}

func (v *NullableListUAMForServiceAccountResponse) Set(val *ListUAMForServiceAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListUAMForServiceAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListUAMForServiceAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListUAMForServiceAccountResponse(val *ListUAMForServiceAccountResponse) *NullableListUAMForServiceAccountResponse {
	return &NullableListUAMForServiceAccountResponse{value: val, isSet: true}
}

func (v NullableListUAMForServiceAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListUAMForServiceAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
