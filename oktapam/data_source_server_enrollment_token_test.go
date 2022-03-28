package oktapam

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDatasourceServerEnrollmentToken(t *testing.T) {
	resourceName := "data.oktapam_server_enrollment_token.test_server_enrollment_tokens"
	projectName := fmt.Sprintf("test-acc-datasource-server-enrollment-token-project-%s", randSeq(10))
	description1 := fmt.Sprintf("Datasource Acceptance Test Token - 1: %s", randSeq(10))
	description2 := fmt.Sprintf("Datasource Acceptance Test Token - 2: %s", randSeq(10))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: createTestAccDatasourceServerEnrollmentTokenInitConfig(projectName, description1, description2),
			},
			{
				Config: createTestAccDatasourceServerEnrollmentTokenConfig(projectName),
				Check:  resource.TestCheckResourceAttr(resourceName, "server_enrollment_tokens.#", "2"),
			},
		},
	})
}

const testAccDatasourceServerEnrollmentTokenConfigFormat = `
data "oktapam_server_enrollment_token" "test_server_enrollment_tokens" {
	project_name = "%s"
}
`

func createTestAccDatasourceServerEnrollmentTokenConfig(projectName string) string {
	return fmt.Sprintf(testAccDatasourceServerEnrollmentTokenConfigFormat, projectName)
}

const testAccDatasourceServerEnrollmentTokenInitConfigFormat = `
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

resource "oktapam_server_enrollment_token" "test-server-enrollment-token-2" {
    project_name = oktapam_project.test-server-enrollment-token-project.name
	description  = "%s"
	depends_on = [oktapam_project.test-server-enrollment-token-project]
}
`

func createTestAccDatasourceServerEnrollmentTokenInitConfig(projectName string, description1, description2 string) string {
	return fmt.Sprintf(testAccDatasourceServerEnrollmentTokenInitConfigFormat, projectName, description1, description2)
}
