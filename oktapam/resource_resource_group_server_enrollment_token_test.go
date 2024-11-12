package oktapam

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"
)

func TestAccResourceGroupServerEnrollmentToken(t *testing.T) {
	checkTeamApplicable(t, true)
	resourceName := "oktapam_resource_group_server_enrollment_token.test_acc_resource_group_server_enrollment_token"
	resourceGroupProjectName := fmt.Sprintf("test_acc_resource_group_project_%s", randSeq())
	resourceGroupName := fmt.Sprintf("test_acc_resource_group_%s", randSeq())
	delegatedAdminGroupName := fmt.Sprintf("test_acc_resource_group_dga_%s", randSeq())
	tokenDescription := "test description"

	serverEnrollmentToken := &client.ResourceGroupServerEnrollmentToken{
		Description: &tokenDescription,
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccV6ProviderFactories(),
		// use the resource group check destroy since we create a new one here and deletion of the resource group will cascade delete the projects / tokens
		CheckDestroy: testAccResourceGroupCheckDestroy(resourceGroupName),
		Steps: []resource.TestStep{
			{
				Config: createTestAccResourceGroupServerEnrollmentTokenCreateConfig(delegatedAdminGroupName, resourceGroupName, resourceGroupProjectName, tokenDescription),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccResourceGroupServerEnrollmentTokenCheckExists(resourceName, serverEnrollmentToken),
					resource.TestCheckResourceAttr(
						resourceName, attributes.Description, "test description",
					),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccResourceGroupServerEnrollmentTokenImportStateId(resourceName),
			},
		},
	})
}

func testAccResourceGroupServerEnrollmentTokenCheckExists(rn string, expectedResourceGroupServerEnrollmentToken *client.ResourceGroupServerEnrollmentToken) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[rn]
		if !ok {
			return fmt.Errorf("resource not found: %s", rn)
		}

		resourceGroupID := rs.Primary.Attributes[attributes.ResourceGroup]
		projectID := rs.Primary.Attributes[attributes.Project]
		serverEnrollmentTokenID := rs.Primary.Attributes[attributes.ID]
		pamClient := mustTestAccAPIClients().LocalClient
		resourceGroupServerEnrollmentToken, err := pamClient.GetResourceGroupServerEnrollmentToken(context.Background(), resourceGroupID, projectID, serverEnrollmentTokenID)
		if err != nil {
			return fmt.Errorf("error getting resource group server enrollment token: %w", err)
		} else if resourceGroupServerEnrollmentToken == nil {
			return fmt.Errorf("resource group server enrollment token does not exist")
		} else if utils.IsBlank(resourceGroupServerEnrollmentToken.ID) {
			return fmt.Errorf("server enrollment token for project %s with id %s does not exist", projectID, serverEnrollmentTokenID)
		}
		if *resourceGroupServerEnrollmentToken.Description != *expectedResourceGroupServerEnrollmentToken.Description {
			return fmt.Errorf("expected description does not match returned description for server enrollment token.  expected: %s, got: %s", *expectedResourceGroupServerEnrollmentToken.Description, *resourceGroupServerEnrollmentToken.Description)
		}
		return nil
	}
}

func testAccResourceGroupServerEnrollmentTokenCheckDestroy(resourceGroupName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		pamClient := mustTestAccAPIClients().LocalClient
		resourceGroups, err := pamClient.ListResourceGroups(context.Background())
		if err != nil {
			return fmt.Errorf("error getting resource groups: %w", err)
		}

		for _, rg := range resourceGroups {
			if *rg.Name == resourceGroupName {
				return fmt.Errorf("resource group still exists")
			}
		}

		return nil
	}
}

const testAccResourceGroupServerEnrollmentTokenCreateConfigFormat = `
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
resource "oktapam_resource_group_server_enrollment_token" "test_acc_resource_group_server_enrollment_token" {
	resource_group = oktapam_resource_group.test_acc_resource_group.id
	project        = oktapam_resource_group_project.test_acc_resource_group_project.id
	description    = "%s"
}
`

func createTestAccResourceGroupServerEnrollmentTokenCreateConfig(dgaName, resourceGroupName, projectName, tokenDescription string) string {
	return fmt.Sprintf(testAccResourceGroupServerEnrollmentTokenCreateConfigFormat, dgaName, resourceGroupName, projectName, tokenDescription)
}

func testAccResourceGroupServerEnrollmentTokenImportStateId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("Not found: %s", resourceName)
		}
		return fmt.Sprintf("%s/%s/%s", rs.Primary.Attributes[attributes.ResourceGroup], rs.Primary.Attributes[attributes.Project], rs.Primary.ID), nil
	}
}
