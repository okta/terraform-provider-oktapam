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

// checks if the DatabaseRoleDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatabaseRoleDetails{}

// DatabaseRoleDetails A database role that maps an identifier to one or more accounts in the database
type DatabaseRoleDetails struct {
	// Type of this role
	Type string `json:"_type"`
	// Name of this database role
	Name string `json:"name"`
	// List of accounts associated with this role
	Accounts []string `json:"accounts,omitempty"`
}

// NewDatabaseRoleDetails instantiates a new DatabaseRoleDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseRoleDetails(type_ string, name string) *DatabaseRoleDetails {
	this := DatabaseRoleDetails{}
	this.Type = type_
	this.Name = name
	return &this
}

// NewDatabaseRoleDetailsWithDefaults instantiates a new DatabaseRoleDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseRoleDetailsWithDefaults() *DatabaseRoleDetails {
	this := DatabaseRoleDetails{}
	return &this
}

// GetType returns the Type field value
func (o *DatabaseRoleDetails) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DatabaseRoleDetails) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DatabaseRoleDetails) SetType(v string) *DatabaseRoleDetails {
	o.Type = v
	return o
}

// GetName returns the Name field value
func (o *DatabaseRoleDetails) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DatabaseRoleDetails) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DatabaseRoleDetails) SetName(v string) *DatabaseRoleDetails {
	o.Name = v
	return o
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *DatabaseRoleDetails) GetAccounts() []string {
	if o == nil || IsNil(o.Accounts) {
		var ret []string
		return ret
	}
	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseRoleDetails) GetAccountsOk() ([]string, bool) {
	if o == nil || IsNil(o.Accounts) {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *DatabaseRoleDetails) HasAccounts() bool {
	if o != nil && !IsNil(o.Accounts) {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []string and assigns it to the Accounts field.
func (o *DatabaseRoleDetails) SetAccounts(v []string) *DatabaseRoleDetails {
	o.Accounts = v
	return o
}

func (o DatabaseRoleDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatabaseRoleDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_type"] = o.Type
	toSerialize["name"] = o.Name
	if !IsNil(o.Accounts) {
		toSerialize["accounts"] = o.Accounts
	}
	return toSerialize, nil
}

type NullableDatabaseRoleDetails struct {
	value *DatabaseRoleDetails
	isSet bool
}

func (v NullableDatabaseRoleDetails) Get() *DatabaseRoleDetails {
	return v.value
}

func (v *NullableDatabaseRoleDetails) Set(val *DatabaseRoleDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseRoleDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseRoleDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseRoleDetails(val *DatabaseRoleDetails) *NullableDatabaseRoleDetails {
	return &NullableDatabaseRoleDetails{value: val, isSet: true}
}

func (v NullableDatabaseRoleDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseRoleDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
