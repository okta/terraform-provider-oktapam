package sdkv2

import (
	"context"
	"fmt"
	"sort"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/kylelemons/godebug/pretty"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/sdkv2"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"
)

func TestAccResourceGroupPasswordSettings(t *testing.T) {
	checkTeamApplicable(t, true)
	resourceName := "oktapam_password_settings.test_acc_password_settings"
	resourceGroupName := fmt.Sprintf("test_acc_resource_group_%s", randSeq())
	projectName := fmt.Sprintf("test_acc_resource_group_project_%s", randSeq())
	delegatedAdminGroupName := fmt.Sprintf("test_acc_resource_group_dga_%s", randSeq())

	initialPasswordSettings := &client.PasswordSettings{
		ManagedPrivilegedAccountsConfig:   []string{"root", "pamadmin"},
		EnablePeriodicRotation:            utils.AsBoolPtrZero(true, true),
		PeriodicRotationDurationInSeconds: utils.AsIntPtrZero(3600, true),
		MinLengthInBytes:                  utils.AsIntPtrZero(12, true),
		MaxLengthInBytes:                  utils.AsIntPtrZero(16, true),
		CharacterOptions: &client.CharacterOptions{
			UpperCase:   utils.AsBoolPtrZero(true, true),
			LowerCase:   utils.AsBoolPtrZero(true, true),
			Digits:      utils.AsBoolPtrZero(true, true),
			Punctuation: utils.AsBoolPtrZero(true, true),
		},
	}
	updatedPasswordSettings := &client.PasswordSettings{
		ManagedPrivilegedAccountsConfig:   []string{"root"},
		EnablePeriodicRotation:            utils.AsBoolPtrZero(false, true),
		PeriodicRotationDurationInSeconds: utils.AsIntPtrZero(0, true),
		MinLengthInBytes:                  utils.AsIntPtrZero(8, true),
		MaxLengthInBytes:                  utils.AsIntPtrZero(12, true),
		CharacterOptions: &client.CharacterOptions{
			UpperCase:   utils.AsBoolPtrZero(true, true),
			LowerCase:   utils.AsBoolPtrZero(true, true),
			Digits:      utils.AsBoolPtrZero(false, true),
			Punctuation: utils.AsBoolPtrZero(false, true),
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		// use the resource group check destroy since we create a new one here and deletion of the resource group will cascade delete the project / password settings
		CheckDestroy: testAccResourceGroupCheckDestroy(resourceGroupName),
		Steps: []resource.TestStep{
			{
				Config: createTestAccPasswordSettingsCreateConfig(delegatedAdminGroupName, resourceGroupName, projectName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccPasswordSettingsCheckExists(resourceName, initialPasswordSettings),
					resource.TestCheckResourceAttr(
						resourceName, attributes.EnablePeriodicRotation, "true",
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.PeriodicRotationDurationInSeconds, "3600",
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.MinLength, "12",
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.MaxLength, "16",
					),
					resource.TestCheckResourceAttr(
						resourceName, fmt.Sprintf("%s.0.%s", attributes.CharacterOptions, attributes.LowerCase), "true",
					),
					resource.TestCheckResourceAttr(
						resourceName, fmt.Sprintf("%s.0.%s", attributes.CharacterOptions, attributes.UpperCase), "true",
					),
					resource.TestCheckResourceAttr(
						resourceName, fmt.Sprintf("%s.0.%s", attributes.CharacterOptions, attributes.Digits), "true",
					),
					resource.TestCheckResourceAttr(
						resourceName, fmt.Sprintf("%s.0.%s", attributes.CharacterOptions, attributes.Punctuation), "true",
					),
				),
			},
			{
				Config: createTestAccPasswordSettingsUpdateConfig(delegatedAdminGroupName, resourceGroupName, projectName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccPasswordSettingsCheckExists(resourceName, updatedPasswordSettings),
					resource.TestCheckResourceAttr(
						resourceName, attributes.EnablePeriodicRotation, "false",
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.PeriodicRotationDurationInSeconds, "0",
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.MinLength, "8",
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.MaxLength, "12",
					),
					resource.TestCheckResourceAttr(
						resourceName, fmt.Sprintf("%s.0.%s", attributes.CharacterOptions, attributes.LowerCase), "true",
					),
					resource.TestCheckResourceAttr(
						resourceName, fmt.Sprintf("%s.0.%s", attributes.CharacterOptions, attributes.UpperCase), "true",
					),
					resource.TestCheckResourceAttr(
						resourceName, fmt.Sprintf("%s.0.%s", attributes.CharacterOptions, attributes.Digits), "false",
					),
					resource.TestCheckResourceAttr(
						resourceName, fmt.Sprintf("%s.0.%s", attributes.CharacterOptions, attributes.Punctuation), "false",
					),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccPasswordSettingsImportStateId(resourceName),
			},
		},
	})
}

func testAccPasswordSettingsCheckExists(rn string, expectedPasswordSettings *client.PasswordSettings) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[rn]
		if !ok {
			return fmt.Errorf("resource not found: %s", rn)
		}

		id := rs.Primary.ID
		resourceGroupID := rs.Primary.Attributes[attributes.ResourceGroup]
		projectID := rs.Primary.Attributes[attributes.Project]
		expectedID := sdkv2.FormatPasswordSettingsID(resourceGroupID, projectID)
		if id != expectedID {
			return fmt.Errorf("unexpected id: %s, expected: %s", id, expectedID)
		}
		pamClient := client.GetLocalClientFromMetadata(testAccProvider.Meta())
		passwordSettings, err := pamClient.GetPasswordSettings(context.Background(), resourceGroupID, projectID)
		if err != nil {
			return fmt.Errorf("error getting password settings: %w", err)
		} else if passwordSettings == nil {
			return fmt.Errorf("password settings does not exist")
		}
		insertComputedValuesForPasswordSettings(expectedPasswordSettings, passwordSettings)

		comparison := pretty.Compare(expectedPasswordSettings, passwordSettings)
		if comparison != "" {
			return fmt.Errorf("expected password settings does not match returned password settings.\n%s", comparison)
		}
		return nil
	}
}

