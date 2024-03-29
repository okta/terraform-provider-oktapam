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

// checks if the GetVaultJWKSResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetVaultJWKSResponse{}

// GetVaultJWKSResponse struct for GetVaultJWKSResponse
type GetVaultJWKSResponse struct {
	Keys []RawJSONWebKey `json:"keys,omitempty"`
}

// NewGetVaultJWKSResponse instantiates a new GetVaultJWKSResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetVaultJWKSResponse() *GetVaultJWKSResponse {
	this := GetVaultJWKSResponse{}
	return &this
}

// NewGetVaultJWKSResponseWithDefaults instantiates a new GetVaultJWKSResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetVaultJWKSResponseWithDefaults() *GetVaultJWKSResponse {
	this := GetVaultJWKSResponse{}
	return &this
}

// GetKeys returns the Keys field value if set, zero value otherwise.
func (o *GetVaultJWKSResponse) GetKeys() []RawJSONWebKey {
	if o == nil || IsNil(o.Keys) {
		var ret []RawJSONWebKey
		return ret
	}
	return o.Keys
}

// GetKeysOk returns a tuple with the Keys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVaultJWKSResponse) GetKeysOk() ([]RawJSONWebKey, bool) {
	if o == nil || IsNil(o.Keys) {
		return nil, false
	}
	return o.Keys, true
}

// HasKeys returns a boolean if a field has been set.
func (o *GetVaultJWKSResponse) HasKeys() bool {
	if o != nil && !IsNil(o.Keys) {
		return true
	}

	return false
}

// SetKeys gets a reference to the given []RawJSONWebKey and assigns it to the Keys field.
func (o *GetVaultJWKSResponse) SetKeys(v []RawJSONWebKey) *GetVaultJWKSResponse {
	o.Keys = v
	return o
}

func (o GetVaultJWKSResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetVaultJWKSResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Keys) {
		toSerialize["keys"] = o.Keys
	}
	return toSerialize, nil
}

type NullableGetVaultJWKSResponse struct {
	value *GetVaultJWKSResponse
	isSet bool
}

func (v NullableGetVaultJWKSResponse) Get() *GetVaultJWKSResponse {
	return v.value
}

func (v *NullableGetVaultJWKSResponse) Set(val *GetVaultJWKSResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetVaultJWKSResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetVaultJWKSResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetVaultJWKSResponse(val *GetVaultJWKSResponse) *NullableGetVaultJWKSResponse {
	return &NullableGetVaultJWKSResponse{value: val, isSet: true}
}

func (v NullableGetVaultJWKSResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetVaultJWKSResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
