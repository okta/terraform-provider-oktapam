package oktapam

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

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

	// Config 2: list using filter that will only return one token
	data1Name := "data1"
	data1FullName := fmt.Sprintf("%s.%s", prefix, data1Name)
	listOneConfig := testAccDatasourceGatewaySetupTokensConfig(data1Name, fmt.Sprintf("%s: 1", identifier))

	// Config 3: list using filter that returns both
	data2Name := "data2"
	data2FullName := fmt.Sprintf("%s.%s", prefix, data2Name)
	listBothConfig := testAccDatasourceGatewaySetupTokensConfig(data2Name, identifier)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccGatewaySetupTokenCheckDestroy(identifier),
		Steps: []resource.TestStep{
			{
				Config: initConfig,
			},
			{
				Config: listOneConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					check(),
					resource.TestCheckResourceAttr(data1FullName, fmt.Sprintf("%s.#", attributes.IDs), "1"),
				),
			},
			{
				Config: listBothConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					check(),
					resource.TestCheckResourceAttr(data2FullName, fmt.Sprintf("%s.#", attributes.IDs), "2"),
				),
			},
		},
	})
}

func check() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		fmt.Println("STATE:", s.RootModule().String())
		return nil
	}
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
		labelStrings = append(labelStrings, fmt.Sprintf("\t%s = \"%s\"", k, v))
	}
	labelBlock := strings.Join(labelStrings, "\n")
	return fmt.Sprintf(testAccDatasourceGatewaySetupTokensInitConfigFormat, description1, labelBlock, description2, labelBlock)
}

const testAccDatasourceGatewaySetupTokensFormat = `
data "oktapam_gateway_setup_tokens" "%s" {
	description_contains = "%s"
}
`

func testAccDatasourceGatewaySetupTokensConfig(name, descriptionContains string) string {
	return fmt.Sprintf(testAccDatasourceGatewaySetupTokensFormat, name, descriptionContains)
}
