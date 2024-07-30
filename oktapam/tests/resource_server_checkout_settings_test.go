package sdkv2

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
)

const testAccServerCheckoutSettingsCreateConfigFormat = `
provider "oktapam" {}
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
resource "oktapam_server_checkout_settings" "test_acc_server_checkout_settings" {
	resource_group = oktapam_resource_group.test_acc_resource_group.id
	project = oktapam_resource_group_project.test_acc_resource_group_project.id
	checkout_required = true
	checkout_duration_in_seconds = 900
}	
`

func createTestAccServerCheckoutSettingsCreateConfig(dgaName, resourceGroupName, projectName string) string {
	return fmt.Sprintf(testAccServerCheckoutSettingsCreateConfigFormat, dgaName, resourceGroupName, projectName)
}

const testAccServerCheckoutSettingsUpdateConfigFormat = `
provider "oktapam" {}
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
resource "oktapam_server_checkout_settings" "test_acc_server_checkout_settings" {
	resource_group = oktapam_resource_group.test_acc_resource_group.id
	project = oktapam_resource_group_project.test_acc_resource_group_project.id
	checkout_required = true
	checkout_duration_in_seconds = 3600
	include_list = ["vaulted_account_1", "vaulted_account_2"]
}	
`

func createTestAccServerCheckoutSettingsUpdateConfig(dgaName, resourceGroupName, projectName string) string {
	return fmt.Sprintf(testAccServerCheckoutSettingsUpdateConfigFormat, dgaName, resourceGroupName, projectName)
}

const testAccServerCheckoutSettingsDeleteConfigFormat = `
provider "oktapam" {}
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
func TestAccServerCheckoutSettingsSource(t *testing.T) {
	checkTeamApplicable(t, true)
	resourceName := "oktapam_server_checkout_settings.test_acc_server_checkout_settings"
	resourceGroupName := fmt.Sprintf("test_acc_resource_group_%s", randSeq())
	projectName := fmt.Sprintf("test_acc_resource_group_project_%s", randSeq())
	delegatedAdminGroupName := fmt.Sprintf("test_acc_resource_group_dga_%s", randSeq())
	defaultCheckoutDuration := int32(900)
	updatedCheckoutDuration := int32(3600)

	initialServerCheckoutSettings := &pam.ResourceCheckoutSettings{
		CheckoutRequired:          true,
		CheckoutDurationInSeconds: &defaultCheckoutDuration,
	}

	updatedServerCheckoutSettings := &pam.ResourceCheckoutSettings{
		CheckoutRequired:          true,
		CheckoutDurationInSeconds: &updatedCheckoutDuration,
		IncludeList:               []string{"vaulted_account_1", "vaulted_account_2"},
	}

	_, providers, accProviders := testAccFrameworkMuxProviders(context.Background(), t)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: accProviders,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: createTestAccServerCheckoutSettingsCreateConfig(delegatedAdminGroupName, resourceGroupName, projectName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccServerCheckoutSettingsCheckExists(providers, resourceName, initialServerCheckoutSettings),
				),
			},
			// Update testing
			{
				Config: createTestAccServerCheckoutSettingsUpdateConfig(delegatedAdminGroupName, resourceGroupName, projectName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccServerCheckoutSettingsCheckExists(providers, resourceName, updatedServerCheckoutSettings),
				),
			},
			// Delete testing
			{
				Config: createTestAccServerCheckoutSettingsDeleteConfig(delegatedAdminGroupName, resourceGroupName, projectName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccServerCheckoutSettingsCheckDeleted(providers, resourceName),
				),
			},
		},
	})
}

func testAccServerCheckoutSettingsCheckExists(providers *compositeDualProviderStruct, resourceName string, expectedServerCheckoutSettings *pam.ResourceCheckoutSettings) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Server checkout settings not found: %s", resourceName)
		}

		resourceGroupID := rs.Primary.Attributes[attributes.ResourceGroup]
		projectID := rs.Primary.Attributes[attributes.Project]

		client := providers.fwProvider.SDKClientWrapper
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

func testAccServerCheckoutSettingsCheckDeleted(providers *compositeDualProviderStruct, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[resourceName]
		if ok {
			return fmt.Errorf("Server checkout settings still exists: %s", resourceName)
		}
		return nil
	}
}
