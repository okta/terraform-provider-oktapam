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
	resourceName := "test_acc_cloud_connection"
	externalId := uuid.New().String()
	initConfig := createTestAccDataSourceCloudConnectionInitConfig(identifier, externalId)
	fetchConfig := testAccDataSourceCloudConnectionConfig("cloud-connections", identifier+"-1", resourceName)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccV6ProviderFactories,
		CheckDestroy:             testAccCloudConnectionsCheckDestroy(identifier + "-1"),
		Steps: []resource.TestStep{
			{
				Config: initConfig,
			},
			{
				Config: fmt.Sprintf("%s\n%s", initConfig, fetchConfig),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.oktapam_cloud_connection."+resourceName, attributes.Name, fmt.Sprintf("%s-1", identifier)),
					resource.TestCheckResourceAttr("data.oktapam_cloud_connection."+resourceName, fmt.Sprintf("%s.0.aws.0.%s", attributes.CloudConnectionDetails, attributes.CloudConnectionAccountId), "123456789012"),
					resource.TestCheckResourceAttr("data.oktapam_cloud_connection."+resourceName, fmt.Sprintf("%s.0.aws.0.%s", attributes.CloudConnectionDetails, attributes.CloudConnectionRoleARN), "arn:aws:iam::123456789012:role/MyRole"),
					resource.TestCheckResourceAttr("data.oktapam_cloud_connection."+resourceName, fmt.Sprintf("%s.0.aws.0.%s", attributes.CloudConnectionDetails, attributes.CloudConnectionExternalId), externalId),
				),
			},
		},
	})
}

func createTestAccDataSourceCloudConnectionInitConfig(identifier, externalId string) string {
	const format = `
	resource "oktapam_cloud_connection" "test-cloud-connection-1" {
		name = "%s-1"
		cloud_connection_details {
			aws {
				account_id = "123456789012"
				role_arn = "arn:aws:iam::123456789012:role/MyRole"
				external_id = "%s"
			}
		}
	}
	`
	return fmt.Sprintf(format, identifier, externalId)
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
