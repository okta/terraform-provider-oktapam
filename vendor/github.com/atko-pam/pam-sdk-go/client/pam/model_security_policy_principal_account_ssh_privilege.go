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

// checks if the SecurityPolicyPrincipalAccountSSHPrivilege type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityPolicyPrincipalAccountSSHPrivilege{}

// SecurityPolicyPrincipalAccountSSHPrivilege SecurityPolicyPrincipalAccountSSHPrivilege indicates that the principal has the ability to SSH as the principal user on the specified targets. The principal account is the \"classic ASA\" account that is created and managed by the server agent.
type SecurityPolicyPrincipalAccountSSHPrivilege struct {
	SecurityPolicyPrivilege
	PrincipalAccountSsh   bool          `json:"principal_account_ssh"`
	AdminLevelPermissions *bool         `json:"admin_level_permissions,omitempty"`
	SudoDisplayName       *string       `json:"sudo_display_name,omitempty"`
	SudoCommandBundles    []NamedObject `json:"sudo_command_bundles,omitempty"`
}

// NewSecurityPolicyPrincipalAccountSSHPrivilege instantiates a new SecurityPolicyPrincipalAccountSSHPrivilege object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityPolicyPrincipalAccountSSHPrivilege(principalAccountSsh bool, type_ SecurityPolicyRulePrivilegeType) *SecurityPolicyPrincipalAccountSSHPrivilege {
	this := SecurityPolicyPrincipalAccountSSHPrivilege{}
	this.Type = type_
	this.PrincipalAccountSsh = principalAccountSsh
	return &this
}

// NewSecurityPolicyPrincipalAccountSSHPrivilegeWithDefaults instantiates a new SecurityPolicyPrincipalAccountSSHPrivilege object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityPolicyPrincipalAccountSSHPrivilegeWithDefaults() *SecurityPolicyPrincipalAccountSSHPrivilege {
	this := SecurityPolicyPrincipalAccountSSHPrivilege{}
	return &this
}

// GetPrincipalAccountSsh returns the PrincipalAccountSsh field value
func (o *SecurityPolicyPrincipalAccountSSHPrivilege) GetPrincipalAccountSsh() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PrincipalAccountSsh
}

// GetPrincipalAccountSshOk returns a tuple with the PrincipalAccountSsh field value
// and a boolean to check if the value has been set.
func (o *SecurityPolicyPrincipalAccountSSHPrivilege) GetPrincipalAccountSshOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrincipalAccountSsh, true
}

// SetPrincipalAccountSsh sets field value
func (o *SecurityPolicyPrincipalAccountSSHPrivilege) SetPrincipalAccountSsh(v bool) *SecurityPolicyPrincipalAccountSSHPrivilege {
	o.PrincipalAccountSsh = v
	return o
}

// GetAdminLevelPermissions returns the AdminLevelPermissions field value if set, zero value otherwise.
func (o *SecurityPolicyPrincipalAccountSSHPrivilege) GetAdminLevelPermissions() bool {
	if o == nil || IsNil(o.AdminLevelPermissions) {
		var ret bool
		return ret
	}
	return *o.AdminLevelPermissions
}

// GetAdminLevelPermissionsOk returns a tuple with the AdminLevelPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityPolicyPrincipalAccountSSHPrivilege) GetAdminLevelPermissionsOk() (*bool, bool) {
	if o == nil || IsNil(o.AdminLevelPermissions) {
		return nil, false
	}
	return o.AdminLevelPermissions, true
}

// HasAdminLevelPermissions returns a boolean if a field has been set.
func (o *SecurityPolicyPrincipalAccountSSHPrivilege) HasAdminLevelPermissions() bool {
	if o != nil && !IsNil(o.AdminLevelPermissions) {
		return true
	}

	return false
}

// SetAdminLevelPermissions gets a reference to the given bool and assigns it to the AdminLevelPermissions field.
func (o *SecurityPolicyPrincipalAccountSSHPrivilege) SetAdminLevelPermissions(v bool) *SecurityPolicyPrincipalAccountSSHPrivilege {
	o.AdminLevelPermissions = &v
	return o
}

// GetSudoDisplayName returns the SudoDisplayName field value if set, zero value otherwise.
func (o *SecurityPolicyPrincipalAccountSSHPrivilege) GetSudoDisplayName() string {
	if o == nil || IsNil(o.SudoDisplayName) {
		var ret string
		return ret
	}
	return *o.SudoDisplayName
}

