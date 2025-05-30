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

// checks if the CloudEntitlementResourcesDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudEntitlementResourcesDetails{}

// CloudEntitlementResourcesDetails Details of the associated resource
type CloudEntitlementResourcesDetails struct {
	// The name of the resource
	Name *string `json:"name,omitempty"`
	// An AWS account ID
	AccountId string `json:"account_id"`
	// The specific region where the resource is stored
	Region string `json:"region"`
	// The AWS UUID of the resource
	AwsId string `json:"aws_id"`
	// An AWS role ARN associated with the resource
	Arn *string `json:"arn,omitempty"`
	// The UUID of the resource
	ResourceId string `json:"resource_id"`
	// The type of resource returned by the job run. Currently only returns `rds`.
	ResourceType string                                     `json:"resource_type"`
	OrgDetails   CloudEntitlementResourcesDetailsOrgDetails `json:"org_details"`
	// One or more resources nested within the parent resource. Not all resources include children.
	Children []CloudEntitlementResourcesDetailsChildrenInner `json:"children,omitempty"`
}

// NewCloudEntitlementResourcesDetails instantiates a new CloudEntitlementResourcesDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudEntitlementResourcesDetails(accountId string, region string, awsId string, resourceId string, resourceType string, orgDetails CloudEntitlementResourcesDetailsOrgDetails) *CloudEntitlementResourcesDetails {
	this := CloudEntitlementResourcesDetails{}
	this.AccountId = accountId
	this.Region = region
	this.AwsId = awsId
	this.ResourceId = resourceId
	this.ResourceType = resourceType
	this.OrgDetails = orgDetails
	return &this
}

// NewCloudEntitlementResourcesDetailsWithDefaults instantiates a new CloudEntitlementResourcesDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudEntitlementResourcesDetailsWithDefaults() *CloudEntitlementResourcesDetails {
	this := CloudEntitlementResourcesDetails{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CloudEntitlementResourcesDetails) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudEntitlementResourcesDetails) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CloudEntitlementResourcesDetails) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CloudEntitlementResourcesDetails) SetName(v string) *CloudEntitlementResourcesDetails {
	o.Name = &v
	return o
}

// GetAccountId returns the AccountId field value
func (o *CloudEntitlementResourcesDetails) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *CloudEntitlementResourcesDetails) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *CloudEntitlementResourcesDetails) SetAccountId(v string) *CloudEntitlementResourcesDetails {
	o.AccountId = v
	return o
}

// GetRegion returns the Region field value
func (o *CloudEntitlementResourcesDetails) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *CloudEntitlementResourcesDetails) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *CloudEntitlementResourcesDetails) SetRegion(v string) *CloudEntitlementResourcesDetails {
	o.Region = v
	return o
}

// GetAwsId returns the AwsId field value
func (o *CloudEntitlementResourcesDetails) GetAwsId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AwsId
}

// GetAwsIdOk returns a tuple with the AwsId field value
// and a boolean to check if the value has been set.
func (o *CloudEntitlementResourcesDetails) GetAwsIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsId, true
}

// SetAwsId sets field value
func (o *CloudEntitlementResourcesDetails) SetAwsId(v string) *CloudEntitlementResourcesDetails {
	o.AwsId = v
	return o
}

// GetArn returns the Arn field value if set, zero value otherwise.
func (o *CloudEntitlementResourcesDetails) GetArn() string {
	if o == nil || IsNil(o.Arn) {
		var ret string
		return ret
	}
	return *o.Arn
}

// GetArnOk returns a tuple with the Arn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudEntitlementResourcesDetails) GetArnOk() (*string, bool) {
	if o == nil || IsNil(o.Arn) {
		return nil, false
	}
	return o.Arn, true
}

// HasArn returns a boolean if a field has been set.
func (o *CloudEntitlementResourcesDetails) HasArn() bool {
	if o != nil && !IsNil(o.Arn) {
		return true
	}

	return false
}

