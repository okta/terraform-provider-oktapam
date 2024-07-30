package oktapam

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/kylelemons/godebug/pretty"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"
)

func TestAccResourceGroupProject(t *testing.T) {
	checkTeamApplicable(t, true)
	resourceName := "oktapam_resource_group_project.test_acc_resource_group_project"
	initialResourceGroupProjectName := fmt.Sprintf("test_acc_resource_group_project_%s", randSeq())
	resourceGroupName := fmt.Sprintf("test_acc_resource_group_%s", randSeq())
	delegatedAdminGroupName := fmt.Sprintf("test_acc_resource_group_dga_%s", randSeq())

	updatedResourceGroupProjectName := fmt.Sprintf("test_acc_resource_group_project_%s", randSeq())

	initialResourceGroupProject := &client.ResourceGroupProject{
		Name:               &initialResourceGroupProjectName,
		SSHCertificateType: utils.AsStringPtr("CERT_TYPE_ED25519_01"),
		AccountDiscovery:   utils.AsBoolPtrZero(false, true),
	}
	updatedResourceGroupProject := &client.ResourceGroupProject{
		Name:               &updatedResourceGroupProjectName,
		SSHCertificateType: utils.AsStringPtr("CERT_TYPE_RSA_01"),
		GatewaySelector:    utils.AsStringPtr("env=test"),
		AccountDiscovery:   utils.AsBoolPtrZero(true, true),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		// use the resource group check destroy since we create a new one here and deletion of the resource group will cascade delete the projects
		CheckDestroy: testAccResourceGroupCheckDestroy(resourceGroupName),
		Steps: []resource.TestStep{
			{
				Config: createTestAccResourceGroupProjectCreateConfig(delegatedAdminGroupName, resourceGroupName, initialResourceGroupProjectName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccResourceGroupProjectCheckExists(resourceName, initialResourceGroupProject),
					resource.TestCheckResourceAttr(
						resourceName, attributes.Name, initialResourceGroupProjectName,
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.SSHCertificateType, "CERT_TYPE_ED25519_01",
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.AccountDiscovery, "false",
					),
				),
			},
			{
				Config: createTestAccResourceGroupProjectUpdateConfig(delegatedAdminGroupName, resourceGroupName, updatedResourceGroupProjectName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccResourceGroupProjectCheckExists(resourceName, updatedResourceGroupProject),
					resource.TestCheckResourceAttr(
						resourceName, attributes.Name, updatedResourceGroupProjectName,
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.SSHCertificateType, "CERT_TYPE_RSA_01",
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.AccountDiscovery, "true",
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.GatewaySelector, "env=test",
					),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccResourceGroupProjectImportStateId(resourceName),
			},
		},
	})
}

func testAccResourceGroupProjectCheckExists(rn string, expectedResourceGroupProject *client.ResourceGroupProject) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[rn]
		if !ok {
			return fmt.Errorf("resource not found: %s", rn)
		}

		resourceGroupID := rs.Primary.Attributes[attributes.ResourceGroup]
		projectID := rs.Primary.Attributes[attributes.ID]
		pamClient := getLocalClientFromMetadata(testAccProvider.Meta())
		resourceGroupProject, err := pamClient.GetResourceGroupProject(context.Background(), resourceGroupID, projectID, false)
		if err != nil {
			return fmt.Errorf("error getting resource group project: %w", err)
		} else if resourceGroupProject == nil {
			return fmt.Errorf("resource group project does not exist")
		}

		err = insertComputedValuesForResourceGroupProject(expectedResourceGroupProject, resourceGroupProject)
		if err != nil {
			return err
		}
		comparison := pretty.Compare(expectedResourceGroupProject, resourceGroupProject)
		if comparison != "" {
			return fmt.Errorf("expected resource group project does not match returned resource group project.\n%s", comparison)
		}
		return nil
	}
}

func insertComputedValuesForResourceGroupProject(expectedResourceGroupProject, actualResourceGroupProject *client.ResourceGroupProject) error {
	actualResourceGroupProject.ID = expectedResourceGroupProject.ID
	actualResourceGroupProject.ResourceGroupID = expectedResourceGroupProject.ResourceGroupID
	actualResourceGroupProject.Team = expectedResourceGroupProject.Team

	return nil
}

const testAccResourceGroupProjectCreateConfigFormat = `
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
}
`

func createTestAccResourceGroupProjectCreateConfig(dgaName, resourceGroupName, projectName string) string {
	return fmt.Sprintf(testAccResourceGroupProjectCreateConfigFormat, dgaName, resourceGroupName, projectName)
}

const testAccResourceGroupProjectUpdateConfigFormat = `
resource "oktapam_group" "test_resource_group_dga_group" {
	name = "%s"
}
resource "oktapam_resource_group" "test_acc_resource_group" {
	name = "%s"
	description = "test resource group"
	delegated_resource_admin_groups = [oktapam_group.test_resource_group_dga_group.id]	
}
resource "oktapam_resource_group_project" "test_acc_resource_group_project" {
	name                 = "%s"
	resource_group       = oktapam_resource_group.test_acc_resource_group.id
	gateway_selector     = "env=test"
	ssh_certificate_type = "CERT_TYPE_RSA_01"
	account_discovery 	 = true
}
`

func createTestAccResourceGroupProjectUpdateConfig(dgaName, resourceGroupName, projectName string) string {
	return fmt.Sprintf(testAccResourceGroupProjectUpdateConfigFormat, dgaName, resourceGroupName, projectName)
}

func testAccResourceGroupProjectImportStateId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("Not found: %s", resourceName)
		}
		return fmt.Sprintf("%s/%s", rs.Primary.Attributes[attributes.ResourceGroup], rs.Primary.Attributes[attributes.ID]), nil
	}
}
