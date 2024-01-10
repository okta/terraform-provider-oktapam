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

// checks if the Project type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Project{}

// Project struct for Project
type Project struct {
	// Whether to create Server Users for Users in this Project. Defaults to `false`. If `false`, you must ensure that accounts exist on the Server for each User.
	CreateServerUsers NullableBool `json:"create_server_users,omitempty"`
	// A timestamp indicating when the Project was deleted
	DeletedAt NullableTime `json:"deleted_at,omitempty"`
	// Whether to force the project to use a shared SSH account
	ForceSharedSshUsers *bool `json:"force_shared_ssh_users,omitempty"`
	// A comma separated list of labels used to match to enrolled Gateways. Labels should use the following format: `key=value`
	GatewaySelector *string `json:"gateway_selector,omitempty"`
	// The UUID of the Project
	Id string `json:"id"`
	// The UUID of the Resource Group where the Project is located
	ResourceGroupId *string `json:"resource_group_id,omitempty"`
	// The number of active resources within this project
	ActiveResourceCount *int32 `json:"active_resource_count,omitempty"`
	// The number of stale resources within this project
	StaleResourceCount *int32 `json:"stale_resource_count,omitempty"`
	// The name of the Project
	Name string `json:"name"`
	// The GID used when creating a Server User
	NextUnixGid NullableInt32 `json:"next_unix_gid,omitempty"`
	// The UID used when creating a Server User
	NextUnixUid NullableInt32 `json:"next_unix_uid,omitempty"`
	// If `true`, the Project requires preauthorization before a User can access a Server. Default is `false`.
	RequirePreauthForCreds NullableBool `json:"require_preauth_for_creds,omitempty"`
	// If `true`, creates persistent user accounts and home folders on Servers in this Project for every user on your Team. By default, on-demand accounts are only created when a user accesses a Server.
	PersistentServerUserAccounts NullableBool `json:"persistent_server_user_accounts,omitempty"`
	// Whether to manage existing local accounts on the server
	ServerAccountManagement *bool `json:"server_account_management,omitempty"`
	// The shared username to use for root accounts
	SharedAdminUserName *string `json:"shared_admin_user_name,omitempty"`
	// The shared username to use for non-root accounts
	SharedStandardUserName *string `json:"shared_standard_user_name,omitempty"`
	// The type of signature algorithm used for authentication keys. Possible values: `CERT_TYPE_ED25519_01`, `CERT_TYPE_RSA_01`, `CERT_TYPE_ECDSA_521_01`, `CERT_TYPE_ECDSA_384_01`, `CERT_TYPE_ECDSA_256_01`. Default is `CERT_TYPE_ED25519_01`.
	SshCertificateType NullableString `json:"ssh_certificate_type,omitempty"`
	// The Team associated with the Project
	Team string `json:"team"`
	// The time period in seconds before an on-demand user account expires and is removed from a Server. `0` if disabled. This is the default,
	UserOnDemandPeriod NullableInt64 `json:"user_on_demand_period,omitempty"`
}

// NewProject instantiates a new Project object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProject(id string, name string, team string) *Project {
	this := Project{}
	this.Id = id
	this.Name = name
	this.Team = team
	return &this
}

// NewProjectWithDefaults instantiates a new Project object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectWithDefaults() *Project {
	this := Project{}
	return &this
}

// GetCreateServerUsers returns the CreateServerUsers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Project) GetCreateServerUsers() bool {
	if o == nil || IsNil(o.CreateServerUsers.Get()) {
		var ret bool
		return ret
	}
	return *o.CreateServerUsers.Get()
}

// GetCreateServerUsersOk returns a tuple with the CreateServerUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Project) GetCreateServerUsersOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreateServerUsers.Get(), o.CreateServerUsers.IsSet()
}

// HasCreateServerUsers returns a boolean if a field has been set.
func (o *Project) HasCreateServerUsers() bool {
	if o != nil && o.CreateServerUsers.IsSet() {
		return true
	}

	return false
}

// SetCreateServerUsers gets a reference to the given NullableBool and assigns it to the CreateServerUsers field.
func (o *Project) SetCreateServerUsers(v bool) *Project {
	o.CreateServerUsers.Set(&v)
	return o
}

