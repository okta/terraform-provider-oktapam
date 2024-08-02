package oktapam

import (
	"context"
	"fmt"
	"sort"
	"testing"

	"github.com/atko-pam/pam-sdk-go/client/pam"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/kylelemons/godebug/pretty"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"
)

func TestAccResourceGroupDatabasePasswordSettings(t *testing.T) {
	checkTeamApplicable(t, true)
	resourceName := "oktapam_database_password_settings.test_acc_db_password_settings"
	resourceGroupName := fmt.Sprintf("test_acc_resource_group_%s", randSeq())
	projectName := fmt.Sprintf("test_acc_resource_group_project_%s", randSeq())
	delegatedAdminGroupName := fmt.Sprintf("test_acc_resource_group_dga_%s", randSeq())

	initialPasswordSettings := &pam.PasswordPolicy{
		ManagedPrivilegedAccountsConfig:   nil, // unused for databases
		EnablePeriodicRotation:            true,
		PeriodicRotationDurationInSeconds: utils.AsInt32PtrZero(3600, true),
		MinLengthInBytes:                  12,
		MaxLengthInBytes:                  16,
		CharacterOptions: pam.PasswordPolicyCharacterOptions{
			UpperCase:          utils.AsBoolPtrZero(true, true),
			LowerCase:          utils.AsBoolPtrZero(true, true),
			Digits:             utils.AsBoolPtrZero(true, true),
			Punctuation:        utils.AsBoolPtrZero(true, true),
			RequireFromEachSet: utils.AsBoolPtrZero(false, true),
		},
	}
	updatedPasswordSettings := &pam.PasswordPolicy{
		ManagedPrivilegedAccountsConfig:   nil,
		EnablePeriodicRotation:            false,
		PeriodicRotationDurationInSeconds: utils.AsInt32PtrZero(0, true),
		MinLengthInBytes:                  8,
		MaxLengthInBytes:                  12,
		CharacterOptions: pam.PasswordPolicyCharacterOptions{
			UpperCase:          utils.AsBoolPtrZero(true, true),
			LowerCase:          utils.AsBoolPtrZero(true, true),
			Digits:             utils.AsBoolPtrZero(false, true),
			Punctuation:        utils.AsBoolPtrZero(false, true),
			RequireFromEachSet: utils.AsBoolPtrZero(true, true),
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccV6ProviderFactories,
		// use the resource group check destroy since we create a new one here and deletion of the resource group will cascade delete the project / password settings
		CheckDestroy: testAccResourceGroupCheckDestroy(resourceGroupName),
		Steps: []resource.TestStep{
			{
				Config: createTestAccDatabasePasswordSettingsCreateConfig(delegatedAdminGroupName, resourceGroupName, projectName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccDatabasePasswordSettingsCheckExists(resourceName, initialPasswordSettings),
					resource.TestCheckResourceAttr(resourceName, attributes.EnablePeriodicRotation, "true"),
					resource.TestCheckResourceAttr(resourceName, attributes.PeriodicRotationDurationInSeconds, "3600"),
					resource.TestCheckResourceAttr(resourceName, attributes.MinLength, "12"),
					resource.TestCheckResourceAttr(resourceName, attributes.MaxLength, "16"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s", attributes.CharacterOptions, attributes.LowerCase), "true"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s", attributes.CharacterOptions, attributes.UpperCase), "true"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s", attributes.CharacterOptions, attributes.Digits), "true"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s", attributes.CharacterOptions, attributes.Punctuation), "true"),
				),
			},
			{
				Config: createTestAccDatabasePasswordSettingsUpdateConfig(delegatedAdminGroupName, resourceGroupName, projectName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccDatabasePasswordSettingsCheckExists(resourceName, updatedPasswordSettings),
					resource.TestCheckResourceAttr(resourceName, attributes.EnablePeriodicRotation, "false"),
					resource.TestCheckResourceAttr(resourceName, attributes.PeriodicRotationDurationInSeconds, "0"),
					resource.TestCheckResourceAttr(resourceName, attributes.MinLength, "8"),
					resource.TestCheckResourceAttr(resourceName, attributes.MaxLength, "12"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s", attributes.CharacterOptions, attributes.LowerCase), "true"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s", attributes.CharacterOptions, attributes.UpperCase), "true"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s", attributes.CharacterOptions, attributes.Digits), "false"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s", attributes.CharacterOptions, attributes.Punctuation), "false"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccDatabasePasswordSettingsImportStateId(resourceName),
			},
		},
	})
}

