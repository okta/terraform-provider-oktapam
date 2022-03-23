package oktapam

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/client"
	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/utils"
)

func TestAccDatasourceGroup(t *testing.T) {
	resourceName := "data.oktapam_group.test_groups"
	groupNamePrefix := fmt.Sprintf("test-acc-datasource-group-%s", randSeq(10))
	expectedGroups := map[string]client.Group{
		groupNamePrefix + "-one": {
			Name:  utils.AsStringPtr(groupNamePrefix + "-one"),
			Roles: make([]string, 0),
		},
		groupNamePrefix + "-two": {
			Name:  utils.AsStringPtr(groupNamePrefix + "-two"),
			Roles: []string{"access_user"},
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: createTestAccDatasourceGroupCreateConfig(groupNamePrefix),
			},
			{
				Config: createTestAccDatasourceGroupsConfig(groupNamePrefix),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "groups.#", "2"),
					testAccDatasourceGroupsCheck(resourceName, expectedGroups),
				),
			},
		},
	})
}

func testAccDatasourceGroupsCheck(rn string, expectedGroups map[string]client.Group) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[rn]
		if !ok {
			return fmt.Errorf("resource not found: %s", rn)
		}
		mappings, err := getIndexMappingFromResource(rs, "groups", "name", len(expectedGroups))
		if err != nil {
			return fmt.Errorf("error mapping resources to indices: %w", err)
		}

		attributes := rs.Primary.Attributes
		for name, group := range expectedGroups {
			idx, ok := mappings[name]
			if !ok {
				return fmt.Errorf("could not find resource with name: %s", name)
			}

			name, ok := attributes[fmt.Sprintf("groups.%s.name", idx)]
			if !ok {
				return fmt.Errorf("name attribute not set for group %q", name)
			}
			if name != *group.Name {
				return fmt.Errorf("name attributes did not match for group %q, expected %q, got %q", name, *group.Name, name)
			}

			expectedNumRoles := len(group.Roles)
			numRoles, ok := attributes[fmt.Sprintf("groups.%s.roles.#", idx)]
			if !ok {
				return fmt.Errorf("roles attribute not set for group %q", name)
			}
			if numRoles != fmt.Sprint(expectedNumRoles) {
				return fmt.Errorf("mismatch of number of roles set for group %q, expected %d, got %s", name, expectedNumRoles, numRoles)
			}

			roles, err := getArrayFromResource(rs, fmt.Sprintf("groups.%s.roles", idx), expectedNumRoles)
			if err != nil {
				return fmt.Errorf("error while retrieving roles from state: %w", err)
			}

			for _, role := range group.Roles {
				if !arrayContainsString(roles, role) {
					return fmt.Errorf("expected to have role %q set for group %q", role, name)
				}
			}

		}

		return nil
	}
}

const testAccDatasourceGroupsConfigFormat = `
data "oktapam_group" "test_groups" {
	contains = "%s"
}
`

func createTestAccDatasourceGroupsConfig(prefix string) string {
	return fmt.Sprintf(testAccDatasourceGroupsConfigFormat, prefix)
}

const testAccDatasourceGroupCreateConfigFormat = `
resource "oktapam_group" "test-group-one" {
	name = "%s-one"
	roles = []
}
resource "oktapam_group" "test-group-two" {
	name = "%s-two"
	roles = ["access_user"]
}
`

func createTestAccDatasourceGroupCreateConfig(groupName string) string {
	return fmt.Sprintf(testAccDatasourceGroupCreateConfigFormat, groupName, groupName)
}
