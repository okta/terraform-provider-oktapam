package oktaasa

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/kylelemons/godebug/pretty"
	"github.com/terraform-providers/terraform-provider-oktaasa/oktaasa/client"
	"github.com/terraform-providers/terraform-provider-oktaasa/oktaasa/utils"
)

func TestAccProject(t *testing.T) {
	resourceName := "oktaasa_project.test-project"
	projectName := fmt.Sprintf("test-acc-project-%s", randSeq(10))
	initialProject := client.Project{
		Name:                   &projectName,
		NextUnixUID:            utils.AsIntPtr(60120),
		NextUnixGID:            utils.AsIntPtr(63020),
		ForceSharedSSHUsers:    utils.AsBoolPtrZero(false, true),
		CreateServerUsers:      utils.AsBoolPtrZero(false, true),
		ForwardTraffic:         utils.AsBoolPtrZero(false, true),
		RequirePreAuthForCreds: utils.AsBoolPtrZero(false, true),
		RDPSessionRecording:    utils.AsBoolPtrZero(false, true),
		SSHSessionRecording:    utils.AsBoolPtrZero(false, true),
		ADJoinedUsers:          utils.AsBoolPtrZero(false, true),
	}
	updatedProject := client.Project{
		Name:                   &projectName,
		NextUnixUID:            utils.AsIntPtr(61200),
		NextUnixGID:            utils.AsIntPtr(63400),
		ForceSharedSSHUsers:    utils.AsBoolPtrZero(false, true),
		CreateServerUsers:      utils.AsBoolPtrZero(true, true),
		ForwardTraffic:         utils.AsBoolPtrZero(true, true),
		RequirePreAuthForCreds: utils.AsBoolPtrZero(false, true),
		RDPSessionRecording:    utils.AsBoolPtrZero(true, true),
		SSHSessionRecording:    utils.AsBoolPtrZero(true, true),
		ADJoinedUsers:          utils.AsBoolPtrZero(false, true),
		GatewaySelector:        utils.AsStringPtr("env=test"),
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccProjectCheckDestroy(projectName),
		Steps: []resource.TestStep{
			{
				Config: createTestAccProjectCreateConfig(projectName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectCheckExists(resourceName, initialProject),
					resource.TestCheckResourceAttr(
						resourceName, "project_name", projectName,
					),
					resource.TestCheckResourceAttr(
						resourceName, "next_unix_uid", "60120",
					),
					resource.TestCheckResourceAttr(
						resourceName, "next_unix_gid", "63020",
					),
				),
			},
			{
				Config: createTestAccProjectUpdateConfig(projectName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectCheckExists(resourceName, updatedProject),
					resource.TestCheckResourceAttr(
						resourceName, "project_name", projectName,
					),
					resource.TestCheckResourceAttr(
						resourceName, "next_unix_uid", "61200",
					),
					resource.TestCheckResourceAttr(
						resourceName, "next_unix_gid", "63400",
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

func testAccProjectCheckExists(rn string, expectedProject client.Project) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[rn]
		if !ok {
			return fmt.Errorf("resource not found: %s", rn)
		}

		resourceID := rs.Primary.ID
		if resourceID == "" {
			return fmt.Errorf("resource id not set")
		}
		if *expectedProject.Name != resourceID {
			return fmt.Errorf("resource id not set to expected value.  expected %s, got %s", *expectedProject.Name, resourceID)
		}

		client := testAccProvider.Meta().(client.OktaASAClient)
		proj, err := client.GetProject(context.Background(), *expectedProject.Name, false)
		if err != nil {
			return fmt.Errorf("error getting project :%w", err)
		}
		if proj == nil || utils.IsBlank(proj.ID) {
			return fmt.Errorf("project %s does not exist", *expectedProject.Name)
		}
		expectedProject.ID = proj.ID
		expectedProject.Team = &client.Team
		comparison := pretty.Compare(proj, expectedProject)
		if comparison != "" {
			return fmt.Errorf("expected project does not match returned project.\n%s", comparison)
		}
		return nil
	}
}

func testAccProjectCheckDestroy(projectName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(client.OktaASAClient)
		proj, err := client.GetProject(context.Background(), projectName, false)
		if err != nil {
			return fmt.Errorf("error getting project: %w", err)
		}

		if proj != nil && proj.Exists() {
			return fmt.Errorf("project still exists")
		}

		return nil
	}
}

const testAccProjectCreateConfigFormat = `
resource "oktaasa_project" "test-project" {
    project_name = "%s"
  	next_unix_uid = 60120
  	next_unix_gid = 63020
}`

func createTestAccProjectCreateConfig(projectName string) string {
	return fmt.Sprintf(testAccProjectCreateConfigFormat, projectName)
}

const testAccProjectUpdateConfigFormat = `
resource "oktaasa_project" "test-project" {
    project_name              = "%s"
  	next_unix_uid             = 61200
  	next_unix_gid             = 63400
	create_server_users       = true
	forward_traffic           = true
	rdp_session_recording     = true
	ssh_session_recording     = true
	gateway_selector          = "env=test"
}`

func createTestAccProjectUpdateConfig(projectName string) string {
	return fmt.Sprintf(testAccProjectUpdateConfigFormat, projectName)
}
