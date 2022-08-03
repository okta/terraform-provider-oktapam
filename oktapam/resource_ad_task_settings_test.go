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

func TestAccADTaskSettings(t *testing.T) {
	adTaskResourceName := "oktapam_ad_task_settings.test_acc_ad_task_settings"
	adConnectionResourceName := "oktapam_ad_connection.test_acc_ad_connection"

	nameIdentifier := randSeq(10)
	adTaskName := fmt.Sprintf("test_acc_ad_task_settings_%s", nameIdentifier)
	adConnectionName := fmt.Sprintf("test_acc_ad_connection_%s", nameIdentifier)
	projectName := fmt.Sprintf("test_acc_project_%s", nameIdentifier)
	domainName := fmt.Sprintf("%s.example.com", nameIdentifier)

	//Update schedule
	updatedFreq := 24
	updatedStartHourUTC := 6

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccADTaskCheckDestroy(adConnectionResourceName, adTaskResourceName),
		Steps: []resource.TestStep{
			{
				//Step 1: Create AD Task Settings
				Config: createTestAccADTaskSettingsCreateConfig(adConnectionName, projectName, adTaskName, domainName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccADTaskCheckExists(adConnectionResourceName, adTaskResourceName),
					resource.TestCheckResourceAttr(adTaskResourceName, attributes.Name, adTaskName),
					resource.TestCheckResourceAttr(adTaskResourceName, attributes.IsActive, "true"),
					//Check if there is exactly one ad rule assignment
					//.#: Number of elements in list or set
					resource.TestCheckResourceAttr(adTaskResourceName, fmt.Sprintf("%s.#", attributes.ADRuleAssignments), "1"),
				),
			},
			{
				//Step 2: Update AD Task Settings Schedule to 24 hours with additional start hour attribute
				Config: createTestAccADTaskSettingsUpdateScheduleConfig(adTaskName, projectName, adTaskName, domainName, updatedFreq, updatedStartHourUTC),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(adTaskResourceName, attributes.IsActive, "true"),
					resource.TestCheckResourceAttr(adTaskResourceName, attributes.Frequency, strconv.Itoa(updatedFreq)),
					resource.TestCheckResourceAttr(adTaskResourceName, attributes.StartHourUTC, strconv.Itoa(updatedStartHourUTC)),
					//Check if there is exactly one ad rule assignment
					//.#: Number of elements in list or set
					resource.TestCheckResourceAttr(adTaskResourceName, fmt.Sprintf("%s.#", attributes.ADRuleAssignments), "1"),
				),
			},
			{
				//Step 3: Update AD Task Rules. AD Task Rule Assignments is immutable attribute and force creating new AD Task Settings resource
				Config: createTestAccADTaskSettingsUpdateRulesConfig(adTaskName, projectName, domainName, adTaskName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(adTaskResourceName, attributes.IsActive, "true"),
					//Check if there are exactly two ad rule assignments
					//.#: Number of elements in list or set
					resource.TestCheckResourceAttr(adTaskResourceName, fmt.Sprintf("%s.#", attributes.ADRuleAssignments), "2"),
				),
			},
			{
				//Step 4: Deactivate AD Task Settings
				Config: createTestAccADTaskSettingsUpdateStatusConfig(adTaskName, projectName, adTaskName, domainName, updatedFreq, updatedStartHourUTC),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(adTaskResourceName, attributes.IsActive, "false"),
				),
			},
			{
				ResourceName:      adTaskResourceName,
				ImportState:       true,
				ImportStateVerify: true,
				//Ignoring attributes for diff comparison. These attributes are passed while creation but not returned while reading resource state back
				ImportStateVerifyIgnore: []string{attributes.ADConnectionID},
			},
		},
	})
}

