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

// checks if the SaasApplicationInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SaasApplicationInstance{}

// SaasApplicationInstance SaaS Application
type SaasApplicationInstance struct {
	// The unique identifier of the SaaS Application Instance
	AppInstanceId string `json:"app_instance_id"`
	// The name of the SaaS Application Instance
	AppInstanceName string `json:"app_instance_name"`
	// Name of the global app that instance belong to. It comes from the application definition in the OIN app catalog.
	GlobalAppName string `json:"global_app_name"`
	// The URL of the logo for the global SaaD application. It comes from the application definition in the OIN app catalog.
	AppLogoUrl *string `json:"app_logo_url,omitempty"`
	// The URL pointing to the login page of the SaaS application
	AppLoginUrl *string `json:"app_login_url,omitempty"`
}

// NewSaasApplicationInstance instantiates a new SaasApplicationInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSaasApplicationInstance(appInstanceId string, appInstanceName string, globalAppName string) *SaasApplicationInstance {
	this := SaasApplicationInstance{}
	this.AppInstanceId = appInstanceId
	this.AppInstanceName = appInstanceName
	this.GlobalAppName = globalAppName
	return &this
}

// NewSaasApplicationInstanceWithDefaults instantiates a new SaasApplicationInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSaasApplicationInstanceWithDefaults() *SaasApplicationInstance {
	this := SaasApplicationInstance{}
	return &this
}

// GetAppInstanceId returns the AppInstanceId field value
func (o *SaasApplicationInstance) GetAppInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppInstanceId
}

// GetAppInstanceIdOk returns a tuple with the AppInstanceId field value
// and a boolean to check if the value has been set.
func (o *SaasApplicationInstance) GetAppInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppInstanceId, true
}

// SetAppInstanceId sets field value
func (o *SaasApplicationInstance) SetAppInstanceId(v string) *SaasApplicationInstance {
	o.AppInstanceId = v
	return o
}

// GetAppInstanceName returns the AppInstanceName field value
func (o *SaasApplicationInstance) GetAppInstanceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppInstanceName
}

// GetAppInstanceNameOk returns a tuple with the AppInstanceName field value
// and a boolean to check if the value has been set.
func (o *SaasApplicationInstance) GetAppInstanceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppInstanceName, true
}

// SetAppInstanceName sets field value
func (o *SaasApplicationInstance) SetAppInstanceName(v string) *SaasApplicationInstance {
	o.AppInstanceName = v
	return o
}

// GetGlobalAppName returns the GlobalAppName field value
func (o *SaasApplicationInstance) GetGlobalAppName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GlobalAppName
}

// GetGlobalAppNameOk returns a tuple with the GlobalAppName field value
// and a boolean to check if the value has been set.
func (o *SaasApplicationInstance) GetGlobalAppNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GlobalAppName, true
}

// SetGlobalAppName sets field value
func (o *SaasApplicationInstance) SetGlobalAppName(v string) *SaasApplicationInstance {
	o.GlobalAppName = v
	return o
}

// GetAppLogoUrl returns the AppLogoUrl field value if set, zero value otherwise.
func (o *SaasApplicationInstance) GetAppLogoUrl() string {
	if o == nil || IsNil(o.AppLogoUrl) {
		var ret string
		return ret
	}
	return *o.AppLogoUrl
}

// GetAppLogoUrlOk returns a tuple with the AppLogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaasApplicationInstance) GetAppLogoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AppLogoUrl) {
		return nil, false
	}
	return o.AppLogoUrl, true
}

// HasAppLogoUrl returns a boolean if a field has been set.
func (o *SaasApplicationInstance) HasAppLogoUrl() bool {
	if o != nil && !IsNil(o.AppLogoUrl) {
		return true
	}

	return false
}

// SetAppLogoUrl gets a reference to the given string and assigns it to the AppLogoUrl field.
func (o *SaasApplicationInstance) SetAppLogoUrl(v string) *SaasApplicationInstance {
	o.AppLogoUrl = &v
	return o
}

// GetAppLoginUrl returns the AppLoginUrl field value if set, zero value otherwise.
func (o *SaasApplicationInstance) GetAppLoginUrl() string {
	if o == nil || IsNil(o.AppLoginUrl) {
		var ret string
		return ret
	}
	return *o.AppLoginUrl
}

// GetAppLoginUrlOk returns a tuple with the AppLoginUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaasApplicationInstance) GetAppLoginUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AppLoginUrl) {
		return nil, false
	}
	return o.AppLoginUrl, true
}

// HasAppLoginUrl returns a boolean if a field has been set.
func (o *SaasApplicationInstance) HasAppLoginUrl() bool {
	if o != nil && !IsNil(o.AppLoginUrl) {
		return true
	}

	return false
}

// SetAppLoginUrl gets a reference to the given string and assigns it to the AppLoginUrl field.
func (o *SaasApplicationInstance) SetAppLoginUrl(v string) *SaasApplicationInstance {
	o.AppLoginUrl = &v
	return o
}

func (o SaasApplicationInstance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SaasApplicationInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["app_instance_id"] = o.AppInstanceId
	toSerialize["app_instance_name"] = o.AppInstanceName
	toSerialize["global_app_name"] = o.GlobalAppName
	if !IsNil(o.AppLogoUrl) {
		toSerialize["app_logo_url"] = o.AppLogoUrl
	}
	if !IsNil(o.AppLoginUrl) {
		toSerialize["app_login_url"] = o.AppLoginUrl
	}
	return toSerialize, nil
}

type NullableSaasApplicationInstance struct {
	value *SaasApplicationInstance
	isSet bool
}

func (v NullableSaasApplicationInstance) Get() *SaasApplicationInstance {
	return v.value
}

func (v *NullableSaasApplicationInstance) Set(val *SaasApplicationInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableSaasApplicationInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableSaasApplicationInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSaasApplicationInstance(val *SaasApplicationInstance) *NullableSaasApplicationInstance {
	return &NullableSaasApplicationInstance{value: val, isSet: true}
}

func (v NullableSaasApplicationInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSaasApplicationInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}