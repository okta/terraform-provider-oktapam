package oktapam

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/kylelemons/godebug/pretty"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"
)

func getValidServerID() string {
	if idFromEnv := os.Getenv("TF_ACC_VALID_SERVER_ID"); idFromEnv != "" {
		return idFromEnv
	}
	return "7d7f2456-e670-4d80-9e13-0d2f7557aaea"
}

func TestAccSecurityPolicy(t *testing.T) {
	checkTeamApplicable(t, true)
	resourceName := "oktapam_security_policy.test_acc_security_policy"
	randIdentifier := randSeq()
	secretsResourceName := "oktapam_security_policy.test_acc_secrets_security_policy"
	securityPolicyName := fmt.Sprintf("test_acc_security_policy_%s", randSeq())
	secretsSecurityPolicyName := fmt.Sprintf("test_acc_secrets_security_policy_%s", randSeq())
	secretsSecurityPolicyName := fmt.Sprintf("test_acc_secrets_security_policy_%s", randSeq())
	group1Name := fmt.Sprintf("test_acc_security_policy_group1_%s", randSeq())
	group2Name := fmt.Sprintf("test_acc_security_policy_group2_%s", randSeq())
	validServerID := getValidServerID()

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
									Id: utils.AsStringPtr(validServerID),
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
						PrivilegeType: client.PasswordCheckoutSSHPrivilegeType,
						PrivilegeValue: &client.PasswordCheckoutSSHPrivilege{
							Enabled: utils.AsBoolPtr(true),
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
									Id: utils.AsStringPtr(validServerID),
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
						PrivilegeType: client.PrincipalAccountRDPPrivilegeType,
						PrivilegeValue: &client.PrincipalAccountRDPPrivilege{
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

	secretsSecurityPolicy := &client.SecurityPolicy{
		Name:        &secretsSecurityPolicyName,
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
				ResourceType: client.SecretBasedResourceSelectorType,
				ResourceSelector: &client.SecretBasedResourceSelector{
					Selectors: []client.SecretBasedResourceSubSelectorContainer{
						{
							SelectorType: client.SecretFolderSubSelectorType,
							Selector: &client.SecretFolderSubSelector{
								SecretFolderID: client.NamedObject{
									Type: client.SecretFolderNamedObjectType,
								},
							},
						},
					},
				},
				Privileges: []*client.SecurityPolicyRulePrivilegeContainer{
					{
						PrivilegeType: client.SecretPrivilegeType,
						PrivilegeValue: &client.SecretPrivilege{
							List:         utils.AsBoolPtr(true),
							FolderCreate: utils.AsBoolPtr(true),
							FolderUpdate: utils.AsBoolPtr(true),
							FolderDelete: utils.AsBoolPtr(true),
							SecretCreate: utils.AsBoolPtr(true),
							SecretUpdate: utils.AsBoolPtr(true),
							SecretReveal: utils.AsBoolPtr(true),
							SecretDelete: utils.AsBoolPtr(true),
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

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccSecurityPolicyCheckDestroy(securityPolicyName),
		Steps: []resource.TestStep{
			{
				// Ensure that we get an error when we try to create a policy with invalid config.
				Config:      createTestAccSecurityPolicyInvalidAdminPrivilegesConfig(group1Name, securityPolicyName, validServerID),
				ExpectError: regexp.MustCompile("admin_level_permissions can not be enabled when principal account rdp privilege is not enabled"),
			},
			{
				// Ensure that we get an error when we try to create a policy with invalid config.
				Config:      createTestAccSecurityPolicyInvalidRDPAndSSHConfig(group1Name, securityPolicyName, validServerID),
				ExpectError: regexp.MustCompile("cannot mix SSH and RDP privileges within a Security Policy Rule"),
			},
			{
				Config: createTestAccSecurityPolicyCreateConfig(group1Name, group2Name, securityPolicyName, validServerID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccSecurityPolicyCheckExists(resourceName, initialSecurityPolicy),
					resource.TestCheckResourceAttr(
						resourceName, attributes.Name, securityPolicyName,
					),
				),
			},
			{
				Config: createTestAccSecurityPolicyUpdateConfig(group1Name, securityPolicyName, validServerID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccSecurityPolicyCheckExists(resourceName, updatedSecurityPolicy),
					resource.TestCheckResourceAttr(
						resourceName, attributes.Name, securityPolicyName,
					),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: createTestAccSecurityPolicySecretsCreateConfig(randIdentifier, group1Name, group2Name, secretsSecurityPolicyName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccSecurityPolicyCheckExists(secretsResourceName, secretsSecurityPolicy),
					resource.TestCheckResourceAttr(
						secretsResourceName, attributes.Name, secretsSecurityPolicyName,
					),
				),
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
			return fmt.Errorf("expected security policy does not match returned security policy.\n%s", comparison)
			return fmt.Errorf("expected security policy does not match returned security policy.\n%s", comparison)
		}
		return nil
	}
}

func insertComputedValuesForSecurityPolicy(expectedSecurityPolicy, actualSecurityPolicy *client.SecurityPolicy) error {
	expectedSecurityPolicy.ID = actualSecurityPolicy.ID
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

	expectedRulesByName := make(map[string]*client.SecurityPolicyRule, len(expectedSecurityPolicy.Rules))

	for _, rule := range expectedSecurityPolicy.Rules {
		rule.SecurityPolicyID = actualSecurityPolicy.ID

		expectedRulesByName[*rule.Name] = rule
	}

	for _, rule := range actualSecurityPolicy.Rules {
		rule := rule
		if expectedRule, ok := expectedRulesByName[*rule.Name]; ok {
			if err := insertComputedValuesForSecurityPolicyRule(expectedRule, rule); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("found unexpected rule with name: %s", *rule.Name)
		}
	}

	return nil
}

func insertComputedValuesForSecurityPolicyRule(expectedRule *client.SecurityPolicyRule, actualRule *client.SecurityPolicyRule) error {
	expectedRule.ID = actualRule.ID
	if expectedRule.ResourceType != actualRule.ResourceType {
		return fmt.Errorf("resource types did not match between expected and actual for rule %s", *expectedRule.Name)
	}

	switch expectedRule.ResourceType {
	case client.ServerBasedResourceSelectorType:
		expectedResourceSelector := expectedRule.ResourceSelector.(*client.ServerBasedResourceSelector)
		actualResourceSelector := actualRule.ResourceSelector.(*client.ServerBasedResourceSelector)

		if len(expectedResourceSelector.Selectors) != len(actualResourceSelector.Selectors) {
			return fmt.Errorf("number of selectors did not match between expected and actual for rule %s", *expectedRule.Name)
		}

		for idx, expectedSel := range expectedResourceSelector.Selectors {
			actualSel := actualResourceSelector.Selectors[idx]

			if expectedSel.SelectorType != actualSel.SelectorType {
				return fmt.Errorf("server based selector type did not match within rule %s", *expectedRule.Name)
			}

			if expectedSel.SelectorType == client.IndividualServerAccountSubSelectorType {
				expectedSubSel := expectedSel.Selector.(*client.IndividualServerAccountSubSelector)
				actualSubSel := actualSel.Selector.(*client.IndividualServerAccountSubSelector)

				expectedSubSel.Server = fillNamedObjectValues(expectedSubSel.Server, actualSubSel.Server)
			} else if expectedSel.SelectorType == client.IndividualServerSubSelectorType {
				expectedSubSel := expectedSel.Selector.(*client.IndividualServerSubSelector)
				actualSubSel := actualSel.Selector.(*client.IndividualServerSubSelector)

				expectedSubSel.Server = fillNamedObjectValues(expectedSubSel.Server, actualSubSel.Server)
			}
		}
	case client.SecretBasedResourceSelectorType:
		expectedResourceSelector := expectedRule.ResourceSelector.(*client.SecretBasedResourceSelector)
		actualResourceSelector := actualRule.ResourceSelector.(*client.SecretBasedResourceSelector)

		if len(expectedResourceSelector.Selectors) != len(actualResourceSelector.Selectors) {
			return fmt.Errorf("number of selectors did not match between expected and actual for rule %s", *expectedRule.Name)
		}

		for idx, expectedSel := range expectedResourceSelector.Selectors {
			actualSel := actualResourceSelector.Selectors[idx]

			if expectedSel.SelectorType != actualSel.SelectorType {
				return fmt.Errorf("server based selector type did not match within rule %s", *expectedRule.Name)
			}

			if expectedSel.SelectorType == client.SecretFolderSubSelectorType {
				expectedSubSel := expectedSel.Selector.(*client.SecretFolderSubSelector)
				actualSubSel := actualSel.Selector.(*client.SecretFolderSubSelector)

				expectedSubSel.SecretFolderID = fillNamedObjectValues(expectedSubSel.SecretFolderID, actualSubSel.SecretFolderID)
			} else if expectedSel.SelectorType == client.SecretSubSelectorType {
				expectedSubSel := expectedSel.Selector.(*client.SecretSubSelector)
				actualSubSel := actualSel.Selector.(*client.SecretSubSelector)

				expectedSubSel.SecretID = fillNamedObjectValues(expectedSubSel.SecretID, actualSubSel.SecretID)
			}
		}

	default:
		return fmt.Errorf("cannot insert computed values for rule for selector type of: %s", string(expectedRule.ResourceType))
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
					server_id = "%s"
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
			password_checkout_ssh {
				enabled = true
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

func createTestAccSecurityPolicyCreateConfig(groupName1, groupName2, securityPolicyName string, serverID string) string {
	return fmt.Sprintf(testAccSecurityPolicyCreateConfigFormat, groupName1, groupName2, securityPolicyName, serverID)
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
					server_id = "%s"
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
			principal_account_rdp {
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

const testAccSecurityPolicyInvalidRDPAndSSHConfigFormat = `
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
					server_id = "%s"
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
			principal_account_rdp {
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

const testAccSecurityPolicyInvalidRDPAndSSHConfigFormat = `
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
					server_id = "%s"
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
			principal_account_rdp {
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

const testAccSecurityPolicyInvalidRDPAndSSHConfigFormat = `
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
					server_id = "%s"
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
const testAccSecurityPolicyInvalidAdminPrivilegesConfigFormat = `
const testAccSecurityPolicyInvalidAdminPrivilegesConfigFormat = `
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
					server_id = "%s"
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
			principal_account_rdp {
			principal_account_rdp {
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

func createTestAccSecurityPolicySecretsCreateConfig(identifier, groupName1, groupName2, securityPolicyName string) string {
	return fmt.Sprintf(testAccSecurityPolicyCreateSecretsConfigFormat, identifier, identifier, identifier, identifier, groupName1, groupName2, securityPolicyName)
}

const testAccSecurityPolicyCreateSecretsConfigFormat = `
resource "oktapam_group" "test_resource_group_dga_group" {
	name = "sp-test-dra-group-%s"
}
resource "oktapam_resource_group" "test_acc_resource_group" {
	name = "sp-test-rg-%s"
	description = "test resource group"
	delegated_resource_admin_groups = [oktapam_group.test_resource_group_dga_group.id]	
}
resource "oktapam_resource_group_project" "test_acc_resource_group_project" {
	name = "sp-test-rg-project-%s"
	resource_group = oktapam_resource_group.test_acc_resource_group.id
}
resource "oktapam_secret_folder" "test_acc_secret_folder_top_level" {
	name = "sp-test-secret-folder-%s"
	description = "updated top-level folder for test"
	resource_group = oktapam_resource_group.test_acc_resource_group.id
	project = oktapam_resource_group_project.test_acc_resource_group_project.id
}
resource "oktapam_group" "test_security_policy_group1" {
	name = "%s"
}
resource "oktapam_group" "test_security_policy_group2" {
	name = "%s"
}
resource "oktapam_security_policy" "test_acc_secrets_security_policy" {
	name = "%s"
	description = "test description"
	active = true
	principals {
		groups = [oktapam_group.test_security_policy_group1.id, oktapam_group.test_security_policy_group2.id]
	}
	rule {
		name = "first rule"
		resources {
			secrets {
				secret_folder {
					secret_folder_id = oktapam_secret_folder.test_acc_secret_folder_top_level.id
			  	}
			}
		}
		privileges {
			secret {
				list = true
			  	folder_create = true
			  	folder_delete = true
			  	folder_update = true
			  	secret_create = true
			  	secret_delete = true
			  	secret_reveal = true
			  	secret_update = true
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

func createTestAccSecurityPolicyUpdateConfig(group1Name string, securityPolicyName string, serverID string) string {
	return fmt.Sprintf(testAccSecurityPolicyUpdateConfigFormat, group1Name, securityPolicyName, serverID)
}

func createTestAccSecurityPolicyInvalidAdminPrivilegesConfig(group1Name string, securityPolicyName string, serverID string) string {
	return fmt.Sprintf(testAccSecurityPolicyInvalidAdminPrivilegesConfigFormat, group1Name, securityPolicyName, serverID)
}

func createTestAccSecurityPolicyInvalidRDPAndSSHConfig(group1Name string, securityPolicyName string, serverID string) string {
	return fmt.Sprintf(testAccSecurityPolicyInvalidRDPAndSSHConfigFormat, group1Name, securityPolicyName, serverID)
}
