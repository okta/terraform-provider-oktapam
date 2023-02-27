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
	/** ADUserSyncTaskSettings is immutable. Only updateable attributes are active status and task frequency/start time.
		ADUserSyncTaskSettings resource has update endpoint, but internally it's a transactional delete & then create
	    operation. Terraform resource id change and deletion of existing resource confuse terraform. Not sure if there is
	    any way to handle it properly. For now, I have marked all the immutable attributes as "ForceNew". If update
	    ADUserSyncTaskSettings endpoint has additional logic then we need to revisit this.
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
				Required:    true,
				ForceNew:    true,
				Description: descriptions.ADUserSyncTaskSettingsUPNField,
			},
			attributes.SIDField: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: descriptions.ADUserSyncTaskSettingsSIDField,
			},
			attributes.IsActive: {
				Type:        schema.TypeBool,
				Required:    true,
				Description: descriptions.ADUserSyncTaskSettingsIsActive,
			},
			/*attributes.RunTest: {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: descriptions.ADTaskRunTest,
			},*/
		},
		Importer: &schema.ResourceImporter{
			StateContext: importADUserSyncTaskSettingsState,
		},
	}
}

/*var additionalAttributeMappingResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		attributes.Label: {
			Type:        schema.TypeString,
			Required:    true,
			Description: descriptions.AdditionalAttributeMappingLabel,
		},
		attributes.Value: {
			Type:        schema.TypeString,
			Required:    true,
			Description: descriptions.AdditionalAttributeMappingValue,
		},
		attributes.IsGuid: {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: descriptions.AdditionalAttributeMappingIsGuid,
		},
	},
}*/

/*var adRuleAssignmentsResource = &schema.Resource{
	Description: descriptions.ResourceADRuleAssignment,
	Schema: map[string]*schema.Schema{
		attributes.ID: {
			Type:     schema.TypeString,
			Computed: true,
		},
		attributes.ADRuleAssignmentsBaseDN: {
			Type:        schema.TypeString,
			Required:    true,
			Description: descriptions.ADRuleAssignmentsBaseDN,
		},
		attributes.ADRuleAssignmentsLDAPQueryFilter: {
			Type:        schema.TypeString,
			Required:    true,
			Description: descriptions.ADRuleAssignmentsLDAPQueryFilter,
		},
		attributes.ADRuleAssignmentsProjectID: {
			Type:        schema.TypeString,
			Required:    true,
			Description: descriptions.ADRuleAssignmentsProjectID,
		},
		attributes.ADRuleAssignmentsPriority: {
			Type:     schema.TypeInt,
			Required: true,
		},
	},
}*/

func resourceADUserSyncTaskSettingsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaPAMClient)

	adConnID := d.Get(attributes.ADConnectionID).(string)

	//Build ADUserSyncTaskSettings Api Object
	//adUserSyncTaskSettingsReq := expandADUserSyncTaskSettings(d)
	adUserSyncTaskSettingsReq := client.ADUserSyncTaskSettings{
		Name:            getStringPtr(attributes.Name, d, false),
		Frequency:       getIntPtr(attributes.Frequency, d, false),
		StartHourUTC:    getIntPtr(attributes.StartHourUTC, d, false),
		BaseDN:          getStringPtr(attributes.BaseDN, d, false),
		LDAPQueryFilter: getStringPtr(attributes.LDAPQueryFilter, d, false),
		UPNField:        getStringPtr(attributes.UPNField, d, false),
		SIDField:        getStringPtr(attributes.SIDField, d, false),
		IsActive:        getBoolPtr(attributes.IsActive, d, true),
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
	var diags diag.Diagnostics
	c := m.(client.OktaPAMClient)

	adConnID := d.Get(attributes.ADConnectionID).(string)
	adUserSyncTaskSettingsID := d.Id()

	adUserSyncTaskSettings, err := c.GetADUserSyncTaskSettings(ctx, adConnID, adUserSyncTaskSettingsID)
	if err != nil {
		return diag.FromErr(err)
	}

	if adUserSyncTaskSettings != nil && utils.IsNonEmpty(adUserSyncTaskSettings.ID) {
		//for key, value := range flattenADUserSyncTaskSettings(adUserSyncTaskSettings) {
		for key, value := range adUserSyncTaskSettings.ToResourceMap() {
			_ = d.Set(key, value)
		}
	} else {
		logging.Infof("ADUserSyncTaskSettings %s does not exist", adUserSyncTaskSettingsID)
	}

	return diags
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
		if active.(bool) {
			err := c.ActivateADUserSyncTaskSettings(ctx, adConnID, adUserSyncTaskSettingsID)
			if err != nil {
				return diag.FromErr(err)
			}
		} else {
			err := c.DeactivateADUserSyncTaskSettings(ctx, adConnID, adUserSyncTaskSettingsID)
			if err != nil {
				return diag.FromErr(err)
			}
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

func expandADUserSyncTaskSettings(d *schema.ResourceData) client.ADUserSyncTaskSettings {
	/*var altNamesAttrs []string
	if v, ok := d.GetOk(attributes.AltNamesAttributes); ok && v.(*schema.Set).Len() > 0 {
		altNamesAttrs = utils.ExpandStringSet(v.(*schema.Set))
	}

	var attrMapping []*client.ADAdditionalAttribute
	if v, ok := d.GetOk(attributes.AdditionalAttributeMapping); ok && v.(*schema.Set).Len() > 0 {
		attrMapping = expandAdditionalAttributeMappings(v.(*schema.Set).List())
	}

	var ruleAssignments []*client.ADRuleAssignment
	if v, ok := d.GetOk(attributes.ADRuleAssignments); ok && v.(*schema.Set).Len() > 0 {
		ruleAssignments = expandADRuleAssignments(v.(*schema.Set).List())
	}*/

	adUserSyncTaskSettings := client.ADUserSyncTaskSettings{
		Name: getStringPtr(attributes.Name, d, false),
		/*Frequency:                  getIntPtr(attributes.Frequency, d, false),
		StartHourUTC:               getIntPtr(attributes.StartHourUTC, d, false),
		IsActive:                   getBoolPtr(attributes.IsActive, d, true),
		RunTest:                    getBoolPtr(attributes.RunTest, d, true),
		HostnameAttribute:          getStringPtr(attributes.HostnameAttribute, d, false),
		AccessAddressAttribute:     getStringPtr(attributes.AccessAddressAttribute, d, false),
		OSAttribute:                getStringPtr(attributes.OSAttribute, d, false),
		BastionAttribute:           getStringPtr(attributes.BastionAttribute, d, false),
		AltNamesAttributes:         altNamesAttrs,
		AdditionalAttributeMapping: attrMapping,
		RuleAssignments:            ruleAssignments,*/
	}

	return adUserSyncTaskSettings
}

/*func expandAdditionalAttributeMappings(tfList []interface{}) []*client.ADAdditionalAttribute {
	apiObjects := make([]*client.ADAdditionalAttribute, 0, len(tfList))

	for _, tfMapRaw := range tfList {
		tfMap, ok := tfMapRaw.(map[string]interface{})

		if !ok {
			continue
		}

		apiObject := &client.ADAdditionalAttribute{
			Label:  tfMap[attributes.Label].(string),
			Value:  tfMap[attributes.Value].(string),
			IsGuid: tfMap[attributes.IsGuid].(bool),
		}

		apiObjects = append(apiObjects, apiObject)
	}

	return apiObjects
}*/

/*func expandADRuleAssignments(tfList []interface{}) []*client.ADRuleAssignment {
	apiObjects := make([]*client.ADRuleAssignment, 0, len(tfList))
	for _, tfMapRaw := range tfList {
		tfMap, ok := tfMapRaw.(map[string]interface{})

		if !ok {
			continue
		}

		apiObject := &client.ADRuleAssignment{
			BaseDN:          tfMap[attributes.ADRuleAssignmentsBaseDN].(string),
			LDAPQueryFilter: tfMap[attributes.ADRuleAssignmentsLDAPQueryFilter].(string),
			ProjectID:       tfMap[attributes.ADRuleAssignmentsProjectID].(string),
			Priority:        tfMap[attributes.ADRuleAssignmentsPriority].(int),
		}

		apiObjects = append(apiObjects, apiObject)
	}

	return apiObjects
}*/

// Convert API Object to flat map for saving in terraform state
func flattenADUserSyncTaskSettings(userSyncTaskSettings *client.ADUserSyncTaskSettings) map[string]interface{} {
	m := make(map[string]interface{}, 2)

	if userSyncTaskSettings.ID != nil {
		m[attributes.ID] = *userSyncTaskSettings.ID
	}
	if userSyncTaskSettings.Name != nil {
		m[attributes.Name] = *userSyncTaskSettings.Name
	}
	/*if taskSettings.Frequency != nil {
		m[attributes.Frequency] = *taskSettings.Frequency
	}
	if taskSettings.IsActive != nil {
		m[attributes.IsActive] = *taskSettings.IsActive
	}
	if taskSettings.RunTest != nil {
		m[attributes.RunTest] = *taskSettings.IsActive
	}
	if taskSettings.StartHourUTC != nil {
		m[attributes.StartHourUTC] = *taskSettings.StartHourUTC
	}
	if taskSettings.HostnameAttribute != nil {
		m[attributes.HostnameAttribute] = *taskSettings.HostnameAttribute
	}
	if taskSettings.OSAttribute != nil {
		m[attributes.OSAttribute] = taskSettings.OSAttribute
	}
	if taskSettings.AccessAddressAttribute != nil {
		m[attributes.AccessAddressAttribute] = *taskSettings.AccessAddressAttribute
	}
	if taskSettings.BastionAttribute != nil {
		m[attributes.BastionAttribute] = *taskSettings.BastionAttribute
	}
	if taskSettings.AltNamesAttributes != nil {
		m[attributes.AltNamesAttributes] = utils.ConvertStringSliceToSet(taskSettings.AltNamesAttributes)
	}
	if taskSettings.AdditionalAttributeMapping != nil {
		attrMappings := taskSettings.AdditionalAttributeMapping
		var flattenedAttrMap []interface{}
		for _, attrMapping := range attrMappings {
			flattenedAttrMap = append(flattenedAttrMap, map[string]interface{}{
				attributes.Label:  attrMapping.Label,
				attributes.Value:  attrMapping.Value,
				attributes.IsGuid: attrMapping.IsGuid,
			})
		}
		m[attributes.AdditionalAttributeMapping] = schema.NewSet(schema.HashResource(additionalAttributeMappingResource), flattenedAttrMap)
	}
	if taskSettings.RuleAssignments != nil {
		rules := taskSettings.RuleAssignments
		var flattenedRules []interface{}
		for _, rule := range rules {
			flattenedRules = append(flattenedRules, map[string]interface{}{
				attributes.ID:                               rule.ID,
				attributes.ADRuleAssignmentsBaseDN:          rule.BaseDN,
				attributes.ADRuleAssignmentsLDAPQueryFilter: rule.LDAPQueryFilter,
				attributes.ADRuleAssignmentsProjectID:       rule.ProjectID,
				attributes.ADRuleAssignmentsPriority:        rule.Priority,
			})
		}
		m[attributes.ADRuleAssignments] = schema.NewSet(schema.HashResource(adRuleAssignmentsResource), flattenedRules)
	}*/

	return m
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
