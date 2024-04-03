package client

import (
	"testing"

	"github.com/okta/terraform-provider-oktapam/oktapam/utils"
)

func TestValidateCloudConnectionName(t *testing.T) {
	data1 := CloudConnection{
		Name:     utils.AsStringPtr("test_test"),
		Provider: utils.AsStringPtr("aws"),
		CloudConnectionDetails: &CloudConnectionDetails{
			AccountId:  utils.AsStringPtr("123456789012"),
			ExternalId: utils.AsStringPtr("external_id_value"),
			RoleArn:    utils.AsStringPtr("arn:aws:iam::123456789012:role/role_name"),
		},
	}
	result1 := validateCloudConnectionData(data1)
	if result1 != true {
		t.Error("Expected name validation to equal true")
	}

	data2 := CloudConnection{
		Name:     utils.AsStringPtr("test!@#$%"),
		Provider: utils.AsStringPtr("aws"),
		CloudConnectionDetails: &CloudConnectionDetails{
			AccountId:  utils.AsStringPtr("123456789012"),
			ExternalId: utils.AsStringPtr("external_id_value"),
			RoleArn:    utils.AsStringPtr("arn:aws:iam::123456789012:role/role_name"),
		},
	}
	result2 := validateCloudConnectionData(data2)
	if result2 != false {
		t.Error("Expected cloud connection name validation to equal false")
	}

	data3 := CloudConnection{
		Name:     utils.AsStringPtr("test test"),
		Provider: utils.AsStringPtr("aws"),
		CloudConnectionDetails: &CloudConnectionDetails{
			AccountId:  utils.AsStringPtr("123456789012"),
			ExternalId: utils.AsStringPtr("external_id_value"),
			RoleArn:    utils.AsStringPtr("arn:aws:iam::123456789012:role/role_name"),
		},
	}
	result3 := validateCloudConnectionData(data3)
	if result3 != false {
		t.Error("Expected cloud connection name validation to equal false")
	}
}

func TestValidateCloudConnectionAccountId(t *testing.T) {
	data1 := CloudConnection{
		Name:     utils.AsStringPtr("test"),
		Provider: utils.AsStringPtr("aws"),
		CloudConnectionDetails: &CloudConnectionDetails{
			AccountId:  utils.AsStringPtr("123456789012"),
			ExternalId: utils.AsStringPtr("external_id_value"),
			RoleArn:    utils.AsStringPtr("arn:aws:iam::123456789012:role/role_name"),
		},
	}
	result1 := validateCloudConnectionData(data1)
	if result1 != true {
		t.Error("Expected cloud connection account id validation to equal true")
	}

	data2 := CloudConnection{
		Name:     utils.AsStringPtr("test"),
		Provider: utils.AsStringPtr("aws"),
		CloudConnectionDetails: &CloudConnectionDetails{
			AccountId:  utils.AsStringPtr("12345678901200"),
			ExternalId: utils.AsStringPtr("external_id_value"),
			RoleArn:    utils.AsStringPtr("arn:aws:iam::123456789012:role/role_name"),
		},
	}
	result2 := validateCloudConnectionData(data2)
	if result2 != false {
		t.Error("Expected cloud connection account id validation to equal false")
	}

	data3 := CloudConnection{
		Name:     utils.AsStringPtr("test"),
		Provider: utils.AsStringPtr("aws"),
		CloudConnectionDetails: &CloudConnectionDetails{
			AccountId:  utils.AsStringPtr("a12345678901"),
			ExternalId: utils.AsStringPtr("external_id_value"),
			RoleArn:    utils.AsStringPtr("arn:aws:iam::123456789012:role/role_name"),
		},
	}
	result3 := validateCloudConnectionData(data3)
	if result3 != false {
		t.Error("Expected cloud connection account id validation to equal false")
	}
}

func TestValidateCloudConnectionProvider(t *testing.T) {
	data1 := CloudConnection{
		Name:     utils.AsStringPtr("test"),
		Provider: utils.AsStringPtr("aws"),
		CloudConnectionDetails: &CloudConnectionDetails{
			AccountId:  utils.AsStringPtr("123456789012"),
			ExternalId: utils.AsStringPtr("external_id_value"),
			RoleArn:    utils.AsStringPtr("arn:aws:iam::123456789012:role/role_name"),
		},
	}
	result1 := validateCloudConnectionData(data1)
	if result1 != true {
		t.Error("Expected cloud connection provider validation to equal true")
	}

	data2 := CloudConnection{
		Name:     utils.AsStringPtr("test"),
		Provider: utils.AsStringPtr("tmp"),
		CloudConnectionDetails: &CloudConnectionDetails{
			AccountId:  utils.AsStringPtr("12345678901200"),
			ExternalId: utils.AsStringPtr("external_id_value"),
			RoleArn:    utils.AsStringPtr("arn:aws:iam::123456789012:role/role_name"),
		},
	}
	result2 := validateCloudConnectionData(data2)
	if result2 != false {
		t.Error("Expected cloud connection provider validation to equal false")
	}
}

func TestValidateCloudConnectionExternalID(t *testing.T) {
	data1 := CloudConnection{
		Name:     utils.AsStringPtr("test"),
		Provider: utils.AsStringPtr("aws"),
		CloudConnectionDetails: &CloudConnectionDetails{
			AccountId:  utils.AsStringPtr("123456789012"),
			ExternalId: utils.AsStringPtr("external_id_value"),
			RoleArn:    utils.AsStringPtr("arn:aws:iam::123456789012:role/role_name"),
		},
	}
	result1 := validateCloudConnectionData(data1)
	if result1 != true {
		t.Error("Expected cloud connection external id validation to equal true")
	}

	data2 := CloudConnection{
		Name:     utils.AsStringPtr("test"),
		Provider: utils.AsStringPtr("tmp"),
		CloudConnectionDetails: &CloudConnectionDetails{
			AccountId:  utils.AsStringPtr("12345678901200"),
			ExternalId: utils.AsStringPtr(""),
			RoleArn:    utils.AsStringPtr("arn:aws:iam::123456789012:role/role_name"),
		},
	}
	result2 := validateCloudConnectionData(data2)
	if result2 != false {
		t.Error("Expected cloud connection external id validation to equal false")
	}
}

func TestValidateCloudConnectionRoleARN(t *testing.T) {
	data1 := CloudConnection{
		Name:     utils.AsStringPtr("test"),
		Provider: utils.AsStringPtr("aws"),
		CloudConnectionDetails: &CloudConnectionDetails{
			AccountId:  utils.AsStringPtr("123456789012"),
			ExternalId: utils.AsStringPtr("external_id_value"),
			RoleArn:    utils.AsStringPtr("arn:aws:iam::123456789012:role/role_name"),
		},
	}
	result1 := validateCloudConnectionData(data1)
	if result1 != true {
		t.Error("Expected cloud connection role arn validation to equal true")
	}

	data2 := CloudConnection{
		Name:     utils.AsStringPtr("test"),
		Provider: utils.AsStringPtr("tmp"),
		CloudConnectionDetails: &CloudConnectionDetails{
			AccountId:  utils.AsStringPtr("12345678901200"),
			ExternalId: utils.AsStringPtr("external_id_value"),
			RoleArn:    utils.AsStringPtr(""),
		},
	}
	result2 := validateCloudConnectionData(data2)
	if result2 != false {
		t.Error("Expected cloud connection external id validation to equal false")
	}
}
