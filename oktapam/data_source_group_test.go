package oktapam

import (
	"fmt"
	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/constants/attributes"
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
		groupNamePrefix + "-1": {
			Name:  utils.AsStringPtr(groupNamePrefix + "-1"),
			Roles: make([]string, 0),
		},
		groupNamePrefix + "-2": {
			Name:  utils.AsStringPtr(groupNamePrefix + "-2"),
			Roles: []string{"access_user"},
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: createTestAccDatasourceGroupInitConfig(groupNamePrefix),
			},
			{
				Config: createTestAccDatasourceGroupsConfig(groupNamePrefix),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.#", attributes.Groups), "2"),
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
		mappings, err := getIndexMappingFromResource(rs, attributes.Groups, attributes.Name, len(expectedGroups))
		if err != nil {
			return fmt.Errorf("error mapping resources to indices: %w", err)
		}

		primaryAttributes := rs.Primary.Attributes
		for name, group := range expectedGroups {
			idx, ok := mappings[name]
			if !ok {
				return fmt.Errorf("could not find resource with %s: %s", attributes.Name, name)
			}

			name, ok := primaryAttributes[fmt.Sprintf("%s.%s.%s", attributes.Groups, idx, attributes.Name)]
			if !ok {
				return fmt.Errorf("%s attribute not set for group %q", attributes.Name, name)
			}
			if name != *group.Name {
				return fmt.Errorf("%s attributes did not match for group %q, expected %q, got %q", attributes.Name, name, *group.Name, name)
			}

			expectedNumRoles := len(group.Roles)
			numRoles, ok := primaryAttributes[fmt.Sprintf("%s.%s.%s.#", attributes.Groups, idx, attributes.Roles)]
			if !ok {
				return fmt.Errorf("%s attribute not set for group %q", attributes.Roles, name)
			}
			if numRoles != fmt.Sprint(expectedNumRoles) {
				return fmt.Errorf("mismatch of number of %s set for group %q, expected %d, got %s", attributes.Roles, name, expectedNumRoles, numRoles)
			}

			roles, err := getArrayFromResource(rs, fmt.Sprintf("%s.%s.%s", attributes.Groups, idx, attributes.Roles), expectedNumRoles)
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

const testAccDatasourceGroupInitConfigFormat = `
resource "oktapam_group" "test-group-1" {
	name = "%s-1"
	roles = []
}
resource "oktapam_group" "test-group-2" {
	name = "%s-2"
	roles = ["access_user"]
}
`

func createTestAccDatasourceGroupInitConfig(groupName string) string {
	return fmt.Sprintf(testAccDatasourceGroupInitConfigFormat, groupName, groupName)
}
