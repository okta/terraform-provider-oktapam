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

// checks if the UserAccessMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserAccessMethod{}

// UserAccessMethod struct for UserAccessMethod
type UserAccessMethod struct {
	// The user account that will be used to access the resource
	Identity *string `json:"identity,omitempty"`
	// The user credential used to access the resource
	AccessCredential *string `json:"access_credential,omitempty"`
	// If `true`, the connection is brokered by the server agent
	Brokered *bool `json:"brokered,omitempty"`
	// A short description used to identify the access method to users interface to help the user pick this access method vs other ones
	ShortText *string `json:"short_text,omitempty"`
	// The ID of the resource
	ServerId *string `json:"server_id,omitempty"`
	// A list of required conditions. Each condition must be met to use this access method
	Conditionals []UserAccessConditional `json:"conditionals,omitempty"`
	// The ID of an existing Security Policy Rule used to filter user access methods
	SecurityPolicyRuleId *string `json:"security_policy_rule_id,omitempty"`
	// A list of rule IDs that result in identical user access methods
	RuleIds []string `json:"rule_ids,omitempty"`
	// The type of access method
	UserAccessType *string                  `json:"user_access_type,omitempty"`
	Details        *UserAccessMethodDetails `json:"details,omitempty"`
	// Collection of all the sudo-related commands a user can access in a single string format
	SudoCommandBundles *string `json:"sudo_command_bundles,omitempty"`
	// This is a nickname set on the policy rule that contains one or more sudo command bundles. This helps Users choose the appropriate user access method based on the sudo permissions it grants.
	SudoDisplayName *string `json:"sudo_display_name,omitempty"`
}

// NewUserAccessMethod instantiates a new UserAccessMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAccessMethod() *UserAccessMethod {
	this := UserAccessMethod{}
	return &this
}

// NewUserAccessMethodWithDefaults instantiates a new UserAccessMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAccessMethodWithDefaults() *UserAccessMethod {
	this := UserAccessMethod{}
	return &this
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *UserAccessMethod) GetIdentity() string {
	if o == nil || IsNil(o.Identity) {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAccessMethod) GetIdentityOk() (*string, bool) {
	if o == nil || IsNil(o.Identity) {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *UserAccessMethod) HasIdentity() bool {
	if o != nil && !IsNil(o.Identity) {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *UserAccessMethod) SetIdentity(v string) *UserAccessMethod {
	o.Identity = &v
	return o
}

// GetAccessCredential returns the AccessCredential field value if set, zero value otherwise.
func (o *UserAccessMethod) GetAccessCredential() string {
	if o == nil || IsNil(o.AccessCredential) {
		var ret string
		return ret
	}
	return *o.AccessCredential
}

// GetAccessCredentialOk returns a tuple with the AccessCredential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAccessMethod) GetAccessCredentialOk() (*string, bool) {
	if o == nil || IsNil(o.AccessCredential) {
		return nil, false
	}
	return o.AccessCredential, true
}

// HasAccessCredential returns a boolean if a field has been set.
func (o *UserAccessMethod) HasAccessCredential() bool {
	if o != nil && !IsNil(o.AccessCredential) {
		return true
	}

	return false
}

// SetAccessCredential gets a reference to the given string and assigns it to the AccessCredential field.
func (o *UserAccessMethod) SetAccessCredential(v string) *UserAccessMethod {
	o.AccessCredential = &v
	return o
}

// GetBrokered returns the Brokered field value if set, zero value otherwise.
func (o *UserAccessMethod) GetBrokered() bool {
	if o == nil || IsNil(o.Brokered) {
		var ret bool
		return ret
	}
	return *o.Brokered
}

// GetBrokeredOk returns a tuple with the Brokered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAccessMethod) GetBrokeredOk() (*bool, bool) {
	if o == nil || IsNil(o.Brokered) {
		return nil, false
	}
	return o.Brokered, true
}

// HasBrokered returns a boolean if a field has been set.
func (o *UserAccessMethod) HasBrokered() bool {
	if o != nil && !IsNil(o.Brokered) {
		return true
	}

	return false
}

// SetBrokered gets a reference to the given bool and assigns it to the Brokered field.
func (o *UserAccessMethod) SetBrokered(v bool) *UserAccessMethod {
	o.Brokered = &v
	return o
}

// GetShortText returns the ShortText field value if set, zero value otherwise.
func (o *UserAccessMethod) GetShortText() string {
	if o == nil || IsNil(o.ShortText) {
		var ret string
		return ret
	}
	return *o.ShortText
}

// GetShortTextOk returns a tuple with the ShortText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAccessMethod) GetShortTextOk() (*string, bool) {
	if o == nil || IsNil(o.ShortText) {
		return nil, false
	}
	return o.ShortText, true
}

