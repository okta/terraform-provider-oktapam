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

// checks if the ADServerSync type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ADServerSync{}

// ADServerSync struct for ADServerSync
type ADServerSync struct {
	// The UUID of the server sync job
	Id *string `json:"id,omitempty"`
	// The name of the server sync job
	Name *string `json:"name,omitempty"`
	// The AD attribute that defines the hostname for a server. This is used to identify discovered servers.
	HostNameAttribute *string `json:"host_name_attribute,omitempty"`
	// The AD attribute that defines an IP address or DNS name for a server. This is used by the Gateway to connect to discovered servers.
	AccessAddressAttribute *string `json:"access_address_attribute,omitempty"`
	// The AD attribute that defines the operating system of discovered servers
	OsAttribute *string `json:"os_attribute,omitempty"`
	// The AD attribute that defines the bastion host for a server. Clients use this bastion to tunnel traffic to discovered servers.
	BastionAttribute           *string                                 `json:"bastion_attribute,omitempty"`
	AdditionalAttributeMapping *ADServerSyncAdditionalAttributeMapping `json:"additional_attribute_mapping,omitempty"`
	// If `true`, the AD attribute is identified as a Globally Unique Identifier (GUID)
	IsGuid *bool `json:"is_guid,omitempty"`
	// Indicates how often the server sync job runs. Possible values: `1`, `6`, `12`, `24`.
	Frequency *string `json:"frequency,omitempty"`
	// A UTC timestamp that indicates the hour range when the server sync job runs. Only used if `frequency`is set to 24.
	StartHourUtc *string `json:"start_hour_utc,omitempty"`
	// If `true`, enables the server sync job
	IsActive *bool `json:"is_active,omitempty"`
	// The AD attribute that defines alternative hostnames or DNS entries. These are used to resolve discovered servers.
	AltNamesAttributes *string                      `json:"alt_names_attributes,omitempty"`
	RuleAssignments    *ADServerSyncRuleAssignments `json:"rule_assignments,omitempty"`
	// The results from the most recent server sync job
	RecentTaskResult *string `json:"RecentTaskResult,omitempty"`
}

// NewADServerSync instantiates a new ADServerSync object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewADServerSync() *ADServerSync {
	this := ADServerSync{}
	return &this
}

// NewADServerSyncWithDefaults instantiates a new ADServerSync object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewADServerSyncWithDefaults() *ADServerSync {
	this := ADServerSync{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ADServerSync) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADServerSync) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ADServerSync) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ADServerSync) SetId(v string) *ADServerSync {
	o.Id = &v
	return o
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ADServerSync) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADServerSync) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ADServerSync) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ADServerSync) SetName(v string) *ADServerSync {
	o.Name = &v
	return o
}

// GetHostNameAttribute returns the HostNameAttribute field value if set, zero value otherwise.
func (o *ADServerSync) GetHostNameAttribute() string {
	if o == nil || IsNil(o.HostNameAttribute) {
		var ret string
		return ret
	}
	return *o.HostNameAttribute
}

// GetHostNameAttributeOk returns a tuple with the HostNameAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADServerSync) GetHostNameAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.HostNameAttribute) {
		return nil, false
	}
	return o.HostNameAttribute, true
}

// HasHostNameAttribute returns a boolean if a field has been set.
func (o *ADServerSync) HasHostNameAttribute() bool {
	if o != nil && !IsNil(o.HostNameAttribute) {
		return true
	}

	return false
}

// SetHostNameAttribute gets a reference to the given string and assigns it to the HostNameAttribute field.
func (o *ADServerSync) SetHostNameAttribute(v string) *ADServerSync {
	o.HostNameAttribute = &v
	return o
}

// GetAccessAddressAttribute returns the AccessAddressAttribute field value if set, zero value otherwise.
func (o *ADServerSync) GetAccessAddressAttribute() string {
	if o == nil || IsNil(o.AccessAddressAttribute) {
		var ret string
		return ret
	}
	return *o.AccessAddressAttribute
}

