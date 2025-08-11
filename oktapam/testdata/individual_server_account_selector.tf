resource "oktapam_security_policy_v2" "individual_server_account_policy" {
  name       = "test individual server account selector"
  active     = true
  principals = { 
    user_groups = [
      { id = "user_group_1" },
      { id = "user_group_2" }
    ]
  }
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