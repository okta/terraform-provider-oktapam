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

// checks if the OktaUniversalDirectoryAccountWithSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OktaUniversalDirectoryAccountWithSettings{}

// OktaUniversalDirectoryAccountWithSettings struct for OktaUniversalDirectoryAccountWithSettings
type OktaUniversalDirectoryAccountWithSettings struct {
	// The project level account settings that are enabled for this account
	AccountSettingsEnabled []AccountSetting `json:"account_settings_enabled,omitempty"`
	// The UUID of the Okta Universal Directory Account
	Id *string `json:"id,omitempty"`
	// A human-readable name for the Okta Universal Directory Account
	Name *string `json:"name,omitempty"`
	// The username used to log into Okta
	Username *string `json:"username,omitempty"`
	// A brief description of the Okta Universal Directory Account
	Description  *string                     `json:"description,omitempty"`
	Status       *ServiceAccountStatus       `json:"status,omitempty"`
	StatusDetail *ServiceAccountStatusDetail `json:"status_detail,omitempty"`
	SyncStatus   *ServiceAccountSyncStatus   `json:"sync_status,omitempty"`
	// Whether the password for the Okta Universal Directory Account can be rotated using Okta Lifecycle Management
	LcmSyncPossible *bool `json:"lcm_sync_possible,omitempty"`
	// The Okta user ID for the Okta Universal Directory Account
	OktaUserId *string `json:"okta_user_id,omitempty"`
	// A timestamp that indicates when the OPA managed Okta Universal Directory Account was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// A timestamp that indicates when the OPA managed Okta Universal Directory Account was updated
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewOktaUniversalDirectoryAccountWithSettings instantiates a new OktaUniversalDirectoryAccountWithSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOktaUniversalDirectoryAccountWithSettings() *OktaUniversalDirectoryAccountWithSettings {
	this := OktaUniversalDirectoryAccountWithSettings{}
	return &this
}

// NewOktaUniversalDirectoryAccountWithSettingsWithDefaults instantiates a new OktaUniversalDirectoryAccountWithSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOktaUniversalDirectoryAccountWithSettingsWithDefaults() *OktaUniversalDirectoryAccountWithSettings {
	this := OktaUniversalDirectoryAccountWithSettings{}
	return &this
}

// GetAccountSettingsEnabled returns the AccountSettingsEnabled field value if set, zero value otherwise.
func (o *OktaUniversalDirectoryAccountWithSettings) GetAccountSettingsEnabled() []AccountSetting {
	if o == nil || IsNil(o.AccountSettingsEnabled) {
		var ret []AccountSetting
		return ret
	}
	return o.AccountSettingsEnabled
}

// GetAccountSettingsEnabledOk returns a tuple with the AccountSettingsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaUniversalDirectoryAccountWithSettings) GetAccountSettingsEnabledOk() ([]AccountSetting, bool) {
	if o == nil || IsNil(o.AccountSettingsEnabled) {
		return nil, false
	}
	return o.AccountSettingsEnabled, true
}

// HasAccountSettingsEnabled returns a boolean if a field has been set.
func (o *OktaUniversalDirectoryAccountWithSettings) HasAccountSettingsEnabled() bool {
	if o != nil && !IsNil(o.AccountSettingsEnabled) {
		return true
	}

	return false
}

