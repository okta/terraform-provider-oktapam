package oktapam

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDatasourceProjectGroupFetch(t *testing.T) {
	checkTeamApplicable(t, false)
	resourceName := "oktapam_project_group.test-1"
	dataSourceName := "data.oktapam_project_group.target"

	identifier := randSeq()

	testConfig := createTestAccDatasourceProjectGroupInitConfig(identifier)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccV6ProviderFactories,
		CheckDestroy:             testAccProjectGroupsCheckDestroy(identifier + "-1"),
		Steps: []resource.TestStep{
			{
				Config: testConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					checkResourcesEqual(resourceName, dataSourceName),
				),
			},
		},
	})
}

// NOTE: This config (1) creates two new resources (2) lists the existing resources with the matching identifier
// and (3) get the new resource as a data source.
// The test then compares the resource with its data source to ensure they are equal.

const testAccDatasourceProjectGroupInitListFetchConfigFormat = `
resource "oktapam_project" "test-1" {
	name = "%s-1"
  	next_unix_uid = 60120
  	next_unix_gid = 63020
}

resource "oktapam_group" "test-1" {
	name = "%s-1"
	roles = ["access_user"]
}

resource "oktapam_group" "test-2" {
	name = "%s-2"
	roles = ["access_user"]
}

resource "oktapam_project_group" "test-1" {
	project_name = oktapam_project.test-1.name
	group_name = oktapam_group.test-1.name
	server_access = true
	server_admin = true
	create_server_group = true
}

resource "oktapam_project_group" "test-2" {
	project_name = oktapam_project.test-1.name
	group_name = oktapam_group.test-2.name
	server_access = true
	server_admin = true
	create_server_group = false
}

data "oktapam_project_groups" "list" {
	depends_on = [oktapam_project_group.test-1]
	project_name = oktapam_project.test-1.name
	create_server_group = true
}

data "oktapam_project_group" "target" {
	depends_on = [data.oktapam_project_groups.list]
	project_name = oktapam_project.test-1.name
	group_name = data.oktapam_project_groups.list.group_names[0]
}
`

func createTestAccDatasourceProjectGroupInitConfig(identifier string) string {
	return fmt.Sprintf(testAccDatasourceProjectGroupInitListFetchConfigFormat, identifier, identifier, identifier)
}
