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

func TestAccDatasourceProject(t *testing.T) {
	resourceName := "data.oktapam_project.test_projects"
	projectNamePrefix := fmt.Sprintf("test-acc-datasource-project-%s", randSeq(10))
	expectedProjects := map[string]client.Project{
		projectNamePrefix + "-one": {
			Name:        utils.AsStringPtr(projectNamePrefix + "-one"),
			NextUnixUID: utils.AsIntPtr(60120),
			NextUnixGID: utils.AsIntPtr(63020),
		},
		projectNamePrefix + "-two": {
			Name:        utils.AsStringPtr(projectNamePrefix + "-two"),
			NextUnixUID: utils.AsIntPtr(60220),
			NextUnixGID: utils.AsIntPtr(63120),
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: createTestAccDatasourceProjectCreateConfig(projectNamePrefix),
			},
			{
				Config: createTestAccDatasourceProjectsConfig(projectNamePrefix),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.#", attributes.Projects), "2"),
					testAccDatasourceProjectsCheck(resourceName, expectedProjects),
				),
			},
		},
	})
}

func testAccDatasourceProjectsCheck(rn string, expectedProjects map[string]client.Project) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[rn]
		if !ok {
			return fmt.Errorf("resource not found: %s", rn)
		}
		mappings, err := getIndexMappingFromResource(rs, attributes.Projects, attributes.Name, len(expectedProjects))
		if err != nil {
			return fmt.Errorf("error mapping resources to indices: %w", err)
		}
		primaryAttributes := rs.Primary.Attributes
		for name, project := range expectedProjects {
			// tests some attributes to ensure we are obtaining some attributes that were set by the original create resource and some that were computed
			idx, ok := mappings[name]
			if !ok {
				return fmt.Errorf("could not find resource with %s: %s", attributes.Name, name)
			}
			if name != *project.Name {
				return fmt.Errorf("%s attributes did not match for project %q, expected %q, got %q", attributes.Name, name, *project.Name, name)
			}
			nextUnixUID, ok := primaryAttributes[fmt.Sprintf("%s.%s.%s", attributes.Projects, idx, attributes.NextUnixUID)]
			if !ok {
				return fmt.Errorf("%s attribute not set for project %q", attributes.NextUnixUID, name)
			}
			expectedNextUnixUID := fmt.Sprintf("%d", *project.NextUnixUID)
			if nextUnixUID != expectedNextUnixUID {
				return fmt.Errorf("mismatch for %s value, expected %q, got %q", attributes.NextUnixUID, expectedNextUnixUID, nextUnixUID)
			}
			nextUnixGID, ok := primaryAttributes[fmt.Sprintf("%s.%s.%s", attributes.Projects, idx, attributes.NextUnixGID)]
			if !ok {
				return fmt.Errorf("%s attribute not set for project %q", attributes.NextUnixGID, name)
			}
			expectedNextUnixGID := fmt.Sprintf("%d", *project.NextUnixGID)
			if nextUnixGID != expectedNextUnixGID {
				return fmt.Errorf("mismatch for %s value, expected %q, got %q", attributes.NextUnixGID, expectedNextUnixGID, nextUnixGID)
			}
			createServerUsers, ok := primaryAttributes[fmt.Sprintf("%s.%s.%s", attributes.Projects, idx, attributes.CreateServerUsers)]
			if !ok {
				return fmt.Errorf("%s attribute not set for project %q", attributes.CreateServerUsers, name)
			}
			expectedCreateServerUsers := project.CreateServerUsers != nil && *project.CreateServerUsers
			if createServerUsers != fmt.Sprint(expectedCreateServerUsers) {
				return fmt.Errorf("mismatch for %s attribute, expected %q, got %q", attributes.CreateServerUsers, fmt.Sprint(expectedCreateServerUsers), createServerUsers)
			}
		}
		return nil
	}
}

const testAccDatasourceProjectsConfigFormat = `
data "oktapam_project" "test_projects" {
	contains = "%s"
}
`

func createTestAccDatasourceProjectsConfig(prefix string) string {
	return fmt.Sprintf(testAccDatasourceProjectsConfigFormat, prefix)
}

const testAccDatasourceProjectCreateConfigFormat = `
resource "oktapam_project" "test-project-one" {
	name = "%s-one"
  	next_unix_uid = 60120
  	next_unix_gid = 63020
}

resource "oktapam_project" "test-project-two" {
	name = "%s-two"
  	next_unix_uid = 60220
  	next_unix_gid = 63120
}
`

func createTestAccDatasourceProjectCreateConfig(groupName string) string {
	return fmt.Sprintf(testAccDatasourceProjectCreateConfigFormat, groupName, groupName)
}
