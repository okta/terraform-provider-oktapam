package oktapam

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"
)

func resourceTeamSettings() *schema.Resource {
	return &schema.Resource{
		Description:   descriptions.ResourceTeamSettings,
		CreateContext: resourceTeamSettingsCreate,
		ReadContext:   resourceTeamSettingsRead,
		UpdateContext: resourceTeamSettingsUpdate,
		DeleteContext: resourceTeamSettingsDelete,
		Schema: map[string]*schema.Schema{
			attributes.ID: {
				Type:     schema.TypeString,
				Computed: true,
				Description:descriptions.TeamSettingsID,
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
	d.SetId(c.Team)

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
			err = d.Set(key, value)
			if err != nil {
				return diag.FromErr(err)
			}
		}
	} else {
		return diag.Errorf("Team settings does not exist for the team %s", c.Team)
	}

	return diags
}

func resourceTeamSettingsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaPAMClient)

	if d.HasChanges(attributes.ReactivateUsersViaIDP, attributes.ApproveDeviceWithoutInteraction, attributes.PostDeviceEnrollmentURL, attributes.PostLoginURL,
		attributes.PostLogoutURL, attributes.UserProvisioningExactUserName, attributes.ClientSessionDuration, attributes.WebSessionDuration, attributes.IncludeUserSID) {

		//Build API Client Request Object
		teamSettingsRequest := client.TeamSettings{
			Team:                            getStringPtr(c.Team, d, false),
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
		d.SetId(c.Team)
	}

	return resourceADConnectionRead(ctx, d, m)
}

func resourceTeamSettingsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(client.OktaPAMClient)
	teamSettingsRequest := &client.TeamSettings{
		Team: getStringPtr(c.Team, d, false),
	}
	const defaultValAlways string = "Always"
	const defaultValIf_Available = "If_Available"

	//Get the remote team settings
	settings, err := c.GetTeamSettings(ctx)

	// Get the terraform managed team settings

	// Compare if the attributes are same
	// For the same attributes, set the default value
	// For different attributes, leave as is
	oldVal, newVal := d.GetChange(attributes.IncludeUserSID)
	logging.Debugf("includeuserSID old val: %s ; includeuserSID new val: %s",oldVal, newVal)
	if oldVal==settings.IncludeUserSID {
		teamSettingsRequest.IncludeUserSID = utils.AsStringPtr(defaultValAlways)
	}
	teamSettingsRequest.IncludeUserSID =  utils.AsStringPtr(defaultValIf_Available)
	//approveDeviceWithoutInteractionHasChange := d.HasChange(attributes.ApproveDeviceWithoutInteraction)
	//postDeviceEnrollmentURLHasChange := d.HasChange(attributes.PostDeviceEnrollmentURL)
	//postLogoutURLHasChange := d.HasChange(attributes.PostLogoutURL)
	//postLoginURLHasChange := d.HasChange(attributes.PostLoginURL)
	//userProvisioningExactUserNameHasChange := d.HasChange(attributes.UserProvisioningExactUserName)
	//clientSessionDurationHasChange := d.HasChange(attributes.ClientSessionDuration)
	//webSessionDurationHasChange := d.HasChange(attributes.WebSessionDuration)
	//includeUserSIDHasChange := d.HasChange(attributes.IncludeUserSID)
	err = c.UpdateTeamSettings(ctx,*teamSettingsRequest)
	err = c.DeleteTeamSettings(ctx)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return diags
}
