package oktapam

//TODO: Commenting this because can't create gatway using terraform and has dependency on it
//import (
//	"context"
//	"fmt"
//	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
//	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/constants/attributes"
//	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/logging"
//	"testing"
//
//	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
//	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/client"
//	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/utils"
//)
//
//func TestAccADConnection(t *testing.T) {
//	resourceName := "oktapam_ad_connection.test-ad-connection"
//	connectionName := fmt.Sprintf("test-acc-ad-connection-%s", randSeq(10))
//
//	resource.Test(t, resource.TestCase{
//		PreCheck:          func() { testAccPreCheck(t) },
//		ProviderFactories: testAccProviders,
//		CheckDestroy:      utils.CreateCheckResourceDestroy(providerADConnectionKey, adConnectionExists),
//		Steps: []resource.TestStep{
//			{
//				Config: createTestAccADConnectionCreateConfig(connectionName),
//				Check: resource.ComposeAggregateTestCheckFunc(
//					utils.CheckResourceExists(resourceName, adConnectionExists),
//					resource.TestCheckResourceAttr(
//						resourceName, attributes.Name, connectionName,
//					),
//					resource.TestCheckResourceAttr(
//						resourceName, attributes.Domain, "test.com",
//					),
//				),
//			},
//			{
//				Config: createTestAccADConnectionUpdateConfig(connectionName),
//				Check: resource.ComposeAggregateTestCheckFunc(
//					utils.CheckResourceExists(resourceName, adConnectionExists),
//					resource.TestCheckResourceAttr(
//						resourceName, attributes.Name, connectionName,
//					),
//					resource.TestCheckResourceAttr(
//						resourceName, attributes.Domain, "updated.test.com",
//					),
//				),
//			},
//			{
//				ResourceName:      resourceName,
//				ImportState:       true,
//				ImportStateVerify: false,
//			},
//		},
//	})
//}
//
//func adConnectionExists(id string) (bool, error) {
//	client := testAccProvider.Meta().(client.OktaPAMClient)
//	logging.Debugf("Checking if resource deleted %s", id)
//	adConnection, err := client.GetADConnection(context.Background(), id, false)
//
//	return adConnection != nil && adConnection.Exists() && err == nil, err
//}
//
//const testAccADConnectionCreateConfigFormat = `
//data oktapam_gateway "gateways" {
//}
//
//resource "oktapam_ad_connection" "test-ad-connection" {
//name                     = "%s"
//gateway_id               = data.oktapam_gateway.gateways.gateways[0].id
//domain                   = "test.com"
//service_account_username = "account@test.com"
//service_account_password = "password"
//use_passwordless         = false
//domain_controllers       = ["dc1.test.com", "dc2.test.com"]
//}`
//
//func createTestAccADConnectionCreateConfig(adConnectionName string) string {
//	logging.Debugf("creating config")
//	return fmt.Sprintf(testAccADConnectionCreateConfigFormat, adConnectionName)
//}
//
//const testAccADConnectionUpdateConfigFormat = `
//data oktapam_gateway "gateways" {
//}
//
//resource "oktapam_ad_connection" "test-ad-connection" {
//name                     = "%s"
//gateway_id               = data.oktapam_gateway.gateways.gateways[0].id
//domain                   = "updated.test.com"
//service_account_username = "account@test.com"
//service_account_password = "password"
//use_passwordless         = false
//domain_controllers       = ["dc1.test.com", "dc2.test.com"]
//}`
//
//func createTestAccADConnectionUpdateConfig(adConnectionName string) string {
//	return fmt.Sprintf(testAccADConnectionUpdateConfigFormat, adConnectionName)
//}
//
//func testAccADConnectionCheckDestroy() resource.TestCheckFunc {
//	return func(s *terraform.State) error {
//		for _, rs := range s.RootModule().Resources {
//			if rs.Type != providerADConnectionKey {
//				continue
//			}
//			id := rs.Primary.ID
//
//			client := testAccProvider.Meta().(client.OktaPAMClient)
//			logging.Debugf("Checking if resource deleted %s", id)
//			adConnection, err := client.GetADConnection(context.Background(), id, false)
//
//			if err != nil {
//				return fmt.Errorf("error getting ad connection: %w", err)
//			}
//
//			if adConnection != nil && adConnection.Exists() {
//				return fmt.Errorf("adConnection still exists")
//			}
//		}
//
//		return nil
//	}
//}
