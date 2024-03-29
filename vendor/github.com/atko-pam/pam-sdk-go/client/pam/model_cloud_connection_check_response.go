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

// checks if the CloudConnectionCheckResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudConnectionCheckResponse{}

// CloudConnectionCheckResponse struct for CloudConnectionCheckResponse
type CloudConnectionCheckResponse struct {
	// The status of the Cloud Connection test
	Status string `json:"status"`
	// Details of the Cloud Connection test
	Details string `json:"details"`
}

// NewCloudConnectionCheckResponse instantiates a new CloudConnectionCheckResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudConnectionCheckResponse(status string, details string) *CloudConnectionCheckResponse {
	this := CloudConnectionCheckResponse{}
	this.Status = status
	this.Details = details
	return &this
}

// NewCloudConnectionCheckResponseWithDefaults instantiates a new CloudConnectionCheckResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudConnectionCheckResponseWithDefaults() *CloudConnectionCheckResponse {
	this := CloudConnectionCheckResponse{}
	return &this
}

// GetStatus returns the Status field value
func (o *CloudConnectionCheckResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CloudConnectionCheckResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CloudConnectionCheckResponse) SetStatus(v string) *CloudConnectionCheckResponse {
	o.Status = v
	return o
}

// GetDetails returns the Details field value
func (o *CloudConnectionCheckResponse) GetDetails() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *CloudConnectionCheckResponse) GetDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *CloudConnectionCheckResponse) SetDetails(v string) *CloudConnectionCheckResponse {
	o.Details = v
	return o
}

func (o CloudConnectionCheckResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudConnectionCheckResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["details"] = o.Details
	return toSerialize, nil
}

type NullableCloudConnectionCheckResponse struct {
	value *CloudConnectionCheckResponse
	isSet bool
}

func (v NullableCloudConnectionCheckResponse) Get() *CloudConnectionCheckResponse {
	return v.value
}

func (v *NullableCloudConnectionCheckResponse) Set(val *CloudConnectionCheckResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudConnectionCheckResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudConnectionCheckResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudConnectionCheckResponse(val *CloudConnectionCheckResponse) *NullableCloudConnectionCheckResponse {
	return &NullableCloudConnectionCheckResponse{value: val, isSet: true}
}

func (v NullableCloudConnectionCheckResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudConnectionCheckResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
