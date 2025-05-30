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

// checks if the SaasAppAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SaasAppAccount{}

// SaasAppAccount struct for SaasAppAccount
type SaasAppAccount struct {
	// The UUID of the SaaS Application Account
	Id *string `json:"id,omitempty"`
	// A human-readable name for the SaaS Application Account
	Name *string `json:"name,omitempty"`
	// The username used to log into the SaaS Application
	Username *string `json:"username,omitempty"`
	// A brief description of the SaaS Application Account
	Description  *string                     `json:"description,omitempty"`
	Status       *ServiceAccountStatus       `json:"status,omitempty"`
	StatusDetail *ServiceAccountStatusDetail `json:"status_detail,omitempty"`
	SyncStatus   *ServiceAccountSyncStatus   `json:"sync_status,omitempty"`
	// Whether the password for the SaaS Application Account can be rotated using Okta Lifecycle Management
	LcmSyncPossible *bool `json:"lcm_sync_possible,omitempty"`
	// A URL pointing to the logo of the SaaS Application
	LogoUrl *string `json:"logo_url,omitempty"`
	// The name of the SaaS Application instance
	ApplicationInstanceName *string `json:"application_instance_name,omitempty"`
	// A URL pointing to the login page of the SaaS Application
	LoginUrl *string `json:"login_url,omitempty"`
	// The Okta app instance ID of the SaaS Application
	ApplicationInstanceId *string `json:"application_instance_id,omitempty"`
	// A timestamp that indicates when the OPA managed SaaS Application Account was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// A timestamp that indicates when the OPA managed SaaS Application Account was updated
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewSaasAppAccount instantiates a new SaasAppAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSaasAppAccount() *SaasAppAccount {
	this := SaasAppAccount{}
	return &this
}

// NewSaasAppAccountWithDefaults instantiates a new SaasAppAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSaasAppAccountWithDefaults() *SaasAppAccount {
	this := SaasAppAccount{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SaasAppAccount) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaasAppAccount) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SaasAppAccount) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SaasAppAccount) SetId(v string) *SaasAppAccount {
	o.Id = &v
	return o
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SaasAppAccount) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaasAppAccount) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SaasAppAccount) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SaasAppAccount) SetName(v string) *SaasAppAccount {
	o.Name = &v
	return o
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *SaasAppAccount) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaasAppAccount) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *SaasAppAccount) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *SaasAppAccount) SetUsername(v string) *SaasAppAccount {
	o.Username = &v
	return o
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SaasAppAccount) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaasAppAccount) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SaasAppAccount) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SaasAppAccount) SetDescription(v string) *SaasAppAccount {
	o.Description = &v
	return o
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SaasAppAccount) GetStatus() ServiceAccountStatus {
	if o == nil || IsNil(o.Status) {
		var ret ServiceAccountStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaasAppAccount) GetStatusOk() (*ServiceAccountStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SaasAppAccount) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ServiceAccountStatus and assigns it to the Status field.
func (o *SaasAppAccount) SetStatus(v ServiceAccountStatus) *SaasAppAccount {
	o.Status = &v
	return o
}

// GetStatusDetail returns the StatusDetail field value if set, zero value otherwise.
func (o *SaasAppAccount) GetStatusDetail() ServiceAccountStatusDetail {
	if o == nil || IsNil(o.StatusDetail) {
		var ret ServiceAccountStatusDetail
		return ret
	}
	return *o.StatusDetail
}

// GetStatusDetailOk returns a tuple with the StatusDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaasAppAccount) GetStatusDetailOk() (*ServiceAccountStatusDetail, bool) {
	if o == nil || IsNil(o.StatusDetail) {
		return nil, false
	}
	return o.StatusDetail, true
}

// HasStatusDetail returns a boolean if a field has been set.
func (o *SaasAppAccount) HasStatusDetail() bool {
	if o != nil && !IsNil(o.StatusDetail) {
		return true
	}

	return false
}

