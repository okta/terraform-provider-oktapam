package oktapam

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDatasourceTeamSettingsFetch(t *testing.T) {
	resourceName := "oktapam_team_settings.test_team_setting-1"
	dataSourceName := "data.oktapam_team_settings.target"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDatasourceTeamSettingsInitListFetchConfigFormat,
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

const testAccDatasourceTeamSettingsInitListFetchConfigFormat = `
resource "oktapam_team_settings" "test_team_setting-1" {
  reactivate_users_via_idp           = "false"
  client_session_duration            = 3600
  web_session_duration               = 3600
  include_user_sid                   = "Never"
  post_logout_url                    = "https://test1.com"
  post_login_url                     = "https://test1.com"
  approve_device_without_interaction = "false"
}

data "oktapam_team_settings" "target" {
	id = "asa"
}
`
