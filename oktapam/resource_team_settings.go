package oktapam

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
)

func resourceTeamSettings() *schema.Resource {
	return &schema.Resource{
		Description:   descriptions.ResourceTeamSettings,
		CreateContext: resourceTeamSettingsCreate,
		ReadContext:   resourceTeamSettingsRead,
		UpdateContext: resourceTeamSettingsUpdate,
		DeleteContext: resourceTeamSettingsDelete,
		Schema: map[string]*schema.Schema{
			attributes.Team: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: descriptions.TeamName,
			},
			attributes.ReactivateUsersViaIDP: {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: descriptions.ReactivateUsersViaIDP,
			},
			attributes.ApproveDeviceWithoutInteraction: {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: descriptions.ApproveDeviceWithoutInteraction,
			},
			attributes.PostDeviceEnrollmentURL: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions.PostDeviceEnrollmentURL,
			},
			attributes.PostLoginURL: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions.PostLoginURL,
			},
			attributes.PostLogoutURL: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions.PostLogoutURL,
			},
			attributes.UserProvisioningExactUserName: {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: descriptions.UserProvisioningExactUserName,
			},
			attributes.ClientSessionDuration: {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: descriptions.ClientSessionDuration,
			},
			attributes.WebSessionDuration: {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: descriptions.WebSessionDuration,
			},
			attributes.IncludeUserSID: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions.IncludeUserSID,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceTeamSettingsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaPAMClient)

	settings := client.TeamSettings{
		ReactivateUsersViaIDP:           getBoolPtr(attributes.ReactivateUsersViaIDP, d, true),
		ApproveDeviceWithoutInteraction: getBoolPtr(attributes.ApproveDeviceWithoutInteraction, d, true),
		PostDeviceEnrollmentURL:         getStringPtr(attributes.PostDeviceEnrollmentURL, d, true),
		PostLogoutURL:                   getStringPtr(attributes.PostLogoutURL, d, true),
		PostLoginURL:                    getStringPtr(attributes.PostLoginURL, d, true),
		UserProvisioningExactUserName:   getBoolPtr(attributes.UserProvisioningExactUserName, d, true),
		ClientSessionDuration:           getIntPtr(attributes.ClientSessionDuration, d, true),
		WebSessionDuration:              getIntPtr(attributes.WebSessionDuration, d, true),
		IncludeUserSID:                  getStringPtr(attributes.IncludeUserSID, d, true),
	}

	err := c.UpdateTeamSettings(ctx, settings)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceTeamSettingsRead(ctx, d, m)
}

func resourceTeamSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(client.OktaPAMClient)

	settings, err := c.GetTeamSettings(ctx)
	if err != nil {
		return diag.FromErr(err)
	}

	if settings != nil {
		for key, value := range settings.ToResourceMap() {
			_ = d.Set(key, value)
		}
	}

	return diags
}

func resourceTeamSettingsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaPAMClient)

	if d.HasChanges(attributes.ReactivateUsersViaIDP, attributes.ApproveDeviceWithoutInteraction, attributes.PostDeviceEnrollmentURL, attributes.PostLoginURL,
		attributes.PostLogoutURL, attributes.UserProvisioningExactUserName, attributes.ClientSessionDuration, attributes.WebSessionDuration, attributes.IncludeUserSID) {

		//Build API Client Request Object
		teamSettingsRequest := client.TeamSettings{
			Team:                            getStringPtr(attributes.TeamName, d, false),
			ReactivateUsersViaIDP:           getBoolPtr(attributes.ReactivateUsersViaIDP, d, false),
			ApproveDeviceWithoutInteraction: getBoolPtr(attributes.ApproveDeviceWithoutInteraction, d, false),
			PostDeviceEnrollmentURL:         getStringPtr(attributes.PostDeviceEnrollmentURL, d, false),
			PostLogoutURL:                   getStringPtr(attributes.PostLogoutURL, d, false),
			PostLoginURL:                    getStringPtr(attributes.PostLoginURL, d, true),
			UserProvisioningExactUserName:   getBoolPtr(attributes.UserProvisioningExactUserName, d, false),
			ClientSessionDuration:           getIntPtr(attributes.ClientSessionDuration, d, false),
			WebSessionDuration:              getIntPtr(attributes.WebSessionDuration, d, false),
			IncludeUserSID:                  getStringPtr(attributes.IncludeUserSID, d, false),
		}

		err := c.UpdateTeamSettings(ctx, teamSettingsRequest)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceADConnectionRead(ctx, d, m)
}

func resourceTeamSettingsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//var diags diag.Diagnostics
	c := m.(client.OktaPAMClient)
	teamName := d.Get(attributes.TeamName).(string)

	err := c.DeleteTeamSettings(ctx, teamName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}
