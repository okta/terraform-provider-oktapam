package oktapam

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/typed_strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceADUserSyncTaskSettingsList() *schema.Resource {
	return &schema.Resource{
		Description: descriptions.SourceADUserSyncTaskSettingsList,
		ReadContext: dataSourceADUserSyncTaskSettingsListRead,
		Schema: map[string]*schema.Schema{
			// Query parameter values
			attributes.ADUserSyncTaskSettingsStatus: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions.FilterADUserSyncTaskSettingsStatus,
				ValidateFunc: validation.StringInSlice(
					[]string{
						typed_strings.TaskSettingsStatusActive.String(),
						typed_strings.TaskSettingsStatusInActive.String(),
					}, false),
			},
			attributes.ADConnectionID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions.FilterConnectionID,
			},
			// Return value
			attributes.ADUserSyncTaskSettingsList: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: descriptions.SourceADUserSyncTaskSettingsList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						attributes.ADConnectionID: {
							Type:        schema.TypeString,
							Required:    true,
							Description: descriptions.ADConnectionID,
						},
						attributes.Name: {
							Type:        schema.TypeString,
							Optional:    true,
							Description: descriptions.Name,
						},
						attributes.ID: {
							Type:     schema.TypeString,
							Required: true,
						},
						attributes.Frequency: {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: descriptions.ADUserSyncTaskSettingsFrequency,
						},
						attributes.StartHourUTC: {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: descriptions.ADUserSyncTaskSettingsStartHourUTC,
						},
						attributes.BaseDN: {
							Type:        schema.TypeString,
							Optional:    true,
							Description: descriptions.ADUserSyncTaskSettingsBaseDN,
						},
						attributes.LDAPQueryFilter: {
							Type:        schema.TypeString,
							Optional:    true,
							Description: descriptions.ADUserSyncTaskSettingsLDAPQueryFilter,
						},
						attributes.UPNField: {
							Type:        schema.TypeString,
							Optional:    true,
							Description: descriptions.ADUserSyncTaskSettingsUPNField,
						},
						attributes.SIDField: {
							Type:        schema.TypeString,
							Optional:    true,
							Description: descriptions.ADUserSyncTaskSettingsSIDField,
						},
						attributes.IsActive: {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: descriptions.ADUserSyncTaskSettingsIsActive,
						},
						attributes.RunTest: {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: descriptions.ADUserSyncTaskSettingsRunTest,
						},
					},
				},
			},
		},
	}
}

func dataSourceADUserSyncTaskSettingsListRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := m.(client.OktaPAMClient)
	parameters := client.ListADUserSyncTaskSettingsParameters{}

	connectionID := d.Get(attributes.ADConnectionID).(string)
	if connectionID == "" {
		return diag.Errorf("%s cannot be blank", attributes.ADConnectionID)
	}

	//Extract parameters
	status := d.Get(attributes.ADUserSyncTaskSettingsStatus).(string)
	parameters.Status = status

	adUserSyncTaskSettingsList, err := c.ListADUserSyncTaskSettings(ctx, connectionID, parameters)
	if err != nil {
		return diag.FromErr(err)
	}

	adUserSyncTaskSettingsData := make([]map[string]any, len(adUserSyncTaskSettingsList))
	for idx, adUserSyncTaskSettings := range adUserSyncTaskSettingsList {
		adUserSyncTaskSettingsData[idx] = adUserSyncTaskSettings.ToResourceMap()
	}

	if err := d.Set(attributes.ADUserSyncTaskSettingsList, adUserSyncTaskSettingsData); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resource.UniqueId())
	return nil
}
