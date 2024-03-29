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

// checks if the CloudEntitlementResources type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudEntitlementResources{}

// CloudEntitlementResources struct for CloudEntitlementResources
type CloudEntitlementResources struct {
	// The UUID of the resource
	Id string `json:"id"`
	// If true, indicates the resource is at risk
	AtRisk  *bool                             `json:"at_risk,omitempty"`
	Details *CloudEntitlementResourcesDetails `json:"details,omitempty"`
}

// NewCloudEntitlementResources instantiates a new CloudEntitlementResources object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudEntitlementResources(id string) *CloudEntitlementResources {
	this := CloudEntitlementResources{}
	this.Id = id
	return &this
}

// NewCloudEntitlementResourcesWithDefaults instantiates a new CloudEntitlementResources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudEntitlementResourcesWithDefaults() *CloudEntitlementResources {
	this := CloudEntitlementResources{}
	return &this
}

// GetId returns the Id field value
func (o *CloudEntitlementResources) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CloudEntitlementResources) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CloudEntitlementResources) SetId(v string) *CloudEntitlementResources {
	o.Id = v
	return o
}

// GetAtRisk returns the AtRisk field value if set, zero value otherwise.
func (o *CloudEntitlementResources) GetAtRisk() bool {
	if o == nil || IsNil(o.AtRisk) {
		var ret bool
		return ret
	}
	return *o.AtRisk
}

// GetAtRiskOk returns a tuple with the AtRisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudEntitlementResources) GetAtRiskOk() (*bool, bool) {
	if o == nil || IsNil(o.AtRisk) {
		return nil, false
	}
	return o.AtRisk, true
}

// HasAtRisk returns a boolean if a field has been set.
func (o *CloudEntitlementResources) HasAtRisk() bool {
	if o != nil && !IsNil(o.AtRisk) {
		return true
	}

	return false
}

// SetAtRisk gets a reference to the given bool and assigns it to the AtRisk field.
func (o *CloudEntitlementResources) SetAtRisk(v bool) *CloudEntitlementResources {
	o.AtRisk = &v
	return o
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *CloudEntitlementResources) GetDetails() CloudEntitlementResourcesDetails {
	if o == nil || IsNil(o.Details) {
		var ret CloudEntitlementResourcesDetails
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudEntitlementResources) GetDetailsOk() (*CloudEntitlementResourcesDetails, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *CloudEntitlementResources) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given CloudEntitlementResourcesDetails and assigns it to the Details field.
func (o *CloudEntitlementResources) SetDetails(v CloudEntitlementResourcesDetails) *CloudEntitlementResources {
	o.Details = &v
	return o
}

func (o CloudEntitlementResources) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudEntitlementResources) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.AtRisk) {
		toSerialize["at_risk"] = o.AtRisk
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableCloudEntitlementResources struct {
	value *CloudEntitlementResources
	isSet bool
}

func (v NullableCloudEntitlementResources) Get() *CloudEntitlementResources {
	return v.value
}

func (v *NullableCloudEntitlementResources) Set(val *CloudEntitlementResources) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudEntitlementResources) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudEntitlementResources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudEntitlementResources(val *CloudEntitlementResources) *NullableCloudEntitlementResources {
	return &NullableCloudEntitlementResources{value: val, isSet: true}
}

func (v NullableCloudEntitlementResources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudEntitlementResources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
