resource "oktapam_security_policy_v2" "devenv_security_policy" {
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
                account_selector_type = "username"
                account_selector = {
                  usernames = ["root", "pamadmin"]
                }
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
        { /* this is invalid*/ },
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
                account_selector_type = "none"
                account_selector = {}
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
                account_selector_type = "username"
                account_selector = {
                  usernames = ["root", "pamadmin"]
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
            admin_level_permissions = true
          }
        }
      ]

      conditions = [
        {
          mfa = {
            acr_values                   = "urn:okta:loa:2fa:any"
            re_auth_frequency_in_seconds = 3600
          }
        }
      ]
    }
  ]
}

# Create a sudo command bundle
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
