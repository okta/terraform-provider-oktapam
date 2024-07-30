package oktapam

import (
	"context"
	"fmt"
	"reflect"
	"regexp"
	"testing"

	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
)

// testAccServerCheckoutSettingsBaseConfigFormat base config to create group, resource group and project
const testAccServerCheckoutSettingsBaseConfigFormat = `
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
`

const testAccServerCheckoutSettingsCreateConfigFormat = `
resource "oktapam_server_checkout_settings" "test_acc_server_checkout_settings" {
	resource_group = oktapam_resource_group.test_acc_resource_group.id
	project = oktapam_resource_group_project.test_acc_resource_group_project.id
	checkout_required = true
	checkout_duration_in_seconds = 900
}
`

const testAccServerCheckoutSettingsUpdateConfigFormat = `
resource "oktapam_server_checkout_settings" "test_acc_server_checkout_settings" {
	resource_group = oktapam_resource_group.test_acc_resource_group.id
	project = oktapam_resource_group_project.test_acc_resource_group_project.id
	checkout_required = true
	checkout_duration_in_seconds = 3600
	include_list = ["vaulted_account_1", "vaulted_account_2"]
}
`

const testAccServerCheckoutSettingsUpdateWithExcludeListConfigFormat = `
resource "oktapam_server_checkout_settings" "test_acc_server_checkout_settings" {
	resource_group = oktapam_resource_group.test_acc_resource_group.id
	project = oktapam_resource_group_project.test_acc_resource_group_project.id
	checkout_required = true
	checkout_duration_in_seconds = 3600
	exclude_list = ["excluded_account_1", "excluded_account_2"]
}
`

const testAccServerCheckoutSettingsUpdateWithBothListsConfigFormat = `
resource "oktapam_server_checkout_settings" "test_acc_server_checkout_settings" {
	resource_group = oktapam_resource_group.test_acc_resource_group.id
	project = oktapam_resource_group_project.test_acc_resource_group_project.id
	checkout_required = true
	checkout_duration_in_seconds = 7200
	include_list = ["included_account_1", "included_account_2"]
	exclude_list = ["excluded_account_1", "excluded_account_2"]
}
`

const testAccServerCheckoutSettingsUpdateProjectConfigFormat = `
resource "oktapam_resource_group_project" "test_acc_updated_project" {
	name = "%s"
	resource_group = oktapam_resource_group.test_acc_resource_group.id
	ssh_certificate_type  = "CERT_TYPE_ED25519_01"
	account_discovery     = true
}

resource "oktapam_server_checkout_settings" "test_acc_server_checkout_settings" {
	resource_group = oktapam_resource_group.test_acc_resource_group.id
	project = oktapam_resource_group_project.test_acc_updated_project.id
	checkout_required = true
	checkout_duration_in_seconds = 5400
}
`

const testAccServerCheckoutSettingsDeleteConfigFormat = `
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
`

func createTestAccServerCheckoutSettingsDeleteConfig(dgaName, resourceGroupName, projectName string) string {
	return fmt.Sprintf(testAccServerCheckoutSettingsDeleteConfigFormat, dgaName, resourceGroupName, projectName)
}
func TestAccServerCheckoutSettings(t *testing.T) {
	checkTeamApplicable(t, true)
	resourceName := "oktapam_server_checkout_settings.test_acc_server_checkout_settings"
	resourceGroupName := fmt.Sprintf("test_acc_resource_group_%s", randSeq())
	projectName := fmt.Sprintf("test_acc_resource_group_project_%s", randSeq())
	updatedProjectName := fmt.Sprintf("test_acc_updated_resource_group_project_%s", randSeq())
	delegatedAdminGroupName := fmt.Sprintf("test_acc_resource_group_dga_%s", randSeq())
	defaultCheckoutDuration := int32(900)
	updatedCheckoutDuration := int32(3600)
	updatedProjectCheckoutDuration := int32(5400)

	initialServerCheckoutSettings := &pam.ResourceCheckoutSettings{
		CheckoutRequired:          true,
		CheckoutDurationInSeconds: &defaultCheckoutDuration,
	}

	updatedServerCheckoutSettings := &pam.ResourceCheckoutSettings{
		CheckoutRequired:          true,
		CheckoutDurationInSeconds: &updatedCheckoutDuration,
		IncludeList:               []string{"vaulted_account_1", "vaulted_account_2"},
	}

	excludeListServerCheckoutSettings := &pam.ResourceCheckoutSettings{
		CheckoutRequired:          true,
		CheckoutDurationInSeconds: &updatedCheckoutDuration,
		ExcludeList:               []string{"excluded_account_1", "excluded_account_2"},
	}

	updatedProjectServerCheckoutSettings := &pam.ResourceCheckoutSettings{
		CheckoutRequired:          true,
		CheckoutDurationInSeconds: &updatedProjectCheckoutDuration,
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create testing
			{
				Config: createServerCheckoutSettingsCreateConfig(delegatedAdminGroupName, resourceGroupName, projectName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccServerCheckoutSettingsCheckExists(resourceName, initialServerCheckoutSettings),
				),
			},
			// Update testing
			{
				Config: createServerCheckoutSettingsUpdateConfig(delegatedAdminGroupName, resourceGroupName, projectName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccServerCheckoutSettingsCheckExists(resourceName, updatedServerCheckoutSettings),
				),
			},
			// Update testing with exclude list
			{
				Config: createServerCheckoutSettingsUpdateWithExcludeListConfig(delegatedAdminGroupName, resourceGroupName, projectName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccServerCheckoutSettingsCheckExists(resourceName, excludeListServerCheckoutSettings),
				),
			},
			// Update with both lists
			{
				Config:      createServerCheckoutSettingsUpdateWithBothListsConfig(delegatedAdminGroupName, resourceGroupName, projectName),
				ExpectError: regexp.MustCompile(`400 Bad Request`),
			},

			// Update project
			{
				Config: createServerCheckoutSettingsUpdateProjectConfig(delegatedAdminGroupName, resourceGroupName, projectName, updatedProjectName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccServerCheckoutSettingsCheckExists(resourceName, updatedProjectServerCheckoutSettings),
				),
			},
			// Delete testing
			{
				Config: createTestAccServerCheckoutSettingsDeleteConfig(delegatedAdminGroupName, resourceGroupName, projectName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccServerCheckoutSettingsCheckDeleted(resourceName),
				),
			},
		},
	})
}

