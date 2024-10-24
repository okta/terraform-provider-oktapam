package oktapam

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDatasourceGroupList(t *testing.T) {
	prefix := "data.oktapam_groups"

	// Generate details
	identifier := randSeq()

	role := "access_user"
	if isExecutingPAMTest() {
		role = "end_user"
	}

	// Config 1: create two tokens
	initConfig := createTestAccDatasourceGroupsInitConfig(identifier, role)

	// Config 2: list using filter that returns both
	dataName := "data1"
	dataFullName := fmt.Sprintf("%s.%s", prefix, dataName)
	listConfig := testAccGroupsContainsConfig(dataName, identifier+"-1")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccV6ProviderFactories,
		CheckDestroy:             testAccGroupsCheckDestroy(identifier),
		Steps: []resource.TestStep{
			{
				Config: initConfig,
			},
			{
				Config: listConfig,
				Check:  resource.TestCheckResourceAttr(dataFullName, fmt.Sprintf("%s.#", attributes.Names), "1"),
			},
		},
	})
}

func testAccGroupsCheckDestroy(identifier string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		c := getTestAccAPIClients().LocalClient

		params := client.ListGroupsParameters{
			Contains: identifier,
		}

		groups, err := c.ListGroups(context.Background(), params)
		if err != nil {
			return fmt.Errorf("error getting tokens: %w", err)
		}

		if len(groups) > 0 {
			return fmt.Errorf("groups still exists")
		}

		return nil
	}
}

const testAccDatasourceGroupsInitConfigFormat = `
resource "oktapam_group" "test-group-1" {
	name = "%s-1"
	roles = ["%s"]
}

resource "oktapam_group" "test-group-2" {
	name = "%s-2"
	roles = ["%s"]
}
`

func createTestAccDatasourceGroupsInitConfig(identifier string, role string) string {
	return fmt.Sprintf(testAccDatasourceGroupsInitConfigFormat, identifier, role, identifier, role)
}

const testAccDatasourceGroupsContainsFormat = `
data "oktapam_groups" "%s" {
	contains = "%s"
}
`

func testAccGroupsContainsConfig(name, contains string) string {
	return fmt.Sprintf(testAccDatasourceGroupsContainsFormat, name, contains)
}
