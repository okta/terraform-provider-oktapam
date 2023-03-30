package oktapam

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"
)

func resourceADUserSyncTaskSettings() *schema.Resource {
	/*
	 ADUserSyncTaskSettings is immutable. Only updatable attributes are active status and task frequency/start time.
	 ADUserSyncTaskSettings resource has update endpoint, but internally it's a transactional delete & then create
	 operation. Terraform resource id change and deletion of existing resource confuse terraform. That is why all the
	 immutable attributes are marked here as "ForceNew".
	 TF Reference - https://www.terraform.io/plugin/sdkv2/schemas/schema-behaviors#forcenew
	*/
	return &schema.Resource{
		Description:   descriptions.ResourceADUserSyncTaskSettings,
		CreateContext: resourceADUserSyncTaskSettingsCreate,
		ReadContext:   resourceADUserSyncTaskSettingsRead,
		UpdateContext: resourceADUserSyncTaskSettingsUpdate,
		DeleteContext: resourceADUserSyncTaskSettingsDelete,
		Schema: map[string]*schema.Schema{
			attributes.ADConnectionID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions.ADConnectionID,
			},
			attributes.Name: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: descriptions.Name,
			},
			attributes.ID: {
				Type:     schema.TypeString,
				Computed: true,
			},
			attributes.Frequency: {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntBetween(1, 24),
				Description:  descriptions.ADUserSyncTaskSettingsFrequency,
			},
			attributes.StartHourUTC: {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntBetween(0, 23),
				Description:  descriptions.ADUserSyncTaskSettingsStartHourUTC,
			},
			attributes.BaseDN: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: descriptions.ADUserSyncTaskSettingsBaseDN,
			},
			attributes.LDAPQueryFilter: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: descriptions.ADUserSyncTaskSettingsLDAPQueryFilter,
			},
			attributes.UPNField: {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "userPrincipalName",
				ForceNew:    true,
				Description: descriptions.ADUserSyncTaskSettingsUPNField,
			},
			attributes.SIDField: {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "objectSID",
				ForceNew:    true,
				Description: descriptions.ADUserSyncTaskSettingsSIDField,
			},
			attributes.IsActive: {
				Type:        schema.TypeBool,
				Required:    true,
				Description: descriptions.ADUserSyncTaskSettingsIsActive,
			},
			attributes.RunTest: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: descriptions.ADUserSyncTaskSettingsRunTest,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: importADUserSyncTaskSettingsState,
		},
	}
}

func resourceADUserSyncTaskSettingsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaPAMClient)

	adConnID := d.Get(attributes.ADConnectionID).(string)

	//Build ADUserSyncTaskSettings Api Object
	adUserSyncTaskSettingsReq := client.ADUserSyncTaskSettings{
		Name:            getStringPtr(attributes.Name, d, false),
		Frequency:       getIntPtr(attributes.Frequency, d, false),
		StartHourUTC:    getIntPtr(attributes.StartHourUTC, d, false),
		BaseDN:          getStringPtr(attributes.BaseDN, d, false),
		LDAPQueryFilter: getStringPtr(attributes.LDAPQueryFilter, d, false),
		UPNField:        getStringPtr(attributes.UPNField, d, false),
		SIDField:        getStringPtr(attributes.SIDField, d, false),
		IsActive:        getBoolPtr(attributes.IsActive, d, true),
		RunTest:         getBoolPtr(attributes.RunTest, d, true),
	}

	//Call api client
	if createdADUSTS, err := c.CreateADUserSyncTaskSettings(ctx, adConnID, adUserSyncTaskSettingsReq); err != nil {
		return diag.FromErr(err)
	} else if createdADUSTS == nil || !createdADUSTS.Exists() {
		d.SetId("")
	} else {
		//Set returned id
		d.SetId(*createdADUSTS.ID)
	}

	return resourceADUserSyncTaskSettingsRead(ctx, d, m)
}

func resourceADUserSyncTaskSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaPAMClient)

	adConnID := d.Get(attributes.ADConnectionID).(string)
	adUserSyncTaskSettingsID := d.Id()

	adUserSyncTaskSettings, err := c.GetADUserSyncTaskSettings(ctx, adConnID, adUserSyncTaskSettingsID)
	if err != nil {
		return diag.FromErr(err)
	}

	if adUserSyncTaskSettings != nil && utils.IsNonEmpty(adUserSyncTaskSettings.ID) {
		for key, value := range adUserSyncTaskSettings.ToResourceMap() {
			_ = d.Set(key, value)
		}
	} else {
		logging.Infof("ADUserSyncTaskSettings %s does not exist", adUserSyncTaskSettingsID)
	}

	return nil
}

func resourceADUserSyncTaskSettingsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaPAMClient)

	adConnID := d.Get(attributes.ADConnectionID).(string)
	adUserSyncTaskSettingsID := d.Id()

	// Make API calls for schedule change and active status change as necessary

	if d.HasChanges(attributes.Frequency, attributes.StartHourUTC) {
		schedule := client.ADUserSyncTaskSettingsSchedule{
			Frequency:    getIntPtr(attributes.Frequency, d, false),
			StartHourUTC: getIntPtr(attributes.StartHourUTC, d, false),
		}
		err := c.UpdateADUserSyncTaskSettingsSchedule(ctx, adConnID, adUserSyncTaskSettingsID, schedule)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if _, active := d.GetChange(attributes.IsActive); d.HasChange(attributes.IsActive) && active != nil {
		isActive := active.(bool)
		newState := client.ADUserSyncTaskSettingsState{
			IsActive: &isActive,
		}
		err := c.UpdateADUserSyncTaskSettingsState(ctx, adConnID, adUserSyncTaskSettingsID, newState)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceADUserSyncTaskSettingsRead(ctx, d, m)
}

func resourceADUserSyncTaskSettingsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(client.OktaPAMClient)

	adConnID := d.Get(attributes.ADConnectionID).(string)
	adUserSyncTaskSettingsID := d.Id()

	err := c.DeleteADUserSyncTaskSettings(ctx, adConnID, adUserSyncTaskSettingsID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return diags
}

func parseADUserSyncTaskSettingsResourceID(resourceId string) (string, string, error) {
	split := strings.Split(resourceId, "/")
	if len(split) != 2 {
		return "", "", fmt.Errorf("expected format: <connection_id>/<id>, received: %s", resourceId)
	}
	return split[0], split[1], nil
}

func importADUserSyncTaskSettingsState(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// d.Id() here is the last argument passed to the `terraform import RESOURCE_TYPE.RESOURCE_NAME RESOURCE_ID` command
	// Id provided for import is in the format <connection_id>/<id>
	adConnectionID, adUserSyncTaskSettingsID, err := parseADUserSyncTaskSettingsResourceID(d.Id())

	if err != nil {
		return nil, fmt.Errorf("invalid resource import specifier; %w", err)
	}

	if err := d.Set(attributes.ADConnectionID, adConnectionID); err != nil {
		return nil, err
	}

	d.SetId(adUserSyncTaskSettingsID)
	return []*schema.ResourceData{d}, nil
}
