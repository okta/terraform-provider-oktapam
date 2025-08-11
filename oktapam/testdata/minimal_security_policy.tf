resource "oktapam_security_policy_v2" "test" {
  type        = "default"
  name        = "minimal test policy"
  description = "Minimal policy for testing import functionality"
  active      = true
  principals = {
    user_groups = [
      { id = "user_group_1" }
    ]
  }

  rules = [
    {
      name          = "minimal rule"
      resource_type = "server_based_resource"
      resource_selector = {
        server_based_resource = {
          selectors = [
            {
              individual_server = {
                server = "test-server-id"
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
        }
      ]
    }
  ]
}