// SetCreateServerUsersNil sets the value for CreateServerUsers to be an explicit nil
func (o *Project) SetCreateServerUsersNil() *Project {
	o.CreateServerUsers.Set(nil)
	return o
}

// UnsetCreateServerUsers ensures that no value is present for CreateServerUsers, not even an explicit nil
func (o *Project) UnsetCreateServerUsers() *Project {
	o.CreateServerUsers.Unset()
	return o
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Project) GetDeletedAt() time.Time {
	if o == nil || IsNil(o.DeletedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Project) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *Project) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *Project) SetDeletedAt(v time.Time) *Project {
	o.DeletedAt.Set(&v)
	return o
}

// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *Project) SetDeletedAtNil() *Project {
	o.DeletedAt.Set(nil)
	return o
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *Project) UnsetDeletedAt() *Project {
	o.DeletedAt.Unset()
	return o
}

// GetForceSharedSshUsers returns the ForceSharedSshUsers field value if set, zero value otherwise.
func (o *Project) GetForceSharedSshUsers() bool {
	if o == nil || IsNil(o.ForceSharedSshUsers) {
		var ret bool
		return ret
	}
	return *o.ForceSharedSshUsers
}

// GetForceSharedSshUsersOk returns a tuple with the ForceSharedSshUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetForceSharedSshUsersOk() (*bool, bool) {
	if o == nil || IsNil(o.ForceSharedSshUsers) {
		return nil, false
	}
	return o.ForceSharedSshUsers, true
}

// HasForceSharedSshUsers returns a boolean if a field has been set.
func (o *Project) HasForceSharedSshUsers() bool {
	if o != nil && !IsNil(o.ForceSharedSshUsers) {
		return true
	}

	return false
}

// SetForceSharedSshUsers gets a reference to the given bool and assigns it to the ForceSharedSshUsers field.
func (o *Project) SetForceSharedSshUsers(v bool) *Project {
	o.ForceSharedSshUsers = &v
	return o
}

// GetGatewaySelector returns the GatewaySelector field value if set, zero value otherwise.
func (o *Project) GetGatewaySelector() string {
	if o == nil || IsNil(o.GatewaySelector) {
		var ret string
		return ret
	}
	return *o.GatewaySelector
}

// GetGatewaySelectorOk returns a tuple with the GatewaySelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetGatewaySelectorOk() (*string, bool) {
	if o == nil || IsNil(o.GatewaySelector) {
		return nil, false
	}
	return o.GatewaySelector, true
}

// HasGatewaySelector returns a boolean if a field has been set.
func (o *Project) HasGatewaySelector() bool {
	if o != nil && !IsNil(o.GatewaySelector) {
		return true
	}

	return false
}

// SetGatewaySelector gets a reference to the given string and assigns it to the GatewaySelector field.
func (o *Project) SetGatewaySelector(v string) *Project {
	o.GatewaySelector = &v
	return o
}

// GetId returns the Id field value
func (o *Project) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Project) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Project) SetId(v string) *Project {
	o.Id = v
	return o
}

// GetResourceGroupId returns the ResourceGroupId field value if set, zero value otherwise.
func (o *Project) GetResourceGroupId() string {
	if o == nil || IsNil(o.ResourceGroupId) {
		var ret string
		return ret
	}
	return *o.ResourceGroupId
}

// GetResourceGroupIdOk returns a tuple with the ResourceGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetResourceGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceGroupId) {
		return nil, false
	}
	return o.ResourceGroupId, true
}

// HasResourceGroupId returns a boolean if a field has been set.
func (o *Project) HasResourceGroupId() bool {
	if o != nil && !IsNil(o.ResourceGroupId) {
		return true
	}

	return false
}

// SetResourceGroupId gets a reference to the given string and assigns it to the ResourceGroupId field.
func (o *Project) SetResourceGroupId(v string) *Project {
	o.ResourceGroupId = &v
	return o
}

