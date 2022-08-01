package oktapam

import (
	"context"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
)

func dataSourceProjects() *schema.Resource {
	return &schema.Resource{
		Description: descriptions.SourceProjects,
		ReadContext: dataSourceProjectList,
		Schema: map[string]*schema.Schema{
			// Query parameter values
			attributes.Self: {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: descriptions.FilterSelf,
			},
			attributes.Contains: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions.FilterContains,
			},
			// Return values
			attributes.Names: {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceProjectList(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(client.OktaPAMClient)
	parameters := client.ListProjectsParameters{}
	if contains, ok := d.GetOk(attributes.Contains); ok {
		parameters.Contains = contains.(string)
	}

	self, err := getOkBool(attributes.Self, d)
	if err != nil {
		return diag.FromErr(err)
	}
	parameters.Self = self

	projectsList, err := c.ListProjects(ctx, parameters)
	if err != nil {
		return diag.FromErr(err)
	}

	names := make([]string, len(projectsList))
	for idx, proj := range projectsList {
		names[idx] = *proj.Name
	}

	if err := d.Set(attributes.Names, names); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(c.Team)
	return diags
}
