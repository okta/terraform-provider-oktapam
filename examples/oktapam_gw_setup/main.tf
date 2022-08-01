terraform {
  required_version = ">= 0.12.0"

  required_providers {
    oktapam = {
      version = ">=0.1.4"
      source  = "okta.com/pam/oktapam"
    }
    google = {
      source = "hashicorp/google"
      version = "4.30.0"
    }
  }
}

provider "google" {
    # Configuration options
    project	= "asa-demo-316514"
    region	= "us-central1"
    zone	= "us-central1-a"
}

