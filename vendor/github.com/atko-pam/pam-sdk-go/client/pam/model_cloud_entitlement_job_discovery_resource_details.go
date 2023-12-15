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

// checks if the CloudEntitlementJobDiscoveryResourceDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudEntitlementJobDiscoveryResourceDetails{}

// CloudEntitlementJobDiscoveryResourceDetails The specific details that define how the job discovers resources
type CloudEntitlementJobDiscoveryResourceDetails struct {
	// Defines the specific resource types the job attempts to discover
	Type string `json:"type"`
	// The specific entitlements the job attempts to discover
	Action []string `json:"action"`
}

// NewCloudEntitlementJobDiscoveryResourceDetails instantiates a new CloudEntitlementJobDiscoveryResourceDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudEntitlementJobDiscoveryResourceDetails(type_ string, action []string) *CloudEntitlementJobDiscoveryResourceDetails {
	this := CloudEntitlementJobDiscoveryResourceDetails{}
	this.Type = type_
	this.Action = action
	return &this
}

// NewCloudEntitlementJobDiscoveryResourceDetailsWithDefaults instantiates a new CloudEntitlementJobDiscoveryResourceDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudEntitlementJobDiscoveryResourceDetailsWithDefaults() *CloudEntitlementJobDiscoveryResourceDetails {
	this := CloudEntitlementJobDiscoveryResourceDetails{}
	return &this
}

// GetType returns the Type field value
func (o *CloudEntitlementJobDiscoveryResourceDetails) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CloudEntitlementJobDiscoveryResourceDetails) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CloudEntitlementJobDiscoveryResourceDetails) SetType(v string) *CloudEntitlementJobDiscoveryResourceDetails {
	o.Type = v
	return o
}

// GetAction returns the Action field value
func (o *CloudEntitlementJobDiscoveryResourceDetails) GetAction() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *CloudEntitlementJobDiscoveryResourceDetails) GetActionOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Action, true
}

// SetAction sets field value
func (o *CloudEntitlementJobDiscoveryResourceDetails) SetAction(v []string) *CloudEntitlementJobDiscoveryResourceDetails {
	o.Action = v
	return o
}

func (o CloudEntitlementJobDiscoveryResourceDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudEntitlementJobDiscoveryResourceDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["action"] = o.Action
	return toSerialize, nil
}

type NullableCloudEntitlementJobDiscoveryResourceDetails struct {
	value *CloudEntitlementJobDiscoveryResourceDetails
	isSet bool
}

func (v NullableCloudEntitlementJobDiscoveryResourceDetails) Get() *CloudEntitlementJobDiscoveryResourceDetails {
	return v.value
}

func (v *NullableCloudEntitlementJobDiscoveryResourceDetails) Set(val *CloudEntitlementJobDiscoveryResourceDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudEntitlementJobDiscoveryResourceDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudEntitlementJobDiscoveryResourceDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudEntitlementJobDiscoveryResourceDetails(val *CloudEntitlementJobDiscoveryResourceDetails) *NullableCloudEntitlementJobDiscoveryResourceDetails {
	return &NullableCloudEntitlementJobDiscoveryResourceDetails{value: val, isSet: true}
}

func (v NullableCloudEntitlementJobDiscoveryResourceDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudEntitlementJobDiscoveryResourceDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}