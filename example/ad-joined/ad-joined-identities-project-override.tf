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

// Create group with `access_user` permissions
resource "oktapam_group" "test_group" {
  name  = "group-one"
  roles = ["access_user"]
}

// Attach group to project
resource "oktapam_project_group" "test_project_group" {
  project_name  = oktapam_project.test-project.name
  group_name    = oktapam_group.test_group.name
  server_access = true
  server_admin  = false
  depends_on = [
    oktapam_project.test-project,
    oktapam_group.test_group,
  ]
}

// We are not creating users in ASA. Provisioned from Okta tenant.
// We can disable/delete provisioned users
resource "oktapam_user" "test_user" {
  name   = "test-user"
  status = "ACTIVE" //Possible values ACTIVE / DISABLED / DELETED
}

//Assign group to user
resource "oktapam_group_user" "test_group_user" {
  group_name = oktapam_group.test_group.name
  user_name  = oktapam_user.test_user.name
  depends_on = [
    oktapam_group.test_group,
    oktapam_user.test_user
  ]
}

resource "oktapam_project_user_attributes" "test-project-user" {
  project_name = oktapam_project.test-project.name
  user_name    = oktapam_user.test_user
  //Override user attributes at project level
  ad_identities              = ["dev@test.com", "qa@test.com"]
  ad_passwordless_identities = ["dev@test.com", "qa@test.com"]
}