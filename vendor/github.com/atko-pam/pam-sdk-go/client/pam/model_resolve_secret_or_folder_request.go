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

// checks if the ResolveSecretOrFolderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResolveSecretOrFolderRequest{}

// ResolveSecretOrFolderRequest struct for ResolveSecretOrFolderRequest
type ResolveSecretOrFolderRequest struct {
	ResourceGroup SecretResolveParent `json:"resource_group"`
	Project       SecretResolveParent `json:"project"`
	// The path to the parent directory. Don't use this parameter if the request also includes an `id`.
	ParentFolderPath NullableString `json:"parent_folder_path,omitempty"`
	// The ID of the Secret or Secret Folder. Don't use this parameter if the request also includes a `parent_folder_path`, `secret_folder_name`, and `secret_name`.
	Id NullableString `json:"id,omitempty"`
	// The name of the Secret Folder. Don't use this parameter if the request also includes an `id` or `secret_name`.
	SecretFolderName *string `json:"secret_folder_name,omitempty"`
	// The name of the Secret. Don't use this parameter if the request also includes an `id` or `secret_folder_name`.
	SecretName *string `json:"secret_name,omitempty"`
}

// NewResolveSecretOrFolderRequest instantiates a new ResolveSecretOrFolderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResolveSecretOrFolderRequest(resourceGroup SecretResolveParent, project SecretResolveParent) *ResolveSecretOrFolderRequest {
	this := ResolveSecretOrFolderRequest{}
	this.ResourceGroup = resourceGroup
	this.Project = project
	return &this
}

// NewResolveSecretOrFolderRequestWithDefaults instantiates a new ResolveSecretOrFolderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResolveSecretOrFolderRequestWithDefaults() *ResolveSecretOrFolderRequest {
	this := ResolveSecretOrFolderRequest{}
	return &this
}

// GetResourceGroup returns the ResourceGroup field value
func (o *ResolveSecretOrFolderRequest) GetResourceGroup() SecretResolveParent {
	if o == nil {
		var ret SecretResolveParent
		return ret
	}

	return o.ResourceGroup
}

// GetResourceGroupOk returns a tuple with the ResourceGroup field value
// and a boolean to check if the value has been set.
func (o *ResolveSecretOrFolderRequest) GetResourceGroupOk() (*SecretResolveParent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceGroup, true
}

// SetResourceGroup sets field value
func (o *ResolveSecretOrFolderRequest) SetResourceGroup(v SecretResolveParent) *ResolveSecretOrFolderRequest {
	o.ResourceGroup = v
	return o
}

// GetProject returns the Project field value
func (o *ResolveSecretOrFolderRequest) GetProject() SecretResolveParent {
	if o == nil {
		var ret SecretResolveParent
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *ResolveSecretOrFolderRequest) GetProjectOk() (*SecretResolveParent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value
func (o *ResolveSecretOrFolderRequest) SetProject(v SecretResolveParent) *ResolveSecretOrFolderRequest {
	o.Project = v
	return o
}

// GetParentFolderPath returns the ParentFolderPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResolveSecretOrFolderRequest) GetParentFolderPath() string {
	if o == nil || IsNil(o.ParentFolderPath.Get()) {
		var ret string
		return ret
	}
	return *o.ParentFolderPath.Get()
}

// GetParentFolderPathOk returns a tuple with the ParentFolderPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResolveSecretOrFolderRequest) GetParentFolderPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentFolderPath.Get(), o.ParentFolderPath.IsSet()
}

// HasParentFolderPath returns a boolean if a field has been set.
func (o *ResolveSecretOrFolderRequest) HasParentFolderPath() bool {
	if o != nil && o.ParentFolderPath.IsSet() {
		return true
	}

	return false
}

// SetParentFolderPath gets a reference to the given NullableString and assigns it to the ParentFolderPath field.
func (o *ResolveSecretOrFolderRequest) SetParentFolderPath(v string) *ResolveSecretOrFolderRequest {
	o.ParentFolderPath.Set(&v)
	return o
}

// SetParentFolderPathNil sets the value for ParentFolderPath to be an explicit nil
func (o *ResolveSecretOrFolderRequest) SetParentFolderPathNil() *ResolveSecretOrFolderRequest {
	o.ParentFolderPath.Set(nil)
	return o
}

