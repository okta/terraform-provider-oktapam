package oktapam

import (
	"context"
	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/constants"
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
			constants.AttrDescriptionContains: constants.AttributeSchemas[constants.AttrDescriptionContains],
			constants.ResGatewaySetupTokens:   constants.ResourceSchemas[constants.ResGatewaySetupTokens],
		},
	}
}

func dataSourceGatewaySetupTokenRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	contains := d.Get(constants.AttrDescriptionContains).(string)
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