// HasShortText returns a boolean if a field has been set.
func (o *UserAccessMethod) HasShortText() bool {
	if o != nil && !IsNil(o.ShortText) {
		return true
	}

	return false
}

// SetShortText gets a reference to the given string and assigns it to the ShortText field.
func (o *UserAccessMethod) SetShortText(v string) *UserAccessMethod {
	o.ShortText = &v
	return o
}

// GetServerId returns the ServerId field value if set, zero value otherwise.
func (o *UserAccessMethod) GetServerId() string {
	if o == nil || IsNil(o.ServerId) {
		var ret string
		return ret
	}
	return *o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAccessMethod) GetServerIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServerId) {
		return nil, false
	}
	return o.ServerId, true
}

// HasServerId returns a boolean if a field has been set.
func (o *UserAccessMethod) HasServerId() bool {
	if o != nil && !IsNil(o.ServerId) {
		return true
	}

	return false
}

// SetServerId gets a reference to the given string and assigns it to the ServerId field.
func (o *UserAccessMethod) SetServerId(v string) *UserAccessMethod {
	o.ServerId = &v
	return o
}

// GetConditionals returns the Conditionals field value if set, zero value otherwise.
func (o *UserAccessMethod) GetConditionals() []UserAccessConditional {
	if o == nil || IsNil(o.Conditionals) {
		var ret []UserAccessConditional
		return ret
	}
	return o.Conditionals
}

// GetConditionalsOk returns a tuple with the Conditionals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAccessMethod) GetConditionalsOk() ([]UserAccessConditional, bool) {
	if o == nil || IsNil(o.Conditionals) {
		return nil, false
	}
	return o.Conditionals, true
}

// HasConditionals returns a boolean if a field has been set.
func (o *UserAccessMethod) HasConditionals() bool {
	if o != nil && !IsNil(o.Conditionals) {
		return true
	}

	return false
}

// SetConditionals gets a reference to the given []UserAccessConditional and assigns it to the Conditionals field.
func (o *UserAccessMethod) SetConditionals(v []UserAccessConditional) *UserAccessMethod {
	o.Conditionals = v
	return o
}

// GetSecurityPolicyRuleId returns the SecurityPolicyRuleId field value if set, zero value otherwise.
func (o *UserAccessMethod) GetSecurityPolicyRuleId() string {
	if o == nil || IsNil(o.SecurityPolicyRuleId) {
		var ret string
		return ret
	}
	return *o.SecurityPolicyRuleId
}

// GetSecurityPolicyRuleIdOk returns a tuple with the SecurityPolicyRuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAccessMethod) GetSecurityPolicyRuleIdOk() (*string, bool) {
	if o == nil || IsNil(o.SecurityPolicyRuleId) {
		return nil, false
	}
	return o.SecurityPolicyRuleId, true
}

// HasSecurityPolicyRuleId returns a boolean if a field has been set.
func (o *UserAccessMethod) HasSecurityPolicyRuleId() bool {
	if o != nil && !IsNil(o.SecurityPolicyRuleId) {
		return true
	}

	return false
}

// SetSecurityPolicyRuleId gets a reference to the given string and assigns it to the SecurityPolicyRuleId field.
func (o *UserAccessMethod) SetSecurityPolicyRuleId(v string) *UserAccessMethod {
	o.SecurityPolicyRuleId = &v
	return o
}

// GetRuleIds returns the RuleIds field value if set, zero value otherwise.
func (o *UserAccessMethod) GetRuleIds() []string {
	if o == nil || IsNil(o.RuleIds) {
		var ret []string
		return ret
	}
	return o.RuleIds
}

// GetRuleIdsOk returns a tuple with the RuleIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAccessMethod) GetRuleIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.RuleIds) {
		return nil, false
	}
	return o.RuleIds, true
}

// HasRuleIds returns a boolean if a field has been set.
func (o *UserAccessMethod) HasRuleIds() bool {
	if o != nil && !IsNil(o.RuleIds) {
		return true
	}

	return false
}

// SetRuleIds gets a reference to the given []string and assigns it to the RuleIds field.
func (o *UserAccessMethod) SetRuleIds(v []string) *UserAccessMethod {
	o.RuleIds = v
	return o
}

// GetUserAccessType returns the UserAccessType field value if set, zero value otherwise.
func (o *UserAccessMethod) GetUserAccessType() string {
	if o == nil || IsNil(o.UserAccessType) {
		var ret string
		return ret
	}
	return *o.UserAccessType
}

// GetUserAccessTypeOk returns a tuple with the UserAccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAccessMethod) GetUserAccessTypeOk() (*string, bool) {
	if o == nil || IsNil(o.UserAccessType) {
		return nil, false
	}
	return o.UserAccessType, true
}

