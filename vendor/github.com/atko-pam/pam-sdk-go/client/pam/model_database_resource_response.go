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

// checks if the DatabaseResourceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatabaseResourceResponse{}

// DatabaseResourceResponse struct for DatabaseResourceResponse
type DatabaseResourceResponse struct {
	// The ID of the Database Resource
	Id string `json:"id"`
	// The canonical name of the Database Resource
	CanonicalName string       `json:"canonical_name"`
	DatabaseType  DatabaseType `json:"database_type"`
	// The ID of the recipe book which contains the SQL queries used by the Gateway. The recipe book permits customers to override the default SQL used by the system, allowing for customized query execution.
	RecipeBookId                    *string                         `json:"recipe_book_id,omitempty"`
	ManagementConnectionDetailsType ManagementConnectionDetailsType `json:"management_connection_details_type"`
	ManagementConnectionDetails     ManagementConnectionDetails     `json:"management_connection_details"`
	// A selector is composed of key-value pairs and is used to dynamically allocate tasks to Gateways for Database Resources. Only those Gateways that match the criteria defined by the selector are eligible to claim and execute the work.
	ManagementGatewaySelector *map[string]string `json:"management_gateway_selector,omitempty"`
	// The ID of the selector used to identify the Gateway associated with the Database Resource
	ManagementGatewaySelectorId string `json:"management_gateway_selector_id"`
	// A timestamp that indicates when the Database Resource was created
	CreatedAt time.Time `json:"created_at"`
	// A timestamp that indicates when the Database Resource was updated
	UpdatedAt time.Time `json:"updated_at"`
}

// NewDatabaseResourceResponse instantiates a new DatabaseResourceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseResourceResponse(id string, canonicalName string, databaseType DatabaseType, managementConnectionDetailsType ManagementConnectionDetailsType, managementConnectionDetails ManagementConnectionDetails, managementGatewaySelectorId string, createdAt time.Time, updatedAt time.Time) *DatabaseResourceResponse {
	this := DatabaseResourceResponse{}
	this.Id = id
	this.CanonicalName = canonicalName
	this.DatabaseType = databaseType
	this.ManagementConnectionDetailsType = managementConnectionDetailsType
	this.ManagementConnectionDetails = managementConnectionDetails
	this.ManagementGatewaySelectorId = managementGatewaySelectorId
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewDatabaseResourceResponseWithDefaults instantiates a new DatabaseResourceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseResourceResponseWithDefaults() *DatabaseResourceResponse {
	this := DatabaseResourceResponse{}
	return &this
}

// GetId returns the Id field value
func (o *DatabaseResourceResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DatabaseResourceResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DatabaseResourceResponse) SetId(v string) *DatabaseResourceResponse {
	o.Id = v
	return o
}

// GetCanonicalName returns the CanonicalName field value
func (o *DatabaseResourceResponse) GetCanonicalName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CanonicalName
}

// GetCanonicalNameOk returns a tuple with the CanonicalName field value
// and a boolean to check if the value has been set.
func (o *DatabaseResourceResponse) GetCanonicalNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanonicalName, true
}

// SetCanonicalName sets field value
func (o *DatabaseResourceResponse) SetCanonicalName(v string) *DatabaseResourceResponse {
	o.CanonicalName = v
	return o
}

// GetDatabaseType returns the DatabaseType field value
func (o *DatabaseResourceResponse) GetDatabaseType() DatabaseType {
	if o == nil {
		var ret DatabaseType
		return ret
	}

	return o.DatabaseType
}

// GetDatabaseTypeOk returns a tuple with the DatabaseType field value
// and a boolean to check if the value has been set.
func (o *DatabaseResourceResponse) GetDatabaseTypeOk() (*DatabaseType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatabaseType, true
}

// SetDatabaseType sets field value
func (o *DatabaseResourceResponse) SetDatabaseType(v DatabaseType) *DatabaseResourceResponse {
	o.DatabaseType = v
	return o
}

// GetRecipeBookId returns the RecipeBookId field value if set, zero value otherwise.
func (o *DatabaseResourceResponse) GetRecipeBookId() string {
	if o == nil || IsNil(o.RecipeBookId) {
		var ret string
		return ret
	}
	return *o.RecipeBookId
}

// GetRecipeBookIdOk returns a tuple with the RecipeBookId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseResourceResponse) GetRecipeBookIdOk() (*string, bool) {
	if o == nil || IsNil(o.RecipeBookId) {
		return nil, false
	}
	return o.RecipeBookId, true
}

// HasRecipeBookId returns a boolean if a field has been set.
func (o *DatabaseResourceResponse) HasRecipeBookId() bool {
	if o != nil && !IsNil(o.RecipeBookId) {
		return true
	}

	return false
}

// SetRecipeBookId gets a reference to the given string and assigns it to the RecipeBookId field.
func (o *DatabaseResourceResponse) SetRecipeBookId(v string) *DatabaseResourceResponse {
	o.RecipeBookId = &v
	return o
}

// GetManagementConnectionDetailsType returns the ManagementConnectionDetailsType field value
func (o *DatabaseResourceResponse) GetManagementConnectionDetailsType() ManagementConnectionDetailsType {
	if o == nil {
		var ret ManagementConnectionDetailsType
		return ret
	}

	return o.ManagementConnectionDetailsType
}

