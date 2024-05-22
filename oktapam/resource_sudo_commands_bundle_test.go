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

func TestAccResourceSudoCommandsBundle(t *testing.T) {
	checkTeamApplicable(t, true)
	resourceName := "oktapam_sudo_commands_bundle.test_acc_sudo_commands_bundle"
	sudoCommandsBundleName := fmt.Sprintf("test-sudo-commands-bundle-%s", randSeq())

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      utils.CreateCheckResourceDestroy(providerSudoCommandsBundleKey, sudoCommandsBundleExists),
		Steps: []resource.TestStep{
			{
				Config: createTestAccSudoCommandsBundleCreateConfig(sudoCommandsBundleName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, attributes.Name, sudoCommandsBundleName),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s", attributes.StructuredCommands, attributes.StructuredCommandType), "executable"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s", attributes.StructuredCommands, attributes.StructuredCommand), "/bin/run.sh"),
				),
			},
		},
	})
}

func sudoCommandsBundleExists(id string) (bool, error) {
	client := getLocalClientFromMetadata(testAccProvider.Meta())
	logging.Debugf("Checking if resource deleted %s", id)
	scb, err := client.GetSudoCommandsBundle(context.Background(), id)
	return scb != nil && scb.Exists() && err == nil, err
}

func createTestAccSudoCommandsBundleCreateConfig(name string) string {
	const format = `
	resource "oktapam_sudo_commands_bundle" "test_acc_sudo_commands_bundle" {
		name = "%s"
		structured_commands {
			structured_command {
				command = "/bin/run.sh"
				command_type = "executable"
			}
		}
	}
	`
	return fmt.Sprintf(format, name)
}
