package oktapam

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
)

func TestAccDatasourceResourceGroupServerEnrollmentTokensFetch(t *testing.T) {
	checkTeamApplicable(t, true)

	// Generate details
	identifier := randSeq()

	// config to create the resources
	initConfig := createTestAccDatasourceResourceGroupServerEnrollmentTokenInitConfig(identifier)

	// config for the datasources
	fetchConfig := testAccDatasourceResourceGroupServerEnrollmentTokenConfig(identifier, identifier)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccResourceGroupServerEnrollmentTokenCheckDestroy(identifier),
		Steps: []resource.TestStep{
			{
				Config: initConfig,
			},
			{
				Config: fmt.Sprintf("%s\n%s", initConfig, fetchConfig),
				Check:  resource.TestCheckResourceAttr("data.oktapam_resource_group_server_enrollment_token.resource_group_server_enrollment_token", attributes.Description, fmt.Sprintf("test rg server enrollment token %s", identifier)),
			},
		},
	})
}

const testAccDatasourceResourceGroupServerEnrollmentTokenCreateConfigFormat = `
resource "oktapam_group" "test_group" {
	name = "dra-test-group-%s"
}
resource "oktapam_resource_group" "test-resource-group" {
	name = "%s"
	description = "terraform test rg"
	delegated_resource_admin_groups = [oktapam_group.test_group.id]
}
resource "oktapam_resource_group_project" "test-resource-group-project" {
	name = "%s"
	resource_group =      oktapam_resource_group.test-resource-group.id
	next_unix_uid         = 60120
	next_unix_gid         = 63020
	ssh_certificate_type  = "CERT_TYPE_ED25519_01"
}
resource "oktapam_resource_group_server_enrollment_token" "test-resource-group-server-enrollment-token" {
	resource_group       = oktapam_resource_group.test-resource-group.id
	project = oktapam_resource_group_project.test-resource-group-project.id
	description = "test rg server enrollment token %s"
}
`

func createTestAccDatasourceResourceGroupServerEnrollmentTokenInitConfig(identifier string) string {
	return fmt.Sprintf(testAccDatasourceResourceGroupServerEnrollmentTokenCreateConfigFormat, identifier, identifier, identifier, identifier)
}

const testAccDatasourceResourceGroupServerEnrollmentTokenFormat = `
data "oktapam_resource_groups" "resource_groups" {
	name = "%s"
}
data "oktapam_resource_group_projects" "resource_group_projects" {
	resource_group = data.oktapam_resource_groups.resource_groups.ids[0]
	name = "%s"
}
data "oktapam_resource_group_server_enrollment_tokens" "resource_group_server_enrollment_tokens" {
	resource_group = data.oktapam_resource_groups.resource_groups.ids[0]
	project = data.oktapam_resource_group_projects.resource_group_projects.ids[0]
}
data "oktapam_resource_group_server_enrollment_token" "resource_group_server_enrollment_token" {
	id = data.oktapam_resource_group_server_enrollment_tokens.resource_group_server_enrollment_tokens.ids[0]
	resource_group = data.oktapam_resource_groups.resource_groups.ids[0]
	project = data.oktapam_resource_group_projects.resource_group_projects.ids[0]
}
`

func testAccDatasourceResourceGroupServerEnrollmentTokenConfig(resourceGroupName, projectName string) string {
	return fmt.Sprintf(testAccDatasourceResourceGroupServerEnrollmentTokenFormat, resourceGroupName, projectName)
}
