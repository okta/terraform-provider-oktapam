package oktapam

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/kylelemons/godebug/pretty"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
)

func TestAccDatabaseResource(t *testing.T) {
	checkTeamApplicable(t, true)
	resourceName := "oktapam_database.test_acc_database_resource"
	groupName := fmt.Sprintf("test_acc_group_%s", randSeq())
	resourceGroupName := fmt.Sprintf("test_acc_resource_group_%s", randSeq())
	projectName := fmt.Sprintf("test_acc_project_%s", randSeq())

	initialDatabase := &pam.DatabaseResourceResponse{
		CanonicalName:                   "MyCanonicalName",
		DatabaseType:                    "mysql.level1",
		RecipeBookId:                    nil,
		ManagementConnectionDetailsType: MySqlBasicAuth,
		ManagementConnectionDetails: pam.ManagementConnectionDetails{
			MySQLBasicAuthManagementConnectionDetails: &pam.MySQLBasicAuthManagementConnectionDetails{
				Hostname: "mysql.example.org",
				Port:     "3306",
				AuthDetails: pam.MySQLBasicAuthDetails{
					Username:    "user",
					PasswordJwe: nil,
					SecretId:    nil,
				},
			},
		},
		ManagementGatewaySelector:   nil,
		ManagementGatewaySelectorId: "",
	}

	updatedDatabase := &pam.DatabaseResourceResponse{
		CanonicalName:                   "MyCanonicalName",
		DatabaseType:                    "mysql.level1",
		RecipeBookId:                    nil,
		ManagementConnectionDetailsType: MySqlBasicAuth,
		ManagementConnectionDetails: pam.ManagementConnectionDetails{
			MySQLBasicAuthManagementConnectionDetails: &pam.MySQLBasicAuthManagementConnectionDetails{
				Hostname: "mysql.example.org",
				Port:     "3306",
				AuthDetails: pam.MySQLBasicAuthDetails{
					Username:    "user",
					PasswordJwe: nil, // this is only used for POST/PUT, not GET
					SecretId:    nil, // Generated UUID - unknown until Read
				},
			},
		},
		ManagementGatewaySelector:   &map[string]string{"type": "db_management"},
		ManagementGatewaySelectorId: "", // Generated UUID - unknown until Read
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccV6ProviderFactories,
		CheckDestroy:             testAccResourceGroupCheckDestroy(resourceGroupName),
		Steps: []resource.TestStep{
			// negative cases must go first
			{
				Config:      createTestDatabaseInvalidCreateConfig(groupName, resourceGroupName, projectName),
				ExpectError: regexp.MustCompile("No more than 1 \"basic_auth\" blocks are allowed"),
			},
			// positive cases next
			{
				Config: createTestDatabaseCreateConfig(groupName, resourceGroupName, projectName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testDatabaseCheckExists(resourceName, initialDatabase),
					resource.TestCheckResourceAttr(resourceName, attributes.RecipeBook, ""),
					resource.TestCheckResourceAttr(resourceName, attributes.ManagementGatewaySelectorID, ""),
					resource.TestCheckNoResourceAttr(resourceName, attributes.ManagementGatewaySelector),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.#", attributes.ManagementConnectionDetails), "1"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%%", attributes.ManagementConnectionDetails), "1"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s.#", attributes.ManagementConnectionDetails, attributes.MySQL), "1"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s.0.%s.0.%%", attributes.ManagementConnectionDetails, attributes.MySQL, attributes.BasicAuth), "3"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s.0.%s.0.%s", attributes.ManagementConnectionDetails, attributes.MySQL, attributes.BasicAuth, attributes.Username), "user"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s.0.%s.0.%s", attributes.ManagementConnectionDetails, attributes.MySQL, attributes.BasicAuth, attributes.Password), ""),
					// Cannot use TestCheckNoResourceAttr here so just check it is empty
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s.0.%s.0.%s", attributes.ManagementConnectionDetails, attributes.MySQL, attributes.BasicAuth, attributes.Secret), ""),
				),
			},
			{
				// adds a selector and password
				Config: createTestDatabaseUpdateConfig(groupName, resourceGroupName, projectName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testDatabaseCheckExists(resourceName, updatedDatabase),
					resource.TestCheckResourceAttr(resourceName, attributes.RecipeBook, ""),
					resource.TestCheckResourceAttrSet(resourceName, attributes.ManagementGatewaySelectorID),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.%%", attributes.ManagementGatewaySelector), "1"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.#", attributes.ManagementConnectionDetails), "1"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%%", attributes.ManagementConnectionDetails), "1"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s.#", attributes.ManagementConnectionDetails, attributes.MySQL), "1"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s.0.%s.0.%%", attributes.ManagementConnectionDetails, attributes.MySQL, attributes.BasicAuth), "3"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s.0.%s.0.%s", attributes.ManagementConnectionDetails, attributes.MySQL, attributes.BasicAuth, attributes.Username), "user"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s.0.%s.0.%s", attributes.ManagementConnectionDetails, attributes.MySQL, attributes.BasicAuth, attributes.Password), "my-pass"),
					// This will fail if the value is an empty string.
					resource.TestCheckResourceAttrSet(resourceName, fmt.Sprintf("%s.0.%s.0.%s.0.%s", attributes.ManagementConnectionDetails, attributes.MySQL, attributes.BasicAuth, attributes.Secret)),
				),
			},
			{
				// verify an existing selector can be removed
				Config: createTestDatabaseRemoveSelectorConfig(groupName, resourceGroupName, projectName),
				Check: resource.ComposeAggregateTestCheckFunc(
					// selector is set but empty
					resource.TestCheckResourceAttrSet(resourceName, attributes.ManagementGatewaySelectorID),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.%%", attributes.ManagementGatewaySelector), "0"),
				),
			},
			{
				// verify the password can be modified
				Config: createTestDatabaseUpdatePasswordConfig(groupName, resourceGroupName, projectName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s.0.%s.0.%%", attributes.ManagementConnectionDetails, attributes.MySQL, attributes.BasicAuth), "3"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s.0.%s.0.%s", attributes.ManagementConnectionDetails, attributes.MySQL, attributes.BasicAuth, attributes.Username), "user"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s.0.%s.0.%s", attributes.ManagementConnectionDetails, attributes.MySQL, attributes.BasicAuth, attributes.Password), "new-pass"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"management_connection_details.0.mysql.0.basic_auth.0.password"},
				ImportStateIdFunc:       testAccDatabaseImportStateId(resourceName),
			},
		},
	})
}

