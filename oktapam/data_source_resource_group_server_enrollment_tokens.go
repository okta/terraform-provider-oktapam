package oktapam

import (
	"context"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceResourceGroupServerEnrollmentTokens() *schema.Resource {
	return &schema.Resource{
		Description: descriptions.SourceResourceGroupServerEnrollmentTokens,
		ReadContext: dataSourceResourceGroupServerEnrollmentTokenList,
		Schema: map[string]*schema.Schema{
			// Query parameter values
			attributes.ResourceGroup: {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions.ResourceGroupID,
			},
			attributes.Project: {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions.ProjectID,
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

func dataSourceResourceGroupServerEnrollmentTokenList(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := getLocalClientFromMetadata(m)
	resourceGroupID := d.Get(attributes.ResourceGroup).(string)
	if resourceGroupID == "" {
		return diag.Errorf("%s cannot be blank", attributes.ResourceGroup)
	}
	projectID := d.Get(attributes.Project).(string)
	if projectID == "" {
		return diag.Errorf("%s cannot be blank", attributes.Project)
	}
	tokensList, err := c.ListResourceGroupServerEnrollmentTokens(ctx, resourceGroupID, projectID)
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
	d.SetId(projectID)
	return nil
}
