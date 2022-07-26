package oktapam

import (
	"fmt"
	"strings"
	"testing"

	"github.com/kylelemons/godebug/pretty"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDatasourceGatewaySetupTokenFetch(t *testing.T) {
	resourceName := "oktapam_gateway_setup_token.test-gateway-setup-token-1"
	dataSourceName := "data.oktapam_gateway_setup_token.target_token"

	identifier := randSeq(10)
	description := fmt.Sprintf("Acceptance Test Setup Token Set %s: 1", identifier)
	labels := constructLabels(10)

	testConfig := createTestAccDatasourceGatewaySetupTokenInitConfig(identifier, description, labels)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccGatewaySetupTokenCheckDestroy(identifier),
		Steps: []resource.TestStep{
			{
				Config: testConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					checkResourcesEqual(resourceName, dataSourceName),
				),
			},
		},
	})
}

func constructLabels(length int) map[string]string {
	labels := make(map[string]string)
	for i := 0; i < length; i++ {
		key := fmt.Sprintf("key-%s", randSeq(10))
		value := fmt.Sprintf("value-%s", randSeq(10))
		labels[key] = value
	}
	return labels
}

func checkResourcesEqual(resourceName1 string, resourceName2 string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resource1, ok := s.RootModule().Resources[resourceName1]
		if !ok {
			return fmt.Errorf("resource 1 not found: %s", resourceName1)
		}

		resource2, ok := s.RootModule().Resources[resourceName2]
		if !ok {
			return fmt.Errorf("resource 2 not found: %s", resourceName2)
		}

		comparison := pretty.Compare(resource1.Primary.Attributes, resource2.Primary.Attributes)
		if comparison != "" {
			return fmt.Errorf("resources are not equal")
		}
		return nil
	}
}

// NOTE: This config (1) creates a new token (2) lists the existing tokens with the matching description
// (a unique identifier is passed in, so the list will return only one id) and (3) get the new token as a oktapam_gateway_setup_token.
// The test then compares the token resource with this data source to ensure they are equal.

const testAccDatasourceGatewaySetupTokenInitListFetchConfigFormat = `
resource "oktapam_gateway_setup_token" "test-gateway-setup-token-1" {
	description = "%s"
	labels = {
		%s
	}
}

data "oktapam_gateway_setup_tokens" "token_list" {
	depends_on = [oktapam_gateway_setup_token.test-gateway-setup-token-1]
	description_contains = "%s"
}

data "oktapam_gateway_setup_token" "target_token" {
	id = data.oktapam_gateway_setup_tokens.token_list.ids[0]
}
`

func createTestAccDatasourceGatewaySetupTokenInitConfig(identifier string, description1 string, labels map[string]string) string {
	labelStrings := make([]string, 0, len(labels))
	for k, v := range labels {
		labelStrings = append(labelStrings, fmt.Sprintf("\t%s = %q", k, v))
	}
	labelBlock := strings.Join(labelStrings, "\n")
	return fmt.Sprintf(testAccDatasourceGatewaySetupTokenInitListFetchConfigFormat, description1, labelBlock, identifier)
}
