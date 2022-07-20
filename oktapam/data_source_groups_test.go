package oktapam

import (
	"fmt"
	"testing"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDatasourceGroupList(t *testing.T) {
	prefix := "data.oktapam_groups"

	// Generate details
	identifier := randSeq(10)

	// Config 1: create two tokens
	initConfig := createTestAccDatasourceGroupsInitConfig(identifier)

	// Config 2: list using filter that returns both
	dataName := "data1"
	dataFullName := fmt.Sprintf("%s.%s", prefix, dataName)
	listConfig := testAccGroupsContainsConfig(dataName, identifier)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccGatewaySetupTokenCheckDestroy(identifier),
		Steps: []resource.TestStep{
			{
				Config: initConfig,
			},
			{
				Config: listConfig,
				Check:  resource.TestCheckResourceAttr(dataFullName, fmt.Sprintf("%s.#", attributes.Names), "2"),
			},
		},
	})
}

const testAccDatasourceGroupsInitConfigFormat = `
resource "oktapam_group" "test-group-1" {
	name = "%s-1"
	roles = ["access_user"]
}

resource "oktapam_group" "test-group-2" {
	name = "%s-2"
	roles = ["access_user"]
}
`

func createTestAccDatasourceGroupsInitConfig(identifier string) string {
	return fmt.Sprintf(testAccDatasourceGroupsInitConfigFormat, identifier, identifier)
}

const testAccDatasourceGroupsContainsFormat = `
data "oktapam_groups" "%s" {
	contains = "%s"
}
`

func testAccGroupsContainsConfig(name, contains string) string {
	return fmt.Sprintf(testAccDatasourceGroupsContainsFormat, name, contains)
}
