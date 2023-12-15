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

// checks if the UserAccessMethodServerDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserAccessMethodServerDetails{}

// UserAccessMethodServerDetails struct for UserAccessMethodServerDetails
type UserAccessMethodServerDetails struct {
	// The ID of the server resource
	ServerId *string `json:"server_id,omitempty"`
	// The hostname of the server resource
	ServerHostName *string `json:"server_host_name,omitempty"`
	// The username that will be used to access the server resource
	Identity *string `json:"identity,omitempty"`
	// If `true`, the connection is brokered by the server agent
	Brokered         *bool                 `json:"brokered,omitempty"`
	AccessCredential *AccessCredentialType `json:"access_credential,omitempty"`
}

// NewUserAccessMethodServerDetails instantiates a new UserAccessMethodServerDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAccessMethodServerDetails() *UserAccessMethodServerDetails {
	this := UserAccessMethodServerDetails{}
	return &this
}

// NewUserAccessMethodServerDetailsWithDefaults instantiates a new UserAccessMethodServerDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAccessMethodServerDetailsWithDefaults() *UserAccessMethodServerDetails {
	this := UserAccessMethodServerDetails{}
	return &this
}

// GetServerId returns the ServerId field value if set, zero value otherwise.
func (o *UserAccessMethodServerDetails) GetServerId() string {
	if o == nil || IsNil(o.ServerId) {
		var ret string
		return ret
	}
	return *o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAccessMethodServerDetails) GetServerIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServerId) {
		return nil, false
	}
	return o.ServerId, true
}

// HasServerId returns a boolean if a field has been set.
func (o *UserAccessMethodServerDetails) HasServerId() bool {
	if o != nil && !IsNil(o.ServerId) {
		return true
	}

	return false
}

// SetServerId gets a reference to the given string and assigns it to the ServerId field.
func (o *UserAccessMethodServerDetails) SetServerId(v string) *UserAccessMethodServerDetails {
	o.ServerId = &v
	return o
}

// GetServerHostName returns the ServerHostName field value if set, zero value otherwise.
func (o *UserAccessMethodServerDetails) GetServerHostName() string {
	if o == nil || IsNil(o.ServerHostName) {
		var ret string
		return ret
	}
	return *o.ServerHostName
}

// GetServerHostNameOk returns a tuple with the ServerHostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAccessMethodServerDetails) GetServerHostNameOk() (*string, bool) {
	if o == nil || IsNil(o.ServerHostName) {
		return nil, false
	}
	return o.ServerHostName, true
}

// HasServerHostName returns a boolean if a field has been set.
func (o *UserAccessMethodServerDetails) HasServerHostName() bool {
	if o != nil && !IsNil(o.ServerHostName) {
		return true
	}

	return false
}

// SetServerHostName gets a reference to the given string and assigns it to the ServerHostName field.
func (o *UserAccessMethodServerDetails) SetServerHostName(v string) *UserAccessMethodServerDetails {
	o.ServerHostName = &v
	return o
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *UserAccessMethodServerDetails) GetIdentity() string {
	if o == nil || IsNil(o.Identity) {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAccessMethodServerDetails) GetIdentityOk() (*string, bool) {
	if o == nil || IsNil(o.Identity) {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *UserAccessMethodServerDetails) HasIdentity() bool {
	if o != nil && !IsNil(o.Identity) {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *UserAccessMethodServerDetails) SetIdentity(v string) *UserAccessMethodServerDetails {
	o.Identity = &v
	return o
}

// GetBrokered returns the Brokered field value if set, zero value otherwise.
func (o *UserAccessMethodServerDetails) GetBrokered() bool {
	if o == nil || IsNil(o.Brokered) {
		var ret bool
		return ret
	}
	return *o.Brokered
}

// GetBrokeredOk returns a tuple with the Brokered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAccessMethodServerDetails) GetBrokeredOk() (*bool, bool) {
	if o == nil || IsNil(o.Brokered) {
		return nil, false
	}
	return o.Brokered, true
}

// HasBrokered returns a boolean if a field has been set.
func (o *UserAccessMethodServerDetails) HasBrokered() bool {
	if o != nil && !IsNil(o.Brokered) {
		return true
	}

	return false
}

// SetBrokered gets a reference to the given bool and assigns it to the Brokered field.
func (o *UserAccessMethodServerDetails) SetBrokered(v bool) *UserAccessMethodServerDetails {
	o.Brokered = &v
	return o
}

// GetAccessCredential returns the AccessCredential field value if set, zero value otherwise.
func (o *UserAccessMethodServerDetails) GetAccessCredential() AccessCredentialType {
	if o == nil || IsNil(o.AccessCredential) {
		var ret AccessCredentialType
		return ret
	}
	return *o.AccessCredential
}

// GetAccessCredentialOk returns a tuple with the AccessCredential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAccessMethodServerDetails) GetAccessCredentialOk() (*AccessCredentialType, bool) {
	if o == nil || IsNil(o.AccessCredential) {
		return nil, false
	}
	return o.AccessCredential, true
}

// HasAccessCredential returns a boolean if a field has been set.
func (o *UserAccessMethodServerDetails) HasAccessCredential() bool {
	if o != nil && !IsNil(o.AccessCredential) {
		return true
	}

	return false
}

// SetAccessCredential gets a reference to the given AccessCredentialType and assigns it to the AccessCredential field.
func (o *UserAccessMethodServerDetails) SetAccessCredential(v AccessCredentialType) *UserAccessMethodServerDetails {
	o.AccessCredential = &v
	return o
}

func (o UserAccessMethodServerDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserAccessMethodServerDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServerId) {
		toSerialize["server_id"] = o.ServerId
	}
	if !IsNil(o.ServerHostName) {
		toSerialize["server_host_name"] = o.ServerHostName
	}
	if !IsNil(o.Identity) {
		toSerialize["identity"] = o.Identity
	}
	if !IsNil(o.Brokered) {
		toSerialize["brokered"] = o.Brokered
	}
	if !IsNil(o.AccessCredential) {
		toSerialize["access_credential"] = o.AccessCredential
	}
	return toSerialize, nil
}

type NullableUserAccessMethodServerDetails struct {
	value *UserAccessMethodServerDetails
	isSet bool
}

func (v NullableUserAccessMethodServerDetails) Get() *UserAccessMethodServerDetails {
	return v.value
}

func (v *NullableUserAccessMethodServerDetails) Set(val *UserAccessMethodServerDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAccessMethodServerDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAccessMethodServerDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAccessMethodServerDetails(val *UserAccessMethodServerDetails) *NullableUserAccessMethodServerDetails {
	return &NullableUserAccessMethodServerDetails{value: val, isSet: true}
}

func (v NullableUserAccessMethodServerDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAccessMethodServerDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}