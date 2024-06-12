package oktapam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
)

func dataSourceSudoCommandsBundles() *schema.Resource {
	return &schema.Resource{
		Description: descriptions.SourceSudoCommandsBundles,
		ReadContext: dataSourceSudoCommandsBundlesList,
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

func dataSourceSudoCommandsBundlesList(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	var diags diag.Diagnostics
	c := getSDKClientFromMetadata(m)

	resp, err := client.ListSudoCommandsBundles(ctx, c)
	if err != nil {
		return diag.FromErr(err)
	}

	sudoCommandsBundles := make([]string, 0, len(resp.GetList()))
	for _, sudoCommandsBundle := range resp.GetList() {
		sudoCommandsBundles = append(sudoCommandsBundles, *sudoCommandsBundle.Id)
	}

	if err = d.Set(attributes.IDs, sudoCommandsBundles); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	d.SetId(c.Team)

	return diags
}
