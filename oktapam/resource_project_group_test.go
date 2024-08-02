package oktapam

import (
	"context"
	"fmt"
	"testing"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/kylelemons/godebug/pretty"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"
)

func TestAccProjectGroup(t *testing.T) {
	checkTeamApplicable(t, false)
	resourceName := "oktapam_project_group.test_acc_project_group"
	projectName := fmt.Sprintf("test_acc_project_group_project_%s", randSeq())
	groupName := fmt.Sprintf("test_acc_project_group_group_%s", randSeq())

	initialProjectGroup := client.ProjectGroup{
		Project:      &projectName,
		Group:        &groupName,
		ServerAccess: true,
		ServerAdmin:  true,
	}
	updatedProjectGroup := client.ProjectGroup{
		Project:           &projectName,
		Group:             &groupName,
		ServerAccess:      true,
		ServerAdmin:       false,
		CreateServerGroup: true,
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccV6ProviderFactories,
		CheckDestroy:             testAccProjectGroupCheckDestroy(initialProjectGroup),
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
				ResourceName: resourceName,
				ImportState:  true,
				//Used to dynamically generate the ID from the terraform state.
				//Terraform Resource ID is set to ASA ProjectGroup UUID but read requires "Project Name" and "Group Name" to retrieve the resource
				ImportStateIdFunc: testAccProjectGroupImportStateId(resourceName),
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

		project := rs.Primary.Attributes[attributes.ProjectName]
		group := rs.Primary.Attributes[attributes.GroupName]

		pamClient := testAccAPIClients.LocalClient
		projectGroup, err := pamClient.GetProjectGroup(context.Background(), project, group)
		if err != nil {
			return fmt.Errorf("error getting project group: %w", err)
		} else if projectGroup == nil {
			return fmt.Errorf("project group does not exist")
		}

		//"ID" is computed after resource creation, so make it same to avoid comparison diff.
		expectedProjectGroup.ID = projectGroup.ID
		comparison := pretty.Compare(expectedProjectGroup, projectGroup)
		if comparison != "" {
			return fmt.Errorf("expected project group does not match returned project.\n%s", comparison)
		}
		return nil
	}
}

func testAccProjectGroupCheckDestroy(projectGroup client.ProjectGroup) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		pamClient := testAccAPIClients.LocalClient
		pg, err := pamClient.GetProjectGroup(context.Background(), *projectGroup.Project, *projectGroup.Group)
		if err != nil {
			return fmt.Errorf("error getting project group: %w", err)
		}
		if pg != nil && utils.IsNonEmpty(pg.Project) && utils.IsNonEmpty(pg.Group) && utils.IsBlank(pg.DeletedAt) && utils.IsBlank(pg.RemovedAt) {
			return fmt.Errorf("project group still exists")
		}
		group, err := pamClient.GetGroup(context.Background(), *projectGroup.Group, false)
		if err != nil {
			return fmt.Errorf("error getting group associated with project group: %w", err)
		}
		if group != nil && utils.IsNonEmpty(group.ID) {
			return fmt.Errorf("group still exists: %s", *projectGroup.Group)
		}
		project, err := pamClient.GetProject(context.Background(), *projectGroup.Project, false)
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
resource "oktapam_project" "test_project_group_project" {
    name = "%s"
  	next_unix_uid = 60120
  	next_unix_gid = 63020
}
resource "oktapam_group" "test_project_group_group" {
    name = "%s"
}
resource "oktapam_project_group" "test_acc_project_group" {
    project_name = oktapam_project.test_project_group_project.name
  	group_name = oktapam_group.test_project_group_group.name
	server_access = true
	server_admin = true
	create_server_group = false
}`

func createTestAccProjectGroupCreateConfig(projectGroup client.ProjectGroup) string {
	return fmt.Sprintf(testAccProjectGroupCreateConfigFormat, *projectGroup.Project, *projectGroup.Group)
}

const testAccProjectGroupUpdateConfigFormat = `
resource "oktapam_project" "test_project_group_project" {
    name = "%s"
  	next_unix_uid = 60120
  	next_unix_gid = 63020
}
resource "oktapam_group" "test_project_group_group" {
    name = "%s"
}
resource "oktapam_project_group" "test_acc_project_group" {
    project_name = oktapam_project.test_project_group_project.name
  	group_name = oktapam_group.test_project_group_group.name
	server_access = true
	server_admin = false
	create_server_group = true
}`

func createTestAccProjectGroupUpdateConfig(projectGroup client.ProjectGroup) string {
	return fmt.Sprintf(testAccProjectGroupUpdateConfigFormat, *projectGroup.Project, *projectGroup.Group)
}

func testAccProjectGroupImportStateId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("Not found: %s", resourceName)
		}
		return fmt.Sprintf("%s/%s", rs.Primary.Attributes[attributes.ProjectName], rs.Primary.Attributes[attributes.GroupName]), nil
	}
}
