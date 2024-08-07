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

// checks if the PrivilegedAccountDetailsAppAccountAllOfDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrivilegedAccountDetailsAppAccountAllOfDetails{}

// PrivilegedAccountDetailsAppAccountAllOfDetails struct for PrivilegedAccountDetailsAppAccountAllOfDetails
type PrivilegedAccountDetailsAppAccountAllOfDetails struct {
	Credentials PrivilegedAccountCredentials `json:"credentials"`
	// The Okta app instance ID of the SaaS Application
	OktaApplicationId string `json:"okta_application_id"`
	// The instance name of the SaaS Application
	AppInstanceName *string `json:"app_instance_name,omitempty"`
}

// NewPrivilegedAccountDetailsAppAccountAllOfDetails instantiates a new PrivilegedAccountDetailsAppAccountAllOfDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivilegedAccountDetailsAppAccountAllOfDetails(credentials PrivilegedAccountCredentials, oktaApplicationId string) *PrivilegedAccountDetailsAppAccountAllOfDetails {
	this := PrivilegedAccountDetailsAppAccountAllOfDetails{}
	this.Credentials = credentials
	this.OktaApplicationId = oktaApplicationId
	return &this
}

// NewPrivilegedAccountDetailsAppAccountAllOfDetailsWithDefaults instantiates a new PrivilegedAccountDetailsAppAccountAllOfDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivilegedAccountDetailsAppAccountAllOfDetailsWithDefaults() *PrivilegedAccountDetailsAppAccountAllOfDetails {
	this := PrivilegedAccountDetailsAppAccountAllOfDetails{}
	return &this
}

// GetCredentials returns the Credentials field value
func (o *PrivilegedAccountDetailsAppAccountAllOfDetails) GetCredentials() PrivilegedAccountCredentials {
	if o == nil {
		var ret PrivilegedAccountCredentials
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *PrivilegedAccountDetailsAppAccountAllOfDetails) GetCredentialsOk() (*PrivilegedAccountCredentials, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Credentials, true
}

// SetCredentials sets field value
func (o *PrivilegedAccountDetailsAppAccountAllOfDetails) SetCredentials(v PrivilegedAccountCredentials) *PrivilegedAccountDetailsAppAccountAllOfDetails {
	o.Credentials = v
	return o
}

// GetOktaApplicationId returns the OktaApplicationId field value
func (o *PrivilegedAccountDetailsAppAccountAllOfDetails) GetOktaApplicationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OktaApplicationId
}

// GetOktaApplicationIdOk returns a tuple with the OktaApplicationId field value
// and a boolean to check if the value has been set.
func (o *PrivilegedAccountDetailsAppAccountAllOfDetails) GetOktaApplicationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OktaApplicationId, true
}

// SetOktaApplicationId sets field value
func (o *PrivilegedAccountDetailsAppAccountAllOfDetails) SetOktaApplicationId(v string) *PrivilegedAccountDetailsAppAccountAllOfDetails {
	o.OktaApplicationId = v
	return o
}

// GetAppInstanceName returns the AppInstanceName field value if set, zero value otherwise.
func (o *PrivilegedAccountDetailsAppAccountAllOfDetails) GetAppInstanceName() string {
	if o == nil || IsNil(o.AppInstanceName) {
		var ret string
		return ret
	}
	return *o.AppInstanceName
}

// GetAppInstanceNameOk returns a tuple with the AppInstanceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivilegedAccountDetailsAppAccountAllOfDetails) GetAppInstanceNameOk() (*string, bool) {
	if o == nil || IsNil(o.AppInstanceName) {
		return nil, false
	}
	return o.AppInstanceName, true
}

// HasAppInstanceName returns a boolean if a field has been set.
func (o *PrivilegedAccountDetailsAppAccountAllOfDetails) HasAppInstanceName() bool {
	if o != nil && !IsNil(o.AppInstanceName) {
		return true
	}

	return false
}

// SetAppInstanceName gets a reference to the given string and assigns it to the AppInstanceName field.
func (o *PrivilegedAccountDetailsAppAccountAllOfDetails) SetAppInstanceName(v string) *PrivilegedAccountDetailsAppAccountAllOfDetails {
	o.AppInstanceName = &v
	return o
}

func (o PrivilegedAccountDetailsAppAccountAllOfDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrivilegedAccountDetailsAppAccountAllOfDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["credentials"] = o.Credentials
	toSerialize["okta_application_id"] = o.OktaApplicationId
	if !IsNil(o.AppInstanceName) {
		toSerialize["app_instance_name"] = o.AppInstanceName
	}
	return toSerialize, nil
}

type NullablePrivilegedAccountDetailsAppAccountAllOfDetails struct {
	value *PrivilegedAccountDetailsAppAccountAllOfDetails
	isSet bool
}

func (v NullablePrivilegedAccountDetailsAppAccountAllOfDetails) Get() *PrivilegedAccountDetailsAppAccountAllOfDetails {
	return v.value
}

func (v *NullablePrivilegedAccountDetailsAppAccountAllOfDetails) Set(val *PrivilegedAccountDetailsAppAccountAllOfDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivilegedAccountDetailsAppAccountAllOfDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivilegedAccountDetailsAppAccountAllOfDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivilegedAccountDetailsAppAccountAllOfDetails(val *PrivilegedAccountDetailsAppAccountAllOfDetails) *NullablePrivilegedAccountDetailsAppAccountAllOfDetails {
	return &NullablePrivilegedAccountDetailsAppAccountAllOfDetails{value: val, isSet: true}
}

func (v NullablePrivilegedAccountDetailsAppAccountAllOfDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivilegedAccountDetailsAppAccountAllOfDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
