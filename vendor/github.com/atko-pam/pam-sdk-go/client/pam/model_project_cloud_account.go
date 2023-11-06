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
	"time"
)

// checks if the ProjectCloudAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectCloudAccount{}

// ProjectCloudAccount struct for ProjectCloudAccount
type ProjectCloudAccount struct {
	AccountDetails ProjectCloudAccountAccountDetails `json:"account_details"`
	// The provider-specific account ID
	AccountId string `json:"account_id"`
	// The Client ID used to authorize the EventBridge connection
	ClientId NullableString `json:"client_id"`
	// The Client secret used to authorize the EventBridge connection
	ClientSecret NullableString `json:"client_secret"`
	// A timestamp indicating when the Cloud Account was deleted. `null` if not deleted.
	DeletedAt time.Time `json:"deleted_at"`
	// The human-readable description of the Cloud Account
	Description NullableString `json:"description,omitempty"`
	// The UUID of the Cloud Account
	Id string `json:"id"`
	// The name of the associated Project
	ProjectName string               `json:"project_name"`
	Provider    CloudAccountProvider `json:"provider"`
	// A timestamp indicating when the secret was last rotated. `null` if unrotated.
	SecretLastRotated time.Time `json:"secret_last_rotated"`
	// A timestamp indicating when the Cloud Account was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// NewProjectCloudAccount instantiates a new ProjectCloudAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectCloudAccount(accountDetails ProjectCloudAccountAccountDetails, accountId string, clientId NullableString, clientSecret NullableString, deletedAt time.Time, id string, projectName string, provider CloudAccountProvider, secretLastRotated time.Time, updatedAt time.Time) *ProjectCloudAccount {
	this := ProjectCloudAccount{}
	this.AccountDetails = accountDetails
	this.AccountId = accountId
	this.ClientId = clientId
	this.ClientSecret = clientSecret
	this.DeletedAt = deletedAt
	this.Id = id
	this.ProjectName = projectName
	this.Provider = provider
	this.SecretLastRotated = secretLastRotated
	this.UpdatedAt = updatedAt
	return &this
}

// NewProjectCloudAccountWithDefaults instantiates a new ProjectCloudAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectCloudAccountWithDefaults() *ProjectCloudAccount {
	this := ProjectCloudAccount{}
	return &this
}

// GetAccountDetails returns the AccountDetails field value
func (o *ProjectCloudAccount) GetAccountDetails() ProjectCloudAccountAccountDetails {
	if o == nil {
		var ret ProjectCloudAccountAccountDetails
		return ret
	}

	return o.AccountDetails
}

// GetAccountDetailsOk returns a tuple with the AccountDetails field value
// and a boolean to check if the value has been set.
func (o *ProjectCloudAccount) GetAccountDetailsOk() (*ProjectCloudAccountAccountDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountDetails, true
}

// SetAccountDetails sets field value
func (o *ProjectCloudAccount) SetAccountDetails(v ProjectCloudAccountAccountDetails) *ProjectCloudAccount {
	o.AccountDetails = v
	return o
}

// GetAccountId returns the AccountId field value
func (o *ProjectCloudAccount) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *ProjectCloudAccount) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *ProjectCloudAccount) SetAccountId(v string) *ProjectCloudAccount {
	o.AccountId = v
	return o
}

// GetClientId returns the ClientId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ProjectCloudAccount) GetClientId() string {
	if o == nil || o.ClientId.Get() == nil {
		var ret string
		return ret
	}

	return *o.ClientId.Get()
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectCloudAccount) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientId.Get(), o.ClientId.IsSet()
}

// SetClientId sets field value
func (o *ProjectCloudAccount) SetClientId(v string) *ProjectCloudAccount {
	o.ClientId.Set(&v)
	return o
}

// GetClientSecret returns the ClientSecret field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ProjectCloudAccount) GetClientSecret() string {
	if o == nil || o.ClientSecret.Get() == nil {
		var ret string
		return ret
	}

	return *o.ClientSecret.Get()
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectCloudAccount) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientSecret.Get(), o.ClientSecret.IsSet()
}