// GetActiveResourceCount returns the ActiveResourceCount field value if set, zero value otherwise.
func (o *Project) GetActiveResourceCount() int32 {
	if o == nil || IsNil(o.ActiveResourceCount) {
		var ret int32
		return ret
	}
	return *o.ActiveResourceCount
}

// GetActiveResourceCountOk returns a tuple with the ActiveResourceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetActiveResourceCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ActiveResourceCount) {
		return nil, false
	}
	return o.ActiveResourceCount, true
}

// HasActiveResourceCount returns a boolean if a field has been set.
func (o *Project) HasActiveResourceCount() bool {
	if o != nil && !IsNil(o.ActiveResourceCount) {
		return true
	}

	return false
}

// SetActiveResourceCount gets a reference to the given int32 and assigns it to the ActiveResourceCount field.
func (o *Project) SetActiveResourceCount(v int32) *Project {
	o.ActiveResourceCount = &v
	return o
}

// GetStaleResourceCount returns the StaleResourceCount field value if set, zero value otherwise.
func (o *Project) GetStaleResourceCount() int32 {
	if o == nil || IsNil(o.StaleResourceCount) {
		var ret int32
		return ret
	}
	return *o.StaleResourceCount
}

// GetStaleResourceCountOk returns a tuple with the StaleResourceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetStaleResourceCountOk() (*int32, bool) {
	if o == nil || IsNil(o.StaleResourceCount) {
		return nil, false
	}
	return o.StaleResourceCount, true
}

// HasStaleResourceCount returns a boolean if a field has been set.
func (o *Project) HasStaleResourceCount() bool {
	if o != nil && !IsNil(o.StaleResourceCount) {
		return true
	}

	return false
}

// SetStaleResourceCount gets a reference to the given int32 and assigns it to the StaleResourceCount field.
func (o *Project) SetStaleResourceCount(v int32) *Project {
	o.StaleResourceCount = &v
	return o
}

// GetName returns the Name field value
func (o *Project) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Project) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Project) SetName(v string) *Project {
	o.Name = v
	return o
}

// GetNextUnixGid returns the NextUnixGid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Project) GetNextUnixGid() int32 {
	if o == nil || IsNil(o.NextUnixGid.Get()) {
		var ret int32
		return ret
	}
	return *o.NextUnixGid.Get()
}

// GetNextUnixGidOk returns a tuple with the NextUnixGid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Project) GetNextUnixGidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextUnixGid.Get(), o.NextUnixGid.IsSet()
}

// HasNextUnixGid returns a boolean if a field has been set.
func (o *Project) HasNextUnixGid() bool {
	if o != nil && o.NextUnixGid.IsSet() {
		return true
	}

	return false
}

// SetNextUnixGid gets a reference to the given NullableInt32 and assigns it to the NextUnixGid field.
func (o *Project) SetNextUnixGid(v int32) *Project {
	o.NextUnixGid.Set(&v)
	return o
}

// SetNextUnixGidNil sets the value for NextUnixGid to be an explicit nil
func (o *Project) SetNextUnixGidNil() *Project {
	o.NextUnixGid.Set(nil)
	return o
}

// UnsetNextUnixGid ensures that no value is present for NextUnixGid, not even an explicit nil
func (o *Project) UnsetNextUnixGid() *Project {
	o.NextUnixGid.Unset()
	return o
}

// GetNextUnixUid returns the NextUnixUid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Project) GetNextUnixUid() int32 {
	if o == nil || IsNil(o.NextUnixUid.Get()) {
		var ret int32
		return ret
	}
	return *o.NextUnixUid.Get()
}

// GetNextUnixUidOk returns a tuple with the NextUnixUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Project) GetNextUnixUidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextUnixUid.Get(), o.NextUnixUid.IsSet()
}

// HasNextUnixUid returns a boolean if a field has been set.
func (o *Project) HasNextUnixUid() bool {
	if o != nil && o.NextUnixUid.IsSet() {
		return true
	}

	return false
}

// SetNextUnixUid gets a reference to the given NullableInt32 and assigns it to the NextUnixUid field.
func (o *Project) SetNextUnixUid(v int32) *Project {
	o.NextUnixUid.Set(&v)
	return o
}