func testAccDatabasePasswordSettingsCheckExists(rn string, expectedPasswordSettings *pam.PasswordPolicy) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[rn]
		if !ok {
			return fmt.Errorf("resource not found: %s", rn)
		}

		id := rs.Primary.ID
		resourceGroupID := rs.Primary.Attributes[attributes.ResourceGroup]
		projectID := rs.Primary.Attributes[attributes.Project]
		expectedID := formatPasswordSettingsID(resourceGroupID, projectID)
		if id != expectedID {
			return fmt.Errorf("unexpected id: %s, expected: %s", id, expectedID)
		}
		pamClient := testAccAPIClients.SDKClient
		req := pamClient.SDKClient.ProjectsAPI.GetProjectPasswordPolicyForDatabaseResources(context.Background(), pamClient.Team, resourceGroupID, projectID)
		passwordSettings, _, err := req.Execute()
		if err != nil {
			return fmt.Errorf("error getting password settings: %w", err)
		} else if passwordSettings == nil {
			return fmt.Errorf("password settings does not exist")
		}
		insertComputedValuesForPasswordPolicy(expectedPasswordSettings, passwordSettings)

		comparison := pretty.Compare(expectedPasswordSettings, passwordSettings)
		if comparison != "" {
			return fmt.Errorf("expected password settings does not match returned password settings.\n%s", comparison)
		}
		return nil
	}
}

const testAccDatabasePasswordSettingsCreateConfigFormat = `
resource "oktapam_group" "test_resource_group_dga_group" {
	name = "%s"
}
resource "oktapam_resource_group" "test_acc_resource_group" {
	name = "%s"
	description = "test resource group"
	delegated_resource_admin_groups = [oktapam_group.test_resource_group_dga_group.id]	
}
resource "oktapam_resource_group_project" "test_acc_resource_group_project" {
	name = "%s"
	resource_group = oktapam_resource_group.test_acc_resource_group.id
	ssh_certificate_type  = "CERT_TYPE_ED25519_01"
	account_discovery     = true
}
resource "oktapam_database_password_settings" "test_acc_db_password_settings" {
	resource_group = oktapam_resource_group.test_acc_resource_group.id
	project = oktapam_resource_group_project.test_acc_resource_group_project.id
	enable_periodic_rotation    = true
	periodic_rotation_duration_in_seconds = 3600
    min_length = 12
	max_length = 16
	character_options {
		upper_case = true
		lower_case = true
		digits = true
		punctuation = true
	}
}
`

func createTestAccDatabasePasswordSettingsCreateConfig(dgaName, resourceGroupName, projectName string) string {
	return fmt.Sprintf(testAccDatabasePasswordSettingsCreateConfigFormat, dgaName, resourceGroupName, projectName)
}

const testAccDatabasePasswordSettingsUpdateConfigFormat = `
resource "oktapam_group" "test_resource_group_dga_group" {
	name = "%s"
}
resource "oktapam_resource_group" "test_acc_resource_group" {
	name = "%s"
	description = "test resource group"
	delegated_resource_admin_groups = [oktapam_group.test_resource_group_dga_group.id]	
}
resource "oktapam_resource_group_project" "test_acc_resource_group_project" {
	name = "%s"
	resource_group = oktapam_resource_group.test_acc_resource_group.id
	ssh_certificate_type  = "CERT_TYPE_ED25519_01"
	account_discovery     = true
}
resource "oktapam_database_password_settings" "test_acc_db_password_settings" {
	resource_group = oktapam_resource_group.test_acc_resource_group.id
	project = oktapam_resource_group_project.test_acc_resource_group_project.id
	enable_periodic_rotation    = false
    min_length = 8
	max_length = 12
	character_options {
		upper_case = true
		lower_case = true
		digits = false
		punctuation = false
		require_from_each_set = true
	}
}
`

func createTestAccDatabasePasswordSettingsUpdateConfig(dgaName, resourceGroupName, projectName string) string {
	return fmt.Sprintf(testAccDatabasePasswordSettingsUpdateConfigFormat, dgaName, resourceGroupName, projectName)
}

func testAccDatabasePasswordSettingsImportStateId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("Not found: %s", resourceName)
		}
		return fmt.Sprintf("%s/%s", rs.Primary.Attributes[attributes.ResourceGroup], rs.Primary.Attributes[attributes.Project]), nil
	}
}

func insertComputedValuesForPasswordPolicy(expectedPasswordSettings, actualPasswordSettings *pam.PasswordPolicy) error {
	sort.Strings(expectedPasswordSettings.ManagedPrivilegedAccountsConfig)
	sort.Strings(actualPasswordSettings.ManagedPrivilegedAccountsConfig)
	return nil
}
