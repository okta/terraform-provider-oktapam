package oktapam

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/client"
)

func dataSourceGatewaySetupTokens() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGatewaySetupTokenRead,
		Schema: map[string]*schema.Schema{
			"description_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"gateway_setup_tokens": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"created_at": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"labels": {
							Type: schema.TypeMap,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGatewaySetupTokenRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	contains := d.Get("description_contains").(string)
	c := m.(client.OktaPAMClient)
	tokensList, err := c.ListGatewaySetupTokens(ctx, contains)
	if err != nil {
		return diag.FromErr(err)
	}
	tokens := make([]map[string]interface{}, len(tokensList))
	for idx, token := range tokensList {
		tokens[idx] = token.ToResourceMap()
	}
	if err := d.Set("gateway_setup_tokens", tokens); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return nil
}
