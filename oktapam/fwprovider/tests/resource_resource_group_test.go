package tests

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"testing"
)

func TestAccResourceGroupBasic(t *testing.T) {
	checkTeamApplicable(t, true)
	_, _, accProviders := testAccFrameworkMuxProviders(context.Background(), t)
	resourceName := "oktapam_resource_group_fwk.test_acc_resource_group"
	initialResourceGroupName := fmt.Sprintf("test_acc_resource_group_%s", randSeq())
	delegatedAdminGroup1Name := fmt.Sprintf("test_acc_resource_group_dga1_%s", randSeq())

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: accProviders,
		Steps: []resource.TestStep{
			{
				Config: createTestAccResourceGroupCreateConfig(delegatedAdminGroup1Name, initialResourceGroupName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						resourceName, attributes.Name, initialResourceGroupName,
					),
				),
			},
		},
	})
}

const testAccResourceGroupCreateConfigFormat = `
resource "oktapam_group" "test_resource_group_dga1_group" {
	name = "%s"
}
resource "oktapam_resource_group_fwk" "test_acc_resource_group" {
	name = "%s"
	description = "initial description"
	delegated_resource_admin_groups = [{
      id = oktapam_group.test_resource_group_dga1_group.id
      name = oktapam_group.test_resource_group_dga1_group.name
      type = "user_group"
   }]
   depends_on = [
   oktapam_group.test_resource_group_dga1_group
 ]
}
`

func createTestAccResourceGroupCreateConfig(dga1Name, resourceGroupName string) string {
	return fmt.Sprintf(testAccResourceGroupCreateConfigFormat, dga1Name, resourceGroupName)
}