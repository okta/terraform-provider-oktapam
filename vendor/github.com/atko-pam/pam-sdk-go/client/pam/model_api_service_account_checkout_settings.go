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

// checks if the APIServiceAccountCheckoutSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &APIServiceAccountCheckoutSettings{}

// APIServiceAccountCheckoutSettings struct for APIServiceAccountCheckoutSettings
type APIServiceAccountCheckoutSettings struct {
	// Indicates whether a checkout is mandatory for accessing resources within the project. If `true`, checkout is enforced for all applicable resources by default. If `false`, checkout is not required, and resources are accessible without it.
	CheckoutRequired bool `json:"checkout_required"`
	// The duration in seconds for the checkout. If the checkout is enabled, the duration is the maximum time a user can access the resource before the checkout expires.
	CheckoutDurationInSeconds int32 `json:"checkout_duration_in_seconds"`
	// If provided, only the account identifiers listed are required to perform a checkout to access the resource. This list is only considered if `checkout_required` is set to `true`. Only one of `include_list` and `exclude_list` can be specified in a request since they are mutually exclusive.
	IncludeList []ServiceAccountSettingNameObject `json:"include_list,omitempty"`
	// If provided, only the account identifiers listed are excluded from the checkout requirement. This list is only considered if `checkout_required` is set to `true`. Only one of `include_list` and `exclude_list` can be specified in a request since they are mutually exclusive.
	ExcludeList []ServiceAccountSettingNameObject `json:"exclude_list,omitempty"`
}

// NewAPIServiceAccountCheckoutSettings instantiates a new APIServiceAccountCheckoutSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIServiceAccountCheckoutSettings(checkoutRequired bool, checkoutDurationInSeconds int32) *APIServiceAccountCheckoutSettings {
	this := APIServiceAccountCheckoutSettings{}
	this.CheckoutRequired = checkoutRequired
	this.CheckoutDurationInSeconds = checkoutDurationInSeconds
	return &this
}

// NewAPIServiceAccountCheckoutSettingsWithDefaults instantiates a new APIServiceAccountCheckoutSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIServiceAccountCheckoutSettingsWithDefaults() *APIServiceAccountCheckoutSettings {
	this := APIServiceAccountCheckoutSettings{}
	return &this
}

// GetCheckoutRequired returns the CheckoutRequired field value
func (o *APIServiceAccountCheckoutSettings) GetCheckoutRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CheckoutRequired
}

// GetCheckoutRequiredOk returns a tuple with the CheckoutRequired field value
// and a boolean to check if the value has been set.
func (o *APIServiceAccountCheckoutSettings) GetCheckoutRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CheckoutRequired, true
}

// SetCheckoutRequired sets field value
func (o *APIServiceAccountCheckoutSettings) SetCheckoutRequired(v bool) *APIServiceAccountCheckoutSettings {
	o.CheckoutRequired = v
	return o
}

// GetCheckoutDurationInSeconds returns the CheckoutDurationInSeconds field value
func (o *APIServiceAccountCheckoutSettings) GetCheckoutDurationInSeconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CheckoutDurationInSeconds
}

// GetCheckoutDurationInSecondsOk returns a tuple with the CheckoutDurationInSeconds field value
// and a boolean to check if the value has been set.
func (o *APIServiceAccountCheckoutSettings) GetCheckoutDurationInSecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CheckoutDurationInSeconds, true
}

// SetCheckoutDurationInSeconds sets field value
func (o *APIServiceAccountCheckoutSettings) SetCheckoutDurationInSeconds(v int32) *APIServiceAccountCheckoutSettings {
	o.CheckoutDurationInSeconds = v
	return o
}

// GetIncludeList returns the IncludeList field value if set, zero value otherwise.
func (o *APIServiceAccountCheckoutSettings) GetIncludeList() []ServiceAccountSettingNameObject {
	if o == nil || IsNil(o.IncludeList) {
		var ret []ServiceAccountSettingNameObject
		return ret
	}
	return o.IncludeList
}

// GetIncludeListOk returns a tuple with the IncludeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIServiceAccountCheckoutSettings) GetIncludeListOk() ([]ServiceAccountSettingNameObject, bool) {
	if o == nil || IsNil(o.IncludeList) {
		return nil, false
	}
	return o.IncludeList, true
}

// HasIncludeList returns a boolean if a field has been set.
func (o *APIServiceAccountCheckoutSettings) HasIncludeList() bool {
	if o != nil && !IsNil(o.IncludeList) {
		return true
	}

	return false
}

// SetIncludeList gets a reference to the given []ServiceAccountSettingNameObject and assigns it to the IncludeList field.
func (o *APIServiceAccountCheckoutSettings) SetIncludeList(v []ServiceAccountSettingNameObject) *APIServiceAccountCheckoutSettings {
	o.IncludeList = v
	return o
}

// GetExcludeList returns the ExcludeList field value if set, zero value otherwise.
func (o *APIServiceAccountCheckoutSettings) GetExcludeList() []ServiceAccountSettingNameObject {
	if o == nil || IsNil(o.ExcludeList) {
		var ret []ServiceAccountSettingNameObject
		return ret
	}
	return o.ExcludeList
}

// GetExcludeListOk returns a tuple with the ExcludeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIServiceAccountCheckoutSettings) GetExcludeListOk() ([]ServiceAccountSettingNameObject, bool) {
	if o == nil || IsNil(o.ExcludeList) {
		return nil, false
	}
	return o.ExcludeList, true
}

// HasExcludeList returns a boolean if a field has been set.
func (o *APIServiceAccountCheckoutSettings) HasExcludeList() bool {
	if o != nil && !IsNil(o.ExcludeList) {
		return true
	}

	return false
}

// SetExcludeList gets a reference to the given []ServiceAccountSettingNameObject and assigns it to the ExcludeList field.
func (o *APIServiceAccountCheckoutSettings) SetExcludeList(v []ServiceAccountSettingNameObject) *APIServiceAccountCheckoutSettings {
	o.ExcludeList = v
	return o
}

func (o APIServiceAccountCheckoutSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o APIServiceAccountCheckoutSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["checkout_required"] = o.CheckoutRequired
	toSerialize["checkout_duration_in_seconds"] = o.CheckoutDurationInSeconds
	if !IsNil(o.IncludeList) {
		toSerialize["include_list"] = o.IncludeList
	}
	if !IsNil(o.ExcludeList) {
		toSerialize["exclude_list"] = o.ExcludeList
	}
	return toSerialize, nil
}

type NullableAPIServiceAccountCheckoutSettings struct {
	value *APIServiceAccountCheckoutSettings
	isSet bool
}

func (v NullableAPIServiceAccountCheckoutSettings) Get() *APIServiceAccountCheckoutSettings {
	return v.value
}

func (v *NullableAPIServiceAccountCheckoutSettings) Set(val *APIServiceAccountCheckoutSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIServiceAccountCheckoutSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIServiceAccountCheckoutSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIServiceAccountCheckoutSettings(val *APIServiceAccountCheckoutSettings) *NullableAPIServiceAccountCheckoutSettings {
	return &NullableAPIServiceAccountCheckoutSettings{value: val, isSet: true}
}

func (v NullableAPIServiceAccountCheckoutSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIServiceAccountCheckoutSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}