// SetArn gets a reference to the given string and assigns it to the Arn field.
func (o *CloudEntitlementResourcesDetails) SetArn(v string) *CloudEntitlementResourcesDetails {
	o.Arn = &v
	return o
}

// GetResourceId returns the ResourceId field value
func (o *CloudEntitlementResourcesDetails) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *CloudEntitlementResourcesDetails) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *CloudEntitlementResourcesDetails) SetResourceId(v string) *CloudEntitlementResourcesDetails {
	o.ResourceId = v
	return o
}

// GetResourceType returns the ResourceType field value
func (o *CloudEntitlementResourcesDetails) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *CloudEntitlementResourcesDetails) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *CloudEntitlementResourcesDetails) SetResourceType(v string) *CloudEntitlementResourcesDetails {
	o.ResourceType = v
	return o
}

// GetOrgDetails returns the OrgDetails field value
func (o *CloudEntitlementResourcesDetails) GetOrgDetails() CloudEntitlementResourcesDetailsOrgDetails {
	if o == nil {
		var ret CloudEntitlementResourcesDetailsOrgDetails
		return ret
	}

	return o.OrgDetails
}

// GetOrgDetailsOk returns a tuple with the OrgDetails field value
// and a boolean to check if the value has been set.
func (o *CloudEntitlementResourcesDetails) GetOrgDetailsOk() (*CloudEntitlementResourcesDetailsOrgDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgDetails, true
}

// SetOrgDetails sets field value
func (o *CloudEntitlementResourcesDetails) SetOrgDetails(v CloudEntitlementResourcesDetailsOrgDetails) *CloudEntitlementResourcesDetails {
	o.OrgDetails = v
	return o
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *CloudEntitlementResourcesDetails) GetChildren() []CloudEntitlementResourcesDetailsChildrenInner {
	if o == nil || IsNil(o.Children) {
		var ret []CloudEntitlementResourcesDetailsChildrenInner
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudEntitlementResourcesDetails) GetChildrenOk() ([]CloudEntitlementResourcesDetailsChildrenInner, bool) {
	if o == nil || IsNil(o.Children) {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *CloudEntitlementResourcesDetails) HasChildren() bool {
	if o != nil && !IsNil(o.Children) {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []CloudEntitlementResourcesDetailsChildrenInner and assigns it to the Children field.
func (o *CloudEntitlementResourcesDetails) SetChildren(v []CloudEntitlementResourcesDetailsChildrenInner) *CloudEntitlementResourcesDetails {
	o.Children = v
	return o
}

func (o CloudEntitlementResourcesDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudEntitlementResourcesDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["account_id"] = o.AccountId
	toSerialize["region"] = o.Region
	toSerialize["aws_id"] = o.AwsId
	if !IsNil(o.Arn) {
		toSerialize["arn"] = o.Arn
	}
	toSerialize["resource_id"] = o.ResourceId
	toSerialize["resource_type"] = o.ResourceType
	toSerialize["org_details"] = o.OrgDetails
	if !IsNil(o.Children) {
		toSerialize["children"] = o.Children
	}
	return toSerialize, nil
}

type NullableCloudEntitlementResourcesDetails struct {
	value *CloudEntitlementResourcesDetails
	isSet bool
}

func (v NullableCloudEntitlementResourcesDetails) Get() *CloudEntitlementResourcesDetails {
	return v.value
}

func (v *NullableCloudEntitlementResourcesDetails) Set(val *CloudEntitlementResourcesDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudEntitlementResourcesDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudEntitlementResourcesDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudEntitlementResourcesDetails(val *CloudEntitlementResourcesDetails) *NullableCloudEntitlementResourcesDetails {
	return &NullableCloudEntitlementResourcesDetails{value: val, isSet: true}
}

func (v NullableCloudEntitlementResourcesDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudEntitlementResourcesDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
