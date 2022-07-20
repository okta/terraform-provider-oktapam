package oktapam

import (
	"context"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
)

func dataSourceProjectGroups() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceProjectGroupList,
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
			attributes.DisconnectedModeOnOnly: {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: descriptions.FilterDisconnectedModeOnOnly,
			},
			// Return values
			attributes.GroupNames: {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceProjectGroupList(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

	disconnectedModeOnOnly, err := getOkBool(attributes.DisconnectedModeOnOnly, d)
	if err != nil {
		return diag.FromErr(err)
	}
	parameters.DisconnectedModeOnOnly = disconnectedModeOnOnly

	projectGroups, err := c.ListProjectGroups(ctx, project, parameters)
	if err != nil {
		return diag.FromErr(err)
	}

	groupNames := make([]string, len(projectGroups))
	for idx, projectGroup := range projectGroups {
		groupNames[idx] = *projectGroup.Group
	}

	if err := d.Set(attributes.GroupNames, groupNames); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resource.UniqueId())
	return nil
}
