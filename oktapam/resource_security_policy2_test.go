package oktapam

import (
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"testing"
)

// language=Terraform
const securityPolicyTerraform = `resource "oktapam_security_policy_v2" "tilt_security_policy" {
  name   = "tilt-security-policy"
  active = true
  principals = {
    groups = ["user_group_id_1", "user_group_id_2"]
  }
  # rule with vaulted account and user level access
  rules = [
    {
      name = "linux server account and user level access"
      resources = {
        servers = {
          server_label = {
            accounts = ["root", "pamadmin"]
            server_selector = {
              labels = {
                "system.os_type" = "linux"
              }
            }
          }
        }
      }


      privileges = {
        password_checkout_ssh = {
          password_checkout_ssh = true
        }
        principal_account_ssh = {
          principal_account_ssh   = true
          admin_level_permissions = false
        }
      }
    },

    #rule with ssh privilege and sudo access
    {
      name = "linux server with sudo"
      resources = {
        servers = {
          label_selectors = {
            server_labels = {
              "system.os_type" = "linux"
            }
          }
        }
      }

      privileges = {
        principal_account_ssh = {
          principal_account_ssh = true
          sudo_display_name     = "sudo-display-name for end user"
          sudo_command_bundles = [
            oktapam_sudo_command_bundle.tilt_sudo_create_directories.id,
            oktapam_sudo_command_bundle.tilt_sudo_remove_directories.id
          ]
        }
      }
    },
    # rule with ssh privilege and admin + mfa every 1hr
    {
      name = "linux server account and admin level access"
      resources = {
        servers = {
          label_selectors = {
            server_labels = {
              "system.os_type" = "linux"
            }
            accounts = ["root", "pamadmin"]
          }
        }
      }
      privileges = {
        password_checkout_ssh = {
          password_checkout_ssh = true
        }
        principal_account_ssh = {
          principal_account_ssh   = true
          admin_level_permissions = true
        }
      }
      conditions = {
        mfa = {
          acr_values                  = "urn:okta:loa:2fa:any"
          reauth_frequency_in_seconds = 3600
        }
      }
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

const terraformConfig = sudoCommandBundlesTerraform + "\n" + securityPolicyTerraform

func TestSimple(t *testing.T) {

	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: testAccV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: terraformConfig,
				Check: func(s *terraform.State) error {
					return nil
				},
			},
		},
	})
}
