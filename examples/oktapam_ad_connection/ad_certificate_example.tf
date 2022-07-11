resource "oktapam_ad_certificate_signing_request" "tilt_ad_csr" {
  count        = var.enable_ad_join ? 1 : 0
  display_name = "test-csr-1"
  common_name  = "tilt"
  details {
    organization        = "test"
    organizational_unit = "asa"
    locality            = "SF"
    province            = "CA"
    country             = "US"
  }
}

data "local_file" "signed_cert" {
  filename = "<path>"
}

resource "oktapam_ad_certificate_object" "test-csr" {
  certificate_id = oktapam_ad_certificate_signing_request.test-csr.id
  source = data.local_file.signed_cert.content
  file_name = "test.crt"
}