package oktaasa

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/terraform-providers/terraform-provider-oktaasa/oktaasa/client"
	"github.com/terraform-providers/terraform-provider-oktaasa/oktaasa/utils"
)

func TestAccServerEnrollmentToken(t *testing.T) {
	resourceName := "oktaasa_server_enrollment_token.test-server-enrollment-token"
	description := fmt.Sprintf("Acceptance Test Token: %s", randSeq(10))
	projectName := fmt.Sprintf("test-acc-server-enrollment-token-project-%s", randSeq(10))
	enrollmentToken := client.ServerEnrollmentToken{
		Project:     &projectName,
		Description: &description,
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccServerEnrollmentTokenCheckDestroy(enrollmentToken),
		Steps: []resource.TestStep{
			{
				Config: createTestAccServerEnrollmentTokenCreateConfig(enrollmentToken),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccServerEnrollmentTokenCheckExists(resourceName, enrollmentToken),
					resource.TestCheckResourceAttr(
						resourceName, "project_name", projectName,
					),
					resource.TestCheckResourceAttr(
						resourceName, "description", description,
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

func testAccServerEnrollmentTokenCheckExists(rn string, expectedServerEnrollmentToken client.ServerEnrollmentToken) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[rn]
		if !ok {
			return fmt.Errorf("resource not found: %s", rn)
		}

		resourceID := rs.Primary.ID
		projectName, tokenID, err := parseServerEnrollmentTokenResourceID(resourceID)
		if err != nil {
			return fmt.Errorf("error parsing resource id: %w", err)
		}
		if projectName != *expectedServerEnrollmentToken.Project {
			return fmt.Errorf("resource ID did not have the expected project. expected %s, got %s", *expectedServerEnrollmentToken.Project, projectName)
		}

		client := testAccProvider.Meta().(client.OktaASAClient)
		token, err := client.GetServerEnrollmentToken(context.Background(), projectName, tokenID)
		if err != nil {
			return fmt.Errorf("error getting server enrollment token: %w", err)
		}
		if token == nil || utils.IsBlank(token.ID) {
			return fmt.Errorf("server enrollment token for project %s with id %s does not exist", projectName, tokenID)
		}
		if *token.Description != *expectedServerEnrollmentToken.Description {
			return fmt.Errorf("expected description does not match returned description for server enrollment token.  expected: %s, got: %s", *expectedServerEnrollmentToken.Description, *token.Description)
		}

		return nil
	}
}

func testAccServerEnrollmentTokenCheckDestroy(token client.ServerEnrollmentToken) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(client.OktaASAClient)
		project, err := client.GetProject(context.Background(), *token.Project, false)
		if err != nil {
			return fmt.Errorf("error getting project associated with server enrollment token: %w", err)
		}
		if project == nil || utils.IsBlank(project.ID) {
			// if all resources are removed, the project will be removed, which also removes the enrollment token
			return nil
		}

		tokens, err := client.ListServerEnrollmentTokens(context.Background(), *token.Project)
		if err != nil {
			return fmt.Errorf("error getting tokens: %w", err)
		}

		for _, t := range tokens {
			if token.Description == t.Description {
				return fmt.Errorf("token still exists")
			}
		}
		// the token was removed, but the project still exists
		return fmt.Errorf("project still exists")
	}
}

const testAccServerEnrollmentTokenCreateConfigFormat = `
resource "oktaasa_project" "test-server-enrollment-token-project" {
    name = "%s"
  	next_unix_uid = 60120
  	next_unix_gid = 63020
}
resource "oktaasa_server_enrollment_token" "test-server-enrollment-token" {
    project_name = oktaasa_project.test-server-enrollment-token-project.name
	description  = "%s"
	depends_on = [oktaasa_project.test-server-enrollment-token-project]
}
`

func createTestAccServerEnrollmentTokenCreateConfig(token client.ServerEnrollmentToken) string {
	return fmt.Sprintf(testAccServerEnrollmentTokenCreateConfigFormat, *token.Project, *token.Description)
}
