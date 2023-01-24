package oktapam

import (
	"context"
	"fmt"
	"testing"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/kylelemons/godebug/pretty"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
)

func TestAccGateway(t *testing.T) {
	resourceName1 := "oktapam_gateway_setup_token.test_gateway_setup_token"
	resourceName2 := "oktapam_gateway.test_gateway"

	identifier := randSeq()
	description := fmt.Sprintf("Acceptance Test Setup Token: %s", identifier)
	labels := constructLabels(10)

	setupToken := client.GatewaySetupToken{
		Description: &description,
		Details:     &client.GatewaySetupTokenLabelDetails{Labels: labels},
	}

	gateway := client.Gateway{
		Labels: labels,
	}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccGatewayCheckDestroy(identifier),
		Steps: []resource.TestStep{
			{
				Config: createTestAccGatewaySetupTokenCreateConfig(setupToken),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccGatewaySetupTokenCheckExists(resourceName1, setupToken),
					resource.TestCheckResourceAttr(
						resourceName1, attributes.Description, description,
					),
					testAccGatewayCheckExists(resourceName2, gateway),
					resource.TestCheckResourceAttr(
						resourceName2, attributes.Labels, description,
					),
				),
			},
			{
				ResourceName:      resourceName2,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccGatewayCheckExists(rn string, expectedGateway client.Gateway) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[rn]
		if !ok {
			return fmt.Errorf("resource not found: %s", rn)
		}

		resourceID := rs.Primary.ID

		c := testAccProvider.Meta().(client.OktaPAMClient)
		gateway, err := c.GetGateway(context.Background(), resourceID)
		if err != nil {
			return fmt.Errorf("error getting gateway setup token: %w", err)
		}
		if gateway == nil {
			return fmt.Errorf("gateway with id %s does not exist", resourceID)
		}

		if gateway.Description != expectedGateway.Description {
			return fmt.Errorf("expected description does not match returned description for gateway setup token.  expected: %s, got: %s", gateway.Description, expectedGateway.Description)
		}
		labelsCompare := pretty.Compare(expectedGateway, gateway)
		if labelsCompare != "" {
			return fmt.Errorf("expected labels does not match returned labels for gateway setup token:\n%s", labelsCompare)
		}

		return nil
	}
}

func testAccGatewayCheckDestroy(identifier string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		c := testAccProvider.Meta().(client.OktaPAMClient)
		params := client.ListGatewayParameters{
			Contains: identifier,
		}
		gateways, err := c.ListGateways(context.Background(), params)
		if err != nil {
			return fmt.Errorf("error getting tokens: %w", err)
		}

		if len(gateways) > 0 {
			return fmt.Errorf("gateway still exists")
		}

		return nil
	}
}