func testAccServerCheckoutSettingsCheckExists(resourceName string, expectedServerCheckoutSettings *pam.ResourceCheckoutSettings) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Server checkout settings not found: %s", resourceName)
		}

		resourceGroupID := rs.Primary.Attributes[attributes.ResourceGroup]
		projectID := rs.Primary.Attributes[attributes.Project]

		client := getTestAccAPIClients().SDKClient
		serverCheckoutSettings, _, err := client.SDKClient.ProjectsAPI.FetchResourceGroupServerBasedProjectCheckoutSettings(context.Background(), client.Team, resourceGroupID, projectID).Execute()
		if err != nil {
			return fmt.Errorf("Error fetching server checkout settings: %s", err)
		}

		if serverCheckoutSettings.CheckoutRequired != expectedServerCheckoutSettings.CheckoutRequired {
			return fmt.Errorf("Server checkout settings checkout required does not match: %t != %t", serverCheckoutSettings.CheckoutRequired, expectedServerCheckoutSettings.CheckoutRequired)
		}

		if *serverCheckoutSettings.CheckoutDurationInSeconds != *expectedServerCheckoutSettings.CheckoutDurationInSeconds {
			return fmt.Errorf("Server checkout settings checkout duration in seconds does not match: %d != %d", *serverCheckoutSettings.CheckoutDurationInSeconds, *expectedServerCheckoutSettings.CheckoutDurationInSeconds)
		}

		if !reflect.DeepEqual(serverCheckoutSettings.IncludeList, expectedServerCheckoutSettings.IncludeList) {
			return fmt.Errorf("Server checkout settings include list does not match: %v != %v", serverCheckoutSettings.IncludeList, expectedServerCheckoutSettings.IncludeList)
		}
		if !reflect.DeepEqual(serverCheckoutSettings.ExcludeList, expectedServerCheckoutSettings.ExcludeList) {
			return fmt.Errorf("Server checkout settings exclude list does not match: %v != %v", serverCheckoutSettings.ExcludeList, expectedServerCheckoutSettings.ExcludeList)
		}
		return nil
	}
}

func testAccServerCheckoutSettingsCheckDeleted(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[resourceName]
		if ok {
			return fmt.Errorf("Server checkout settings still exists: %s", resourceName)
		}
		return nil
	}
}

func createServerCheckoutSettingsBaseConfig(delegatedAdminGroupName string, resourceGroupName string, projectName string) string {
	return fmt.Sprintf(testAccServerCheckoutSettingsBaseConfigFormat, delegatedAdminGroupName, resourceGroupName, projectName)
}

func createServerCheckoutSettingsCreateConfig(delegatedAdminGroupName string, resourceGroupName string, projectName string) string {
	return createServerCheckoutSettingsBaseConfig(delegatedAdminGroupName, resourceGroupName, projectName) + testAccServerCheckoutSettingsCreateConfigFormat
}

func createServerCheckoutSettingsUpdateConfig(delegatedAdminGroupName string, resourceGroupName string, projectName string) string {
	return createServerCheckoutSettingsBaseConfig(delegatedAdminGroupName, resourceGroupName, projectName) + testAccServerCheckoutSettingsUpdateConfigFormat
}

func createServerCheckoutSettingsUpdateWithExcludeListConfig(delegatedAdminGroupName string, resourceGroupName string, projectName string) string {
	return createServerCheckoutSettingsBaseConfig(delegatedAdminGroupName, resourceGroupName, projectName) + testAccServerCheckoutSettingsUpdateWithExcludeListConfigFormat
}

func createServerCheckoutSettingsUpdateWithBothListsConfig(delegatedAdminGroupName string, resourceGroupName string, projectName string) string {
	return createServerCheckoutSettingsBaseConfig(delegatedAdminGroupName, resourceGroupName, projectName) + testAccServerCheckoutSettingsUpdateWithBothListsConfigFormat
}

func createServerCheckoutSettingsUpdateProjectConfig(delegatedAdminGroupName string, resourceGroupName string, projectName string, updatedProjectName string) string {
	return fmt.Sprintf(createServerCheckoutSettingsBaseConfig(delegatedAdminGroupName, resourceGroupName, projectName)+testAccServerCheckoutSettingsUpdateProjectConfigFormat, updatedProjectName)
}
