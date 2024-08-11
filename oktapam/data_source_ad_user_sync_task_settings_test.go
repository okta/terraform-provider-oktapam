package oktapam

import (
	"fmt"
	"testing"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
)

// Creates a managed user-sync task-settings, then read it using data source, use list function as we need to find ID of
// the user sync task settings first
func TestAccDataSourceADUserSyncTaskSettings(t *testing.T) {
	checkTeamApplicable(t, false)
	adUserSyncTaskSettingsResourceName := "oktapam_ad_user_sync_task_settings.test_acc_ad_user_sync_task_settings"
	adConnectionResourceName := "oktapam_ad_connection.test_acc_ad_connection"
	dataSourceResourceName := "data.oktapam_ad_user_sync_task_settings.test_acc_datasource_ad_user_sync_task_settings"
	nameIdentifier := randSeq()
	adUserSyncTaskName := fmt.Sprintf("test_acc_ad_user_sync_task_settings_%s", nameIdentifier)
	adConnectionName := fmt.Sprintf("test_acc_ad_connection_%s", nameIdentifier)
	projectName := fmt.Sprintf("test_acc_project_%s", nameIdentifier)
	//Only one connection can exist per domain per team
	domainName := fmt.Sprintf("%s.example.com", nameIdentifier)

	//Build required pre-req config. AD User Sync Tasks Settings require AD Connection and Project
	preConfig := createTestAccADUserSyncTaskSettingsPreConfig(adConnectionName, projectName, domainName)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccV6ProviderFactories,
		CheckDestroy:             testAccADUserSyncTaskCheckDestroy(adConnectionResourceName),
		Steps: []resource.TestStep{
			{
				Config: createTestAccDataSourceADUserSyncTaskSettingsInitConfig(preConfig, adUserSyncTaskName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceResourceName, attributes.Name,
						adUserSyncTaskSettingsResourceName, attributes.Name),
					resource.TestCheckResourceAttrPair(dataSourceResourceName, attributes.Frequency,
						adUserSyncTaskSettingsResourceName, attributes.Frequency),
					resource.TestCheckResourceAttrPair(dataSourceResourceName, attributes.SIDField,
						adUserSyncTaskSettingsResourceName, attributes.SIDField),
					resource.TestCheckResourceAttrPair(dataSourceResourceName, attributes.UPNField,
						adUserSyncTaskSettingsResourceName, attributes.UPNField),
					resource.TestCheckResourceAttrPair(dataSourceResourceName, attributes.BaseDN,
						adUserSyncTaskSettingsResourceName, attributes.BaseDN),
					resource.TestCheckResourceAttrPair(dataSourceResourceName, attributes.LDAPQueryFilter,
						adUserSyncTaskSettingsResourceName, attributes.LDAPQueryFilter),
				),
			},
		},
	})
}

const testAccDataSourceADUserSyncTaskSettingsInitFormat = `
resource "oktapam_ad_user_sync_task_settings" "test_acc_ad_user_sync_task_settings" {
    connection_id            = oktapam_ad_connection.test_acc_ad_connection.id
    name                     = "%[1]s"
    is_active                = true
    frequency                = 12 # Every 12 hours Note: If 24 hours then start_hour_utc is required
    sid_field                = "objectSID"
    upn_field                = "userPrincipalName"
    base_dn                  = "dc=tilt,dc=scaleft,dc=com"
    ldap_query_filter        = "(objectclass=user)"
}

data "oktapam_ad_user_sync_task_settings_id_list" "adusts_list" {
    depends_on = [oktapam_ad_user_sync_task_settings.test_acc_ad_user_sync_task_settings]
    connection_id = oktapam_ad_connection.test_acc_ad_connection.id
}

data "oktapam_ad_user_sync_task_settings" "test_acc_datasource_ad_user_sync_task_settings" {
    id = data.oktapam_ad_user_sync_task_settings_id_list.adusts_list.ad_user_sync_task_settings_id_list[0]
    connection_id = oktapam_ad_connection.test_acc_ad_connection.id
}
`

func createTestAccDataSourceADUserSyncTaskSettingsInitConfig(preConfig string, adUserSyncTaskName string) string {
	logging.Debugf("creating config")
	return preConfig + fmt.Sprintf(testAccDataSourceADUserSyncTaskSettingsInitFormat, adUserSyncTaskName)
}
