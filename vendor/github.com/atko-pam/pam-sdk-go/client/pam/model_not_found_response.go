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

// checks if the NotFoundResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotFoundResponse{}

// NotFoundResponse JSON response for accessing an invalid endpoint
type NotFoundResponse struct {
	Nottype *string `json:"nottype,omitempty"`
	Code    *int32  `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewNotFoundResponse instantiates a new NotFoundResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotFoundResponse() *NotFoundResponse {
	this := NotFoundResponse{}
	return &this
}

// NewNotFoundResponseWithDefaults instantiates a new NotFoundResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotFoundResponseWithDefaults() *NotFoundResponse {
	this := NotFoundResponse{}
	return &this
}

// GetNottype returns the Nottype field value if set, zero value otherwise.
func (o *NotFoundResponse) GetNottype() string {
	if o == nil || IsNil(o.Nottype) {
		var ret string
		return ret
	}
	return *o.Nottype
}

// GetNottypeOk returns a tuple with the Nottype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotFoundResponse) GetNottypeOk() (*string, bool) {
	if o == nil || IsNil(o.Nottype) {
		return nil, false
	}
	return o.Nottype, true
}

// HasNottype returns a boolean if a field has been set.
func (o *NotFoundResponse) HasNottype() bool {
	if o != nil && !IsNil(o.Nottype) {
		return true
	}

	return false
}

// SetNottype gets a reference to the given string and assigns it to the Nottype field.
func (o *NotFoundResponse) SetNottype(v string) *NotFoundResponse {
	o.Nottype = &v
	return o
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *NotFoundResponse) GetCode() int32 {
	if o == nil || IsNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotFoundResponse) GetCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *NotFoundResponse) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *NotFoundResponse) SetCode(v int32) *NotFoundResponse {
	o.Code = &v
	return o
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *NotFoundResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotFoundResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *NotFoundResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *NotFoundResponse) SetMessage(v string) *NotFoundResponse {
	o.Message = &v
	return o
}

func (o NotFoundResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotFoundResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Nottype) {
		toSerialize["nottype"] = o.Nottype
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableNotFoundResponse struct {
	value *NotFoundResponse
	isSet bool
}

func (v NullableNotFoundResponse) Get() *NotFoundResponse {
	return v.value
}

func (v *NullableNotFoundResponse) Set(val *NotFoundResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNotFoundResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNotFoundResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotFoundResponse(val *NotFoundResponse) *NullableNotFoundResponse {
	return &NullableNotFoundResponse{value: val, isSet: true}
}

func (v NullableNotFoundResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotFoundResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
