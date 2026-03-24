package oktapam

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"sync"
	"testing"

	"github.com/atko-pam/pam-sdk-go/client/pam"
	fwresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/jarcoal/httpmock"
	"github.com/okta/terraform-provider-oktapam/oktapam/convert"
	"github.com/stretchr/testify/require"
)

func entityId(body []byte) string {
	sum := sha256.Sum256(body)
	hash := fmt.Sprintf("%x", sum)
	// Convert hash to UUID format: xxxxxxxx-xxxx-4xxx-8xxx-xxxxxxxxxxxx
	// Take first 32 characters of hash and format as UUID
	if len(hash) >= 32 {
		return fmt.Sprintf("%s-%s-4%s-8%s-%s",
			hash[0:8],
			hash[8:12],
			hash[13:16],
			hash[17:20],
			hash[20:32])
	}
	// Fallback to a fixed UUID if hash is too short (shouldn't happen)
	return "12345678-1234-4123-8123-123456789abc"
}

func setupHTTPMock(t *testing.T, entities map[string]any) {

	var entitiesLock sync.Mutex

	prefix := "/v1/teams/httpmock-test-team"

	// This is a quick hack that unmarshals a given entity, assigns it an ID then stores it in a map. It then marshals
	// the ID-enhanced entity as a response. The ID is generated from hashing the contents because the order of
	// requests from Terraform is not guaranteed.
	httpmock.RegisterRegexpResponder(http.MethodPost, regexp.MustCompile(prefix+`/(sudo_command_bundles|security_policy)`),
		func(request *http.Request) (*http.Response, error) {
			var created any

			body, err := io.ReadAll(request.Body)
			require.NoError(t, err)

			id := entityId(body)
			entityType := httpmock.MustGetSubmatch(request, 1)

			switch entityType {
			case "sudo_command_bundles":
				var sudoCommandBundle pam.SudoCommandBundle
				require.NoError(t, json.Unmarshal(body, &sudoCommandBundle))
				sudoCommandBundle.SetId(id)
				created = sudoCommandBundle
			case "security_policy":
				var securityPolicy pam.SecurityPolicy
				require.NoError(t, json.Unmarshal(body, &securityPolicy))
				securityPolicy.SetId(id)
				created = securityPolicy
			default:
				panic("don't know how to create " + entityType)
			}
			entitiesLock.Lock()
			defer entitiesLock.Unlock()
			entities[id] = created
			return httpmock.NewJsonResponse(http.StatusCreated, created)
		})

	// This is the cousin of the quick hack above that looks up the entity in the map and marshals whatever is there as
	// a response.
	httpmock.RegisterRegexpResponder(http.MethodGet, regexp.MustCompile(prefix+`/(sudo_command_bundles|security_policy)/(.*)`),
		func(request *http.Request) (*http.Response, error) {
			id := httpmock.MustGetSubmatch(request, 2)
			entityType := httpmock.MustGetSubmatch(request, 1)
			entitiesLock.Lock()
			defer entitiesLock.Unlock()

			// First, try to find the entity by the exact ID
			if entity, exists := entities[id]; exists {
				return httpmock.NewJsonResponse(http.StatusOK, entity)
			}

			// If not found by exact ID, this might be an import test with a UUID.
			// For import tests, we need to return a mock entity with the requested ID
			// instead of the hash-based ID that was used during creation.
			switch entityType {
			case "security_policy":
				// Return a mock security policy for import testing
				// We'll take the first security policy from entities and clone it with the requested ID
				for _, entity := range entities {
					if securityPolicy, ok := entity.(pam.SecurityPolicy); ok {
						// Create a proper deep clone by marshaling and unmarshaling
						jsonData, err := json.Marshal(securityPolicy)
						if err != nil {
							return httpmock.NewStringResponse(http.StatusInternalServerError, "Marshal error"), err
						}

						var clonedPolicy pam.SecurityPolicy
						err = json.Unmarshal(jsonData, &clonedPolicy)
						if err != nil {
							return httpmock.NewStringResponse(http.StatusInternalServerError, "Unmarshal error"), err
						}

						// Set the requested ID
						clonedPolicy.SetId(id)
						return httpmock.NewJsonResponse(http.StatusOK, clonedPolicy)
					}
				}
				// Return 404 if no security policy exists to clone
				return httpmock.NewStringResponse(http.StatusNotFound, "Entity not found"), nil
			case "sudo_command_bundles":
				// Handle sudo command bundles if needed
				for _, entity := range entities {
					if sudoBundle, ok := entity.(pam.SudoCommandBundle); ok {
						clonedBundle := sudoBundle
						clonedBundle.SetId(id)
						return httpmock.NewJsonResponse(http.StatusOK, clonedBundle)
					}
				}
			}

			// If entity not found, return 404
			return httpmock.NewStringResponse(http.StatusNotFound, "Entity not found"), nil
		})

	// This doesn't actually delete anything, if you need it, feel free to add it in.
	httpmock.RegisterRegexpResponder(http.MethodDelete, regexp.MustCompile(`.*`),
		httpmock.NewStringResponder(http.StatusNoContent, ""),
	)
}

