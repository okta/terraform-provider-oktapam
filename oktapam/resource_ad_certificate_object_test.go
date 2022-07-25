package oktapam

import (
	"context"
	"fmt"
	"testing"

	"github.com/okta/terraform-provider-oktapam/oktapam/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
)

const (
	certificateStatusValid = "valid"
)

func TestAccADCertificateObject(t *testing.T) {
	certificateUploadResourceName := "oktapam_ad_certificate_object.test-upload-cert"
	csrDisplayName := fmt.Sprintf("test-acc-csr-%s", randSeq(10))

	// Hashicorp TLS provider generate CA certificate and create signed certificate. It is added as external provider dependency
	tlsExternalProvider := resource.ExternalProvider{
		VersionConstraint: ">=4.0.0",
		Source:            "registry.terraform.io/hashicorp/tls",
	}
	externalProviders := map[string]resource.ExternalProvider{}
	externalProviders["tls"] = tlsExternalProvider

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		ExternalProviders: externalProviders,
		CheckDestroy:      utils.CreateCheckResourceDestroy(providerADCertificateSigningRequestKey, adCertificateExists),
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
resource "oktapam_ad_certificate_signing_request" "test-csr" {
  display_name = "%s"
  common_name = "testacc"
  details {
    organization = "test"
    organizational_unit = "asa"
    locality = "SF"
    province = "CA"
    country = "US"
  }
}

resource "tls_private_key" "ca-private-key" {
  algorithm   = "RSA"
  rsa_bits = "4096"
}

resource "tls_self_signed_cert" "ca-cert" {
  private_key_pem = tls_private_key.ca-private-key.private_key_pem

  subject {
    common_name  = "AccTest.com"
    organization = "Acc Test"
  }

  validity_period_hours = 12
  is_ca_certificate = true

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

resource "oktapam_ad_certificate_object" "test-upload-cert" {
  certificate_id = oktapam_ad_certificate_signing_request.test-csr.id
  source = tls_locally_signed_cert.signed-cert.cert_pem
}
`

func createTestAccADCertificateUploadConfig(csrDisplayName string) string {
	return fmt.Sprintf(testAccADCertificateUploadConfigFormat, csrDisplayName)
}

func adCertificateExists(id string) (bool, error) {
	client := testAccProvider.Meta().(client.OktaPAMClient)
	logging.Debugf("Checking if resource deleted %s", id)
	adCertificate, err := client.GetADSmartcardCertificate(context.Background(), id)

	return adCertificate != nil && adCertificate.Exists() && err == nil, err
}
