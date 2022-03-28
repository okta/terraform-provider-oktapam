package oktapam

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/client"
	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/utils"
)

func TestAccDatasourceProjectGroup(t *testing.T) {
	resourceName := "data.oktapam_project_group.test_project_groups"
	projectName := fmt.Sprintf("test-acc-project-group-project-%s", randSeq(10))
	group1Name := fmt.Sprintf("test-acc-project-group-group-1-%s", randSeq(10))
	group2Name := fmt.Sprintf("test-acc-project-group-group-2-%s", randSeq(10))

	expectedGroups := map[string]client.ProjectGroup{
		group1Name: {
			Project:          utils.AsStringPtr(projectName),
			Group:            utils.AsStringPtr(group1Name),
			ServerAccess:     true,
			ServerAdmin:      true,
			CreateServerGoup: false,
		},
		group2Name: {
			Project:          utils.AsStringPtr(projectName),
			Group:            utils.AsStringPtr(group2Name),
			ServerAccess:     true,
			ServerAdmin:      false,
			CreateServerGoup: false,
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: createTestAccDatasourceProjectGroupInitConfig(projectName, group1Name, group2Name),
			},
			{
				Config: testAccDatasourceProjectGroupConfig(projectName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "project_groups.#", "2"),
					testAccDatasourceProjectGroupsCheck(resourceName, expectedGroups),
				),
			},
		},
	})
}
func testAccDatasourceProjectGroupsCheck(rn string, expectedProjects map[string]client.ProjectGroup) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[rn]
		if !ok {
			return fmt.Errorf("resource not found: %s", rn)
		}
		mappings, err := getIndexMappingFromResource(rs, "project_groups", "group_name", len(expectedProjects))
		if err != nil {
			return fmt.Errorf("error mapping resources to indices: %w", err)
		}
		attributes := rs.Primary.Attributes
		for name, projectGroup := range expectedProjects {
			// tests some attributes to ensure we are obtaining some attributes that were set by the original create resource and some that were computed
			idx, ok := mappings[name]
			if !ok {
				return fmt.Errorf("could not find resource with name: %s", name)
			}
			projectName, ok := attributes[fmt.Sprintf("project_groups.%s.project_name", idx)]
			if !ok {
				return fmt.Errorf("project_name attribute not set for project group with group %q", name)
			}
			if projectName != *projectGroup.Project {
				return fmt.Errorf("mismatch for project_name value, expected %q, got %q", *projectGroup.Project, projectName)
			}

			groupName, ok := attributes[fmt.Sprintf("project_groups.%s.group_name", idx)]
			if !ok {
				return fmt.Errorf("group_name attribute not set for project group with group %q", name)
			}
			if groupName != *projectGroup.Group {
				return fmt.Errorf("mismatch for group_name value, expected %q, got %q", *projectGroup.Group, groupName)
			}

			serverAccess, ok := attributes[fmt.Sprintf("project_groups.%s.server_access", idx)]
			if !ok {
				return fmt.Errorf("server_access attribute not set for project group with group %q", name)
			}
			if serverAccess != fmt.Sprint(projectGroup.ServerAccess) {
				return fmt.Errorf("mismatch for server_access attribute, expected %q, got %q", fmt.Sprint(projectGroup.ServerAccess), serverAccess)
			}

			serverAdmin, ok := attributes[fmt.Sprintf("project_groups.%s.server_admin", idx)]
			if !ok {
				return fmt.Errorf("server_admin attribute not set for project group with group %q", name)
			}
			if serverAdmin != fmt.Sprint(projectGroup.ServerAdmin) {
				return fmt.Errorf("mismatch for server_admin attribute, expected %q, got %q", fmt.Sprint(projectGroup.ServerAdmin), serverAdmin)
			}

			createServerGroup, ok := attributes[fmt.Sprintf("project_groups.%s.create_server_group", idx)]
			if !ok {
				return fmt.Errorf("create_server_group attribute not set for project group with group %q", name)
			}
			if createServerGroup != fmt.Sprint(projectGroup.CreateServerGoup) {
				return fmt.Errorf("mismatch for create_server_group attribute, expected %q, got %q", fmt.Sprint(projectGroup.CreateServerGoup), createServerGroup)
			}

		}

		return nil
	}
}

const testAccDatasourceProjectGroupFormat = `
data "oktapam_project_group" "test_project_groups" {
	project_name = "%s"
}
`

func testAccDatasourceProjectGroupConfig(projectName string) string {
	return fmt.Sprintf(testAccDatasourceProjectGroupFormat, projectName)
}

const testAccDatasourceProjectGroupInitConfigFormat = `
resource "oktapam_project" "test-project-group-project" {
	name = "%s"
	next_unix_uid = 60120
	next_unix_gid = 63020
}
resource "oktapam_group" "test-project-group-group-1" {
	name = "%s"
}
resource "oktapam_group" "test-project-group-group-2" {
	name = "%s"
}
resource "oktapam_project_group" "test-acc-project-group-project-group-1" {
	project_name = oktapam_project.test-project-group-project.name
	group_name = oktapam_group.test-project-group-group-1.name
	server_access = true
	server_admin = true
	create_server_group = false
}
resource "oktapam_project_group" "test-acc-project-group-project-group-2" {
	project_name = oktapam_project.test-project-group-project.name
	group_name = oktapam_group.test-project-group-group-2.name
	server_access = true
	server_admin = false
	create_server_group = false
}
`

func createTestAccDatasourceProjectGroupInitConfig(projectName, group1Name, group2Name string) string {
	return fmt.Sprintf(testAccDatasourceProjectGroupInitConfigFormat, projectName, group1Name, group2Name)
}
