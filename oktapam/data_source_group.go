package oktapam

import (
	"context"
	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/constants"
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
			constants.AttrContains:               constants.AttributeSchemas[constants.AttrContains],
			constants.AttrIncludeDeleted:         constants.AttributeSchemas[constants.AttrIncludeDeleted],
			constants.AttrOnlyIncludeDeleted:     constants.AttributeSchemas[constants.AttrOnlyIncludeDeleted],
			constants.AttrDisconnectedModeOnOnly: constants.AttributeSchemas[constants.AttrDisconnectedModeOnOnly],
			constants.ResGroups:                  constants.ResourceSchemas[constants.ResGroups],
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
