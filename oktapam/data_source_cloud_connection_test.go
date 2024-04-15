package oktapam

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
)

func TestAccDataSourceCloudConnection(t *testing.T) {
	checkTeamApplicable(t, true)
	identifier := randSeq()
	initConfig := createTestAccDataSourceCloudConnectionInitConfig(identifier)
	fetchConfig := testAccDataSourceCloudConnectionConfig("data1", identifier+"-1", "data2")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccCloudConnectionsCheckDestroy(identifier+"-1"),
		Steps: []resource.TestStep{
			{
				Config: initConfig,
			},
			{
				Config: fmt.Sprintf("%s\n%s", initConfig, fetchConfig),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.oktapam_cloud_connection.data2", attributes.Name, fmt.Sprintf("%s-1", identifier)),
					resource.TestCheckResourceAttr("data.oktapam_cloud_connection.data2", attributes.CloudConnectionProvider, "aws"),
				),
			},
		},
	})
}

func createTestAccDataSourceCloudConnectionInitConfig(identifier string) string {
	uuid := uuid.New().String()
	const format = `
	resource "oktapam_cloud_connection" "test-cloud-connection-1" {
		name = "%s-1"
		cloud_connection_provider = "aws"
		cloud_connection_details {
			account_id = "123456789000"
			role_arn = "arn:aws:iam::123456789012:role/MyRole"
			external_id = "%s"
		}
	}
	`
	return fmt.Sprintf(format, identifier, uuid)
}

func testAccDataSourceCloudConnectionConfig(connectionsResourceName, name, cloudConnectionName string) string {
	const format = `
	data "oktapam_cloud_connections" "%s" {
		name = "%s"
	}
	
	data "oktapam_cloud_connection" "%s" {
		id = data.oktapam_cloud_connections.%s.ids[0]
	}
	`
	return fmt.Sprintf(format, connectionsResourceName, name, cloudConnectionName, connectionsResourceName)
}
