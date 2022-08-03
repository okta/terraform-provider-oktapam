package oktapam

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"
)

func TestAccADConnection(t *testing.T) {
	resourceName := "oktapam_ad_connection.test_acc_ad_connection"
	nameIdentifier := randSeq(10)
	connectionName := fmt.Sprintf("test_acc_ad_connection-%s", nameIdentifier)
	domainName := fmt.Sprintf("%s.example.com", nameIdentifier)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      utils.CreateCheckResourceDestroy(providerADConnectionKey, adConnectionExists),
		Steps: []resource.TestStep{
			{
				Config: createTestAccADConnectionCreateConfig(connectionName, domainName),
				Check: resource.ComposeAggregateTestCheckFunc(
					utils.CheckResourceExists(resourceName, adConnectionExists),
					resource.TestCheckResourceAttr(
						resourceName, attributes.Name, connectionName,
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.Domain, domainName,
					),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.#", attributes.DomainControllers), "2"),
				),
			},
			{
				Config: createTestAccADConnectionUpdateConfig(connectionName, domainName),
				Check: resource.ComposeAggregateTestCheckFunc(
					utils.CheckResourceExists(resourceName, adConnectionExists),
					resource.TestCheckResourceAttr(
						resourceName, attributes.Name, connectionName,
					),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.#", attributes.DomainControllers), "3"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				//Ignoring attributes for diff comparison. These attributes are passed while creation but not returned while reading resource state back
				ImportStateVerifyIgnore: []string{attributes.ServiceAccountPassword, attributes.UsePasswordless},
			},
		},
	})
}

func adConnectionExists(id string) (bool, error) {
	client := testAccProvider.Meta().(client.OktaPAMClient)
	logging.Debugf("Checking if resource deleted %s", id)
	adConnection, err := client.GetADConnection(context.Background(), id, false)

	return adConnection != nil && adConnection.Exists() && err == nil, err
}

const testAccADConnectionCreateConfigFormat = `
data "oktapam_gateways" "gateways" {
}

resource "oktapam_ad_connection" "test_acc_ad_connection" {
 name                     = "%[1]s"
 gateway_id               = data.oktapam_gateways.gateways.gateways[0].id
 domain                   = "%[2]s"
 service_account_username = "user@example.com"
 service_account_password = "password"
 use_passwordless         = false
 domain_controllers       = ["dc1.example.com", "dc2.example.com"]
}
`

func createTestAccADConnectionCreateConfig(adConnectionName string, domainName string) string {
	logging.Debugf("creating config")
	return fmt.Sprintf(testAccADConnectionCreateConfigFormat, adConnectionName, domainName)
}

const testAccADConnectionUpdateConfigFormat = `
data "oktapam_gateways" "gateways" {
}

resource "oktapam_ad_connection" "test_acc_ad_connection" {
name                     = "%[1]s"
gateway_id               = data.oktapam_gateways.gateways.gateways[0].id
domain                   = "%[2]s"
service_account_username = "account@example.com"
service_account_password = "password"
use_passwordless         = false
domain_controllers       = ["dc1.example.com", "dc2.example.com", "dc3.example.com"]
}`

func createTestAccADConnectionUpdateConfig(adConnectionName string, domainName string) string {
	return fmt.Sprintf(testAccADConnectionUpdateConfigFormat, adConnectionName, domainName)
}
