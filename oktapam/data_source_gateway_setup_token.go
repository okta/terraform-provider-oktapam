package oktapam

import (
	"context"
	"strconv"
	"time"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
)

func dataSourceGatewaySetupTokens() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGatewaySetupTokenRead,
		Schema: map[string]*schema.Schema{
			// Query parameter values
			attributes.DescriptionContains: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions.FilterDescriptionContains,
			},
			// Return value
			attributes.GatewaySetupTokens: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: descriptions.SourceGatewaySetupTokens,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						attributes.ID: {
							Type:     schema.TypeString,
							Computed: true,
							// Description is autogenerated
						},
						attributes.CreatedAt: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: descriptions.CreatedAt,
						},
						attributes.Description: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: descriptions.Description,
						},
						attributes.Labels: {
							Type: schema.TypeMap,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Computed:    true,
							Description: descriptions.Labels,
						},
					},
				},
			},
		},
	}
}

func dataSourceGatewaySetupTokenRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	contains := d.Get(attributes.DescriptionContains).(string)
	c := m.(client.OktaPAMClient)
	tokensList, err := c.ListGatewaySetupTokens(ctx, contains)
	if err != nil {
		return diag.FromErr(err)
	}
	tokens := make([]map[string]interface{}, len(tokensList))
	for idx, token := range tokensList {
		tokens[idx] = token.ToResourceMap()
	}
	if err := d.Set(attributes.GatewaySetupTokens, tokens); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return nil
}