// SetStatusDetail gets a reference to the given ServiceAccountStatusDetail and assigns it to the StatusDetail field.
func (o *SaasAppAccount) SetStatusDetail(v ServiceAccountStatusDetail) *SaasAppAccount {
	o.StatusDetail = &v
	return o
}

// GetSyncStatus returns the SyncStatus field value if set, zero value otherwise.
func (o *SaasAppAccount) GetSyncStatus() ServiceAccountSyncStatus {
	if o == nil || IsNil(o.SyncStatus) {
		var ret ServiceAccountSyncStatus
		return ret
	}
	return *o.SyncStatus
}

// GetSyncStatusOk returns a tuple with the SyncStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaasAppAccount) GetSyncStatusOk() (*ServiceAccountSyncStatus, bool) {
	if o == nil || IsNil(o.SyncStatus) {
		return nil, false
	}
	return o.SyncStatus, true
}

// HasSyncStatus returns a boolean if a field has been set.
func (o *SaasAppAccount) HasSyncStatus() bool {
	if o != nil && !IsNil(o.SyncStatus) {
		return true
	}

	return false
}

// SetSyncStatus gets a reference to the given ServiceAccountSyncStatus and assigns it to the SyncStatus field.
func (o *SaasAppAccount) SetSyncStatus(v ServiceAccountSyncStatus) *SaasAppAccount {
	o.SyncStatus = &v
	return o
}

// GetLcmSyncPossible returns the LcmSyncPossible field value if set, zero value otherwise.
func (o *SaasAppAccount) GetLcmSyncPossible() bool {
	if o == nil || IsNil(o.LcmSyncPossible) {
		var ret bool
		return ret
	}
	return *o.LcmSyncPossible
}

// GetLcmSyncPossibleOk returns a tuple with the LcmSyncPossible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaasAppAccount) GetLcmSyncPossibleOk() (*bool, bool) {
	if o == nil || IsNil(o.LcmSyncPossible) {
		return nil, false
	}
	return o.LcmSyncPossible, true
}

// HasLcmSyncPossible returns a boolean if a field has been set.
func (o *SaasAppAccount) HasLcmSyncPossible() bool {
	if o != nil && !IsNil(o.LcmSyncPossible) {
		return true
	}

	return false
}

// SetLcmSyncPossible gets a reference to the given bool and assigns it to the LcmSyncPossible field.
func (o *SaasAppAccount) SetLcmSyncPossible(v bool) *SaasAppAccount {
	o.LcmSyncPossible = &v
	return o
}

// GetLogoUrl returns the LogoUrl field value if set, zero value otherwise.
func (o *SaasAppAccount) GetLogoUrl() string {
	if o == nil || IsNil(o.LogoUrl) {
		var ret string
		return ret
	}
	return *o.LogoUrl
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaasAppAccount) GetLogoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LogoUrl) {
		return nil, false
	}
	return o.LogoUrl, true
}

// HasLogoUrl returns a boolean if a field has been set.
func (o *SaasAppAccount) HasLogoUrl() bool {
	if o != nil && !IsNil(o.LogoUrl) {
		return true
	}

	return false
}

// SetLogoUrl gets a reference to the given string and assigns it to the LogoUrl field.
func (o *SaasAppAccount) SetLogoUrl(v string) *SaasAppAccount {
	o.LogoUrl = &v
	return o
}

// GetApplicationInstanceName returns the ApplicationInstanceName field value if set, zero value otherwise.
func (o *SaasAppAccount) GetApplicationInstanceName() string {
	if o == nil || IsNil(o.ApplicationInstanceName) {
		var ret string
		return ret
	}
	return *o.ApplicationInstanceName
}

// GetApplicationInstanceNameOk returns a tuple with the ApplicationInstanceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaasAppAccount) GetApplicationInstanceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationInstanceName) {
		return nil, false
	}
	return o.ApplicationInstanceName, true
}

// HasApplicationInstanceName returns a boolean if a field has been set.
func (o *SaasAppAccount) HasApplicationInstanceName() bool {
	if o != nil && !IsNil(o.ApplicationInstanceName) {
		return true
	}

	return false
}

