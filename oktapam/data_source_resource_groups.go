package oktapam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
)

func dataSourceResourceGroups() *schema.Resource {
	return &schema.Resource{
		Description: descriptions.SourceResourceGroups,
		ReadContext: dataSourceResourceGroupList,
		Schema: map[string]*schema.Schema{
			attributes.Name: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions.FilterName,
			},

			attributes.IDs: {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceResourceGroupList(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(client.OktaPAMClient)

	var nameFilter string
	if f, ok := d.GetOk(attributes.Name); ok {
		nameFilter = f.(string)
	}

	resourceGroupsList, err := c.ListResourceGroups(ctx)
	if err != nil {
		return diag.FromErr(err)
	}

	resourceGroups := make([]string, 0, len(resourceGroupsList))
	for _, rg := range resourceGroupsList {
		if nameFilter == "" || *rg.Name == nameFilter {
			resourceGroups = append(resourceGroups, *rg.ID)
		}
	}

	if err = d.Set(attributes.IDs, resourceGroups); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	d.SetId(c.Team)

	return diags
}
