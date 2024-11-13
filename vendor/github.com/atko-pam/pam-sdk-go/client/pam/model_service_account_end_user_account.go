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

// checks if the ServiceAccountEndUserAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceAccountEndUserAccount{}

// ServiceAccountEndUserAccount Service account details
type ServiceAccountEndUserAccount struct {
	// The UUID of the service account
	Id *string `json:"id,omitempty"`
	// The human-readable name for the service account
	Name *string `json:"name,omitempty"`
	// The username associated with the service account
	Username *string `json:"username,omitempty"`
	// Whether the password for the service account can be rotated using Okta Lifecycle Management
	LcmSyncPossible *bool `json:"lcm_sync_possible,omitempty"`
	// Current availability status of the account
	AvailabilityStatus  *string                     `json:"availability_status,omitempty"`
	AccountStatus       *ServiceAccountStatus       `json:"account_status,omitempty"`
	AccountStatusDetail *ServiceAccountStatusDetail `json:"account_status_detail,omitempty"`
}

// NewServiceAccountEndUserAccount instantiates a new ServiceAccountEndUserAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountEndUserAccount() *ServiceAccountEndUserAccount {
	this := ServiceAccountEndUserAccount{}
	return &this
}

// NewServiceAccountEndUserAccountWithDefaults instantiates a new ServiceAccountEndUserAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountEndUserAccountWithDefaults() *ServiceAccountEndUserAccount {
	this := ServiceAccountEndUserAccount{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServiceAccountEndUserAccount) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountEndUserAccount) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServiceAccountEndUserAccount) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServiceAccountEndUserAccount) SetId(v string) *ServiceAccountEndUserAccount {
	o.Id = &v
	return o
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServiceAccountEndUserAccount) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountEndUserAccount) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceAccountEndUserAccount) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceAccountEndUserAccount) SetName(v string) *ServiceAccountEndUserAccount {
	o.Name = &v
	return o
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ServiceAccountEndUserAccount) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountEndUserAccount) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ServiceAccountEndUserAccount) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ServiceAccountEndUserAccount) SetUsername(v string) *ServiceAccountEndUserAccount {
	o.Username = &v
	return o
}

// GetLcmSyncPossible returns the LcmSyncPossible field value if set, zero value otherwise.
func (o *ServiceAccountEndUserAccount) GetLcmSyncPossible() bool {
	if o == nil || IsNil(o.LcmSyncPossible) {
		var ret bool
		return ret
	}
	return *o.LcmSyncPossible
}

// GetLcmSyncPossibleOk returns a tuple with the LcmSyncPossible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountEndUserAccount) GetLcmSyncPossibleOk() (*bool, bool) {
	if o == nil || IsNil(o.LcmSyncPossible) {
		return nil, false
	}
	return o.LcmSyncPossible, true
}

// HasLcmSyncPossible returns a boolean if a field has been set.
func (o *ServiceAccountEndUserAccount) HasLcmSyncPossible() bool {
	if o != nil && !IsNil(o.LcmSyncPossible) {
		return true
	}

	return false
}

// SetLcmSyncPossible gets a reference to the given bool and assigns it to the LcmSyncPossible field.
func (o *ServiceAccountEndUserAccount) SetLcmSyncPossible(v bool) *ServiceAccountEndUserAccount {
	o.LcmSyncPossible = &v
	return o
}

// GetAvailabilityStatus returns the AvailabilityStatus field value if set, zero value otherwise.
func (o *ServiceAccountEndUserAccount) GetAvailabilityStatus() string {
	if o == nil || IsNil(o.AvailabilityStatus) {
		var ret string
		return ret
	}
	return *o.AvailabilityStatus
}

// GetAvailabilityStatusOk returns a tuple with the AvailabilityStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountEndUserAccount) GetAvailabilityStatusOk() (*string, bool) {
	if o == nil || IsNil(o.AvailabilityStatus) {
		return nil, false
	}
	return o.AvailabilityStatus, true
}

// HasAvailabilityStatus returns a boolean if a field has been set.
func (o *ServiceAccountEndUserAccount) HasAvailabilityStatus() bool {
	if o != nil && !IsNil(o.AvailabilityStatus) {
		return true
	}

	return false
}

