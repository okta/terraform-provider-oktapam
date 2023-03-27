package oktapam

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDatasourceTeamSettingsFetch(t *testing.T) {
	resourceName := "oktapam_team_settings.test_team_setting-1"
	dataSourceName := "data.oktapam_team_settings.target"
	team := os.Getenv(teamSchemaEnvVar)

	testConfig := createTestAccDatasourceTeamSettingsInitConfig(team)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					checkResourcesEqual(resourceName, dataSourceName),
				),
			},
		},
	})
}

// NOTE: This config (1) creates two new resources (2) lists the existing resources with the matching identifier
// and (3) get the new resource as a data source.
// The test then compares the resource with its data source to ensure they are equal.

const testAccDatasourceTeamSettingsInitConfigFormat = `
resource "oktapam_team_settings" "test_team_setting-1" {
  reactivate_users_via_idp           = "false"
  include_user_sid                   = "Never"
  post_logout_url                    = "https://okta.com"
}
data "oktapam_team_settings" "target" {
  id = "%s"
}
`

func createTestAccDatasourceTeamSettingsInitConfig(identifier string) string {
	return fmt.Sprintf(testAccDatasourceTeamSettingsInitConfigFormat, identifier)
}