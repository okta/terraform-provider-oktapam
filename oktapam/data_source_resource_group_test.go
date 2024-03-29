package oktapam

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
)

func TestAccDatasourceResourceGroupFetch(t *testing.T) {
	checkTeamApplicable(t, true)

	// Generate details
	identifier := randSeq()

	// config to create the resources
	initConfig := createTestAccDatasourceResourceGroupInitConfig(identifier)

	// config for the datasources
	fetchConfig := testAccDatasourceResourceGroupConfig("data1", identifier+"-1", "data2")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccResourceGroupsCheckDestroy(identifier+"-1", identifier+"-2"),
		Steps: []resource.TestStep{
			{
				Config: initConfig,
			},
			{
				Config: fmt.Sprintf("%s\n%s", initConfig, fetchConfig),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.oktapam_resource_group.data2", attributes.Name, fmt.Sprintf("%s-1", identifier)),
					resource.TestCheckResourceAttr("data.oktapam_resource_group.data2", attributes.Description, "terraform test rg"),
				),
			},
		},
	})
}

const testAccDatasourceResourceGroupCreateConfigFormat = `
resource "oktapam_group" "test_group" {
	name = "dra-test-group-%s"
}
resource "oktapam_resource_group" "test-resource-group-1" {
	name = "%s-1"
	description = "terraform test rg"
	delegated_resource_admin_groups = [oktapam_group.test_group.id]
}
`

func createTestAccDatasourceResourceGroupInitConfig(identifier string) string {
	return fmt.Sprintf(testAccDatasourceResourceGroupCreateConfigFormat, identifier, identifier)
}

const testAccDatasourceResourceGroupFormat = `
data "oktapam_resource_groups" "%s" {
	name = "%s"
}

data "oktapam_resource_group" "%s" {
	id = data.oktapam_resource_groups.%s.ids[0]
}
`

// "data1", identifier+"-1", "data2"
func testAccDatasourceResourceGroupConfig(groupsResourceName, name, groupResourceName string) string {
	return fmt.Sprintf(testAccDatasourceResourceGroupFormat, groupsResourceName, name, groupResourceName, groupsResourceName)
}