func createTestDatabaseCreateConfig(groupName, resourceGroupName, projectName string) string {
	return fmt.Sprintf(testDatabaseResourceCreateConfigFormat, groupName, resourceGroupName, projectName)
}
func createTestDatabaseUpdateConfig(groupName, resourceGroupName, projectName string) string {
	return fmt.Sprintf(testDatabaseResourceUpdateConfigFormat, groupName, resourceGroupName, projectName)
}

func createTestDatabaseUpdatePasswordConfig(groupName, resourceGroupName, projectName string) string {
	return fmt.Sprintf(testDatabaseResourceUpdatePasswordConfigFormat, groupName, resourceGroupName, projectName)
}

func createTestDatabaseRemoveSelectorConfig(groupName, resourceGroupName, projectName string) string {
	return fmt.Sprintf(testDatabaseResourceRemoveSelectorConfigFormat, groupName, resourceGroupName, projectName)
}

func createTestDatabaseInvalidCreateConfig(groupName, resourceGroupName, projectName string) string {
	return fmt.Sprintf(testDatabaseResourceInvalidCreateConfigFormat, groupName, resourceGroupName, projectName)
}

// multiple auth_details blocks should be invalid.
const testDatabaseResourceInvalidCreateConfigFormat = `
resource "oktapam_group" "test_acc_resource_group_dga_group" {
	name = "%s"
}
resource "oktapam_resource_group" "test_acc_resource_group" {
	name = "%s"
	description = "test resource group"
	delegated_resource_admin_groups = [oktapam_group.test_acc_resource_group_dga_group.id]
}
resource "oktapam_resource_group_project" "test_acc_resource_group_project" {
	name                 = "%s"
	resource_group       = oktapam_resource_group.test_acc_resource_group.id
	gateway_selector     = "env=test"
	ssh_certificate_type = "CERT_TYPE_RSA_01"
	account_discovery 	 = true
}
resource "oktapam_database" "test_acc_database_resource" {
	resource_group = oktapam_resource_group.test_acc_resource_group.id
	project = oktapam_resource_group_project.test_acc_resource_group_project.id
	canonical_name = "MyCanonicalName"
	database_type = "mysql.level1"
	management_connection_details {
		mysql {
			hostname = "mysql.example.org"
			port = "3306"
			basic_auth {
				username = "user"
				password = "my-pass"
			}
			basic_auth {
				username = "user2"
				password = "my-pass2"
			}
		}
	}
	management_gateway_selector = {
		"type": "db_management"
	}
}
`

