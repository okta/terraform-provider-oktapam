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
	"time"
)

// checks if the SecretFolderResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecretFolderResponse{}

// SecretFolderResponse struct for SecretFolderResponse
type SecretFolderResponse struct {
	// The UUID of the Secret Folder
	Id string `json:"id"`
	// The name of the Secret Folder
	Name string `json:"name"`
	// A description of the Secret Folder
	Description NullableString `json:"description,omitempty"`
	Path        []SecretPath   `json:"path,omitempty"`
	// A timestamp indicating when the Secret Folder was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The username of the User who created the Secret Folder
	CreatedBy *string `json:"created_by,omitempty"`
	// A timestamp indicating when the Secret Folder was last updated
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The username of the User who last updated the Secret Folder
	UpdatedBy *string `json:"updated_by,omitempty"`
}

// NewSecretFolderResponse instantiates a new SecretFolderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretFolderResponse(id string, name string) *SecretFolderResponse {
	this := SecretFolderResponse{}
	this.Id = id
	this.Name = name
	return &this
}

// NewSecretFolderResponseWithDefaults instantiates a new SecretFolderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretFolderResponseWithDefaults() *SecretFolderResponse {
	this := SecretFolderResponse{}
	return &this
}

// GetId returns the Id field value
func (o *SecretFolderResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SecretFolderResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SecretFolderResponse) SetId(v string) *SecretFolderResponse {
	o.Id = v
	return o
}

// GetName returns the Name field value
func (o *SecretFolderResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SecretFolderResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SecretFolderResponse) SetName(v string) *SecretFolderResponse {
	o.Name = v
	return o
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecretFolderResponse) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecretFolderResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *SecretFolderResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *SecretFolderResponse) SetDescription(v string) *SecretFolderResponse {
	o.Description.Set(&v)
	return o
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *SecretFolderResponse) SetDescriptionNil() *SecretFolderResponse {
	o.Description.Set(nil)
	return o
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *SecretFolderResponse) UnsetDescription() *SecretFolderResponse {
	o.Description.Unset()
	return o
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *SecretFolderResponse) GetPath() []SecretPath {
	if o == nil || IsNil(o.Path) {
		var ret []SecretPath
		return ret
	}
	return o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretFolderResponse) GetPathOk() ([]SecretPath, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *SecretFolderResponse) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given []SecretPath and assigns it to the Path field.
func (o *SecretFolderResponse) SetPath(v []SecretPath) *SecretFolderResponse {
	o.Path = v
	return o
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SecretFolderResponse) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretFolderResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SecretFolderResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *SecretFolderResponse) SetCreatedAt(v time.Time) *SecretFolderResponse {
	o.CreatedAt = &v
	return o
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *SecretFolderResponse) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretFolderResponse) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *SecretFolderResponse) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *SecretFolderResponse) SetCreatedBy(v string) *SecretFolderResponse {
	o.CreatedBy = &v
	return o
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *SecretFolderResponse) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretFolderResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *SecretFolderResponse) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *SecretFolderResponse) SetUpdatedAt(v time.Time) *SecretFolderResponse {
	o.UpdatedAt = &v
	return o
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *SecretFolderResponse) GetUpdatedBy() string {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretFolderResponse) GetUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *SecretFolderResponse) HasUpdatedBy() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *SecretFolderResponse) SetUpdatedBy(v string) *SecretFolderResponse {
	o.UpdatedBy = &v
	return o
}

func (o SecretFolderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecretFolderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["created_by"] = o.CreatedBy
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.UpdatedBy) {
		toSerialize["updated_by"] = o.UpdatedBy
	}
	return toSerialize, nil
}

type NullableSecretFolderResponse struct {
	value *SecretFolderResponse
	isSet bool
}

func (v NullableSecretFolderResponse) Get() *SecretFolderResponse {
	return v.value
}

func (v *NullableSecretFolderResponse) Set(val *SecretFolderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretFolderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretFolderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretFolderResponse(val *SecretFolderResponse) *NullableSecretFolderResponse {
	return &NullableSecretFolderResponse{value: val, isSet: true}
}

func (v NullableSecretFolderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretFolderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
