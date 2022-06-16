//Gateway is mandatory for AD-Joined
// Separate gateway for AD Connection and RDP
data "oktapam_gateway" "ad_gateway" {
  id = "db778975-1776-4057-8360-6b9488c8f528"
}

// Create project with `forward_traffic` enabled & gateway selector
resource "oktapam_project" "test_project" {
  name                 = "test_project"
  next_unix_uid        = 60120
  next_unix_gid        = 63020
  create_server_users  = true
  ssh_certificate_type = "CERT_TYPE_ED25519_01"

  // Forward through gateway, gateway must have matching label
  forward_traffic  = true
  gateway_selector = "env=test"
}

//Active Directory Connection
resource "oktapam_ad_connection" "testad" {
  name                     = "ad-connection"
  gateway_id               = data.oktapam_gateway.ad_gateway.id
  domain                   = "dev-test.example.com"
  service_account_username = "ldap@dev-test"
  service_account_password = "secret"
  use_passwordless         = false                  // Boolean flag to indicate if passwordless feature is enabled. Other option is to implicitly set it if cert ifd provided
  domain_controller        = ["dc1.com", "dc2.com"] //Optional: DC used to query the domain
}

//AD Joined Server Discovery
//Each Connection can have multiple server sync Jobs but only one active at a time
resource "oktapam_ad_task_settings" "hourly_job" {
  connection                 = oktapam_ad_connection.testad.id
  name                       = "daily-job"
  is_active                  = true
  frequency                  = 12 # Every 12 hours
  host_name_attribute        = var.ad_connection_task.hostname_attribute
  access_address_attribute   = var.ad_connection_task.access_address_attribute
  operating_system_attribute = var.ad_connection_task.operating_system_attribute
  bastion_attribute          = var.ad_connection_task.bastion_attribute
  alt_names_attributes       = var.ad_connection_task.alt_names_attributes
  depends_on                 = [
    oktapam_ad_connection.testad
  ]
  rule_assignments {
    //Each Job can have multiple sync rules with different priorities for server discovery
    base_dn           = "ou=real,dc=dev-test,dc=sudo,dc=wtf"
    ldap_query_filter = "(objectclass=computer)"
    project_id        = "d7f82a2a-b547-49b3-b28f-8ebe948ad689"
    priority          = 1
  }
}








