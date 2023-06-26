package oktapam

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
)

func TestAccDatasourceSecurityPoliciesList(t *testing.T) {
	checkTeamApplicable(t, true)
	prefix := "data.oktapam_security_policies"

	// Generate details
	identifier := randSeq()

	// config to create the resources
	initConfig := createTestAccDatasourceSecurityPoliciesInitConfig(identifier)

	// config for the datasources
	dataSource1Name := fmt.Sprintf("%s.%s", prefix, "data1")
	dataSource2Name := fmt.Sprintf("%s.%s", prefix, "data2")
	list1Config := testAccDatasourceSecurityPoliciesConfig("data1", identifier+"-1")
	list2Config := testAccDatasourceSecurityPoliciesConfig("data2", identifier+"-2")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccSecurityPoliciesCheckDestroy(identifier+"-1", identifier+"-2"),
		Steps: []resource.TestStep{
			{
				Config: initConfig,
			},
			{
				Config: fmt.Sprintf("%s\n%s", initConfig, list1Config),
				Check:  resource.TestCheckResourceAttr(dataSource1Name, fmt.Sprintf("%s.#", attributes.IDs), "1"),
			},
			{
				Config: fmt.Sprintf("%s\n%s", initConfig, list2Config),
				Check:  resource.TestCheckResourceAttr(dataSource2Name, fmt.Sprintf("%s.#", attributes.IDs), "1"),
			},
		},
	})
}

func testAccSecurityPoliciesCheckDestroy(identifiers ...string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		c := testAccProvider.Meta().(client.OktaPAMClient)

		securityPolicies, err := c.ListSecurityPolicies(context.Background())
		if err != nil {
			return fmt.Errorf("error getting security policies: %w", err)
		}

		m := make(map[string]bool, len(identifiers))
		for _, id := range identifiers {
			m[id] = true
		}

		for _, rg := range securityPolicies {
			if _, ok := m[*rg.Name]; ok {
				return fmt.Errorf("security policies still exists")
			}
		}

		return nil
	}
}

const testAccDatasourceSecurityPoliciesCreateConfigFormat = `
resource "oktapam_group" "test_group" {
	name = "sp-test-group-%s"
}
resource "oktapam_security_policy" "test_ds_security_policies_1" {
	name = "%s-1"
	description = "first description"
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
			}
			principal_account_ssh {
				enabled = true
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
resource "oktapam_security_policy" "test_ds_security_policies_2" {
	name = "%s-2"
	description = "second description"
	active = true
	principals {
		groups = [oktapam_group.test_group.id]
	}
	rule {
		name = "first rule"
		resources {
			servers {
				server_account {
					server_id = "9103335f-1147-407b-84d7-dbfc57f75c99"
					account   = "pamadmin"
				}
				label_selectors {
					server_labels = {
						"system.os_type" = "windows"
					}
					accounts = ["pamadmin"]
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
				enabled = false
			}
			principal_account_ssh {
				enabled = false
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
			gateway {
				traffic_forwarding = true
				session_recording  = true
			}
		}
	}
}
`

func createTestAccDatasourceSecurityPoliciesInitConfig(identifier string) string {
	return fmt.Sprintf(testAccDatasourceSecurityPoliciesCreateConfigFormat, identifier, identifier, identifier)
}

const testAccDatasourceSecurityPoliciesFormat = `
data "oktapam_security_policies" "%s" {
	name = "%s"
}
`

func testAccDatasourceSecurityPoliciesConfig(resourceName, name string) string {
	return fmt.Sprintf(testAccDatasourceSecurityPoliciesFormat, resourceName, name)
}
