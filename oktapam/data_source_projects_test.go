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

func TestAccDatasourceProjectList(t *testing.T) {
	prefix := "data.oktapam_projects"

	// Generate details
	identifier := randSeq(10)

	// Config 1: create two tokens
	initConfig := createTestAccDatasourceProjectsInitConfig(identifier)

	// Config 2: list using filter that returns both
	dataName := "data1"
	dataFullName1 := fmt.Sprintf("%s.%s", prefix, dataName)
	listConfig := testAccProjectsContainsConfig(dataName, identifier+"-1")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccProjectsCheckDestroy(identifier),
		Steps: []resource.TestStep{
			{
				Config: initConfig,
			},
			{
				Config: listConfig,
				Check:  resource.TestCheckResourceAttr(dataFullName1, fmt.Sprintf("%s.#", attributes.Names), "1"),
			},
		},
	})
}

func testAccProjectsCheckDestroy(identifier string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		c := testAccProvider.Meta().(client.OktaPAMClient)

		params := client.ListProjectsParameters{
			Contains: identifier,
		}

		projects, err := c.ListProjects(context.Background(), params)
		if err != nil {
			return fmt.Errorf("error getting projects: %w", err)
		}

		if len(projects) > 0 {
			return fmt.Errorf("projects still exists")
		}

		return nil
	}
}

const testAccDatasourceProjectsCreateConfigFormat = `
resource "oktapam_project" "test-project-1" {
	name = "%s-1"
  	next_unix_uid = 60120
  	next_unix_gid = 63020
}

resource "oktapam_project" "test-project-2" {
	name = "%s-2"
  	next_unix_uid = 60220
  	next_unix_gid = 63120
}
`

func createTestAccDatasourceProjectsInitConfig(identifier string) string {
	return fmt.Sprintf(testAccDatasourceProjectsCreateConfigFormat, identifier, identifier)
}

const testAccDatasourceProjectsContainsFormat = `
data "oktapam_projects" "%s" {
	contains = "%s"
}
`

func testAccProjectsContainsConfig(name, contains string) string {
	return fmt.Sprintf(testAccDatasourceProjectsContainsFormat, name, contains)
}
