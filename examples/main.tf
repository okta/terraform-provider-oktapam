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

resource "oktapam_project" "test-project-one" {
  name = "project-one"
  next_unix_uid = 60120
  next_unix_gid = 63020
  ssh_certificate_type = "CERT_TYPE_RSA_01"
}

resource "oktapam_project" "test-project-two" {
  name = "project-two"
  next_unix_uid = 60220
  next_unix_gid = 63120
  ssh_certificate_type = "CERT_TYPE_ED25519_01"
}