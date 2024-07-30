package oktapam

import (
	"context"
	"fmt"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceResourceGroupProjects() *schema.Resource {
	return &schema.Resource{
		Description: descriptions.SourceResourceGroupProjects,
		ReadContext: dataSourceResourceGroupProjectsList,
		Schema: map[string]*schema.Schema{
			// Query parameter values
			attributes.ResourceGroup: {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions.ResourceGroupID,
			},
			attributes.Name: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions.FilterName,
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

func dataSourceResourceGroupProjectsList(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	var diags diag.Diagnostics

	c := getLocalClientFromMetadata(m)

	resourceGroupIDI, ok := d.GetOk(attributes.ResourceGroup)
	if !ok {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("missing required %s parameter", attributes.ResourceGroup),
		})
	}

	resourceGroupID := resourceGroupIDI.(string)

	var nameFilter string
	if f, ok := d.GetOk(attributes.Name); ok {
		nameFilter = f.(string)
	}

	projectsList, err := c.ListResourceGroupProjects(ctx, resourceGroupID)
	if err != nil {
		return diag.FromErr(err)
	}

	ids := make([]string, 0, len(projectsList))
	for _, proj := range projectsList {
		if nameFilter == "" || nameFilter == *proj.Name {
			ids = append(ids, *proj.ID)
		}
	}

	if err := d.Set(attributes.IDs, ids); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(c.Team)
	return diags
}
