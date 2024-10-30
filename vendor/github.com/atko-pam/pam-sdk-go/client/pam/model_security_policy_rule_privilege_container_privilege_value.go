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
	"fmt"
)

// SecurityPolicyRulePrivilegeContainerPrivilegeValue - An object that indicates whether the privilege type is allowed. The key must match the specified privilege type and the value must be set to `true`.
type SecurityPolicyRulePrivilegeContainerPrivilegeValue struct {
	SecurityPolicyPasswordCheckoutDatabasePrivilege *SecurityPolicyPasswordCheckoutDatabasePrivilege
	SecurityPolicyPasswordCheckoutRDPPrivilege      *SecurityPolicyPasswordCheckoutRDPPrivilege
	SecurityPolicyPasswordCheckoutSSHPrivilege      *SecurityPolicyPasswordCheckoutSSHPrivilege
	SecurityPolicyPrincipalAccountRDPPrivilege      *SecurityPolicyPrincipalAccountRDPPrivilege
	SecurityPolicyPrincipalAccountSSHPrivilege      *SecurityPolicyPrincipalAccountSSHPrivilege
	SecurityPolicyRevealPasswordPrivilege           *SecurityPolicyRevealPasswordPrivilege
	SecurityPolicySecretPrivilege                   *SecurityPolicySecretPrivilege
	SecurityPolicyUpdatePasswordPrivilege           *SecurityPolicyUpdatePasswordPrivilege
}

// SecurityPolicyPasswordCheckoutDatabasePrivilegeAsSecurityPolicyRulePrivilegeContainerPrivilegeValue is a convenience function that returns SecurityPolicyPasswordCheckoutDatabasePrivilege wrapped in SecurityPolicyRulePrivilegeContainerPrivilegeValue
func SecurityPolicyPasswordCheckoutDatabasePrivilegeAsSecurityPolicyRulePrivilegeContainerPrivilegeValue(v *SecurityPolicyPasswordCheckoutDatabasePrivilege) SecurityPolicyRulePrivilegeContainerPrivilegeValue {
	return SecurityPolicyRulePrivilegeContainerPrivilegeValue{
		SecurityPolicyPasswordCheckoutDatabasePrivilege: v,
	}
}

// SecurityPolicyPasswordCheckoutRDPPrivilegeAsSecurityPolicyRulePrivilegeContainerPrivilegeValue is a convenience function that returns SecurityPolicyPasswordCheckoutRDPPrivilege wrapped in SecurityPolicyRulePrivilegeContainerPrivilegeValue
func SecurityPolicyPasswordCheckoutRDPPrivilegeAsSecurityPolicyRulePrivilegeContainerPrivilegeValue(v *SecurityPolicyPasswordCheckoutRDPPrivilege) SecurityPolicyRulePrivilegeContainerPrivilegeValue {
	return SecurityPolicyRulePrivilegeContainerPrivilegeValue{
		SecurityPolicyPasswordCheckoutRDPPrivilege: v,
	}
}

// SecurityPolicyPasswordCheckoutSSHPrivilegeAsSecurityPolicyRulePrivilegeContainerPrivilegeValue is a convenience function that returns SecurityPolicyPasswordCheckoutSSHPrivilege wrapped in SecurityPolicyRulePrivilegeContainerPrivilegeValue
func SecurityPolicyPasswordCheckoutSSHPrivilegeAsSecurityPolicyRulePrivilegeContainerPrivilegeValue(v *SecurityPolicyPasswordCheckoutSSHPrivilege) SecurityPolicyRulePrivilegeContainerPrivilegeValue {
	return SecurityPolicyRulePrivilegeContainerPrivilegeValue{
		SecurityPolicyPasswordCheckoutSSHPrivilege: v,
	}
}

