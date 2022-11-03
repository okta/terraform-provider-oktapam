package oktapam

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
)

func TestAccDatasourceADConnections(t *testing.T) {
	nameIdentifier := randSeq(10)
	prefix := "test_acc_datasource_ad_connection"
	adConnectionTFResourceName := "oktapam_ad_connection.test_acc_ad_connection"
	adConnectionName := fmt.Sprintf("%s-%s", prefix, nameIdentifier)
	domainName := fmt.Sprintf("%s-%s.example.com", prefix, nameIdentifier)
	datasourceTFResourceName := "data.oktapam_ad_connections.test_acc_datasource_ad_connections"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		Steps: []resource.TestStep{
			{
				//Check if datasource returns AD Connections based on the certificateId filter
				Config: createTestAccDatasourceADConnectionsConfig(adConnectionName, domainName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceTFResourceName, fmt.Sprintf("%s.#", attributes.ADConnections), "1"),
					resource.TestCheckResourceAttrPair(datasourceTFResourceName, fmt.Sprintf("%s.0.%s", attributes.ADConnections, attributes.Name),
						adConnectionTFResourceName, attributes.Name),
				),
			},
		},
	})
}

const testAccDatasourceADConnectionsConfigFormat = `
data "oktapam_gateways" "gateways" {
}

resource "oktapam_ad_certificate_request" "self_signed_cert" {
  display_name = "test_acc"
  common_name  = "test_acc"
  type         = "self_signed"
  details {
	ttl_days = 1
  }
}

resource "oktapam_ad_connection" "test_acc_ad_connection" {
	 name                     = "%[1]s"
	 gateway_id               = data.oktapam_gateways.gateways.gateways[0].id
	 domain                   = "%[2]s"
	 service_account_username = "user@testacc.example.com"
	 service_account_password = "password"
	 use_passwordless         = true
     certificate_id  		  = oktapam_ad_certificate_request.self_signed_cert.id
}

data "oktapam_ad_connections" "test_acc_datasource_ad_connections" {
    certificate_id = oktapam_ad_certificate_request.self_signed_cert.id
    depends_on = [oktapam_ad_connection.test_acc_ad_connection, oktapam_ad_certificate_request.self_signed_cert]
}
`

func createTestAccDatasourceADConnectionsConfig(adConnectionName string, domainName string) string {
	return fmt.Sprintf(testAccDatasourceADConnectionsConfigFormat, adConnectionName, domainName)
}
