package oktapam

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDatasourceServerEnrollmentTokenFetch(t *testing.T) {
	resourceName := "oktapam_server_enrollment_token.test_server_enrollment_token_1"
	dataSourceName := "data.oktapam_server_enrollment_token.target_token"

	identifier := randSeq()
	projectName := fmt.Sprintf("test_acc_datasource_server_enrollment_token_project_%s", identifier)
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

const testAccDatasourceServerEnrollmentTokenInitListFetchConfigFormat = `
resource "oktapam_project" "test_server_enrollment_token_project" {
    name = "%s"
  	next_unix_uid = 60120
  	next_unix_gid = 63020
}

resource "oktapam_server_enrollment_token" "test_server_enrollment_token_1" {
    project_name = oktapam_project.test_server_enrollment_token_project.name
	description  = "%s"
}

data "oktapam_server_enrollment_tokens" "token_list" {
	project_name = oktapam_project.test_server_enrollment_token_project.name
	depends_on = [oktapam_server_enrollment_token.test_server_enrollment_token_1]
}

data "oktapam_server_enrollment_token" "target_token" {
	project_name = oktapam_project.test_server_enrollment_token_project.name
	id = data.oktapam_server_enrollment_tokens.token_list.ids[0]
}
`

func createTestAccDatasourceServerEnrollmentTokenInitListFetchConfig(projectName string, description string) string {
	return fmt.Sprintf(testAccDatasourceServerEnrollmentTokenInitListFetchConfigFormat, projectName, description)
}
