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

// checks if the ServiceWithUserName type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceWithUserName{}

// ServiceWithUserName struct for ServiceWithUserName
type ServiceWithUserName struct {
	// The UUID of the Service
	Id string `json:"id"`
	// The UUID of the Server associated with the Service
	ServerId string `json:"server_id"`
	// The UID of the Server User associated with the Service User
	ServerUid string `json:"server_uid"`
	// The UUID of the Service User associated with the Service
	ServiceUserId string `json:"service_user_id"`
	// The username of the Service User associated with the Service
	ServiceUserName string `json:"service_user_name"`
}

// NewServiceWithUserName instantiates a new ServiceWithUserName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceWithUserName(id string, serverId string, serverUid string, serviceUserId string, serviceUserName string) *ServiceWithUserName {
	this := ServiceWithUserName{}
	this.Id = id
	this.ServerId = serverId
	this.ServerUid = serverUid
	this.ServiceUserId = serviceUserId
	this.ServiceUserName = serviceUserName
	return &this
}

// NewServiceWithUserNameWithDefaults instantiates a new ServiceWithUserName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceWithUserNameWithDefaults() *ServiceWithUserName {
	this := ServiceWithUserName{}
	return &this
}

// GetId returns the Id field value
func (o *ServiceWithUserName) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ServiceWithUserName) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ServiceWithUserName) SetId(v string) *ServiceWithUserName {
	o.Id = v
	return o
}

// GetServerId returns the ServerId field value
func (o *ServiceWithUserName) GetServerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value
// and a boolean to check if the value has been set.
func (o *ServiceWithUserName) GetServerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerId, true
}

// SetServerId sets field value
func (o *ServiceWithUserName) SetServerId(v string) *ServiceWithUserName {
	o.ServerId = v
	return o
}

// GetServerUid returns the ServerUid field value
func (o *ServiceWithUserName) GetServerUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerUid
}

// GetServerUidOk returns a tuple with the ServerUid field value
// and a boolean to check if the value has been set.
func (o *ServiceWithUserName) GetServerUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerUid, true
}

// SetServerUid sets field value
func (o *ServiceWithUserName) SetServerUid(v string) *ServiceWithUserName {
	o.ServerUid = v
	return o
}

// GetServiceUserId returns the ServiceUserId field value
func (o *ServiceWithUserName) GetServiceUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceUserId
}

// GetServiceUserIdOk returns a tuple with the ServiceUserId field value
// and a boolean to check if the value has been set.
func (o *ServiceWithUserName) GetServiceUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceUserId, true
}

// SetServiceUserId sets field value
func (o *ServiceWithUserName) SetServiceUserId(v string) *ServiceWithUserName {
	o.ServiceUserId = v
	return o
}

// GetServiceUserName returns the ServiceUserName field value
func (o *ServiceWithUserName) GetServiceUserName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceUserName
}

// GetServiceUserNameOk returns a tuple with the ServiceUserName field value
// and a boolean to check if the value has been set.
func (o *ServiceWithUserName) GetServiceUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceUserName, true
}

// SetServiceUserName sets field value
func (o *ServiceWithUserName) SetServiceUserName(v string) *ServiceWithUserName {
	o.ServiceUserName = v
	return o
}

func (o ServiceWithUserName) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceWithUserName) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["server_id"] = o.ServerId
	toSerialize["server_uid"] = o.ServerUid
	toSerialize["service_user_id"] = o.ServiceUserId
	toSerialize["service_user_name"] = o.ServiceUserName
	return toSerialize, nil
}

type NullableServiceWithUserName struct {
	value *ServiceWithUserName
	isSet bool
}

func (v NullableServiceWithUserName) Get() *ServiceWithUserName {
	return v.value
}

func (v *NullableServiceWithUserName) Set(val *ServiceWithUserName) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceWithUserName) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceWithUserName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceWithUserName(val *ServiceWithUserName) *NullableServiceWithUserName {
	return &NullableServiceWithUserName{value: val, isSet: true}
}

func (v NullableServiceWithUserName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceWithUserName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
