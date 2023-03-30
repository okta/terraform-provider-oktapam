package oktapam

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/typed_strings"
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
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
				Description: descriptions.PostDeviceEnrollmentURL,
			},
			attributes.PostLoginURL: {
				Type:        schema.TypeString,
				Optional:    true,
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
				Description: descriptions.PostLoginURL,
			},
			attributes.PostLogoutURL: {
				Type:        schema.TypeString,
				Optional:    true,
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
				Description: descriptions.PostLogoutURL,
			},
			attributes.UserProvisioningExactUserName: {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: descriptions.UserProvisioningExactUserName,
			},
			attributes.ClientSessionDuration: {
				Type:        schema.TypeInt,
				Computed: true,
				Optional:    true,
				ValidateFunc: validation.IntBetween(60*60, 25*60*60),
				Description: descriptions.ClientSessionDuration,
			},
			attributes.WebSessionDuration: {
				Type:        schema.TypeInt,
				Computed: true,
				Optional:    true,
				ValidateFunc: validation.IntBetween(30*60, 25*60*60),
				Description: descriptions.WebSessionDuration,
			},
			attributes.IncludeUserSID: {
				Type:        schema.TypeString,
				Computed: true,
				Optional:    true,
				ValidateFunc: validation.StringInSlice(
					[]string{
						typed_strings.IncludeUserSIDNever.String(),
						typed_strings.IncludeUserSIDAlways.String(),
						typed_strings.IncludeUserSIDIfAvailable.String(),
					}, false),
				Description: descriptions.IncludeUserSID,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceTeamSettingsCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := m.(client.OktaPAMClient)

	settings := client.TeamSettings{
		ReactivateUsersViaIDP:           getBoolPtr(attributes.ReactivateUsersViaIDP, d, false),
		ApproveDeviceWithoutInteraction: getBoolPtr(attributes.ApproveDeviceWithoutInteraction, d, false),
		PostDeviceEnrollmentURL:         getStringPtr(attributes.PostDeviceEnrollmentURL, d, false),
		PostLogoutURL:                   getStringPtr(attributes.PostLogoutURL, d, false),
		PostLoginURL:                    getStringPtr(attributes.PostLoginURL, d, false),
		UserProvisioningExactUserName:   getBoolPtr(attributes.UserProvisioningExactUserName, d, false),
		ClientSessionDuration:           getIntPtr(attributes.ClientSessionDuration, d, false),
		WebSessionDuration:              getIntPtr(attributes.WebSessionDuration, d, false),
		IncludeUserSID:                  getStringPtr(attributes.IncludeUserSID, d, false),
	}
	if err := c.CreateTeamSettings(ctx, settings); err!=nil {
		return diag.FromErr(err)
	}
	d.SetId(c.Team)

	return resourceTeamSettingsRead(ctx, d, m)
}

func resourceTeamSettingsRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
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

	return nil
}

func resourceTeamSettingsUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := m.(client.OktaPAMClient)

	changed := false
	updates := make(map[string]any)

	changeableAttributes := []string{
		attributes.ReactivateUsersViaIDP,
		attributes.ApproveDeviceWithoutInteraction,
		attributes.PostDeviceEnrollmentURL,
		attributes.PostLoginURL,
		attributes.PostLogoutURL,
		attributes.UserProvisioningExactUserName,
		attributes.ClientSessionDuration,
		attributes.WebSessionDuration,
		attributes.IncludeUserSID,
	}

	for _, attribute := range changeableAttributes {
		// Check if the value is nil
		if d.HasChange(attribute) {
			updates[attribute] = d.Get(attribute)
			changed = true
		}
	}

	if changed {
		err := c.UpdateTeamSettings(ctx, updates)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceTeamSettingsRead(ctx, d, m)
}

func resourceTeamSettingsDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	var diags diag.Diagnostics
	diag := diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Can't delete team settings resource.",
		Detail: fmt.Sprintf("Team settings resource does not support delete operation. If you don't want to " +
			"manage team settings via terraform, then remove it from the terraform state manually using " +
			"\"terraform state rm\" command."),
	}
	return append(diags,diag)
}