// SetAvailabilityStatus gets a reference to the given string and assigns it to the AvailabilityStatus field.
func (o *ServiceAccountEndUserAccount) SetAvailabilityStatus(v string) *ServiceAccountEndUserAccount {
	o.AvailabilityStatus = &v
	return o
}

// GetAccountStatus returns the AccountStatus field value if set, zero value otherwise.
func (o *ServiceAccountEndUserAccount) GetAccountStatus() ServiceAccountStatus {
	if o == nil || IsNil(o.AccountStatus) {
		var ret ServiceAccountStatus
		return ret
	}
	return *o.AccountStatus
}

// GetAccountStatusOk returns a tuple with the AccountStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountEndUserAccount) GetAccountStatusOk() (*ServiceAccountStatus, bool) {
	if o == nil || IsNil(o.AccountStatus) {
		return nil, false
	}
	return o.AccountStatus, true
}

// HasAccountStatus returns a boolean if a field has been set.
func (o *ServiceAccountEndUserAccount) HasAccountStatus() bool {
	if o != nil && !IsNil(o.AccountStatus) {
		return true
	}

	return false
}

// SetAccountStatus gets a reference to the given ServiceAccountStatus and assigns it to the AccountStatus field.
func (o *ServiceAccountEndUserAccount) SetAccountStatus(v ServiceAccountStatus) *ServiceAccountEndUserAccount {
	o.AccountStatus = &v
	return o
}

// GetAccountStatusDetail returns the AccountStatusDetail field value if set, zero value otherwise.
func (o *ServiceAccountEndUserAccount) GetAccountStatusDetail() ServiceAccountStatusDetail {
	if o == nil || IsNil(o.AccountStatusDetail) {
		var ret ServiceAccountStatusDetail
		return ret
	}
	return *o.AccountStatusDetail
}

// GetAccountStatusDetailOk returns a tuple with the AccountStatusDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountEndUserAccount) GetAccountStatusDetailOk() (*ServiceAccountStatusDetail, bool) {
	if o == nil || IsNil(o.AccountStatusDetail) {
		return nil, false
	}
	return o.AccountStatusDetail, true
}

// HasAccountStatusDetail returns a boolean if a field has been set.
func (o *ServiceAccountEndUserAccount) HasAccountStatusDetail() bool {
	if o != nil && !IsNil(o.AccountStatusDetail) {
		return true
	}

	return false
}

// SetAccountStatusDetail gets a reference to the given ServiceAccountStatusDetail and assigns it to the AccountStatusDetail field.
func (o *ServiceAccountEndUserAccount) SetAccountStatusDetail(v ServiceAccountStatusDetail) *ServiceAccountEndUserAccount {
	o.AccountStatusDetail = &v
	return o
}

func (o ServiceAccountEndUserAccount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceAccountEndUserAccount) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.LcmSyncPossible) {
		toSerialize["lcm_sync_possible"] = o.LcmSyncPossible
	}
	if !IsNil(o.AvailabilityStatus) {
		toSerialize["availability_status"] = o.AvailabilityStatus
	}
	if !IsNil(o.AccountStatus) {
		toSerialize["account_status"] = o.AccountStatus
	}
	if !IsNil(o.AccountStatusDetail) {
		toSerialize["account_status_detail"] = o.AccountStatusDetail
	}
	return toSerialize, nil
}

type NullableServiceAccountEndUserAccount struct {
	value *ServiceAccountEndUserAccount
	isSet bool
}

func (v NullableServiceAccountEndUserAccount) Get() *ServiceAccountEndUserAccount {
	return v.value
}

func (v *NullableServiceAccountEndUserAccount) Set(val *ServiceAccountEndUserAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccountEndUserAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccountEndUserAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccountEndUserAccount(val *ServiceAccountEndUserAccount) *NullableServiceAccountEndUserAccount {
	return &NullableServiceAccountEndUserAccount{value: val, isSet: true}
}

func (v NullableServiceAccountEndUserAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccountEndUserAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}