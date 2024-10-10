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

// checks if the ActiveDirectoryAccountRuleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActiveDirectoryAccountRuleResponse{}

// ActiveDirectoryAccountRuleResponse struct for ActiveDirectoryAccountRuleResponse
type ActiveDirectoryAccountRuleResponse struct {
	Id interface{} `json:"id,omitempty"`
	// The name of the Active Directory account rule
	Name *string `json:"name,omitempty"`
	// The type of the Active Directory account rule
	RuleType *string `json:"rule_type,omitempty"`
	// The list of organizational units to discover individual accounts from
	OrganizationalUnits []string `json:"organizational_units,omitempty"`
	// The priority of the Active Directory account rule
	Priority      *int32       `json:"priority,omitempty"`
	ResourceGroup *NamedObject `json:"resource_group,omitempty"`
	Project       *NamedObject `json:"project,omitempty"`
}

// NewActiveDirectoryAccountRuleResponse instantiates a new ActiveDirectoryAccountRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActiveDirectoryAccountRuleResponse() *ActiveDirectoryAccountRuleResponse {
	this := ActiveDirectoryAccountRuleResponse{}
	return &this
}

// NewActiveDirectoryAccountRuleResponseWithDefaults instantiates a new ActiveDirectoryAccountRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActiveDirectoryAccountRuleResponseWithDefaults() *ActiveDirectoryAccountRuleResponse {
	this := ActiveDirectoryAccountRuleResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveDirectoryAccountRuleResponse) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveDirectoryAccountRuleResponse) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ActiveDirectoryAccountRuleResponse) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *ActiveDirectoryAccountRuleResponse) SetId(v interface{}) *ActiveDirectoryAccountRuleResponse {
	o.Id = v
	return o
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ActiveDirectoryAccountRuleResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveDirectoryAccountRuleResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ActiveDirectoryAccountRuleResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ActiveDirectoryAccountRuleResponse) SetName(v string) *ActiveDirectoryAccountRuleResponse {
	o.Name = &v
	return o
}

// GetRuleType returns the RuleType field value if set, zero value otherwise.
func (o *ActiveDirectoryAccountRuleResponse) GetRuleType() string {
	if o == nil || IsNil(o.RuleType) {
		var ret string
		return ret
	}
	return *o.RuleType
}

// GetRuleTypeOk returns a tuple with the RuleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveDirectoryAccountRuleResponse) GetRuleTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RuleType) {
		return nil, false
	}
	return o.RuleType, true
}

// HasRuleType returns a boolean if a field has been set.
func (o *ActiveDirectoryAccountRuleResponse) HasRuleType() bool {
	if o != nil && !IsNil(o.RuleType) {
		return true
	}

	return false
}

// SetRuleType gets a reference to the given string and assigns it to the RuleType field.
func (o *ActiveDirectoryAccountRuleResponse) SetRuleType(v string) *ActiveDirectoryAccountRuleResponse {
	o.RuleType = &v
	return o
}

// GetOrganizationalUnits returns the OrganizationalUnits field value if set, zero value otherwise.
func (o *ActiveDirectoryAccountRuleResponse) GetOrganizationalUnits() []string {
	if o == nil || IsNil(o.OrganizationalUnits) {
		var ret []string
		return ret
	}
	return o.OrganizationalUnits
}

// GetOrganizationalUnitsOk returns a tuple with the OrganizationalUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveDirectoryAccountRuleResponse) GetOrganizationalUnitsOk() ([]string, bool) {
	if o == nil || IsNil(o.OrganizationalUnits) {
		return nil, false
	}
	return o.OrganizationalUnits, true
}

// HasOrganizationalUnits returns a boolean if a field has been set.
func (o *ActiveDirectoryAccountRuleResponse) HasOrganizationalUnits() bool {
	if o != nil && !IsNil(o.OrganizationalUnits) {
		return true
	}

	return false
}

// SetOrganizationalUnits gets a reference to the given []string and assigns it to the OrganizationalUnits field.
func (o *ActiveDirectoryAccountRuleResponse) SetOrganizationalUnits(v []string) *ActiveDirectoryAccountRuleResponse {
	o.OrganizationalUnits = v
	return o
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *ActiveDirectoryAccountRuleResponse) GetPriority() int32 {
	if o == nil || IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveDirectoryAccountRuleResponse) GetPriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *ActiveDirectoryAccountRuleResponse) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *ActiveDirectoryAccountRuleResponse) SetPriority(v int32) *ActiveDirectoryAccountRuleResponse {
	o.Priority = &v
	return o
}

// GetResourceGroup returns the ResourceGroup field value if set, zero value otherwise.
func (o *ActiveDirectoryAccountRuleResponse) GetResourceGroup() NamedObject {
	if o == nil || IsNil(o.ResourceGroup) {
		var ret NamedObject
		return ret
	}
	return *o.ResourceGroup
}

// GetResourceGroupOk returns a tuple with the ResourceGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveDirectoryAccountRuleResponse) GetResourceGroupOk() (*NamedObject, bool) {
	if o == nil || IsNil(o.ResourceGroup) {
		return nil, false
	}
	return o.ResourceGroup, true
}

// HasResourceGroup returns a boolean if a field has been set.
func (o *ActiveDirectoryAccountRuleResponse) HasResourceGroup() bool {
	if o != nil && !IsNil(o.ResourceGroup) {
		return true
	}

	return false
}

// SetResourceGroup gets a reference to the given NamedObject and assigns it to the ResourceGroup field.
func (o *ActiveDirectoryAccountRuleResponse) SetResourceGroup(v NamedObject) *ActiveDirectoryAccountRuleResponse {
	o.ResourceGroup = &v
	return o
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *ActiveDirectoryAccountRuleResponse) GetProject() NamedObject {
	if o == nil || IsNil(o.Project) {
		var ret NamedObject
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveDirectoryAccountRuleResponse) GetProjectOk() (*NamedObject, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *ActiveDirectoryAccountRuleResponse) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given NamedObject and assigns it to the Project field.
func (o *ActiveDirectoryAccountRuleResponse) SetProject(v NamedObject) *ActiveDirectoryAccountRuleResponse {
	o.Project = &v
	return o
}

func (o ActiveDirectoryAccountRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActiveDirectoryAccountRuleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.RuleType) {
		toSerialize["rule_type"] = o.RuleType
	}
	if !IsNil(o.OrganizationalUnits) {
		toSerialize["organizational_units"] = o.OrganizationalUnits
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.ResourceGroup) {
		toSerialize["resource_group"] = o.ResourceGroup
	}
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	return toSerialize, nil
}

type NullableActiveDirectoryAccountRuleResponse struct {
	value *ActiveDirectoryAccountRuleResponse
	isSet bool
}

func (v NullableActiveDirectoryAccountRuleResponse) Get() *ActiveDirectoryAccountRuleResponse {
	return v.value
}

func (v *NullableActiveDirectoryAccountRuleResponse) Set(val *ActiveDirectoryAccountRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableActiveDirectoryAccountRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableActiveDirectoryAccountRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActiveDirectoryAccountRuleResponse(val *ActiveDirectoryAccountRuleResponse) *NullableActiveDirectoryAccountRuleResponse {
	return &NullableActiveDirectoryAccountRuleResponse{value: val, isSet: true}
}

func (v NullableActiveDirectoryAccountRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActiveDirectoryAccountRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
