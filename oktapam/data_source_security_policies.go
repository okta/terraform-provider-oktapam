package oktapam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
)

func dataSourceSecurityPolicies() *schema.Resource {
	return &schema.Resource{
		Description: descriptions.SourceSecurityPolicies,
		ReadContext: dataSourceSecurityPoliciesList,
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

func dataSourceSecurityPoliciesList(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	var diags diag.Diagnostics
	c := getLocalClientFromMetadata(m)

	var nameFilter string
	if f, ok := d.GetOk(attributes.Name); ok {
		nameFilter = f.(string)
	}

	securityPoliciesList, err := c.ListSecurityPolicies(ctx)
	if err != nil {
		return diag.FromErr(err)
	}

	securityPolicies := make([]string, 0, len(securityPoliciesList))
	for _, rg := range securityPoliciesList {
		if nameFilter == "" || *rg.Name == nameFilter {
			securityPolicies = append(securityPolicies, *rg.ID)
		}
	}

	if err = d.Set(attributes.IDs, securityPolicies); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	d.SetId(c.Team)

	return diags
}
