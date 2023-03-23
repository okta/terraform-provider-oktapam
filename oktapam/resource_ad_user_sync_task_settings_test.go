package oktapam

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
)

func TestAccADUserSyncTaskSettings(t *testing.T) {
	adUserSyncTaskSettingsResourceName := "oktapam_ad_user_sync_task_settings.test_acc_ad_user_sync_task_settings"
	adConnectionResourceName := "oktapam_ad_connection.test_acc_ad_connection"

	nameIdentifier := randSeq()
	adUserSyncTaskName := fmt.Sprintf("test_acc_ad_user_sync_task_settings_%s", nameIdentifier)
	adConnectionName := fmt.Sprintf("test_acc_ad_connection_%s", nameIdentifier)
	projectName := fmt.Sprintf("test_acc_project_%s", nameIdentifier)
	//Only one connection can exist per domain per team
	domainName := fmt.Sprintf("%s.example.com", nameIdentifier)

	//Build required pre-req config. AD User Sync Tasks Settings require AD Connection and Project
	preConfig := createTestAccADUserSyncTaskSettingsPreConfig(adConnectionName, projectName, domainName)

	//Update schedule
	updatedFreq := 24
	updatedStartHourUTC := 6

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccADUserSyncTaskCheckDestroy(adConnectionResourceName),
		Steps: []resource.TestStep{
			{
				//Step 1: Create AD User Sync Task Settings
				Config: createTestAccADUserSyncTaskSettingsCreateConfig(preConfig, adUserSyncTaskName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccADUserSyncTaskCheckExists(adUserSyncTaskSettingsResourceName),
					resource.TestCheckResourceAttr(adUserSyncTaskSettingsResourceName, attributes.Name, adUserSyncTaskName),
					resource.TestCheckResourceAttr(adUserSyncTaskSettingsResourceName, attributes.IsActive, "true"),
				),
			},
			{
				//Step 2: Update AD User Sync Task Settings Schedule to 24 hours with additional start hour attribute
				Config: createTestAccADUserSyncTaskSettingsUpdateScheduleConfig(preConfig, adUserSyncTaskName, updatedFreq, updatedStartHourUTC),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(adUserSyncTaskSettingsResourceName, attributes.IsActive, "true"),
					resource.TestCheckResourceAttr(adUserSyncTaskSettingsResourceName, attributes.Frequency, strconv.Itoa(updatedFreq)),
					resource.TestCheckResourceAttr(adUserSyncTaskSettingsResourceName, attributes.StartHourUTC, strconv.Itoa(updatedStartHourUTC)),
				),
			},
			{
				//Step 3: Deactivate AD User Sync Task Settings
				Config: createTestAccADUserSyncTaskSettingsUpdateStateConfig(preConfig, adUserSyncTaskName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(adUserSyncTaskSettingsResourceName, attributes.IsActive, "false"),
				),
			},
			{
				ResourceName:      adUserSyncTaskSettingsResourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccADUserSyncTaskSettingsImportStateId(adUserSyncTaskSettingsResourceName),
			},
		},
	})
}

func testAccADUserSyncTaskCheckExists(adUserSyncTaskSettingsResourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		adUserSyncTaskRS, ok := s.RootModule().Resources[adUserSyncTaskSettingsResourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", adUserSyncTaskSettingsResourceName)
		}
		adConnID := adUserSyncTaskRS.Primary.Attributes[attributes.ADConnectionID]
		adUserSyncTaskSettingsID := adUserSyncTaskRS.Primary.ID

		pamClient := testAccProvider.Meta().(client.OktaPAMClient)
		adUserSyncTaskSettings, err := pamClient.GetADUserSyncTaskSettings(context.Background(), adConnID, adUserSyncTaskSettingsID)
		if err != nil {
			return fmt.Errorf("error getting ad user sync task settings: %w", err)
		} else if !adUserSyncTaskSettings.Exists() {
			return fmt.Errorf("ad user sync task settings: %s does not exist", adUserSyncTaskSettingsID)
		}

		return nil
	}
}

