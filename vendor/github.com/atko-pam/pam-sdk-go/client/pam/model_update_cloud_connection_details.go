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

// checks if the UpdateCloudConnectionDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateCloudConnectionDetails{}

// UpdateCloudConnectionDetails struct for UpdateCloudConnectionDetails
type UpdateCloudConnectionDetails struct {
	// An AWS account ID
	AccountId string `json:"account_id"`
	// An AWS role ARN. The Cloud Connection will use this role to access the associated AWS account.
	RoleArn string `json:"role_arn"`
}

// NewUpdateCloudConnectionDetails instantiates a new UpdateCloudConnectionDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCloudConnectionDetails(accountId string, roleArn string) *UpdateCloudConnectionDetails {
	this := UpdateCloudConnectionDetails{}
	this.AccountId = accountId
	this.RoleArn = roleArn
	return &this
}

// NewUpdateCloudConnectionDetailsWithDefaults instantiates a new UpdateCloudConnectionDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCloudConnectionDetailsWithDefaults() *UpdateCloudConnectionDetails {
	this := UpdateCloudConnectionDetails{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *UpdateCloudConnectionDetails) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *UpdateCloudConnectionDetails) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *UpdateCloudConnectionDetails) SetAccountId(v string) *UpdateCloudConnectionDetails {
	o.AccountId = v
	return o
}

// GetRoleArn returns the RoleArn field value
func (o *UpdateCloudConnectionDetails) GetRoleArn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleArn
}

// GetRoleArnOk returns a tuple with the RoleArn field value
// and a boolean to check if the value has been set.
func (o *UpdateCloudConnectionDetails) GetRoleArnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleArn, true
}

// SetRoleArn sets field value
func (o *UpdateCloudConnectionDetails) SetRoleArn(v string) *UpdateCloudConnectionDetails {
	o.RoleArn = v
	return o
}

func (o UpdateCloudConnectionDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateCloudConnectionDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_id"] = o.AccountId
	toSerialize["role_arn"] = o.RoleArn
	return toSerialize, nil
}

type NullableUpdateCloudConnectionDetails struct {
	value *UpdateCloudConnectionDetails
	isSet bool
}

func (v NullableUpdateCloudConnectionDetails) Get() *UpdateCloudConnectionDetails {
	return v.value
}

func (v *NullableUpdateCloudConnectionDetails) Set(val *UpdateCloudConnectionDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCloudConnectionDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCloudConnectionDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCloudConnectionDetails(val *UpdateCloudConnectionDetails) *NullableUpdateCloudConnectionDetails {
	return &NullableUpdateCloudConnectionDetails{value: val, isSet: true}
}

func (v NullableUpdateCloudConnectionDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCloudConnectionDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}