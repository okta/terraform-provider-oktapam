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

// checks if the UpdateActiveDirectoryConnectionStatusRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateActiveDirectoryConnectionStatusRequest{}

// UpdateActiveDirectoryConnectionStatusRequest struct for UpdateActiveDirectoryConnectionStatusRequest
type UpdateActiveDirectoryConnectionStatusRequest struct {
	Status *UpdateableADConnectionStatus `json:"status,omitempty"`
}

// NewUpdateActiveDirectoryConnectionStatusRequest instantiates a new UpdateActiveDirectoryConnectionStatusRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateActiveDirectoryConnectionStatusRequest() *UpdateActiveDirectoryConnectionStatusRequest {
	this := UpdateActiveDirectoryConnectionStatusRequest{}
	return &this
}

// NewUpdateActiveDirectoryConnectionStatusRequestWithDefaults instantiates a new UpdateActiveDirectoryConnectionStatusRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateActiveDirectoryConnectionStatusRequestWithDefaults() *UpdateActiveDirectoryConnectionStatusRequest {
	this := UpdateActiveDirectoryConnectionStatusRequest{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UpdateActiveDirectoryConnectionStatusRequest) GetStatus() UpdateableADConnectionStatus {
	if o == nil || IsNil(o.Status) {
		var ret UpdateableADConnectionStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateActiveDirectoryConnectionStatusRequest) GetStatusOk() (*UpdateableADConnectionStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UpdateActiveDirectoryConnectionStatusRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given UpdateableADConnectionStatus and assigns it to the Status field.
func (o *UpdateActiveDirectoryConnectionStatusRequest) SetStatus(v UpdateableADConnectionStatus) *UpdateActiveDirectoryConnectionStatusRequest {
	o.Status = &v
	return o
}

func (o UpdateActiveDirectoryConnectionStatusRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateActiveDirectoryConnectionStatusRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableUpdateActiveDirectoryConnectionStatusRequest struct {
	value *UpdateActiveDirectoryConnectionStatusRequest
	isSet bool
}

func (v NullableUpdateActiveDirectoryConnectionStatusRequest) Get() *UpdateActiveDirectoryConnectionStatusRequest {
	return v.value
}

func (v *NullableUpdateActiveDirectoryConnectionStatusRequest) Set(val *UpdateActiveDirectoryConnectionStatusRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateActiveDirectoryConnectionStatusRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateActiveDirectoryConnectionStatusRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateActiveDirectoryConnectionStatusRequest(val *UpdateActiveDirectoryConnectionStatusRequest) *NullableUpdateActiveDirectoryConnectionStatusRequest {
	return &NullableUpdateActiveDirectoryConnectionStatusRequest{value: val, isSet: true}
}

func (v NullableUpdateActiveDirectoryConnectionStatusRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateActiveDirectoryConnectionStatusRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
