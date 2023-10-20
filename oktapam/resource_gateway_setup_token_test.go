package oktapam

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/kylelemons/godebug/pretty"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
)

func TestAccGatewaySetupToken(t *testing.T) {
	resourceName := "oktapam_gateway_setup_token.test_gateway_setup_token"

	identifier := randSeq()
	description := fmt.Sprintf("Acceptance Test Setup Token: %s", identifier)
	labels := constructLabels(10)

	setupToken := client.GatewaySetupToken{
		Description: &description,
		Details:     &client.GatewaySetupTokenLabelDetails{Labels: labels},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccGatewaySetupTokenCheckDestroy(identifier),
		Steps: []resource.TestStep{
			{
				Config: createTestAccGatewaySetupTokenCreateConfig(setupToken),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccGatewaySetupTokenCheckExists(resourceName, setupToken),
					resource.TestCheckResourceAttr(
						resourceName, attributes.Description, description,
					),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccGatewaySetupTokenCheckExists(rn string, expectedGatewaySetupToken client.GatewaySetupToken) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[rn]
		if !ok {
			return fmt.Errorf("resource not found: %s", rn)
		}

		resourceID := rs.Primary.ID

		client := getLocalClientFromMetadata(testAccProvider.Meta())
		token, err := client.GetGatewaySetupToken(context.Background(), resourceID)
		if err != nil {
			return fmt.Errorf("error getting gateway setup token: %w", err)
		}
		if token == nil {
			return fmt.Errorf("gateway setup token for with id %s does not exist", resourceID)
		}
		if *token.Token == "" {
			return fmt.Errorf("gateway setup token value for with id %s does not exist", resourceID)
		}
		if *token.Description != *expectedGatewaySetupToken.Description {
			return fmt.Errorf("expected description does not match returned description for gateway setup token.  expected: %s, got: %s", *expectedGatewaySetupToken.Description, *token.Description)
		}
		labelsCompare := pretty.Compare(expectedGatewaySetupToken.Details.Labels, token.Details.Labels)
		if labelsCompare != "" {
			return fmt.Errorf("expected labels does not match returned labels for gateway setup token:\n%s", labelsCompare)
		}

		return nil
	}
}

func testAccGatewaySetupTokenCheckDestroy(identifier string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := getLocalClientFromMetadata(testAccProvider.Meta())
		tokens, err := client.ListGatewaySetupTokens(context.Background(), identifier)
		if err != nil {
			return fmt.Errorf("error getting tokens: %w", err)
		}

		if len(tokens) > 0 {
			return fmt.Errorf("token still exists")
		}

		return nil
	}
}

const testAccGatewaySetupTokenCreateConfigFormat = `
resource "oktapam_gateway_setup_token" "test_gateway_setup_token" {
	description = "%s"
	labels = {
		%s
	}
}
`

func createTestAccGatewaySetupTokenCreateConfig(token client.GatewaySetupToken) string {
	labelStrings := make([]string, 0, len(token.Details.Labels))
	for k, v := range token.Details.Labels {
		labelStrings = append(labelStrings, fmt.Sprintf("\t%s = \"%s\"", k, v))
	}
	return fmt.Sprintf(testAccGatewaySetupTokenCreateConfigFormat, *token.Description, strings.Join(labelStrings, "\n"))
}
