package oktapam

import (
	"context"
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/constants/attributes"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/kylelemons/godebug/pretty"
	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/client"
)

func TestAccServiceUser(t *testing.T) {
	resourceName := "oktapam_service_user.test-user"
	userName := "resource-test3"
	teamName := "asa" // TODO: Change to tf-provider-testing team
	userType := string(client.UserTypeService)

	constructUser := func(status client.UserStatus) client.User {
		sstatus := string(status)
		return client.User{
			Name:     &userName,
			Status:   &sstatus,
			TeamName: &teamName,
			UserType: &userType,
		}
	}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccServiceUserCheckDestroy(userName), // TODO: can't hard-delete service user
		Steps: []resource.TestStep{
			{
				Config: createTestAccServiceUserUpdateConfig(userName, client.UserStatusActive), // TODO: does this require POST?
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccServiceUserCheckExists(resourceName, constructUser(client.UserStatusActive)),
					resource.TestCheckResourceAttr(
						resourceName, attributes.Name, userName,
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.Status, string(client.UserStatusActive),
					),
				),
			},
			{
				Config: createTestAccServiceUserUpdateConfig(userName, client.UserStatusDisabled),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccServiceUserCheckExists(resourceName, constructUser(client.UserStatusDisabled)),
					resource.TestCheckResourceAttr(
						resourceName, attributes.Name, userName,
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.Status, string(client.UserStatusDisabled),
					),
				),
			},
			{
				Config: createTestAccServiceUserUpdateConfig(userName, client.UserStatusDeleted),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccServiceUserCheckExists(resourceName, constructUser(client.UserStatusDeleted)),
					resource.TestCheckResourceAttr(
						resourceName, attributes.Name, userName,
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.Status, string(client.UserStatusDeleted),
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

func testAccServiceUserCheckExists(rn string, expectedUser client.User) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[rn]
		if !ok {
			return fmt.Errorf("resource not found: %s", rn)
		}

		resourceID := rs.Primary.ID
		if resourceID == "" {
			return fmt.Errorf("resource id not set")
		}
		if expectedUser.Name == nil || *expectedUser.Name != resourceID {
			return fmt.Errorf("resource id not set to expected value. expected %s, got %s", *expectedUser.Name, resourceID)
		}

		c := testAccProvider.Meta().(client.OktaPAMClient)
		user, err := c.GetServiceUser(context.Background(), *expectedUser.Name)
		if err != nil {
			return fmt.Errorf("error getting service user :%w", err)
		}
		if user == nil {
			return fmt.Errorf("service user %s does not exist", *expectedUser.Name)
		}
		expectedUser.ID = user.ID
		expectedUser.DeletedAt = user.DeletedAt
		comparison := pretty.Compare(user, expectedUser)
		if comparison != "" {
			return fmt.Errorf("expected service user does not match returned service user.\n%s", comparison)
		}
		return nil
	}
}

// Users cannot be hard-deleted, therefore check if the status is deleted
func testAccServiceUserCheckDestroy(userName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		c := testAccProvider.Meta().(client.OktaPAMClient)
		user, err := c.GetServiceUser(context.Background(), userName)
		if err != nil {
			return fmt.Errorf("error getting service user: %w", err)
		}

		if user == nil {
			return fmt.Errorf("service user does not exist")
		}

		if *user.Status != string(client.UserStatusDeleted) {
			return fmt.Errorf("service user is not deleted")
		}

		return nil
	}
}

// Create and update are the same
func createTestAccServiceUserUpdateConfig(userName string, status client.UserStatus) string {
	return fmt.Sprintf(testAccServiceUserUpdateConfigFormat, userName, string(status))
}

const testAccServiceUserUpdateConfigFormat = `
resource "oktapam_service_user" "test-user" {
	name = "%s"
    status = "%s"
}
`
