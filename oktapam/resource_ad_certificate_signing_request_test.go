package oktapam

import (
	"context"
	"fmt"
	"testing"

	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
)

const (
	csrStatusCreated = "request_created"
)

func TestAccADCertificateSigningRequest(t *testing.T) {
	resourceName := "oktapam_ad_certificate_signing_request.test-csr"
	csrName := fmt.Sprintf("test-acc-csr-%s", randSeq(10))

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      utils.CreateCheckResourceDestroy(providerADCertificateSigningRequestKey, adCSRExists),
		Steps: []resource.TestStep{
			{
				Config: createTestAccCSRCreateConfig(csrName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, attributes.DisplayName, csrName),
					resource.TestCheckResourceAttr(resourceName, attributes.EnterpriseSigned, "true"),
					resource.TestCheckResourceAttr(resourceName, attributes.Status, csrStatusCreated),
				),
			},
		},
	})
}

func adCSRExists(id string) (bool, error) {
	client := testAccProvider.Meta().(client.OktaPAMClient)
	logging.Debugf("Checking if resource deleted %s", id)
	adCSR, err := client.GetADSmartcardCertificate(context.Background(), id)

	return adCSR != nil && adCSR.Exists() && err == nil, err
}

const testAccCSRCreateConfigFormat = `
resource "oktapam_ad_certificate_signing_request" "test-csr" {
  display_name = "%s"
  common_name = "testacc"
  details {
    organization = "Okta"
    organizational_unit = "Okta Unit"
    locality = "San Francisco"
    province = "CA"
    country = "US"
  }
}
`

func createTestAccCSRCreateConfig(groupName string) string {
	return fmt.Sprintf(testAccCSRCreateConfigFormat, groupName)
}
