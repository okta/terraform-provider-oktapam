resource "oktapam_ad_certificate_request" "test_csr" {
  display_name = "test-csr"
  common_name  = "testcsr"
  type         = "certificate_signing_request"
  details {
    organization        = "test"
    organizational_unit = "okta"
    locality            = "SF"
    province            = "CA"
    country             = "US"
  }
}

# Generates a secure private key and encodes it in PEM (RFC 1421)
resource "tls_private_key" "ca_private_key" {
  key_algorithm = "rsa"
  ecdsa_curve   = "4096"
}

# Creates a self-signed TLS certificate for CA in PEM (RFC 1421) format
resource "tls_self_signed_cert" "ca_cert" {
  private_key_pem = tls_private_key.ca_private_key.private_key_pem
  subject {
    common_name  = "example.com"
    organization = "Example, Inc"
  }
  validity_period_hours = 12
  is_ca_certificate     = true

  allowed_uses = [
    "cert_signing",
    "crl_signing",
    "client_auth"
  ]
}

# Creates a TLS certificate in PEM (RFC 1421) format using a Certificate Signing Request (CSR) and signs it with a provided (local) Certificate Authority (CA)
resource "tls_locally_signed_cert" "signed_cert" {
  cert_request_pem      = oktapam_ad_certificate_request.test_csr.content
  ca_private_key_pem    = tls_private_key.ca_private_key.private_key_pem
  ca_cert_pem           = tls_self_signed_cert.ca_cert.cert_pem
  validity_period_hours = 12
  is_ca_certificate     = true

  allowed_uses = [
    "cert_signing",
    "crl_signing",
    "client_auth"
  ]
}

# Upload signed certificate
resource "oktapam_ad_certificate_object" "upload_cert" {
  certificate_id = oktapam_ad_certificate_request.test_csr.id
  source         = tls_locally_signed_cert.signed_cert.cert_pem
}