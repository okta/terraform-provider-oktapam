package oktapam

import (
	"fmt"
	"testing"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDatasourceServerEnrollmentTokenList(t *testing.T) {
	checkTeamApplicable(t, false)
	dataSourceName := fmt.Sprintf("data.%s.tokens_list", providerServerEnrollmentTokensKey)

	// Generate details
	identifier := randSeq()
	projectName := fmt.Sprintf("test-acc-datasource-server-enrollment-token-project-%s", identifier)
	description1 := fmt.Sprintf("Datasource Acceptance Test Token Set %s, Token 1", identifier)
	description2 := fmt.Sprintf("Datasource Acceptance Test Token Set %s, Token 2", identifier)

	// Generate Config 1: create two tokens
	initConfig := createTestAccDatasourceServerEnrollmentTokenInitConfig(projectName, description1, description2)

	// Generate Config 2: list using filter that returns both
	listConfig := createTestAccDatasourceServerEnrollmentTokenConfig(projectName)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccServerEnrollmentTokenCheckDestroy(projectName, identifier),
		Steps: []resource.TestStep{
			{
				Config: initConfig,
			},
			{
				Config: listConfig,
				Check:  resource.TestCheckResourceAttr(dataSourceName, fmt.Sprintf("%s.#", attributes.IDs), "2"),
			},
		},
	})
}

const testAccDatasourceServerEnrollmentTokenConfigFormat = `
data "oktapam_server_enrollment_tokens" "tokens_list" {
	project_name = "%s"
}
`

func createTestAccDatasourceServerEnrollmentTokenConfig(projectName string) string {
	return fmt.Sprintf(testAccDatasourceServerEnrollmentTokenConfigFormat, projectName)
}

const testAccDatasourceServerEnrollmentTokenInitConfigFormat = `
resource "oktapam_project" "test_server_enrollment_token_project" {
    name = "%s"
  	next_unix_uid = 60120
  	next_unix_gid = 63020
}
resource "oktapam_server_enrollment_token" "test_server_enrollment_token_1" {
    project_name = oktapam_project.test_server_enrollment_token_project.name
	description  = "%s"
}

resource "oktapam_server_enrollment_token" "test_server_enrollment_token_2" {
    project_name = oktapam_project.test_server_enrollment_token_project.name
	description  = "%s"
}
`

func createTestAccDatasourceServerEnrollmentTokenInitConfig(projectName string, description1 string, description2 string) string {
	return fmt.Sprintf(testAccDatasourceServerEnrollmentTokenInitConfigFormat, projectName, description1, description2)
}
