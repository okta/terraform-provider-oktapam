package client

import (
	"testing"
)

func String(s string) *string {
	return &s
}

func TestValidateCloudConnectionName(t *testing.T) {
	data1 := CloudConnection{
		Name:     String("test_test"),
		Provider: String("aws"),
		CloudConnectionDetails: &CloudConnectionDetails{
			AccountId:  String("123456789012"),
			ExternalId: String("external_id_value"),
			RoleArn:    String("arn:aws:iam::123456789012:role/role_name"),
		},
	}
	result1 := validateCloudConnectionData(data1)
	if result1 != true {
		t.Error("Expected name validation to equal true")
	}

	data2 := CloudConnection{
		Name:     String("test!@#$%"),
		Provider: String("aws"),
		CloudConnectionDetails: &CloudConnectionDetails{
			AccountId:  String("123456789012"),
			ExternalId: String("external_id_value"),
			RoleArn:    String("arn:aws:iam::123456789012:role/role_name"),
		},
	}
	result2 := validateCloudConnectionData(data2)
	if result2 != false {
		t.Error("Expected cloud connection name validation to equal false")
	}

	data3 := CloudConnection{
		Name:     String("test test"),
		Provider: String("aws"),
		CloudConnectionDetails: &CloudConnectionDetails{
			AccountId:  String("123456789012"),
			ExternalId: String("external_id_value"),
			RoleArn:    String("arn:aws:iam::123456789012:role/role_name"),
		},
	}
	result3 := validateCloudConnectionData(data3)
	if result3 != false {
		t.Error("Expected cloud connection name validation to equal false")
	}
}

func TestValidateCloudConnectionAccountId(t *testing.T) {
	data1 := CloudConnection{
		Name:     String("test"),
		Provider: String("aws"),
		CloudConnectionDetails: &CloudConnectionDetails{
			AccountId:  String("123456789012"),
			ExternalId: String("external_id_value"),
			RoleArn:    String("arn:aws:iam::123456789012:role/role_name"),
		},
	}
	result1 := validateCloudConnectionData(data1)
	if result1 != true {
		t.Error("Expected cloud connection account id validation to equal true")
	}

	data2 := CloudConnection{
		Name:     String("test"),
		Provider: String("aws"),
		CloudConnectionDetails: &CloudConnectionDetails{
			AccountId:  String("12345678901200"),
			ExternalId: String("external_id_value"),
			RoleArn:    String("arn:aws:iam::123456789012:role/role_name"),
		},
	}
	result2 := validateCloudConnectionData(data2)
	if result2 != false {
		t.Error("Expected cloud connection account id validation to equal false")
	}

	data3 := CloudConnection{
		Name:     String("test"),
		Provider: String("aws"),
		CloudConnectionDetails: &CloudConnectionDetails{
			AccountId:  String("a12345678901"),
			ExternalId: String("external_id_value"),
			RoleArn:    String("arn:aws:iam::123456789012:role/role_name"),
		},
	}
	result3 := validateCloudConnectionData(data3)
	if result3 != false {
		t.Error("Expected cloud connection account id validation to equal false")
	}
}

func TestValidateCloudConnectionProvider(t *testing.T) {
	data1 := CloudConnection{
		Name:     String("test"),
		Provider: String("aws"),
		CloudConnectionDetails: &CloudConnectionDetails{
			AccountId:  String("123456789012"),
			ExternalId: String("external_id_value"),
			RoleArn:    String("arn:aws:iam::123456789012:role/role_name"),
		},
	}
	result1 := validateCloudConnectionData(data1)
	if result1 != true {
		t.Error("Expected cloud connection provider validation to equal true")
	}

	data2 := CloudConnection{
		Name:     String("test"),
		Provider: String("tmp"),
		CloudConnectionDetails: &CloudConnectionDetails{
			AccountId:  String("12345678901200"),
			ExternalId: String("external_id_value"),
			RoleArn:    String("arn:aws:iam::123456789012:role/role_name"),
		},
	}
	result2 := validateCloudConnectionData(data2)
	if result2 != false {
		t.Error("Expected cloud connection provider validation to equal false")
	}
}

func TestValidateCloudConnectionExternalID(t *testing.T) {
	data1 := CloudConnection{
		Name:     String("test"),
		Provider: String("aws"),
		CloudConnectionDetails: &CloudConnectionDetails{
			AccountId:  String("123456789012"),
			ExternalId: String("external_id_value"),
			RoleArn:    String("arn:aws:iam::123456789012:role/role_name"),
		},
	}
	result1 := validateCloudConnectionData(data1)
	if result1 != true {
		t.Error("Expected cloud connection external id validation to equal true")
	}

	data2 := CloudConnection{
		Name:     String("test"),
		Provider: String("tmp"),
		CloudConnectionDetails: &CloudConnectionDetails{
			AccountId:  String("12345678901200"),
			ExternalId: String(""),
			RoleArn:    String("arn:aws:iam::123456789012:role/role_name"),
		},
	}
	result2 := validateCloudConnectionData(data2)
	if result2 != false {
		t.Error("Expected cloud connection external id validation to equal false")
	}
}

func TestValidateCloudConnectionRoleARN(t *testing.T) {
	data1 := CloudConnection{
		Name:     String("test"),
		Provider: String("aws"),
		CloudConnectionDetails: &CloudConnectionDetails{
			AccountId:  String("123456789012"),
			ExternalId: String("external_id_value"),
			RoleArn:    String("arn:aws:iam::123456789012:role/role_name"),
		},
	}
	result1 := validateCloudConnectionData(data1)
	if result1 != true {
		t.Error("Expected cloud connection role arn validation to equal true")
	}

	data2 := CloudConnection{
		Name:     String("test"),
		Provider: String("tmp"),
		CloudConnectionDetails: &CloudConnectionDetails{
			AccountId:  String("12345678901200"),
			ExternalId: String("external_id_value"),
			RoleArn:    String(""),
		},
	}
	result2 := validateCloudConnectionData(data2)
	if result2 != false {
		t.Error("Expected cloud connection external id validation to equal false")
	}
}
