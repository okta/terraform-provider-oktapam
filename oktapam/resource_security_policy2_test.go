package oktapam

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"io"
	"net/http"
	"os"
	"regexp"
	"testing"

	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

const individualServerAccountSelectorPolicyTerraform = `resource "oktapam_security_policy_v2" "individual_server_account_policy" {
  name   = "test individual server account selector"
  active = true
  principals = { user_groups = ["user_group_1", "user_group_2"] }
  rules = [
    {
      name          = "test"
      resource_type = "server_based_resource"
      resource_selector = {
        server_based_resource = {
          selectors = [
            {
              individual_server_account = {
                server   = "server-id-goes-here"
                username = "root"
              }
            }
          ]
        }
      }
      privileges = [{ password_checkout_ssh = { password_checkout_ssh = true } }]
    }
  ]
}
`

const individualServerSelectorPolicyTerraform = `resource "oktapam_security_policy_v2" "individual_server_policy" {
  name = "test individual server selector"
  active = true
  principals = { user_groups = ["user_group_1", "user_group_2"] }
  rules = [
    {
      name          = "test"
      resource_type = "server_based_resource"
      resource_selector = {
        server_based_resource = {
          selectors = [
            { individual_server = { server = "server-id-goes-here" } }
          ]
        }
      }
      privileges = [{ password_checkout_ssh = { password_checkout_ssh = true } }]
    }
  ]
}


`

func entityId(body []byte) string {
	sum := sha256.Sum256(body)
	return fmt.Sprintf("%x", sum)
}

func setupHTTPMock(t *testing.T, entities map[string]any) {

	prefix := "/v1/teams/httpmock-test-team"

	// This is a quick hack that unmarshals a given entity, assigns it an ID then stores it in a map. It then marshals
	// the ID-enhanced entity as a response.
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
			entities[id] = created
			return httpmock.NewJsonResponse(http.StatusCreated, created)
		})

	// This is the cousin of the quick hack above that looks up the entity in the map and marshals whatever is there as
	// a response.
	httpmock.RegisterRegexpResponder(http.MethodGet, regexp.MustCompile(prefix+`/(sudo_command_bundles|security_policy)/(.*)`),
		func(request *http.Request) (*http.Response, error) {
			id := httpmock.MustGetSubmatch(request, 2)
			return httpmock.NewJsonResponse(http.StatusOK, entities[id])
		})

	httpmock.RegisterRegexpResponder(http.MethodDelete, regexp.MustCompile(`.*`),
		httpmock.NewStringResponder(http.StatusOK, ""),
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

	for _, entity := range entities {
		switch entity.(type) {
		case pam.SecurityPolicy:
			expectedJSONBytes, err := os.ReadFile("testdata/dev_env.json")
			require.NoError(t, err)
			actualJSONBytes, err := json.Marshal(entity)
			require.NoError(t, err)
			require.JSONEq(t, string(expectedJSONBytes), string(actualJSONBytes))
		}
	}
	require.Len(t, entities, 3)
}

// TestSecurityPolicyLoopback_IndividualServer does a test on a resource with an "individual server" resource
// selector.
func TestSecurityPolicyLoopback_IndividualServer(t *testing.T) {
	var entities = make(map[string]any)

	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: httpMockTestV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				PreConfig: func() { setupHTTPMock(t, entities) },
				Config:    individualServerSelectorPolicyTerraform,
				Check: func(s *terraform.State) error {
					return nil
				},
			},
		},
	})
	require.Len(t, entities, 1)
	for _, entity := range entities {
		policy := entity.(pam.SecurityPolicy)
		require.Len(t, policy.Rules, 1)
		require.Len(t, policy.Rules[0].Privileges, 1)
		require.Equal(t, "password_checkout_ssh", string(*policy.Rules[0].Privileges[0].PrivilegeType))
	}
}

// TestSecurityPolicyLoopback_IndividualServerAccount does a loopback test on a resource with an "individual server
// account" resource selector.
func TestSecurityPolicyLoopback_IndividualServerAccount(t *testing.T) {
	var entities = make(map[string]any)

	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: httpMockTestV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				PreConfig: func() { setupHTTPMock(t, entities) },
				Config:    individualServerAccountSelectorPolicyTerraform,
				Check: func(s *terraform.State) error {
					return nil
				},
			},
		},
	})
}