// UnsetParentFolderPath ensures that no value is present for ParentFolderPath, not even an explicit nil
func (o *ResolveSecretOrFolderRequest) UnsetParentFolderPath() *ResolveSecretOrFolderRequest {
	o.ParentFolderPath.Unset()
	return o
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResolveSecretOrFolderRequest) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResolveSecretOrFolderRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ResolveSecretOrFolderRequest) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *ResolveSecretOrFolderRequest) SetId(v string) *ResolveSecretOrFolderRequest {
	o.Id.Set(&v)
	return o
}

// SetIdNil sets the value for Id to be an explicit nil
func (o *ResolveSecretOrFolderRequest) SetIdNil() *ResolveSecretOrFolderRequest {
	o.Id.Set(nil)
	return o
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ResolveSecretOrFolderRequest) UnsetId() *ResolveSecretOrFolderRequest {
	o.Id.Unset()
	return o
}

// GetSecretFolderName returns the SecretFolderName field value if set, zero value otherwise.
func (o *ResolveSecretOrFolderRequest) GetSecretFolderName() string {
	if o == nil || IsNil(o.SecretFolderName) {
		var ret string
		return ret
	}
	return *o.SecretFolderName
}

// GetSecretFolderNameOk returns a tuple with the SecretFolderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResolveSecretOrFolderRequest) GetSecretFolderNameOk() (*string, bool) {
	if o == nil || IsNil(o.SecretFolderName) {
		return nil, false
	}
	return o.SecretFolderName, true
}

// HasSecretFolderName returns a boolean if a field has been set.
func (o *ResolveSecretOrFolderRequest) HasSecretFolderName() bool {
	if o != nil && !IsNil(o.SecretFolderName) {
		return true
	}

	return false
}

// SetSecretFolderName gets a reference to the given string and assigns it to the SecretFolderName field.
func (o *ResolveSecretOrFolderRequest) SetSecretFolderName(v string) *ResolveSecretOrFolderRequest {
	o.SecretFolderName = &v
	return o
}

// GetSecretName returns the SecretName field value if set, zero value otherwise.
func (o *ResolveSecretOrFolderRequest) GetSecretName() string {
	if o == nil || IsNil(o.SecretName) {
		var ret string
		return ret
	}
	return *o.SecretName
}

// GetSecretNameOk returns a tuple with the SecretName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResolveSecretOrFolderRequest) GetSecretNameOk() (*string, bool) {
	if o == nil || IsNil(o.SecretName) {
		return nil, false
	}
	return o.SecretName, true
}

// HasSecretName returns a boolean if a field has been set.
func (o *ResolveSecretOrFolderRequest) HasSecretName() bool {
	if o != nil && !IsNil(o.SecretName) {
		return true
	}

	return false
}

// SetSecretName gets a reference to the given string and assigns it to the SecretName field.
func (o *ResolveSecretOrFolderRequest) SetSecretName(v string) *ResolveSecretOrFolderRequest {
	o.SecretName = &v
	return o
}

func (o ResolveSecretOrFolderRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResolveSecretOrFolderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resource_group"] = o.ResourceGroup
	toSerialize["project"] = o.Project
	if o.ParentFolderPath.IsSet() {
		toSerialize["parent_folder_path"] = o.ParentFolderPath.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if !IsNil(o.SecretFolderName) {
		toSerialize["secret_folder_name"] = o.SecretFolderName
	}
	if !IsNil(o.SecretName) {
		toSerialize["secret_name"] = o.SecretName
	}
	return toSerialize, nil
}

type NullableResolveSecretOrFolderRequest struct {
	value *ResolveSecretOrFolderRequest
	isSet bool
}

func (v NullableResolveSecretOrFolderRequest) Get() *ResolveSecretOrFolderRequest {
	return v.value
}

func (v *NullableResolveSecretOrFolderRequest) Set(val *ResolveSecretOrFolderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableResolveSecretOrFolderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableResolveSecretOrFolderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResolveSecretOrFolderRequest(val *ResolveSecretOrFolderRequest) *NullableResolveSecretOrFolderRequest {
	return &NullableResolveSecretOrFolderRequest{value: val, isSet: true}
}

func (v NullableResolveSecretOrFolderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResolveSecretOrFolderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
