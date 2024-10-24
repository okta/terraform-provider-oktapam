package oktapam

import (
	"context"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceADUserSyncTaskSettings() *schema.Resource {
	return &schema.Resource{
		Description: descriptions.SourceADUserSyncTaskSettings,
		ReadContext: dataSourceADUserSyncTaskSettingsRead,
		Schema: map[string]*schema.Schema{
			attributes.ADConnectionID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions.ADConnectionID,
			},
			attributes.Name: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: descriptions.Name,
			},
			attributes.ID: {
				Type:     schema.TypeString,
				Required: true,
			},
			attributes.Frequency: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: descriptions.ADUserSyncTaskSettingsFrequency,
			},
			attributes.StartHourUTC: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: descriptions.ADUserSyncTaskSettingsStartHourUTC,
			},
			attributes.BaseDN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: descriptions.ADUserSyncTaskSettingsBaseDN,
			},
			attributes.LDAPQueryFilter: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: descriptions.ADUserSyncTaskSettingsLDAPQueryFilter,
			},
			attributes.UPNField: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: descriptions.ADUserSyncTaskSettingsUPNField,
			},
			attributes.SIDField: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: descriptions.ADUserSyncTaskSettingsSIDField,
			},
			attributes.IsActive: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: descriptions.ADUserSyncTaskSettingsIsActive,
			},
			attributes.RunTest: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: descriptions.ADUserSyncTaskSettingsRunTest,
			},
		},
	}
}

func dataSourceADUserSyncTaskSettingsRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := getLocalClientFromMetadata(m)
	adUserSyncTaskSettingsId := d.Get(attributes.ID).(string)
	if adUserSyncTaskSettingsId == "" {
		return diag.Errorf("%s cannot be blank", attributes.ID)
	}
	adConnectionID := d.Get(attributes.ADConnectionID).(string)
	if adConnectionID == "" {
		return diag.Errorf("%s cannot be blank", attributes.ADConnectionID)
	}

	adUserSyncTaskSettings, err := c.GetADUserSyncTaskSettings(ctx, adConnectionID, adUserSyncTaskSettingsId)
	if err != nil {
		return diag.FromErr(err)
	}

	if adUserSyncTaskSettings != nil {
		d.SetId(*adUserSyncTaskSettings.ID)
		for key, value := range adUserSyncTaskSettings.ToResourceMap() {
			if err := d.Set(key, value); err != nil {
				return diag.FromErr(err)
			}
		}
	} else {
		logging.Infof("ADUserSyncTaskSettings %s does not exist", adUserSyncTaskSettingsId)
	}
	return nil
}
