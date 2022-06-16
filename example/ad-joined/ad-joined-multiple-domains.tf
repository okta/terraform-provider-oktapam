//Gateway is mandatory for AD-Joined
data "oktapam_gateway" "ubuntu" {
  name = "ubuntu"
}

// Create project with `forward_traffic` enabled & gateway selector
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

//Active Directory Connection with domain 1
resource "oktapam_ad_connection" "connection1" {
  name                 = "ad-connection"
  gateway_id           = data.oktapam_gateway.ubuntu.id
  domain               = "domain1.example.com"
  service_user         = "ldap@dev-test"
  service_user_secret  = "secret"
  passwordless_enabled = false          // Boolean flag to indicate if passwordless feature is enabled. Other option is to implicitly set it if cert ifd provided
  domain_controller    = ["dc1", "dc2"] //Optional: DC used to query the domain
}

//Active Directory Connection with domain 2
resource "oktapam_ad_connection" "connection2" {
  name                 = "ad-connection"
  gateway_id           = data.oktapam_gateway.ubuntu.id
  domain               = "domain1.example.com"
  service_user         = "ldap@dev-test"
  service_user_secret  = "secret"
  passwordless_enabled = false          // Boolean flag to indicate if passwordless feature is enabled. Other option is to implicitly set it if cert ifd provided
  domain_controller    = ["dc1", "dc2"] //Optional: DC used to query the domain
}

//AD Joined Server Discovery
//Each Connection can have multiple server sync Jobs but only one active at a time
resource "oktapam_ad_connection_discovery_job" "daily-job" {
  connection     = oktapam_ad_connection.testad.id
  name           = "daily-job"
  job_status     = "active"
  frequency      = 24
  start_hour_utc = 12 # Only required when frequency is 24 hours
  depends_on = [
    oktapam_ad_connection.testad
  ]
  rule_assignments {
    base_dn           = "ou=real,dc=dev-test,dc=sudo,dc=wtf"
    ldap_query_filter = "(objectclass=computer)"
    project_id        = "d7f82a2a-b547-49b3-b28f-8ebe948ad689"
    priority          = 1
  }
}





