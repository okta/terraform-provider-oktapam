package oktapam

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/kylelemons/godebug/pretty"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"
)

func TestAccResourceGroup(t *testing.T) {
	checkTeamApplicable(t, true)
	resourceName := "oktapam_resource_group.test_acc_resource_group"
	initialResourceGroupName := fmt.Sprintf("test_acc_resource_group_%s", randSeq())
	delegatedAdminGroup1Name := fmt.Sprintf("test_acc_resource_group_dga1_%s", randSeq())
	delegatedAdminGroup2Name := fmt.Sprintf("test_acc_resource_group_dga2_%s", randSeq())

	initialResourceGroup := &client.ResourceGroup{
		Name:        &initialResourceGroupName,
		Description: utils.AsStringPtr("initial description"),
		DelegatedResourceAdminGroups: []client.NamedObject{
			{
				Name: &delegatedAdminGroup1Name,
			},
			{
				Name: &delegatedAdminGroup2Name,
			},
		},
	}

	updatedResourceGroupName := fmt.Sprintf("test_acc_resource_group_%s", randSeq())
	updatedResourceGroup := &client.ResourceGroup{
		Name:        &updatedResourceGroupName,
		Description: utils.AsStringPtr("initial description"),
		DelegatedResourceAdminGroups: []client.NamedObject{
			{
				Name: &delegatedAdminGroup2Name,
			},
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccV6ProviderFactories(),
		CheckDestroy:             testAccResourceGroupCheckDestroy(initialResourceGroupName, updatedResourceGroupName),
		Steps: []resource.TestStep{
			{
				Config: createTestAccResourceGroupCreateConfig(delegatedAdminGroup1Name, delegatedAdminGroup2Name, initialResourceGroupName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccResourceGroupCheckExists(resourceName, initialResourceGroup),
					resource.TestCheckResourceAttr(
						resourceName, attributes.Name, initialResourceGroupName,
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.Description, *initialResourceGroup.Description,
					),
				),
			},
			{
				Config: createTestAccResourceGroupUpdateConfig(delegatedAdminGroup2Name, updatedResourceGroupName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccResourceGroupCheckExists(resourceName, updatedResourceGroup),
					resource.TestCheckResourceAttr(
						resourceName, attributes.Name, updatedResourceGroupName,
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.Description, *updatedResourceGroup.Description,
					),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccResourceGroupCheckExists(rn string, expectedResourceGroup *client.ResourceGroup) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[rn]
		if !ok {
			return fmt.Errorf("resource not found: %s", rn)
		}

		resourceGroupID := rs.Primary.Attributes[attributes.ID]
		//pamClient := getTestAccAPIClients().LocalClient
		pamClient := getTestAccAPIClients().LocalClient
		resourceGroup, err := pamClient.GetResourceGroup(context.Background(), resourceGroupID)
		if err != nil {
			return fmt.Errorf("error getting resource group: %w", err)
		} else if resourceGroup == nil {
			return fmt.Errorf("resource group does not exist")
		}

		err = insertComputedValuesForResourceGroup(expectedResourceGroup, resourceGroup)
		if err != nil {
			return err
		}
		comparison := pretty.Compare(expectedResourceGroup, resourceGroup)
		if comparison != "" {
			return fmt.Errorf("expected resource group does not match returned resource group.\n%s", comparison)
		}
		return nil
	}
}

func insertComputedValuesForResourceGroup(expectedResourceGroup, actualResourceGroup *client.ResourceGroup) error {
	actualResourceGroup.ID = expectedResourceGroup.ID
	actualResourceGroup.TeamID = expectedResourceGroup.TeamID
	if dgaSubs, err := subNamedObjects(expectedResourceGroup.DelegatedResourceAdminGroups, actualResourceGroup.DelegatedResourceAdminGroups, false); err == nil {
		actualResourceGroup.DelegatedResourceAdminGroups = dgaSubs
	} else {
		return err
	}

	return nil
}

func testAccResourceGroupCheckDestroy(resourceGroupNames ...string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		//client := getTestAccAPIClients().LocalClient
		client := getTestAccAPIClients().LocalClient
		resourceGroups, err := client.ListResourceGroups(context.Background())
		if err != nil {
			return fmt.Errorf("error getting resource groups: %w", err)
		}

		m := make(map[string]bool)
		for _, name := range resourceGroupNames {
			m[name] = true
		}

		for _, rg := range resourceGroups {
			if _, ok := m[*rg.Name]; ok {
				return fmt.Errorf("resource group still exists")
			}
		}

		return nil
	}
}

const testAccResourceGroupCreateConfigFormat = `
resource "oktapam_group" "test_resource_group_dga1_group" {
	name = "%s"
}
resource "oktapam_group" "test_resource_group_dga2_group" {
	name = "%s"
}
resource "oktapam_resource_group" "test_acc_resource_group" {
	name = "%s"
	description = "initial description"
	delegated_resource_admin_groups = [oktapam_group.test_resource_group_dga1_group.id, oktapam_group.test_resource_group_dga2_group.id]	
}
`

func createTestAccResourceGroupCreateConfig(dga1Name, dga2Name, resourceGroupName string) string {
	return fmt.Sprintf(testAccResourceGroupCreateConfigFormat, dga1Name, dga2Name, resourceGroupName)
}

const testAccResourceGroupUpdateConfigFormat = `
resource "oktapam_group" "test_resource_group_dga2_group" {
	name = "%s"
}
resource "oktapam_resource_group" "test_acc_resource_group" {
	name = "%s"
	description = "initial description"
	delegated_resource_admin_groups = [oktapam_group.test_resource_group_dga2_group.id]	
}
`

func createTestAccResourceGroupUpdateConfig(dga2Name, resourceGroupName string) string {
	return fmt.Sprintf(testAccResourceGroupUpdateConfigFormat, dga2Name, resourceGroupName)
}
