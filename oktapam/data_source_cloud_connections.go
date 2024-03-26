package oktapam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
)

func dataSourceCloudConnections() *schema.Resource {
	return &schema.Resource{
		Description: descriptions.SourceCloudConnections,
		ReadContext: dataSourceCloudConnectionsList,
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

func dataSourceCloudConnectionsList(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	var diags diag.Diagnostics
	c := getLocalClientFromMetadata(m)

	var nameFilter string
	if f, ok := d.GetOk(attributes.Name); ok {
		nameFilter = f.(string)
	}

	cloudConnectionsList, err := c.ListCloudConnections(ctx)
	if err != nil {
		return diag.FromErr(err)
	}

	cloudConnections := make([]string, 0, len(cloudConnectionsList))
	for _, cloudConnection := range cloudConnectionsList {
		if nameFilter == "" || *cloudConnection.Name == nameFilter {
			cloudConnections = append(cloudConnections, *cloudConnection.ID)
		}
	}

	if err = d.Set(attributes.IDs, cloudConnections); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	d.SetId(c.Team)

	return diags
}
