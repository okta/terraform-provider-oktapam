package oktapam

import (
	"fmt"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
)

//TestAccDataSourceADUserSyncTaskSettingsList creates a few  managed user-sync task-settings, then reads them as list,
// so that it can be checked for length and its elements accessed.
func TestAccDataSourceADUserSyncTaskSettingsList(t *testing.T) {
	adConnectionResourceName := "oktapam_ad_connection.test_acc_ad_connection"
	dataSourceResourceName := "data.oktapam_ad_user_sync_task_settings_list.test_acc_data_source_ad_user_sync_task_settings_list"
	nameIdentifier := randSeq()
	adUserSyncTaskNamePrefix := fmt.Sprintf("test_acc_ad_user_sync_task_settings_%s", nameIdentifier)
	reTaskName := regexp.MustCompile(`test_acc_ad_user_sync_task_settings_[0-9A-Z]*_[1-3]`)
	adConnectionName := fmt.Sprintf("test_acc_ad_connection_%s", nameIdentifier)
	projectName := fmt.Sprintf("test_acc_project_%s", nameIdentifier)
	//Only one connection can exist per domain per team
	domainName := fmt.Sprintf("%s.example.com", nameIdentifier)

	//Build required pre-req config. AD User Sync Tasks Settings require AD Connection and Project
	preConfig := createTestAccADUserSyncTaskSettingsPreConfig(adConnectionName, projectName, domainName)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccADUserSyncTaskCheckDestroy(adConnectionResourceName),
		Steps: []resource.TestStep{
			{
				Config: createTestAccDataSourceADUserSyncTaskSettingsListInitConfig(preConfig, adUserSyncTaskNamePrefix),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceResourceName, fmt.Sprintf("%s.#", attributes.ADUserSyncTaskSettingsList), "3"),
					resource.TestMatchResourceAttr(dataSourceResourceName, fmt.Sprintf("%s.0.%s", attributes.ADUserSyncTaskSettingsList, attributes.Name), reTaskName),
					resource.TestMatchResourceAttr(dataSourceResourceName, fmt.Sprintf("%s.1.%s", attributes.ADUserSyncTaskSettingsList, attributes.Name), reTaskName),
					resource.TestMatchResourceAttr(dataSourceResourceName, fmt.Sprintf("%s.2.%s", attributes.ADUserSyncTaskSettingsList, attributes.Name), reTaskName),
				),
			},
		},
	})
}

const testAccDataSourceADUserSyncTaskSettingsListInitFormat = `
resource "oktapam_ad_user_sync_task_settings" "test_acc_ad_user_sync_task_settings_1" {
    connection_id            = oktapam_ad_connection.test_acc_ad_connection.id
    name                     = "%[1]s"
    is_active                = true
    frequency                = 12 # Every 12 hours Note: If 24 hours then start_hour_utc is required
    sid_field                = "objectSID"
    upn_field                = "userPrincipalName"
    base_dn                  = "dc=tilt,dc=scaleft,dc=com"
    ldap_query_filter        = "(objectclass=user)"
}

resource "oktapam_ad_user_sync_task_settings" "test_acc_ad_user_sync_task_settings_2" {
    depends_on               = [oktapam_ad_user_sync_task_settings.test_acc_ad_user_sync_task_settings_1]
    connection_id            = oktapam_ad_connection.test_acc_ad_connection.id
    name                     = "%[2]s"
    is_active                = true
    frequency                = 12 # Every 12 hours Note: If 24 hours then start_hour_utc is required
    sid_field                = "objectSID1"
    upn_field                = "userPrincipalName1"
    base_dn                  = "dc=tilt,dc=scaleft,dc=com"
    ldap_query_filter        = "(objectclass=user)"
}

resource "oktapam_ad_user_sync_task_settings" "test_acc_ad_user_sync_task_settings_3" {
    depends_on               = [oktapam_ad_user_sync_task_settings.test_acc_ad_user_sync_task_settings_2]
    connection_id            = oktapam_ad_connection.test_acc_ad_connection.id
    name                     = "%[3]s"
    is_active                = true
    frequency                = 12 # Every 12 hours Note: If 24 hours then start_hour_utc is required
    sid_field                = "objectSID2"
    upn_field                = "userPrincipalName2"
    base_dn                  = "dc=tilt,dc=scaleft,dc=com"
    ldap_query_filter        = "(objectclass=user)"
}

data "oktapam_ad_user_sync_task_settings_list" "test_acc_data_source_ad_user_sync_task_settings_list" {
    depends_on = [oktapam_ad_user_sync_task_settings.test_acc_ad_user_sync_task_settings_3]
    connection_id = oktapam_ad_connection.test_acc_ad_connection.id
}
`

func createTestAccDataSourceADUserSyncTaskSettingsListInitConfig(preConfig string, adUserSyncTaskNamePrefix string) string {
	logging.Debugf("creating config")
	return preConfig + fmt.Sprintf(testAccDataSourceADUserSyncTaskSettingsListInitFormat, adUserSyncTaskNamePrefix+"_1",
		adUserSyncTaskNamePrefix+"_2", adUserSyncTaskNamePrefix+"_3")
}
