package oktapam

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

// language=Terraform
const devEnvSecurityPolicyTerraform = `resource "oktapam_security_policy_v2" "devenv_security_policy" {
  type        = "default"
  name        = "development environment policy"
  description = "An example security policy for dev environment"
  active      = true
  principals = {
    user_groups = ["user_group_id_1", "user_group_id_2"]
  }
  # rule with vaulted account and user level access
  rules = [
    {
      name          = "linux server account and user level access"
      resource_type = "server_based_resource"
      resource_selector = {
        server_based_resource = {
          selectors = [
            {
              server_label = {
                account_selector = ["root", "pamadmin"]
                server_selector = {
                  labels = {
                    "system.os_type" = "linux"
                  }
                }
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
        }, {
          principal_account_ssh = {
            principal_account_ssh   = true
            admin_level_permissions = false
          }
        }
      ]
    },

    #rule with ssh privilege and sudo access
    {
      name          = "linux server with sudo"
      resource_type = "server_based_resource"
      resource_selector = {
        server_based_resource = {
          selectors = [
            {
              server_label = {
                server_selector = {
                  labels = {
                    "system.os_type" = "linux"
                  }
                }
              }
            }
          ]
        }
      }

      privileges = [
        {
          principal_account_ssh = {
            principal_account_ssh = true
            sudo_display_name     = "sudo-display-name for end user"
            sudo_command_bundles = [
              oktapam_sudo_command_bundle.tilt_sudo_create_directories.id,
              oktapam_sudo_command_bundle.tilt_sudo_remove_directories.id
            ]
          }
        }
      ]
    },
    # rule with ssh privilege and admin + mfa every 1hr
    {
      name          = "linux server account and admin level access"
      resource_type = "server_based_resource"
      resource_selector = {
        server_based_resource = {
          selectors = [
            {
              server_label = {
                server_selector = {
                  labels = {
                    "system.os_type" = "linux"
                  }
                }
                account_selector = ["root", "pamadmin"]
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
        }, {
          principal_account_ssh = {
            principal_account_ssh   = true
            admin_level_permissions = true
          }
        }
      ]

      conditions = [
        {
          mfa = {
            acr_values                  = "urn:okta:loa:2fa:any"
            reauth_frequency_in_seconds = 3600
          }
        }
      ]
    }
  ]
}
`

const sudoCommandBundlesTerraform = `# Create a sudo command bundle
resource "oktapam_sudo_command_bundle" "tilt_sudo_create_directories" {
  name = "create_directories"
  structured_commands {
    command      = "/bin/mkdir"
    command_type = "executable"
    args_type    = "any"
  }
  no_passwd = true
}
resource "oktapam_sudo_command_bundle" "tilt_sudo_remove_directories" {
  name = "remove_directories"
  structured_commands {
    command      = "/bin/rmdir"
    command_type = "executable"
    args_type    = "any"
  }
  no_passwd = true
  add_env = ["HOME"]
}
`

const devEnvTerraformConfig = sudoCommandBundlesTerraform + "\n" + devEnvSecurityPolicyTerraform

func entityId(body []byte) string {
	sum := sha256.Sum256(body)
	return fmt.Sprintf("%x", sum)
}

func setupHTTPMock(t *testing.T) {
	var entities = make(map[string]any)

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

	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: httpMockTestV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				PreConfig: func() { setupHTTPMock(t) },
				Config:    devEnvTerraformConfig,
				Check: func(s *terraform.State) error {
					return nil
				},
			},
		},
	})
}

// TestSecurityPolicyLoopback_IndividualServer does a test on a resource with an "individual server" resource
// selector.
func TestSecurityPolicyLoopback_IndividualServer(t *testing.T) {
	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: httpMockTestV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				PreConfig: func() { setupHTTPMock(t) },
				Config:    individualServerSelectorPolicyTerraform,
				Check: func(s *terraform.State) error {
					return nil
				},
			},
		},
	})
}

// TestSecurityPolicyLoopback_IndividualServerAccount does a loopback test on a resource with an "individual server
// account" resource selector.
func TestSecurityPolicyLoopback_IndividualServerAccount(t *testing.T) {
	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: httpMockTestV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				PreConfig: func() { setupHTTPMock(t) },
				Config:    individualServerAccountSelectorPolicyTerraform,
				Check: func(s *terraform.State) error {
					return nil
				},
			},
		},
	})
}
