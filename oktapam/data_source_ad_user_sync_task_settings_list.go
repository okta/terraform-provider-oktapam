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

func dataSourceADUserSyncTaskSettingsIDList() *schema.Resource {
	return &schema.Resource{
		Description: descriptions.SourceADUserSyncTaskSettingsIDList,
		ReadContext: dataSourceADUserSyncTaskSettingsIDListRead,
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
			attributes.ADUserSyncTaskSettingsIDList: {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceADUserSyncTaskSettingsIDListRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
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
	adUserSyncTaskSettingsIDList := make([]string, len(adUserSyncTaskSettingsList))
	for idx, adUserSyncTaskSettings := range adUserSyncTaskSettingsList {
		adUserSyncTaskSettingsIDList[idx] = *adUserSyncTaskSettings.ID
	}

	if err := d.Set(attributes.ADUserSyncTaskSettingsIDList, adUserSyncTaskSettingsIDList); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resource.UniqueId())
	return nil
}