// GetManagementConnectionDetailsTypeOk returns a tuple with the ManagementConnectionDetailsType field value
// and a boolean to check if the value has been set.
func (o *DatabaseResourceResponse) GetManagementConnectionDetailsTypeOk() (*ManagementConnectionDetailsType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ManagementConnectionDetailsType, true
}

// SetManagementConnectionDetailsType sets field value
func (o *DatabaseResourceResponse) SetManagementConnectionDetailsType(v ManagementConnectionDetailsType) *DatabaseResourceResponse {
	o.ManagementConnectionDetailsType = v
	return o
}

// GetManagementConnectionDetails returns the ManagementConnectionDetails field value
func (o *DatabaseResourceResponse) GetManagementConnectionDetails() ManagementConnectionDetails {
	if o == nil {
		var ret ManagementConnectionDetails
		return ret
	}

	return o.ManagementConnectionDetails
}

// GetManagementConnectionDetailsOk returns a tuple with the ManagementConnectionDetails field value
// and a boolean to check if the value has been set.
func (o *DatabaseResourceResponse) GetManagementConnectionDetailsOk() (*ManagementConnectionDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ManagementConnectionDetails, true
}

// SetManagementConnectionDetails sets field value
func (o *DatabaseResourceResponse) SetManagementConnectionDetails(v ManagementConnectionDetails) *DatabaseResourceResponse {
	o.ManagementConnectionDetails = v
	return o
}

// GetManagementGatewaySelector returns the ManagementGatewaySelector field value if set, zero value otherwise.
func (o *DatabaseResourceResponse) GetManagementGatewaySelector() map[string]string {
	if o == nil || IsNil(o.ManagementGatewaySelector) {
		var ret map[string]string
		return ret
	}
	return *o.ManagementGatewaySelector
}

// GetManagementGatewaySelectorOk returns a tuple with the ManagementGatewaySelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseResourceResponse) GetManagementGatewaySelectorOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ManagementGatewaySelector) {
		return nil, false
	}
	return o.ManagementGatewaySelector, true
}

// HasManagementGatewaySelector returns a boolean if a field has been set.
func (o *DatabaseResourceResponse) HasManagementGatewaySelector() bool {
	if o != nil && !IsNil(o.ManagementGatewaySelector) {
		return true
	}

	return false
}

// SetManagementGatewaySelector gets a reference to the given map[string]string and assigns it to the ManagementGatewaySelector field.
func (o *DatabaseResourceResponse) SetManagementGatewaySelector(v map[string]string) *DatabaseResourceResponse {
	o.ManagementGatewaySelector = &v
	return o
}

// GetManagementGatewaySelectorId returns the ManagementGatewaySelectorId field value
func (o *DatabaseResourceResponse) GetManagementGatewaySelectorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ManagementGatewaySelectorId
}

// GetManagementGatewaySelectorIdOk returns a tuple with the ManagementGatewaySelectorId field value
// and a boolean to check if the value has been set.
func (o *DatabaseResourceResponse) GetManagementGatewaySelectorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ManagementGatewaySelectorId, true
}

// SetManagementGatewaySelectorId sets field value
func (o *DatabaseResourceResponse) SetManagementGatewaySelectorId(v string) *DatabaseResourceResponse {
	o.ManagementGatewaySelectorId = v
	return o
}

// GetCreatedAt returns the CreatedAt field value
func (o *DatabaseResourceResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DatabaseResourceResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *DatabaseResourceResponse) SetCreatedAt(v time.Time) *DatabaseResourceResponse {
	o.CreatedAt = v
	return o
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *DatabaseResourceResponse) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *DatabaseResourceResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *DatabaseResourceResponse) SetUpdatedAt(v time.Time) *DatabaseResourceResponse {
	o.UpdatedAt = v
	return o
}

func (o DatabaseResourceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatabaseResourceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["canonical_name"] = o.CanonicalName
	toSerialize["database_type"] = o.DatabaseType
	if !IsNil(o.RecipeBookId) {
		toSerialize["recipe_book_id"] = o.RecipeBookId
	}
	toSerialize["management_connection_details_type"] = o.ManagementConnectionDetailsType
	toSerialize["management_connection_details"] = o.ManagementConnectionDetails
	if !IsNil(o.ManagementGatewaySelector) {
		toSerialize["management_gateway_selector"] = o.ManagementGatewaySelector
	}
	toSerialize["management_gateway_selector_id"] = o.ManagementGatewaySelectorId
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

type NullableDatabaseResourceResponse struct {
	value *DatabaseResourceResponse
	isSet bool
}

func (v NullableDatabaseResourceResponse) Get() *DatabaseResourceResponse {
	return v.value
}

func (v *NullableDatabaseResourceResponse) Set(val *DatabaseResourceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseResourceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseResourceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseResourceResponse(val *DatabaseResourceResponse) *NullableDatabaseResourceResponse {
	return &NullableDatabaseResourceResponse{value: val, isSet: true}
}

func (v NullableDatabaseResourceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseResourceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}