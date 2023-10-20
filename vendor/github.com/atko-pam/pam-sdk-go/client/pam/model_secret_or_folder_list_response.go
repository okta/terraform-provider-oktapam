/*
Okta Privileged Access

The ScaleFT API is a control plane API for operations in Okta Privileged Access (formerly ScaleFT)

API version: 1.0.0
Contact: support@okta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pam

import (
	"encoding/json"
)

// checks if the SecretOrFolderListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecretOrFolderListResponse{}

// SecretOrFolderListResponse struct for SecretOrFolderListResponse
type SecretOrFolderListResponse struct {
	// The UUID of the Secret Folder
	Id string `json:"id"`
	// The name of the Secret Folder
	Name string `json:"name"`
	// A description of the Secret Folder
	Description   NullableString `json:"description,omitempty"`
	Type          SecretType     `json:"type"`
	ResourceGroup NamedObject    `json:"resource_group"`
	Project       NamedObject    `json:"project"`
}

// NewSecretOrFolderListResponse instantiates a new SecretOrFolderListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretOrFolderListResponse(id string, name string, type_ SecretType, resourceGroup NamedObject, project NamedObject) *SecretOrFolderListResponse {
	this := SecretOrFolderListResponse{}
	this.Id = id
	this.Name = name
	this.Type = type_
	this.ResourceGroup = resourceGroup
	this.Project = project
	return &this
}

// NewSecretOrFolderListResponseWithDefaults instantiates a new SecretOrFolderListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretOrFolderListResponseWithDefaults() *SecretOrFolderListResponse {
	this := SecretOrFolderListResponse{}
	return &this
}

// GetId returns the Id field value
func (o *SecretOrFolderListResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SecretOrFolderListResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SecretOrFolderListResponse) SetId(v string) *SecretOrFolderListResponse {
	o.Id = v
	return o
}

// GetName returns the Name field value
func (o *SecretOrFolderListResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SecretOrFolderListResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SecretOrFolderListResponse) SetName(v string) *SecretOrFolderListResponse {
	o.Name = v
	return o
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecretOrFolderListResponse) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecretOrFolderListResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *SecretOrFolderListResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *SecretOrFolderListResponse) SetDescription(v string) *SecretOrFolderListResponse {
	o.Description.Set(&v)
	return o
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *SecretOrFolderListResponse) SetDescriptionNil() *SecretOrFolderListResponse {
	o.Description.Set(nil)
	return o
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *SecretOrFolderListResponse) UnsetDescription() *SecretOrFolderListResponse {
	o.Description.Unset()
	return o
}

// GetType returns the Type field value
func (o *SecretOrFolderListResponse) GetType() SecretType {
	if o == nil {
		var ret SecretType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SecretOrFolderListResponse) GetTypeOk() (*SecretType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SecretOrFolderListResponse) SetType(v SecretType) *SecretOrFolderListResponse {
	o.Type = v
	return o
}

// GetResourceGroup returns the ResourceGroup field value
func (o *SecretOrFolderListResponse) GetResourceGroup() NamedObject {
	if o == nil {
		var ret NamedObject
		return ret
	}

	return o.ResourceGroup
}

// GetResourceGroupOk returns a tuple with the ResourceGroup field value
// and a boolean to check if the value has been set.
func (o *SecretOrFolderListResponse) GetResourceGroupOk() (*NamedObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceGroup, true
}

// SetResourceGroup sets field value
func (o *SecretOrFolderListResponse) SetResourceGroup(v NamedObject) *SecretOrFolderListResponse {
	o.ResourceGroup = v
	return o
}

// GetProject returns the Project field value
func (o *SecretOrFolderListResponse) GetProject() NamedObject {
	if o == nil {
		var ret NamedObject
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *SecretOrFolderListResponse) GetProjectOk() (*NamedObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value
func (o *SecretOrFolderListResponse) SetProject(v NamedObject) *SecretOrFolderListResponse {
	o.Project = v
	return o
}

func (o SecretOrFolderListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecretOrFolderListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["type"] = o.Type
	toSerialize["resource_group"] = o.ResourceGroup
	toSerialize["project"] = o.Project
	return toSerialize, nil
}

type NullableSecretOrFolderListResponse struct {
	value *SecretOrFolderListResponse
	isSet bool
}

func (v NullableSecretOrFolderListResponse) Get() *SecretOrFolderListResponse {
	return v.value
}

func (v *NullableSecretOrFolderListResponse) Set(val *SecretOrFolderListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretOrFolderListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretOrFolderListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretOrFolderListResponse(val *SecretOrFolderListResponse) *NullableSecretOrFolderListResponse {
	return &NullableSecretOrFolderListResponse{value: val, isSet: true}
}

func (v NullableSecretOrFolderListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretOrFolderListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
