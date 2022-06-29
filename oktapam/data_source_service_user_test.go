package oktapam

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/constants/attributes"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/client"
	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/utils"
)

func TestAccDatasourceServiceUser(t *testing.T) {
	resourceName := "data.oktapam_service_user.test_users"
	userNamePrefix := "tf-test"

	expectedUsers := map[string]client.User{
		userNamePrefix + "1": {
			Name:     utils.AsStringPtr(userNamePrefix + "1"),
			Status:   utils.AsStringPtr(string(client.UserStatusActive)),
			UserType: utils.AsStringPtr(string(client.UserTypeService)),
		},
		userNamePrefix + "2": {
			Name:     utils.AsStringPtr(userNamePrefix + "2"),
			Status:   utils.AsStringPtr(string(client.UserStatusActive)),
			UserType: utils.AsStringPtr(string(client.UserTypeService)),
		},
		userNamePrefix + "3": {
			Name:     utils.AsStringPtr(userNamePrefix + "3"),
			Status:   utils.AsStringPtr(string(client.UserStatusActive)),
			UserType: utils.AsStringPtr(string(client.UserTypeService)),
		},
	}

	emptyUsers := map[string]client.User{}

	steps := []resource.TestStep{
		// TODO: Skip creation, because users cannot be deleted
		/*{
			Config: createTestAccDatasourceServiceUserInitConfig(userNamePrefix),
		},*/
		// Test Contains

		{
			Config: createTestAccDatasourceServiceUsersContainsConfig(userNamePrefix),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.#", attributes.ServiceUsers), "3"),
				testAccDatasourceServiceUsersCheck(resourceName, expectedUsers),
			),
		},
		// Test StartsWith
		{
			Config: createTestAccDatasourceServiceUsersStartsWithConfig(userNamePrefix),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.#", attributes.ServiceUsers), "3"),
				testAccDatasourceServiceUsersCheck(resourceName, expectedUsers),
			),
		},
		// Test Statuses
		{
			Config: createTestAccDatasourceServiceUsersStatusConfig(userNamePrefix, string(client.UserStatusActive)),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.#", attributes.ServiceUsers), "3"),
				testAccDatasourceServiceUsersCheck(resourceName, expectedUsers),
			),
		},
		{
			Config: createTestAccDatasourceServiceUsersStatusConfig(userNamePrefix, string(client.UserStatusDisabled)),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.#", attributes.ServiceUsers), "0"),
				testAccDatasourceServiceUsersCheck(resourceName, emptyUsers),
			),
		},
		{
			Config: createTestAccDatasourceServiceUsersStatusConfig(userNamePrefix, string(client.UserStatusDeleted)),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.#", attributes.ServiceUsers), "0"),
				testAccDatasourceServiceUsersCheck(resourceName, emptyUsers),
			),
		},
	}
	// Status steps
	//steps = append(steps, generateStatusSteps(resourceName, userNamePrefix, expectedUsers)...)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		Steps:             steps,
	})
}

// Generates all combinations (Note: not permutations) for 3 boolean values
/*func generateSets() []map[string]bool {
	statuses := []string{"ACTIVE", "DISABLED", "REMOVED"}
	var sets []map[string]bool
	// Not a comprehensive combination solution
	for index1 := 0; index1 < len(statuses); index1++ {
		set := make(map[string]bool)
		for index2 := 0; index2 < len(statuses); index2++ {
			if index1 == index2 {
				set[statuses[index1]] = true
				sets = append(sets, set)
				continue
			}
			set[statuses[index1]] = true
			set[statuses[index2]] = true
			sets = append(sets, set)
		}
	}
	return sets
}

func generateStatusSteps(resourceName, prefix string, expectedUsers map[string]client.ServiceUser) []resource.TestStep {
	statuses := []string{"ACTIVE", "DISABLED", "REMOVED"}
	sets := generateSets()

	fullSet := make(map[string]bool)
	for _, status := range statuses {
		fullSet[status] = true
	}
	sets = append(sets, fullSet)

	// Generate step for each set
	var steps []resource.TestStep
	for _, s := range sets {

		statusLookup := ""
		expectedLength := 0
		for k, v := range s {
			if v {
				statusLookup += fmt.Sprintf("%s,", k)
				expectedLength++
			}
		}

		steps = append(steps,
			resource.TestStep{
				Config: createTestAccDatasourceServiceUsersStatusConfig(prefix, statusLookup),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.#", attributes.ServiceUsers), fmt.Sprintf("%d", expectedLength)),
					testAccDatasourceServiceUsersCheck(resourceName, expectedUsers),
				),
			},
		)
	}
	return steps
}*/

func testAccDatasourceServiceUsersCheck(rn string, expectedUsers map[string]client.User) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[rn]
		if !ok {
			return fmt.Errorf("resource not found: %s", rn)
		}
		mappings, err := getIndexMappingFromResource(rs, attributes.ServiceUsers, attributes.Name, len(expectedUsers))
		if err != nil {
			return fmt.Errorf("error mapping resources to indices: %w", err)
		}

		primaryAttributes := rs.Primary.Attributes
		for name, user := range expectedUsers {
			idx, ok := mappings[name]
			if !ok {
				return fmt.Errorf("could not find resource with %s: %s", attributes.Name, name)
			}

			name, ok := primaryAttributes[fmt.Sprintf("%s.%s.%s", attributes.ServiceUsers, idx, attributes.Name)]
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

const testAccDatasourceServiceUsersContainsFormat = `
data "oktapam_service_user" "test_users" {
	contains = "%s"
}
`

const testAccDatasourceServiceUsersStartsWithFormat = `
data "oktapam_service_user" "test_users" {
	starts_with = "%s"
}
`

const testAccDatasourceServiceUsersStatusFormat = `
data "oktapam_service_user" "test_users" {
	starts_with = "%s"
	status = ["%s"]
}
`

func createTestAccDatasourceServiceUsersContainsConfig(contains string) string {
	return fmt.Sprintf(testAccDatasourceServiceUsersContainsFormat, contains)
}

func createTestAccDatasourceServiceUsersStartsWithConfig(starsWith string) string {
	return fmt.Sprintf(testAccDatasourceServiceUsersStartsWithFormat, starsWith)
}

func createTestAccDatasourceServiceUsersStatusConfig(prefix, status string) string {
	return fmt.Sprintf(testAccDatasourceServiceUsersStatusFormat, prefix, status)
}

/*
// HERE: Just create the values, remove this when you don't need it
const testAccDatasourceServiceUserInitConfigFormat = `
resource "oktapam_service_user" "test-user-1" {
	name = "%s1"
	status = "ACTIVE"
}
resource "oktapam_service_user" "test-user-2" {
	name = "%s2"
	status = "DISABLED"
}
resource "oktapam_service_user" "test-user-3" {
	name = "%s3"
	status = "REMOVED"
}
`

func createTestAccDatasourceServiceUserInitConfig(userName string) string {
	return fmt.Sprintf(testAccDatasourceServiceUserInitConfigFormat, userName, userName, userName)
}*/
