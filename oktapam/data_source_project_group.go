package oktapam

import (
	"context"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
)

func dataSourceProjectGroup() *schema.Resource {
	return &schema.Resource{
		Description: descriptions.SourceProjectGroups,
		ReadContext: dataSourceProjectGroupsFetch,
		Schema: map[string]*schema.Schema{
			attributes.ProjectName: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: descriptions.ProjectName,
			},
			attributes.GroupName: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
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
	}
}

func dataSourceProjectGroupsFetch(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaPAMClient)

	project, group, err := parseProjectGroupResourceID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	tokenID := d.Get(attributes.ID).(string)
	if tokenID == "" {
		return diag.Errorf("%s cannot be blank", attributes.ID)
	}

	project, err := c.GetProjectGroup(ctx, project, group)
	if err != nil {
		return diag.FromErr(err)
	}

	if token != nil {
		d.SetId(*token.ID)
		for key, value := range token.ToResourceMap() {
			d.Set(key, value)
		}
	} else {
		logging.Infof("gateway setup token %s does not exist", tokenID)
	}

	return nil
}
