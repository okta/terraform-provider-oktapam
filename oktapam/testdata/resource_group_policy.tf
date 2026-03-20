resource "oktapam_security_policy_v2" "resource_group_policy" {
  type           = "default"
  name           = "resource group test policy"
  description    = "Test policy with resource group"
  active         = true
  resource_group = "a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d"
  principals = { user_groups = ["user_group_1"] }
  rules = [
    {
      name          = "basic rule"
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
