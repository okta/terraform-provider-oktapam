resource "oktapam_ad_certificate_signing_request" "test_csr" {
  display_name = "test-csr"
  common_name = "testcsr"
  details {
    organization = "test"
    organizational_unit = "okta"
    locality = "SF"
    province = "CA"
    country = "US"
  }
}

resource "tls_private_key" "ca-private-key" {
  key_algorithm   = "rsa"
  ecdsa_curve = "4096"
}

resource "tls_self_signed_cert" "ca-cert" {
  private_key_pem = tls_private_key.ca-private-key.private_key_pem

  subject {
    common_name  = "test.com"
    organization = "Test Examples, Inc"
  }

  validity_period_hours = 12

  allowed_uses = [
    "cert_signing",
    "crl_signing",
    "client_auth"
  ]
}

resource "tls_locally_signed_cert" "signed-cert" {
  cert_request_pem   = oktapam_ad_certificate_signing_request.test-csr.content
  ca_private_key_pem = tls_private_key.ca-private-key.private_key_pem
  ca_cert_pem        = tls_self_signed_cert.ca-cert.cert_pem

  validity_period_hours = 12
  is_ca_certificate = true

  allowed_uses = [
    "cert_signing",
    "crl_signing",
    "client_auth"
  ]
}

resource "oktapam_ad_certificate_object" "upload-cert" {
  certificate_id = oktapam_ad_certificate_signing_request.test-csr.id
  source = tls_locally_signed_cert.signed-cert.cert_pem
}