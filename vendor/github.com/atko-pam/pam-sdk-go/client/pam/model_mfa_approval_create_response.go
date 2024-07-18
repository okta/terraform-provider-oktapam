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

// checks if the MFAApprovalCreateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MFAApprovalCreateResponse{}

// MFAApprovalCreateResponse struct for MFAApprovalCreateResponse
type MFAApprovalCreateResponse struct {
	MfaId       *string `json:"mfa_id,omitempty"`
	RedirectUrl *string `json:"redirect_url,omitempty"`
	PollingUrl  *string `json:"polling_url,omitempty"`
}

// NewMFAApprovalCreateResponse instantiates a new MFAApprovalCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMFAApprovalCreateResponse() *MFAApprovalCreateResponse {
	this := MFAApprovalCreateResponse{}
	return &this
}

// NewMFAApprovalCreateResponseWithDefaults instantiates a new MFAApprovalCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMFAApprovalCreateResponseWithDefaults() *MFAApprovalCreateResponse {
	this := MFAApprovalCreateResponse{}
	return &this
}

// GetMfaId returns the MfaId field value if set, zero value otherwise.
func (o *MFAApprovalCreateResponse) GetMfaId() string {
	if o == nil || IsNil(o.MfaId) {
		var ret string
		return ret
	}
	return *o.MfaId
}

// GetMfaIdOk returns a tuple with the MfaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MFAApprovalCreateResponse) GetMfaIdOk() (*string, bool) {
	if o == nil || IsNil(o.MfaId) {
		return nil, false
	}
	return o.MfaId, true
}

// HasMfaId returns a boolean if a field has been set.
func (o *MFAApprovalCreateResponse) HasMfaId() bool {
	if o != nil && !IsNil(o.MfaId) {
		return true
	}

	return false
}

// SetMfaId gets a reference to the given string and assigns it to the MfaId field.
func (o *MFAApprovalCreateResponse) SetMfaId(v string) *MFAApprovalCreateResponse {
	o.MfaId = &v
	return o
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise.
func (o *MFAApprovalCreateResponse) GetRedirectUrl() string {
	if o == nil || IsNil(o.RedirectUrl) {
		var ret string
		return ret
	}
	return *o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MFAApprovalCreateResponse) GetRedirectUrlOk() (*string, bool) {
	if o == nil || IsNil(o.RedirectUrl) {
		return nil, false
	}
	return o.RedirectUrl, true
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *MFAApprovalCreateResponse) HasRedirectUrl() bool {
	if o != nil && !IsNil(o.RedirectUrl) {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given string and assigns it to the RedirectUrl field.
func (o *MFAApprovalCreateResponse) SetRedirectUrl(v string) *MFAApprovalCreateResponse {
	o.RedirectUrl = &v
	return o
}

// GetPollingUrl returns the PollingUrl field value if set, zero value otherwise.
func (o *MFAApprovalCreateResponse) GetPollingUrl() string {
	if o == nil || IsNil(o.PollingUrl) {
		var ret string
		return ret
	}
	return *o.PollingUrl
}

// GetPollingUrlOk returns a tuple with the PollingUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MFAApprovalCreateResponse) GetPollingUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PollingUrl) {
		return nil, false
	}
	return o.PollingUrl, true
}

// HasPollingUrl returns a boolean if a field has been set.
func (o *MFAApprovalCreateResponse) HasPollingUrl() bool {
	if o != nil && !IsNil(o.PollingUrl) {
		return true
	}

	return false
}

// SetPollingUrl gets a reference to the given string and assigns it to the PollingUrl field.
func (o *MFAApprovalCreateResponse) SetPollingUrl(v string) *MFAApprovalCreateResponse {
	o.PollingUrl = &v
	return o
}

func (o MFAApprovalCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MFAApprovalCreateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MfaId) {
		toSerialize["mfa_id"] = o.MfaId
	}
	if !IsNil(o.RedirectUrl) {
		toSerialize["redirect_url"] = o.RedirectUrl
	}
	if !IsNil(o.PollingUrl) {
		toSerialize["polling_url"] = o.PollingUrl
	}
	return toSerialize, nil
}

type NullableMFAApprovalCreateResponse struct {
	value *MFAApprovalCreateResponse
	isSet bool
}

func (v NullableMFAApprovalCreateResponse) Get() *MFAApprovalCreateResponse {
	return v.value
}

func (v *NullableMFAApprovalCreateResponse) Set(val *MFAApprovalCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMFAApprovalCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMFAApprovalCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMFAApprovalCreateResponse(val *MFAApprovalCreateResponse) *NullableMFAApprovalCreateResponse {
	return &NullableMFAApprovalCreateResponse{value: val, isSet: true}
}

func (v NullableMFAApprovalCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMFAApprovalCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
