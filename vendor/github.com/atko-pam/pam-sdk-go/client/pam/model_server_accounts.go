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

// checks if the ServerAccounts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerAccounts{}

// ServerAccounts struct for ServerAccounts
type ServerAccounts struct {
	// The UUID of the Server Account
	Id *string `json:"id,omitempty"`
	// The UUID of the Team associated with this Server Account
	TeamId *string `json:"team_id,omitempty"`
	// The UUID of the Project associated with this Server Account
	ProjectId *string `json:"project_id,omitempty"`
	// The UUID of the Server associated with this Server Account
	ServerId *string `json:"server_id,omitempty"`
	// The user account name on the Server
	Login *string `json:"login,omitempty"`
	// The hostname of the Server
	Hostname *string `json:"hostname,omitempty"`
	// A timestamp indicating when the Server Account was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// A timestamp indicating when the Server Account was deleted
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// A timestamp indicating when the Server last reported a change to the Server Account password
	LastPasswordChangeSuccessReportTimestamp *time.Time `json:"last_password_change_success_report_timestamp,omitempty"`
	// A timestamp reported from the Server system clock indicating when the Server last changed the Server Account password
	LastPasswordChangeSystemTimestamp *time.Time `json:"last_password_change_system_timestamp,omitempty"`
	// A timestamp indicating when the Server last reported a failure to change to the Server Account password
	LastPasswordChangeErrorReportTimestamp *time.Time `json:"last_password_change_error_report_timestamp,omitempty"`
	// A timestamp reported from the Server system clock indicating when the Server last failed to change the Server Account password
	LastPasswordChangeErrorSystemTimestamp *time.Time `json:"last_password_change_error_system_timestamp,omitempty"`
	// The type of error message reported during the most recent failure to change the Server Account password
	LastPasswordChangeErrorType *string `json:"last_password_change_error_type,omitempty"`
	// The error message metadata reported during the most recent failure to change the Server Account password
	LastPasswordChangeErrorMetadata *string `json:"last_password_change_error_metadata,omitempty"`
	// If `true`, indicates that the Server Account password is managed by OPA. If `false`, this account was discovered on the Server, but the password is not managed by OPA.
	Managed *bool `json:"managed,omitempty"`
}

// NewServerAccounts instantiates a new ServerAccounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerAccounts() *ServerAccounts {
	this := ServerAccounts{}
	return &this
}

// NewServerAccountsWithDefaults instantiates a new ServerAccounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerAccountsWithDefaults() *ServerAccounts {
	this := ServerAccounts{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServerAccounts) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerAccounts) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServerAccounts) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServerAccounts) SetId(v string) *ServerAccounts {
	o.Id = &v
	return o
}

// GetTeamId returns the TeamId field value if set, zero value otherwise.
func (o *ServerAccounts) GetTeamId() string {
	if o == nil || IsNil(o.TeamId) {
		var ret string
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerAccounts) GetTeamIdOk() (*string, bool) {
	if o == nil || IsNil(o.TeamId) {
		return nil, false
	}
	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *ServerAccounts) HasTeamId() bool {
	if o != nil && !IsNil(o.TeamId) {
		return true
	}

	return false
}

// SetTeamId gets a reference to the given string and assigns it to the TeamId field.
func (o *ServerAccounts) SetTeamId(v string) *ServerAccounts {
	o.TeamId = &v
	return o
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *ServerAccounts) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerAccounts) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *ServerAccounts) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *ServerAccounts) SetProjectId(v string) *ServerAccounts {
	o.ProjectId = &v
	return o
}

// GetServerId returns the ServerId field value if set, zero value otherwise.
func (o *ServerAccounts) GetServerId() string {
	if o == nil || IsNil(o.ServerId) {
		var ret string
		return ret
	}
	return *o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerAccounts) GetServerIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServerId) {
		return nil, false
	}
	return o.ServerId, true
}

// HasServerId returns a boolean if a field has been set.
func (o *ServerAccounts) HasServerId() bool {
	if o != nil && !IsNil(o.ServerId) {
		return true
	}

	return false
}

// SetServerId gets a reference to the given string and assigns it to the ServerId field.
func (o *ServerAccounts) SetServerId(v string) *ServerAccounts {
	o.ServerId = &v
	return o
}