// SetApplicationInstanceName gets a reference to the given string and assigns it to the ApplicationInstanceName field.
func (o *SaasAppAccount) SetApplicationInstanceName(v string) *SaasAppAccount {
	o.ApplicationInstanceName = &v
	return o
}

// GetLoginUrl returns the LoginUrl field value if set, zero value otherwise.
func (o *SaasAppAccount) GetLoginUrl() string {
	if o == nil || IsNil(o.LoginUrl) {
		var ret string
		return ret
	}
	return *o.LoginUrl
}

// GetLoginUrlOk returns a tuple with the LoginUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaasAppAccount) GetLoginUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LoginUrl) {
		return nil, false
	}
	return o.LoginUrl, true
}

// HasLoginUrl returns a boolean if a field has been set.
func (o *SaasAppAccount) HasLoginUrl() bool {
	if o != nil && !IsNil(o.LoginUrl) {
		return true
	}

	return false
}

// SetLoginUrl gets a reference to the given string and assigns it to the LoginUrl field.
func (o *SaasAppAccount) SetLoginUrl(v string) *SaasAppAccount {
	o.LoginUrl = &v
	return o
}

// GetApplicationInstanceId returns the ApplicationInstanceId field value if set, zero value otherwise.
func (o *SaasAppAccount) GetApplicationInstanceId() string {
	if o == nil || IsNil(o.ApplicationInstanceId) {
		var ret string
		return ret
	}
	return *o.ApplicationInstanceId
}

// GetApplicationInstanceIdOk returns a tuple with the ApplicationInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaasAppAccount) GetApplicationInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationInstanceId) {
		return nil, false
	}
	return o.ApplicationInstanceId, true
}

// HasApplicationInstanceId returns a boolean if a field has been set.
func (o *SaasAppAccount) HasApplicationInstanceId() bool {
	if o != nil && !IsNil(o.ApplicationInstanceId) {
		return true
	}

	return false
}

// SetApplicationInstanceId gets a reference to the given string and assigns it to the ApplicationInstanceId field.
func (o *SaasAppAccount) SetApplicationInstanceId(v string) *SaasAppAccount {
	o.ApplicationInstanceId = &v
	return o
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SaasAppAccount) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaasAppAccount) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SaasAppAccount) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *SaasAppAccount) SetCreatedAt(v time.Time) *SaasAppAccount {
	o.CreatedAt = &v
	return o
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *SaasAppAccount) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaasAppAccount) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *SaasAppAccount) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *SaasAppAccount) SetUpdatedAt(v time.Time) *SaasAppAccount {
	o.UpdatedAt = &v
	return o
}

func (o SaasAppAccount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SaasAppAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StatusDetail) {
		toSerialize["status_detail"] = o.StatusDetail
	}
	if !IsNil(o.SyncStatus) {
		toSerialize["sync_status"] = o.SyncStatus
	}
	if !IsNil(o.LcmSyncPossible) {
		toSerialize["lcm_sync_possible"] = o.LcmSyncPossible
	}
	if !IsNil(o.LogoUrl) {
		toSerialize["logo_url"] = o.LogoUrl
	}
	if !IsNil(o.ApplicationInstanceName) {
		toSerialize["application_instance_name"] = o.ApplicationInstanceName
	}
	if !IsNil(o.LoginUrl) {
		toSerialize["login_url"] = o.LoginUrl
	}
	if !IsNil(o.ApplicationInstanceId) {
		toSerialize["application_instance_id"] = o.ApplicationInstanceId
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableSaasAppAccount struct {
	value *SaasAppAccount
	isSet bool
}

func (v NullableSaasAppAccount) Get() *SaasAppAccount {
	return v.value
}

func (v *NullableSaasAppAccount) Set(val *SaasAppAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableSaasAppAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableSaasAppAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSaasAppAccount(val *SaasAppAccount) *NullableSaasAppAccount {
	return &NullableSaasAppAccount{value: val, isSet: true}
}

func (v NullableSaasAppAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSaasAppAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
