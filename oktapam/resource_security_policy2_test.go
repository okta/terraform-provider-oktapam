package oktapam

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"sync"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"

	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

func entityId(body []byte) string {
	sum := sha256.Sum256(body)
	return fmt.Sprintf("%x", sum)
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
			entitiesLock.Lock()
			defer entitiesLock.Unlock()
			return httpmock.NewJsonResponse(http.StatusOK, entities[id])
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
