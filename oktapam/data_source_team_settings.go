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
			attributes.ID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions.TeamSettingsID,
			},
			attributes.ReactivateUsersViaIDP: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: descriptions.ReactivateUsersViaIDP,
			},
			attributes.ApproveDeviceWithoutInteraction: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: descriptions.ApproveDeviceWithoutInteraction,
			},
			attributes.UserProvisioningExactUserName: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: descriptions.UserProvisioningExactUserName,
			},
			attributes.ClientSessionDuration: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: descriptions.ClientSessionDuration,
			},
			attributes.WebSessionDuration: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: descriptions.WebSessionDuration,
			},
			attributes.IncludeUserSID: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: descriptions.IncludeUserSID,
			},
		},
	}
}

func dataSourceTeamSettingsFetch(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(client.OktaPAMClient)
	name := d.Get(attributes.ID).(string)
	if name == "" {
		return diag.Errorf("%s cannot be blank", attributes.ID)
	}

	settings, err := c.GetTeamSettings(ctx)
	if err != nil {
		return diag.FromErr(err)
	}

	if settings != nil {
		for key, value := range settings.ToResourceMap() {
			if err := d.Set(key, value); err != nil {
				diags = append(diags, diag.FromErr(err)...)
			}
		}
		d.SetId(name)
	} else {
		return diag.Errorf("Team settings does not exist for the team %s", c.Team)
	}
	return diags
}
