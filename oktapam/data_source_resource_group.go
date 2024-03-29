package oktapam

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
)

func dataSourceResourceGroup() *schema.Resource {
	return &schema.Resource{
		Description: descriptions.SourceResourceGroup,
		ReadContext: dataSourceResourceGroupFetch,
		Schema: map[string]*schema.Schema{
			attributes.ID: {
				Type:     schema.TypeString,
				Required: true,
				// Description is autogenerated
			},
			attributes.Name: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: descriptions.Name,
			},
			attributes.Description: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: descriptions.Description,
			},
			attributes.DelegatedResourceAdminGroups: {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed:    true,
				Description: descriptions.DelegatedAdminGroups,
			},
		},
	}
}

func dataSourceResourceGroupFetch(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	var diags diag.Diagnostics
	c := getLocalClientFromMetadata(m)

	idI, ok := d.GetOk(attributes.ID)
	if !ok {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("missing required %s parameter", attributes.ID),
		})
	}

	id := idI.(string)

	rg, err := c.GetResourceGroup(ctx, id)
	if err != nil {
		return diag.FromErr(err)
	}

	if rg != nil {
		d.SetId(*rg.ID)
		for key, value := range rg.ToResourceMap() {
			if err := d.Set(key, value); err != nil {
				diags = append(diags, diag.FromErr(err)...)
			}
		}
	} else {
		logging.Infof("resource group with id %s does not exist", id)
	}

	return diags
}
