package oktapam

import (
	"context"
	"fmt"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/kylelemons/godebug/pretty"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
)

func TestAccGroup(t *testing.T) {
	resourceName := "oktapam_group.test-group"
	groupName := fmt.Sprintf("test-acc-group-%s", randSeq(10))
	initialGroup := client.Group{
		Name:  &groupName,
		Roles: make([]string, 0),
	}
	updatedGroup := client.Group{
		Name:  &groupName,
		Roles: []string{"access_user"},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		ProviderFactories:    testAccProviders,
		CheckDestroy: testAccGroupCheckDestroy(groupName),
		Steps: []resource.TestStep{
			{
				Config: createTestAccGroupCreateConfig(groupName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccGroupCheckExists(resourceName, initialGroup),
					resource.TestCheckResourceAttr(
						resourceName, attributes.Name, groupName,
					),
				),
			},
			{
				Config: createTestAccGroupUpdateConfig(groupName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccGroupCheckExists(resourceName, updatedGroup),
					resource.TestCheckResourceAttr(
						resourceName, attributes.Name, groupName,
					),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccGroupCheckExists(rn string, expectedGroup client.Group) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[rn]
		if !ok {
			return fmt.Errorf("resource not found: %s", rn)
		}

		resourceID := rs.Primary.ID
		if resourceID == "" {
			return fmt.Errorf("resource id not set")
		}
		if *expectedGroup.Name != resourceID {
			return fmt.Errorf("resource id not set to expected value.  expected %s, got %s", *expectedGroup.Name, resourceID)
		}

		client := testAccProvider.Meta().(client.OktaPAMClient)
		group, err := client.GetGroup(context.Background(), *expectedGroup.Name, false)
		if err != nil {
			return fmt.Errorf("error getting group :%w", err)
		}
		if group == nil {
			return fmt.Errorf("group %s does not exist", *expectedGroup.Name)
		}
		expectedGroup.ID = group.ID
		comparison := pretty.Compare(group, expectedGroup)
		if comparison != "" {
			return fmt.Errorf("expected group does not match returned group.\n%s", comparison)
		}
		return nil
	}
}

func testAccGroupCheckDestroy(groupName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(client.OktaPAMClient)
		group, err := client.GetGroup(context.Background(), groupName, false)
		if err != nil {
			return fmt.Errorf("error getting group: %w", err)
		}

		if group != nil && group.Exists() {
			return fmt.Errorf("group still exists")
		}

		return nil
	}
}

const testAccGroupCreateConfigFormat = `
resource "oktapam_group" "test-group" {
	name = "%s"
	roles = []
}
`

func createTestAccGroupCreateConfig(groupName string) string {
	return fmt.Sprintf(testAccGroupCreateConfigFormat, groupName)
}

const testAccGroupUpdateConfigFormat = `
resource "oktapam_group" "test-group" {
	name = "%s"
	roles = ["access_user"]
}
`

func createTestAccGroupUpdateConfig(groupName string) string {
	return fmt.Sprintf(testAccGroupUpdateConfigFormat, groupName)
}