// GetLogin returns the Login field value if set, zero value otherwise.
func (o *ServerAccounts) GetLogin() string {
	if o == nil || IsNil(o.Login) {
		var ret string
		return ret
	}
	return *o.Login
}

// GetLoginOk returns a tuple with the Login field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerAccounts) GetLoginOk() (*string, bool) {
	if o == nil || IsNil(o.Login) {
		return nil, false
	}
	return o.Login, true
}

// HasLogin returns a boolean if a field has been set.
func (o *ServerAccounts) HasLogin() bool {
	if o != nil && !IsNil(o.Login) {
		return true
	}

	return false
}

// SetLogin gets a reference to the given string and assigns it to the Login field.
func (o *ServerAccounts) SetLogin(v string) *ServerAccounts {
	o.Login = &v
	return o
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *ServerAccounts) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerAccounts) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *ServerAccounts) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *ServerAccounts) SetHostname(v string) *ServerAccounts {
	o.Hostname = &v
	return o
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ServerAccounts) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerAccounts) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ServerAccounts) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ServerAccounts) SetCreatedAt(v time.Time) *ServerAccounts {
	o.CreatedAt = &v
	return o
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *ServerAccounts) GetDeletedAt() time.Time {
	if o == nil || IsNil(o.DeletedAt) {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerAccounts) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *ServerAccounts) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given time.Time and assigns it to the DeletedAt field.
func (o *ServerAccounts) SetDeletedAt(v time.Time) *ServerAccounts {
	o.DeletedAt = &v
	return o
}

// GetLastPasswordChangeSuccessReportTimestamp returns the LastPasswordChangeSuccessReportTimestamp field value if set, zero value otherwise.
func (o *ServerAccounts) GetLastPasswordChangeSuccessReportTimestamp() time.Time {
	if o == nil || IsNil(o.LastPasswordChangeSuccessReportTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.LastPasswordChangeSuccessReportTimestamp
}

// GetLastPasswordChangeSuccessReportTimestampOk returns a tuple with the LastPasswordChangeSuccessReportTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerAccounts) GetLastPasswordChangeSuccessReportTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastPasswordChangeSuccessReportTimestamp) {
		return nil, false
	}
	return o.LastPasswordChangeSuccessReportTimestamp, true
}

// HasLastPasswordChangeSuccessReportTimestamp returns a boolean if a field has been set.
func (o *ServerAccounts) HasLastPasswordChangeSuccessReportTimestamp() bool {
	if o != nil && !IsNil(o.LastPasswordChangeSuccessReportTimestamp) {
		return true
	}

	return false
}

// SetLastPasswordChangeSuccessReportTimestamp gets a reference to the given time.Time and assigns it to the LastPasswordChangeSuccessReportTimestamp field.
func (o *ServerAccounts) SetLastPasswordChangeSuccessReportTimestamp(v time.Time) *ServerAccounts {
	o.LastPasswordChangeSuccessReportTimestamp = &v
	return o
}

// GetLastPasswordChangeSystemTimestamp returns the LastPasswordChangeSystemTimestamp field value if set, zero value otherwise.
func (o *ServerAccounts) GetLastPasswordChangeSystemTimestamp() time.Time {
	if o == nil || IsNil(o.LastPasswordChangeSystemTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.LastPasswordChangeSystemTimestamp
}

// GetLastPasswordChangeSystemTimestampOk returns a tuple with the LastPasswordChangeSystemTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerAccounts) GetLastPasswordChangeSystemTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastPasswordChangeSystemTimestamp) {
		return nil, false
	}
	return o.LastPasswordChangeSystemTimestamp, true
}

// HasLastPasswordChangeSystemTimestamp returns a boolean if a field has been set.
func (o *ServerAccounts) HasLastPasswordChangeSystemTimestamp() bool {
	if o != nil && !IsNil(o.LastPasswordChangeSystemTimestamp) {
		return true
	}

	return false
}

// SetLastPasswordChangeSystemTimestamp gets a reference to the given time.Time and assigns it to the LastPasswordChangeSystemTimestamp field.
func (o *ServerAccounts) SetLastPasswordChangeSystemTimestamp(v time.Time) *ServerAccounts {
	o.LastPasswordChangeSystemTimestamp = &v
	return o
}

