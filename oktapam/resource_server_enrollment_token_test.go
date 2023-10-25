package oktapam

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"
)

func TestAccServerEnrollmentToken(t *testing.T) {
	checkTeamApplicable(t, false)
	resourceName := "oktapam_server_enrollment_token.test_server_enrollment_token"
	identifier := randSeq()
	description := fmt.Sprintf("Acceptance Test Token Set %s, Token 1", identifier)
	projectName := fmt.Sprintf("test_acc_server_enrollment_token_project_%s", identifier)
	enrollmentToken := client.ServerEnrollmentToken{
		Project:     &projectName,
		Description: &description,
	}
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccServerEnrollmentTokenCheckDestroy(projectName, identifier),
		Steps: []resource.TestStep{
			{
				Config: createTestAccServerEnrollmentTokenCreateConfig(enrollmentToken),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccServerEnrollmentTokenCheckExists(resourceName, enrollmentToken),
					resource.TestCheckResourceAttr(
						resourceName, attributes.ProjectName, projectName,
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.Description, description,
					),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccServerEnrollmentTokenImportStateId(resourceName),
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

		serverEnrollmentTokenId := rs.Primary.ID
		projectName := rs.Primary.Attributes[attributes.ProjectName]
		if projectName != *expectedServerEnrollmentToken.Project {
			return fmt.Errorf("resource id did not have the expected project. expected %s, got %s", *expectedServerEnrollmentToken.Project, projectName)
		}

		client := getLocalClientFromMetadata(testAccProvider.Meta())
		token, err := client.GetServerEnrollmentToken(context.Background(), projectName, serverEnrollmentTokenId)
		if err != nil {
			return fmt.Errorf("error getting server enrollment token: %w", err)
		}
		if token == nil || utils.IsBlank(token.ID) {
			return fmt.Errorf("server enrollment token for project %s with id %s does not exist", projectName, serverEnrollmentTokenId)
		}
		if *token.Description != *expectedServerEnrollmentToken.Description {
			return fmt.Errorf("expected description does not match returned description for server enrollment token.  expected: %s, got: %s", *expectedServerEnrollmentToken.Description, *token.Description)
		}

		return nil
	}
}

func testAccServerEnrollmentTokenCheckDestroy(projectName string, identifier string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := getLocalClientFromMetadata(testAccProvider.Meta())
		project, err := client.GetProject(context.Background(), projectName, false)
		if err != nil {
			return fmt.Errorf("error getting project associated with server enrollment token: %w", err)
		}
		if project == nil || utils.IsBlank(project.ID) {
			// if all resources are removed, the project will be removed, which also removes the enrollment token
			return nil
		}

		tokens, err := client.ListServerEnrollmentTokens(context.Background(), projectName)
		if err != nil {
			return fmt.Errorf("error getting tokens: %w", err)
		}

		for _, token := range tokens {
			if strings.Contains(*token.Description, identifier) {
				return fmt.Errorf("token still exists")
			}
		}

		// the token was removed, but the project still exists
		return fmt.Errorf("project still exists")
	}
}

const testAccServerEnrollmentTokenCreateConfigFormat = `
resource "oktapam_project" "test_server_enrollment_token_project" {
    name = "%s"
  	next_unix_uid = 60120
  	next_unix_gid = 63020
}
resource "oktapam_server_enrollment_token" "test_server_enrollment_token" {
    project_name = oktapam_project.test_server_enrollment_token_project.name
	description  = "%s"
}
`

func createTestAccServerEnrollmentTokenCreateConfig(token client.ServerEnrollmentToken) string {
	return fmt.Sprintf(testAccServerEnrollmentTokenCreateConfigFormat, *token.Project, *token.Description)
}

func testAccServerEnrollmentTokenImportStateId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("Not found: %s", resourceName)
		}
		return fmt.Sprintf("%s/%s", rs.Primary.Attributes[attributes.ProjectName], rs.Primary.ID), nil
	}
}
