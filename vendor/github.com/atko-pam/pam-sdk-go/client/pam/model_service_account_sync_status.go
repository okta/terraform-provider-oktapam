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

// ServiceAccountSyncStatus Describes the current sync status for actively managed accounts
type ServiceAccountSyncStatus string

// List of ServiceAccountSyncStatus
const (
	ServiceAccountSyncStatus_NOT_SYNCED  ServiceAccountSyncStatus = "NOT_SYNCED"
	ServiceAccountSyncStatus_SYNCED      ServiceAccountSyncStatus = "SYNCED"
	ServiceAccountSyncStatus_SYNCING     ServiceAccountSyncStatus = "SYNCING"
	ServiceAccountSyncStatus_SYNC_FAILED ServiceAccountSyncStatus = "SYNC_FAILED"
)

// All allowed values of ServiceAccountSyncStatus enum
var AllowedServiceAccountSyncStatusEnumValues = []ServiceAccountSyncStatus{
	"NOT_SYNCED",
	"SYNCED",
	"SYNCING",
	"SYNC_FAILED",
}

func (v *ServiceAccountSyncStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ServiceAccountSyncStatus(value)

	*v = enumTypeValue

	return nil
}

// NewServiceAccountSyncStatusFromValue returns a pointer to a valid ServiceAccountSyncStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewServiceAccountSyncStatusFromValue(v string) (*ServiceAccountSyncStatus, error) {
	ev := ServiceAccountSyncStatus(v)

	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ServiceAccountSyncStatus) IsValid() bool {
	for _, existing := range AllowedServiceAccountSyncStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceAccountSyncStatus value
func (v ServiceAccountSyncStatus) Ptr() *ServiceAccountSyncStatus {
	return &v
}

type NullableServiceAccountSyncStatus struct {
	value *ServiceAccountSyncStatus
	isSet bool
}

func (v NullableServiceAccountSyncStatus) Get() *ServiceAccountSyncStatus {
	return v.value
}

func (v *NullableServiceAccountSyncStatus) Set(val *ServiceAccountSyncStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccountSyncStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccountSyncStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccountSyncStatus(val *ServiceAccountSyncStatus) *NullableServiceAccountSyncStatus {
	return &NullableServiceAccountSyncStatus{value: val, isSet: true}
}

func (v NullableServiceAccountSyncStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccountSyncStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
