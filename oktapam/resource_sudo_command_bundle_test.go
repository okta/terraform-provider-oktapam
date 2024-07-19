package oktapam

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"
)

func TestAccResourceSudoCommandBundle(t *testing.T) {
	checkTeamApplicable(t, true)
	resourceName := "oktapam_sudo_command_bundle.test_acc_sudo_command_bundle"
	sudoCommandBundleName := fmt.Sprintf("test-sudo-command-bundle-%s", randSeq())

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      utils.CreateCheckResourceDestroy(providerSudoCommandBundleKey, sudoCommandBundleExists),
		Steps: []resource.TestStep{
			{
				Config: createTestAccSudoCommandBundleCreateConfig(sudoCommandBundleName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, attributes.Name, sudoCommandBundleName),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s", attributes.StructuredCommands, attributes.StructuredCommandType), "executable"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s", attributes.StructuredCommands, attributes.StructuredCommand), "/bin/run.sh"),
				),
			},
			{
				Config: createTestAccSudoCommandBundleUpdateConfig(sudoCommandBundleName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, attributes.Name, sudoCommandBundleName),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s", attributes.StructuredCommands, attributes.StructuredCommandType), "directory"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s", attributes.StructuredCommands, attributes.StructuredCommand), "/bin/"),
				),
			},
		},
	})
}

func sudoCommandBundleExists(id string) (bool, error) {
	c := getSDKClientFromMetadata(testAccProvider.Meta())
	logging.Debugf("Checking if resource deleted %s", id)
	scb, err := client.GetSudoCommandBundle(context.Background(), c, id)
	return scb != nil && *scb.Id != "" && err == nil, err
}

func createTestAccSudoCommandBundleCreateConfig(name string) string {
	const format = `
	resource "oktapam_sudo_command_bundle" "test_acc_sudo_command_bundle" {
		name = "%s"
		structured_commands {
			command       = "/bin/run.sh"
			command_type  = "executable"
			args_type     = "any"
  		}
	}
	`
	return fmt.Sprintf(format, name)
}

func createTestAccSudoCommandBundleUpdateConfig(name string) string {
	const format = `
	resource "oktapam_sudo_command_bundle" "test_acc_sudo_command_bundle" {
		name = "%s"
		structured_commands {
			command       = "/bin/"
			command_type  = "directory"
  		}
	}
	`
	return fmt.Sprintf(format, name)
}
