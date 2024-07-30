package sdkv2

import (
	"context"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
)

func dataSourceGroups() *schema.Resource {
	return &schema.Resource{
		Description: descriptions.SourceGroups,
		ReadContext: dataSourceGroupList,
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
			// Return values
			attributes.Names: {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceGroupList(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	var diags diag.Diagnostics

	c := client.GetLocalClientFromMetadata(m)
	parameters := client.ListGroupsParameters{}

	if contains, ok := d.GetOk(attributes.Contains); ok {
		parameters.Contains = contains.(string)
	}

	includeDeleted, err := GetOkBoolFromResource(attributes.IncludeDeleted, d)
	if err != nil {
		return diag.FromErr(err)
	}
	parameters.IncludeDeleted = includeDeleted

	onlyIncludeDeleted, err := GetOkBoolFromResource(attributes.OnlyIncludeDeleted, d)
	if err != nil {
		return diag.FromErr(err)
	}
	parameters.OnlyIncludeDeleted = onlyIncludeDeleted

	disconnectedModeOnOnly, err := GetOkBoolFromResource(attributes.DisconnectedModeOnOnly, d)
	if err != nil {
		return diag.FromErr(err)
	}
	parameters.DisconnectedModeOnOnly = disconnectedModeOnOnly

	groupsList, err := c.ListGroups(ctx, parameters)
	if err != nil {
		return diag.FromErr(err)
	}
	names := make([]string, len(groupsList))
	for idx, proj := range groupsList {
		names[idx] = *proj.Name
	}

	if err := d.Set(attributes.Names, names); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(c.Team)
	return diags
}
