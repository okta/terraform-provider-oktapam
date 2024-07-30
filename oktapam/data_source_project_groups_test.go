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

func TestAccDatasourceProjectGroupList(t *testing.T) {
	checkTeamApplicable(t, false)
	prefix := "data.oktapam_project_groups"

	// Generate details
	identifier := randSeq()

	// Config 1: create two project groups
	initConfig := createTestAccDatasourceProjectGroupsInitConfig(identifier)

	// Config 2: list using filter that returns both
	dataName := "data1"
	dataFullName1 := fmt.Sprintf("%s.%s", prefix, dataName)
	listConfig := testAccProjectGroupsConfig(dataName, identifier+"-1")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccProjectGroupsCheckDestroy(identifier + "-1"),
		Steps: []resource.TestStep{
			{
				Config: initConfig,
			},
			{
				Config: listConfig,
				Check:  resource.TestCheckResourceAttr(dataFullName1, fmt.Sprintf("%s.#", attributes.GroupNames), "2"),
			},
		},
	})
}

func testAccProjectGroupsCheckDestroy(projectName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		c := getLocalClientFromMetadata(testAccProvider.Meta())

		params := client.ListProjectGroupsParameters{}

		projects, err := c.ListProjectGroups(context.Background(), projectName, params)
		if err != nil {
			return fmt.Errorf("error getting projects: %w", err)
		}

		if len(projects) > 0 {
			return fmt.Errorf("projects still exists")
		}

		return nil
	}
}

const testAccDatasourceProjectGroupsCreateConfigFormat = `
resource "oktapam_project" "test-project-1" {
	name = "%s-1"
  	next_unix_uid = 60120
  	next_unix_gid = 63020
}

resource "oktapam_group" "test-group-1" {
	name = "%s-1"
	roles = ["access_user"]
}

resource "oktapam_group" "test-group-2" {
	name = "%s-2"
	roles = ["access_user"]
}

resource "oktapam_project_group" "test-project-group-1" {
	project_name = oktapam_project.test-project-1.name
	group_name = oktapam_group.test-group-1.name
	server_access = true
	server_admin = true
	create_server_group = false
}

resource "oktapam_project_group" "test-project-group-2" {
	project_name = oktapam_project.test-project-1.name
	group_name = oktapam_group.test-group-2.name
	server_access = true
	server_admin = true
	create_server_group = false
}
`

func createTestAccDatasourceProjectGroupsInitConfig(identifier string) string {
	return fmt.Sprintf(testAccDatasourceProjectGroupsCreateConfigFormat, identifier, identifier, identifier)
}

const testAccDatasourceProjectGroupsFormat = `
data "oktapam_project_groups" "%s" {
	project_name = "%s"
}
`

func testAccProjectGroupsConfig(name, projectName string) string {
	return fmt.Sprintf(testAccDatasourceProjectGroupsFormat, name, projectName)
}