func testAccADTaskCheckExists(adConnectionResourceName string, adTaskSettingsResourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		adTaskRS, ok := s.RootModule().Resources[adTaskSettingsResourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", adTaskSettingsResourceName)
		}
		adConnId, adTaskSettingsId, err := parseADTaskSettingsResourceID(adTaskRS.Primary.ID)
		if err != nil {
			return err
		}

		pamClient := testAccProvider.Meta().(client.OktaPAMClient)
		adTaskSettings, err := pamClient.GetADTaskSettings(context.Background(), adConnId, adTaskSettingsId)
		if err != nil {
			return fmt.Errorf("error getting ad task settings: %w", err)
		} else if !adTaskSettings.Exists() {
			return fmt.Errorf("ad task settings does not exist")
		}

		return nil
	}
}

func testAccADTaskCheckDestroy(adConnectionResourceName string, adTaskSettingsResourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		//Get ad connection id and ad task id
		adTaskRS, ok := s.RootModule().Resources[adTaskSettingsResourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", adTaskSettingsResourceName)
		}
		adConnId, adTaskSettingsId, err := parseADTaskSettingsResourceID(adTaskRS.Primary.ID)
		if err != nil {
			return err
		}

		pamClient := testAccProvider.Meta().(client.OktaPAMClient)
		adTask, err := pamClient.GetADTaskSettings(context.Background(), adConnId, adTaskSettingsId)
		if err != nil {
			return fmt.Errorf("error getting ad task settings: %w", err)
		}
		if adTask != nil && adTask.Exists() {
			return fmt.Errorf("ad task settings: %s still exists", adTaskSettingsId)
		}

		return nil
	}
}

// Query or Create Prerequisites Resources(Gateway, Project, AD Connection) for AD Task Settings
const testAccADTaskPreConfigFormat = `

# Query pre-existing gateway. It doesn't matter for acceptance testing if gateway is enabled for RDP.
data "oktapam_gateways" "gateways" {
}

resource "oktapam_ad_connection" "test_acc_ad_connection" {
name                     = "%[1]s"
gateway_id               = data.oktapam_gateways.gateways.gateways[0].id
# Domain name is unique for a team and no two connections can have the same domain. To avoid conflicts adding random suffix
domain                   = "example-%[3]s.com"
service_account_username = "user@test.com"
service_account_password = "password"
use_passwordless         = false
domain_controllers       = ["dc1.example.com", "dc2.example.com"]
}

# Create project with forward_traffic enabled
resource "oktapam_project" "test_acc_project" {
 name                 = "%[2]s"
 next_unix_uid        = 60120
 next_unix_gid        = 63020
 create_server_users  = true
 ssh_certificate_type = "CERT_TYPE_ED25519_01"

 // Forward through gateway, gateway must have matching label
 forward_traffic  = true
 gateway_selector = "env=test"
}
`

const testAccADTaskCreateConfigFormat = `
resource "oktapam_ad_task_settings" "test_acc_ad_task_settings" {
 connection_id            = oktapam_ad_connection.test_acc_ad_connection.id
 name                     = "%[1]s"
 is_active                = true
 frequency                = 12 # Every 12 hours Note: If 24 hours then start_hour_utc is required
 host_name_attribute      = "dNSHostName"
 access_address_attribute = "dNSHostName"
 os_attribute             = "operatingSystem"
 rule_assignments {
   base_dn           = "dc=example,dc=com"
   ldap_query_filter = "(objectclass=computer)"
   project_id        = oktapam_project.test_acc_project.project_id
   priority          = 1
 } 
 # Force dependency on project so that resources destroy works correctly. Project Id is fetched from data and not resource.	
 depends_on = [oktapam_project.test_acc_project]
}
`
func createTestAccADTaskSettingsCreateConfig(adConnectionName string, projectName string, adTaskName string, domainName string) string {
	logging.Debugf("creating config")
	return fmt.Sprintf(testAccADTaskPreConfigFormat, adConnectionName, projectName, domainName) +
		fmt.Sprintf(testAccADTaskCreateConfigFormat, adTaskName)
}

