package oktapam

import (
	"fmt"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDatasourceGatewaySetupToken(t *testing.T) {
	resourceName := "data.oktapam_gateway_setup_token.test_gateway_setup_tokens"
	identifier := randSeq(10)
	description1 := fmt.Sprintf("Acceptance Test Setup Token - 1: %s", identifier)
	description2 := fmt.Sprintf("Acceptance Test Setup Token - 2: %s", identifier)
	labels := make(map[string]string)
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key-%s", randSeq(10))
		value := fmt.Sprintf("value-%s", randSeq(10))
		labels[key] = value
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: createTestAccDatasourceGatewaySetupTokenInitConfig(description1, description2, labels),
			},
			{
				Config: testAccDatasourceGatewaySetupTokenConfig(identifier),
				Check:  resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("%s.#", attributes.GatewaySetupTokens), "2"),
			},
		},
	})
}

const testAccDatasourceGatewaySetupTokenFormat = `
data "oktapam_gateway_setup_token" "test_gateway_setup_tokens" {
	description_contains = "%s"
}
`

func testAccDatasourceGatewaySetupTokenConfig(descriptionContains string) string {
	return fmt.Sprintf(testAccDatasourceGatewaySetupTokenFormat, descriptionContains)
}

const testAccDatasourceGatewaySetupTokenInitConfigFormat = `
resource "oktapam_gateway_setup_token" "test-gateway-setup-token-1" {
	description = "%s"
	labels = {
%s
	}
}

resource "oktapam_gateway_setup_token" "test-gateway-setup-token-2" {
	description = "%s"
	labels = {
%s
	}
}
`

func createTestAccDatasourceGatewaySetupTokenInitConfig(description1 string, description2 string, labels map[string]string) string {
	labelStrings := make([]string, 0, len(labels))
	for k, v := range labels {
		labelStrings = append(labelStrings, fmt.Sprintf("\t%s = \"%s\"", k, v))
	}
	labelBlock := strings.Join(labelStrings, "\n")
	return fmt.Sprintf(testAccDatasourceGatewaySetupTokenInitConfigFormat, description1, labelBlock, description2, labelBlock)
}
