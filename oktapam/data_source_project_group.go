package oktapam

import (
	"context"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceProjectGroup() *schema.Resource {
	return &schema.Resource{
		Description: descriptions.SourceProjectGroup,
		ReadContext: dataSourceProjectGroupFetch,
		Schema: map[string]*schema.Schema{
			attributes.ProjectName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions.ProjectName,
			},
			attributes.GroupName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions.GroupName,
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
	}
}

func dataSourceProjectGroupFetch(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := getLocalClientFromMetadata(m)

	projectName := d.Get(attributes.ProjectName).(string)
	groupName := d.Get(attributes.GroupName).(string)

	projectGroup, err := c.GetProjectGroup(ctx, projectName, groupName)
	if err != nil {
		return diag.FromErr(err)
	}

	if projectGroup != nil {
		d.SetId(*projectGroup.ID)
		resourceMap, err := projectGroup.ToResourceMap()
		if err != nil {
			return diag.FromErr(err)
		}

		for key, value := range resourceMap {
			if err := d.Set(key, value); err != nil {
				return diag.FromErr(err)
			}
		}
	} else {
		logging.Infof("project group belonging to project %s and group %s does not exist", *projectGroup.Project, *projectGroup.Group)
	}

	return nil
}
