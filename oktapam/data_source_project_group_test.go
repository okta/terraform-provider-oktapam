package oktapam

import (
	"fmt"
	"testing"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"
)

func TestAccDatasourceProjectGroup(t *testing.T) {
	resourceName := "data.oktapam_project_group.test_project_groups"
	projectName := fmt.Sprintf("test-acc-project-group-project-%s", randSeq(10))
	group1Name := fmt.Sprintf("test-acc-project-group-group-1-%s", randSeq(10))
	group2Name := fmt.Sprintf("test-acc-project-group-group-2-%s", randSeq(10))

	expectedGroups := map[string]client.ProjectGroup{
		group1Name: {
			Project:           utils.AsStringPtr(projectName),
			Group:             utils.AsStringPtr(group1Name),
			ServerAccess:      true,
			ServerAdmin:       true,
			CreateServerGroup: false,
		},
		group2Name: {
			Project:           utils.AsStringPtr(projectName),
			Group:             utils.AsStringPtr(group2Name),
			ServerAccess:      true,
			ServerAdmin:       false,
			CreateServerGroup: false,
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: createTestAccDatasourceProjectGroupInitConfig(projectName, group1Name, group2Name),
			},
			{
				Config: testAccDatasourceProjectGroupConfig(projectName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.#", attributes.ProjectGroups), "2"),
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
		mappings, err := getIndexMappingFromResource(rs, attributes.ProjectGroups, attributes.GroupName, len(expectedProjects))
		if err != nil {
			return fmt.Errorf("error mapping resources to indices: %w", err)
		}
		primaryAttributes := rs.Primary.Attributes
		for name, projectGroup := range expectedProjects {
			// tests some attributes to ensure we are obtaining some attributes that were set by the original create resource and some that were computed
			idx, ok := mappings[name]
			if !ok {
				return fmt.Errorf("could not find resource with name: %s", name)
			}
			projectName, ok := primaryAttributes[fmt.Sprintf("%s.%s.%s", attributes.ProjectGroups, idx, attributes.ProjectName)]
			if !ok {
				return fmt.Errorf("%s attribute not set for project group with group %q", attributes.ProjectName, name)
			}
			if projectName != *projectGroup.Project {
				return fmt.Errorf("mismatch for %s value, expected %q, got %q", attributes.ProjectName, *projectGroup.Project, projectName)
			}

			groupName, ok := primaryAttributes[fmt.Sprintf("%s.%s.%s", attributes.ProjectGroups, idx, attributes.GroupName)]
			if !ok {
				return fmt.Errorf("%s attribute not set for project group with group %q", attributes.GroupName, name)
			}
			if groupName != *projectGroup.Group {
				return fmt.Errorf("mismatch for %s value, expected %q, got %q", attributes.GroupName, *projectGroup.Group, groupName)
			}

			serverAccess, ok := primaryAttributes[fmt.Sprintf("%s.%s.%s", attributes.ProjectGroups, idx, attributes.ServerAccess)]
			if !ok {
				return fmt.Errorf("%s attribute not set for project group with group %q", attributes.ServerAccess, name)
			}
			if serverAccess != fmt.Sprint(projectGroup.ServerAccess) {
				return fmt.Errorf("mismatch for %s attribute, expected %q, got %q", attributes.ServerAccess, fmt.Sprint(projectGroup.ServerAccess), serverAccess)
			}

			serverAdmin, ok := primaryAttributes[fmt.Sprintf("%s.%s.%s", attributes.ProjectGroups, idx, attributes.ServerAdmin)]
			if !ok {
				return fmt.Errorf("%s attribute not set for project group with group %q", attributes.ServerAccess, name)
			}
			if serverAdmin != fmt.Sprint(projectGroup.ServerAdmin) {
				return fmt.Errorf("mismatch for %s attribute, expected %q, got %q", attributes.ServerAccess, fmt.Sprint(projectGroup.ServerAdmin), serverAdmin)
			}

			createServerGroup, ok := primaryAttributes[fmt.Sprintf("%s.%s.%s", attributes.ProjectGroups, idx, attributes.CreateServerGroup)]
			if !ok {
				return fmt.Errorf("%s attribute not set for project group with group %q", attributes.CreateServerGroup, name)
			}
			if createServerGroup != fmt.Sprint(projectGroup.CreateServerGroup) {
				return fmt.Errorf("mismatch for %s attribute, expected %q, got %q", attributes.CreateServerGroup, fmt.Sprint(projectGroup.CreateServerGroup), createServerGroup)
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

func createTestAccDatasourceProjectGroupInitConfig(projectName string, group1Name string, group2Name string) string {
	return fmt.Sprintf(testAccDatasourceProjectGroupInitConfigFormat, projectName, group1Name, group2Name)
}
