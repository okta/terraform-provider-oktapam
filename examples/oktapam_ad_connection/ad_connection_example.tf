
# Create project with `forward_traffic` enabled
resource "oktapam_project" "test-project" {
  name                 = "project-one"
  next_unix_uid        = 60120
  next_unix_gid        = 63020
  create_server_users  = true
  ssh_certificate_type = "CERT_TYPE_ED25519_01"

  // Forward through gateway, gateway must have matching label
  forward_traffic  = true
  gateway_selector = "env=test"
}

# Creating gateway is not supported. Get the gateway id using datasource
data "oktapam_gateway" "gateway-list" {
  contains = "ubuntu" # Filter gateway that contains given name
}

# Create active Directory Connection
resource "oktapam_ad_connection" "test-ad-connection" {
  name                     = "test-ad-connection"
  gateway_id               = data.oktapam_gateway.gateways.gateways[0].id
  domain                   = "dev-test.test.com"
  service_account_username = "ldap@dev-test"
  service_account_password = "secret"
  use_passwordless         = true
  domain_controllers       = ["dc1.com", "dc2.com"] //Optional: DC used to query the domain
}

# AD Joined Server Discovery
# Each Connection can have multiple server sync Jobs but only one active at a time
resource "oktapam_ad_task_settings" "test-ad-task-settings" {
  connection_id            = oktapam_ad_connection.test-ad-connection.id
  name                     = "daily-job-1"
  is_active                = true
  frequency                = 12 # Every 12 hours Note: If 24 hours then start_hour_utc is required
  host_name_attribute      = "dNSHostName"
  access_address_attribute = "dNSHostName"
  os_attribute             = "operatingSystem"
  rule_assignments {
    base_dn           = "ou=real,dc=dev-test,dc=sudo,dc=wtf"
    ldap_query_filter = "(objectclass=computer)"
    project_id        = oktapam_project.test-project.id
    priority          = 1
  }
}

