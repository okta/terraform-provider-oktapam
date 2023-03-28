package oktapam

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
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
		CheckDestroy:      testAccTeamSettingCheckDestroy(),
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

func testAccTeamSettingCheckDestroy() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(client.OktaPAMClient)
		teamSettings, err := client.GetTeamSettings(context.Background())
		if err!=nil {
			return fmt.Errorf("error getting team settings: %w", err)
		}
		if teamSettings == nil{
			return fmt.Errorf("team settings got deleted even when it can not be deleted")
		}

		return nil
	}
}

// NOTE: This config creates a team settings resource and gets the new resource as a data source.
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