// Loopback tests are designed to convert a Terraform resource into the SDK representation, marshal it, pretend
// to call an HTTP call which will unmarshal it, assign it an ID, then marshal a response and send that back. The
// Terraform provider will in turn unmarshal from JSON to SDK, then convert from SDK to Terraform model. At its
// root, loopback tests are Garbage In/Garbage Out. For example, if you don't take a resource field and set it in
// the SDK, _and_ you're not converting it from the SDK back to the Terraform resource, then nobody will notice. You
// can improve that by including a canonical expectation of what the JSON should look like and require.JSONEq checking.

// TestSecurityPolicyLoopback_DevEnv uses an example policy from our development environment.
func TestSecurityPolicyLoopback_DevEnv(t *testing.T) {
	var entities = make(map[string]any)

	setupHTTPMock(t, entities)

	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: httpMockTestV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				ConfigFile: config.StaticFile("testdata/dev_env.tf"),
			},
			{
				ConfigFile: config.StaticFile("testdata/dev_env.tf"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				}},
		},
	})
	checkSecurityPolicyJSON(t, entities, "testdata/dev_env.json")
}

// TestSecurityPolicyImport_DevEnv tests importing an existing security policy using the dev_env configuration
func TestSecurityPolicyImport_DevEnv(t *testing.T) {
	var entities = make(map[string]any)

	setupHTTPMock(t, entities)

	// First create a security policy to import
	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: httpMockTestV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				ConfigFile: config.StaticFile("testdata/dev_env.tf"),
			},
			{
				ResourceName:      "oktapam_security_policy_v2.devenv_security_policy",
				ImportState:       true, // runs `terraform import` using the state's ID by default
				ImportStateVerify: true, // refreshes and verify attributes match after import
			},
		},
	})
}

// TestSecurityPolicyImport_IndividualServer tests importing a security policy with individual server selector
func TestSecurityPolicyImport_IndividualServer(t *testing.T) {
	var entities = make(map[string]any)

	setupHTTPMock(t, entities)

	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: httpMockTestV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				ConfigFile: config.StaticFile("testdata/individual_server_selector.tf"),
			},
			{
				ResourceName:      "oktapam_security_policy_v2.individual_server_policy",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// TestSecurityPolicyImport_InvalidID tests importing with an invalid (non-UUID) ID
func TestSecurityPolicyImport_InvalidID(t *testing.T) {
	var entities = make(map[string]any)

	setupHTTPMock(t, entities)

	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: httpMockTestV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				ConfigFile: config.StaticFile("testdata/minimal_security_policy.tf"),
			},
			{
				ResourceName:  "oktapam_security_policy_v2.test",
				ImportState:   true,
				ImportStateId: "invalid-id", // explicitly set to an invalid ID instead of using the state's ID
				ExpectError:   regexp.MustCompile("Security policy import requires a valid UUID"),
			},
		},
	})
}

// TestSecurityPolicyLoopback_IndividualServer does a test on a resource with an "individual server" resource
// selector.
func TestSecurityPolicyLoopback_IndividualServer(t *testing.T) {
	var entities = make(map[string]any)

	setupHTTPMock(t, entities)
	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: httpMockTestV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				ConfigFile: config.StaticFile("testdata/individual_server_selector.tf"),
			},
			{
				ConfigFile: config.StaticFile("testdata/individual_server_selector.tf"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
	checkSecurityPolicyJSON(t, entities, "testdata/individual_server_selector.json")
}

// TestSecurityPolicyLoopback_IndividualServerAccount does a loopback test on a resource with an "individual server
// account" resource selector.
func TestSecurityPolicyLoopback_IndividualServerAccount(t *testing.T) {
	var entities = make(map[string]any)

	setupHTTPMock(t, entities)
	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: httpMockTestV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				ConfigFile: config.StaticFile("testdata/individual_server_account_selector.tf"),
			},
			{
				ConfigFile: config.StaticFile("testdata/individual_server_account_selector.tf"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})

	checkSecurityPolicyJSON(t, entities, "testdata/individual_server_account_selector.json")
}

