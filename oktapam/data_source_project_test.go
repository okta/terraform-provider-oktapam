package oktapam

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDatasourceProjectFetch(t *testing.T) {
	checkTeamApplicable(t, false)
	resourceName := "oktapam_project.test-1"
	dataSourceName := "data.oktapam_project.target"

	identifier := randSeq()

	testConfig := createTestAccDatasourceProjectInitConfig(identifier)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccV6ProviderFactories,
		CheckDestroy:             testAccProjectsCheckDestroy(identifier),
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

const testAccDatasourceProjectInitListFetchConfigFormat = `
resource "oktapam_project" "test-1" {
	name = "%s-1"
  	next_unix_uid = 60120
  	next_unix_gid = 63020
}

resource "oktapam_project" "test-2" {
	name = "%s-2"
  	next_unix_uid = 60220
  	next_unix_gid = 63120
}

data "oktapam_projects" "list" {
	depends_on = [oktapam_project.test-1]
	contains = "%s"
}

data "oktapam_project" "target" {
	depends_on = [data.oktapam_projects.list]
	name = data.oktapam_projects.list.names[0]
}
`

func createTestAccDatasourceProjectInitConfig(identifier string) string {
	return fmt.Sprintf(testAccDatasourceProjectInitListFetchConfigFormat, identifier, identifier, identifier)
}
