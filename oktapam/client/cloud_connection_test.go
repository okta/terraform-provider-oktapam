package client

import (
	"testing"

	"github.com/okta/terraform-provider-oktapam/oktapam/utils"
)

func TestValidateCloudConnectionName(t *testing.T) {
	// testing invalid name regexp, underscore is not allowed
	data1 := CloudConnection{
		Name:     utils.AsStringPtr("test_test"),
		Provider: utils.AsStringPtr("aws"),
		CloudConnectionDetails: &CloudConnectionDetails{
			AccountId:  utils.AsStringPtr("123456789012"),
			ExternalId: utils.AsStringPtr("550e8400-e29b-41d4-a716-446655440000"),
			RoleArn:    utils.AsStringPtr("arn:aws:iam::123456789012:role/role_name"),
		},
	}
	result1, _ := isCloudConnectionDataValid(data1)
	if !result1 {
		t.Error("Expected name validation to equal true")
	}

	// testing invalid name regexp, special characters are not allowed
	data2 := CloudConnection{
		Name:     utils.AsStringPtr("test!@#$%"),
		Provider: utils.AsStringPtr("aws"),
		CloudConnectionDetails: &CloudConnectionDetails{
			AccountId:  utils.AsStringPtr("123456789012"),
			ExternalId: utils.AsStringPtr("550e8400-e29b-41d4-a716-446655440000"),
			RoleArn:    utils.AsStringPtr("arn:aws:iam::123456789012:role/role_name"),
		},
	}
	result2, _ := isCloudConnectionDataValid(data2)
	if result2 {
		t.Error("Expected cloud connection name validation to equal false")
	}

	// testing invalid name regexp, space is not allowed
	data3 := CloudConnection{
		Name:     utils.AsStringPtr("test test"),
		Provider: utils.AsStringPtr("aws"),
		CloudConnectionDetails: &CloudConnectionDetails{
			AccountId:  utils.AsStringPtr("123456789012"),
			ExternalId: utils.AsStringPtr("550e8400-e29b-41d4-a716-446655440000"),
			RoleArn:    utils.AsStringPtr("arn:aws:iam::123456789012:role/role_name"),
		},
	}
	result3, _ := isCloudConnectionDataValid(data3)
	if result3 {
		t.Error("Expected cloud connection name validation to equal false")
	}

	// testing empty name, should be invalid
	data4 := CloudConnection{
		Name:     utils.AsStringPtr(""),
		Provider: utils.AsStringPtr("aws"),
		CloudConnectionDetails: &CloudConnectionDetails{
			AccountId:  utils.AsStringPtr("123456789012"),
			ExternalId: utils.AsStringPtr("550e8400-e29b-41d4-a716-446655440000"),
			RoleArn:    utils.AsStringPtr("arn:aws:iam::123456789012:role/role_name"),
		},
	}
	result4, _ := isCloudConnectionDataValid(data4)
	if result4 {
		t.Error("Expected cloud connection name validation to equal false")
	}
}

func TestValidateCloudConnectionAccountId(t *testing.T) {
	// account id should be 12 digits
	data1 := CloudConnection{
		Name:     utils.AsStringPtr("test"),
		Provider: utils.AsStringPtr("aws"),
		CloudConnectionDetails: &CloudConnectionDetails{
			AccountId:  utils.AsStringPtr("123456789012"),
			ExternalId: utils.AsStringPtr("550e8400-e29b-41d4-a716-446655440000"),
			RoleArn:    utils.AsStringPtr("arn:aws:iam::123456789012:role/role_name"),
		},
	}
	result1, _ := isCloudConnectionDataValid(data1)
	if !result1 {
		t.Error("Expected cloud connection account id validation to equal true")
	}

	// account id is 14 digits
	data2 := CloudConnection{
		Name:     utils.AsStringPtr("test"),
		Provider: utils.AsStringPtr("aws"),
		CloudConnectionDetails: &CloudConnectionDetails{
			AccountId:  utils.AsStringPtr("12345678901200"),
			ExternalId: utils.AsStringPtr("550e8400-e29b-41d4-a716-446655440000"),
			RoleArn:    utils.AsStringPtr("arn:aws:iam::123456789012:role/role_name"),
		},
	}
	result2, _ := isCloudConnectionDataValid(data2)
	if result2 {
		t.Error("Expected cloud connection account id validation to equal false")
	}

	// account id is not all digits
	data3 := CloudConnection{
		Name:     utils.AsStringPtr("test"),
		Provider: utils.AsStringPtr("aws"),
		CloudConnectionDetails: &CloudConnectionDetails{
			AccountId:  utils.AsStringPtr("a12345678901"),
			ExternalId: utils.AsStringPtr("550e8400-e29b-41d4-a716-446655440000"),
			RoleArn:    utils.AsStringPtr("arn:aws:iam::123456789012:role/role_name"),
		},
	}
	result3, _ := isCloudConnectionDataValid(data3)
	if result3 {
		t.Error("Expected cloud connection account id validation to equal false")
	}
}