// A valid config with no selector or password.
const testDatabaseResourceCreateConfigFormat = `
resource "oktapam_group" "test_acc_resource_group_dga_group" {
	name = "%s"
}
resource "oktapam_resource_group" "test_acc_resource_group" {
	name = "%s"
	description = "test resource group"
	delegated_resource_admin_groups = [oktapam_group.test_acc_resource_group_dga_group.id]
}
resource "oktapam_resource_group_project" "test_acc_resource_group_project" {
	name                 = "%s"
	resource_group       = oktapam_resource_group.test_acc_resource_group.id
	gateway_selector     = "env=test"
	ssh_certificate_type = "CERT_TYPE_RSA_01"
	account_discovery 	 = true
}
resource "oktapam_database" "test_acc_database_resource" {
	resource_group = oktapam_resource_group.test_acc_resource_group.id
	project = oktapam_resource_group_project.test_acc_resource_group_project.id
	canonical_name = "MyCanonicalName"
	database_type = "mysql.level1"
	management_connection_details {
		mysql {
			hostname = "mysql.example.org"
			port = "3306"
			basic_auth {
				username = "user"
			}
		}
	}
}
`

// Adds a selector and password.
const testDatabaseResourceUpdateConfigFormat = `
resource "oktapam_group" "test_acc_resource_group_dga_group" {
	name = "%s"
}
resource "oktapam_resource_group" "test_acc_resource_group" {
	name = "%s"
	description = "test resource group"
	delegated_resource_admin_groups = [oktapam_group.test_acc_resource_group_dga_group.id]
}
resource "oktapam_resource_group_project" "test_acc_resource_group_project" {
	name                 = "%s"
	resource_group       = oktapam_resource_group.test_acc_resource_group.id
	gateway_selector     = "env=test"
	ssh_certificate_type = "CERT_TYPE_RSA_01"
	account_discovery 	 = true
}
resource "oktapam_database" "test_acc_database_resource" {
	resource_group = oktapam_resource_group.test_acc_resource_group.id
	project = oktapam_resource_group_project.test_acc_resource_group_project.id
	canonical_name = "MyCanonicalName"
	database_type = "mysql.level1"
	management_connection_details {
		mysql {
			hostname = "mysql.example.org"
			port = "3306"
			basic_auth {
				username = "user"
				password = "my-pass"
			}
		}
	}
	management_gateway_selector = {
		"type": "db_management"
	}
}
`

// updates the mysql password
const testDatabaseResourceUpdatePasswordConfigFormat = `
resource "oktapam_group" "test_acc_resource_group_dga_group" {
	name = "%s"
}
resource "oktapam_resource_group" "test_acc_resource_group" {
	name = "%s"
	description = "test resource group"
	delegated_resource_admin_groups = [oktapam_group.test_acc_resource_group_dga_group.id]
}
resource "oktapam_resource_group_project" "test_acc_resource_group_project" {
	name                 = "%s"
	resource_group       = oktapam_resource_group.test_acc_resource_group.id
	gateway_selector     = "env=test"
	ssh_certificate_type = "CERT_TYPE_RSA_01"
	account_discovery 	 = true
}
resource "oktapam_database" "test_acc_database_resource" {
	resource_group = oktapam_resource_group.test_acc_resource_group.id
	project = oktapam_resource_group_project.test_acc_resource_group_project.id
	canonical_name = "MyCanonicalName"
	database_type = "mysql.level1"
	management_connection_details {
		mysql {
			hostname = "mysql.example.org"
			port = "3306"
			basic_auth {
				username = "user"
				password = "new-pass"
			}
		}
	}
	management_gateway_selector = {
		"type": "db_management"
	}
}
`

