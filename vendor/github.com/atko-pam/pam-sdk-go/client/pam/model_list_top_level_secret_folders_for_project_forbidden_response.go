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

// checks if the ListTopLevelSecretFoldersForProjectForbiddenResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListTopLevelSecretFoldersForProjectForbiddenResponse{}

// ListTopLevelSecretFoldersForProjectForbiddenResponse struct for ListTopLevelSecretFoldersForProjectForbiddenResponse
type ListTopLevelSecretFoldersForProjectForbiddenResponse struct {
	List []UserAccessMethod `json:"list,omitempty"`
}

// NewListTopLevelSecretFoldersForProjectForbiddenResponse instantiates a new ListTopLevelSecretFoldersForProjectForbiddenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTopLevelSecretFoldersForProjectForbiddenResponse() *ListTopLevelSecretFoldersForProjectForbiddenResponse {
	this := ListTopLevelSecretFoldersForProjectForbiddenResponse{}
	return &this
}

// NewListTopLevelSecretFoldersForProjectForbiddenResponseWithDefaults instantiates a new ListTopLevelSecretFoldersForProjectForbiddenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTopLevelSecretFoldersForProjectForbiddenResponseWithDefaults() *ListTopLevelSecretFoldersForProjectForbiddenResponse {
	this := ListTopLevelSecretFoldersForProjectForbiddenResponse{}
	return &this
}

// GetList returns the List field value if set, zero value otherwise.
func (o *ListTopLevelSecretFoldersForProjectForbiddenResponse) GetList() []UserAccessMethod {
	if o == nil || IsNil(o.List) {
		var ret []UserAccessMethod
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTopLevelSecretFoldersForProjectForbiddenResponse) GetListOk() ([]UserAccessMethod, bool) {
	if o == nil || IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *ListTopLevelSecretFoldersForProjectForbiddenResponse) HasList() bool {
	if o != nil && !IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given []UserAccessMethod and assigns it to the List field.
func (o *ListTopLevelSecretFoldersForProjectForbiddenResponse) SetList(v []UserAccessMethod) *ListTopLevelSecretFoldersForProjectForbiddenResponse {
	o.List = v
	return o
}

func (o ListTopLevelSecretFoldersForProjectForbiddenResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListTopLevelSecretFoldersForProjectForbiddenResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.List) {
		toSerialize["list"] = o.List
	}
	return toSerialize, nil
}

type NullableListTopLevelSecretFoldersForProjectForbiddenResponse struct {
	value *ListTopLevelSecretFoldersForProjectForbiddenResponse
	isSet bool
}

func (v NullableListTopLevelSecretFoldersForProjectForbiddenResponse) Get() *ListTopLevelSecretFoldersForProjectForbiddenResponse {
	return v.value
}

func (v *NullableListTopLevelSecretFoldersForProjectForbiddenResponse) Set(val *ListTopLevelSecretFoldersForProjectForbiddenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListTopLevelSecretFoldersForProjectForbiddenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListTopLevelSecretFoldersForProjectForbiddenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTopLevelSecretFoldersForProjectForbiddenResponse(val *ListTopLevelSecretFoldersForProjectForbiddenResponse) *NullableListTopLevelSecretFoldersForProjectForbiddenResponse {
	return &NullableListTopLevelSecretFoldersForProjectForbiddenResponse{value: val, isSet: true}
}

func (v NullableListTopLevelSecretFoldersForProjectForbiddenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTopLevelSecretFoldersForProjectForbiddenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
