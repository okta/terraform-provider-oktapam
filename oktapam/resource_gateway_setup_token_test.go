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
	resourceName := "oktapam_gateway_setup_token.test-gateway-setup-token"
	description := fmt.Sprintf("Acceptance Test Setup Token: %s", randSeq(10))
	labels := make(map[string]string)
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key-%s", randSeq(10))
		value := fmt.Sprintf("value-%s", randSeq(10))
		labels[key] = value
	}
	setupToken := client.GatewaySetupToken{
		Description: &description,
		Details:     &client.GatewaySetupTokenLabelDetails{Labels: labels},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccGatewaySetupTokenCheckDestroy(setupToken),
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

		client := testAccProvider.Meta().(client.OktaPAMClient)
		token, err := client.GetGatewaySetupToken(context.Background(), resourceID)
		if err != nil {
			return fmt.Errorf("error getting gateway setup token: %w", err)
		}
		if token == nil {
			return fmt.Errorf("gateway setup token for with id %s does not exist", resourceID)
		}
		if token.Token == nil {
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

func testAccGatewaySetupTokenCheckDestroy(token client.GatewaySetupToken) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(client.OktaPAMClient)
		tokens, err := client.ListGatewaySetupTokens(context.Background(), *token.Description)
		if err != nil {
			return fmt.Errorf("error getting tokens: %w", err)
		}

		for _, t := range tokens {
			if token.Description == t.Description {
				return fmt.Errorf("token still exists")
			}
		}
		return nil
	}
}

const testAccGatewaySetupTokenCreateConfigFormat = `
resource "oktapam_gateway_setup_token" "test-gateway-setup-token" {
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
