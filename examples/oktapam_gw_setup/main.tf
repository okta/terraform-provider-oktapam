terraform {
  required_version = ">= 0.12.0"

  required_providers {
    oktapam = {
      version = ">=0.2.0"
      source  = "okta/oktapam"
    }
    google = {
      source = "hashicorp/google"
      version = "4.30.0"
    }
  }
}

provider "google" {
    # Configuration options
    project	= "asa-demo-project"
    region	= "us-central1"
    zone	= "us-central1-a"
}

