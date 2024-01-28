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

// checks if the RevealSecretForbiddenResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RevealSecretForbiddenResponse{}

// RevealSecretForbiddenResponse struct for RevealSecretForbiddenResponse
type RevealSecretForbiddenResponse struct {
	List []UserAccessMethod `json:"list,omitempty"`
}

// NewRevealSecretForbiddenResponse instantiates a new RevealSecretForbiddenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRevealSecretForbiddenResponse() *RevealSecretForbiddenResponse {
	this := RevealSecretForbiddenResponse{}
	return &this
}

// NewRevealSecretForbiddenResponseWithDefaults instantiates a new RevealSecretForbiddenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevealSecretForbiddenResponseWithDefaults() *RevealSecretForbiddenResponse {
	this := RevealSecretForbiddenResponse{}
	return &this
}

// GetList returns the List field value if set, zero value otherwise.
func (o *RevealSecretForbiddenResponse) GetList() []UserAccessMethod {
	if o == nil || IsNil(o.List) {
		var ret []UserAccessMethod
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevealSecretForbiddenResponse) GetListOk() ([]UserAccessMethod, bool) {
	if o == nil || IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *RevealSecretForbiddenResponse) HasList() bool {
	if o != nil && !IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given []UserAccessMethod and assigns it to the List field.
func (o *RevealSecretForbiddenResponse) SetList(v []UserAccessMethod) *RevealSecretForbiddenResponse {
	o.List = v
	return o
}

func (o RevealSecretForbiddenResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RevealSecretForbiddenResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.List) {
		toSerialize["list"] = o.List
	}
	return toSerialize, nil
}

type NullableRevealSecretForbiddenResponse struct {
	value *RevealSecretForbiddenResponse
	isSet bool
}

func (v NullableRevealSecretForbiddenResponse) Get() *RevealSecretForbiddenResponse {
	return v.value
}

func (v *NullableRevealSecretForbiddenResponse) Set(val *RevealSecretForbiddenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRevealSecretForbiddenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRevealSecretForbiddenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevealSecretForbiddenResponse(val *RevealSecretForbiddenResponse) *NullableRevealSecretForbiddenResponse {
	return &NullableRevealSecretForbiddenResponse{value: val, isSet: true}
}

func (v NullableRevealSecretForbiddenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevealSecretForbiddenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}