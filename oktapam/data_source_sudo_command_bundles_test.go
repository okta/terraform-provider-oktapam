package oktapam

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
)

func TestAccDataSourceSudoCommandBundlesList(t *testing.T) {
	checkTeamApplicable(t, true)
	prefix := "data.oktapam_sudo_command_bundles"
	identifier := randSeq()
	initConfig := createTestAccDataSourceSudoCommandBundlesInitConfig(identifier)

	sudoCommandsBundle1Name := fmt.Sprintf("%s.%s", prefix, "data1")
	sudoCommandsBundle2Name := fmt.Sprintf("%s.%s", prefix, "data2")
	list1Config := testAccDataSourceSudoCommandBundlesConfig("data1", identifier+"-1")
	list2Config := testAccDataSourceSudoCommandBundlesConfig("data2", identifier+"-2")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccV6ProviderFactories,
		CheckDestroy:             testAccSudoCommandBundlesCheckDestroy(identifier+"-1", identifier+"-2"),
		Steps: []resource.TestStep{
			{
				Config: initConfig,
			},
			{
				Config: fmt.Sprintf("%s\n%s", initConfig, list1Config),
				Check:  resource.TestCheckResourceAttr(sudoCommandsBundle1Name, fmt.Sprintf("%s.#", attributes.IDs), "1"),
			},
			{
				Config: fmt.Sprintf("%s\n%s", initConfig, list2Config),
				Check:  resource.TestCheckResourceAttr(sudoCommandsBundle2Name, fmt.Sprintf("%s.#", attributes.IDs), "1"),
			},
		},
	})
}

func createTestAccDataSourceSudoCommandBundlesInitConfig(identifier string) string {
	const format = `
	resource "oktapam_sudo_command_bundle" "test-sudo-commands-bundle-1" {
		name = "%s-1"
		structured_commands {
			command = "/bin/run.sh"
			command_type = "executable"
            args_type = "custom"
			args = "ls"
		}
	}
	resource "oktapam_sudo_command_bundle" "test-sudo-commands-bundle-2" {
		name = "%s-2"
		structured_commands {
			command = "/bin/find"
			command_type = "raw"
		}
	}
	`
	return fmt.Sprintf(format, identifier, identifier)
}

func testAccDataSourceSudoCommandBundlesConfig(resourceName, name string) string {
	const format = `
	data "oktapam_sudo_command_bundles" "%s" {
		name = "%s"
	}
	`
	return fmt.Sprintf(format, resourceName, name)
}

func testAccSudoCommandBundlesCheckDestroy(identifiers ...string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		c := getTestAccAPIClients().SDKClient

		respList, err := client.ListSudoCommandBundles(context.Background(), c)
		if err != nil {
			return fmt.Errorf("error getting sudo commands bundles: %w", err)
		}

		m := make(map[string]bool, len(identifiers))
		for _, id := range identifiers {
			m[id] = true
		}

		for _, scb := range respList {
			if _, ok := m[scb.Name]; ok {
				return fmt.Errorf("sudo commands bundles still exists")
			}
		}

		return nil
	}
}