// GetAccessAddressAttributeOk returns a tuple with the AccessAddressAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADServerSync) GetAccessAddressAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.AccessAddressAttribute) {
		return nil, false
	}
	return o.AccessAddressAttribute, true
}

// HasAccessAddressAttribute returns a boolean if a field has been set.
func (o *ADServerSync) HasAccessAddressAttribute() bool {
	if o != nil && !IsNil(o.AccessAddressAttribute) {
		return true
	}

	return false
}

// SetAccessAddressAttribute gets a reference to the given string and assigns it to the AccessAddressAttribute field.
func (o *ADServerSync) SetAccessAddressAttribute(v string) *ADServerSync {
	o.AccessAddressAttribute = &v
	return o
}

// GetOsAttribute returns the OsAttribute field value if set, zero value otherwise.
func (o *ADServerSync) GetOsAttribute() string {
	if o == nil || IsNil(o.OsAttribute) {
		var ret string
		return ret
	}
	return *o.OsAttribute
}

// GetOsAttributeOk returns a tuple with the OsAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADServerSync) GetOsAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.OsAttribute) {
		return nil, false
	}
	return o.OsAttribute, true
}

// HasOsAttribute returns a boolean if a field has been set.
func (o *ADServerSync) HasOsAttribute() bool {
	if o != nil && !IsNil(o.OsAttribute) {
		return true
	}

	return false
}

// SetOsAttribute gets a reference to the given string and assigns it to the OsAttribute field.
func (o *ADServerSync) SetOsAttribute(v string) *ADServerSync {
	o.OsAttribute = &v
	return o
}

// GetBastionAttribute returns the BastionAttribute field value if set, zero value otherwise.
func (o *ADServerSync) GetBastionAttribute() string {
	if o == nil || IsNil(o.BastionAttribute) {
		var ret string
		return ret
	}
	return *o.BastionAttribute
}

// GetBastionAttributeOk returns a tuple with the BastionAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADServerSync) GetBastionAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.BastionAttribute) {
		return nil, false
	}
	return o.BastionAttribute, true
}

// HasBastionAttribute returns a boolean if a field has been set.
func (o *ADServerSync) HasBastionAttribute() bool {
	if o != nil && !IsNil(o.BastionAttribute) {
		return true
	}

	return false
}

// SetBastionAttribute gets a reference to the given string and assigns it to the BastionAttribute field.
func (o *ADServerSync) SetBastionAttribute(v string) *ADServerSync {
	o.BastionAttribute = &v
	return o
}

// GetAdditionalAttributeMapping returns the AdditionalAttributeMapping field value if set, zero value otherwise.
func (o *ADServerSync) GetAdditionalAttributeMapping() ADServerSyncAdditionalAttributeMapping {
	if o == nil || IsNil(o.AdditionalAttributeMapping) {
		var ret ADServerSyncAdditionalAttributeMapping
		return ret
	}
	return *o.AdditionalAttributeMapping
}

// GetAdditionalAttributeMappingOk returns a tuple with the AdditionalAttributeMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADServerSync) GetAdditionalAttributeMappingOk() (*ADServerSyncAdditionalAttributeMapping, bool) {
	if o == nil || IsNil(o.AdditionalAttributeMapping) {
		return nil, false
	}
	return o.AdditionalAttributeMapping, true
}

// HasAdditionalAttributeMapping returns a boolean if a field has been set.
func (o *ADServerSync) HasAdditionalAttributeMapping() bool {
	if o != nil && !IsNil(o.AdditionalAttributeMapping) {
		return true
	}

	return false
}

// SetAdditionalAttributeMapping gets a reference to the given ADServerSyncAdditionalAttributeMapping and assigns it to the AdditionalAttributeMapping field.
func (o *ADServerSync) SetAdditionalAttributeMapping(v ADServerSyncAdditionalAttributeMapping) *ADServerSync {
	o.AdditionalAttributeMapping = &v
	return o
}

