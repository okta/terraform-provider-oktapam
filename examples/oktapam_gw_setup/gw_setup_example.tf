# Create gateway setup token
resource "oktapam_gateway_setup_token" "test-gw" {
  description          = "tf-gw"
  labels               = {env:"prod"}
}

data "template_file" "startup_script" {
	template = file("${path.module}/startup.sh")
	vars = {
		gwtoken = oktapam_gateway_setup_token.test-gw.token
	}
}


# Create a GCP instance that will host the gateway
resource "google_compute_instance" "default" {
  # (resource arguments)
  name 		= "tfgw1"
  machine_type  = "e2-medium"
  zone		= "us-central1-a"
  boot_disk {
	initialize_params {
		image = "ubuntu-2004-focal-v20220712"
	}
  }
  network_interface {
	network = "default"
	subnetwork = "asa-private-subnet-2"
        access_config {}
  }
  
  # Startup script that puts the gateway setup token on the server 
  # and installs the gateway package
  metadata_startup_script = data.template_file.startup_script.rendered
}

# Create a Project and configure it to use the gateway we just created
resource "oktapam_project" "tfproject" {
	name = "tfproject"
	gateway_selector = "env.prod"
	ssh_certificate_type = "CERT_TYPE_ECDSA_256_01"
	forward_traffic = true
}