// SetNextUnixUidNil sets the value for NextUnixUid to be an explicit nil
func (o *Project) SetNextUnixUidNil() *Project {
	o.NextUnixUid.Set(nil)
	return o
}

// UnsetNextUnixUid ensures that no value is present for NextUnixUid, not even an explicit nil
func (o *Project) UnsetNextUnixUid() *Project {
	o.NextUnixUid.Unset()
	return o
}

// GetRequirePreauthForCreds returns the RequirePreauthForCreds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Project) GetRequirePreauthForCreds() bool {
	if o == nil || IsNil(o.RequirePreauthForCreds.Get()) {
		var ret bool
		return ret
	}
	return *o.RequirePreauthForCreds.Get()
}

// GetRequirePreauthForCredsOk returns a tuple with the RequirePreauthForCreds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Project) GetRequirePreauthForCredsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequirePreauthForCreds.Get(), o.RequirePreauthForCreds.IsSet()
}

// HasRequirePreauthForCreds returns a boolean if a field has been set.
func (o *Project) HasRequirePreauthForCreds() bool {
	if o != nil && o.RequirePreauthForCreds.IsSet() {
		return true
	}

	return false
}

// SetRequirePreauthForCreds gets a reference to the given NullableBool and assigns it to the RequirePreauthForCreds field.
func (o *Project) SetRequirePreauthForCreds(v bool) *Project {
	o.RequirePreauthForCreds.Set(&v)
	return o
}

// SetRequirePreauthForCredsNil sets the value for RequirePreauthForCreds to be an explicit nil
func (o *Project) SetRequirePreauthForCredsNil() *Project {
	o.RequirePreauthForCreds.Set(nil)
	return o
}

// UnsetRequirePreauthForCreds ensures that no value is present for RequirePreauthForCreds, not even an explicit nil
func (o *Project) UnsetRequirePreauthForCreds() *Project {
	o.RequirePreauthForCreds.Unset()
	return o
}

// GetPersistentServerUserAccounts returns the PersistentServerUserAccounts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Project) GetPersistentServerUserAccounts() bool {
	if o == nil || IsNil(o.PersistentServerUserAccounts.Get()) {
		var ret bool
		return ret
	}
	return *o.PersistentServerUserAccounts.Get()
}

// GetPersistentServerUserAccountsOk returns a tuple with the PersistentServerUserAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Project) GetPersistentServerUserAccountsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.PersistentServerUserAccounts.Get(), o.PersistentServerUserAccounts.IsSet()
}

// HasPersistentServerUserAccounts returns a boolean if a field has been set.
func (o *Project) HasPersistentServerUserAccounts() bool {
	if o != nil && o.PersistentServerUserAccounts.IsSet() {
		return true
	}

	return false
}

// SetPersistentServerUserAccounts gets a reference to the given NullableBool and assigns it to the PersistentServerUserAccounts field.
func (o *Project) SetPersistentServerUserAccounts(v bool) *Project {
	o.PersistentServerUserAccounts.Set(&v)
	return o
}

// SetPersistentServerUserAccountsNil sets the value for PersistentServerUserAccounts to be an explicit nil
func (o *Project) SetPersistentServerUserAccountsNil() *Project {
	o.PersistentServerUserAccounts.Set(nil)
	return o
}

// UnsetPersistentServerUserAccounts ensures that no value is present for PersistentServerUserAccounts, not even an explicit nil
func (o *Project) UnsetPersistentServerUserAccounts() *Project {
	o.PersistentServerUserAccounts.Unset()
	return o
}

// GetServerAccountManagement returns the ServerAccountManagement field value if set, zero value otherwise.
func (o *Project) GetServerAccountManagement() bool {
	if o == nil || IsNil(o.ServerAccountManagement) {
		var ret bool
		return ret
	}
	return *o.ServerAccountManagement
}

// GetServerAccountManagementOk returns a tuple with the ServerAccountManagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetServerAccountManagementOk() (*bool, bool) {
	if o == nil || IsNil(o.ServerAccountManagement) {
		return nil, false
	}
	return o.ServerAccountManagement, true
}

