package oktapam

import (
	"context"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
)

func dataSourceServerEnrollmentTokens() *schema.Resource {
	return &schema.Resource{
		Description: descriptions.SourceServerEnrollmentTokens,
		ReadContext: dataSourceServerEnrollmentTokenList,
		Schema: map[string]*schema.Schema{
			// Query parameter values
			attributes.ProjectName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions.FilterProjectName,
			},
			// Return value
			attributes.IDs: {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceServerEnrollmentTokenList(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := m.(client.OktaPAMClient)
	project := d.Get(attributes.ProjectName).(string)
	if project == "" {
		return diag.Errorf("%s cannot be blank", attributes.ProjectName)
	}
	tokensList, err := c.ListServerEnrollmentTokens(ctx, project)
	if err != nil {
		return diag.FromErr(err)
	}

	ids := make([]string, len(tokensList))
	for i, token := range tokensList {
		ids[i] = *token.ID
	}

	if err := d.Set(attributes.IDs, ids); err != nil {
		return diag.FromErr(err)
	}

	// Server enrollment tokens correspond to a Project
	d.SetId(project)
	return nil
}
