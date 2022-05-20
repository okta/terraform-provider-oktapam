terraform {
  required_version = ">= 0.12.0"

  required_providers {
    oktapam = {
      version = "0.1.0"
      source  = "hashicorp.com/okta/oktapam"
    }
  }
}

provider "oktapam" {
  // Required
  oktapam_key = var.oktapam_key
  oktapam_secret = var.oktapam_secret
  oktapam_team = var.oktapam_team

  // Optional
  oktapam_api_host = var.oktapam_api_host
}

// Create project with `forward_traffic` enabled
resource "oktapam_project" "test-project" {
  name = "project-one"
  next_unix_uid = 60120
  next_unix_gid = 63020
  create_server_users = true
  ssh_certificate_type = "CERT_TYPE_ED25519_01"

  // Forward through gateway, gateway must have matching label
  forward_traffic = true
  gateway_selector = "env=test"
}

// Create group with `access_user` permissions
resource "oktapam_group" "test-group" {
  name = "group-one"
  roles = ["access_user"]
}

// Attach group to project
resource "oktapam_project_group" "test-project-group" {
  project_name = oktapam_project.test-project.name
  group_name = oktapam_group.test-group.name
  server_access = true
  server_admin = false
  depends_on = [
    oktapam_project.test-project,
    oktapam_group.test-group,
  ]
}

// Create server enrollment token for project
resource "oktapam_server_enrollment_token" "test-server-token" {
  description = "token for server enrollment"
  project_name = oktapam_project.test-project.name
}

// Create gateway used for traffic-forwarding
resource "oktapam_gateway_setup_token" "test-gateway-token" {
  description = "token for gateway enrollment"

  // Gateway is assigned to project if labels satisfy all gateway_selectors
  labels = {env = "test"}
}

// End result:
// If the resulting server enrollment token and gateway setup token are added to an sftd agent
// and sft-gatewayd agent, users on `group-one` could access `sftd` via `sft-gatewayd`.
// This is left as an exercise to the user.
