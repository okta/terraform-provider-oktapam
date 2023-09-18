package oktapam

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
)

func TestAccDatasourceSecurityPolicyFetch(t *testing.T) {
	checkTeamApplicable(t, true)

	// Generate details
	identifier := randSeq()

	// config to create the resources
	initConfig := createTestAccDatasourceSecurityPolicyInitConfig(identifier)

	// config for the datasources
	fetchConfig := testAccDatasourceSecurityPolicyConfig(identifier)

	resourceName := "data.oktapam_security_policy.security_policy"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccSecurityPoliciesCheckDestroy(identifier),
		Steps: []resource.TestStep{
			{
				Config: initConfig,
			},
			{
				Config: fmt.Sprintf("%s\n%s", initConfig, fetchConfig),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, attributes.Name, identifier),
					resource.TestCheckResourceAttr(resourceName, attributes.Description, "terraform test sp"),
					resource.TestCheckResourceAttr(resourceName, attributes.Active, "true"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.#", attributes.Principals), "1"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s.#", attributes.Principals, attributes.Groups), "1"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.#", attributes.Rule), "1"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s.#", attributes.Rule, attributes.Conditions), "1"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s.0.%s.#", attributes.Rule, attributes.Conditions, attributes.AccessRequest), "2"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s.0.%s.0.%s", attributes.Rule, attributes.Conditions, attributes.AccessRequest, attributes.ExpiresAfterSeconds), "1200"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s.0.%s.0.%s", attributes.Rule, attributes.Conditions, attributes.AccessRequest, attributes.RequestTypeId), "abcd"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s.0.%s.0.%s", attributes.Rule, attributes.Conditions, attributes.AccessRequest, attributes.RequestTypeName), "foo"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s.0.%s.1.%s", attributes.Rule, attributes.Conditions, attributes.AccessRequest, attributes.ExpiresAfterSeconds), "1800"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s.0.%s.1.%s", attributes.Rule, attributes.Conditions, attributes.AccessRequest, attributes.RequestTypeId), "wxyz"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s.0.%s.1.%s", attributes.Rule, attributes.Conditions, attributes.AccessRequest, attributes.RequestTypeName), "bar"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s.0.%s.#", attributes.Rule, attributes.Conditions, attributes.Gateway), "0"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s", attributes.Rule, attributes.Name), "first rule"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s.#", attributes.Rule, attributes.Privileges), "1"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s.0.%s.#", attributes.Rule, attributes.Privileges, attributes.PasswordCheckoutRDP), "1"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s.0.%s.0.%s", attributes.Rule, attributes.Privileges, attributes.PasswordCheckoutRDP, attributes.Enabled), "true"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s.0.%s.#", attributes.Rule, attributes.Privileges, attributes.PasswordCheckoutSSH), "1"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s.0.%s.0.%s", attributes.Rule, attributes.Privileges, attributes.PasswordCheckoutSSH, attributes.Enabled), "true"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s.0.%s.#", attributes.Rule, attributes.Privileges, attributes.PrincipalAccountRDP), "1"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s.0.%s.0.%s", attributes.Rule, attributes.Privileges, attributes.PrincipalAccountRDP, attributes.Enabled), "true"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s.0.%s.0.%s", attributes.Rule, attributes.Privileges, attributes.PrincipalAccountRDP, attributes.AdminLevelPermissions), "true"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s.0.%s.#", attributes.Rule, attributes.Privileges, attributes.PrincipalAccountSSH), "1"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s.0.%s.0.%s", attributes.Rule, attributes.Privileges, attributes.PrincipalAccountSSH, attributes.Enabled), "true"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s.0.%s.0.%s", attributes.Rule, attributes.Privileges, attributes.PrincipalAccountSSH, attributes.AdminLevelPermissions), "false"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s.#", attributes.Rule, attributes.Resources), "1"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s.0.%s.#", attributes.Rule, attributes.Resources, attributes.Servers), "1"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s.0.%s.0.%s.#", attributes.Rule, attributes.Resources, attributes.Servers, attributes.LabelSelectors), "1"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s.0.%s.0.%s.0.%s.#", attributes.Rule, attributes.Resources, attributes.Servers, attributes.LabelSelectors, attributes.Accounts), "2"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s.0.%s.0.%s.0.%s.0", attributes.Rule, attributes.Resources, attributes.Servers, attributes.LabelSelectors, attributes.Accounts), "root"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s.0.%s.0.%s.0.%s.1", attributes.Rule, attributes.Resources, attributes.Servers, attributes.LabelSelectors, attributes.Accounts), "pamadmin"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s.0.%s.0.%s.0.%s.system.os_type", attributes.Rule, attributes.Resources, attributes.Servers, attributes.LabelSelectors, attributes.ServerLabels), "linux"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s.0.%s.0.%s.#", attributes.Rule, attributes.Resources, attributes.Servers, attributes.Server), "1"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s.0.%s.0.%s.0.%s", attributes.Rule, attributes.Resources, attributes.Servers, attributes.Server, attributes.ServerID), "9103335f-1147-407b-84d7-dbfc57f75c99"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.0.%s.0.%s.0.%s.#", attributes.Rule, attributes.Resources, attributes.Servers, attributes.ServerAccount), "0"),
				),
			},
		},
	})
}

const testAccDatasourceSecurityPolicyCreateConfigFormat = `
resource "oktapam_group" "test_group" {
	name = "sp-test-group-%s"
}
resource "oktapam_security_policy" "test_ds_security_policies" {
	name = "%s"
	description = "terraform test sp"
	active = true
	principals {
		groups = [oktapam_group.test_group.id]
	}
	rule {
		name = "first rule"
		resources {
			servers {
				server {
					server_id = "9103335f-1147-407b-84d7-dbfc57f75c99"
				}
				label_selectors {
					server_labels = {
						"system.os_type" = "linux"
					}
					accounts = ["root", "pamadmin"]
				}
			}
		}
		privileges {
			password_checkout_rdp {
				enabled = true
			}
			password_checkout_ssh {
				enabled = true
			}
			principal_account_rdp {
				enabled = true
				admin_level_permissions = true
			}
			principal_account_ssh {
				enabled = true
				admin_level_permissions = false
			}
		}
		conditions {
			access_request {
				request_type_id = "abcd"
				request_type_name = "foo"
				expires_after_seconds = 1200
			}
			access_request {
				request_type_id = "wxyz"
				request_type_name = "bar"
				expires_after_seconds = 1800
			}
		}
	}
}
`

func createTestAccDatasourceSecurityPolicyInitConfig(identifier string) string {
	return fmt.Sprintf(testAccDatasourceSecurityPolicyCreateConfigFormat, identifier, identifier)
}

const testAccDatasourceSecurityPolicyFormat = `
data "oktapam_security_policies" "security_policies" {
	name = "%s"
}

data "oktapam_security_policy" "security_policy" {
	id = data.oktapam_security_policies.security_policies.ids[0]
}
`

func testAccDatasourceSecurityPolicyConfig(name string) string {
	return fmt.Sprintf(testAccDatasourceSecurityPolicyFormat, name)
}
