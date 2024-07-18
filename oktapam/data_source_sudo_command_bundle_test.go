package oktapam

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
)

func TestAccDataSourceSudoCommandBundle(t *testing.T) {
	checkTeamApplicable(t, true)
	identifier := randSeq()
	resourceName := "test_acc_sudo_commands_bundle"
	initConfig := createTestAccDataSourceSudoCommandBundleInitConfig(identifier)
	fetchConfig := testAccDataSourceSudoCommandBundleConfig("sudo-commands-bundle", identifier+"-1", resourceName)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccSudoCommandBundlesCheckDestroy(identifier + "-1"),
		Steps: []resource.TestStep{
			{
				Config: initConfig,
			},
			{
				Config: fmt.Sprintf("%s\n%s", initConfig, fetchConfig),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.oktapam_sudo_commands_bundle."+resourceName, attributes.Name, fmt.Sprintf("%s-1", identifier)),
					resource.TestCheckResourceAttr("data.oktapam_sudo_commands_bundle."+resourceName, fmt.Sprintf("%s.0.%s", attributes.StructuredCommands, attributes.StructuredCommandType), "executable"),
					resource.TestCheckResourceAttr("data.oktapam_sudo_commands_bundle."+resourceName, fmt.Sprintf("%s.0.%s", attributes.StructuredCommands, attributes.StructuredCommand), "/bin/run.sh"),
					resource.TestCheckResourceAttr("data.oktapam_sudo_commands_bundle."+resourceName, fmt.Sprintf("%s.0.%s", attributes.StructuredCommands, attributes.StructuredCommandArgsType), "custom"),
					resource.TestCheckResourceAttr("data.oktapam_sudo_commands_bundle."+resourceName, fmt.Sprintf("%s.0.%s", attributes.StructuredCommands, attributes.StructuredCommandArgs), "ls"),
				),
			},
		},
	})
}

func createTestAccDataSourceSudoCommandBundleInitConfig(identifier string) string {
	const format = `
	resource "oktapam_sudo_commands_bundle" "test-sudo-commands-bundle-1" {
		name = "%s-1"
		structured_commands {
			command       = "/bin/run.sh"
			command_type  = "executable"
            args_type = "custom"
			args = "ls"
		}
	}
	`
	return fmt.Sprintf(format, identifier)
}

func testAccDataSourceSudoCommandBundleConfig(resourceName, name, SudoCommandsBundleName string) string {
	const format = `
	data "oktapam_sudo_commands_bundles" "%s" {
		name = "%s"
	}
	
	data "oktapam_sudo_commands_bundle" "%s" {
		id = data.oktapam_sudo_commands_bundles.%s.ids[0]
	}
	`
	return fmt.Sprintf(format, resourceName, name, SudoCommandsBundleName, resourceName)
}
