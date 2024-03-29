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

// checks if the ResourceGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceGroup{}

// ResourceGroup struct for ResourceGroup
type ResourceGroup struct {
	// The UUID of the Resource Group
	Id *string `json:"id,omitempty"`
	// The name of the Resource Group
	Name *string `json:"name,omitempty"`
	// The UUID of the Team associated with the Resource Group
	TeamId *string `json:"team_id,omitempty"`
	// A description of the Resource Group
	Description *string `json:"description,omitempty"`
	// An array of OPA groups that own this Resource Group. This param is only available to users with the `resource_admin` role.
	DelegatedResourceAdminGroups []NamedObject `json:"delegated_resource_admin_groups,omitempty"`
}

// NewResourceGroup instantiates a new ResourceGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceGroup() *ResourceGroup {
	this := ResourceGroup{}
	return &this
}

// NewResourceGroupWithDefaults instantiates a new ResourceGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceGroupWithDefaults() *ResourceGroup {
	this := ResourceGroup{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ResourceGroup) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceGroup) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ResourceGroup) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ResourceGroup) SetId(v string) *ResourceGroup {
	o.Id = &v
	return o
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ResourceGroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceGroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ResourceGroup) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ResourceGroup) SetName(v string) *ResourceGroup {
	o.Name = &v
	return o
}

// GetTeamId returns the TeamId field value if set, zero value otherwise.
func (o *ResourceGroup) GetTeamId() string {
	if o == nil || IsNil(o.TeamId) {
		var ret string
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceGroup) GetTeamIdOk() (*string, bool) {
	if o == nil || IsNil(o.TeamId) {
		return nil, false
	}
	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *ResourceGroup) HasTeamId() bool {
	if o != nil && !IsNil(o.TeamId) {
		return true
	}

	return false
}

// SetTeamId gets a reference to the given string and assigns it to the TeamId field.
func (o *ResourceGroup) SetTeamId(v string) *ResourceGroup {
	o.TeamId = &v
	return o
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ResourceGroup) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceGroup) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ResourceGroup) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ResourceGroup) SetDescription(v string) *ResourceGroup {
	o.Description = &v
	return o
}

// GetDelegatedResourceAdminGroups returns the DelegatedResourceAdminGroups field value if set, zero value otherwise.
func (o *ResourceGroup) GetDelegatedResourceAdminGroups() []NamedObject {
	if o == nil || IsNil(o.DelegatedResourceAdminGroups) {
		var ret []NamedObject
		return ret
	}
	return o.DelegatedResourceAdminGroups
}

// GetDelegatedResourceAdminGroupsOk returns a tuple with the DelegatedResourceAdminGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceGroup) GetDelegatedResourceAdminGroupsOk() ([]NamedObject, bool) {
	if o == nil || IsNil(o.DelegatedResourceAdminGroups) {
		return nil, false
	}
	return o.DelegatedResourceAdminGroups, true
}

// HasDelegatedResourceAdminGroups returns a boolean if a field has been set.
func (o *ResourceGroup) HasDelegatedResourceAdminGroups() bool {
	if o != nil && !IsNil(o.DelegatedResourceAdminGroups) {
		return true
	}

	return false
}

// SetDelegatedResourceAdminGroups gets a reference to the given []NamedObject and assigns it to the DelegatedResourceAdminGroups field.
func (o *ResourceGroup) SetDelegatedResourceAdminGroups(v []NamedObject) *ResourceGroup {
	o.DelegatedResourceAdminGroups = v
	return o
}

func (o ResourceGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.TeamId) {
		toSerialize["team_id"] = o.TeamId
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DelegatedResourceAdminGroups) {
		toSerialize["delegated_resource_admin_groups"] = o.DelegatedResourceAdminGroups
	}
	return toSerialize, nil
}

type NullableResourceGroup struct {
	value *ResourceGroup
	isSet bool
}

func (v NullableResourceGroup) Get() *ResourceGroup {
	return v.value
}

func (v *NullableResourceGroup) Set(val *ResourceGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceGroup(val *ResourceGroup) *NullableResourceGroup {
	return &NullableResourceGroup{value: val, isSet: true}
}

func (v NullableResourceGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