const testAccADTaskUpdateScheduleConfigFormat = `
resource "oktapam_ad_task_settings" "test_acc_ad_task_settings" {
 connection_id            = oktapam_ad_connection.test_acc_ad_connection.id
 name                     = "%[1]s"
 is_active                = true
 frequency                = "%[2]d" # Every 12 hours Note: If 24 hours then start_hour_utc is required
 start_hour_utc 		  = "%[3]d"
 host_name_attribute      = "dNSHostName"
 access_address_attribute = "dNSHostName"
 os_attribute             = "operatingSystem"
 rule_assignments {
   base_dn           = "dc=example,dc=com"
   ldap_query_filter = "(objectclass=computer)"
   project_id        = oktapam_project.test_acc_project.project_id
   priority          = 1
 }
 # Force dependency on project so that resources destroy works correctly. Project Id is fetched from data and not resource.	
 depends_on = [oktapam_project.test_acc_project]
}
`

func createTestAccADTaskSettingsUpdateScheduleConfig(adConnectionName string, projectName string, adTaskName string, domainName string, updatedFreq int, updatedStartHourUTC int) string {
	logging.Debugf("creating config")
	return fmt.Sprintf(testAccADTaskPreConfigFormat, adConnectionName, projectName, domainName) +
		fmt.Sprintf(testAccADTaskUpdateScheduleConfigFormat, adTaskName, updatedFreq, updatedStartHourUTC)
}

const testAccADTaskUpdateStatusConfigFormat = `
resource "oktapam_ad_task_settings" "test_acc_ad_task_settings" {
 connection_id            = oktapam_ad_connection.test_acc_ad_connection.id
 name                     = "%[1]s"
 is_active                = false
 frequency                = "%[2]d" # Every 12 hours Note: If 24 hours then start_hour_utc is required
 start_hour_utc 		  = "%[3]d"
 host_name_attribute      = "dNSHostName"
 access_address_attribute = "dNSHostName"
 os_attribute             = "operatingSystem"
 rule_assignments {
   base_dn           = "dc=example,dc=com"
   ldap_query_filter = "(objectclass=computer)"
   project_id        = oktapam_project.test_acc_project.project_id
   priority          = 1
 }
 # Force dependency on project so that resources destroy works correctly. Project Id is fetched from data and not resource.	
 depends_on = [oktapam_project.test_acc_project]
}
`

func createTestAccADTaskSettingsUpdateStatusConfig(adConnectionName string, projectName string, adTaskName string, domainName string, updatedFreq int, updatedStartHourUTC int) string {
	logging.Debugf("creating config")
	return fmt.Sprintf(testAccADTaskPreConfigFormat, adConnectionName, projectName, domainName) +
		fmt.Sprintf(testAccADTaskUpdateStatusConfigFormat, adTaskName, updatedFreq, updatedStartHourUTC)
}

const testAccADTaskUpdateRulesConfigFormat = `
resource "oktapam_ad_task_settings" "test_acc_ad_task_settings" {
 connection_id            = oktapam_ad_connection.test_acc_ad_connection.id
 name                     = "%[1]s"
 is_active                = true
 frequency                = "24" # Every 12 hours Note: If 24 hours then start_hour_utc is required
 start_hour_utc 		  = "6"
 host_name_attribute      = "dNSHostName"
 access_address_attribute = "dNSHostName"
 os_attribute             = "operatingSystem"
 rule_assignments {
   base_dn           = "dc=example,dc=com"
   ldap_query_filter = "(objectclass=computer)"
   project_id        = oktapam_project.test_acc_project.project_id
   priority          = 1
 }
 rule_assignments {
   base_dn           = "dc=example,dc=com"
   ldap_query_filter = "(objectclass=computer)"
   project_id        = oktapam_project.test_acc_project.project_id
   priority          = 2
 }
 # Force dependency on project so that resources destroy works correctly. Project Id is fetched from data and not resource.	
 depends_on = [oktapam_project.test_acc_project]
}
`

func createTestAccADTaskSettingsUpdateRulesConfig(adConnectionName string, projectName string, adTaskName string, domainName string) string {
	logging.Debugf("creating config")
	return fmt.Sprintf(testAccADTaskPreConfigFormat, adConnectionName, projectName, domainName) +
		fmt.Sprintf(testAccADTaskUpdateRulesConfigFormat, adTaskName)
}