// TestSecurityPolicyLoopback_InvalidPrivileges1 ensures that at least one privilege must be put in the privilege
// container - in this case we have an empty stanza for a privilege container, which errors out.
func TestSecurityPolicyLoopback_InvalidPrivileges1(t *testing.T) {
	var entities = make(map[string]any)

	setupHTTPMock(t, entities)
	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: httpMockTestV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				ConfigFile:  config.StaticFile("testdata/invalid_privileges_1.tf"),
				ExpectError: regexp.MustCompile(".*privilege listed in policy rule.*"),
			},
		},
	})
}

// TestSecurityPolicyLoopback_InvalidConditions1 similar to privileges, an empty stanza as a condition in the
// conditions container is invalid in a policy rule. This then errors out.
func TestSecurityPolicyLoopback_InvalidConditions1(t *testing.T) {
	var entities = make(map[string]any)

	setupHTTPMock(t, entities)
	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: httpMockTestV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				ConfigFile:  config.StaticFile("testdata/invalid_conditions_1.tf"),
				ExpectError: regexp.MustCompile(".*condition listed in policy rule.*"),
			},
		},
	})
}

// TestSecurityPolicyLoopback_InvalidServerBasedSubSelectors1 similar to conditions and privileges, an empty stanza
// as a "server based sub selector" is invalid. This test ensures it errors out if one is found.
func TestSecurityPolicyLoopback_InvalidServerBasedSubSelectors1(t *testing.T) {
	var entities = make(map[string]any)

	setupHTTPMock(t, entities)
	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: httpMockTestV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				ConfigFile:  config.StaticFile("testdata/invalid_server_based_sub_selector_1.tf"),
				ExpectError: regexp.MustCompile(".*selector listed in policy rule.*"),
			},
		},
	})
}

// TestSecurityPolicyRevealPasswordPrivilege tests creating a security policy with reveal password privilege
func TestSecurityPolicyRevealPasswordPrivilege(t *testing.T) {
	var entities = make(map[string]any)
	setupHTTPMock(t, entities)
	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: httpMockTestV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				ConfigFile: config.StaticFile("testdata/reveal_password_privilege.tf"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("oktapam_security_policy_v2.reveal_password_test", "name", "reveal-password-test-policy"),
					resource.TestCheckResourceAttr("oktapam_security_policy_v2.reveal_password_test", "description", "Test policy for reveal password privilege"),
					resource.TestCheckResourceAttr("oktapam_security_policy_v2.reveal_password_test", "active", "true"),
					resource.TestCheckResourceAttr("oktapam_security_policy_v2.reveal_password_test", "rules.0.name", "Reveal password access rule"),
					resource.TestCheckResourceAttr("oktapam_security_policy_v2.reveal_password_test", "rules.0.privileges.0.reveal_password.reveal_password", "true"),
				),
			},
		},
	})
	checkSecurityPolicyJSON(t, entities, "testdata/reveal_password_privilege.json")
}

// TestSecurityPolicyImportRevealPasswordPrivilege tests importing a security policy with reveal password privilege
func TestSecurityPolicyImportRevealPasswordPrivilege(t *testing.T) {
	var entities = make(map[string]any)
	setupHTTPMock(t, entities)
	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: httpMockTestV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				ConfigFile: config.StaticFile("testdata/reveal_password_privilege.tf"),
			},
			{
				ResourceName:      "oktapam_security_policy_v2.reveal_password_test",
				ImportState:       true,
				ImportStateId:     "22b69692-dc1b-48c5-89b9-5b90b69f81c2", // Valid UUID format
				ImportStateVerify: true,
			},
		},
	})
}

