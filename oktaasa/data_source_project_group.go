package oktaasa

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-oktaasa/oktaasa/client"
)

func dataSourceProjectGroups() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceProjectGroupsRead,
		Schema: map[string]*schema.Schema{
			"project_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"include_removed": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"create_server_group": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"has_selectors": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"has_no_selectors": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"offline_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"project_groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"project_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"group_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"group_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"deleted_at": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"removed_at": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"create_server_group": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"server_access": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"server_admin": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"servers_selector": {
							Type: schema.TypeMap,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceProjectGroupsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaASAClient)
	project := d.Get("project_name").(string)
	if project == "" {
		return diag.Errorf("project_name cannot be blank")
	}
	parameters := client.ListProjectGroupsParameters{}
	includeRemoved, err := getOkBool("include_removed", d)
	if err != nil {
		return diag.FromErr(err)
	}
	parameters.IncludeRemoved = includeRemoved

	createServerGroup, err := getOkBool("create_server_group", d)
	if err != nil {
		return diag.FromErr(err)
	}
	parameters.CreateServerGroup = createServerGroup

	hasSelectors, err := getOkBool("has_selectors", d)
	if err != nil {
		return diag.FromErr(err)
	}
	parameters.HasSelectors = hasSelectors

	hasNoSelectors, err := getOkBool("has_no_selectors", d)
	if err != nil {
		return diag.FromErr(err)
	}
	parameters.HasNoSelectors = hasNoSelectors

	offlineEnabled, err := getOkBool("offline_enabled", d)
	if err != nil {
		return diag.FromErr(err)
	}
	parameters.OfflineEnabled = offlineEnabled

	assignmentsList, err := c.ListProjectGroups(ctx, project, parameters)
	if err != nil {
		return diag.FromErr(err)
	}

	assignments := make([]map[string]interface{}, len(assignmentsList))
	for idx, assignment := range assignmentsList {
		m, err := assignment.ToResourceMap()
		if err != nil {
			return diag.FromErr(err)
		}
		assignments[idx] = m
	}

	if err := d.Set("project_groups", assignments); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return nil
}
