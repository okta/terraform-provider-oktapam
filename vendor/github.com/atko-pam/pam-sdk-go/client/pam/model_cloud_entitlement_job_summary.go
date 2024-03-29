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
	"time"
)

// checks if the CloudEntitlementJobSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudEntitlementJobSummary{}

// CloudEntitlementJobSummary struct for CloudEntitlementJobSummary
type CloudEntitlementJobSummary struct {
	// The status of the job run
	Status string `json:"status"`
	// A timestamp indicating when the job was last run
	LastUpdatedAt time.Time                         `json:"last_updated_at"`
	Errors        CloudEntitlementJobSummaryErrors  `json:"errors"`
	Details       CloudEntitlementJobSummaryDetails `json:"details"`
}

// NewCloudEntitlementJobSummary instantiates a new CloudEntitlementJobSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudEntitlementJobSummary(status string, lastUpdatedAt time.Time, errors CloudEntitlementJobSummaryErrors, details CloudEntitlementJobSummaryDetails) *CloudEntitlementJobSummary {
	this := CloudEntitlementJobSummary{}
	this.Status = status
	this.LastUpdatedAt = lastUpdatedAt
	this.Errors = errors
	this.Details = details
	return &this
}

// NewCloudEntitlementJobSummaryWithDefaults instantiates a new CloudEntitlementJobSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudEntitlementJobSummaryWithDefaults() *CloudEntitlementJobSummary {
	this := CloudEntitlementJobSummary{}
	return &this
}

// GetStatus returns the Status field value
func (o *CloudEntitlementJobSummary) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CloudEntitlementJobSummary) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CloudEntitlementJobSummary) SetStatus(v string) *CloudEntitlementJobSummary {
	o.Status = v
	return o
}

// GetLastUpdatedAt returns the LastUpdatedAt field value
func (o *CloudEntitlementJobSummary) GetLastUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdatedAt
}

// GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field value
// and a boolean to check if the value has been set.
func (o *CloudEntitlementJobSummary) GetLastUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedAt, true
}

// SetLastUpdatedAt sets field value
func (o *CloudEntitlementJobSummary) SetLastUpdatedAt(v time.Time) *CloudEntitlementJobSummary {
	o.LastUpdatedAt = v
	return o
}

// GetErrors returns the Errors field value
func (o *CloudEntitlementJobSummary) GetErrors() CloudEntitlementJobSummaryErrors {
	if o == nil {
		var ret CloudEntitlementJobSummaryErrors
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *CloudEntitlementJobSummary) GetErrorsOk() (*CloudEntitlementJobSummaryErrors, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Errors, true
}

// SetErrors sets field value
func (o *CloudEntitlementJobSummary) SetErrors(v CloudEntitlementJobSummaryErrors) *CloudEntitlementJobSummary {
	o.Errors = v
	return o
}

// GetDetails returns the Details field value
func (o *CloudEntitlementJobSummary) GetDetails() CloudEntitlementJobSummaryDetails {
	if o == nil {
		var ret CloudEntitlementJobSummaryDetails
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *CloudEntitlementJobSummary) GetDetailsOk() (*CloudEntitlementJobSummaryDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *CloudEntitlementJobSummary) SetDetails(v CloudEntitlementJobSummaryDetails) *CloudEntitlementJobSummary {
	o.Details = v
	return o
}

func (o CloudEntitlementJobSummary) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudEntitlementJobSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["last_updated_at"] = o.LastUpdatedAt
	toSerialize["errors"] = o.Errors
	toSerialize["details"] = o.Details
	return toSerialize, nil
}

type NullableCloudEntitlementJobSummary struct {
	value *CloudEntitlementJobSummary
	isSet bool
}

func (v NullableCloudEntitlementJobSummary) Get() *CloudEntitlementJobSummary {
	return v.value
}

func (v *NullableCloudEntitlementJobSummary) Set(val *CloudEntitlementJobSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudEntitlementJobSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudEntitlementJobSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudEntitlementJobSummary(val *CloudEntitlementJobSummary) *NullableCloudEntitlementJobSummary {
	return &NullableCloudEntitlementJobSummary{value: val, isSet: true}
}

func (v NullableCloudEntitlementJobSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudEntitlementJobSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
