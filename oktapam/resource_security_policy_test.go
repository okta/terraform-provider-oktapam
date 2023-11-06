package oktapam

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/kylelemons/godebug/pretty"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"
)

func TestAccSecurityPolicy(t *testing.T) {
	checkTeamApplicable(t, true)
	resourceName := "oktapam_security_policy.test_acc_security_policy"
	securityPolicyName := fmt.Sprintf("test_acc_security_policy_%s", randSeq())
	group1Name := fmt.Sprintf("test_acc_security_policy_group1_%s", randSeq())
	group2Name := fmt.Sprintf("test_acc_security_policy_group2_%s", randSeq())

	initialSecurityPolicy := &client.SecurityPolicy{
		Name:        &securityPolicyName,
		Active:      utils.AsBoolPtr(true),
		Description: utils.AsStringPtr("test description"),
		Principals: &client.SecurityPolicyPrincipals{
			UserGroups: []client.NamedObject{
				{
					Name: utils.AsStringPtr(group1Name),
				},
				{
					Name: utils.AsStringPtr(group2Name),
				},
			},
		},
		Rules: []*client.SecurityPolicyRule{
			{
				Name:         utils.AsStringPtr("first rule"),
				ResourceType: client.ServerBasedResourceSelectorType,
				ResourceSelector: &client.ServerBasedResourceSelector{
					Selectors: []client.ServerBasedResourceSubSelectorContainer{
						{
							SelectorType: client.IndividualServerSubSelectorType,
							Selector: &client.IndividualServerSubSelector{
								Server: client.NamedObject{
									Id: utils.AsStringPtr("9103335f-1147-407b-84d7-dbfc57f75c99"),
								},
							},
						},
						{
							SelectorType: client.ServerLabelServerSubSelectorType,
							Selector: &client.ServerLabelBasedSubSelector{
								ServerSelector: &client.ServerLabelServerSelector{
									Labels: map[string]string{
										"system.os_type": "linux",
									},
								},
								AccountSelectorType: client.UsernameAccountSelectorType,
								AccountSelector: &client.UsernameAccountSelector{
									Usernames: []string{"root", "pamadmin"},
								},
							},
						},
					},
				},
				Privileges: []*client.SecurityPolicyRulePrivilegeContainer{
					{
						PrivilegeType: client.PasswordCheckoutRDPPrivilegeType,
						PrivilegeValue: &client.PasswordCheckoutRDPPrivilege{
							Enabled: utils.AsBoolPtr(true),
						},
					},
					{
						PrivilegeType: client.PasswordCheckoutSSHPrivilegeType,
						PrivilegeValue: &client.PasswordCheckoutSSHPrivilege{
							Enabled: utils.AsBoolPtr(true),
						},
					},
					{
						PrivilegeType: client.PrincipalAccountRDPPrivilegeType,
						PrivilegeValue: &client.PrincipalAccountRDPPrivilege{
							Enabled:               utils.AsBoolPtr(true),
							AdminLevelPermissions: utils.AsBoolPtr(true),
						},
					},
					{
						PrivilegeType: client.PrincipalAccountSSHPrivilegeType,
						PrivilegeValue: &client.PrincipalAccountSSHPrivilege{
							Enabled:               utils.AsBoolPtr(true),
							AdminLevelPermissions: utils.AsBoolPtr(true),
						},
					},
				},
				Conditions: []*client.SecurityPolicyRuleConditionContainer{
					{
						ConditionType: client.AccessRequestConditionType,
						ConditionValue: &client.AccessRequestCondition{
							RequestTypeID:       utils.AsStringPtr("abcd"),
							RequestTypeName:     utils.AsStringPtr("foo"),
							ExpiresAfterSeconds: utils.AsIntPtr(1200),
						},
					},
					{
						ConditionType: client.AccessRequestConditionType,
						ConditionValue: &client.AccessRequestCondition{
							RequestTypeID:       utils.AsStringPtr("wxyz"),
							RequestTypeName:     utils.AsStringPtr("bar"),
							ExpiresAfterSeconds: utils.AsIntPtr(1800),
						},
					},
				},
			},
		},
	}

	updatedSecurityPolicy := &client.SecurityPolicy{
		Name:        &securityPolicyName,
		Active:      utils.AsBoolPtr(true),
		Description: utils.AsStringPtr("updated description"),
		Principals: &client.SecurityPolicyPrincipals{
			UserGroups: []client.NamedObject{
				{
					Name: utils.AsStringPtr(group1Name),
				},
			},
		},
		Rules: []*client.SecurityPolicyRule{
			{
				Name:         utils.AsStringPtr("first rule"),
				ResourceType: client.ServerBasedResourceSelectorType,
				ResourceSelector: &client.ServerBasedResourceSelector{
					Selectors: []client.ServerBasedResourceSubSelectorContainer{
						{
							SelectorType: client.IndividualServerAccountSubSelectorType,
							Selector: &client.IndividualServerAccountSubSelector{
								Server: client.NamedObject{
									Id: utils.AsStringPtr("9103335f-1147-407b-84d7-dbfc57f75c99"),
								},
								Username: utils.AsStringPtr("pamadmin"),
							},
						},
						{
							SelectorType: client.ServerLabelServerSubSelectorType,
							Selector: &client.ServerLabelBasedSubSelector{
								ServerSelector: &client.ServerLabelServerSelector{
									Labels: map[string]string{
										"system.os_type": "windows",
									},
								},
								AccountSelectorType: client.UsernameAccountSelectorType,
								AccountSelector: &client.UsernameAccountSelector{
									Usernames: []string{"pamadmin"},
								},
							},
						},
					},
				},
				Privileges: []*client.SecurityPolicyRulePrivilegeContainer{
					{
						PrivilegeType: client.PasswordCheckoutRDPPrivilegeType,
						PrivilegeValue: &client.PasswordCheckoutRDPPrivilege{
							Enabled: utils.AsBoolPtr(true),
						},
					},
					{
						PrivilegeType: client.PasswordCheckoutSSHPrivilegeType,
						PrivilegeValue: &client.PasswordCheckoutSSHPrivilege{
							Enabled: utils.AsBoolPtr(true),
						},
					},
					{
						PrivilegeType: client.PrincipalAccountRDPPrivilegeType,
						PrivilegeValue: &client.PrincipalAccountRDPPrivilege{
							Enabled:               utils.AsBoolPtrZero(false, true),
							AdminLevelPermissions: utils.AsBoolPtrZero(false, true),
						},
					},
					{
						PrivilegeType: client.PrincipalAccountSSHPrivilegeType,
						PrivilegeValue: &client.PrincipalAccountSSHPrivilege{
							Enabled:               utils.AsBoolPtrZero(false, true),
							AdminLevelPermissions: utils.AsBoolPtrZero(false, true),
						},
					},
				},
				Conditions: []*client.SecurityPolicyRuleConditionContainer{
					{
						ConditionType: client.AccessRequestConditionType,
						ConditionValue: &client.AccessRequestCondition{
							RequestTypeID:       utils.AsStringPtr("abcd"),
							RequestTypeName:     utils.AsStringPtr("foo"),
							ExpiresAfterSeconds: utils.AsIntPtr(1200),
						},
					},
					{
						ConditionType: client.AccessRequestConditionType,
						ConditionValue: &client.AccessRequestCondition{
							RequestTypeID:       utils.AsStringPtr("wxyz"),
							RequestTypeName:     utils.AsStringPtr("bar"),
							ExpiresAfterSeconds: utils.AsIntPtr(1800),
						},
					},
					{
						ConditionType: client.GatewayConditionType,
						ConditionValue: &client.GatewayCondition{
							TrafficForwarding: utils.AsBoolPtr(true),
							SessionRecording:  utils.AsBoolPtr(true),
						},
					},
				},
			},
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccSecurityPolicyCheckDestroy(securityPolicyName),
		Steps: []resource.TestStep{
			{
				// Ensure that we get an error when we try to create a policy with invalid config.
				Config:      createTestAccSecurityPolicyInvalidConfig(group1Name, securityPolicyName),
				ExpectError: regexp.MustCompile(fmt.Sprintf("admin_level_permissions can not be enabled when principal account ssh privilege is not enabled")),
			},
			{
				Config: createTestAccSecurityPolicyCreateConfig(group1Name, group2Name, securityPolicyName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccSecurityPolicyCheckExists(resourceName, initialSecurityPolicy),
					resource.TestCheckResourceAttr(
						resourceName, attributes.Name, securityPolicyName,
					),
				),
			},
			{
				Config: createTestAccSecurityPolicyUpdateConfig(group1Name, securityPolicyName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccSecurityPolicyCheckExists(resourceName, updatedSecurityPolicy),
					resource.TestCheckResourceAttr(
						resourceName, attributes.Name, securityPolicyName,
					),
				),
			},
			{
				// Ensure that we get an error when we try to update an existing policy with invalid config.
				Config:      createTestAccSecurityPolicyInvalidConfig(group1Name, securityPolicyName),
				ExpectError: regexp.MustCompile(fmt.Sprintf("admin_level_permissions can not be enabled when principal account ssh privilege is not enabled")),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccSecurityPolicyCheckExists(rn string, expectedSecurityPolicy *client.SecurityPolicy) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[rn]
		if !ok {
			return fmt.Errorf("resource not found: %s", rn)
		}

		id := rs.Primary.ID
		pamClient := getLocalClientFromMetadata(testAccProvider.Meta())
		securityPolicy, err := pamClient.GetSecurityPolicy(context.Background(), id)
		if err != nil {
			return fmt.Errorf("error getting security policy: %w", err)
		} else if securityPolicy == nil {
			return fmt.Errorf("security policy does not exist")
		}
		err = insertComputedValuesForSecurityPolicy(expectedSecurityPolicy, securityPolicy)
		if err != nil {
			return err
		}

		comparison := pretty.Compare(expectedSecurityPolicy, securityPolicy)
		if comparison != "" {
			return fmt.Errorf("expected password settings does not match returned password settings.\n%s", comparison)
		}
		return nil
	}
}

func insertComputedValuesForSecurityPolicy(expectedSecurityPolicy, actualSecurityPolicy *client.SecurityPolicy) error {
	actualSecurityPolicy.ID = expectedSecurityPolicy.ID
	if expectedSecurityPolicy.Principals != nil && actualSecurityPolicy.Principals != nil {
		if userGroups, err := subNamedObjects(expectedSecurityPolicy.Principals.UserGroups, actualSecurityPolicy.Principals.UserGroups, false); err == nil {
			actualSecurityPolicy.Principals.UserGroups = userGroups
		} else {
			return err
		}
	}

	if len(actualSecurityPolicy.Rules) != len(expectedSecurityPolicy.Rules) {
		return fmt.Errorf("invalid number of rules in security policy.  expected %d, got %d", len(expectedSecurityPolicy.Rules), len(actualSecurityPolicy.Rules))
	}

	for _, rule := range actualSecurityPolicy.Rules {
		rule.SecurityPolicyID = expectedSecurityPolicy.ID
	}
	return nil
}

func testAccSecurityPolicyCheckDestroy(securityPolicyName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := getLocalClientFromMetadata(testAccProvider.Meta())
		securityPolicies, err := client.ListSecurityPolicies(context.Background())
		if err != nil {
			return fmt.Errorf("error getting security policies: %w", err)
		}

		for _, rg := range securityPolicies {
			if *rg.Name == securityPolicyName {
				return fmt.Errorf("resource group still exists")
			}
		}

		return nil
	}
}

const testAccSecurityPolicyCreateConfigFormat = `
resource "oktapam_group" "test_security_policy_group1" {
	name = "%s"
}
resource "oktapam_group" "test_security_policy_group2" {
	name = "%s"
}
resource "oktapam_security_policy" "test_acc_security_policy" {
	name = "%s"
	description = "test description"
	active = true
	principals {
		groups = [oktapam_group.test_security_policy_group1.id, oktapam_group.test_security_policy_group2.id]
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
				admin_level_permissions = true
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

func createTestAccSecurityPolicyCreateConfig(groupName1, groupName2, securityPolicyName string) string {
	return fmt.Sprintf(testAccSecurityPolicyCreateConfigFormat, groupName1, groupName2, securityPolicyName)
}

const testAccSecurityPolicyUpdateConfigFormat = `
resource "oktapam_group" "test_security_policy_group1" {
	name = "%s"
}
resource "oktapam_security_policy" "test_acc_security_policy" {
	name = "%s"
	description = "updated description"
	active = true
	principals {
		groups = [oktapam_group.test_security_policy_group1.id]
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
				admin_level_permissions = false
			}
			principal_account_ssh {
				enabled = false
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
			gateway {
				traffic_forwarding = true
				session_recording  = true
			}
		}
	}
}
`

// When the principal accounts are not enabled, terraform should not allow the users to enable
// admin level permissions for those accounts.
const testAccSecurityPolicyInvalidConfigFormat = `
resource "oktapam_group" "test_security_policy_group1" {
	name = "%s"
}
resource "oktapam_security_policy" "test_acc_security_policy" {
	name = "%s"
	description = "updated description"
	active = true
	principals {
		groups = [oktapam_group.test_security_policy_group1.id]
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
			principal_account_ssh {
				enabled = false
				admin_level_permissions = true
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

func createTestAccSecurityPolicyUpdateConfig(group1Name, securityPolicyName string) string {
	return fmt.Sprintf(testAccSecurityPolicyUpdateConfigFormat, group1Name, securityPolicyName)
}

func createTestAccSecurityPolicyInvalidConfig(group1Name string, securityPolicyName string) string {
	return fmt.Sprintf(testAccSecurityPolicyInvalidConfigFormat, group1Name, securityPolicyName)
}
