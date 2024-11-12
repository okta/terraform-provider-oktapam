package oktapam

import (
	"context"
	"fmt"
	"testing"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/kylelemons/godebug/pretty"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
)

func TestAccGroup(t *testing.T) {
	resourceName := "oktapam_group.test_group"
	groupName := fmt.Sprintf("test_acc_group_%s", randSeq())
	roleName := "access_user"
	if isExecutingPAMTest() {
		roleName = "end_user"
	}

	initialGroup := client.Group{
		Name:  &groupName,
		Roles: make([]string, 0),
	}
	updatedGroup := client.Group{
		Name:  &groupName,
		Roles: []string{roleName},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccV6ProviderFactories(),
		CheckDestroy:             testAccGroupCheckDestroy(groupName),
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
				Config: createTestAccGroupUpdateConfig(groupName, roleName),
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
				ImportStateIdFunc: testAccGroupImportStateId(resourceName),
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

		client := getTestAccAPIClients().LocalClient
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
		client := getTestAccAPIClients().LocalClient
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
resource "oktapam_group" "test_group" {
	name = "%s"
	roles = []
}
`

func createTestAccGroupCreateConfig(groupName string) string {
	return fmt.Sprintf(testAccGroupCreateConfigFormat, groupName)
}

const testAccGroupUpdateConfigFormat = `
resource "oktapam_group" "test_group" {
	name = "%s"
	roles = ["%s"]
}
`

func createTestAccGroupUpdateConfig(groupName, roleName string) string {
	return fmt.Sprintf(testAccGroupUpdateConfigFormat, groupName, roleName)
}

func testAccGroupImportStateId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("Not found: %s", resourceName)
		}
		return rs.Primary.Attributes[attributes.Name], nil
	}
}
