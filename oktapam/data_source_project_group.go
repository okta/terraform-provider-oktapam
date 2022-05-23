package oktapam

import (
	"context"
	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/constants/descriptions"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/client"
)

func dataSourceProjectGroups() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceProjectGroupsRead,
		Schema: map[string]*schema.Schema{
			// Query parameter values
			attributes.ProjectName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions.FilterProjectName,
			},
			attributes.IncludeRemoved: {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: descriptions.FilterIncludeRemoved,
			},
			attributes.CreateServerGroup: {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: descriptions.FilterCreateServerGroup,
			},
			attributes.HasSelectors: {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: descriptions.FilterHasSelectors,
			},
			attributes.HasNoSelectors: {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: descriptions.FilterHasNoSelectors,
			},
			attributes.OfflineEnabled: { // TODO: update this after fixing endpoint
				Type:        schema.TypeBool,
				Optional:    true,
				Description: descriptions.FilterOfflineEnabled,
			},
			// Return value
			attributes.ProjectGroups: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: descriptions.SourceProjectGroups,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						attributes.ProjectName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: descriptions.ProjectName,
						},
						attributes.GroupName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: descriptions.GroupName,
						},
						attributes.GroupID: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: descriptions.GroupID,
						},
						attributes.DeletedAt: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: descriptions.DeletedAt,
						},
						attributes.RemovedAt: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: descriptions.RemovedAt,
						},
						attributes.CreateServerGroup: {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: descriptions.CreateServerGroup,
						},
						attributes.ServerAccess: {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: descriptions.ServerAccess,
						},
						attributes.ServerAdmin: {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: descriptions.ServerAdmin,
						},
						attributes.ServersSelector: {
							Type: schema.TypeMap,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Computed:    true,
							Description: descriptions.ServersSelector,
						},
					},
				},
			},
		},
	}
}

func dataSourceProjectGroupsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaPAMClient)
	project := d.Get(attributes.ProjectName).(string)
	if project == "" {
		return diag.Errorf("%s cannot be blank", attributes.ProjectName)
	}
	parameters := client.ListProjectGroupsParameters{}
	includeRemoved, err := getOkBool(attributes.IncludeRemoved, d)
	if err != nil {
		return diag.FromErr(err)
	}
	parameters.IncludeRemoved = includeRemoved

	createServerGroup, err := getOkBool(attributes.CreateServerGroup, d)
	if err != nil {
		return diag.FromErr(err)
	}
	parameters.CreateServerGroup = createServerGroup

	hasSelectors, err := getOkBool(attributes.HasSelectors, d)
	if err != nil {
		return diag.FromErr(err)
	}
	parameters.HasSelectors = hasSelectors

	hasNoSelectors, err := getOkBool(attributes.HasNoSelectors, d)
	if err != nil {
		return diag.FromErr(err)
	}
	parameters.HasNoSelectors = hasNoSelectors

	offlineEnabled, err := getOkBool(attributes.OfflineEnabled, d)
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

	if err := d.Set(attributes.ProjectGroups, assignments); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return nil
}
