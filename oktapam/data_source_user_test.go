package oktapam

import (
	"fmt"
	"strings"
	"testing"

	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/constants/attributes"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/client"
	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/utils"
)

func TestAccDatasourceServiceUser(t *testing.T) {
	resourceName := "data.oktapam_user.test_users"
	userNamePrefix := "tf-testing-service-user"

	userName1 := userNamePrefix + "1"
	userName2 := userNamePrefix + "2"
	humanName1 := "steven.elleman"

	expectedUsers := map[string]client.User{
		userName1: {
			Name:     utils.AsStringPtr(userName1),
			Status:   utils.AsStringPtr(string(client.UserStatusActive)),
			UserType: utils.AsStringPtr(string(client.UserTypeService)),
		},
		userName2: {
			Name:     utils.AsStringPtr(userName2),
			Status:   utils.AsStringPtr(string(client.UserStatusDisabled)),
			UserType: utils.AsStringPtr(string(client.UserTypeService)),
		},
	}

	activeUsers := map[string]client.User{
		userName1: expectedUsers[userName1],
	}

	disabledUsers := map[string]client.User{
		userName2: expectedUsers[userName2],
	}

	humanUsers := map[string]client.User{
		humanName1: {
			Name:     utils.AsStringPtr(humanName1),
			Status:   utils.AsStringPtr(string(client.UserStatusActive)),
			UserType: utils.AsStringPtr(string(client.UserTypeHuman)),
		},
	}

	steps := []resource.TestStep{
		// Human user
		{
			Config: createTestAccDatasourceHumanUsersConfig(),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.#", attributes.Users), "1"),
				testAccDatasourceServiceUsersCheck(resourceName, humanUsers),
			),
		},
		// Test Contains
		{
			Config: createTestAccDatasourceServiceUsersContainsConfig(userNamePrefix),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.#", attributes.Users), "2"),
				testAccDatasourceServiceUsersCheck(resourceName, expectedUsers),
			),
		},
		// Test StartsWith
		{
			Config: createTestAccDatasourceServiceUsersStartsWithConfig(userNamePrefix),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.#", attributes.Users), "2"),
				testAccDatasourceServiceUsersCheck(resourceName, expectedUsers),
			),
		},
		// Test Statuses
		{
			Config: createTestAccDatasourceServiceUsersStatusConfig(userNamePrefix, []string{string(client.UserStatusActive)}),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.#", attributes.Users), "1"),
				testAccDatasourceServiceUsersCheck(resourceName, activeUsers),
			),
		},
		{
			Config: createTestAccDatasourceServiceUsersStatusConfig(userNamePrefix, []string{string(client.UserStatusDisabled)}),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.#", attributes.Users), "1"),
				testAccDatasourceServiceUsersCheck(resourceName, disabledUsers),
			),
		},
		{
			Config: createTestAccDatasourceServiceUsersStatusConfig(userNamePrefix, []string{string(client.UserStatusDisabled)}),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.#", attributes.Users), "1"),
				testAccDatasourceServiceUsersCheck(resourceName, disabledUsers),
			),
		},
		{
			Config: createTestAccDatasourceServiceUsersStatusConfig(userNamePrefix, []string{string(client.UserStatusActive), string(client.UserStatusDisabled)}),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.#", attributes.Users), "2"),
				testAccDatasourceServiceUsersCheck(resourceName, expectedUsers),
			),
		},
	}
	// Status steps

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		Steps:             steps,
	})
}

func testAccDatasourceServiceUsersCheck(rn string, expectedUsers map[string]client.User) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[rn]
		if !ok {
			return fmt.Errorf("resource not found: %s", rn)
		}
		mappings, err := getIndexMappingFromResource(rs, attributes.Users, attributes.Name, len(expectedUsers))
		if err != nil {
			return fmt.Errorf("error mapping resources to indices: %w", err)
		}

		primaryAttributes := rs.Primary.Attributes
		for name, user := range expectedUsers {
			idx, ok := mappings[name]
			if !ok {
				return fmt.Errorf("could not find resource with %s: %s", attributes.Name, name)
			}

			name, ok := primaryAttributes[fmt.Sprintf("%s.%s.%s", attributes.Users, idx, attributes.Name)]
			if !ok {
				return fmt.Errorf("%s attribute not set for service user %q", attributes.Name, name)
			}
			if name != *user.Name {
				return fmt.Errorf("%s attributes did not match for service user %q, expected %q, got %q", attributes.Name, name, *user.Name, name)
			}
		}

		return nil
	}
}

const testAccDatasourceHumanUsersFormat = `
data "oktapam_user" "test_users" {
	user_type = "human"
	contains = "ste"
}
`

const testAccDatasourceServiceUsersContainsFormat = `
data "oktapam_user" "test_users" {
	user_type = "service"
	contains = "%s"
}
`

const testAccDatasourceServiceUsersStartsWithFormat = `
data "oktapam_user" "test_users" {
	user_type = "service"
	starts_with = "%s"
}
`

const testAccDatasourceServiceUsersStatusFormat = `
data "oktapam_user" "test_users" {
	user_type = "service"
	starts_with = "%s"
	status = [%s]
}
`

func createTestAccDatasourceHumanUsersConfig() string {
	return fmt.Sprintf(testAccDatasourceHumanUsersFormat)
}

func createTestAccDatasourceServiceUsersContainsConfig(contains string) string {
	return fmt.Sprintf(testAccDatasourceServiceUsersContainsFormat, contains)
}

func createTestAccDatasourceServiceUsersStartsWithConfig(starsWith string) string {
	return fmt.Sprintf(testAccDatasourceServiceUsersStartsWithFormat, starsWith)
}

func createTestAccDatasourceServiceUsersStatusConfig(prefix string, statuses []string) string {
	for i, status := range statuses {
		status = fmt.Sprintf("%q", status)
		statuses[i] = status
	}

	status := strings.Join(statuses, ",")
	return fmt.Sprintf(testAccDatasourceServiceUsersStatusFormat, prefix, status)
}
