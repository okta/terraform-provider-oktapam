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

// checks if the GatewaySetupToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewaySetupToken{}

// GatewaySetupToken struct for GatewaySetupToken
type GatewaySetupToken struct {
	// The ID of the Gateway Setup Token
	Id string `json:"id"`
	// The name for the Gateway Setup Token
	Name *string `json:"name,omitempty"`
	// The description for the Gateway Setup Token
	Description      NullableString                    `json:"description,omitempty"`
	Details          GatewaySetupTokenDetails          `json:"details"`
	RegistrationType GatewaySetupTokenRegistrationType `json:"registration_type"`
	// A timestamp indicating when the Gateway Setup Token was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The token used to enroll the Gateway
	Token *string `json:"token,omitempty"`
}

// NewGatewaySetupToken instantiates a new GatewaySetupToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewaySetupToken(id string, details GatewaySetupTokenDetails, registrationType GatewaySetupTokenRegistrationType) *GatewaySetupToken {
	this := GatewaySetupToken{}
	this.Id = id
	this.Details = details
	this.RegistrationType = registrationType
	return &this
}

// NewGatewaySetupTokenWithDefaults instantiates a new GatewaySetupToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewaySetupTokenWithDefaults() *GatewaySetupToken {
	this := GatewaySetupToken{}
	return &this
}

// GetId returns the Id field value
func (o *GatewaySetupToken) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GatewaySetupToken) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GatewaySetupToken) SetId(v string) *GatewaySetupToken {
	o.Id = v
	return o
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GatewaySetupToken) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewaySetupToken) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GatewaySetupToken) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GatewaySetupToken) SetName(v string) *GatewaySetupToken {
	o.Name = &v
	return o
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GatewaySetupToken) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GatewaySetupToken) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *GatewaySetupToken) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *GatewaySetupToken) SetDescription(v string) *GatewaySetupToken {
	o.Description.Set(&v)
	return o
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *GatewaySetupToken) SetDescriptionNil() *GatewaySetupToken {
	o.Description.Set(nil)
	return o
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *GatewaySetupToken) UnsetDescription() *GatewaySetupToken {
	o.Description.Unset()
	return o
}

// GetDetails returns the Details field value
func (o *GatewaySetupToken) GetDetails() GatewaySetupTokenDetails {
	if o == nil {
		var ret GatewaySetupTokenDetails
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *GatewaySetupToken) GetDetailsOk() (*GatewaySetupTokenDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *GatewaySetupToken) SetDetails(v GatewaySetupTokenDetails) *GatewaySetupToken {
	o.Details = v
	return o
}

// GetRegistrationType returns the RegistrationType field value
func (o *GatewaySetupToken) GetRegistrationType() GatewaySetupTokenRegistrationType {
	if o == nil {
		var ret GatewaySetupTokenRegistrationType
		return ret
	}

	return o.RegistrationType
}

// GetRegistrationTypeOk returns a tuple with the RegistrationType field value
// and a boolean to check if the value has been set.
func (o *GatewaySetupToken) GetRegistrationTypeOk() (*GatewaySetupTokenRegistrationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegistrationType, true
}

// SetRegistrationType sets field value
func (o *GatewaySetupToken) SetRegistrationType(v GatewaySetupTokenRegistrationType) *GatewaySetupToken {
	o.RegistrationType = v
	return o
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GatewaySetupToken) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewaySetupToken) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GatewaySetupToken) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *GatewaySetupToken) SetCreatedAt(v time.Time) *GatewaySetupToken {
	o.CreatedAt = &v
	return o
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewaySetupToken) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewaySetupToken) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewaySetupToken) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewaySetupToken) SetToken(v string) *GatewaySetupToken {
	o.Token = &v
	return o
}

func (o GatewaySetupToken) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewaySetupToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["details"] = o.Details
	toSerialize["registration_type"] = o.RegistrationType
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

type NullableGatewaySetupToken struct {
	value *GatewaySetupToken
	isSet bool
}

func (v NullableGatewaySetupToken) Get() *GatewaySetupToken {
	return v.value
}

func (v *NullableGatewaySetupToken) Set(val *GatewaySetupToken) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewaySetupToken) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewaySetupToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewaySetupToken(val *GatewaySetupToken) *NullableGatewaySetupToken {
	return &NullableGatewaySetupToken{value: val, isSet: true}
}

func (v NullableGatewaySetupToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewaySetupToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}