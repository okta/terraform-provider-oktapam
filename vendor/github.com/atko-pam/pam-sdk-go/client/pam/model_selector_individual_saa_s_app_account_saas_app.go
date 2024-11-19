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

// checks if the SelectorIndividualSaaSAppAccountSaasApp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SelectorIndividualSaaSAppAccountSaasApp{}

// SelectorIndividualSaaSAppAccountSaasApp A named object representing the SaaS Application
type SelectorIndividualSaaSAppAccountSaasApp struct {
	// The UUID of the object
	Id *string `json:"id,omitempty"`
	// The human-readable name of the object
	Name *string          `json:"name,omitempty"`
	Type *NamedObjectType `json:"type,omitempty"`
	// Boolean value determining if the named object with the given id is valid or not.
	Missing *bool `json:"missing,omitempty"`
}

// NewSelectorIndividualSaaSAppAccountSaasApp instantiates a new SelectorIndividualSaaSAppAccountSaasApp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelectorIndividualSaaSAppAccountSaasApp() *SelectorIndividualSaaSAppAccountSaasApp {
	this := SelectorIndividualSaaSAppAccountSaasApp{}
	return &this
}

// NewSelectorIndividualSaaSAppAccountSaasAppWithDefaults instantiates a new SelectorIndividualSaaSAppAccountSaasApp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelectorIndividualSaaSAppAccountSaasAppWithDefaults() *SelectorIndividualSaaSAppAccountSaasApp {
	this := SelectorIndividualSaaSAppAccountSaasApp{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SelectorIndividualSaaSAppAccountSaasApp) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectorIndividualSaaSAppAccountSaasApp) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SelectorIndividualSaaSAppAccountSaasApp) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SelectorIndividualSaaSAppAccountSaasApp) SetId(v string) *SelectorIndividualSaaSAppAccountSaasApp {
	o.Id = &v
	return o
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SelectorIndividualSaaSAppAccountSaasApp) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectorIndividualSaaSAppAccountSaasApp) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SelectorIndividualSaaSAppAccountSaasApp) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SelectorIndividualSaaSAppAccountSaasApp) SetName(v string) *SelectorIndividualSaaSAppAccountSaasApp {
	o.Name = &v
	return o
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SelectorIndividualSaaSAppAccountSaasApp) GetType() NamedObjectType {
	if o == nil || IsNil(o.Type) {
		var ret NamedObjectType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectorIndividualSaaSAppAccountSaasApp) GetTypeOk() (*NamedObjectType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SelectorIndividualSaaSAppAccountSaasApp) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given NamedObjectType and assigns it to the Type field.
func (o *SelectorIndividualSaaSAppAccountSaasApp) SetType(v NamedObjectType) *SelectorIndividualSaaSAppAccountSaasApp {
	o.Type = &v
	return o
}

// GetMissing returns the Missing field value if set, zero value otherwise.
func (o *SelectorIndividualSaaSAppAccountSaasApp) GetMissing() bool {
	if o == nil || IsNil(o.Missing) {
		var ret bool
		return ret
	}
	return *o.Missing
}

// GetMissingOk returns a tuple with the Missing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectorIndividualSaaSAppAccountSaasApp) GetMissingOk() (*bool, bool) {
	if o == nil || IsNil(o.Missing) {
		return nil, false
	}
	return o.Missing, true
}

// HasMissing returns a boolean if a field has been set.
func (o *SelectorIndividualSaaSAppAccountSaasApp) HasMissing() bool {
	if o != nil && !IsNil(o.Missing) {
		return true
	}

	return false
}

// SetMissing gets a reference to the given bool and assigns it to the Missing field.
func (o *SelectorIndividualSaaSAppAccountSaasApp) SetMissing(v bool) *SelectorIndividualSaaSAppAccountSaasApp {
	o.Missing = &v
	return o
}

func (o SelectorIndividualSaaSAppAccountSaasApp) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SelectorIndividualSaaSAppAccountSaasApp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Missing) {
		toSerialize["missing"] = o.Missing
	}
	return toSerialize, nil
}

type NullableSelectorIndividualSaaSAppAccountSaasApp struct {
	value *SelectorIndividualSaaSAppAccountSaasApp
	isSet bool
}

func (v NullableSelectorIndividualSaaSAppAccountSaasApp) Get() *SelectorIndividualSaaSAppAccountSaasApp {
	return v.value
}

func (v *NullableSelectorIndividualSaaSAppAccountSaasApp) Set(val *SelectorIndividualSaaSAppAccountSaasApp) {
	v.value = val
	v.isSet = true
}

func (v NullableSelectorIndividualSaaSAppAccountSaasApp) IsSet() bool {
	return v.isSet
}

func (v *NullableSelectorIndividualSaaSAppAccountSaasApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelectorIndividualSaaSAppAccountSaasApp(val *SelectorIndividualSaaSAppAccountSaasApp) *NullableSelectorIndividualSaaSAppAccountSaasApp {
	return &NullableSelectorIndividualSaaSAppAccountSaasApp{value: val, isSet: true}
}

func (v NullableSelectorIndividualSaaSAppAccountSaasApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelectorIndividualSaaSAppAccountSaasApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}