package oktaasa

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-oktaasa/oktaasa/client"
	"github.com/terraform-providers/terraform-provider-oktaasa/oktaasa/logging"
)

func resourceGatewaySetupToken() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGatewaySetupTokenCreate,
		ReadContext:   resourceGatewaySetupTokenRead,
		DeleteContext: resourceGatewaySetupTokenDelete,
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
				Required: true,
				ForceNew: true,
			},
			"labels": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
				ForceNew: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceGatewaySetupTokenRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaASAClient)

	tokenID := d.Id()
	token, err := c.GetGatewaySetupToken(ctx, tokenID)
	if err != nil {
		return diag.FromErr(err)
	}

	if token == nil {
		logging.Debugf("setup token id was blank")
		d.SetId("")
		return nil
	}

	for key, value := range token.ToResourceMap() {
		logging.Debugf("setting %s to %v", key, value)
		d.Set(key, value)
	}

	return nil
}

func resourceGatewaySetupTokenCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaASAClient)

	description := getStringPtr("description", d, true)
	labels := d.Get("labels").(map[string]interface{})

	labelsMap := make(map[string]string, len(labels))

	for k, v := range labels {
		labelsMap[k] = fmt.Sprint(v)
	}

	tokenSpec := client.GatewaySetupToken{
		Description: description,
		Details:     &client.GatewaySetupTokenLabelDetails{Labels: labelsMap},
	}
	if createdToken, err := c.CreateGatewaySetupToken(ctx, tokenSpec); err != nil {
		return diag.FromErr(err)
	} else if createdToken == nil {
		d.SetId("")
	} else {
		d.SetId(*createdToken.ID)
	}

	return resourceGatewaySetupTokenRead(ctx, d, m)
}

func resourceGatewaySetupTokenDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaASAClient)

	id := d.Id()
	err := c.DeleteGatewaySetupToken(ctx, id)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