// GetLastPasswordChangeErrorReportTimestamp returns the LastPasswordChangeErrorReportTimestamp field value if set, zero value otherwise.
func (o *ServerAccounts) GetLastPasswordChangeErrorReportTimestamp() time.Time {
	if o == nil || IsNil(o.LastPasswordChangeErrorReportTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.LastPasswordChangeErrorReportTimestamp
}

// GetLastPasswordChangeErrorReportTimestampOk returns a tuple with the LastPasswordChangeErrorReportTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerAccounts) GetLastPasswordChangeErrorReportTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastPasswordChangeErrorReportTimestamp) {
		return nil, false
	}
	return o.LastPasswordChangeErrorReportTimestamp, true
}

// HasLastPasswordChangeErrorReportTimestamp returns a boolean if a field has been set.
func (o *ServerAccounts) HasLastPasswordChangeErrorReportTimestamp() bool {
	if o != nil && !IsNil(o.LastPasswordChangeErrorReportTimestamp) {
		return true
	}

	return false
}

// SetLastPasswordChangeErrorReportTimestamp gets a reference to the given time.Time and assigns it to the LastPasswordChangeErrorReportTimestamp field.
func (o *ServerAccounts) SetLastPasswordChangeErrorReportTimestamp(v time.Time) *ServerAccounts {
	o.LastPasswordChangeErrorReportTimestamp = &v
	return o
}

// GetLastPasswordChangeErrorSystemTimestamp returns the LastPasswordChangeErrorSystemTimestamp field value if set, zero value otherwise.
func (o *ServerAccounts) GetLastPasswordChangeErrorSystemTimestamp() time.Time {
	if o == nil || IsNil(o.LastPasswordChangeErrorSystemTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.LastPasswordChangeErrorSystemTimestamp
}

// GetLastPasswordChangeErrorSystemTimestampOk returns a tuple with the LastPasswordChangeErrorSystemTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerAccounts) GetLastPasswordChangeErrorSystemTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastPasswordChangeErrorSystemTimestamp) {
		return nil, false
	}
	return o.LastPasswordChangeErrorSystemTimestamp, true
}

// HasLastPasswordChangeErrorSystemTimestamp returns a boolean if a field has been set.
func (o *ServerAccounts) HasLastPasswordChangeErrorSystemTimestamp() bool {
	if o != nil && !IsNil(o.LastPasswordChangeErrorSystemTimestamp) {
		return true
	}

	return false
}

// SetLastPasswordChangeErrorSystemTimestamp gets a reference to the given time.Time and assigns it to the LastPasswordChangeErrorSystemTimestamp field.
func (o *ServerAccounts) SetLastPasswordChangeErrorSystemTimestamp(v time.Time) *ServerAccounts {
	o.LastPasswordChangeErrorSystemTimestamp = &v
	return o
}

// GetLastPasswordChangeErrorType returns the LastPasswordChangeErrorType field value if set, zero value otherwise.
func (o *ServerAccounts) GetLastPasswordChangeErrorType() string {
	if o == nil || IsNil(o.LastPasswordChangeErrorType) {
		var ret string
		return ret
	}
	return *o.LastPasswordChangeErrorType
}

// GetLastPasswordChangeErrorTypeOk returns a tuple with the LastPasswordChangeErrorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerAccounts) GetLastPasswordChangeErrorTypeOk() (*string, bool) {
	if o == nil || IsNil(o.LastPasswordChangeErrorType) {
		return nil, false
	}
	return o.LastPasswordChangeErrorType, true
}

// HasLastPasswordChangeErrorType returns a boolean if a field has been set.
func (o *ServerAccounts) HasLastPasswordChangeErrorType() bool {
	if o != nil && !IsNil(o.LastPasswordChangeErrorType) {
		return true
	}

	return false
}

// SetLastPasswordChangeErrorType gets a reference to the given string and assigns it to the LastPasswordChangeErrorType field.
func (o *ServerAccounts) SetLastPasswordChangeErrorType(v string) *ServerAccounts {
	o.LastPasswordChangeErrorType = &v
	return o
}

// GetLastPasswordChangeErrorMetadata returns the LastPasswordChangeErrorMetadata field value if set, zero value otherwise.
func (o *ServerAccounts) GetLastPasswordChangeErrorMetadata() string {
	if o == nil || IsNil(o.LastPasswordChangeErrorMetadata) {
		var ret string
		return ret
	}
	return *o.LastPasswordChangeErrorMetadata
}

