package oktapam

import (
	"fmt"
	"testing"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/config"

	"github.com/okta/terraform-provider-oktapam/oktapam/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
)

const (
	certificateStatusValid = "valid"
)

func TestAccADCertificateObject(t *testing.T) {
	checkTeamApplicable(t, false)
	certificateUploadResourceName := "oktapam_ad_certificate_object.test_upload_cert"
	csrDisplayName := fmt.Sprintf("test-acc-csr-%s", randSeq())

	// Hashicorp TLS provider generate CA certificate and create signed certificate. It is added as external provider dependency
	tlsExternalProvider := resource.ExternalProvider{
		VersionConstraint: ">=4.0.0",
		Source:            "registry.terraform.io/hashicorp/tls",
	}
	externalProviders := map[string]resource.ExternalProvider{}
	externalProviders["tls"] = tlsExternalProvider

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccV6ProviderFactories,
		ExternalProviders:        externalProviders,
		CheckDestroy:             utils.CreateCheckResourceDestroy(config.ProviderADCertificateRequestKey, adCertificateExists),
		Steps: []resource.TestStep{
			{
				Config: createTestAccADCertificateUploadConfig(csrDisplayName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(certificateUploadResourceName, attributes.EnterpriseSigned, "true"),
					resource.TestCheckResourceAttr(certificateUploadResourceName, attributes.Status, certificateStatusValid),
				),
			},
		},
	})
}

const testAccADCertificateUploadConfigFormat = `
resource "oktapam_ad_certificate_request" "test_csr" {
  display_name = "%s"
  common_name  = "testacc"
  type         = "certificate_signing_request"
  details {
    organization        = "Okta"
    organizational_unit = "Okta Unit"
    locality            = "San Francisco"
    province            = "CA"
    country             = "US"
  }
}

resource "tls_private_key" "ca_private_key" {
  algorithm   = "RSA"
  rsa_bits = "4096"
}

resource "tls_self_signed_cert" "ca_cert" {
  private_key_pem = tls_private_key.ca_private_key.private_key_pem

  subject {
    common_name  = "example.com"
    organization = "Example, Inc"
  }

  validity_period_hours = 12
  is_ca_certificate = true

  allowed_uses = [
    "cert_signing",
    "crl_signing",
    "client_auth"
  ]
}

resource "tls_locally_signed_cert" "signed_cert" {
  cert_request_pem   = oktapam_ad_certificate_request.test_csr.content
  ca_private_key_pem = tls_private_key.ca_private_key.private_key_pem
  ca_cert_pem        = tls_self_signed_cert.ca_cert.cert_pem

  validity_period_hours = 12
  is_ca_certificate = true

  allowed_uses = [
    "cert_signing",
    "crl_signing",
    "client_auth"
  ]
}

resource "oktapam_ad_certificate_object" "test_upload_cert" {
  certificate_id = oktapam_ad_certificate_request.test_csr.id
  source = tls_locally_signed_cert.signed_cert.cert_pem
}
`

func createTestAccADCertificateUploadConfig(csrDisplayName string) string {
	return fmt.Sprintf(testAccADCertificateUploadConfigFormat, csrDisplayName)
}
