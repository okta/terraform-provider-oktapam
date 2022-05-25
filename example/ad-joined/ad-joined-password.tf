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

//Active Directory Connection with passwordless certificate configured in team settings
resource "oktapam_connection" "ad-connection" {
  name                 = "ad-connection"
  gateway_id           = data.oktapam_gateway.ubuntu.id
  domain               = "dev-test.example.com"
  service_user         = "ldap@dev-test"
  service_user_secret  = "secret"
  passwordless_enabled = false
  domain_controller    = ["dc1", "dc2"] //Optional: DC used to query the domain
}

//AD Joined Server Discovery
//Each Connection can have multiple server sync Jobs but only one active at a time
resource "oktapam_connection_server_sync_job" "daily-job" {
  connection = oktapam_connection.ad-connection.id
  name       = "daily-job"
  job_status = "active"
  depends_on = [
    oktapam_connection.ad-connection
  ]
}

resource "oktapam_connection_server_sync_job" "daily-job-schedule" {
  server_sync_job_id = oktapam_connection_server_sync_job.daily-job.id
  schedule_type      = "hourly" //Possible values hourly, daily, every 6 hours, 12 hours
  interval           = "1"
  depends_on = [
    oktapam_connection_server_sync_job.daily-job
  ]
}

//Each Job can have multiple sync rules with different priorities for server discovery
resource "oktapam_connection_server_sync_job_rules" "rule1" {
  server_sync_job_id  = oktapam_connection_server_sync_job.daily-job.id
  name                = "rule1"
  priority            = "1"
  baseDN              = "ou=technology,dc=dev-test,dc=com"
  ldapQuery           = "(objectCategory=Computer)"
  assigned_project_id = oktapam_project.test-project.id
  depends_on = [
    oktapam_connection_server_sync_job.daily-job
  ]
}