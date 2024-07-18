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
	"time"
)

// checks if the CheckedOutResourceByUserDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckedOutResourceByUserDetails{}

// CheckedOutResourceByUserDetails struct for CheckedOutResourceByUserDetails
type CheckedOutResourceByUserDetails struct {
	// The UUID of the resource
	ResourceId   string               `json:"resource_id"`
	ResourceType CheckoutResourceType `json:"resource_type"`
	// The timestamp when the resource was checked out
	CheckoutAt time.Time `json:"checkout_at"`
	// The timestamp when the resource lease expires
	CheckoutExpiryAt time.Time `json:"checkout_expiry_at"`
	// The UUID of the User or system responsible for the last checkin
	CheckinBy *string `json:"checkin_by,omitempty"`
	// The timestamp when the resource checkin process started
	CheckinStartAt *time.Time   `json:"checkin_start_at,omitempty"`
	CheckinType    *CheckinType `json:"checkin_type,omitempty"`
	// The name of the resource. Only returned if a `resource_type` query was specified.
	ResourceName    *string                                         `json:"resource_name,omitempty"`
	ResourceDetails *CheckedOutResourceByUserDetailsResourceDetails `json:"resource_details,omitempty"`
}

// NewCheckedOutResourceByUserDetails instantiates a new CheckedOutResourceByUserDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckedOutResourceByUserDetails(resourceId string, resourceType CheckoutResourceType, checkoutAt time.Time, checkoutExpiryAt time.Time) *CheckedOutResourceByUserDetails {
	this := CheckedOutResourceByUserDetails{}
	this.ResourceId = resourceId
	this.ResourceType = resourceType
	this.CheckoutAt = checkoutAt
	this.CheckoutExpiryAt = checkoutExpiryAt
	return &this
}

// NewCheckedOutResourceByUserDetailsWithDefaults instantiates a new CheckedOutResourceByUserDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckedOutResourceByUserDetailsWithDefaults() *CheckedOutResourceByUserDetails {
	this := CheckedOutResourceByUserDetails{}
	return &this
}

// GetResourceId returns the ResourceId field value
func (o *CheckedOutResourceByUserDetails) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *CheckedOutResourceByUserDetails) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *CheckedOutResourceByUserDetails) SetResourceId(v string) *CheckedOutResourceByUserDetails {
	o.ResourceId = v
	return o
}

// GetResourceType returns the ResourceType field value
func (o *CheckedOutResourceByUserDetails) GetResourceType() CheckoutResourceType {
	if o == nil {
		var ret CheckoutResourceType
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *CheckedOutResourceByUserDetails) GetResourceTypeOk() (*CheckoutResourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *CheckedOutResourceByUserDetails) SetResourceType(v CheckoutResourceType) *CheckedOutResourceByUserDetails {
	o.ResourceType = v
	return o
}

// GetCheckoutAt returns the CheckoutAt field value
func (o *CheckedOutResourceByUserDetails) GetCheckoutAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CheckoutAt
}

// GetCheckoutAtOk returns a tuple with the CheckoutAt field value
// and a boolean to check if the value has been set.
func (o *CheckedOutResourceByUserDetails) GetCheckoutAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CheckoutAt, true
}

// SetCheckoutAt sets field value
func (o *CheckedOutResourceByUserDetails) SetCheckoutAt(v time.Time) *CheckedOutResourceByUserDetails {
	o.CheckoutAt = v
	return o
}

// GetCheckoutExpiryAt returns the CheckoutExpiryAt field value
func (o *CheckedOutResourceByUserDetails) GetCheckoutExpiryAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CheckoutExpiryAt
}

// GetCheckoutExpiryAtOk returns a tuple with the CheckoutExpiryAt field value
// and a boolean to check if the value has been set.
func (o *CheckedOutResourceByUserDetails) GetCheckoutExpiryAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CheckoutExpiryAt, true
}

// SetCheckoutExpiryAt sets field value
func (o *CheckedOutResourceByUserDetails) SetCheckoutExpiryAt(v time.Time) *CheckedOutResourceByUserDetails {
	o.CheckoutExpiryAt = v
	return o
}

// GetCheckinBy returns the CheckinBy field value if set, zero value otherwise.
func (o *CheckedOutResourceByUserDetails) GetCheckinBy() string {
	if o == nil || IsNil(o.CheckinBy) {
		var ret string
		return ret
	}
	return *o.CheckinBy
}

// GetCheckinByOk returns a tuple with the CheckinBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckedOutResourceByUserDetails) GetCheckinByOk() (*string, bool) {
	if o == nil || IsNil(o.CheckinBy) {
		return nil, false
	}
	return o.CheckinBy, true
}

// HasCheckinBy returns a boolean if a field has been set.
func (o *CheckedOutResourceByUserDetails) HasCheckinBy() bool {
	if o != nil && !IsNil(o.CheckinBy) {
		return true
	}

	return false
}

// SetCheckinBy gets a reference to the given string and assigns it to the CheckinBy field.
func (o *CheckedOutResourceByUserDetails) SetCheckinBy(v string) *CheckedOutResourceByUserDetails {
	o.CheckinBy = &v
	return o
}

// GetCheckinStartAt returns the CheckinStartAt field value if set, zero value otherwise.
func (o *CheckedOutResourceByUserDetails) GetCheckinStartAt() time.Time {
	if o == nil || IsNil(o.CheckinStartAt) {
		var ret time.Time
		return ret
	}
	return *o.CheckinStartAt
}