// GetSudoDisplayNameOk returns a tuple with the SudoDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityPolicyPrincipalAccountSSHPrivilege) GetSudoDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.SudoDisplayName) {
		return nil, false
	}
	return o.SudoDisplayName, true
}

// HasSudoDisplayName returns a boolean if a field has been set.
func (o *SecurityPolicyPrincipalAccountSSHPrivilege) HasSudoDisplayName() bool {
	if o != nil && !IsNil(o.SudoDisplayName) {
		return true
	}

	return false
}

// SetSudoDisplayName gets a reference to the given string and assigns it to the SudoDisplayName field.
func (o *SecurityPolicyPrincipalAccountSSHPrivilege) SetSudoDisplayName(v string) *SecurityPolicyPrincipalAccountSSHPrivilege {
	o.SudoDisplayName = &v
	return o
}

// GetSudoCommandBundles returns the SudoCommandBundles field value if set, zero value otherwise.
func (o *SecurityPolicyPrincipalAccountSSHPrivilege) GetSudoCommandBundles() []NamedObject {
	if o == nil || IsNil(o.SudoCommandBundles) {
		var ret []NamedObject
		return ret
	}
	return o.SudoCommandBundles
}

// GetSudoCommandBundlesOk returns a tuple with the SudoCommandBundles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityPolicyPrincipalAccountSSHPrivilege) GetSudoCommandBundlesOk() ([]NamedObject, bool) {
	if o == nil || IsNil(o.SudoCommandBundles) {
		return nil, false
	}
	return o.SudoCommandBundles, true
}

// HasSudoCommandBundles returns a boolean if a field has been set.
func (o *SecurityPolicyPrincipalAccountSSHPrivilege) HasSudoCommandBundles() bool {
	if o != nil && !IsNil(o.SudoCommandBundles) {
		return true
	}

	return false
}

// SetSudoCommandBundles gets a reference to the given []NamedObject and assigns it to the SudoCommandBundles field.
func (o *SecurityPolicyPrincipalAccountSSHPrivilege) SetSudoCommandBundles(v []NamedObject) *SecurityPolicyPrincipalAccountSSHPrivilege {
	o.SudoCommandBundles = v
	return o
}

func (o SecurityPolicyPrincipalAccountSSHPrivilege) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityPolicyPrincipalAccountSSHPrivilege) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedSecurityPolicyPrivilege, errSecurityPolicyPrivilege := json.Marshal(o.SecurityPolicyPrivilege)
	if errSecurityPolicyPrivilege != nil {
		return map[string]interface{}{}, errSecurityPolicyPrivilege
	}
	errSecurityPolicyPrivilege = json.Unmarshal([]byte(serializedSecurityPolicyPrivilege), &toSerialize)
	if errSecurityPolicyPrivilege != nil {
		return map[string]interface{}{}, errSecurityPolicyPrivilege
	}
	toSerialize["principal_account_ssh"] = o.PrincipalAccountSsh
	if !IsNil(o.AdminLevelPermissions) {
		toSerialize["admin_level_permissions"] = o.AdminLevelPermissions
	}
	if !IsNil(o.SudoDisplayName) {
		toSerialize["sudo_display_name"] = o.SudoDisplayName
	}
	if !IsNil(o.SudoCommandBundles) {
		toSerialize["sudo_command_bundles"] = o.SudoCommandBundles
	}
	return toSerialize, nil
}

type NullableSecurityPolicyPrincipalAccountSSHPrivilege struct {
	value *SecurityPolicyPrincipalAccountSSHPrivilege
	isSet bool
}

func (v NullableSecurityPolicyPrincipalAccountSSHPrivilege) Get() *SecurityPolicyPrincipalAccountSSHPrivilege {
	return v.value
}

func (v *NullableSecurityPolicyPrincipalAccountSSHPrivilege) Set(val *SecurityPolicyPrincipalAccountSSHPrivilege) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityPolicyPrincipalAccountSSHPrivilege) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityPolicyPrincipalAccountSSHPrivilege) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityPolicyPrincipalAccountSSHPrivilege(val *SecurityPolicyPrincipalAccountSSHPrivilege) *NullableSecurityPolicyPrincipalAccountSSHPrivilege {
	return &NullableSecurityPolicyPrincipalAccountSSHPrivilege{value: val, isSet: true}
}

func (v NullableSecurityPolicyPrincipalAccountSSHPrivilege) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityPolicyPrincipalAccountSSHPrivilege) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
