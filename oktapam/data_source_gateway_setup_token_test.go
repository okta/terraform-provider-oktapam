package oktapam

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDatasourceGatewaySetupToken(t *testing.T) {
	resourceName := "data.oktapam_gateway_setup_token.test_gateway_setup_tokens"
	identifier := randSeq(10)
	descriptionOne := fmt.Sprintf("Acceptance Test Setup Token - One: %s", identifier)
	descriptionTwo := fmt.Sprintf("Acceptance Test Setup Token - Two: %s", identifier)
	labels := make(map[string]string)
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key-%s", randSeq(10))
		value := fmt.Sprintf("value-%s", randSeq(10))
		labels[key] = value
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: createTestAccDatasourceGatewaySetupTokenCreateConfig(descriptionOne, descriptionTwo, labels),
			},
			{
				Config: testAccDatasourceGatewaySetupTokenConfig(identifier),
				Check:  resource.TestCheckResourceAttr(resourceName, "gateway_setup_tokens.#", "2"),
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

const testAccDatasourceGatewaySetupTokenCreateConfigFormat = `
resource "oktapam_gateway_setup_token" "test-gateway-setup-token-one" {
	description = "%s"
	labels = {
%s
	}
}

resource "oktapam_gateway_setup_token" "test-gateway-setup-token-two" {
	description = "%s"
	labels = {
%s
	}
}
`

func createTestAccDatasourceGatewaySetupTokenCreateConfig(descriptionOne string, descriptionTwo string, labels map[string]string) string {
	labelStrings := make([]string, 0, len(labels))
	for k, v := range labels {
		labelStrings = append(labelStrings, fmt.Sprintf("\t%s = \"%s\"", k, v))
	}
	labelBlock := strings.Join(labelStrings, "\n")
	return fmt.Sprintf(testAccDatasourceGatewaySetupTokenCreateConfigFormat, descriptionOne, labelBlock, descriptionTwo, labelBlock)
}