// GetCheckinStartAtOk returns a tuple with the CheckinStartAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckedOutResourceByUserDetails) GetCheckinStartAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CheckinStartAt) {
		return nil, false
	}
	return o.CheckinStartAt, true
}

// HasCheckinStartAt returns a boolean if a field has been set.
func (o *CheckedOutResourceByUserDetails) HasCheckinStartAt() bool {
	if o != nil && !IsNil(o.CheckinStartAt) {
		return true
	}

	return false
}

// SetCheckinStartAt gets a reference to the given time.Time and assigns it to the CheckinStartAt field.
func (o *CheckedOutResourceByUserDetails) SetCheckinStartAt(v time.Time) *CheckedOutResourceByUserDetails {
	o.CheckinStartAt = &v
	return o
}

// GetCheckinType returns the CheckinType field value if set, zero value otherwise.
func (o *CheckedOutResourceByUserDetails) GetCheckinType() CheckinType {
	if o == nil || IsNil(o.CheckinType) {
		var ret CheckinType
		return ret
	}
	return *o.CheckinType
}

// GetCheckinTypeOk returns a tuple with the CheckinType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckedOutResourceByUserDetails) GetCheckinTypeOk() (*CheckinType, bool) {
	if o == nil || IsNil(o.CheckinType) {
		return nil, false
	}
	return o.CheckinType, true
}

// HasCheckinType returns a boolean if a field has been set.
func (o *CheckedOutResourceByUserDetails) HasCheckinType() bool {
	if o != nil && !IsNil(o.CheckinType) {
		return true
	}

	return false
}

// SetCheckinType gets a reference to the given CheckinType and assigns it to the CheckinType field.
func (o *CheckedOutResourceByUserDetails) SetCheckinType(v CheckinType) *CheckedOutResourceByUserDetails {
	o.CheckinType = &v
	return o
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise.
func (o *CheckedOutResourceByUserDetails) GetResourceName() string {
	if o == nil || IsNil(o.ResourceName) {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckedOutResourceByUserDetails) GetResourceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceName) {
		return nil, false
	}
	return o.ResourceName, true
}

// HasResourceName returns a boolean if a field has been set.
func (o *CheckedOutResourceByUserDetails) HasResourceName() bool {
	if o != nil && !IsNil(o.ResourceName) {
		return true
	}

	return false
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *CheckedOutResourceByUserDetails) SetResourceName(v string) *CheckedOutResourceByUserDetails {
	o.ResourceName = &v
	return o
}

// GetResourceDetails returns the ResourceDetails field value if set, zero value otherwise.
func (o *CheckedOutResourceByUserDetails) GetResourceDetails() CheckedOutResourceByUserDetailsResourceDetails {
	if o == nil || IsNil(o.ResourceDetails) {
		var ret CheckedOutResourceByUserDetailsResourceDetails
		return ret
	}
	return *o.ResourceDetails
}

// GetResourceDetailsOk returns a tuple with the ResourceDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckedOutResourceByUserDetails) GetResourceDetailsOk() (*CheckedOutResourceByUserDetailsResourceDetails, bool) {
	if o == nil || IsNil(o.ResourceDetails) {
		return nil, false
	}
	return o.ResourceDetails, true
}

// HasResourceDetails returns a boolean if a field has been set.
func (o *CheckedOutResourceByUserDetails) HasResourceDetails() bool {
	if o != nil && !IsNil(o.ResourceDetails) {
		return true
	}

	return false
}

// SetResourceDetails gets a reference to the given CheckedOutResourceByUserDetailsResourceDetails and assigns it to the ResourceDetails field.
func (o *CheckedOutResourceByUserDetails) SetResourceDetails(v CheckedOutResourceByUserDetailsResourceDetails) *CheckedOutResourceByUserDetails {
	o.ResourceDetails = &v
	return o
}

func (o CheckedOutResourceByUserDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckedOutResourceByUserDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resource_id"] = o.ResourceId
	toSerialize["resource_type"] = o.ResourceType
	toSerialize["checkout_at"] = o.CheckoutAt
	toSerialize["checkout_expiry_at"] = o.CheckoutExpiryAt
	if !IsNil(o.CheckinBy) {
		toSerialize["checkin_by"] = o.CheckinBy
	}
	if !IsNil(o.CheckinStartAt) {
		toSerialize["checkin_start_at"] = o.CheckinStartAt
	}
	if !IsNil(o.CheckinType) {
		toSerialize["checkin_type"] = o.CheckinType
	}
	if !IsNil(o.ResourceName) {
		toSerialize["resource_name"] = o.ResourceName
	}
	if !IsNil(o.ResourceDetails) {
		toSerialize["resource_details"] = o.ResourceDetails
	}
	return toSerialize, nil
}

type NullableCheckedOutResourceByUserDetails struct {
	value *CheckedOutResourceByUserDetails
	isSet bool
}

func (v NullableCheckedOutResourceByUserDetails) Get() *CheckedOutResourceByUserDetails {
	return v.value
}

func (v *NullableCheckedOutResourceByUserDetails) Set(val *CheckedOutResourceByUserDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckedOutResourceByUserDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckedOutResourceByUserDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckedOutResourceByUserDetails(val *CheckedOutResourceByUserDetails) *NullableCheckedOutResourceByUserDetails {
	return &NullableCheckedOutResourceByUserDetails{value: val, isSet: true}
}

func (v NullableCheckedOutResourceByUserDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckedOutResourceByUserDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