// GetLastPasswordChangeErrorMetadataOk returns a tuple with the LastPasswordChangeErrorMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerAccounts) GetLastPasswordChangeErrorMetadataOk() (*string, bool) {
	if o == nil || IsNil(o.LastPasswordChangeErrorMetadata) {
		return nil, false
	}
	return o.LastPasswordChangeErrorMetadata, true
}

// HasLastPasswordChangeErrorMetadata returns a boolean if a field has been set.
func (o *ServerAccounts) HasLastPasswordChangeErrorMetadata() bool {
	if o != nil && !IsNil(o.LastPasswordChangeErrorMetadata) {
		return true
	}

	return false
}

// SetLastPasswordChangeErrorMetadata gets a reference to the given string and assigns it to the LastPasswordChangeErrorMetadata field.
func (o *ServerAccounts) SetLastPasswordChangeErrorMetadata(v string) *ServerAccounts {
	o.LastPasswordChangeErrorMetadata = &v
	return o
}

// GetManaged returns the Managed field value if set, zero value otherwise.
func (o *ServerAccounts) GetManaged() bool {
	if o == nil || IsNil(o.Managed) {
		var ret bool
		return ret
	}
	return *o.Managed
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerAccounts) GetManagedOk() (*bool, bool) {
	if o == nil || IsNil(o.Managed) {
		return nil, false
	}
	return o.Managed, true
}

// HasManaged returns a boolean if a field has been set.
func (o *ServerAccounts) HasManaged() bool {
	if o != nil && !IsNil(o.Managed) {
		return true
	}

	return false
}

// SetManaged gets a reference to the given bool and assigns it to the Managed field.
func (o *ServerAccounts) SetManaged(v bool) *ServerAccounts {
	o.Managed = &v
	return o
}

func (o ServerAccounts) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerAccounts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.TeamId) {
		toSerialize["team_id"] = o.TeamId
	}
	if !IsNil(o.ProjectId) {
		toSerialize["project_id"] = o.ProjectId
	}
	if !IsNil(o.ServerId) {
		toSerialize["server_id"] = o.ServerId
	}
	if !IsNil(o.Login) {
		toSerialize["login"] = o.Login
	}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.DeletedAt) {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	if !IsNil(o.LastPasswordChangeSuccessReportTimestamp) {
		toSerialize["last_password_change_success_report_timestamp"] = o.LastPasswordChangeSuccessReportTimestamp
	}
	if !IsNil(o.LastPasswordChangeSystemTimestamp) {
		toSerialize["last_password_change_system_timestamp"] = o.LastPasswordChangeSystemTimestamp
	}
	if !IsNil(o.LastPasswordChangeErrorReportTimestamp) {
		toSerialize["last_password_change_error_report_timestamp"] = o.LastPasswordChangeErrorReportTimestamp
	}
	if !IsNil(o.LastPasswordChangeErrorSystemTimestamp) {
		toSerialize["last_password_change_error_system_timestamp"] = o.LastPasswordChangeErrorSystemTimestamp
	}
	if !IsNil(o.LastPasswordChangeErrorType) {
		toSerialize["last_password_change_error_type"] = o.LastPasswordChangeErrorType
	}
	if !IsNil(o.LastPasswordChangeErrorMetadata) {
		toSerialize["last_password_change_error_metadata"] = o.LastPasswordChangeErrorMetadata
	}
	if !IsNil(o.Managed) {
		toSerialize["managed"] = o.Managed
	}
	return toSerialize, nil
}

type NullableServerAccounts struct {
	value *ServerAccounts
	isSet bool
}

func (v NullableServerAccounts) Get() *ServerAccounts {
	return v.value
}

func (v *NullableServerAccounts) Set(val *ServerAccounts) {
	v.value = val
	v.isSet = true
}

func (v NullableServerAccounts) IsSet() bool {
	return v.isSet
}

func (v *NullableServerAccounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerAccounts(val *ServerAccounts) *NullableServerAccounts {
	return &NullableServerAccounts{value: val, isSet: true}
}

func (v NullableServerAccounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerAccounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
