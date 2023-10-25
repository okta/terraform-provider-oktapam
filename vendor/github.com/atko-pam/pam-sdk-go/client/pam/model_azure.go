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
)

// checks if the Azure type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Azure{}

// Azure struct for Azure
type Azure struct {
	// IP Address
	ExternalIpV4 NullableString `json:"external_ip_v4,omitempty"`
	// IP Address
	InternalIpV4 NullableString `json:"internal_ip_v4,omitempty"`
	Location     NullableString `json:"location,omitempty"`
	VmSize       NullableString `json:"vm_size,omitempty"`
	Vmid         NullableString `json:"vmid,omitempty"`
}

// NewAzure instantiates a new Azure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzure() *Azure {
	this := Azure{}
	return &this
}

// NewAzureWithDefaults instantiates a new Azure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureWithDefaults() *Azure {
	this := Azure{}
	return &this
}

// GetExternalIpV4 returns the ExternalIpV4 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Azure) GetExternalIpV4() string {
	if o == nil || IsNil(o.ExternalIpV4.Get()) {
		var ret string
		return ret
	}
	return *o.ExternalIpV4.Get()
}

// GetExternalIpV4Ok returns a tuple with the ExternalIpV4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Azure) GetExternalIpV4Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalIpV4.Get(), o.ExternalIpV4.IsSet()
}

// HasExternalIpV4 returns a boolean if a field has been set.
func (o *Azure) HasExternalIpV4() bool {
	if o != nil && o.ExternalIpV4.IsSet() {
		return true
	}

	return false
}

// SetExternalIpV4 gets a reference to the given NullableString and assigns it to the ExternalIpV4 field.
func (o *Azure) SetExternalIpV4(v string) *Azure {
	o.ExternalIpV4.Set(&v)
	return o
}

// SetExternalIpV4Nil sets the value for ExternalIpV4 to be an explicit nil
func (o *Azure) SetExternalIpV4Nil() *Azure {
	o.ExternalIpV4.Set(nil)
	return o
}

// UnsetExternalIpV4 ensures that no value is present for ExternalIpV4, not even an explicit nil
func (o *Azure) UnsetExternalIpV4() *Azure {
	o.ExternalIpV4.Unset()
	return o
}

// GetInternalIpV4 returns the InternalIpV4 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Azure) GetInternalIpV4() string {
	if o == nil || IsNil(o.InternalIpV4.Get()) {
		var ret string
		return ret
	}
	return *o.InternalIpV4.Get()
}

// GetInternalIpV4Ok returns a tuple with the InternalIpV4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Azure) GetInternalIpV4Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InternalIpV4.Get(), o.InternalIpV4.IsSet()
}

// HasInternalIpV4 returns a boolean if a field has been set.
func (o *Azure) HasInternalIpV4() bool {
	if o != nil && o.InternalIpV4.IsSet() {
		return true
	}

	return false
}

// SetInternalIpV4 gets a reference to the given NullableString and assigns it to the InternalIpV4 field.
func (o *Azure) SetInternalIpV4(v string) *Azure {
	o.InternalIpV4.Set(&v)
	return o
}

// SetInternalIpV4Nil sets the value for InternalIpV4 to be an explicit nil
func (o *Azure) SetInternalIpV4Nil() *Azure {
	o.InternalIpV4.Set(nil)
	return o
}

// UnsetInternalIpV4 ensures that no value is present for InternalIpV4, not even an explicit nil
func (o *Azure) UnsetInternalIpV4() *Azure {
	o.InternalIpV4.Unset()
	return o
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Azure) GetLocation() string {
	if o == nil || IsNil(o.Location.Get()) {
		var ret string
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Azure) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *Azure) HasLocation() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableString and assigns it to the Location field.
func (o *Azure) SetLocation(v string) *Azure {
	o.Location.Set(&v)
	return o
}

// SetLocationNil sets the value for Location to be an explicit nil
func (o *Azure) SetLocationNil() *Azure {
	o.Location.Set(nil)
	return o
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *Azure) UnsetLocation() *Azure {
	o.Location.Unset()
	return o
}

// GetVmSize returns the VmSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Azure) GetVmSize() string {
	if o == nil || IsNil(o.VmSize.Get()) {
		var ret string
		return ret
	}
	return *o.VmSize.Get()
}

// GetVmSizeOk returns a tuple with the VmSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Azure) GetVmSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VmSize.Get(), o.VmSize.IsSet()
}

// HasVmSize returns a boolean if a field has been set.
func (o *Azure) HasVmSize() bool {
	if o != nil && o.VmSize.IsSet() {
		return true
	}

	return false
}

// SetVmSize gets a reference to the given NullableString and assigns it to the VmSize field.
func (o *Azure) SetVmSize(v string) *Azure {
	o.VmSize.Set(&v)
	return o
}

// SetVmSizeNil sets the value for VmSize to be an explicit nil
func (o *Azure) SetVmSizeNil() *Azure {
	o.VmSize.Set(nil)
	return o
}

// UnsetVmSize ensures that no value is present for VmSize, not even an explicit nil
func (o *Azure) UnsetVmSize() *Azure {
	o.VmSize.Unset()
	return o
}

// GetVmid returns the Vmid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Azure) GetVmid() string {
	if o == nil || IsNil(o.Vmid.Get()) {
		var ret string
		return ret
	}
	return *o.Vmid.Get()
}

// GetVmidOk returns a tuple with the Vmid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Azure) GetVmidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vmid.Get(), o.Vmid.IsSet()
}

// HasVmid returns a boolean if a field has been set.
func (o *Azure) HasVmid() bool {
	if o != nil && o.Vmid.IsSet() {
		return true
	}

	return false
}

// SetVmid gets a reference to the given NullableString and assigns it to the Vmid field.
func (o *Azure) SetVmid(v string) *Azure {
	o.Vmid.Set(&v)
	return o
}

// SetVmidNil sets the value for Vmid to be an explicit nil
func (o *Azure) SetVmidNil() *Azure {
	o.Vmid.Set(nil)
	return o
}

// UnsetVmid ensures that no value is present for Vmid, not even an explicit nil
func (o *Azure) UnsetVmid() *Azure {
	o.Vmid.Unset()
	return o
}

func (o Azure) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Azure) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ExternalIpV4.IsSet() {
		toSerialize["external_ip_v4"] = o.ExternalIpV4.Get()
	}
	if o.InternalIpV4.IsSet() {
		toSerialize["internal_ip_v4"] = o.InternalIpV4.Get()
	}
	if o.Location.IsSet() {
		toSerialize["location"] = o.Location.Get()
	}
	if o.VmSize.IsSet() {
		toSerialize["vm_size"] = o.VmSize.Get()
	}
	if o.Vmid.IsSet() {
		toSerialize["vmid"] = o.Vmid.Get()
	}
	return toSerialize, nil
}

type NullableAzure struct {
	value *Azure
	isSet bool
}

func (v NullableAzure) Get() *Azure {
	return v.value
}

func (v *NullableAzure) Set(val *Azure) {
	v.value = val
	v.isSet = true
}

func (v NullableAzure) IsSet() bool {
	return v.isSet
}

func (v *NullableAzure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzure(val *Azure) *NullableAzure {
	return &NullableAzure{value: val, isSet: true}
}

func (v NullableAzure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
