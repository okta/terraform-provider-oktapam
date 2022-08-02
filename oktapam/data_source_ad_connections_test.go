package oktapam

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
)

func TestAccDatasourceADConnections(t *testing.T) {
	adConnectionNamePrefix := fmt.Sprintf("test_acc_datasource_ad_connection_%s", randSeq(10))
	adConnectionName1 := fmt.Sprintf("%s_1", adConnectionNamePrefix)

	adConnectionResNamePrefix := "oktapam_ad_connection.test_acc_ad_connection"
	adConnectionResName1 := fmt.Sprintf("%s_1", adConnectionResNamePrefix)

	resourceNamePrefix := "data.oktapam_ad_connections.test_acc_datasource_ad_connections"
	dataSource1 := fmt.Sprintf("%s_1", resourceNamePrefix)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		Steps: []resource.TestStep{
			{
				//Check if datasource returns pre-existing AD Connections
				Config: createTestAccDatasourceADConnectionsConfig(adConnectionName1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSource1, fmt.Sprintf("%s.#", attributes.ADConnections), "1"),
					resource.TestCheckResourceAttrPair(dataSource1, fmt.Sprintf("%s.0.%s", attributes.ADConnections, attributes.Name),
						adConnectionResName1, attributes.Name),
				),
			},
		},
	})
}

const testAccDatasourceADConnectionsConfigFormat = `
data oktapam_gateway "gateways" {
}

resource "oktapam_ad_connection" "test_acc_ad_connection_1" {
	 name                     = "%[1]s"
	 gateway_id               = data.oktapam_gateway.gateways.gateways[0].id
	 domain                   = "testacc.example.com"
	 service_account_username = "user@testacc.example.com"
	 service_account_password = "password"
	 use_passwordless         = false
}

data "oktapam_ad_connections" "test_acc_datasource_ad_connections_1" {
    gateway_id = data.oktapam_gateway.gateways.gateways[0].id
    depends_on = [oktapam_ad_connection.test_acc_ad_connection_1]
}
`

func createTestAccDatasourceADConnectionsConfig(adConnectionName1 string) string {
	return fmt.Sprintf(testAccDatasourceADConnectionsConfigFormat, adConnectionName1)
}
