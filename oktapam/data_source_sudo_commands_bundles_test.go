package oktapam

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
)

func TestAccDataSourceSudoCommandsBundlesList(t *testing.T) {
	checkTeamApplicable(t, true)
	prefix := "data.oktapam_sudo_commands_bundles"
	identifier := randSeq()
	initConfig := createTestAccDataSourceSudoCommandsBundlesInitConfig(identifier)

	sudoCommandsBundle1Name := fmt.Sprintf("%s.%s", prefix, "data1")
	sudoCommandsBundle2Name := fmt.Sprintf("%s.%s", prefix, "data2")
	list1Config := testAccDataSourceSudoCommandsBundlesConfig("data1", identifier+"-1")
	list2Config := testAccDataSourceSudoCommandsBundlesConfig("data2", identifier+"-2")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccSudoCommandsBundlesCheckDestroy(identifier+"-1", identifier+"-2"),
		Steps: []resource.TestStep{
			{
				Config: initConfig,
			},
			{
				Config: fmt.Sprintf("%s\n%s", initConfig, list1Config),
				Check:  resource.TestCheckResourceAttr(sudoCommandsBundle1Name, fmt.Sprintf("%s.#", attributes.IDs), "2"),
			},
			{
				Config: fmt.Sprintf("%s\n%s", initConfig, list2Config),
				Check:  resource.TestCheckResourceAttr(sudoCommandsBundle2Name, fmt.Sprintf("%s.#", attributes.IDs), "2"),
			},
		},
	})
}

func createTestAccDataSourceSudoCommandsBundlesInitConfig(identifier string) string {
	const format = `
	resource "oktapam_sudo_commands_bundle" "test-sudo-commands-bundle-1" {
		name = "%s-1"
		structured_commands {
			command = "/bin/run.sh"
			command_type = "executable"
            args_type = "custom"
			args = "ls"
		}
	}

	resource "oktapam_sudo_commands_bundle" "test-sudo-commands-bundle-2" {
		name = "%s-2"
		structured_commands {
			command = "/bin/find"
			command_type = "raw"
		}
	}
	`
	return fmt.Sprintf(format, identifier, identifier)
}

func testAccDataSourceSudoCommandsBundlesConfig(resourceName, name string) string {
	const format = `
	data "oktapam_sudo_commands_bundles" "%s" {
		name = "%s"
	}
	`
	return fmt.Sprintf(format, resourceName, name)
}

func testAccSudoCommandsBundlesCheckDestroy(identifiers ...string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		c := getSDKClientFromMetadata(testAccProvider.Meta())

		resp, err := client.ListSudoCommandsBundles(context.Background(), c)
		if err != nil {
			return fmt.Errorf("error getting sudo commands bundles: %w", err)
		}

		m := make(map[string]bool, len(identifiers))
		for _, id := range identifiers {
			m[id] = true
		}

		for _, scb := range resp.GetList() {
			if _, ok := m[scb.Name]; ok {
				return fmt.Errorf("sudo commands bundles still exists")
			}
		}

		return nil
	}
}