// HasUserAccessType returns a boolean if a field has been set.
func (o *UserAccessMethod) HasUserAccessType() bool {
	if o != nil && !IsNil(o.UserAccessType) {
		return true
	}

	return false
}

// SetUserAccessType gets a reference to the given string and assigns it to the UserAccessType field.
func (o *UserAccessMethod) SetUserAccessType(v string) *UserAccessMethod {
	o.UserAccessType = &v
	return o
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *UserAccessMethod) GetDetails() UserAccessMethodDetails {
	if o == nil || IsNil(o.Details) {
		var ret UserAccessMethodDetails
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAccessMethod) GetDetailsOk() (*UserAccessMethodDetails, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *UserAccessMethod) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given UserAccessMethodDetails and assigns it to the Details field.
func (o *UserAccessMethod) SetDetails(v UserAccessMethodDetails) *UserAccessMethod {
	o.Details = &v
	return o
}

// GetSudoCommandBundles returns the SudoCommandBundles field value if set, zero value otherwise.
func (o *UserAccessMethod) GetSudoCommandBundles() string {
	if o == nil || IsNil(o.SudoCommandBundles) {
		var ret string
		return ret
	}
	return *o.SudoCommandBundles
}

// GetSudoCommandBundlesOk returns a tuple with the SudoCommandBundles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAccessMethod) GetSudoCommandBundlesOk() (*string, bool) {
	if o == nil || IsNil(o.SudoCommandBundles) {
		return nil, false
	}
	return o.SudoCommandBundles, true
}

// HasSudoCommandBundles returns a boolean if a field has been set.
func (o *UserAccessMethod) HasSudoCommandBundles() bool {
	if o != nil && !IsNil(o.SudoCommandBundles) {
		return true
	}

	return false
}

// SetSudoCommandBundles gets a reference to the given string and assigns it to the SudoCommandBundles field.
func (o *UserAccessMethod) SetSudoCommandBundles(v string) *UserAccessMethod {
	o.SudoCommandBundles = &v
	return o
}

// GetSudoDisplayName returns the SudoDisplayName field value if set, zero value otherwise.
func (o *UserAccessMethod) GetSudoDisplayName() string {
	if o == nil || IsNil(o.SudoDisplayName) {
		var ret string
		return ret
	}
	return *o.SudoDisplayName
}

// GetSudoDisplayNameOk returns a tuple with the SudoDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAccessMethod) GetSudoDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.SudoDisplayName) {
		return nil, false
	}
	return o.SudoDisplayName, true
}

// HasSudoDisplayName returns a boolean if a field has been set.
func (o *UserAccessMethod) HasSudoDisplayName() bool {
	if o != nil && !IsNil(o.SudoDisplayName) {
		return true
	}

	return false
}

// SetSudoDisplayName gets a reference to the given string and assigns it to the SudoDisplayName field.
func (o *UserAccessMethod) SetSudoDisplayName(v string) *UserAccessMethod {
	o.SudoDisplayName = &v
	return o
}

func (o UserAccessMethod) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserAccessMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Identity) {
		toSerialize["identity"] = o.Identity
	}
	if !IsNil(o.AccessCredential) {
		toSerialize["access_credential"] = o.AccessCredential
	}
	if !IsNil(o.Brokered) {
		toSerialize["brokered"] = o.Brokered
	}
	if !IsNil(o.ShortText) {
		toSerialize["short_text"] = o.ShortText
	}
	if !IsNil(o.ServerId) {
		toSerialize["server_id"] = o.ServerId
	}
	if !IsNil(o.Conditionals) {
		toSerialize["conditionals"] = o.Conditionals
	}
	if !IsNil(o.SecurityPolicyRuleId) {
		toSerialize["security_policy_rule_id"] = o.SecurityPolicyRuleId
	}
	if !IsNil(o.RuleIds) {
		toSerialize["rule_ids"] = o.RuleIds
	}
	if !IsNil(o.UserAccessType) {
		toSerialize["user_access_type"] = o.UserAccessType
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.SudoCommandBundles) {
		toSerialize["sudo_command_bundles"] = o.SudoCommandBundles
	}
	if !IsNil(o.SudoDisplayName) {
		toSerialize["sudo_display_name"] = o.SudoDisplayName
	}
	return toSerialize, nil
}

type NullableUserAccessMethod struct {
	value *UserAccessMethod
	isSet bool
}

func (v NullableUserAccessMethod) Get() *UserAccessMethod {
	return v.value
}

func (v *NullableUserAccessMethod) Set(val *UserAccessMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAccessMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAccessMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAccessMethod(val *UserAccessMethod) *NullableUserAccessMethod {
	return &NullableUserAccessMethod{value: val, isSet: true}
}

func (v NullableUserAccessMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAccessMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