// SetClientSecret sets field value
func (o *ProjectCloudAccount) SetClientSecret(v string) *ProjectCloudAccount {
	o.ClientSecret.Set(&v)
	return o
}

// GetDeletedAt returns the DeletedAt field value
func (o *ProjectCloudAccount) GetDeletedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value
// and a boolean to check if the value has been set.
func (o *ProjectCloudAccount) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeletedAt, true
}

// SetDeletedAt sets field value
func (o *ProjectCloudAccount) SetDeletedAt(v time.Time) *ProjectCloudAccount {
	o.DeletedAt = v
	return o
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectCloudAccount) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectCloudAccount) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ProjectCloudAccount) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ProjectCloudAccount) SetDescription(v string) *ProjectCloudAccount {
	o.Description.Set(&v)
	return o
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ProjectCloudAccount) SetDescriptionNil() *ProjectCloudAccount {
	o.Description.Set(nil)
	return o
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ProjectCloudAccount) UnsetDescription() *ProjectCloudAccount {
	o.Description.Unset()
	return o
}

// GetId returns the Id field value
func (o *ProjectCloudAccount) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProjectCloudAccount) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProjectCloudAccount) SetId(v string) *ProjectCloudAccount {
	o.Id = v
	return o
}

// GetProjectName returns the ProjectName field value
func (o *ProjectCloudAccount) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *ProjectCloudAccount) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *ProjectCloudAccount) SetProjectName(v string) *ProjectCloudAccount {
	o.ProjectName = v
	return o
}

// GetProvider returns the Provider field value
func (o *ProjectCloudAccount) GetProvider() CloudAccountProvider {
	if o == nil {
		var ret CloudAccountProvider
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *ProjectCloudAccount) GetProviderOk() (*CloudAccountProvider, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *ProjectCloudAccount) SetProvider(v CloudAccountProvider) *ProjectCloudAccount {
	o.Provider = v
	return o
}

// GetSecretLastRotated returns the SecretLastRotated field value
func (o *ProjectCloudAccount) GetSecretLastRotated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.SecretLastRotated
}

// GetSecretLastRotatedOk returns a tuple with the SecretLastRotated field value
// and a boolean to check if the value has been set.
func (o *ProjectCloudAccount) GetSecretLastRotatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretLastRotated, true
}

// SetSecretLastRotated sets field value
func (o *ProjectCloudAccount) SetSecretLastRotated(v time.Time) *ProjectCloudAccount {
	o.SecretLastRotated = v
	return o
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ProjectCloudAccount) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ProjectCloudAccount) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ProjectCloudAccount) SetUpdatedAt(v time.Time) *ProjectCloudAccount {
	o.UpdatedAt = v
	return o
}

func (o ProjectCloudAccount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectCloudAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_details"] = o.AccountDetails
	toSerialize["account_id"] = o.AccountId
	toSerialize["client_id"] = o.ClientId.Get()
	toSerialize["client_secret"] = o.ClientSecret.Get()
	toSerialize["deleted_at"] = o.DeletedAt
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["id"] = o.Id
	toSerialize["project_name"] = o.ProjectName
	toSerialize["provider"] = o.Provider
	toSerialize["secret_last_rotated"] = o.SecretLastRotated
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

type NullableProjectCloudAccount struct {
	value *ProjectCloudAccount
	isSet bool
}

func (v NullableProjectCloudAccount) Get() *ProjectCloudAccount {
	return v.value
}

func (v *NullableProjectCloudAccount) Set(val *ProjectCloudAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectCloudAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectCloudAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectCloudAccount(val *ProjectCloudAccount) *NullableProjectCloudAccount {
	return &NullableProjectCloudAccount{value: val, isSet: true}
}

func (v NullableProjectCloudAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectCloudAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
