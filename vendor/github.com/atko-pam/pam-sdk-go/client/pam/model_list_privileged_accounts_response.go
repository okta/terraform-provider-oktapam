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

// checks if the ListPrivilegedAccountsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListPrivilegedAccountsResponse{}

// ListPrivilegedAccountsResponse struct for ListPrivilegedAccountsResponse
type ListPrivilegedAccountsResponse struct {
	List []PrivilegedAccount `json:"list,omitempty"`
}

// NewListPrivilegedAccountsResponse instantiates a new ListPrivilegedAccountsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPrivilegedAccountsResponse() *ListPrivilegedAccountsResponse {
	this := ListPrivilegedAccountsResponse{}
	return &this
}

// NewListPrivilegedAccountsResponseWithDefaults instantiates a new ListPrivilegedAccountsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPrivilegedAccountsResponseWithDefaults() *ListPrivilegedAccountsResponse {
	this := ListPrivilegedAccountsResponse{}
	return &this
}

// GetList returns the List field value if set, zero value otherwise.
func (o *ListPrivilegedAccountsResponse) GetList() []PrivilegedAccount {
	if o == nil || IsNil(o.List) {
		var ret []PrivilegedAccount
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPrivilegedAccountsResponse) GetListOk() ([]PrivilegedAccount, bool) {
	if o == nil || IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *ListPrivilegedAccountsResponse) HasList() bool {
	if o != nil && !IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given []PrivilegedAccount and assigns it to the List field.
func (o *ListPrivilegedAccountsResponse) SetList(v []PrivilegedAccount) *ListPrivilegedAccountsResponse {
	o.List = v
	return o
}

func (o ListPrivilegedAccountsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListPrivilegedAccountsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.List) {
		toSerialize["list"] = o.List
	}
	return toSerialize, nil
}

type NullableListPrivilegedAccountsResponse struct {
	value *ListPrivilegedAccountsResponse
	isSet bool
}

func (v NullableListPrivilegedAccountsResponse) Get() *ListPrivilegedAccountsResponse {
	return v.value
}

func (v *NullableListPrivilegedAccountsResponse) Set(val *ListPrivilegedAccountsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListPrivilegedAccountsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListPrivilegedAccountsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPrivilegedAccountsResponse(val *ListPrivilegedAccountsResponse) *NullableListPrivilegedAccountsResponse {
	return &NullableListPrivilegedAccountsResponse{value: val, isSet: true}
}

func (v NullableListPrivilegedAccountsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPrivilegedAccountsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
