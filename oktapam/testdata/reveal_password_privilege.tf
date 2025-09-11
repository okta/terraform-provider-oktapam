resource "oktapam_security_policy_v2" "reveal_password_test" {
  type        = "default"
  name        = "reveal-password-test-policy"
  description = "Test policy for reveal password privilege"
  active      = true
  principals = { user_groups = ["user_group_1", "user_group_2"] }
  rules = [
    {
      name          = "Reveal password access rule"
      resource_type = "server_based_resource"
      resource_selector = {
        server_based_resource = {
          selectors = [
            {
              individual_server = {
                server = "test-server-id"
                account_selector_type = "all_accounts"
                account_selector = {
                  all_accounts = true
                }
              }
            }
          ]
        }
      }

      privileges = [
        {
          reveal_password = {
            reveal_password = true
          }
        }
      ]
    }
  ]
}