func TestValidateCloudConnectionProvider(t *testing.T) {
	// valid provider name "aws"
	data1 := CloudConnection{
		Name:     utils.AsStringPtr("test"),
		Provider: utils.AsStringPtr("aws"),
		CloudConnectionDetails: &CloudConnectionDetails{
			AccountId:  utils.AsStringPtr("123456789012"),
			ExternalId: utils.AsStringPtr("550e8400-e29b-41d4-a716-446655440000"),
			RoleArn:    utils.AsStringPtr("arn:aws:iam::123456789012:role/role_name"),
		},
	}
	result1, _ := isCloudConnectionDataValid(data1)
	if !result1 {
		t.Error("Expected cloud connection provider validation to equal true")
	}

	// invalid provider name, the only accepted value is "aws"
	data2 := CloudConnection{
		Name:     utils.AsStringPtr("test"),
		Provider: utils.AsStringPtr("tmp"),
		CloudConnectionDetails: &CloudConnectionDetails{
			AccountId:  utils.AsStringPtr("12345678901200"),
			ExternalId: utils.AsStringPtr("550e8400-e29b-41d4-a716-446655440000"),
			RoleArn:    utils.AsStringPtr("arn:aws:iam::123456789012:role/role_name"),
		},
	}
	result2, _ := isCloudConnectionDataValid(data2)
	if result2 {
		t.Error("Expected cloud connection provider validation to equal false")
	}
}

func TestValidateCloudConnectionExternalID(t *testing.T) {
	// invalid external id, external id should be a UUID
	data1 := CloudConnection{
		Name:     utils.AsStringPtr("test"),
		Provider: utils.AsStringPtr("aws"),
		CloudConnectionDetails: &CloudConnectionDetails{
			AccountId:  utils.AsStringPtr("123456789012"),
			ExternalId: utils.AsStringPtr("invalid_external_id_value"),
			RoleArn:    utils.AsStringPtr("arn:aws:iam::123456789012:role/role_name"),
		},
	}
	result1, _ := isCloudConnectionDataValid(data1)
	if result1 {
		t.Error("Expected cloud connection external id validation to equal false")
	}

	// empty external id is not valid
	data2 := CloudConnection{
		Name:     utils.AsStringPtr("test"),
		Provider: utils.AsStringPtr("tmp"),
		CloudConnectionDetails: &CloudConnectionDetails{
			AccountId:  utils.AsStringPtr("123456789012"),
			ExternalId: utils.AsStringPtr(""),
			RoleArn:    utils.AsStringPtr("arn:aws:iam::123456789012:role/role_name"),
		},
	}
	result2, _ := isCloudConnectionDataValid(data2)
	if result2 {
		t.Error("Expected cloud connection external id validation to equal false")
	}

	// external id should be a valid UUID
	data3 := CloudConnection{
		Name:     utils.AsStringPtr("test"),
		Provider: utils.AsStringPtr("aws"),
		CloudConnectionDetails: &CloudConnectionDetails{
			AccountId:  utils.AsStringPtr("123456789012"),
			ExternalId: utils.AsStringPtr("550e8400-e29b-41d4-a716-446655440000"),
			RoleArn:    utils.AsStringPtr("arn:aws:iam::123456789012:role/role_name"),
		},
	}
	result3, _ := isCloudConnectionDataValid(data3)
	if !result3 {
		t.Error("Expected cloud connection external id validation to equal true")
	}
}

func TestValidateCloudConnectionRoleARN(t *testing.T) {
	data1 := CloudConnection{
		Name:     utils.AsStringPtr("test"),
		Provider: utils.AsStringPtr("aws"),
		CloudConnectionDetails: &CloudConnectionDetails{
			AccountId:  utils.AsStringPtr("123456789012"),
			ExternalId: utils.AsStringPtr("550e8400-e29b-41d4-a716-446655440000"),
			RoleArn:    utils.AsStringPtr("arn:aws:iam::123456789012:role/role_name"),
		},
	}
	result1, _ := isCloudConnectionDataValid(data1)
	if !result1 {
		t.Error("Expected cloud connection role arn validation to equal true")
	}

	// empty string is not a valid value for role arn
	data2 := CloudConnection{
		Name:     utils.AsStringPtr("test"),
		Provider: utils.AsStringPtr("tmp"),
		CloudConnectionDetails: &CloudConnectionDetails{
			AccountId:  utils.AsStringPtr("12345678901200"),
			ExternalId: utils.AsStringPtr("550e8400-e29b-41d4-a716-446655440000"),
			RoleArn:    utils.AsStringPtr(""),
		},
	}
	result2, _ := isCloudConnectionDataValid(data2)
	if result2 {
		t.Error("Expected cloud connection external id validation to equal false")
	}
}
