package oktapam

import (
	"fmt"
	"strings"
	"testing"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDatasourceGatewaySetupTokenList(t *testing.T) {
	prefix := "data.oktapam_gateway_setup_tokens"

	// Generate details
	identifier := randSeq(10)
	description1 := fmt.Sprintf("Acceptance Test Setup Token Set %s: 1", identifier)
	description2 := fmt.Sprintf("Acceptance Test Setup Token Set %s: 2", identifier)
	labels := constructLabels(10)

	// Config 1: create two tokens
	initConfig := createTestAccDatasourceGatewaySetupTokensInitConfig(description1, description2, labels)

	// Config 2: list using filter that returns both
	dataName := "data1"
	dataFullName := fmt.Sprintf("%s.%s", prefix, dataName)
	listConfig := testAccDatasourceGatewaySetupTokensConfig(dataName, identifier)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccGatewaySetupTokenCheckDestroy(identifier),
		Steps: []resource.TestStep{
			{
				Config: initConfig,
			},
			{
				Config: listConfig,
				Check:  resource.TestCheckResourceAttr(dataFullName, fmt.Sprintf("%s.#", attributes.IDs), "2"),
			},
		},
	})
}

const testAccDatasourceGatewaySetupTokensInitConfigFormat = `
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

func createTestAccDatasourceGatewaySetupTokensInitConfig(description1 string, description2 string, labels map[string]string) string {
	labelStrings := make([]string, 0, len(labels))
	for k, v := range labels {
		labelStrings = append(labelStrings, fmt.Sprintf("\t%s = %q", k, v))
	}
	labelBlock := strings.Join(labelStrings, "\n")
	return fmt.Sprintf(testAccDatasourceGatewaySetupTokensInitConfigFormat, description1, labelBlock, description2, labelBlock)
}

const testAccDatasourceGatewaySetupTokensFormat = `
data "oktapam_gateway_setup_tokens" "%s" {
	description_contains = "%s"
}
`

func testAccDatasourceGatewaySetupTokensConfig(name string, descriptionContains string) string {
	return fmt.Sprintf(testAccDatasourceGatewaySetupTokensFormat, name, descriptionContains)
}
