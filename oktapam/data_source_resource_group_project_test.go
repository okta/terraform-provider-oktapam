package oktapam

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
)

func TestAccDatasourceResourceGroupProjectFetch(t *testing.T) {
	checkTeamApplicable(t, true)

	// Generate details
	identifier := randSeq()

	// config to create the resources
	initConfig := createTestAccDatasourceResourceGroupProjectInitConfig(identifier)

	// config for the datasources
	fetchConfig := testAccDatasourceResourceGroupProjectConfig(identifier)

	resourceName := "data.oktapam_resource_group_project.rg_project"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccResourceGroupsCheckDestroy(identifier),
		Steps: []resource.TestStep{
			{
				Config: initConfig,
			},
			{
				Config: fmt.Sprintf("%s\n%s", initConfig, fetchConfig),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						resourceName, attributes.Name, identifier,
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.NextUnixUID, "60120",
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.NextUnixGID, "63020",
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.SSHCertificateType, "CERT_TYPE_ED25519_01",
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.AccountDiscovery, "false",
					),
				),
			},
		},
	})
}

const testAccDatasourceResourceGroupProjectCreateConfigFormat = `
resource "oktapam_group" "test_group" {
	name = "dra-test-group-%s"
}
resource "oktapam_resource_group" "test-resource-group" {
	name = "%s"
	description = "terraform test rg"
	delegated_resource_admin_groups = [oktapam_group.test_group.id]
}
resource "oktapam_resource_group_project" "test-resource-group-project" {
	name = "%s"
	resource_group =      oktapam_resource_group.test-resource-group.id
	next_unix_uid         = 60120
	next_unix_gid         = 63020
	ssh_certificate_type  = "CERT_TYPE_ED25519_01"
}
`

func createTestAccDatasourceResourceGroupProjectInitConfig(identifier string) string {
	return fmt.Sprintf(testAccDatasourceResourceGroupProjectCreateConfigFormat, identifier, identifier, identifier)
}

const testAccDatasourceResourceGroupProjectFormat = `
data "oktapam_resource_groups" "resource_groups" {
	name = "%s"
}

data "oktapam_resource_group_projects" "rg_projects" {
	resource_group = data.oktapam_resource_groups.resource_groups.ids[0]
	name = "%s"
}

data "oktapam_resource_group_project" "rg_project" {
	resource_group = data.oktapam_resource_groups.resource_groups.ids[0]
	id = data.oktapam_resource_group_projects.rg_projects.ids[0]
}
`

func testAccDatasourceResourceGroupProjectConfig(identifier string) string {
	return fmt.Sprintf(testAccDatasourceResourceGroupProjectFormat, identifier, identifier)
}
