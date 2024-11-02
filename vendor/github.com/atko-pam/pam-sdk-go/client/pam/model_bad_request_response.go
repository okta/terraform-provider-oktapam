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

// checks if the BadRequestResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BadRequestResponse{}

// BadRequestResponse JSON response for a bad request
type BadRequestResponse struct {
	Nottype *string `json:"nottype,omitempty"`
	Code    *int32  `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewBadRequestResponse instantiates a new BadRequestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBadRequestResponse() *BadRequestResponse {
	this := BadRequestResponse{}
	return &this
}

// NewBadRequestResponseWithDefaults instantiates a new BadRequestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBadRequestResponseWithDefaults() *BadRequestResponse {
	this := BadRequestResponse{}
	return &this
}

// GetNottype returns the Nottype field value if set, zero value otherwise.
func (o *BadRequestResponse) GetNottype() string {
	if o == nil || IsNil(o.Nottype) {
		var ret string
		return ret
	}
	return *o.Nottype
}

// GetNottypeOk returns a tuple with the Nottype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BadRequestResponse) GetNottypeOk() (*string, bool) {
	if o == nil || IsNil(o.Nottype) {
		return nil, false
	}
	return o.Nottype, true
}

// HasNottype returns a boolean if a field has been set.
func (o *BadRequestResponse) HasNottype() bool {
	if o != nil && !IsNil(o.Nottype) {
		return true
	}

	return false
}

// SetNottype gets a reference to the given string and assigns it to the Nottype field.
func (o *BadRequestResponse) SetNottype(v string) *BadRequestResponse {
	o.Nottype = &v
	return o
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *BadRequestResponse) GetCode() int32 {
	if o == nil || IsNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BadRequestResponse) GetCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *BadRequestResponse) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *BadRequestResponse) SetCode(v int32) *BadRequestResponse {
	o.Code = &v
	return o
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *BadRequestResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BadRequestResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *BadRequestResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *BadRequestResponse) SetMessage(v string) *BadRequestResponse {
	o.Message = &v
	return o
}

func (o BadRequestResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BadRequestResponse) ToMap() (map[string]interface{}, error) {
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

type NullableBadRequestResponse struct {
	value *BadRequestResponse
	isSet bool
}

func (v NullableBadRequestResponse) Get() *BadRequestResponse {
	return v.value
}

func (v *NullableBadRequestResponse) Set(val *BadRequestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBadRequestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBadRequestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBadRequestResponse(val *BadRequestResponse) *NullableBadRequestResponse {
	return &NullableBadRequestResponse{value: val, isSet: true}
}

func (v NullableBadRequestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBadRequestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
