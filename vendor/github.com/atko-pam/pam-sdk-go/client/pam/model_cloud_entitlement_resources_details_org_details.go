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

// checks if the CloudEntitlementResourcesDetailsOrgDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudEntitlementResourcesDetailsOrgDetails{}

// CloudEntitlementResourcesDetailsOrgDetails The details of the associated AWS organization
type CloudEntitlementResourcesDetailsOrgDetails struct {
	// The UUID of the AWS organization
	Orgid string `json:"orgid"`
	// The name of the AWS organization
	Name string `json:"name"`
}

// NewCloudEntitlementResourcesDetailsOrgDetails instantiates a new CloudEntitlementResourcesDetailsOrgDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudEntitlementResourcesDetailsOrgDetails(orgid string, name string) *CloudEntitlementResourcesDetailsOrgDetails {
	this := CloudEntitlementResourcesDetailsOrgDetails{}
	this.Orgid = orgid
	this.Name = name
	return &this
}

// NewCloudEntitlementResourcesDetailsOrgDetailsWithDefaults instantiates a new CloudEntitlementResourcesDetailsOrgDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudEntitlementResourcesDetailsOrgDetailsWithDefaults() *CloudEntitlementResourcesDetailsOrgDetails {
	this := CloudEntitlementResourcesDetailsOrgDetails{}
	return &this
}

// GetOrgid returns the Orgid field value
func (o *CloudEntitlementResourcesDetailsOrgDetails) GetOrgid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Orgid
}

// GetOrgidOk returns a tuple with the Orgid field value
// and a boolean to check if the value has been set.
func (o *CloudEntitlementResourcesDetailsOrgDetails) GetOrgidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Orgid, true
}

// SetOrgid sets field value
func (o *CloudEntitlementResourcesDetailsOrgDetails) SetOrgid(v string) *CloudEntitlementResourcesDetailsOrgDetails {
	o.Orgid = v
	return o
}

// GetName returns the Name field value
func (o *CloudEntitlementResourcesDetailsOrgDetails) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CloudEntitlementResourcesDetailsOrgDetails) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CloudEntitlementResourcesDetailsOrgDetails) SetName(v string) *CloudEntitlementResourcesDetailsOrgDetails {
	o.Name = v
	return o
}

func (o CloudEntitlementResourcesDetailsOrgDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudEntitlementResourcesDetailsOrgDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["orgid"] = o.Orgid
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableCloudEntitlementResourcesDetailsOrgDetails struct {
	value *CloudEntitlementResourcesDetailsOrgDetails
	isSet bool
}

func (v NullableCloudEntitlementResourcesDetailsOrgDetails) Get() *CloudEntitlementResourcesDetailsOrgDetails {
	return v.value
}

func (v *NullableCloudEntitlementResourcesDetailsOrgDetails) Set(val *CloudEntitlementResourcesDetailsOrgDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudEntitlementResourcesDetailsOrgDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudEntitlementResourcesDetailsOrgDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudEntitlementResourcesDetailsOrgDetails(val *CloudEntitlementResourcesDetailsOrgDetails) *NullableCloudEntitlementResourcesDetailsOrgDetails {
	return &NullableCloudEntitlementResourcesDetailsOrgDetails{value: val, isSet: true}
}

func (v NullableCloudEntitlementResourcesDetailsOrgDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudEntitlementResourcesDetailsOrgDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
