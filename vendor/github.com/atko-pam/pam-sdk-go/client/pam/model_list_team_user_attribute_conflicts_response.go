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

// checks if the ListTeamUserAttributeConflictsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListTeamUserAttributeConflictsResponse{}

// ListTeamUserAttributeConflictsResponse struct for ListTeamUserAttributeConflictsResponse
type ListTeamUserAttributeConflictsResponse struct {
	List []TeamUserAttributeConflict `json:"list,omitempty"`
}

// NewListTeamUserAttributeConflictsResponse instantiates a new ListTeamUserAttributeConflictsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTeamUserAttributeConflictsResponse() *ListTeamUserAttributeConflictsResponse {
	this := ListTeamUserAttributeConflictsResponse{}
	return &this
}

// NewListTeamUserAttributeConflictsResponseWithDefaults instantiates a new ListTeamUserAttributeConflictsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTeamUserAttributeConflictsResponseWithDefaults() *ListTeamUserAttributeConflictsResponse {
	this := ListTeamUserAttributeConflictsResponse{}
	return &this
}

// GetList returns the List field value if set, zero value otherwise.
func (o *ListTeamUserAttributeConflictsResponse) GetList() []TeamUserAttributeConflict {
	if o == nil || IsNil(o.List) {
		var ret []TeamUserAttributeConflict
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTeamUserAttributeConflictsResponse) GetListOk() ([]TeamUserAttributeConflict, bool) {
	if o == nil || IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *ListTeamUserAttributeConflictsResponse) HasList() bool {
	if o != nil && !IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given []TeamUserAttributeConflict and assigns it to the List field.
func (o *ListTeamUserAttributeConflictsResponse) SetList(v []TeamUserAttributeConflict) *ListTeamUserAttributeConflictsResponse {
	o.List = v
	return o
}

func (o ListTeamUserAttributeConflictsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListTeamUserAttributeConflictsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.List) {
		toSerialize["list"] = o.List
	}
	return toSerialize, nil
}

type NullableListTeamUserAttributeConflictsResponse struct {
	value *ListTeamUserAttributeConflictsResponse
	isSet bool
}

func (v NullableListTeamUserAttributeConflictsResponse) Get() *ListTeamUserAttributeConflictsResponse {
	return v.value
}

func (v *NullableListTeamUserAttributeConflictsResponse) Set(val *ListTeamUserAttributeConflictsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListTeamUserAttributeConflictsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListTeamUserAttributeConflictsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTeamUserAttributeConflictsResponse(val *ListTeamUserAttributeConflictsResponse) *NullableListTeamUserAttributeConflictsResponse {
	return &NullableListTeamUserAttributeConflictsResponse{value: val, isSet: true}
}

func (v NullableListTeamUserAttributeConflictsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTeamUserAttributeConflictsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}