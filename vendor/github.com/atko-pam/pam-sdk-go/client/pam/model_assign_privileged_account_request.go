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

// checks if the AssignPrivilegedAccountRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssignPrivilegedAccountRequest{}

// AssignPrivilegedAccountRequest struct for AssignPrivilegedAccountRequest
type AssignPrivilegedAccountRequest struct {
	// The UUID of the Resource Group to assign the Privileged Account to
	ResourceGroupId *string `json:"resource_group_id,omitempty"`
	// The UUID of the Project to assign the Privileged Account to
	ProjectId *string `json:"project_id,omitempty"`
}

// NewAssignPrivilegedAccountRequest instantiates a new AssignPrivilegedAccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssignPrivilegedAccountRequest() *AssignPrivilegedAccountRequest {
	this := AssignPrivilegedAccountRequest{}
	return &this
}

// NewAssignPrivilegedAccountRequestWithDefaults instantiates a new AssignPrivilegedAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssignPrivilegedAccountRequestWithDefaults() *AssignPrivilegedAccountRequest {
	this := AssignPrivilegedAccountRequest{}
	return &this
}

// GetResourceGroupId returns the ResourceGroupId field value if set, zero value otherwise.
func (o *AssignPrivilegedAccountRequest) GetResourceGroupId() string {
	if o == nil || IsNil(o.ResourceGroupId) {
		var ret string
		return ret
	}
	return *o.ResourceGroupId
}

// GetResourceGroupIdOk returns a tuple with the ResourceGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignPrivilegedAccountRequest) GetResourceGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceGroupId) {
		return nil, false
	}
	return o.ResourceGroupId, true
}

// HasResourceGroupId returns a boolean if a field has been set.
func (o *AssignPrivilegedAccountRequest) HasResourceGroupId() bool {
	if o != nil && !IsNil(o.ResourceGroupId) {
		return true
	}

	return false
}

// SetResourceGroupId gets a reference to the given string and assigns it to the ResourceGroupId field.
func (o *AssignPrivilegedAccountRequest) SetResourceGroupId(v string) *AssignPrivilegedAccountRequest {
	o.ResourceGroupId = &v
	return o
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *AssignPrivilegedAccountRequest) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignPrivilegedAccountRequest) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *AssignPrivilegedAccountRequest) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *AssignPrivilegedAccountRequest) SetProjectId(v string) *AssignPrivilegedAccountRequest {
	o.ProjectId = &v
	return o
}

func (o AssignPrivilegedAccountRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssignPrivilegedAccountRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResourceGroupId) {
		toSerialize["resource_group_id"] = o.ResourceGroupId
	}
	if !IsNil(o.ProjectId) {
		toSerialize["project_id"] = o.ProjectId
	}
	return toSerialize, nil
}

type NullableAssignPrivilegedAccountRequest struct {
	value *AssignPrivilegedAccountRequest
	isSet bool
}

func (v NullableAssignPrivilegedAccountRequest) Get() *AssignPrivilegedAccountRequest {
	return v.value
}

func (v *NullableAssignPrivilegedAccountRequest) Set(val *AssignPrivilegedAccountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignPrivilegedAccountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignPrivilegedAccountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignPrivilegedAccountRequest(val *AssignPrivilegedAccountRequest) *NullableAssignPrivilegedAccountRequest {
	return &NullableAssignPrivilegedAccountRequest{value: val, isSet: true}
}

func (v NullableAssignPrivilegedAccountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssignPrivilegedAccountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