// SetAccountSettingsEnabled gets a reference to the given []AccountSetting and assigns it to the AccountSettingsEnabled field.
func (o *OktaUniversalDirectoryAccountWithSettings) SetAccountSettingsEnabled(v []AccountSetting) *OktaUniversalDirectoryAccountWithSettings {
	o.AccountSettingsEnabled = v
	return o
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OktaUniversalDirectoryAccountWithSettings) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaUniversalDirectoryAccountWithSettings) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OktaUniversalDirectoryAccountWithSettings) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OktaUniversalDirectoryAccountWithSettings) SetId(v string) *OktaUniversalDirectoryAccountWithSettings {
	o.Id = &v
	return o
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OktaUniversalDirectoryAccountWithSettings) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaUniversalDirectoryAccountWithSettings) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OktaUniversalDirectoryAccountWithSettings) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OktaUniversalDirectoryAccountWithSettings) SetName(v string) *OktaUniversalDirectoryAccountWithSettings {
	o.Name = &v
	return o
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *OktaUniversalDirectoryAccountWithSettings) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaUniversalDirectoryAccountWithSettings) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *OktaUniversalDirectoryAccountWithSettings) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *OktaUniversalDirectoryAccountWithSettings) SetUsername(v string) *OktaUniversalDirectoryAccountWithSettings {
	o.Username = &v
	return o
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OktaUniversalDirectoryAccountWithSettings) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaUniversalDirectoryAccountWithSettings) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OktaUniversalDirectoryAccountWithSettings) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OktaUniversalDirectoryAccountWithSettings) SetDescription(v string) *OktaUniversalDirectoryAccountWithSettings {
	o.Description = &v
	return o
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OktaUniversalDirectoryAccountWithSettings) GetStatus() ServiceAccountStatus {
	if o == nil || IsNil(o.Status) {
		var ret ServiceAccountStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaUniversalDirectoryAccountWithSettings) GetStatusOk() (*ServiceAccountStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OktaUniversalDirectoryAccountWithSettings) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ServiceAccountStatus and assigns it to the Status field.
func (o *OktaUniversalDirectoryAccountWithSettings) SetStatus(v ServiceAccountStatus) *OktaUniversalDirectoryAccountWithSettings {
	o.Status = &v
	return o
}

// GetStatusDetail returns the StatusDetail field value if set, zero value otherwise.
func (o *OktaUniversalDirectoryAccountWithSettings) GetStatusDetail() ServiceAccountStatusDetail {
	if o == nil || IsNil(o.StatusDetail) {
		var ret ServiceAccountStatusDetail
		return ret
	}
	return *o.StatusDetail
}

// GetStatusDetailOk returns a tuple with the StatusDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaUniversalDirectoryAccountWithSettings) GetStatusDetailOk() (*ServiceAccountStatusDetail, bool) {
	if o == nil || IsNil(o.StatusDetail) {
		return nil, false
	}
	return o.StatusDetail, true
}

// HasStatusDetail returns a boolean if a field has been set.
func (o *OktaUniversalDirectoryAccountWithSettings) HasStatusDetail() bool {
	if o != nil && !IsNil(o.StatusDetail) {
		return true
	}

	return false
}

// SetStatusDetail gets a reference to the given ServiceAccountStatusDetail and assigns it to the StatusDetail field.
func (o *OktaUniversalDirectoryAccountWithSettings) SetStatusDetail(v ServiceAccountStatusDetail) *OktaUniversalDirectoryAccountWithSettings {
	o.StatusDetail = &v
	return o
}

// GetSyncStatus returns the SyncStatus field value if set, zero value otherwise.
func (o *OktaUniversalDirectoryAccountWithSettings) GetSyncStatus() ServiceAccountSyncStatus {
	if o == nil || IsNil(o.SyncStatus) {
		var ret ServiceAccountSyncStatus
		return ret
	}
	return *o.SyncStatus
}

// GetSyncStatusOk returns a tuple with the SyncStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaUniversalDirectoryAccountWithSettings) GetSyncStatusOk() (*ServiceAccountSyncStatus, bool) {
	if o == nil || IsNil(o.SyncStatus) {
		return nil, false
	}
	return o.SyncStatus, true
}

// HasSyncStatus returns a boolean if a field has been set.
func (o *OktaUniversalDirectoryAccountWithSettings) HasSyncStatus() bool {
	if o != nil && !IsNil(o.SyncStatus) {
		return true
	}

	return false
}

// SetSyncStatus gets a reference to the given ServiceAccountSyncStatus and assigns it to the SyncStatus field.
func (o *OktaUniversalDirectoryAccountWithSettings) SetSyncStatus(v ServiceAccountSyncStatus) *OktaUniversalDirectoryAccountWithSettings {
	o.SyncStatus = &v
	return o
}

// GetLcmSyncPossible returns the LcmSyncPossible field value if set, zero value otherwise.
func (o *OktaUniversalDirectoryAccountWithSettings) GetLcmSyncPossible() bool {
	if o == nil || IsNil(o.LcmSyncPossible) {
		var ret bool
		return ret
	}
	return *o.LcmSyncPossible
}

