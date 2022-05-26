// Create project with `forward_traffic` enabled & gateway selector
resource "oktapam_project" "test_project" {
  name                 = "project-one"
  next_unix_uid        = 60120
  next_unix_gid        = 63020
  create_server_users  = true
  ssh_certificate_type = "CERT_TYPE_ED25519_01"

  // Forward through gateway, gateway must have matching label
  forward_traffic  = true
  gateway_selector = "env=test"
}

// We are not creating users in ASA. Provisioned from Okta tenant.
// We can disable/delete provisioned users
resource "oktapam_user" "test_user" {
  name   = "test-user"
  status = "ACTIVE" //Possible values ACTIVE / DISABLED / DELETED
}

resource "oktapam_project_user" "test-project-user" {
  project_name = oktapam_project.test-project.name
  user_name    = oktapam_user.test_user
  //Override user attributes at project level
  ad_identities              = ["dev@test.com", "qa@test.com"]
  ad_passwordless_identities = ["dev@test.com", "qa@test.com"]
}