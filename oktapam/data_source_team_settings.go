package oktapam

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
)

func dataSourceTeamSettings() *schema.Resource {
	return &schema.Resource{
		Description: descriptions.SourceTeamSettings,
		ReadContext: dataSourceTeamSettingsFetch,
		Schema: map[string]*schema.Schema{
			attributes.Team: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: descriptions.TeamName,
			},
		},
	}
}

func dataSourceTeamSettingsFetch(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaPAMClient)
	name := d.Get(attributes.Name).(string)
	if name == "" {
		return diag.Errorf("%s cannot be blank", attributes.Name)
	}

	settings, err := c.GetTeamSettings(ctx)
	if err != nil {
		return diag.FromErr(err)
	}

	if settings != nil {
		for key, value := range settings.ToResourceMap() {
			d.Set(key, value)
		}
	}
	//TODO(lehar) : should I return an error when team settings does not exist?
	return nil
}
