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

// checks if the Service type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Service{}

// Service struct for Service
type Service struct {
	// The UUID of the Service
	Id string `json:"id"`
	// The UUID of the Server associated with the Service
	ServerId string `json:"server_id"`
	// The UID of the Server User associated with the Service User
	ServerUid string `json:"server_uid"`
	// The UUID of the Service User associated with the Service
	ServiceUserId string `json:"service_user_id"`
}

// NewService instantiates a new Service object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewService(id string, serverId string, serverUid string, serviceUserId string) *Service {
	this := Service{}
	this.Id = id
	this.ServerId = serverId
	this.ServerUid = serverUid
	this.ServiceUserId = serviceUserId
	return &this
}

// NewServiceWithDefaults instantiates a new Service object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceWithDefaults() *Service {
	this := Service{}
	return &this
}

// GetId returns the Id field value
func (o *Service) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Service) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Service) SetId(v string) *Service {
	o.Id = v
	return o
}

// GetServerId returns the ServerId field value
func (o *Service) GetServerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value
// and a boolean to check if the value has been set.
func (o *Service) GetServerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerId, true
}

// SetServerId sets field value
func (o *Service) SetServerId(v string) *Service {
	o.ServerId = v
	return o
}

// GetServerUid returns the ServerUid field value
func (o *Service) GetServerUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerUid
}

// GetServerUidOk returns a tuple with the ServerUid field value
// and a boolean to check if the value has been set.
func (o *Service) GetServerUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerUid, true
}

// SetServerUid sets field value
func (o *Service) SetServerUid(v string) *Service {
	o.ServerUid = v
	return o
}

// GetServiceUserId returns the ServiceUserId field value
func (o *Service) GetServiceUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceUserId
}

// GetServiceUserIdOk returns a tuple with the ServiceUserId field value
// and a boolean to check if the value has been set.
func (o *Service) GetServiceUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceUserId, true
}

// SetServiceUserId sets field value
func (o *Service) SetServiceUserId(v string) *Service {
	o.ServiceUserId = v
	return o
}

func (o Service) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Service) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["server_id"] = o.ServerId
	toSerialize["server_uid"] = o.ServerUid
	toSerialize["service_user_id"] = o.ServiceUserId
	return toSerialize, nil
}

type NullableService struct {
	value *Service
	isSet bool
}

func (v NullableService) Get() *Service {
	return v.value
}

func (v *NullableService) Set(val *Service) {
	v.value = val
	v.isSet = true
}

func (v NullableService) IsSet() bool {
	return v.isSet
}

func (v *NullableService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableService(val *Service) *NullableService {
	return &NullableService{value: val, isSet: true}
}

func (v NullableService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}