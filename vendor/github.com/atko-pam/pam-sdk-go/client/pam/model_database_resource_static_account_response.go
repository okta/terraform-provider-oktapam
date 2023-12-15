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

// checks if the DatabaseResourceStaticAccountResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatabaseResourceStaticAccountResponse{}

// DatabaseResourceStaticAccountResponse struct for DatabaseResourceStaticAccountResponse
type DatabaseResourceStaticAccountResponse struct {
	// The ID of the Database Static Account
	Id string `json:"id"`
	// The name of the Database Static Account
	Name                   string                 `json:"name"`
	DatabaseAccountType    DatabaseAccountType    `json:"database_account_type"`
	DatabaseAccountDetails DatabaseAccountDetails `json:"database_account_details"`
	// The secret ID of the Database Static Account
	SecretId string `json:"secret_id"`
	// A timestamp that indicates when the Database Static Account was created
	CreatedAt time.Time `json:"created_at"`
	// A timestamp that indicates when the Database Static Account was updated
	UpdatedAt time.Time `json:"updated_at"`
}

// NewDatabaseResourceStaticAccountResponse instantiates a new DatabaseResourceStaticAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseResourceStaticAccountResponse(id string, name string, databaseAccountType DatabaseAccountType, databaseAccountDetails DatabaseAccountDetails, secretId string, createdAt time.Time, updatedAt time.Time) *DatabaseResourceStaticAccountResponse {
	this := DatabaseResourceStaticAccountResponse{}
	this.Id = id
	this.Name = name
	this.DatabaseAccountType = databaseAccountType
	this.DatabaseAccountDetails = databaseAccountDetails
	this.SecretId = secretId
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewDatabaseResourceStaticAccountResponseWithDefaults instantiates a new DatabaseResourceStaticAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseResourceStaticAccountResponseWithDefaults() *DatabaseResourceStaticAccountResponse {
	this := DatabaseResourceStaticAccountResponse{}
	return &this
}

// GetId returns the Id field value
func (o *DatabaseResourceStaticAccountResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DatabaseResourceStaticAccountResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DatabaseResourceStaticAccountResponse) SetId(v string) *DatabaseResourceStaticAccountResponse {
	o.Id = v
	return o
}

// GetName returns the Name field value
func (o *DatabaseResourceStaticAccountResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DatabaseResourceStaticAccountResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DatabaseResourceStaticAccountResponse) SetName(v string) *DatabaseResourceStaticAccountResponse {
	o.Name = v
	return o
}

// GetDatabaseAccountType returns the DatabaseAccountType field value
func (o *DatabaseResourceStaticAccountResponse) GetDatabaseAccountType() DatabaseAccountType {
	if o == nil {
		var ret DatabaseAccountType
		return ret
	}

	return o.DatabaseAccountType
}

// GetDatabaseAccountTypeOk returns a tuple with the DatabaseAccountType field value
// and a boolean to check if the value has been set.
func (o *DatabaseResourceStaticAccountResponse) GetDatabaseAccountTypeOk() (*DatabaseAccountType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatabaseAccountType, true
}

// SetDatabaseAccountType sets field value
func (o *DatabaseResourceStaticAccountResponse) SetDatabaseAccountType(v DatabaseAccountType) *DatabaseResourceStaticAccountResponse {
	o.DatabaseAccountType = v
	return o
}

// GetDatabaseAccountDetails returns the DatabaseAccountDetails field value
func (o *DatabaseResourceStaticAccountResponse) GetDatabaseAccountDetails() DatabaseAccountDetails {
	if o == nil {
		var ret DatabaseAccountDetails
		return ret
	}

	return o.DatabaseAccountDetails
}

// GetDatabaseAccountDetailsOk returns a tuple with the DatabaseAccountDetails field value
// and a boolean to check if the value has been set.
func (o *DatabaseResourceStaticAccountResponse) GetDatabaseAccountDetailsOk() (*DatabaseAccountDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatabaseAccountDetails, true
}

// SetDatabaseAccountDetails sets field value
func (o *DatabaseResourceStaticAccountResponse) SetDatabaseAccountDetails(v DatabaseAccountDetails) *DatabaseResourceStaticAccountResponse {
	o.DatabaseAccountDetails = v
	return o
}

// GetSecretId returns the SecretId field value
func (o *DatabaseResourceStaticAccountResponse) GetSecretId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretId
}

// GetSecretIdOk returns a tuple with the SecretId field value
// and a boolean to check if the value has been set.
func (o *DatabaseResourceStaticAccountResponse) GetSecretIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretId, true
}

// SetSecretId sets field value
func (o *DatabaseResourceStaticAccountResponse) SetSecretId(v string) *DatabaseResourceStaticAccountResponse {
	o.SecretId = v
	return o
}

// GetCreatedAt returns the CreatedAt field value
func (o *DatabaseResourceStaticAccountResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DatabaseResourceStaticAccountResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *DatabaseResourceStaticAccountResponse) SetCreatedAt(v time.Time) *DatabaseResourceStaticAccountResponse {
	o.CreatedAt = v
	return o
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *DatabaseResourceStaticAccountResponse) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *DatabaseResourceStaticAccountResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *DatabaseResourceStaticAccountResponse) SetUpdatedAt(v time.Time) *DatabaseResourceStaticAccountResponse {
	o.UpdatedAt = v
	return o
}

func (o DatabaseResourceStaticAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatabaseResourceStaticAccountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["database_account_type"] = o.DatabaseAccountType
	toSerialize["database_account_details"] = o.DatabaseAccountDetails
	toSerialize["secret_id"] = o.SecretId
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

type NullableDatabaseResourceStaticAccountResponse struct {
	value *DatabaseResourceStaticAccountResponse
	isSet bool
}

func (v NullableDatabaseResourceStaticAccountResponse) Get() *DatabaseResourceStaticAccountResponse {
	return v.value
}

func (v *NullableDatabaseResourceStaticAccountResponse) Set(val *DatabaseResourceStaticAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseResourceStaticAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseResourceStaticAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseResourceStaticAccountResponse(val *DatabaseResourceStaticAccountResponse) *NullableDatabaseResourceStaticAccountResponse {
	return &NullableDatabaseResourceStaticAccountResponse{value: val, isSet: true}
}

func (v NullableDatabaseResourceStaticAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseResourceStaticAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}