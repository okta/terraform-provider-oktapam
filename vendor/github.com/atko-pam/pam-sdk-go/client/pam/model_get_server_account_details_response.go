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

// checks if the GetServerAccountDetailsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetServerAccountDetailsResponse{}

// GetServerAccountDetailsResponse struct for GetServerAccountDetailsResponse
type GetServerAccountDetailsResponse struct {
	Items *ServerAccounts `json:"items,omitempty"`
}

// NewGetServerAccountDetailsResponse instantiates a new GetServerAccountDetailsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetServerAccountDetailsResponse() *GetServerAccountDetailsResponse {
	this := GetServerAccountDetailsResponse{}
	return &this
}

// NewGetServerAccountDetailsResponseWithDefaults instantiates a new GetServerAccountDetailsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetServerAccountDetailsResponseWithDefaults() *GetServerAccountDetailsResponse {
	this := GetServerAccountDetailsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *GetServerAccountDetailsResponse) GetItems() ServerAccounts {
	if o == nil || IsNil(o.Items) {
		var ret ServerAccounts
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServerAccountDetailsResponse) GetItemsOk() (*ServerAccounts, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *GetServerAccountDetailsResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given ServerAccounts and assigns it to the Items field.
func (o *GetServerAccountDetailsResponse) SetItems(v ServerAccounts) *GetServerAccountDetailsResponse {
	o.Items = &v
	return o
}

func (o GetServerAccountDetailsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetServerAccountDetailsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableGetServerAccountDetailsResponse struct {
	value *GetServerAccountDetailsResponse
	isSet bool
}

func (v NullableGetServerAccountDetailsResponse) Get() *GetServerAccountDetailsResponse {
	return v.value
}

func (v *NullableGetServerAccountDetailsResponse) Set(val *GetServerAccountDetailsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetServerAccountDetailsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetServerAccountDetailsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetServerAccountDetailsResponse(val *GetServerAccountDetailsResponse) *NullableGetServerAccountDetailsResponse {
	return &NullableGetServerAccountDetailsResponse{value: val, isSet: true}
}

func (v NullableGetServerAccountDetailsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetServerAccountDetailsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
