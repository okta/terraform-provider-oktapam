package oktapam

import (
	"context"
	"fmt"
	"testing"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/typed_strings"

	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
)

const (
	csrStatusCreated = "request_created"
	certStatusValid  = "valid"
)

func TestAccADCertificateRequest_CSR(t *testing.T) {
	resourceName := "oktapam_ad_certificate_request.test_csr"
	csrName := fmt.Sprintf("test-acc-csr-%s", randSeq())

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      utils.CreateCheckResourceDestroy(providerADCertificateRequestKey, adCertificateExists),
		Steps: []resource.TestStep{
			{
				//Ensure CSR Creation Works
				Config: createTestAccCSRCreateConfig(csrName, typed_strings.ADCertificateTypeSigningRequest),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, attributes.DisplayName, csrName),
					resource.TestCheckResourceAttr(resourceName, attributes.EnterpriseSigned, "true"),
					resource.TestCheckResourceAttr(resourceName, attributes.Status, csrStatusCreated),
				),
			},
		},
	})
}

func TestAccADCertificateRequest_SelfSigned(t *testing.T) {
	resourceName := "oktapam_ad_certificate_request.test_self_signed_cert"
	certName := fmt.Sprintf("test-acc-self-signed-cert-%s", randSeq())

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      utils.CreateCheckResourceDestroy(providerADCertificateRequestKey, adCertificateExists),
		Steps: []resource.TestStep{
			{
				//Ensure Self Signed Cert Creation Works
				Config: createTestAccSelfSignedCertCreateConfig(certName, typed_strings.ADCertificateTypeSelfSigned),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, attributes.DisplayName, certName),
					resource.TestCheckResourceAttr(resourceName, attributes.EnterpriseSigned, "false"),
					resource.TestCheckResourceAttr(resourceName, attributes.Status, certStatusValid),
				),
			},
		},
	})
}

func adCertificateExists(id string) (bool, error) {
	client := testAccProvider.Meta().(client.OktaPAMClient)
	logging.Debugf("Checking if resource deleted %s", id)
	adCertificate, err := client.GetADSmartcardCertificate(context.Background(), id)

	return adCertificate != nil && adCertificate.Exists() && err == nil, err
}

const testAccCSRCreateConfigFormat = `
resource "oktapam_ad_certificate_request" "test_csr" {
  display_name = "%s"
  common_name  = "testacc"
  type         = "%s"
  details {
    organization        = "Okta"
    organizational_unit = "Okta Unit"
    locality            = "San Francisco"
    province            = "CA"
    country             = "US"
  }
}
`

const testAccSelfSignedCertCreateConfigFormat = `
resource "oktapam_ad_certificate_request" "test_self_signed_cert" {
  display_name = "%s"
  common_name  = "testacc"
  type         = "%s"
  details {
	ttl_days = 90
  }
}
`

func createTestAccCSRCreateConfig(displayName string, certificateType typed_strings.ADCertificateType) string {
	return fmt.Sprintf(testAccCSRCreateConfigFormat, displayName, certificateType.String())
}

func createTestAccSelfSignedCertCreateConfig(displayName string, certificateType typed_strings.ADCertificateType) string {
	return fmt.Sprintf(testAccSelfSignedCertCreateConfigFormat, displayName, certificateType.String())
}
