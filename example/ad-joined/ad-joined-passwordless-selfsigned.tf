//Gateway is mandatory for AD-Joined
// Separate gateway for AD Connection and RDP
data "oktapam_gateway" "ad_gateway" {
  id = "db778975-1776-4057-8360-6b9488c8f528"
}

// Create self-signed certificate
resource "oktapam_ad_passwordless_certificate" "self_signed" {
  name        = "dev-test"
  common_name = "test-cn"
  //API is designed to take this argument as a single value in days but we can have something more general too
  #  validity {
  #    type = "DAYS"
  #    value = 180
  #  }
  ttl_days          = 180
  enterprise_signed = false
}

// Create project with `forward_traffic` enabled & gateway selector
resource "oktapam_project" "test_project" {
  name                 = "project-one"
  next_unix_uid        = 60120
  next_unix_gid        = 63020
  create_server_users  = true
  ssh_certificate_type = "CERT_TYPE_ED25519_01"

  // Forward through gateway, gateway must have matching label
  forward_traffic  = true
  gateway_selector = "env=dev"
}

//Active Directory Connection with passwordless certificate configured in team settings
//Active Directory Connection with passwordless certificate configured in team settings
resource "oktapam_ad_connection" "testad" {
  name                        = "ad-connection"
  gateway_id                  = data.oktapam_gateway.ad_gateway.id
  domain                      = "dev-test.example.com"
  service_account_username    = "ldap@dev-test"
  service_account_password    = "secret"
  use_passwordless            = true // Boolean flag to indicate if passwordless feature is enabled. Other option is to implicitly set it if cert ifd provided
  passwordless_certificate_id = oktapam_ad_passwordless_certificate.self_signed.id
  domain_controller           = ["dc1.com", "dc2.com"] //Optional: DC used to query the domain
  depends_on = [
    oktapam_ad_passwordless_certificate.self_signed
  ]
}

//AD Joined Server Discovery
//Each Connection can have multiple server sync Jobs but only one active at a time
//One option is to name it : oktapam_ad_connection_server_sync_job
resource "oktapam_ad_connection_task" "daily_job" {
  connection                 = oktapam_ad_connection.testad.id
  name                       = "daily-job"
  is_active                  = true
  frequency                  = 24
  start_hour_utc             = 12 # Only required when frequency is 24 hours
  hostname_attribute         = var.ad_connection_task.hostname_attribute
  access_address_attribute   = var.ad_connection_task.access_address_attribute
  operating_system_attribute = var.ad_connection_task.operating_system_attribute
  bastion_attribute          = var.ad_connection_task.bastion_attribute
  alt_names_attributes       = var.ad_connection_task.alt_names_attributes
  rule_assignments {
    base_dn           = "ou=real,dc=dev-test,dc=sudo,dc=wtf"
    ldap_query_filter = "(objectclass=computer)"
    project_id        = "d7f82a2a-b547-49b3-b28f-8ebe948ad689"
    priority          = 1
  }
  depends_on = [
    oktapam_ad_connection.testad
  ]
}






