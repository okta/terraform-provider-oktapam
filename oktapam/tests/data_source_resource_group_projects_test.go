package sdkv2

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
)

func TestAccDatasourceResourceGroupProjectsList(t *testing.T) {
	checkTeamApplicable(t, true)

	// Generate details
	identifier := randSeq()

	// config to create the resources
	initConfig := createTestAccDatasourceResourceGroupProjectsInitConfig(identifier)

	// config for the datasources
	list1Config := testAccDatasourceResourceGroupProjectsConfig(identifier, identifier+"-1")
	list2Config := testAccDatasourceResourceGroupProjectsConfig(identifier, identifier+"-2")
	listAllConfig := testAccDatasourceResourceGroupProjectsConfig(identifier, "")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccResourceGroupProjectsCheckDestroy(identifier),
		Steps: []resource.TestStep{
			{
				Config: initConfig,
			},
			{
				Config: fmt.Sprintf("%s\n%s", initConfig, list1Config),
				Check:  resource.TestCheckResourceAttr("data.oktapam_resource_group_projects.resource_group_projects", fmt.Sprintf("%s.#", attributes.IDs), "1"),
			},
			{
				Config: fmt.Sprintf("%s\n%s", initConfig, list2Config),
				Check:  resource.TestCheckResourceAttr("data.oktapam_resource_group_projects.resource_group_projects", fmt.Sprintf("%s.#", attributes.IDs), "1"),
			},
			{
				Config: fmt.Sprintf("%s\n%s", initConfig, listAllConfig),
				Check:  resource.TestCheckResourceAttr("data.oktapam_resource_group_projects.resource_group_projects", fmt.Sprintf("%s.#", attributes.IDs), "2"),
			},
		},
	})
}

func testAccResourceGroupProjectsCheckDestroy(identifiers ...string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		c := client.GetLocalClientFromMetadata(testAccProvider.Meta())

		resourceGroups, err := c.ListResourceGroups(context.Background())
		if err != nil {
			return fmt.Errorf("error getting resource groups: %w", err)
		}

		m := make(map[string]bool, len(identifiers))
		for _, id := range identifiers {
			m[id] = true
		}

		for _, rg := range resourceGroups {
			if _, ok := m[*rg.Name]; ok {
				return fmt.Errorf("resource groups still exists")
			}
		}

		return nil
	}
}

const testAccDatasourceResourceGroupProjectsCreateConfigFormat = `
resource "oktapam_group" "test_group" {
	name = "dra-test-group-%s"
}
resource "oktapam_resource_group" "test-resource-group" {
	name = "%s"
	description = "terraform test rg"
	delegated_resource_admin_groups = [oktapam_group.test_group.id]
}
resource "oktapam_resource_group_project" "test-resource-group-project-1" {
	name = "%s-1"
	resource_group =      oktapam_resource_group.test-resource-group.id
	ssh_certificate_type  = "CERT_TYPE_ED25519_01"
}
resource "oktapam_resource_group_project" "test-resource-group-project-2" {
	name = "%s-2"
	resource_group       = oktapam_resource_group.test-resource-group.id
	gateway_selector     = "env=test"
	ssh_certificate_type = "CERT_TYPE_RSA_01"
	account_discovery 	 = true
}
`

func createTestAccDatasourceResourceGroupProjectsInitConfig(identifier string) string {
	return fmt.Sprintf(testAccDatasourceResourceGroupProjectsCreateConfigFormat, identifier, identifier, identifier, identifier)
}

const testAccDatasourceResourceGroupProjectsFormat = `
data "oktapam_resource_groups" "resource_groups" {
	name = "%s"
}
data "oktapam_resource_group_projects" "resource_group_projects" {
	resource_group = data.oktapam_resource_groups.resource_groups.ids[0]
	name = "%s"
}
`

func testAccDatasourceResourceGroupProjectsConfig(resourceGroupName, projectNameContains string) string {
	return fmt.Sprintf(testAccDatasourceResourceGroupProjectsFormat, resourceGroupName, projectNameContains)
}
