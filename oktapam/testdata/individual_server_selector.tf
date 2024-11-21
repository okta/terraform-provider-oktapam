resource "oktapam_security_policy_v2" "individual_server_policy" {
  name       = "test individual server selector"
  active     = true
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

