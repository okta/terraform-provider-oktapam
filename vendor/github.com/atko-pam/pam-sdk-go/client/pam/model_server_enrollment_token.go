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

// checks if the ServerEnrollmentToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerEnrollmentToken{}

// ServerEnrollmentToken struct for ServerEnrollmentToken
type ServerEnrollmentToken struct {
	// The User that created this Server Enrollment Token
	CreatedByUser string `json:"created_by_user"`
	// A human-readable description of the purpose of this Server Enrollment Token
	Description string `json:"description"`
	// The UUID of a Server Enrollment Token
	Id string `json:"id"`
	// A timestamp indicating when the Server Enrollment Token was created
	IssuedAt time.Time `json:"issued_at"`
	// A token used to enroll a Server
	Token map[string]string `json:"token"`
}

// NewServerEnrollmentToken instantiates a new ServerEnrollmentToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerEnrollmentToken(createdByUser string, description string, id string, issuedAt time.Time, token map[string]string) *ServerEnrollmentToken {
	this := ServerEnrollmentToken{}
	this.CreatedByUser = createdByUser
	this.Description = description
	this.Id = id
	this.IssuedAt = issuedAt
	this.Token = token
	return &this
}

// NewServerEnrollmentTokenWithDefaults instantiates a new ServerEnrollmentToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerEnrollmentTokenWithDefaults() *ServerEnrollmentToken {
	this := ServerEnrollmentToken{}
	return &this
}

// GetCreatedByUser returns the CreatedByUser field value
func (o *ServerEnrollmentToken) GetCreatedByUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedByUser
}

// GetCreatedByUserOk returns a tuple with the CreatedByUser field value
// and a boolean to check if the value has been set.
func (o *ServerEnrollmentToken) GetCreatedByUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedByUser, true
}

// SetCreatedByUser sets field value
func (o *ServerEnrollmentToken) SetCreatedByUser(v string) *ServerEnrollmentToken {
	o.CreatedByUser = v
	return o
}

// GetDescription returns the Description field value
func (o *ServerEnrollmentToken) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ServerEnrollmentToken) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ServerEnrollmentToken) SetDescription(v string) *ServerEnrollmentToken {
	o.Description = v
	return o
}

// GetId returns the Id field value
func (o *ServerEnrollmentToken) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ServerEnrollmentToken) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ServerEnrollmentToken) SetId(v string) *ServerEnrollmentToken {
	o.Id = v
	return o
}

// GetIssuedAt returns the IssuedAt field value
func (o *ServerEnrollmentToken) GetIssuedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.IssuedAt
}

// GetIssuedAtOk returns a tuple with the IssuedAt field value
// and a boolean to check if the value has been set.
func (o *ServerEnrollmentToken) GetIssuedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuedAt, true
}

// SetIssuedAt sets field value
func (o *ServerEnrollmentToken) SetIssuedAt(v time.Time) *ServerEnrollmentToken {
	o.IssuedAt = v
	return o
}

// GetToken returns the Token field value
func (o *ServerEnrollmentToken) GetToken() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *ServerEnrollmentToken) GetTokenOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *ServerEnrollmentToken) SetToken(v map[string]string) *ServerEnrollmentToken {
	o.Token = v
	return o
}

func (o ServerEnrollmentToken) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerEnrollmentToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created_by_user"] = o.CreatedByUser
	toSerialize["description"] = o.Description
	toSerialize["id"] = o.Id
	toSerialize["issued_at"] = o.IssuedAt
	toSerialize["token"] = o.Token
	return toSerialize, nil
}

type NullableServerEnrollmentToken struct {
	value *ServerEnrollmentToken
	isSet bool
}

func (v NullableServerEnrollmentToken) Get() *ServerEnrollmentToken {
	return v.value
}

func (v *NullableServerEnrollmentToken) Set(val *ServerEnrollmentToken) {
	v.value = val
	v.isSet = true
}

func (v NullableServerEnrollmentToken) IsSet() bool {
	return v.isSet
}

func (v *NullableServerEnrollmentToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerEnrollmentToken(val *ServerEnrollmentToken) *NullableServerEnrollmentToken {
	return &NullableServerEnrollmentToken{value: val, isSet: true}
}

func (v NullableServerEnrollmentToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerEnrollmentToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}