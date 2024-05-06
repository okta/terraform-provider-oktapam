package oktapam

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
)

func TestAccDataSourceCloudConnectionsList(t *testing.T) {
	checkTeamApplicable(t, true)
	prefix := "data.oktapam_cloud_connections"
	identifier := randSeq()
	initConfig := createTestAccDataSourceCloudConnectionsInitConfig(identifier)

	cloudConnection1Name := fmt.Sprintf("%s.%s", prefix, "data1")
	cloudConnection2Name := fmt.Sprintf("%s.%s", prefix, "data2")
	list1Config := testAccDataSourceCloudConnectionsConfig("data1", identifier+"-1")
	list2Config := testAccDataSourceCloudConnectionsConfig("data2", identifier+"-2")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccCloudConnectionsCheckDestroy(identifier+"-1", identifier+"-2"),
		Steps: []resource.TestStep{
			{
				Config: initConfig,
			},
			{
				Config: fmt.Sprintf("%s\n%s", initConfig, list1Config),
				Check:  resource.TestCheckResourceAttr(cloudConnection1Name, fmt.Sprintf("%s.#", attributes.IDs), "1"),
			},
			{
				Config: fmt.Sprintf("%s\n%s", initConfig, list2Config),
				Check:  resource.TestCheckResourceAttr(cloudConnection2Name, fmt.Sprintf("%s.#", attributes.IDs), "1"),
			},
		},
	})
}

func createTestAccDataSourceCloudConnectionsInitConfig(identifier string) string {
	uuid := uuid.New().String()
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

	resource "oktapam_cloud_connection" "test-cloud-connection-2" {
		name = "%s-2"
		cloud_connection_details {
			aws {
				account_id = "123456789012"
				role_arn = "arn:aws:iam::123456789012:role/MyRole"
				external_id = "%s"
			}
		}
	}
	`
	return fmt.Sprintf(format, identifier, uuid, identifier, uuid)
}

func testAccDataSourceCloudConnectionsConfig(resourceName, name string) string {
	const format = `
	data "oktapam_cloud_connections" "%s" {
		name = "%s"
	}
	`
	return fmt.Sprintf(format, resourceName, name)
}

func testAccCloudConnectionsCheckDestroy(identifiers ...string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		c := getLocalClientFromMetadata(testAccProvider.Meta())

		cloudConnections, err := c.ListCloudConnections(context.Background())
		if err != nil {
			return fmt.Errorf("error getting cloud connections: %w", err)
		}

		m := make(map[string]bool, len(identifiers))
		for _, id := range identifiers {
			m[id] = true
		}

		for _, rg := range cloudConnections {
			if _, ok := m[*rg.Name]; ok {
				return fmt.Errorf("cloud connections still exists")
			}
		}

		return nil
	}
}
