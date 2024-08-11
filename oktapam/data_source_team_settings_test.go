package oktapam

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/kylelemons/godebug/pretty"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/config"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDatasourceTeamSettingsFetch(t *testing.T) {
	resourceName := "oktapam_team_settings.test_team_setting"
	dataSourceName := "data.oktapam_team_settings.target"
	team := os.Getenv(config.TeamSchemaEnvVar)

	testConfig := createTestAccDatasourceTeamSettingsInitConfig(team)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccV6ProviderFactories,
		CheckDestroy:             testAccTeamSettingCheckDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					checkTeamSettingsResourcesEqual(resourceName, dataSourceName),
				),
			},
		},
	})
}

func testAccTeamSettingCheckDestroy() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := getTestAccAPIClients().LocalClient
		teamSettings, err := client.GetTeamSettings(context.Background())
		if err != nil {
			return fmt.Errorf("error getting team settings: %w", err)
		}
		if teamSettings == nil {
			return fmt.Errorf("team settings got deleted even when it can not be deleted")
		}

		return nil
	}
}

// NOTE: This config creates a team settings resource and gets the new resource as a data source.
// The test then compares the resource with its data source to ensure they are equal.

const testAccDatasourceTeamSettingsInitConfigFormat = `
resource "oktapam_team_settings" "test_team_setting" {
  reactivate_users_via_idp           = false
  approve_device_without_interaction = false
  include_user_sid                   = "Never"
}
data "oktapam_team_settings" "target" {
  depends_on = [oktapam_team_settings.test_team_setting]
  id = "%s"
}
`

func createTestAccDatasourceTeamSettingsInitConfig(identifier string) string {
	return fmt.Sprintf(testAccDatasourceTeamSettingsInitConfigFormat, identifier)
}

func checkTeamSettingsResourcesEqual(resourceName1 string, resourceName2 string) resource.TestCheckFunc {
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
			return fmt.Errorf("resources are not equal: %s", comparison)
		}
		return nil
	}
}
