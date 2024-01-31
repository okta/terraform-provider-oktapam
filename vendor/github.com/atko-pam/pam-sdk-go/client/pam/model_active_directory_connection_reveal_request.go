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

// checks if the ActiveDirectoryConnectionRevealRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActiveDirectoryConnectionRevealRequest{}

// ActiveDirectoryConnectionRevealRequest struct for ActiveDirectoryConnectionRevealRequest
type ActiveDirectoryConnectionRevealRequest struct {
	PublicKey RawJSONWebKey `json:"public_key"`
}

// NewActiveDirectoryConnectionRevealRequest instantiates a new ActiveDirectoryConnectionRevealRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActiveDirectoryConnectionRevealRequest(publicKey RawJSONWebKey) *ActiveDirectoryConnectionRevealRequest {
	this := ActiveDirectoryConnectionRevealRequest{}
	this.PublicKey = publicKey
	return &this
}

// NewActiveDirectoryConnectionRevealRequestWithDefaults instantiates a new ActiveDirectoryConnectionRevealRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActiveDirectoryConnectionRevealRequestWithDefaults() *ActiveDirectoryConnectionRevealRequest {
	this := ActiveDirectoryConnectionRevealRequest{}
	return &this
}

// GetPublicKey returns the PublicKey field value
func (o *ActiveDirectoryConnectionRevealRequest) GetPublicKey() RawJSONWebKey {
	if o == nil {
		var ret RawJSONWebKey
		return ret
	}

	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value
// and a boolean to check if the value has been set.
func (o *ActiveDirectoryConnectionRevealRequest) GetPublicKeyOk() (*RawJSONWebKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicKey, true
}

// SetPublicKey sets field value
func (o *ActiveDirectoryConnectionRevealRequest) SetPublicKey(v RawJSONWebKey) *ActiveDirectoryConnectionRevealRequest {
	o.PublicKey = v
	return o
}

func (o ActiveDirectoryConnectionRevealRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActiveDirectoryConnectionRevealRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["public_key"] = o.PublicKey
	return toSerialize, nil
}

type NullableActiveDirectoryConnectionRevealRequest struct {
	value *ActiveDirectoryConnectionRevealRequest
	isSet bool
}

func (v NullableActiveDirectoryConnectionRevealRequest) Get() *ActiveDirectoryConnectionRevealRequest {
	return v.value
}

func (v *NullableActiveDirectoryConnectionRevealRequest) Set(val *ActiveDirectoryConnectionRevealRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableActiveDirectoryConnectionRevealRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableActiveDirectoryConnectionRevealRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActiveDirectoryConnectionRevealRequest(val *ActiveDirectoryConnectionRevealRequest) *NullableActiveDirectoryConnectionRevealRequest {
	return &NullableActiveDirectoryConnectionRevealRequest{value: val, isSet: true}
}

func (v NullableActiveDirectoryConnectionRevealRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActiveDirectoryConnectionRevealRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
