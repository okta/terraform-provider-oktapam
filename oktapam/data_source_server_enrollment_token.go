package oktapam

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/client"
)

func dataSourceServerEnrollmentTokens() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceServerEnrollmentTokenRead,
		Schema: map[string]*schema.Schema{
			"project_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"server_enrollment_tokens": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"project_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"token": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"issued_at": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"created_by_user": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceServerEnrollmentTokenRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaPAMClient)
	project := d.Get("project_name").(string)
	if project == "" {
		return diag.Errorf("project_name cannot be blank")
	}
	tokensList, err := c.ListServerEnrollmentTokens(ctx, project)
	if err != nil {
		return diag.FromErr(err)
	}

	tokens := make([]map[string]interface{}, len(tokensList))
	for idx, token := range tokensList {
		tokens[idx] = token.ToResourceMap()
	}

	if err := d.Set("server_enrollment_tokens", tokens); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return nil
}
