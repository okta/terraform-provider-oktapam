package oktapam

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
)

func TestAccDatasourceResourceGroupServerEnrollmentTokensList(t *testing.T) {
	checkTeamApplicable(t, true)

	// Generate details
	identifier := randSeq()

	// config to create the resources
	initConfig := createTestAccDatasourceResourceGroupServerEnrollmentTokensInitConfig(identifier)

	// config for the datasources
	listAllConfig := testAccDatasourceResourceGroupServerEnrollmentTokensConfig(identifier, identifier)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccResourceGroupServerEnrollmentTokensCheckDestroy(identifier),
		Steps: []resource.TestStep{
			{
				Config: initConfig,
			},
			{
				Config: fmt.Sprintf("%s\n%s", initConfig, listAllConfig),
				Check:  resource.TestCheckResourceAttr("data.oktapam_resource_group_server_enrollment_tokens.resource_group_server_enrollment_tokens", fmt.Sprintf("%s.#", attributes.IDs), "2"),
			},
		},
	})
}

func testAccResourceGroupServerEnrollmentTokensCheckDestroy(identifiers ...string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		c := testAccProvider.Meta().(client.OktaPAMClient)

		resourceGroups, err := c.ListResourceGroups(context.Background())
		if err != nil {
			return fmt.Errorf("error getting resource groups: %w", err)
		}

		m := make(map[string]bool, len(identifiers))
		for _, id := range identifiers {
			m[id] = true
		}

		for _, rg := range resourceGroups {
			if _, ok := m[*rg.Name]; ok {
				return fmt.Errorf("resource groups still exists")
			}
		}

		return nil
	}
}

const testAccDatasourceResourceGroupServerEnrollmentTokensCreateConfigFormat = `
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
resource "oktapam_resource_group_server_enrollment_token" "test-resource-group-server-enrollment-token-1" {
	resource_group       = oktapam_resource_group.test-resource-group.id
	project = oktapam_resource_group_project.test-resource-group-project.id
	description = "test rg server enrollment token 1"
}
resource "oktapam_resource_group_server_enrollment_token" "test-resource-group-server-enrollment-token-2" {
	resource_group       = oktapam_resource_group.test-resource-group.id
	project = oktapam_resource_group_project.test-resource-group-project.id
	description = "test rg server enrollment token 2"
}
`

func createTestAccDatasourceResourceGroupServerEnrollmentTokensInitConfig(identifier string) string {
	return fmt.Sprintf(testAccDatasourceResourceGroupServerEnrollmentTokensCreateConfigFormat, identifier, identifier, identifier)
}

const testAccDatasourceResourceGroupServerEnrollmentTokensFormat = `
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
`

func testAccDatasourceResourceGroupServerEnrollmentTokensConfig(resourceGroupName, projectName string) string {
	return fmt.Sprintf(testAccDatasourceResourceGroupServerEnrollmentTokensFormat, resourceGroupName, projectName)
}
