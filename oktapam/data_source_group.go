package oktapam

import (
	"context"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
)

func dataSourceGroups() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGroupsRead,
		Schema: map[string]*schema.Schema{
			// Query parameter values
			attributes.Contains: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions.FilterContains,
			},
			attributes.IncludeDeleted: {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: descriptions.FilterIncludeDeleted,
			},
			attributes.OnlyIncludeDeleted: {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: descriptions.FilterOnlyIncludeDeleted,
			},
			attributes.DisconnectedModeOnOnly: {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: descriptions.FilterDisconnectedModeOnOnly,
			},
			// Return value
			attributes.Groups: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: descriptions.SourceGroups,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						attributes.Name: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: descriptions.Name,
						},
						attributes.ID: {
							Type:     schema.TypeString,
							Computed: true,
							// Description is autogenerated
						},
						attributes.DeletedAt: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: descriptions.DeletedAt,
						},
						attributes.Roles: {
							Type:        schema.TypeList,
							Elem:        schema.TypeString,
							Computed:    true,
							Description: descriptions.Roles,
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

	if contains, ok := d.GetOk(attributes.Contains); ok {
		parameters.Contains = contains.(string)
	}

	includeDeleted, err := getOkBool(attributes.IncludeDeleted, d)
	if err != nil {
		return diag.FromErr(err)
	}
	parameters.IncludeDeleted = includeDeleted

	onlyIncludeDeleted, err := getOkBool(attributes.OnlyIncludeDeleted, d)
	if err != nil {
		return diag.FromErr(err)
	}
	parameters.OnlyIncludeDeleted = onlyIncludeDeleted

	disconnectedModeOnOnly, err := getOkBool(attributes.DisconnectedModeOnOnly, d)
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

	if err := d.Set(attributes.Groups, groups); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}