// GetIsGuid returns the IsGuid field value if set, zero value otherwise.
func (o *ADServerSync) GetIsGuid() bool {
	if o == nil || IsNil(o.IsGuid) {
		var ret bool
		return ret
	}
	return *o.IsGuid
}

// GetIsGuidOk returns a tuple with the IsGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADServerSync) GetIsGuidOk() (*bool, bool) {
	if o == nil || IsNil(o.IsGuid) {
		return nil, false
	}
	return o.IsGuid, true
}

// HasIsGuid returns a boolean if a field has been set.
func (o *ADServerSync) HasIsGuid() bool {
	if o != nil && !IsNil(o.IsGuid) {
		return true
	}

	return false
}

// SetIsGuid gets a reference to the given bool and assigns it to the IsGuid field.
func (o *ADServerSync) SetIsGuid(v bool) *ADServerSync {
	o.IsGuid = &v
	return o
}

// GetFrequency returns the Frequency field value if set, zero value otherwise.
func (o *ADServerSync) GetFrequency() string {
	if o == nil || IsNil(o.Frequency) {
		var ret string
		return ret
	}
	return *o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADServerSync) GetFrequencyOk() (*string, bool) {
	if o == nil || IsNil(o.Frequency) {
		return nil, false
	}
	return o.Frequency, true
}

// HasFrequency returns a boolean if a field has been set.
func (o *ADServerSync) HasFrequency() bool {
	if o != nil && !IsNil(o.Frequency) {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given string and assigns it to the Frequency field.
func (o *ADServerSync) SetFrequency(v string) *ADServerSync {
	o.Frequency = &v
	return o
}

// GetStartHourUtc returns the StartHourUtc field value if set, zero value otherwise.
func (o *ADServerSync) GetStartHourUtc() string {
	if o == nil || IsNil(o.StartHourUtc) {
		var ret string
		return ret
	}
	return *o.StartHourUtc
}

// GetStartHourUtcOk returns a tuple with the StartHourUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADServerSync) GetStartHourUtcOk() (*string, bool) {
	if o == nil || IsNil(o.StartHourUtc) {
		return nil, false
	}
	return o.StartHourUtc, true
}

// HasStartHourUtc returns a boolean if a field has been set.
func (o *ADServerSync) HasStartHourUtc() bool {
	if o != nil && !IsNil(o.StartHourUtc) {
		return true
	}

	return false
}

// SetStartHourUtc gets a reference to the given string and assigns it to the StartHourUtc field.
func (o *ADServerSync) SetStartHourUtc(v string) *ADServerSync {
	o.StartHourUtc = &v
	return o
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *ADServerSync) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADServerSync) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *ADServerSync) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *ADServerSync) SetIsActive(v bool) *ADServerSync {
	o.IsActive = &v
	return o
}

// GetAltNamesAttributes returns the AltNamesAttributes field value if set, zero value otherwise.
func (o *ADServerSync) GetAltNamesAttributes() string {
	if o == nil || IsNil(o.AltNamesAttributes) {
		var ret string
		return ret
	}
	return *o.AltNamesAttributes
}

// GetAltNamesAttributesOk returns a tuple with the AltNamesAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADServerSync) GetAltNamesAttributesOk() (*string, bool) {
	if o == nil || IsNil(o.AltNamesAttributes) {
		return nil, false
	}
	return o.AltNamesAttributes, true
}

// HasAltNamesAttributes returns a boolean if a field has been set.
func (o *ADServerSync) HasAltNamesAttributes() bool {
	if o != nil && !IsNil(o.AltNamesAttributes) {
		return true
	}

	return false
}

// SetAltNamesAttributes gets a reference to the given string and assigns it to the AltNamesAttributes field.
func (o *ADServerSync) SetAltNamesAttributes(v string) *ADServerSync {
	o.AltNamesAttributes = &v
	return o
}

// GetRuleAssignments returns the RuleAssignments field value if set, zero value otherwise.
func (o *ADServerSync) GetRuleAssignments() ADServerSyncRuleAssignments {
	if o == nil || IsNil(o.RuleAssignments) {
		var ret ADServerSyncRuleAssignments
		return ret
	}
	return *o.RuleAssignments
}

