package oktapam

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
)

func TestAccDatasourceResourceGroupsList(t *testing.T) {
	checkTeamApplicable(t, true)
	prefix := "data.oktapam_resource_groups"

	// Generate details
	identifier := randSeq()

	// config to create the resources
	initConfig := createTestAccDatasourceResourceGroupsInitConfig(identifier)

	// config for the datasources
	dataSource1Name := fmt.Sprintf("%s.%s", prefix, "data1")
	dataSource2Name := fmt.Sprintf("%s.%s", prefix, "data2")
	list1Config := testAccDatasourceResourceGroupsConfig("data1", identifier+"-1")
	list2Config := testAccDatasourceResourceGroupsConfig("data2", identifier+"-2")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccResourceGroupsCheckDestroy(identifier+"-1", identifier+"-2"),
		Steps: []resource.TestStep{
			{
				Config: initConfig,
			},
			{
				Config: fmt.Sprintf("%s\n%s", initConfig, list1Config),
				Check:  resource.TestCheckResourceAttr(dataSource1Name, fmt.Sprintf("%s.#", attributes.IDs), "1"),
			},
			{
				Config: fmt.Sprintf("%s\n%s", initConfig, list2Config),
				Check:  resource.TestCheckResourceAttr(dataSource2Name, fmt.Sprintf("%s.#", attributes.IDs), "1"),
			},
		},
	})
}

func testAccResourceGroupsCheckDestroy(identifiers ...string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		c := testAccProvider.Meta().(client.OktaPAMClient)

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

const testAccDatasourceResourceGroupsCreateConfigFormat = `
resource "oktapam_group" "test_group-1" {
	name = "dra-test-group-%s-1"
}
resource "oktapam_group" "test_group-2" {
	name = "dra-test-group-%s-2"
}
resource "oktapam_resource_group" "test-resource-group-1" {
	name = "%s-1"
	description = "terraform test rg"
	delegated_resource_admin_groups = [oktapam_group.test_group-1.id]
}

resource "oktapam_resource_group" "test-resource-group-2" {
	name = "%s-2"
	description = "terraform test rg"
	delegated_resource_admin_groups = [oktapam_group.test_group-2.id]
}
`

func createTestAccDatasourceResourceGroupsInitConfig(identifier string) string {
	return fmt.Sprintf(testAccDatasourceResourceGroupsCreateConfigFormat, identifier, identifier, identifier, identifier)
}

const testAccDatasourceResourceGroupsFormat = `
data "oktapam_resource_groups" "%s" {
	name = "%s"
}
`

func testAccDatasourceResourceGroupsConfig(resourceName, name string) string {
	return fmt.Sprintf(testAccDatasourceResourceGroupsFormat, resourceName, name)
}
