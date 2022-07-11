package oktapam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
)

func dataSourceGatewaySetupTokens() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGatewaySetupTokensRead,
		Schema: map[string]*schema.Schema{
			attributes.DescriptionContains: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions.FilterDescriptionContains,
			},
			attributes.IDs: {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceGatewaySetupTokensRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaPAMClient)

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
	d.SetId(resource.UniqueId())
	return nil
}