// GetLcmSyncPossibleOk returns a tuple with the LcmSyncPossible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaUniversalDirectoryAccountWithSettings) GetLcmSyncPossibleOk() (*bool, bool) {
	if o == nil || IsNil(o.LcmSyncPossible) {
		return nil, false
	}
	return o.LcmSyncPossible, true
}

// HasLcmSyncPossible returns a boolean if a field has been set.
func (o *OktaUniversalDirectoryAccountWithSettings) HasLcmSyncPossible() bool {
	if o != nil && !IsNil(o.LcmSyncPossible) {
		return true
	}

	return false
}

// SetLcmSyncPossible gets a reference to the given bool and assigns it to the LcmSyncPossible field.
func (o *OktaUniversalDirectoryAccountWithSettings) SetLcmSyncPossible(v bool) *OktaUniversalDirectoryAccountWithSettings {
	o.LcmSyncPossible = &v
	return o
}

// GetOktaUserId returns the OktaUserId field value if set, zero value otherwise.
func (o *OktaUniversalDirectoryAccountWithSettings) GetOktaUserId() string {
	if o == nil || IsNil(o.OktaUserId) {
		var ret string
		return ret
	}
	return *o.OktaUserId
}

// GetOktaUserIdOk returns a tuple with the OktaUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaUniversalDirectoryAccountWithSettings) GetOktaUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.OktaUserId) {
		return nil, false
	}
	return o.OktaUserId, true
}

// HasOktaUserId returns a boolean if a field has been set.
func (o *OktaUniversalDirectoryAccountWithSettings) HasOktaUserId() bool {
	if o != nil && !IsNil(o.OktaUserId) {
		return true
	}

	return false
}

// SetOktaUserId gets a reference to the given string and assigns it to the OktaUserId field.
func (o *OktaUniversalDirectoryAccountWithSettings) SetOktaUserId(v string) *OktaUniversalDirectoryAccountWithSettings {
	o.OktaUserId = &v
	return o
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *OktaUniversalDirectoryAccountWithSettings) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaUniversalDirectoryAccountWithSettings) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *OktaUniversalDirectoryAccountWithSettings) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *OktaUniversalDirectoryAccountWithSettings) SetCreatedAt(v time.Time) *OktaUniversalDirectoryAccountWithSettings {
	o.CreatedAt = &v
	return o
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *OktaUniversalDirectoryAccountWithSettings) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaUniversalDirectoryAccountWithSettings) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *OktaUniversalDirectoryAccountWithSettings) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *OktaUniversalDirectoryAccountWithSettings) SetUpdatedAt(v time.Time) *OktaUniversalDirectoryAccountWithSettings {
	o.UpdatedAt = &v
	return o
}

func (o OktaUniversalDirectoryAccountWithSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OktaUniversalDirectoryAccountWithSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountSettingsEnabled) {
		toSerialize["account_settings_enabled"] = o.AccountSettingsEnabled
	}
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
	if !IsNil(o.OktaUserId) {
		toSerialize["okta_user_id"] = o.OktaUserId
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableOktaUniversalDirectoryAccountWithSettings struct {
	value *OktaUniversalDirectoryAccountWithSettings
	isSet bool
}

func (v NullableOktaUniversalDirectoryAccountWithSettings) Get() *OktaUniversalDirectoryAccountWithSettings {
	return v.value
}

func (v *NullableOktaUniversalDirectoryAccountWithSettings) Set(val *OktaUniversalDirectoryAccountWithSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableOktaUniversalDirectoryAccountWithSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableOktaUniversalDirectoryAccountWithSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOktaUniversalDirectoryAccountWithSettings(val *OktaUniversalDirectoryAccountWithSettings) *NullableOktaUniversalDirectoryAccountWithSettings {
	return &NullableOktaUniversalDirectoryAccountWithSettings{value: val, isSet: true}
}

func (v NullableOktaUniversalDirectoryAccountWithSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOktaUniversalDirectoryAccountWithSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}