func insertComputedValuesForPasswordSettings(expectedPasswordSettings, actualPasswordSettings *client.PasswordSettings) error {
	sort.Strings(expectedPasswordSettings.ManagedPrivilegedAccountsConfig)
	sort.Strings(actualPasswordSettings.ManagedPrivilegedAccountsConfig)
	return nil
}

const testAccPasswordSettingsCreateConfigFormat = `
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
resource "oktapam_password_settings" "test_acc_password_settings" {
	resource_group = oktapam_resource_group.test_acc_resource_group.id
	project = oktapam_resource_group_project.test_acc_resource_group_project.id
	managed_privileged_accounts = ["root", "pamadmin"]
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

func createTestAccPasswordSettingsCreateConfig(dgaName, resourceGroupName, projectName string) string {
	return fmt.Sprintf(testAccPasswordSettingsCreateConfigFormat, dgaName, resourceGroupName, projectName)
}

const testAccPasswordSettingsUpdateConfigFormat = `
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
resource "oktapam_password_settings" "test_acc_password_settings" {
	resource_group = oktapam_resource_group.test_acc_resource_group.id
	project = oktapam_resource_group_project.test_acc_resource_group_project.id
	managed_privileged_accounts = ["root"]
	enable_periodic_rotation    = false
    min_length = 8
	max_length = 12
	character_options {
		upper_case = true
		lower_case = true
		digits = false
		punctuation = false
	}
}
`

func createTestAccPasswordSettingsUpdateConfig(dgaName, resourceGroupName, projectName string) string {
	return fmt.Sprintf(testAccPasswordSettingsUpdateConfigFormat, dgaName, resourceGroupName, projectName)
}

func testAccPasswordSettingsImportStateId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("Not found: %s", resourceName)
		}
		return fmt.Sprintf("%s/%s", rs.Primary.Attributes[attributes.ResourceGroup], rs.Primary.Attributes[attributes.Project]), nil
	}
}
