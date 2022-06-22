package oktapam

import (
	"context"
	"fmt"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/kylelemons/godebug/pretty"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"
)

func TestAccProjectGroup(t *testing.T) {
	resourceName := "oktapam_project_group.test-acc-project-group"
	projectName := fmt.Sprintf("test-acc-project-group-project-%s", randSeq(10))
	groupName := fmt.Sprintf("test-acc-project-group-group-%s", randSeq(10))

	initialProjectGroup := client.ProjectGroup{
		Project:      &projectName,
		Group:        &groupName,
		ServerAccess: true,
		ServerAdmin:  true,
	}
	updatedProjectGroup := client.ProjectGroup{
		Project:          &projectName,
		Group:            &groupName,
		ServerAccess:     true,
		ServerAdmin:      false,
		CreateServerGoup: true,
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		ProviderFactories:    testAccProviders,
		CheckDestroy: testAccProjectGroupCheckDestroy(initialProjectGroup),
		Steps: []resource.TestStep{
			{
				Config: createTestAccProjectGroupCreateConfig(initialProjectGroup),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectGroupCheckExists(resourceName, initialProjectGroup),
					resource.TestCheckResourceAttr(
						resourceName, attributes.ProjectName, projectName,
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.GroupName, groupName,
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.ServerAccess, "true",
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.ServerAdmin, "true",
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.CreateServerGroup, "false",
					),
				),
			},
			{
				Config: createTestAccProjectGroupUpdateConfig(updatedProjectGroup),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectGroupCheckExists(resourceName, updatedProjectGroup),
					resource.TestCheckResourceAttr(
						resourceName, attributes.ProjectName, projectName,
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.GroupName, groupName,
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.ServerAccess, "true",
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.ServerAdmin, "false",
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.CreateServerGroup, "true",
					),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccProjectGroupCheckExists(rn string, expectedProjectGroup client.ProjectGroup) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[rn]
		if !ok {
			return fmt.Errorf("resource not found: %s", rn)
		}

		resourceID := rs.Primary.ID
		project, group, err := parseProjectGroupResourceID(resourceID)
		if err != nil {
			return fmt.Errorf("error parsing resource id: %w", err)
		}

		client := testAccProvider.Meta().(client.OktaPAMClient)
		projectGroup, err := client.GetProjectGroup(context.Background(), project, group)
		if err != nil {
			return fmt.Errorf("error getting project group: %w", err)
		} else if projectGroup == nil {
			return fmt.Errorf("project group does not exist")
		}
		expectedProjectGroup.GroupID = projectGroup.GroupID

		comparison := pretty.Compare(expectedProjectGroup, projectGroup)
		if comparison != "" {
			return fmt.Errorf("expected project group does not match returned project.\n%s", comparison)
		}
		return nil
	}
}

func testAccProjectGroupCheckDestroy(projectGroup client.ProjectGroup) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(client.OktaPAMClient)
		pg, err := client.GetProjectGroup(context.Background(), *projectGroup.Project, *projectGroup.Group)
		if err != nil {
			return fmt.Errorf("error getting project group: %w", err)
		}
		if pg != nil && utils.IsNonEmpty(pg.Project) && utils.IsNonEmpty(pg.Group) && utils.IsBlank(pg.DeletedAt) && utils.IsBlank(pg.RemovedAt) {
			return fmt.Errorf("project group still exists")
		}
		group, err := client.GetGroup(context.Background(), *projectGroup.Group, false)
		if err != nil {
			return fmt.Errorf("error getting group associated with project group: %w", err)
		}
		if group != nil && utils.IsNonEmpty(group.ID) {
			return fmt.Errorf("group still exists: %s", *projectGroup.Group)
		}
		project, err := client.GetProject(context.Background(), *projectGroup.Project, false)
		if err != nil {
			return fmt.Errorf("error getting project associated with project group: %w", err)
		}
		if project != nil && utils.IsNonEmpty(project.ID) {
			return fmt.Errorf("project still exists: %s", *projectGroup.Project)
		}
		return nil
	}
}

const testAccProjectGroupCreateConfigFormat = `
resource "oktapam_project" "test-project-group-project" {
    name = "%s"
  	next_unix_uid = 60120
  	next_unix_gid = 63020
}
resource "oktapam_group" "test-project-group-group" {
    name = "%s"
}
resource "oktapam_project_group" "test-acc-project-group" {
    project_name = oktapam_project.test-project-group-project.name
  	group_name = oktapam_group.test-project-group-group.name
	server_access = true
	server_admin = true
	create_server_group = false
}`

func createTestAccProjectGroupCreateConfig(projectGroup client.ProjectGroup) string {
	return fmt.Sprintf(testAccProjectGroupCreateConfigFormat, *projectGroup.Project, *projectGroup.Group)
}

const testAccProjectGroupUpdateConfigFormat = `
resource "oktapam_project" "test-project-group-project" {
    name = "%s"
  	next_unix_uid = 60120
  	next_unix_gid = 63020
}
resource "oktapam_group" "test-project-group-group" {
    name = "%s"
}
resource "oktapam_project_group" "test-acc-project-group" {
    project_name = oktapam_project.test-project-group-project.name
  	group_name = oktapam_group.test-project-group-group.name
	server_access = true
	server_admin = false
	create_server_group = true
}`

func createTestAccProjectGroupUpdateConfig(projectGroup client.ProjectGroup) string {
	return fmt.Sprintf(testAccProjectGroupUpdateConfigFormat, *projectGroup.Project, *projectGroup.Group)
}
