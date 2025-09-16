resource "oktapam_group" "test_security_policy_group1" {
  name = "%s"
}
resource "oktapam_group" "test_security_policy_group2" {
  name = "%s"
}
resource "oktapam_sudo_command_bundle" "test_acc_sudo_command_bundle" {
  name = "%s"
  structured_commands {
    command       = "/bin/run.sh"
    command_type  = "executable"
    args_type     = "custom"
    args          = "ls"
  }
}
resource "oktapam_security_policy" "test_acc_security_policy" {
  name = "%s"
  description = "test description"
  active = true
  principals {
    groups = [oktapam_group.test_security_policy_group1.id, oktapam_group.test_security_policy_group2.id]
  }
  rule {
    name = "first rule"
    resources {
      servers {
        server {
          server_id = "%s"
        }
        label_selectors {
          server_labels = {
            "system.os_type" = "linux"
          }
          accounts = ["root", "pamadmin"]
        }
      }
    }
    privileges {
      password_checkout_ssh {
        enabled = true
      }
      principal_account_ssh {
        enabled = true
        admin_level_permissions = false
        sudo_display_name = "foo-uam"
        sudo_command_bundles = [oktapam_sudo_command_bundle.test_acc_sudo_command_bundle.id]
      }
    }
    conditions {
      access_request {
        request_type_id = "abcd"
        request_type_name = "foo"
        expires_after_seconds = 1200
      }
      access_request {
        request_type_id = "wxyz"
        request_type_name = "bar"
        expires_after_seconds = 1800
      }
    }
  }
}