// TestSecurityPolicyLoopback_ResourceGroup tests creating a security policy with a resource_group set.
func TestSecurityPolicyLoopback_ResourceGroup(t *testing.T) {
	var entities = make(map[string]any)

	setupHTTPMock(t, entities)
	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: httpMockTestV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				ConfigFile: config.StaticFile("testdata/resource_group_policy.tf"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("oktapam_security_policy_v2.resource_group_policy", "name", "resource group test policy"),
					resource.TestCheckResourceAttr("oktapam_security_policy_v2.resource_group_policy", "resource_group", "a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d"),
				),
			},
			{
				ConfigFile: config.StaticFile("testdata/resource_group_policy.tf"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
	checkSecurityPolicyJSON(t, entities, "testdata/resource_group_policy.json")
}

// TestSecurityPolicyImport_ResourceGroup tests importing a security policy with a resource_group set.
func TestSecurityPolicyImport_ResourceGroup(t *testing.T) {
	var entities = make(map[string]any)

	setupHTTPMock(t, entities)
	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: httpMockTestV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				ConfigFile: config.StaticFile("testdata/resource_group_policy.tf"),
			},
			{
				ResourceName:      "oktapam_security_policy_v2.resource_group_policy",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// checkSecurityPolicyJSON digs through the entities to find the first pam.SecurityPolicy and ensures
// its contents match the contents of the specified file. This is a bit brittle, do expect things to break
// if you change the .tf file.
func checkSecurityPolicyJSON(t *testing.T, entities map[string]any, jsonFilename string) {
	jsonChecked := false
	for _, entity := range entities {
		switch entity.(type) {
		case pam.SecurityPolicy:
			expectedJSONBytes, err := os.ReadFile(jsonFilename)
			require.NoError(t, err)
			actualJSONBytes, err := json.Marshal(entity)
			require.NoError(t, err)
			require.JSONEq(t, string(expectedJSONBytes), string(actualJSONBytes), "security policy json must match, check that you haven't broken anything")
			jsonChecked = true
		}
	}
	require.True(t, jsonChecked, "security policy json must be checked")
}

func TestSecurityPolicyV2IntegrationTest(t *testing.T) {
	checkTeamApplicable(t, true)
	securityPolicyName := fmt.Sprintf("test_acc_security_policy_%s", randSeq())
	groupName := fmt.Sprintf("test_acc_security_policy_group_%s", randSeq())
	validServerID := getValidServerID()
	policyResourceName := "oktapam_security_policy_v2.test_acc_security_policy_v2"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: createTestAccSecurityPolicyV2CreateConfig(groupName, securityPolicyName, validServerID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(policyResourceName, "name", securityPolicyName),
					resource.TestCheckResourceAttr(policyResourceName, "description", "test description"),
					resource.TestCheckResourceAttr(policyResourceName, "active", "true"),
					resource.TestCheckResourceAttrPair(policyResourceName, "principals.user_groups.0", "oktapam_group.test_security_policy_group", "id"),
					resource.TestCheckResourceAttr(policyResourceName, "rules.0.name", "linux server account and user level access"),
					resource.TestCheckResourceAttr(policyResourceName, "rules.0.resource_type", "server_based_resource"),
					resource.TestCheckResourceAttr(policyResourceName, "rules.0.resource_selector.server_based_resource.selectors.0.server_label.account_selector_type", "username"),
					resource.TestCheckResourceAttr(policyResourceName, "rules.0.resource_selector.server_based_resource.selectors.0.server_label.account_selector.usernames.0", "pamadmin"),
					resource.TestCheckResourceAttr(policyResourceName, "rules.0.resource_selector.server_based_resource.selectors.0.server_label.server_selector.labels.system.os_type", "linux"),
					resource.TestCheckResourceAttr(policyResourceName, "rules.0.resource_selector.server_based_resource.selectors.1.individual_server.server", validServerID),
					resource.TestCheckResourceAttr(policyResourceName, "rules.0.privileges.0.password_checkout_ssh.password_checkout_ssh", "true"),
					resource.TestCheckResourceAttr(policyResourceName, "rules.0.privileges.1.principal_account_ssh.principal_account_ssh", "true"),
					resource.TestCheckResourceAttr(policyResourceName, "rules.0.privileges.1.principal_account_ssh.admin_level_permissions", "false"),
				),
			},
		},
	})
}

const testAccSecurityPolicyV2CreateConfigFormat = `
resource "oktapam_group" "test_security_policy_group" {
	name = "%s"
}

resource "oktapam_security_policy_v2" "test_acc_security_policy_v2" {
    type = "default"
	name = "%s"
	description = "test description"
	active = true
	principals = {
		user_groups = [oktapam_group.test_security_policy_group.id], 
	}
    rules = [
    {
      name          = "linux server account and user level access"
      resource_type = "server_based_resource"
      resource_selector = {
        server_based_resource = {
          selectors = [
            {
              server_label = {
                account_selector_type = "username"
                account_selector = {
                  usernames = ["pamadmin"]
                }
                server_selector = {
                  labels = {
                    "system.os_type" = "linux"
                  }
                }
              }
            },
            {
              individual_server = {
				server = "%s"
			  }
			}
          ]
        }
      }

      privileges = [
        {
          password_checkout_ssh = {
            password_checkout_ssh = true
          }
        },
        {
          principal_account_ssh = {
            principal_account_ssh   = true
            admin_level_permissions = false
          }
        }
      ]
    }
	]
}
`