// HasServerAccountManagement returns a boolean if a field has been set.
func (o *Project) HasServerAccountManagement() bool {
	if o != nil && !IsNil(o.ServerAccountManagement) {
		return true
	}

	return false
}

// SetServerAccountManagement gets a reference to the given bool and assigns it to the ServerAccountManagement field.
func (o *Project) SetServerAccountManagement(v bool) *Project {
	o.ServerAccountManagement = &v
	return o
}

// GetSharedAdminUserName returns the SharedAdminUserName field value if set, zero value otherwise.
func (o *Project) GetSharedAdminUserName() string {
	if o == nil || IsNil(o.SharedAdminUserName) {
		var ret string
		return ret
	}
	return *o.SharedAdminUserName
}

// GetSharedAdminUserNameOk returns a tuple with the SharedAdminUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetSharedAdminUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.SharedAdminUserName) {
		return nil, false
	}
	return o.SharedAdminUserName, true
}

// HasSharedAdminUserName returns a boolean if a field has been set.
func (o *Project) HasSharedAdminUserName() bool {
	if o != nil && !IsNil(o.SharedAdminUserName) {
		return true
	}

	return false
}

// SetSharedAdminUserName gets a reference to the given string and assigns it to the SharedAdminUserName field.
func (o *Project) SetSharedAdminUserName(v string) *Project {
	o.SharedAdminUserName = &v
	return o
}

// GetSharedStandardUserName returns the SharedStandardUserName field value if set, zero value otherwise.
func (o *Project) GetSharedStandardUserName() string {
	if o == nil || IsNil(o.SharedStandardUserName) {
		var ret string
		return ret
	}
	return *o.SharedStandardUserName
}

// GetSharedStandardUserNameOk returns a tuple with the SharedStandardUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetSharedStandardUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.SharedStandardUserName) {
		return nil, false
	}
	return o.SharedStandardUserName, true
}

// HasSharedStandardUserName returns a boolean if a field has been set.
func (o *Project) HasSharedStandardUserName() bool {
	if o != nil && !IsNil(o.SharedStandardUserName) {
		return true
	}

	return false
}

// SetSharedStandardUserName gets a reference to the given string and assigns it to the SharedStandardUserName field.
func (o *Project) SetSharedStandardUserName(v string) *Project {
	o.SharedStandardUserName = &v
	return o
}

// GetSshCertificateType returns the SshCertificateType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Project) GetSshCertificateType() string {
	if o == nil || IsNil(o.SshCertificateType.Get()) {
		var ret string
		return ret
	}
	return *o.SshCertificateType.Get()
}

// GetSshCertificateTypeOk returns a tuple with the SshCertificateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Project) GetSshCertificateTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SshCertificateType.Get(), o.SshCertificateType.IsSet()
}

// HasSshCertificateType returns a boolean if a field has been set.
func (o *Project) HasSshCertificateType() bool {
	if o != nil && o.SshCertificateType.IsSet() {
		return true
	}

	return false
}

// SetSshCertificateType gets a reference to the given NullableString and assigns it to the SshCertificateType field.
func (o *Project) SetSshCertificateType(v string) *Project {
	o.SshCertificateType.Set(&v)
	return o
}

// SetSshCertificateTypeNil sets the value for SshCertificateType to be an explicit nil
func (o *Project) SetSshCertificateTypeNil() *Project {
	o.SshCertificateType.Set(nil)
	return o
}

// UnsetSshCertificateType ensures that no value is present for SshCertificateType, not even an explicit nil
func (o *Project) UnsetSshCertificateType() *Project {
	o.SshCertificateType.Unset()
	return o
}

// GetTeam returns the Team field value
func (o *Project) GetTeam() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Team
}

// GetTeamOk returns a tuple with the Team field value
// and a boolean to check if the value has been set.
func (o *Project) GetTeamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Team, true
}

// SetTeam sets field value
func (o *Project) SetTeam(v string) *Project {
	o.Team = v
	return o
}

// GetUserOnDemandPeriod returns the UserOnDemandPeriod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Project) GetUserOnDemandPeriod() int64 {
	if o == nil || IsNil(o.UserOnDemandPeriod.Get()) {
		var ret int64
		return ret
	}
	return *o.UserOnDemandPeriod.Get()
}