// SecurityPolicyPrincipalAccountRDPPrivilegeAsSecurityPolicyRulePrivilegeContainerPrivilegeValue is a convenience function that returns SecurityPolicyPrincipalAccountRDPPrivilege wrapped in SecurityPolicyRulePrivilegeContainerPrivilegeValue
func SecurityPolicyPrincipalAccountRDPPrivilegeAsSecurityPolicyRulePrivilegeContainerPrivilegeValue(v *SecurityPolicyPrincipalAccountRDPPrivilege) SecurityPolicyRulePrivilegeContainerPrivilegeValue {
	return SecurityPolicyRulePrivilegeContainerPrivilegeValue{
		SecurityPolicyPrincipalAccountRDPPrivilege: v,
	}
}

// SecurityPolicyPrincipalAccountSSHPrivilegeAsSecurityPolicyRulePrivilegeContainerPrivilegeValue is a convenience function that returns SecurityPolicyPrincipalAccountSSHPrivilege wrapped in SecurityPolicyRulePrivilegeContainerPrivilegeValue
func SecurityPolicyPrincipalAccountSSHPrivilegeAsSecurityPolicyRulePrivilegeContainerPrivilegeValue(v *SecurityPolicyPrincipalAccountSSHPrivilege) SecurityPolicyRulePrivilegeContainerPrivilegeValue {
	return SecurityPolicyRulePrivilegeContainerPrivilegeValue{
		SecurityPolicyPrincipalAccountSSHPrivilege: v,
	}
}

// SecurityPolicyRevealPasswordPrivilegeAsSecurityPolicyRulePrivilegeContainerPrivilegeValue is a convenience function that returns SecurityPolicyRevealPasswordPrivilege wrapped in SecurityPolicyRulePrivilegeContainerPrivilegeValue
func SecurityPolicyRevealPasswordPrivilegeAsSecurityPolicyRulePrivilegeContainerPrivilegeValue(v *SecurityPolicyRevealPasswordPrivilege) SecurityPolicyRulePrivilegeContainerPrivilegeValue {
	return SecurityPolicyRulePrivilegeContainerPrivilegeValue{
		SecurityPolicyRevealPasswordPrivilege: v,
	}
}

// SecurityPolicySecretPrivilegeAsSecurityPolicyRulePrivilegeContainerPrivilegeValue is a convenience function that returns SecurityPolicySecretPrivilege wrapped in SecurityPolicyRulePrivilegeContainerPrivilegeValue
func SecurityPolicySecretPrivilegeAsSecurityPolicyRulePrivilegeContainerPrivilegeValue(v *SecurityPolicySecretPrivilege) SecurityPolicyRulePrivilegeContainerPrivilegeValue {
	return SecurityPolicyRulePrivilegeContainerPrivilegeValue{
		SecurityPolicySecretPrivilege: v,
	}
}