func testAccADUserSyncTaskCheckDestroy(adConnectionResourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		adConnectionRS, ok := s.RootModule().Resources[adConnectionResourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", adConnectionResourceName)
		}
		adConnID := adConnectionRS.Primary.ID
		pamClient := testAccProvider.Meta().(client.OktaPAMClient)
		parameters := client.ListADUserSyncTaskSettingsParameters{}
		adUserSyncTaskList, err := pamClient.ListADUserSyncTaskSettings(context.Background(), adConnID, parameters)
		if err != nil {
			return fmt.Errorf("error getting ad user sync task settings list: %w", err)
		}
		if adUserSyncTaskList == nil {
			return fmt.Errorf("user sync task settings list should not be nil")
		}
		if len(adUserSyncTaskList) != 0 {
			return fmt.Errorf("user sync task settings list for the connection should be 0 len")
		}
		return nil
	}
}

// Query or Create Prerequisites Resources(Gateway, Project, AD Connection) for AD Task Settings
const testAccADUserSyncTaskPreConfigFormat = `

# Query pre-existing gateway. It doesn't matter for acceptance testing if gateway is enabled for RDP.
data "oktapam_gateways" "gateways" {
}

resource "oktapam_ad_connection" "test_acc_ad_connection" {
name                     = "%[1]s"
gateway_id               = data.oktapam_gateways.gateways.gateways[0].id
# Domain name is unique for a team and no two connections can have the same domain. To avoid conflicts adding random suffix
domain                   = "example-%[2]s.com"
service_account_username = "user@test.com"
service_account_password = "password"
use_passwordless         = false
domain_controllers       = ["dc1.example.com", "dc2.example.com"]
}

# Create project with forward_traffic enabled
resource "oktapam_project" "test_acc_project" {
 name                 = "%[3]s"
 next_unix_uid        = 60120
 next_unix_gid        = 63020
 create_server_users  = true
 ssh_certificate_type = "CERT_TYPE_ED25519_01"

 // Forward through gateway, gateway must have matching label
 forward_traffic  = true
 gateway_selector = "env=test"
}
`

func createTestAccADUserSyncTaskSettingsPreConfig(adConnectionName string, projectName string, domainName string) string {
	return fmt.Sprintf(testAccADUserSyncTaskPreConfigFormat, adConnectionName, domainName, projectName)
}

const testAccADUserSyncTaskCreateConfigFormat = `
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
`

func createTestAccADUserSyncTaskSettingsCreateConfig(preConfig string, adUserSyncTaskName string) string {
	logging.Debugf("creating config")
	return preConfig + fmt.Sprintf(testAccADUserSyncTaskCreateConfigFormat, adUserSyncTaskName)
}

const testAccADUserSyncTaskUpdateScheduleConfigFormat = `
resource "oktapam_ad_user_sync_task_settings" "test_acc_ad_user_sync_task_settings" {
 connection_id            = oktapam_ad_connection.test_acc_ad_connection.id
 name                     = "%[1]s"
 is_active                = true
 frequency                = "%[2]d" # Every 12 hours Note: If 24 hours then start_hour_utc is required
 start_hour_utc 		  = "%[3]d"
 base_dn                  = "dc=tilt,dc=scaleft,dc=com"
 ldap_query_filter        = "(objectclass=user)"
}
`

func createTestAccADUserSyncTaskSettingsUpdateScheduleConfig(preConfig string, adUserSyncTaskName string, updatedFreq int, updatedStartHourUTC int) string {
	logging.Debugf("creating config")
	return preConfig +
		fmt.Sprintf(testAccADUserSyncTaskUpdateScheduleConfigFormat, adUserSyncTaskName, updatedFreq, updatedStartHourUTC)
}

const testAccADUserSyncTaskUpdateStateConfigFormat = `
resource "oktapam_ad_user_sync_task_settings" "test_acc_ad_user_sync_task_settings" {
 connection_id            = oktapam_ad_connection.test_acc_ad_connection.id
 name                     = "%[1]s"
 is_active                = false
 frequency                = "24" # If 24 hours then start_hour_utc is required
 start_hour_utc 		  = "6"
 base_dn                  = "dc=tilt,dc=scaleft,dc=com"
 ldap_query_filter        = "(objectclass=user)"
}
`

func createTestAccADUserSyncTaskSettingsUpdateStateConfig(preConfig string, adUserSyncTaskName string) string {
	logging.Debugf("creating config")
	return preConfig +
		fmt.Sprintf(testAccADUserSyncTaskUpdateStateConfigFormat, adUserSyncTaskName)
}

func testAccADUserSyncTaskSettingsImportStateId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("Not found: %s", resourceName)
		}
		return fmt.Sprintf("%s/%s", rs.Primary.Attributes[attributes.ADConnectionID], rs.Primary.Attributes[attributes.ID]), nil
	}
}
