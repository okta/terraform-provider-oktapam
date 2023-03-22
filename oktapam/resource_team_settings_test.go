package oktapam

import (
	"testing"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTeamSettings(t *testing.T) {
	resourceName := "oktapam_team_settings.test_team_setting"
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccTeamSettingsCreateConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						resourceName, attributes.ReactivateUsersViaIDP, "true",
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.ClientSessionDuration, "4200",
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.WebSessionDuration, "4200",
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.IncludeUserSID, "Never",
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.PostLogoutURL, "https://test.com",
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.PostLoginURL, "https://test.com",
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.ApproveDeviceWithoutInteraction, "false",
					),
				),
			},
			{
				Config: testAccTeamSettingsUpdateConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						resourceName, attributes.ReactivateUsersViaIDP, "false",
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.ClientSessionDuration, "3600",
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.WebSessionDuration, "3600",
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.IncludeUserSID, "Never",
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.PostLogoutURL, "https://test1.com",
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.PostLoginURL, "https://test1.com",
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.ApproveDeviceWithoutInteraction, "true",
					),
				),
			},
		},
	})
}

const testAccTeamSettingsCreateConfig = `
resource "oktapam_team_settings" "test_team_setting" {
  reactivate_users_via_idp           = true
  client_session_duration            = 4200
  web_session_duration               = 4200
  include_user_sid                   = "Never"
  post_logout_url                    = "https://test.com"
  post_login_url                     = "https://test.com"
  approve_device_without_interaction = false
}
`

const testAccTeamSettingsUpdateConfig = `
resource "oktapam_team_settings" "test_team_setting" {
  reactivate_users_via_idp           = false
  client_session_duration            = 3600
  web_session_duration               = 3600
  include_user_sid                   = "Never"
  post_logout_url                    = "https://test1.com"
  post_login_url                     = "https://test1.com"
  approve_device_without_interaction = true
}`
