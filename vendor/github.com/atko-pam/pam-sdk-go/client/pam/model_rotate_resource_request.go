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

// checks if the RotateResourceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RotateResourceRequest{}

// RotateResourceRequest struct for RotateResourceRequest
type RotateResourceRequest struct {
	// The UUID or ORN of the resource
	ResourceId   string               `json:"resource_id"`
	ResourceType CheckoutResourceType `json:"resource_type"`
}

// NewRotateResourceRequest instantiates a new RotateResourceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRotateResourceRequest(resourceId string, resourceType CheckoutResourceType) *RotateResourceRequest {
	this := RotateResourceRequest{}
	this.ResourceId = resourceId
	this.ResourceType = resourceType
	return &this
}

// NewRotateResourceRequestWithDefaults instantiates a new RotateResourceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRotateResourceRequestWithDefaults() *RotateResourceRequest {
	this := RotateResourceRequest{}
	return &this
}

// GetResourceId returns the ResourceId field value
func (o *RotateResourceRequest) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *RotateResourceRequest) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *RotateResourceRequest) SetResourceId(v string) *RotateResourceRequest {
	o.ResourceId = v
	return o
}

// GetResourceType returns the ResourceType field value
func (o *RotateResourceRequest) GetResourceType() CheckoutResourceType {
	if o == nil {
		var ret CheckoutResourceType
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *RotateResourceRequest) GetResourceTypeOk() (*CheckoutResourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *RotateResourceRequest) SetResourceType(v CheckoutResourceType) *RotateResourceRequest {
	o.ResourceType = v
	return o
}

func (o RotateResourceRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RotateResourceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resource_id"] = o.ResourceId
	toSerialize["resource_type"] = o.ResourceType
	return toSerialize, nil
}

type NullableRotateResourceRequest struct {
	value *RotateResourceRequest
	isSet bool
}

func (v NullableRotateResourceRequest) Get() *RotateResourceRequest {
	return v.value
}

func (v *NullableRotateResourceRequest) Set(val *RotateResourceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRotateResourceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRotateResourceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRotateResourceRequest(val *RotateResourceRequest) *NullableRotateResourceRequest {
	return &NullableRotateResourceRequest{value: val, isSet: true}
}

func (v NullableRotateResourceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRotateResourceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
