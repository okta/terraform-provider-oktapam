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

// checks if the UpdateGroupAttribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateGroupAttribute{}

// UpdateGroupAttribute struct for UpdateGroupAttribute
type UpdateGroupAttribute struct {
	AttributeName  GroupAttributeName               `json:"attribute_name"`
	AttributeValue TeamGroupAttributeAttributeValue `json:"attribute_value"`
}

// NewUpdateGroupAttribute instantiates a new UpdateGroupAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateGroupAttribute(attributeName GroupAttributeName, attributeValue TeamGroupAttributeAttributeValue) *UpdateGroupAttribute {
	this := UpdateGroupAttribute{}
	this.AttributeName = attributeName
	this.AttributeValue = attributeValue
	return &this
}

// NewUpdateGroupAttributeWithDefaults instantiates a new UpdateGroupAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateGroupAttributeWithDefaults() *UpdateGroupAttribute {
	this := UpdateGroupAttribute{}
	return &this
}

// GetAttributeName returns the AttributeName field value
func (o *UpdateGroupAttribute) GetAttributeName() GroupAttributeName {
	if o == nil {
		var ret GroupAttributeName
		return ret
	}

	return o.AttributeName
}

// GetAttributeNameOk returns a tuple with the AttributeName field value
// and a boolean to check if the value has been set.
func (o *UpdateGroupAttribute) GetAttributeNameOk() (*GroupAttributeName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeName, true
}

// SetAttributeName sets field value
func (o *UpdateGroupAttribute) SetAttributeName(v GroupAttributeName) *UpdateGroupAttribute {
	o.AttributeName = v
	return o
}

// GetAttributeValue returns the AttributeValue field value
func (o *UpdateGroupAttribute) GetAttributeValue() TeamGroupAttributeAttributeValue {
	if o == nil {
		var ret TeamGroupAttributeAttributeValue
		return ret
	}

	return o.AttributeValue
}

// GetAttributeValueOk returns a tuple with the AttributeValue field value
// and a boolean to check if the value has been set.
func (o *UpdateGroupAttribute) GetAttributeValueOk() (*TeamGroupAttributeAttributeValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeValue, true
}

// SetAttributeValue sets field value
func (o *UpdateGroupAttribute) SetAttributeValue(v TeamGroupAttributeAttributeValue) *UpdateGroupAttribute {
	o.AttributeValue = v
	return o
}

func (o UpdateGroupAttribute) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateGroupAttribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attribute_name"] = o.AttributeName
	toSerialize["attribute_value"] = o.AttributeValue
	return toSerialize, nil
}

type NullableUpdateGroupAttribute struct {
	value *UpdateGroupAttribute
	isSet bool
}

func (v NullableUpdateGroupAttribute) Get() *UpdateGroupAttribute {
	return v.value
}

func (v *NullableUpdateGroupAttribute) Set(val *UpdateGroupAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateGroupAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateGroupAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateGroupAttribute(val *UpdateGroupAttribute) *NullableUpdateGroupAttribute {
	return &NullableUpdateGroupAttribute{value: val, isSet: true}
}

func (v NullableUpdateGroupAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateGroupAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