// SecurityPolicyUpdatePasswordPrivilegeAsSecurityPolicyRulePrivilegeContainerPrivilegeValue is a convenience function that returns SecurityPolicyUpdatePasswordPrivilege wrapped in SecurityPolicyRulePrivilegeContainerPrivilegeValue
func SecurityPolicyUpdatePasswordPrivilegeAsSecurityPolicyRulePrivilegeContainerPrivilegeValue(v *SecurityPolicyUpdatePasswordPrivilege) SecurityPolicyRulePrivilegeContainerPrivilegeValue {
	return SecurityPolicyRulePrivilegeContainerPrivilegeValue{
		SecurityPolicyUpdatePasswordPrivilege: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SecurityPolicyRulePrivilegeContainerPrivilegeValue) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'SecurityPolicyPasswordCheckoutDatabasePrivilege'
	if jsonDict["_type"] == "SecurityPolicyPasswordCheckoutDatabasePrivilege" {
		// try to unmarshal JSON data into SecurityPolicyPasswordCheckoutDatabasePrivilege
		err = json.Unmarshal(data, &dst.SecurityPolicyPasswordCheckoutDatabasePrivilege)
		if err == nil {
			return nil // data stored in dst.SecurityPolicyPasswordCheckoutDatabasePrivilege, return on the first match
		} else {
			dst.SecurityPolicyPasswordCheckoutDatabasePrivilege = nil
			return fmt.Errorf("failed to unmarshal SecurityPolicyRulePrivilegeContainerPrivilegeValue as SecurityPolicyPasswordCheckoutDatabasePrivilege: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SecurityPolicyPasswordCheckoutRDPPrivilege'
	if jsonDict["_type"] == "SecurityPolicyPasswordCheckoutRDPPrivilege" {
		// try to unmarshal JSON data into SecurityPolicyPasswordCheckoutRDPPrivilege
		err = json.Unmarshal(data, &dst.SecurityPolicyPasswordCheckoutRDPPrivilege)
		if err == nil {
			return nil // data stored in dst.SecurityPolicyPasswordCheckoutRDPPrivilege, return on the first match
		} else {
			dst.SecurityPolicyPasswordCheckoutRDPPrivilege = nil
			return fmt.Errorf("failed to unmarshal SecurityPolicyRulePrivilegeContainerPrivilegeValue as SecurityPolicyPasswordCheckoutRDPPrivilege: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SecurityPolicyPasswordCheckoutSSHPrivilege'
	if jsonDict["_type"] == "SecurityPolicyPasswordCheckoutSSHPrivilege" {
		// try to unmarshal JSON data into SecurityPolicyPasswordCheckoutSSHPrivilege
		err = json.Unmarshal(data, &dst.SecurityPolicyPasswordCheckoutSSHPrivilege)
		if err == nil {
			return nil // data stored in dst.SecurityPolicyPasswordCheckoutSSHPrivilege, return on the first match
		} else {
			dst.SecurityPolicyPasswordCheckoutSSHPrivilege = nil
			return fmt.Errorf("failed to unmarshal SecurityPolicyRulePrivilegeContainerPrivilegeValue as SecurityPolicyPasswordCheckoutSSHPrivilege: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SecurityPolicyPrincipalAccountRDPPrivilege'
	if jsonDict["_type"] == "SecurityPolicyPrincipalAccountRDPPrivilege" {
		// try to unmarshal JSON data into SecurityPolicyPrincipalAccountRDPPrivilege
		err = json.Unmarshal(data, &dst.SecurityPolicyPrincipalAccountRDPPrivilege)
		if err == nil {
			return nil // data stored in dst.SecurityPolicyPrincipalAccountRDPPrivilege, return on the first match
		} else {
			dst.SecurityPolicyPrincipalAccountRDPPrivilege = nil
			return fmt.Errorf("failed to unmarshal SecurityPolicyRulePrivilegeContainerPrivilegeValue as SecurityPolicyPrincipalAccountRDPPrivilege: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SecurityPolicyPrincipalAccountSSHPrivilege'
	if jsonDict["_type"] == "SecurityPolicyPrincipalAccountSSHPrivilege" {
		// try to unmarshal JSON data into SecurityPolicyPrincipalAccountSSHPrivilege
		err = json.Unmarshal(data, &dst.SecurityPolicyPrincipalAccountSSHPrivilege)
		if err == nil {
			return nil // data stored in dst.SecurityPolicyPrincipalAccountSSHPrivilege, return on the first match
		} else {
			dst.SecurityPolicyPrincipalAccountSSHPrivilege = nil
			return fmt.Errorf("failed to unmarshal SecurityPolicyRulePrivilegeContainerPrivilegeValue as SecurityPolicyPrincipalAccountSSHPrivilege: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SecurityPolicyRevealPasswordPrivilege'
	if jsonDict["_type"] == "SecurityPolicyRevealPasswordPrivilege" {
		// try to unmarshal JSON data into SecurityPolicyRevealPasswordPrivilege
		err = json.Unmarshal(data, &dst.SecurityPolicyRevealPasswordPrivilege)
		if err == nil {
			return nil // data stored in dst.SecurityPolicyRevealPasswordPrivilege, return on the first match
		} else {
			dst.SecurityPolicyRevealPasswordPrivilege = nil
			return fmt.Errorf("failed to unmarshal SecurityPolicyRulePrivilegeContainerPrivilegeValue as SecurityPolicyRevealPasswordPrivilege: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SecurityPolicySecretPrivilege'
	if jsonDict["_type"] == "SecurityPolicySecretPrivilege" {
		// try to unmarshal JSON data into SecurityPolicySecretPrivilege
		err = json.Unmarshal(data, &dst.SecurityPolicySecretPrivilege)
		if err == nil {
			return nil // data stored in dst.SecurityPolicySecretPrivilege, return on the first match
		} else {
			dst.SecurityPolicySecretPrivilege = nil
			return fmt.Errorf("failed to unmarshal SecurityPolicyRulePrivilegeContainerPrivilegeValue as SecurityPolicySecretPrivilege: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SecurityPolicyUpdatePasswordPrivilege'
	if jsonDict["_type"] == "SecurityPolicyUpdatePasswordPrivilege" {
		// try to unmarshal JSON data into SecurityPolicyUpdatePasswordPrivilege
		err = json.Unmarshal(data, &dst.SecurityPolicyUpdatePasswordPrivilege)
		if err == nil {
			return nil // data stored in dst.SecurityPolicyUpdatePasswordPrivilege, return on the first match
		} else {
			dst.SecurityPolicyUpdatePasswordPrivilege = nil
			return fmt.Errorf("failed to unmarshal SecurityPolicyRulePrivilegeContainerPrivilegeValue as SecurityPolicyUpdatePasswordPrivilege: %s", err.Error())
		}
	}

	// check if the discriminator value is 'password_checkout_database'
	if jsonDict["_type"] == "password_checkout_database" {
		// try to unmarshal JSON data into SecurityPolicyPasswordCheckoutDatabasePrivilege
		err = json.Unmarshal(data, &dst.SecurityPolicyPasswordCheckoutDatabasePrivilege)
		if err == nil {
			return nil // data stored in dst.SecurityPolicyPasswordCheckoutDatabasePrivilege, return on the first match
		} else {
			dst.SecurityPolicyPasswordCheckoutDatabasePrivilege = nil
			return fmt.Errorf("failed to unmarshal SecurityPolicyRulePrivilegeContainerPrivilegeValue as SecurityPolicyPasswordCheckoutDatabasePrivilege: %s", err.Error())
		}
	}

	// check if the discriminator value is 'password_checkout_rdp'
	if jsonDict["_type"] == "password_checkout_rdp" {
		// try to unmarshal JSON data into SecurityPolicyPasswordCheckoutRDPPrivilege
		err = json.Unmarshal(data, &dst.SecurityPolicyPasswordCheckoutRDPPrivilege)
		if err == nil {
			return nil // data stored in dst.SecurityPolicyPasswordCheckoutRDPPrivilege, return on the first match
		} else {
			dst.SecurityPolicyPasswordCheckoutRDPPrivilege = nil
			return fmt.Errorf("failed to unmarshal SecurityPolicyRulePrivilegeContainerPrivilegeValue as SecurityPolicyPasswordCheckoutRDPPrivilege: %s", err.Error())
		}
	}

	// check if the discriminator value is 'password_checkout_ssh'
	if jsonDict["_type"] == "password_checkout_ssh" {
		// try to unmarshal JSON data into SecurityPolicyPasswordCheckoutSSHPrivilege
		err = json.Unmarshal(data, &dst.SecurityPolicyPasswordCheckoutSSHPrivilege)
		if err == nil {
			return nil // data stored in dst.SecurityPolicyPasswordCheckoutSSHPrivilege, return on the first match
		} else {
			dst.SecurityPolicyPasswordCheckoutSSHPrivilege = nil
			return fmt.Errorf("failed to unmarshal SecurityPolicyRulePrivilegeContainerPrivilegeValue as SecurityPolicyPasswordCheckoutSSHPrivilege: %s", err.Error())
		}
	}

	// check if the discriminator value is 'principal_account_rdp'
	if jsonDict["_type"] == "principal_account_rdp" {
		// try to unmarshal JSON data into SecurityPolicyPrincipalAccountRDPPrivilege
		err = json.Unmarshal(data, &dst.SecurityPolicyPrincipalAccountRDPPrivilege)
		if err == nil {
			return nil // data stored in dst.SecurityPolicyPrincipalAccountRDPPrivilege, return on the first match
		} else {
			dst.SecurityPolicyPrincipalAccountRDPPrivilege = nil
			return fmt.Errorf("failed to unmarshal SecurityPolicyRulePrivilegeContainerPrivilegeValue as SecurityPolicyPrincipalAccountRDPPrivilege: %s", err.Error())
		}
	}

	// check if the discriminator value is 'principal_account_ssh'
	if jsonDict["_type"] == "principal_account_ssh" {
		// try to unmarshal JSON data into SecurityPolicyPrincipalAccountSSHPrivilege
		err = json.Unmarshal(data, &dst.SecurityPolicyPrincipalAccountSSHPrivilege)
		if err == nil {
			return nil // data stored in dst.SecurityPolicyPrincipalAccountSSHPrivilege, return on the first match
		} else {
			dst.SecurityPolicyPrincipalAccountSSHPrivilege = nil
			return fmt.Errorf("failed to unmarshal SecurityPolicyRulePrivilegeContainerPrivilegeValue as SecurityPolicyPrincipalAccountSSHPrivilege: %s", err.Error())
		}
	}

	// check if the discriminator value is 'reveal_password'
	if jsonDict["_type"] == "reveal_password" {
		// try to unmarshal JSON data into SecurityPolicyRevealPasswordPrivilege
		err = json.Unmarshal(data, &dst.SecurityPolicyRevealPasswordPrivilege)
		if err == nil {
			return nil // data stored in dst.SecurityPolicyRevealPasswordPrivilege, return on the first match
		} else {
			dst.SecurityPolicyRevealPasswordPrivilege = nil
			return fmt.Errorf("failed to unmarshal SecurityPolicyRulePrivilegeContainerPrivilegeValue as SecurityPolicyRevealPasswordPrivilege: %s", err.Error())
		}
	}

	// check if the discriminator value is 'secret'
	if jsonDict["_type"] == "secret" {
		// try to unmarshal JSON data into SecurityPolicySecretPrivilege
		err = json.Unmarshal(data, &dst.SecurityPolicySecretPrivilege)
		if err == nil {
			return nil // data stored in dst.SecurityPolicySecretPrivilege, return on the first match
		} else {
			dst.SecurityPolicySecretPrivilege = nil
			return fmt.Errorf("failed to unmarshal SecurityPolicyRulePrivilegeContainerPrivilegeValue as SecurityPolicySecretPrivilege: %s", err.Error())
		}
	}

	// check if the discriminator value is 'update_password'
	if jsonDict["_type"] == "update_password" {
		// try to unmarshal JSON data into SecurityPolicyUpdatePasswordPrivilege
		err = json.Unmarshal(data, &dst.SecurityPolicyUpdatePasswordPrivilege)
		if err == nil {
			return nil // data stored in dst.SecurityPolicyUpdatePasswordPrivilege, return on the first match
		} else {
			dst.SecurityPolicyUpdatePasswordPrivilege = nil
			return fmt.Errorf("failed to unmarshal SecurityPolicyRulePrivilegeContainerPrivilegeValue as SecurityPolicyUpdatePasswordPrivilege: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SecurityPolicyRulePrivilegeContainerPrivilegeValue) MarshalJSON() ([]byte, error) {
	if src.SecurityPolicyPasswordCheckoutDatabasePrivilege != nil {
		return json.Marshal(&src.SecurityPolicyPasswordCheckoutDatabasePrivilege)
	}

	if src.SecurityPolicyPasswordCheckoutRDPPrivilege != nil {
		return json.Marshal(&src.SecurityPolicyPasswordCheckoutRDPPrivilege)
	}

	if src.SecurityPolicyPasswordCheckoutSSHPrivilege != nil {
		return json.Marshal(&src.SecurityPolicyPasswordCheckoutSSHPrivilege)
	}

	if src.SecurityPolicyPrincipalAccountRDPPrivilege != nil {
		return json.Marshal(&src.SecurityPolicyPrincipalAccountRDPPrivilege)
	}

	if src.SecurityPolicyPrincipalAccountSSHPrivilege != nil {
		return json.Marshal(&src.SecurityPolicyPrincipalAccountSSHPrivilege)
	}

	if src.SecurityPolicyRevealPasswordPrivilege != nil {
		return json.Marshal(&src.SecurityPolicyRevealPasswordPrivilege)
	}

	if src.SecurityPolicySecretPrivilege != nil {
		return json.Marshal(&src.SecurityPolicySecretPrivilege)
	}

	if src.SecurityPolicyUpdatePasswordPrivilege != nil {
		return json.Marshal(&src.SecurityPolicyUpdatePasswordPrivilege)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SecurityPolicyRulePrivilegeContainerPrivilegeValue) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.SecurityPolicyPasswordCheckoutDatabasePrivilege != nil {
		return obj.SecurityPolicyPasswordCheckoutDatabasePrivilege
	}

	if obj.SecurityPolicyPasswordCheckoutRDPPrivilege != nil {
		return obj.SecurityPolicyPasswordCheckoutRDPPrivilege
	}

	if obj.SecurityPolicyPasswordCheckoutSSHPrivilege != nil {
		return obj.SecurityPolicyPasswordCheckoutSSHPrivilege
	}

	if obj.SecurityPolicyPrincipalAccountRDPPrivilege != nil {
		return obj.SecurityPolicyPrincipalAccountRDPPrivilege
	}

	if obj.SecurityPolicyPrincipalAccountSSHPrivilege != nil {
		return obj.SecurityPolicyPrincipalAccountSSHPrivilege
	}

	if obj.SecurityPolicyRevealPasswordPrivilege != nil {
		return obj.SecurityPolicyRevealPasswordPrivilege
	}

	if obj.SecurityPolicySecretPrivilege != nil {
		return obj.SecurityPolicySecretPrivilege
	}

	if obj.SecurityPolicyUpdatePasswordPrivilege != nil {
		return obj.SecurityPolicyUpdatePasswordPrivilege
	}

	// all schemas are nil
	return nil
}

type NullableSecurityPolicyRulePrivilegeContainerPrivilegeValue struct {
	value *SecurityPolicyRulePrivilegeContainerPrivilegeValue
	isSet bool
}

func (v NullableSecurityPolicyRulePrivilegeContainerPrivilegeValue) Get() *SecurityPolicyRulePrivilegeContainerPrivilegeValue {
	return v.value
}

func (v *NullableSecurityPolicyRulePrivilegeContainerPrivilegeValue) Set(val *SecurityPolicyRulePrivilegeContainerPrivilegeValue) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityPolicyRulePrivilegeContainerPrivilegeValue) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityPolicyRulePrivilegeContainerPrivilegeValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityPolicyRulePrivilegeContainerPrivilegeValue(val *SecurityPolicyRulePrivilegeContainerPrivilegeValue) *NullableSecurityPolicyRulePrivilegeContainerPrivilegeValue {
	return &NullableSecurityPolicyRulePrivilegeContainerPrivilegeValue{value: val, isSet: true}
}

func (v NullableSecurityPolicyRulePrivilegeContainerPrivilegeValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityPolicyRulePrivilegeContainerPrivilegeValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
