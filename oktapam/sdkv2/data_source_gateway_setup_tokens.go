package sdkv2

import (
	"context"

	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGatewaySetupTokens() *schema.Resource {
	return &schema.Resource{
		Description: descriptions.SourceGatewaySetupTokens,
		ReadContext: dataSourceGatewaySetupTokensRead,
		Schema: map[string]*schema.Schema{
			// Query filter
			attributes.DescriptionContains: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions.FilterDescriptionContains,
			},
			// Return values
			attributes.IDs: {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceGatewaySetupTokensRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := client.GetLocalClientFromMetadata(m)

	contains := d.Get(attributes.DescriptionContains).(string)
	tokens, err := c.ListGatewaySetupTokens(ctx, contains)
	if err != nil {
		return diag.FromErr(err)
	}

	ids := make([]string, len(tokens))
	for i, token := range tokens {
		ids[i] = *token.ID
	}

	if err := d.Set(attributes.IDs, ids); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(c.Team)
	return nil
}
