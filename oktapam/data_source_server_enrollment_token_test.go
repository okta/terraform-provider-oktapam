package oktapam

import (
	"fmt"
	"testing"

	"github.com/kylelemons/godebug/pretty"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDatasourceServerEnrollmentTokenFetch(t *testing.T) {
	resourceName := "oktapam_server_enrollment_token.test-server-enrollment-token-1"
	dataSourceName := "data.oktapam_server_enrollment_token.target_token"

	identifier := randSeq(10)
	projectName := fmt.Sprintf("test-acc-datasource-server-enrollment-token-project-%s", identifier)
	description := fmt.Sprintf("Datasource Acceptance Test Token Set %s, Token 1", identifier)

	testConfig := createTestAccDatasourceServerEnrollmentTokenInitListFetchConfig(projectName, description)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccServerEnrollmentTokenCheckDestroy(projectName, identifier),
		Steps: []resource.TestStep{
			{
				Config: testConfig,
				Check:  checkResourcesEqual(resourceName, dataSourceName),
			},
		},
	})
}

func checkResourcesEqual(resourceName1, resourceName2 string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resource1, ok := s.RootModule().Resources[resourceName1]
		if !ok {
			return fmt.Errorf("resource 1 not found: %s", resourceName1)
		}

		resource2, ok := s.RootModule().Resources[resourceName2]
		if !ok {
			return fmt.Errorf("resource 2 not found: %s", resourceName2)
		}

		comparison := pretty.Compare(resource1.Primary.Attributes, resource2.Primary.Attributes)
		if comparison != "" {
			return fmt.Errorf("resources are not equal")
		}
		return nil
	}
}

const testAccDatasourceServerEnrollmentTokenInitListFetchConfigFormat = `
resource "oktapam_project" "test-server-enrollment-token-project" {
    name = "%s"
  	next_unix_uid = 60120
  	next_unix_gid = 63020
}

resource "oktapam_server_enrollment_token" "test-server-enrollment-token-1" {
    project_name = oktapam_project.test-server-enrollment-token-project.name
	description  = "%s"
	depends_on = [oktapam_project.test-server-enrollment-token-project]
}

data "oktapam_server_enrollment_tokens" "token_list" {
	project_name = oktapam_project.test-server-enrollment-token-project.name
	depends_on = [oktapam_server_enrollment_token.test-server-enrollment-token-1]
}

data "oktapam_server_enrollment_token" "target_token" {
	project_name = oktapam_project.test-server-enrollment-token-project.name
	id = data.oktapam_server_enrollment_tokens.token_list.ids[0]
	depends_on = [data.oktapam_server_enrollment_tokens.token_list]
}
`

func createTestAccDatasourceServerEnrollmentTokenInitListFetchConfig(projectName, description string) string {
	return fmt.Sprintf(testAccDatasourceServerEnrollmentTokenInitListFetchConfigFormat, projectName, description)
}