// GetUserOnDemandPeriodOk returns a tuple with the UserOnDemandPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Project) GetUserOnDemandPeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserOnDemandPeriod.Get(), o.UserOnDemandPeriod.IsSet()
}

// HasUserOnDemandPeriod returns a boolean if a field has been set.
func (o *Project) HasUserOnDemandPeriod() bool {
	if o != nil && o.UserOnDemandPeriod.IsSet() {
		return true
	}

	return false
}

// SetUserOnDemandPeriod gets a reference to the given NullableInt64 and assigns it to the UserOnDemandPeriod field.
func (o *Project) SetUserOnDemandPeriod(v int64) *Project {
	o.UserOnDemandPeriod.Set(&v)
	return o
}

// SetUserOnDemandPeriodNil sets the value for UserOnDemandPeriod to be an explicit nil
func (o *Project) SetUserOnDemandPeriodNil() *Project {
	o.UserOnDemandPeriod.Set(nil)
	return o
}

// UnsetUserOnDemandPeriod ensures that no value is present for UserOnDemandPeriod, not even an explicit nil
func (o *Project) UnsetUserOnDemandPeriod() *Project {
	o.UserOnDemandPeriod.Unset()
	return o
}

func (o Project) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Project) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CreateServerUsers.IsSet() {
		toSerialize["create_server_users"] = o.CreateServerUsers.Get()
	}
	if o.DeletedAt.IsSet() {
		toSerialize["deleted_at"] = o.DeletedAt.Get()
	}
	if !IsNil(o.ForceSharedSshUsers) {
		toSerialize["force_shared_ssh_users"] = o.ForceSharedSshUsers
	}
	if !IsNil(o.GatewaySelector) {
		toSerialize["gateway_selector"] = o.GatewaySelector
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.ResourceGroupId) {
		toSerialize["resource_group_id"] = o.ResourceGroupId
	}
	if !IsNil(o.ActiveResourceCount) {
		toSerialize["active_resource_count"] = o.ActiveResourceCount
	}
	if !IsNil(o.StaleResourceCount) {
		toSerialize["stale_resource_count"] = o.StaleResourceCount
	}
	toSerialize["name"] = o.Name
	if o.NextUnixGid.IsSet() {
		toSerialize["next_unix_gid"] = o.NextUnixGid.Get()
	}
	if o.NextUnixUid.IsSet() {
		toSerialize["next_unix_uid"] = o.NextUnixUid.Get()
	}
	if o.RequirePreauthForCreds.IsSet() {
		toSerialize["require_preauth_for_creds"] = o.RequirePreauthForCreds.Get()
	}
	if o.PersistentServerUserAccounts.IsSet() {
		toSerialize["persistent_server_user_accounts"] = o.PersistentServerUserAccounts.Get()
	}
	if !IsNil(o.ServerAccountManagement) {
		toSerialize["server_account_management"] = o.ServerAccountManagement
	}
	if !IsNil(o.SharedAdminUserName) {
		toSerialize["shared_admin_user_name"] = o.SharedAdminUserName
	}
	if !IsNil(o.SharedStandardUserName) {
		toSerialize["shared_standard_user_name"] = o.SharedStandardUserName
	}
	if o.SshCertificateType.IsSet() {
		toSerialize["ssh_certificate_type"] = o.SshCertificateType.Get()
	}
	toSerialize["team"] = o.Team
	if o.UserOnDemandPeriod.IsSet() {
		toSerialize["user_on_demand_period"] = o.UserOnDemandPeriod.Get()
	}
	return toSerialize, nil
}

type NullableProject struct {
	value *Project
	isSet bool
}

func (v NullableProject) Get() *Project {
	return v.value
}

func (v *NullableProject) Set(val *Project) {
	v.value = val
	v.isSet = true
}

func (v NullableProject) IsSet() bool {
	return v.isSet
}

func (v *NullableProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProject(val *Project) *NullableProject {
	return &NullableProject{value: val, isSet: true}
}

func (v NullableProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
