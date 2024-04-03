package oktapam

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"
)

func TestAccResourceCloudConnection(t *testing.T) {
	checkTeamApplicable(t, true)
	resourceName := "oktapam_cloud_connection.test_acc_cloud_connection"
	cloudConnectionName := fmt.Sprintf("test-cloud-connection-%s", randSeq())

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      utils.CreateCheckResourceDestroy(providerCloudConnectionKey, cloudConnectionExists),
		Steps: []resource.TestStep{
			{
				Config: createTestAccCloudConnectionCreateConfig(cloudConnectionName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, attributes.Name, cloudConnectionName),
					resource.TestCheckResourceAttr(resourceName, attributes.CloudConnectionProvider, "aws"),
				),
			},
		},
	})
}

func cloudConnectionExists(id string) (bool, error) {
	client := getLocalClientFromMetadata(testAccProvider.Meta())
	logging.Debugf("Checking if resource deleted %s", id)
	cloudConnection, err := client.GetCloudConnection(context.Background(), id)
	return cloudConnection != nil && cloudConnection.Exists() && err == nil, err
}

func createTestAccCloudConnectionCreateConfig(cloudConnectionName string) string {
	const format = `
	resource "oktapam_cloud_connection" "test_acc_cloud_connection" {
		name = "%s"
		cloud_connection_provider = "aws"
		cloud_connection_details {
			account_id = "123456789012"
			role_arn = "arn:aws:iam::123456789012:role/MyRole"
			external_id = "3c086859-3674-49d8-96b0-f1942047c0dc"
		}
	}
	`
	return fmt.Sprintf(format, cloudConnectionName)
}