func createTestAccSecurityPolicyV2CreateConfig(groupName, securityPolicyName string, serverID string) string {
	return fmt.Sprintf(testAccSecurityPolicyV2CreateConfigFormat, groupName, securityPolicyName, serverID)
}

// TestUpgradeSecurityPolicyState_RemovesUnknownAttributes verifies that the state upgrader correctly handles old state
// that contains attributes removed from the schema.
//
// Note: The Terraform Plugin Framework stores all schema-defined attributes in state, including optional ones that were
// never configured (as null). When an attribute is later removed from the schema, the stored null entry becomes orphaned
// and causes "unsupported attribute" errors during plan. The state upgrader must strip these.
func TestUpgradeSecurityPolicyState_RemovesUnknownAttributes(t *testing.T) {
	rawStateJSON := `{
		"id": "test-policy-id",
		"name": "test-policy",
		"type": "default",
		"description": "test description",
		"active": true,
		"resource_group": null,
		"principals": {
			"user_groups": ["group1"]
		},
		"rules": [
			{
				"name": "test rule",
				"resource_type": "server_based_resource",
				"override_checkout_duration_in_seconds": null,
				"resource_selector": {
					"server_based_resource": {
						"selectors": [
							{
								"individual_server": null,
								"individual_server_account": null,
								"server_label": {
									"server_selector": {
										"labels": {"system.os_type": "linux"}
									},
									"account_selector_type": "none",
									"account_selector": {
										"usernames": null
									}
								}
							}
						]
					}
				},
				"privileges": [
					{
						"principal_account_ssh": {
							"principal_account_ssh": true,
							"admin_level_permissions": false,
							"sudo_display_name": null,
							"sudo_command_bundles": null
						},
						"password_checkout_ssh": null,
						"reveal_password": null,
						"password_checkout_rdp": null,
						"principal_account_rdp": null,
						"password_checkout_database": null
					}
				],
				"conditions": null
			}
		]
	}`

	ctx := context.Background()

	// First, prove the old state is invalid for the current schema.
	// This is the exact error users would see without the state upgrader.
	currentSchema := schema.Schema{
		Attributes: convert.SecurityPolicySchema(),
	}
	resourceSchemaType := currentSchema.Type().TerraformType(ctx)

	_, err := tftypes.ValueFromJSON([]byte(rawStateJSON), resourceSchemaType)
	require.Error(t, err, "old state with removed attribute should fail without IgnoreUndefinedAttributes")
	require.Contains(t, err.Error(), "unsupported attribute \"password_checkout_database\"")

	// Now run the state upgrader — it should succeed
	req := fwresource.UpgradeStateRequest{
		RawState: &tfprotov6.RawState{
			JSON: []byte(rawStateJSON),
		},
	}
	resp := fwresource.UpgradeStateResponse{}

	upgradeSecurityPolicyState(ctx, req, &resp)

	require.False(t, resp.Diagnostics.HasError(), "state upgrade should not produce errors: %v", resp.Diagnostics.Errors())
	require.NotNil(t, resp.State.Raw.Type(), "upgraded state should have a valid type")

	// Navigate into the upgraded state to verify the privilege container
	var topLevel map[string]tftypes.Value
	require.NoError(t, resp.State.Raw.As(&topLevel))

	var rules []tftypes.Value
	require.NoError(t, topLevel["rules"].As(&rules))
	require.Len(t, rules, 1)

	var rule map[string]tftypes.Value
	require.NoError(t, rules[0].As(&rule))

	var privileges []tftypes.Value
	require.NoError(t, rule["privileges"].As(&privileges))
	require.Len(t, privileges, 1)

	var privilegeContainer map[string]tftypes.Value
	require.NoError(t, privileges[0].As(&privilegeContainer))

	// The removed attribute must be gone
	_, hasRemoved := privilegeContainer["password_checkout_database"]
	require.False(t, hasRemoved, "removed attribute 'password_checkout_database' should have been stripped from state")

	// Valid attributes must be preserved
	_, hasPrincipalSSH := privilegeContainer["principal_account_ssh"]
	require.True(t, hasPrincipalSSH, "principal_account_ssh should be preserved")
	require.False(t, privilegeContainer["principal_account_ssh"].IsNull(), "principal_account_ssh was set and should not be null")

	// The other known privilege types should still be present (as null, since they weren't set)
	for _, knownType := range []string{"password_checkout_ssh", "reveal_password", "password_checkout_rdp", "principal_account_rdp"} {
		val, exists := privilegeContainer[knownType]
		require.True(t, exists, "%s should still be present in state", knownType)
		require.True(t, val.IsNull(), "%s should be null since it was not set", knownType)
	}
}
