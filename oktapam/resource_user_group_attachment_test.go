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

func TestAccUserGroupAttachment(t *testing.T) {
	resourceName := "oktapam_user_group_attachment.test_user_group_attachment"
	groupName := fmt.Sprintf("test_acc_user_group_attachment_%s", randSeq())
	username := fmt.Sprintf("attachment_user_%s", randSeq())

	roleName := "access_user"
	if isExecutingPAMTest() {
		roleName = "end_user"
	}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccUserGroupAttachmentDestroy(groupName, username),
		Steps: []resource.TestStep{
			{
				Config: createTestAccUserGroupAttachmentCreateConfig(groupName, roleName, username),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccUserGroupAttachmentCheckExists(resourceName, groupName, username),
					resource.TestCheckResourceAttr(
						resourceName, attributes.Group, groupName,
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.Username, username,
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.ID, fmt.Sprintf("%s|%s", groupName, username),
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

func testAccUserGroupAttachmentCheckExists(rn string, expectedGroupName string, expectedUsername string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[rn]
		if !ok {
			return fmt.Errorf("resource not found: %s", rn)
		}

		resourceID := rs.Primary.ID
		if resourceID == "" {
			return fmt.Errorf("resource id is not set")
		}

		c := getSDKClientFromMetadata(testAccProvider.Meta())

		if hasAttachment, err := client.GroupContainsUser(context.Background(), c, expectedGroupName, expectedUsername); err != nil {
			return err
		} else if !hasAttachment {
			return fmt.Errorf("user-group attachment is not present")
		}

		return nil
	}
}

func testAccUserGroupAttachmentDestroy(groupName string, username string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		c := getSDKClientFromMetadata(testAccProvider.Meta())

		if hasAttachment, err := client.GroupContainsUser(context.Background(), c, groupName, username); err != nil {
			return err
		} else if hasAttachment {
			return fmt.Errorf("user-group attachment is still present")
		}

		return nil
	}
}

const testAccUserGroupAttachmentCreateConfigFormat = `
resource "oktapam_group" "test_group" {
	name = "%s"
	roles = ["%s"]
}
resource "oktapam_user" "test_user" {
	name = "%s"
	status = "ACTIVE"
	user_type = "service"
}
resource "oktapam_user_group_attachment" "test_user_group_attachment" {
	group = oktapam_group.test_group.name
	username = oktapam_user.test_user.name
}
`

func createTestAccUserGroupAttachmentCreateConfig(groupName string, groupRole string, username string) string {
	return fmt.Sprintf(testAccUserGroupAttachmentCreateConfigFormat, groupName, groupRole, username)
}
