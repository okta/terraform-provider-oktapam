package oktapam

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/errors"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/typed_strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/kylelemons/godebug/pretty"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
)

func TestAccUser(t *testing.T) {
	resource1 := "test_user_1"
	resource2 := "test_user_2"
	resourceType := "oktapam_user"
	resourceName1 := fmt.Sprintf("%s.%s", resourceType, resource1)
	identifier := randSeq()
	userName := "tf_acceptance_test_user_" + identifier
	teamName := DefaultTestTeam
	userType := typed_strings.UserTypeService

	constructUser := func(status typed_strings.UserStatus) client.User {
		return client.User{
			Name:     &userName,
			Status:   &status,
			TeamName: &teamName,
			UserType: &userType,
		}
	}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccServiceUserCheckDestroy(userName),
		Steps: []resource.TestStep{
			{
				Config: createTestAccServiceUserUpdateConfig(resource1, userName, typed_strings.UserStatusActive),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccServiceUserCheckExists(resourceName1, constructUser(typed_strings.UserStatusActive)),
					resource.TestCheckResourceAttr(
						resourceName1, attributes.Name, userName,
					),
					resource.TestCheckResourceAttr(
						resourceName1, attributes.UserType, typed_strings.UserTypeService.String(),
					),
					resource.TestCheckResourceAttr(
						resourceName1, attributes.Status, typed_strings.UserStatusActive.String(),
					),
				),
			},
			{
				ResourceName: resourceName1,
				ImportState:  true,
				//Used to dynamically generate the ID from the terraform state.
				//Terraform Resource ID is set to ASA User UUID but read requires "User Name" and "Type" to retrieve the resource
				ImportStateIdFunc: testAccUserImportStateId(resourceName1),
				ImportStateVerify: true,
			},
			{
				Config: createTestAccServiceUserUpdateConfig(resource1, userName, typed_strings.UserStatusDisabled),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccServiceUserCheckExists(resourceName1, constructUser(typed_strings.UserStatusDisabled)),
					resource.TestCheckResourceAttr(
						resourceName1, attributes.Name, userName,
					),
					resource.TestCheckResourceAttr(
						resourceName1, attributes.UserType, typed_strings.UserTypeService.String(),
					),
					resource.TestCheckResourceAttr(
						resourceName1, attributes.Status, typed_strings.UserStatusDisabled.String(),
					),
				),
			},
			{
				Config: createTestAccServiceUserUpdateConfig(resource1, userName, typed_strings.UserStatusDeleted),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccServiceUserCheckExists(resourceName1, constructUser(typed_strings.UserStatusDeleted)),
					resource.TestCheckResourceAttr(
						resourceName1, attributes.Name, userName,
					),
					resource.TestCheckResourceAttr(
						resourceName1, attributes.UserType, typed_strings.UserTypeService.String(),
					),
					resource.TestCheckResourceAttr(
						resourceName1, attributes.Status, typed_strings.UserStatusDeleted.String(),
					),
				),
			},
			{
				// Ensure attempted Human User creation fails
				Config:      createTestAccHumanUserUpdateConfig(resource2, userName, typed_strings.UserStatusActive),
				ExpectError: regexp.MustCompile(errors.HumanUserCreationError),
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

		c := testAccProvider.Meta().(client.OktaPAMClient)
		user, err := c.GetServiceUser(context.Background(), *expectedUser.Name)
		if err != nil {
			return fmt.Errorf("error getting service user :%w", err)
		}
		if user == nil {
			return fmt.Errorf("service user %s does not exist", *expectedUser.Name)
		}
		expectedUser.DeletedAt = user.DeletedAt
		expectedUser.ID = user.ID
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

		if *user.Status != typed_strings.UserStatusDeleted {
			return fmt.Errorf("service user is not deleted")
		}

		return nil
	}
}

// Create and update are the same
func createTestAccServiceUserUpdateConfig(resourceName string, userName string, status typed_strings.UserStatus) string {
	return fmt.Sprintf(testAccUserUpdateConfigFormat, resourceName, userName, status.String(), typed_strings.UserTypeService)
}

func createTestAccHumanUserUpdateConfig(resourceName string, userName string, status typed_strings.UserStatus) string {
	return fmt.Sprintf(testAccUserUpdateConfigFormat, resourceName, userName, status.String(), typed_strings.UserTypeHuman)
}

const testAccUserUpdateConfigFormat = `
resource "oktapam_user" "%s" {
	name = "%s"
	status = "%s"
	user_type = "%s"
}
`

func testAccUserImportStateId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("Not found: %s", resourceName)
		}
		return fmt.Sprintf("%s/%s", rs.Primary.Attributes[attributes.Name], rs.Primary.Attributes[attributes.UserType]), nil
	}
}
