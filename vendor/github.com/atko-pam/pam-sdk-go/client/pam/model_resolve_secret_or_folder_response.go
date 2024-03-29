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

// checks if the ResolveSecretOrFolderResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResolveSecretOrFolderResponse{}

// ResolveSecretOrFolderResponse struct for ResolveSecretOrFolderResponse
type ResolveSecretOrFolderResponse struct {
	// The UUID of the Secret or Secret Folder
	Id string `json:"id"`
	// The name of the Secret or Secret Folder
	Name string `json:"name"`
	// A description of the Secret or Secret Folder
	Description   NullableString `json:"description,omitempty"`
	Path          []SecretPath   `json:"path,omitempty"`
	ResourceGroup NamedObject    `json:"resource_group"`
	Project       NamedObject    `json:"project"`
	Type          *SecretType    `json:"type,omitempty"`
}

// NewResolveSecretOrFolderResponse instantiates a new ResolveSecretOrFolderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResolveSecretOrFolderResponse(id string, name string, resourceGroup NamedObject, project NamedObject) *ResolveSecretOrFolderResponse {
	this := ResolveSecretOrFolderResponse{}
	this.Id = id
	this.Name = name
	this.ResourceGroup = resourceGroup
	this.Project = project
	return &this
}

// NewResolveSecretOrFolderResponseWithDefaults instantiates a new ResolveSecretOrFolderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResolveSecretOrFolderResponseWithDefaults() *ResolveSecretOrFolderResponse {
	this := ResolveSecretOrFolderResponse{}
	return &this
}

// GetId returns the Id field value
func (o *ResolveSecretOrFolderResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResolveSecretOrFolderResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResolveSecretOrFolderResponse) SetId(v string) *ResolveSecretOrFolderResponse {
	o.Id = v
	return o
}

// GetName returns the Name field value
func (o *ResolveSecretOrFolderResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ResolveSecretOrFolderResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ResolveSecretOrFolderResponse) SetName(v string) *ResolveSecretOrFolderResponse {
	o.Name = v
	return o
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResolveSecretOrFolderResponse) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResolveSecretOrFolderResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ResolveSecretOrFolderResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ResolveSecretOrFolderResponse) SetDescription(v string) *ResolveSecretOrFolderResponse {
	o.Description.Set(&v)
	return o
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ResolveSecretOrFolderResponse) SetDescriptionNil() *ResolveSecretOrFolderResponse {
	o.Description.Set(nil)
	return o
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ResolveSecretOrFolderResponse) UnsetDescription() *ResolveSecretOrFolderResponse {
	o.Description.Unset()
	return o
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ResolveSecretOrFolderResponse) GetPath() []SecretPath {
	if o == nil || IsNil(o.Path) {
		var ret []SecretPath
		return ret
	}
	return o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResolveSecretOrFolderResponse) GetPathOk() ([]SecretPath, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ResolveSecretOrFolderResponse) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given []SecretPath and assigns it to the Path field.
func (o *ResolveSecretOrFolderResponse) SetPath(v []SecretPath) *ResolveSecretOrFolderResponse {
	o.Path = v
	return o
}

// GetResourceGroup returns the ResourceGroup field value
func (o *ResolveSecretOrFolderResponse) GetResourceGroup() NamedObject {
	if o == nil {
		var ret NamedObject
		return ret
	}

	return o.ResourceGroup
}

// GetResourceGroupOk returns a tuple with the ResourceGroup field value
// and a boolean to check if the value has been set.
func (o *ResolveSecretOrFolderResponse) GetResourceGroupOk() (*NamedObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceGroup, true
}

// SetResourceGroup sets field value
func (o *ResolveSecretOrFolderResponse) SetResourceGroup(v NamedObject) *ResolveSecretOrFolderResponse {
	o.ResourceGroup = v
	return o
}

// GetProject returns the Project field value
func (o *ResolveSecretOrFolderResponse) GetProject() NamedObject {
	if o == nil {
		var ret NamedObject
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *ResolveSecretOrFolderResponse) GetProjectOk() (*NamedObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value
func (o *ResolveSecretOrFolderResponse) SetProject(v NamedObject) *ResolveSecretOrFolderResponse {
	o.Project = v
	return o
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ResolveSecretOrFolderResponse) GetType() SecretType {
	if o == nil || IsNil(o.Type) {
		var ret SecretType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResolveSecretOrFolderResponse) GetTypeOk() (*SecretType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ResolveSecretOrFolderResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given SecretType and assigns it to the Type field.
func (o *ResolveSecretOrFolderResponse) SetType(v SecretType) *ResolveSecretOrFolderResponse {
	o.Type = &v
	return o
}

func (o ResolveSecretOrFolderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResolveSecretOrFolderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	toSerialize["resource_group"] = o.ResourceGroup
	toSerialize["project"] = o.Project
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableResolveSecretOrFolderResponse struct {
	value *ResolveSecretOrFolderResponse
	isSet bool
}

func (v NullableResolveSecretOrFolderResponse) Get() *ResolveSecretOrFolderResponse {
	return v.value
}

func (v *NullableResolveSecretOrFolderResponse) Set(val *ResolveSecretOrFolderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableResolveSecretOrFolderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResolveSecretOrFolderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResolveSecretOrFolderResponse(val *ResolveSecretOrFolderResponse) *NullableResolveSecretOrFolderResponse {
	return &NullableResolveSecretOrFolderResponse{value: val, isSet: true}
}

func (v NullableResolveSecretOrFolderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResolveSecretOrFolderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
