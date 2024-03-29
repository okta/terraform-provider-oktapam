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

// checks if the ListCloudConnectionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCloudConnectionsResponse{}

// ListCloudConnectionsResponse struct for ListCloudConnectionsResponse
type ListCloudConnectionsResponse struct {
	List []CloudConnection `json:"list,omitempty"`
}

// NewListCloudConnectionsResponse instantiates a new ListCloudConnectionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCloudConnectionsResponse() *ListCloudConnectionsResponse {
	this := ListCloudConnectionsResponse{}
	return &this
}

// NewListCloudConnectionsResponseWithDefaults instantiates a new ListCloudConnectionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCloudConnectionsResponseWithDefaults() *ListCloudConnectionsResponse {
	this := ListCloudConnectionsResponse{}
	return &this
}

// GetList returns the List field value if set, zero value otherwise.
func (o *ListCloudConnectionsResponse) GetList() []CloudConnection {
	if o == nil || IsNil(o.List) {
		var ret []CloudConnection
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCloudConnectionsResponse) GetListOk() ([]CloudConnection, bool) {
	if o == nil || IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *ListCloudConnectionsResponse) HasList() bool {
	if o != nil && !IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given []CloudConnection and assigns it to the List field.
func (o *ListCloudConnectionsResponse) SetList(v []CloudConnection) *ListCloudConnectionsResponse {
	o.List = v
	return o
}

func (o ListCloudConnectionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListCloudConnectionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.List) {
		toSerialize["list"] = o.List
	}
	return toSerialize, nil
}

type NullableListCloudConnectionsResponse struct {
	value *ListCloudConnectionsResponse
	isSet bool
}

func (v NullableListCloudConnectionsResponse) Get() *ListCloudConnectionsResponse {
	return v.value
}

func (v *NullableListCloudConnectionsResponse) Set(val *ListCloudConnectionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListCloudConnectionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListCloudConnectionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCloudConnectionsResponse(val *ListCloudConnectionsResponse) *NullableListCloudConnectionsResponse {
	return &NullableListCloudConnectionsResponse{value: val, isSet: true}
}

func (v NullableListCloudConnectionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCloudConnectionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
