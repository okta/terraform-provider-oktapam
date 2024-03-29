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

// checks if the CloudEntitlementJobSummaryErrors type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudEntitlementJobSummaryErrors{}

// CloudEntitlementJobSummaryErrors Includes a list of errors associated with the Cloud Entitlement Job run. This is only returned if the job run has a status of 'ERROR'.
type CloudEntitlementJobSummaryErrors struct {
	// The specific error messages returned from the job run
	Message *string `json:"message,omitempty"`
}

// NewCloudEntitlementJobSummaryErrors instantiates a new CloudEntitlementJobSummaryErrors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudEntitlementJobSummaryErrors() *CloudEntitlementJobSummaryErrors {
	this := CloudEntitlementJobSummaryErrors{}
	return &this
}

// NewCloudEntitlementJobSummaryErrorsWithDefaults instantiates a new CloudEntitlementJobSummaryErrors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudEntitlementJobSummaryErrorsWithDefaults() *CloudEntitlementJobSummaryErrors {
	this := CloudEntitlementJobSummaryErrors{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CloudEntitlementJobSummaryErrors) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudEntitlementJobSummaryErrors) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CloudEntitlementJobSummaryErrors) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *CloudEntitlementJobSummaryErrors) SetMessage(v string) *CloudEntitlementJobSummaryErrors {
	o.Message = &v
	return o
}

func (o CloudEntitlementJobSummaryErrors) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudEntitlementJobSummaryErrors) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableCloudEntitlementJobSummaryErrors struct {
	value *CloudEntitlementJobSummaryErrors
	isSet bool
}

func (v NullableCloudEntitlementJobSummaryErrors) Get() *CloudEntitlementJobSummaryErrors {
	return v.value
}

func (v *NullableCloudEntitlementJobSummaryErrors) Set(val *CloudEntitlementJobSummaryErrors) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudEntitlementJobSummaryErrors) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudEntitlementJobSummaryErrors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudEntitlementJobSummaryErrors(val *CloudEntitlementJobSummaryErrors) *NullableCloudEntitlementJobSummaryErrors {
	return &NullableCloudEntitlementJobSummaryErrors{value: val, isSet: true}
}

func (v NullableCloudEntitlementJobSummaryErrors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudEntitlementJobSummaryErrors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
