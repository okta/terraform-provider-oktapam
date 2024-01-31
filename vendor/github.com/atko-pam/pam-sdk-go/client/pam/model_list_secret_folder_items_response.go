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

// checks if the ListSecretFolderItemsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSecretFolderItemsResponse{}

// ListSecretFolderItemsResponse struct for ListSecretFolderItemsResponse
type ListSecretFolderItemsResponse struct {
	List []SecretOrFolderListResponse `json:"list,omitempty"`
}

// NewListSecretFolderItemsResponse instantiates a new ListSecretFolderItemsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSecretFolderItemsResponse() *ListSecretFolderItemsResponse {
	this := ListSecretFolderItemsResponse{}
	return &this
}

// NewListSecretFolderItemsResponseWithDefaults instantiates a new ListSecretFolderItemsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSecretFolderItemsResponseWithDefaults() *ListSecretFolderItemsResponse {
	this := ListSecretFolderItemsResponse{}
	return &this
}

// GetList returns the List field value if set, zero value otherwise.
func (o *ListSecretFolderItemsResponse) GetList() []SecretOrFolderListResponse {
	if o == nil || IsNil(o.List) {
		var ret []SecretOrFolderListResponse
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSecretFolderItemsResponse) GetListOk() ([]SecretOrFolderListResponse, bool) {
	if o == nil || IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *ListSecretFolderItemsResponse) HasList() bool {
	if o != nil && !IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given []SecretOrFolderListResponse and assigns it to the List field.
func (o *ListSecretFolderItemsResponse) SetList(v []SecretOrFolderListResponse) *ListSecretFolderItemsResponse {
	o.List = v
	return o
}

func (o ListSecretFolderItemsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListSecretFolderItemsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.List) {
		toSerialize["list"] = o.List
	}
	return toSerialize, nil
}

type NullableListSecretFolderItemsResponse struct {
	value *ListSecretFolderItemsResponse
	isSet bool
}

func (v NullableListSecretFolderItemsResponse) Get() *ListSecretFolderItemsResponse {
	return v.value
}

func (v *NullableListSecretFolderItemsResponse) Set(val *ListSecretFolderItemsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListSecretFolderItemsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListSecretFolderItemsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSecretFolderItemsResponse(val *ListSecretFolderItemsResponse) *NullableListSecretFolderItemsResponse {
	return &NullableListSecretFolderItemsResponse{value: val, isSet: true}
}

func (v NullableListSecretFolderItemsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSecretFolderItemsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
