package oktapam

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/client"
)

func dataSourceGroups() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGroupsRead,
		Schema: map[string]*schema.Schema{
			"contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"include_deleted": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"only_include_deleted": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"disconnected_mode_on_only": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"deleted_at": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"roles": {
							Type:     schema.TypeList,
							Elem:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGroupsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(client.OktaPAMClient)
	parameters := client.ListGroupsParameters{}

	if contains, ok := d.GetOk("contains"); ok {
		parameters.Contains = contains.(string)
	}

	includeDeleted, err := getOkBool("include_deleted", d)
	if err != nil {
		return diag.FromErr(err)
	}
	parameters.IncludeDeleted = includeDeleted

	onlyIncludeDeleted, err := getOkBool("only_include_deleted", d)
	if err != nil {
		return diag.FromErr(err)
	}
	parameters.OnlyIncludeDeleted = onlyIncludeDeleted

	disconnectedModeOnOnly, err := getOkBool("disconnected_mode_on_only", d)
	if err != nil {
		return diag.FromErr(err)
	}
	parameters.DisconnectedModeOnOnly = disconnectedModeOnOnly

	groupsList, err := c.ListGroups(ctx, parameters)
	if err != nil {
		return diag.FromErr(err)
	}
	groups := make([]map[string]interface{}, len(groupsList))
	for idx, proj := range groupsList {
		groups[idx] = proj.ToResourceMap()
	}

	if err := d.Set("groups", groups); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}
