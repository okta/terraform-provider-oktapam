package oktapam

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDatasourceGroupFetch(t *testing.T) {
	resourceName := "oktapam_group.test-1"
	dataSourceName := "data.oktapam_group.target"

	identifier := randSeq()

	role := "access_user"
	if isExecutingPAMTest() {
		role = "end_user"
	}

	testConfig := createTestAccDatasourceGroupInitConfig(identifier, role)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccV6ProviderFactories,
		CheckDestroy:             testAccGroupsCheckDestroy(identifier),
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

const testAccDatasourceGroupInitListFetchConfigFormat = `
resource "oktapam_group" "test-1" {
	name = "%s-1"
	roles = ["%s"]
}

resource "oktapam_group" "test-2" {
	name = "%s-2"
	roles = ["%s"]
}

data "oktapam_groups" "list" {
	depends_on = [oktapam_group.test-1]
	contains = "%s"
}

data "oktapam_group" "target" {
	depends_on = [data.oktapam_groups.list]
	name = data.oktapam_groups.list.names[0]
}
`

func createTestAccDatasourceGroupInitConfig(identifier string, role string) string {
	return fmt.Sprintf(testAccDatasourceGroupInitListFetchConfigFormat, identifier, role, identifier, role, identifier)
}
