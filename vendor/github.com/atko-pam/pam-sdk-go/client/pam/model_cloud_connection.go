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

// checks if the CloudConnection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudConnection{}

// CloudConnection struct for CloudConnection
type CloudConnection struct {
	// The UUID of the Cloud Connection
	Id *string `json:"id,omitempty"`
	// The name of the Cloud Connection
	Name                   *string                  `json:"name,omitempty"`
	Provider               *CloudConnectionProvider `json:"provider,omitempty"`
	CloudConnectionDetails *CloudConnectionDetails  `json:"cloud_connection_details,omitempty"`
}

// NewCloudConnection instantiates a new CloudConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudConnection() *CloudConnection {
	this := CloudConnection{}
	return &this
}

// NewCloudConnectionWithDefaults instantiates a new CloudConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudConnectionWithDefaults() *CloudConnection {
	this := CloudConnection{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CloudConnection) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudConnection) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CloudConnection) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CloudConnection) SetId(v string) *CloudConnection {
	o.Id = &v
	return o
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CloudConnection) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudConnection) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CloudConnection) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CloudConnection) SetName(v string) *CloudConnection {
	o.Name = &v
	return o
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *CloudConnection) GetProvider() CloudConnectionProvider {
	if o == nil || IsNil(o.Provider) {
		var ret CloudConnectionProvider
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudConnection) GetProviderOk() (*CloudConnectionProvider, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *CloudConnection) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given CloudConnectionProvider and assigns it to the Provider field.
func (o *CloudConnection) SetProvider(v CloudConnectionProvider) *CloudConnection {
	o.Provider = &v
	return o
}

// GetCloudConnectionDetails returns the CloudConnectionDetails field value if set, zero value otherwise.
func (o *CloudConnection) GetCloudConnectionDetails() CloudConnectionDetails {
	if o == nil || IsNil(o.CloudConnectionDetails) {
		var ret CloudConnectionDetails
		return ret
	}
	return *o.CloudConnectionDetails
}

// GetCloudConnectionDetailsOk returns a tuple with the CloudConnectionDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudConnection) GetCloudConnectionDetailsOk() (*CloudConnectionDetails, bool) {
	if o == nil || IsNil(o.CloudConnectionDetails) {
		return nil, false
	}
	return o.CloudConnectionDetails, true
}

// HasCloudConnectionDetails returns a boolean if a field has been set.
func (o *CloudConnection) HasCloudConnectionDetails() bool {
	if o != nil && !IsNil(o.CloudConnectionDetails) {
		return true
	}

	return false
}

// SetCloudConnectionDetails gets a reference to the given CloudConnectionDetails and assigns it to the CloudConnectionDetails field.
func (o *CloudConnection) SetCloudConnectionDetails(v CloudConnectionDetails) *CloudConnection {
	o.CloudConnectionDetails = &v
	return o
}

func (o CloudConnection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudConnection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !IsNil(o.CloudConnectionDetails) {
		toSerialize["cloud_connection_details"] = o.CloudConnectionDetails
	}
	return toSerialize, nil
}

type NullableCloudConnection struct {
	value *CloudConnection
	isSet bool
}

func (v NullableCloudConnection) Get() *CloudConnection {
	return v.value
}

func (v *NullableCloudConnection) Set(val *CloudConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudConnection(val *CloudConnection) *NullableCloudConnection {
	return &NullableCloudConnection{value: val, isSet: true}
}

func (v NullableCloudConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