// Removes the selector set in the initial update.
const testDatabaseResourceRemoveSelectorConfigFormat = `
resource "oktapam_group" "test_acc_resource_group_dga_group" {
	name = "%s"
}
resource "oktapam_resource_group" "test_acc_resource_group" {
	name = "%s"
	description = "test resource group"
	delegated_resource_admin_groups = [oktapam_group.test_acc_resource_group_dga_group.id]
}
resource "oktapam_resource_group_project" "test_acc_resource_group_project" {
	name                 = "%s"
	resource_group       = oktapam_resource_group.test_acc_resource_group.id
	gateway_selector     = "env=test"
	ssh_certificate_type = "CERT_TYPE_RSA_01"
	account_discovery 	 = true
}
resource "oktapam_database" "test_acc_database_resource" {
	resource_group = oktapam_resource_group.test_acc_resource_group.id
	project = oktapam_resource_group_project.test_acc_resource_group_project.id
	canonical_name = "MyCanonicalName"
	database_type = "mysql.level1"
	management_connection_details {
		mysql {
			hostname = "mysql.example.org"
			port = "3306"
			basic_auth {
				username = "user"
				password = "my-pass"
			}
		}
	}
}
`

func testDatabaseCheckExists(rn string, expectedDatabase *pam.DatabaseResourceResponse) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[rn]
		if !ok {
			return fmt.Errorf("resource not found: %s", rn)
		}

		resourceGroupID := rs.Primary.Attributes[attributes.ResourceGroup]
		projectID := rs.Primary.Attributes[attributes.Project]
		databaseID := rs.Primary.Attributes[attributes.ID]
		pamClient := getTestAccAPIClients().SDKClient
		database, _, err := pamClient.SDKClient.DatabaseResourcesAPI.GetDatabaseResource(context.Background(), pamClient.Team, resourceGroupID, projectID, databaseID).Execute()
		if err != nil {
			return fmt.Errorf("error getting database: %w", err)
		} else if database == nil {
			return fmt.Errorf("database does not exist")
		}

		err = insertComputedValuesForDatabase(expectedDatabase, database)
		if err != nil {
			return err
		}

		comparison := pretty.Compare(expectedDatabase, database)
		if comparison != "" {
			return fmt.Errorf("expected database does not match returned database.\n%s", comparison)
		}

		return nil
	}
}

func insertComputedValuesForDatabase(expectedDatabase, actualDatabase *pam.DatabaseResourceResponse) error {
	expectedDatabase.Id = actualDatabase.Id
	expectedDatabase.UpdatedAt = actualDatabase.UpdatedAt
	expectedDatabase.CreatedAt = actualDatabase.CreatedAt
	if actualDatabase.ManagementConnectionDetails.MySQLBasicAuthManagementConnectionDetails != nil && actualDatabase.ManagementConnectionDetails.MySQLBasicAuthManagementConnectionDetails.AuthDetails.SecretId != nil {
		expectedDatabase.ManagementConnectionDetails.MySQLBasicAuthManagementConnectionDetails.AuthDetails.SecretId = actualDatabase.ManagementConnectionDetails.MySQLBasicAuthManagementConnectionDetails.AuthDetails.SecretId
	}
	expectedDatabase.ManagementGatewaySelectorId = actualDatabase.ManagementGatewaySelectorId
	return nil
}

func testAccDatabaseImportStateId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("Not found: %s", resourceName)
		}
		return fmt.Sprintf("%s/%s/%s", rs.Primary.Attributes[attributes.ResourceGroup], rs.Primary.Attributes[attributes.Project], rs.Primary.Attributes[attributes.ID]), nil
	}
}
