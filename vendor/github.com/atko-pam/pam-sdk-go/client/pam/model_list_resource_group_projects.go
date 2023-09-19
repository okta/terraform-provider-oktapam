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

// checks if the ListResourceGroupProjects type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListResourceGroupProjects{}

// ListResourceGroupProjects struct for ListResourceGroupProjects
type ListResourceGroupProjects struct {
	List []Project `json:"list,omitempty"`
}

// NewListResourceGroupProjects instantiates a new ListResourceGroupProjects object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListResourceGroupProjects() *ListResourceGroupProjects {
	this := ListResourceGroupProjects{}
	return &this
}

// NewListResourceGroupProjectsWithDefaults instantiates a new ListResourceGroupProjects object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListResourceGroupProjectsWithDefaults() *ListResourceGroupProjects {
	this := ListResourceGroupProjects{}
	return &this
}

// GetList returns the List field value if set, zero value otherwise.
func (o *ListResourceGroupProjects) GetList() []Project {
	if o == nil || IsNil(o.List) {
		var ret []Project
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListResourceGroupProjects) GetListOk() ([]Project, bool) {
	if o == nil || IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *ListResourceGroupProjects) HasList() bool {
	if o != nil && !IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given []Project and assigns it to the List field.
func (o *ListResourceGroupProjects) SetList(v []Project) *ListResourceGroupProjects {
	o.List = v
	return o
}

func (o ListResourceGroupProjects) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListResourceGroupProjects) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.List) {
		toSerialize["list"] = o.List
	}
	return toSerialize, nil
}

type NullableListResourceGroupProjects struct {
	value *ListResourceGroupProjects
	isSet bool
}

func (v NullableListResourceGroupProjects) Get() *ListResourceGroupProjects {
	return v.value
}

func (v *NullableListResourceGroupProjects) Set(val *ListResourceGroupProjects) {
	v.value = val
	v.isSet = true
}

func (v NullableListResourceGroupProjects) IsSet() bool {
	return v.isSet
}

func (v *NullableListResourceGroupProjects) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListResourceGroupProjects(val *ListResourceGroupProjects) *NullableListResourceGroupProjects {
	return &NullableListResourceGroupProjects{value: val, isSet: true}
}

func (v NullableListResourceGroupProjects) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListResourceGroupProjects) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
