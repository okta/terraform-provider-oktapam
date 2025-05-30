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

// checks if the CloudEntitlementJobAssignedUsersUserDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudEntitlementJobAssignedUsersUserDetails{}

// CloudEntitlementJobAssignedUsersUserDetails The specific details of the assigned user
type CloudEntitlementJobAssignedUsersUserDetails struct {
	// The user credential of the the user
	UserName *string `json:"user_name,omitempty"`
	// The email address associated with the user
	PrimaryEmail *string `json:"primary_email,omitempty"`
	// The display name of the user
	DisplayName *string `json:"display_name,omitempty"`
}

// NewCloudEntitlementJobAssignedUsersUserDetails instantiates a new CloudEntitlementJobAssignedUsersUserDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudEntitlementJobAssignedUsersUserDetails() *CloudEntitlementJobAssignedUsersUserDetails {
	this := CloudEntitlementJobAssignedUsersUserDetails{}
	return &this
}

// NewCloudEntitlementJobAssignedUsersUserDetailsWithDefaults instantiates a new CloudEntitlementJobAssignedUsersUserDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudEntitlementJobAssignedUsersUserDetailsWithDefaults() *CloudEntitlementJobAssignedUsersUserDetails {
	this := CloudEntitlementJobAssignedUsersUserDetails{}
	return &this
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *CloudEntitlementJobAssignedUsersUserDetails) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudEntitlementJobAssignedUsersUserDetails) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *CloudEntitlementJobAssignedUsersUserDetails) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *CloudEntitlementJobAssignedUsersUserDetails) SetUserName(v string) *CloudEntitlementJobAssignedUsersUserDetails {
	o.UserName = &v
	return o
}

// GetPrimaryEmail returns the PrimaryEmail field value if set, zero value otherwise.
func (o *CloudEntitlementJobAssignedUsersUserDetails) GetPrimaryEmail() string {
	if o == nil || IsNil(o.PrimaryEmail) {
		var ret string
		return ret
	}
	return *o.PrimaryEmail
}

// GetPrimaryEmailOk returns a tuple with the PrimaryEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudEntitlementJobAssignedUsersUserDetails) GetPrimaryEmailOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryEmail) {
		return nil, false
	}
	return o.PrimaryEmail, true
}

// HasPrimaryEmail returns a boolean if a field has been set.
func (o *CloudEntitlementJobAssignedUsersUserDetails) HasPrimaryEmail() bool {
	if o != nil && !IsNil(o.PrimaryEmail) {
		return true
	}

	return false
}

// SetPrimaryEmail gets a reference to the given string and assigns it to the PrimaryEmail field.
func (o *CloudEntitlementJobAssignedUsersUserDetails) SetPrimaryEmail(v string) *CloudEntitlementJobAssignedUsersUserDetails {
	o.PrimaryEmail = &v
	return o
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *CloudEntitlementJobAssignedUsersUserDetails) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudEntitlementJobAssignedUsersUserDetails) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CloudEntitlementJobAssignedUsersUserDetails) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *CloudEntitlementJobAssignedUsersUserDetails) SetDisplayName(v string) *CloudEntitlementJobAssignedUsersUserDetails {
	o.DisplayName = &v
	return o
}

func (o CloudEntitlementJobAssignedUsersUserDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudEntitlementJobAssignedUsersUserDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserName) {
		toSerialize["user_name"] = o.UserName
	}
	if !IsNil(o.PrimaryEmail) {
		toSerialize["primary_email"] = o.PrimaryEmail
	}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	return toSerialize, nil
}

type NullableCloudEntitlementJobAssignedUsersUserDetails struct {
	value *CloudEntitlementJobAssignedUsersUserDetails
	isSet bool
}

func (v NullableCloudEntitlementJobAssignedUsersUserDetails) Get() *CloudEntitlementJobAssignedUsersUserDetails {
	return v.value
}

func (v *NullableCloudEntitlementJobAssignedUsersUserDetails) Set(val *CloudEntitlementJobAssignedUsersUserDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudEntitlementJobAssignedUsersUserDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudEntitlementJobAssignedUsersUserDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudEntitlementJobAssignedUsersUserDetails(val *CloudEntitlementJobAssignedUsersUserDetails) *NullableCloudEntitlementJobAssignedUsersUserDetails {
	return &NullableCloudEntitlementJobAssignedUsersUserDetails{value: val, isSet: true}
}

func (v NullableCloudEntitlementJobAssignedUsersUserDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudEntitlementJobAssignedUsersUserDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
