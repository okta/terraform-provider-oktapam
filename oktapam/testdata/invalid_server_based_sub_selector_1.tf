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
    # rule with ssh privilege and admin + mfa every 1hr
    {
      name          = "linux server account and admin level access"
      resource_type = "server_based_resource"
      resource_selector = {
        server_based_resource = {
          selectors = [
            { /* invalid sub selector */ }
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
    }
  ]
}
