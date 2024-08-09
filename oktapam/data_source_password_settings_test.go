package oktapam

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
)

func TestAccDatasourcePasswordSettingsFetch(t *testing.T) {
	checkTeamApplicable(t, true)

	// Generate details
	identifier := randSeq()

	// config to create the resources
	initConfig := createTestAccDatasourcePasswordSettingsInitConfig(identifier)

	// config for the datasources
	fetchConfig := testAccDatasourcePasswordSettingsConfig(identifier)

	resourceName := "data.oktapam_password_settings.pw_settings"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccV6ProviderFactories,
		CheckDestroy:             testAccResourceGroupsCheckDestroy(identifier),
		Steps: []resource.TestStep{
			{
				Config: initConfig,
			},
			{
				Config: fmt.Sprintf("%s\n%s", initConfig, fetchConfig),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						resourceName, attributes.EnablePeriodicRotation, "true",
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.PeriodicRotationDurationInSeconds, "3600",
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.MinLength, "12",
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.MaxLength, "16",
					),
					resource.TestCheckResourceAttr(
						resourceName, fmt.Sprintf("%s.0.%s", attributes.CharacterOptions, attributes.LowerCase), "true",
					),
					resource.TestCheckResourceAttr(
						resourceName, fmt.Sprintf("%s.0.%s", attributes.CharacterOptions, attributes.UpperCase), "true",
					),
					resource.TestCheckResourceAttr(
						resourceName, fmt.Sprintf("%s.0.%s", attributes.CharacterOptions, attributes.Digits), "true",
					),
					resource.TestCheckResourceAttr(
						resourceName, fmt.Sprintf("%s.0.%s", attributes.CharacterOptions, attributes.Punctuation), "true",
					),
				),
			},
		},
	})
}

const testAccDatasourcePasswordSettingsCreateConfigFormat = `
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
	ssh_certificate_type  = "CERT_TYPE_ED25519_01"
	account_discovery     = true
}
resource "oktapam_password_settings" "test-password-settings" {
	resource_group = oktapam_resource_group.test-resource-group.id
	project = oktapam_resource_group_project.test-resource-group-project.id
	managed_privileged_accounts = ["root", "pamadmin"]
	enable_periodic_rotation    = true
	periodic_rotation_duration_in_seconds = 3600
    min_length = 12
	max_length = 16
	character_options {
		upper_case = true
		lower_case = true
		digits = true
		punctuation = true
	}
}
`

func createTestAccDatasourcePasswordSettingsInitConfig(identifier string) string {
	return fmt.Sprintf(testAccDatasourcePasswordSettingsCreateConfigFormat, identifier, identifier, identifier)
}

const testAccDatasourcePasswordSettingsFormat = `
data "oktapam_resource_groups" "resource_groups" {
	name = "%s"
}

data "oktapam_resource_group_projects" "rg_projects" {
	resource_group = data.oktapam_resource_groups.resource_groups.ids[0]
	name = "%s"
}

data "oktapam_password_settings" "pw_settings" {
	resource_group = data.oktapam_resource_groups.resource_groups.ids[0]
	project = data.oktapam_resource_group_projects.rg_projects.ids[0]
}
`

func testAccDatasourcePasswordSettingsConfig(identifier string) string {
	return fmt.Sprintf(testAccDatasourcePasswordSettingsFormat, identifier, identifier)
}
