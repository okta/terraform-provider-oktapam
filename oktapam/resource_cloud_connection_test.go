package oktapam

import (
	"context"
	"fmt"
	"testing"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/config"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"
)

func TestAccResourceCloudConnection(t *testing.T) {
	checkTeamApplicable(t, true)
	resourceName := "oktapam_cloud_connection.test_acc_cloud_connection"
	cloudConnectionName := fmt.Sprintf("test-cloud-connection-%s", randSeq())
	uuid := uuid.New().String()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccV6ProviderFactories,
		CheckDestroy:             utils.CreateCheckResourceDestroy(config.ProviderCloudConnectionKey, cloudConnectionExists),
		Steps: []resource.TestStep{
			{
				Config: createTestAccCloudConnectionCreateConfig(cloudConnectionName, uuid),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, attributes.Name, cloudConnectionName),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.aws.0.%s", attributes.CloudConnectionDetails, attributes.CloudConnectionAccountId), "123456789012"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.aws.0.%s", attributes.CloudConnectionDetails, attributes.CloudConnectionRoleARN), "arn:aws:iam::123456789012:role/MyRole"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.aws.0.%s", attributes.CloudConnectionDetails, attributes.CloudConnectionExternalId), uuid),
				),
			},
		},
	})
}

func cloudConnectionExists(id string) (bool, error) {
	client := getTestAccAPIClients().LocalClient
	logging.Debugf("Checking if resource deleted %s", id)
	cloudConnection, err := client.GetCloudConnection(context.Background(), id)
	return cloudConnection != nil && cloudConnection.Exists() && err == nil, err
}

func createTestAccCloudConnectionCreateConfig(cloudConnectionName string, uuid string) string {
	const format = `
	resource "oktapam_cloud_connection" "test_acc_cloud_connection" {
		name = "%s"
		cloud_connection_details {
			aws {
				account_id = "123456789012"
				role_arn = "arn:aws:iam::123456789012:role/MyRole"
				external_id = "%s"
			}
		}
	}
	`
	return fmt.Sprintf(format, cloudConnectionName, uuid)
}