// GetRuleAssignmentsOk returns a tuple with the RuleAssignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADServerSync) GetRuleAssignmentsOk() (*ADServerSyncRuleAssignments, bool) {
	if o == nil || IsNil(o.RuleAssignments) {
		return nil, false
	}
	return o.RuleAssignments, true
}

// HasRuleAssignments returns a boolean if a field has been set.
func (o *ADServerSync) HasRuleAssignments() bool {
	if o != nil && !IsNil(o.RuleAssignments) {
		return true
	}

	return false
}

// SetRuleAssignments gets a reference to the given ADServerSyncRuleAssignments and assigns it to the RuleAssignments field.
func (o *ADServerSync) SetRuleAssignments(v ADServerSyncRuleAssignments) *ADServerSync {
	o.RuleAssignments = &v
	return o
}

// GetRecentTaskResult returns the RecentTaskResult field value if set, zero value otherwise.
func (o *ADServerSync) GetRecentTaskResult() string {
	if o == nil || IsNil(o.RecentTaskResult) {
		var ret string
		return ret
	}
	return *o.RecentTaskResult
}

// GetRecentTaskResultOk returns a tuple with the RecentTaskResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADServerSync) GetRecentTaskResultOk() (*string, bool) {
	if o == nil || IsNil(o.RecentTaskResult) {
		return nil, false
	}
	return o.RecentTaskResult, true
}

// HasRecentTaskResult returns a boolean if a field has been set.
func (o *ADServerSync) HasRecentTaskResult() bool {
	if o != nil && !IsNil(o.RecentTaskResult) {
		return true
	}

	return false
}

// SetRecentTaskResult gets a reference to the given string and assigns it to the RecentTaskResult field.
func (o *ADServerSync) SetRecentTaskResult(v string) *ADServerSync {
	o.RecentTaskResult = &v
	return o
}

func (o ADServerSync) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ADServerSync) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.HostNameAttribute) {
		toSerialize["host_name_attribute"] = o.HostNameAttribute
	}
	if !IsNil(o.AccessAddressAttribute) {
		toSerialize["access_address_attribute"] = o.AccessAddressAttribute
	}
	if !IsNil(o.OsAttribute) {
		toSerialize["os_attribute"] = o.OsAttribute
	}
	if !IsNil(o.BastionAttribute) {
		toSerialize["bastion_attribute"] = o.BastionAttribute
	}
	if !IsNil(o.AdditionalAttributeMapping) {
		toSerialize["additional_attribute_mapping"] = o.AdditionalAttributeMapping
	}
	if !IsNil(o.IsGuid) {
		toSerialize["is_guid"] = o.IsGuid
	}
	if !IsNil(o.Frequency) {
		toSerialize["frequency"] = o.Frequency
	}
	if !IsNil(o.StartHourUtc) {
		toSerialize["start_hour_utc"] = o.StartHourUtc
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	if !IsNil(o.AltNamesAttributes) {
		toSerialize["alt_names_attributes"] = o.AltNamesAttributes
	}
	if !IsNil(o.RuleAssignments) {
		toSerialize["rule_assignments"] = o.RuleAssignments
	}
	if !IsNil(o.RecentTaskResult) {
		toSerialize["RecentTaskResult"] = o.RecentTaskResult
	}
	return toSerialize, nil
}

type NullableADServerSync struct {
	value *ADServerSync
	isSet bool
}

func (v NullableADServerSync) Get() *ADServerSync {
	return v.value
}

func (v *NullableADServerSync) Set(val *ADServerSync) {
	v.value = val
	v.isSet = true
}

func (v NullableADServerSync) IsSet() bool {
	return v.isSet
}

func (v *NullableADServerSync) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableADServerSync(val *ADServerSync) *NullableADServerSync {
	return &NullableADServerSync{value: val, isSet: true}
}

func (v NullableADServerSync) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableADServerSync) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
