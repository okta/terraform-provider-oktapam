package oktapam

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"
)

func resourceADTaskSettings() *schema.Resource {
	return &schema.Resource{
		Description:   descriptions.ResourceADTaskSettings,
		CreateContext: resourceADTaskSettingsCreate,
		ReadContext:   resourceADTaskSettingsRead,
		UpdateContext: resourceADTaskSettingsUpdate,
		DeleteContext: resourceADTaskSettingsDelete,
		Schema: map[string]*schema.Schema{
			attributes.ADConnectionID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions.ADConnectionID,
			},
			attributes.Name: {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions.Name,
			},
			attributes.Frequency: {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntBetween(1, 24),
				Description:  descriptions.ADTaskFrequency,
			},
			attributes.IsActive: {
				Type:        schema.TypeBool,
				Required:    true,
				Description: descriptions.ADTaskIsActive,
			},
			attributes.RunTest: {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: descriptions.ADTaskRunTest,
			},
			attributes.StartHourUTC: {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntBetween(0, 23),
				Description:  descriptions.ADTaskStartHourUTC,
			},
			attributes.HostnameAttribute: {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions.HostnameAttribute,
			},
			attributes.AccessAddressAttribute: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions.AccessAddressAttribute,
			},
			attributes.OSAttribute: {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions.OSAttribute,
			},
			attributes.BastionAttribute: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions.BastionAttribute,
			},
			attributes.AltNamesAttributes: {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: descriptions.AltNamesAttributes,
			},
			attributes.AdditionalAttributeMapping: {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: descriptions.AdditionalAttributeMapping,
				MaxItems:    10,
				Elem:        additionalAttributeMappingResource,
			},
			attributes.ADRuleAssignments: {
				Type:        schema.TypeSet,
				Required:    true,
				ForceNew:    true,
				Description: descriptions.ADRuleAssignments,
				Elem:        adRuleAssignmentsResource,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

var additionalAttributeMappingResource = &schema.Resource{
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
}

var adRuleAssignmentsResource = &schema.Resource{
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
}

func resourceADTaskSettingsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaPAMClient)

	adConnId := d.Get(attributes.ADConnectionID).(string)

	//Build ADTaskSettings Api Object
	adTaskSettingsReq := expandADTaskSettings(d)

	//Call api client
	if createdADTS, err := c.CreateADTaskSettings(ctx, adConnId, adTaskSettingsReq); err != nil {
		return diag.FromErr(err)
	} else if createdADTS == nil {
		d.SetId("")
	} else {
		//Set returned id
		d.SetId(createADTaskSettingsResourceID(adConnId, *createdADTS.ID))
	}

	return resourceADTaskSettingsRead(ctx, d, m)
}

func resourceADTaskSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(client.OktaPAMClient)

	adConnId, adTaskSettingsId, err := parseADTaskSettingsResourceID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	adTaskSettings, err := c.GetADTaskSettings(ctx, adConnId, adTaskSettingsId)
	if err != nil {
		return diag.FromErr(err)
	}

	if adTaskSettings != nil && utils.IsNonEmpty(adTaskSettings.ID) {
		for key, value := range flattenADTaskSettings(adTaskSettings) {
			_ = d.Set(key, value)
		}
	} else {
		logging.Infof("ADTaskSettings %s does not exist", adTaskSettingsId)
	}

	return diags
}

func resourceADTaskSettingsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaPAMClient)

	adConnId, adTaskSettingsId, err := parseADTaskSettingsResourceID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	if d.HasChangesExcept(attributes.Name,
		attributes.HostnameAttribute,
		attributes.AccessAddressAttribute,
		attributes.OSAttribute,
		attributes.BastionAttribute,
		attributes.AltNamesAttributes,
		attributes.AdditionalAttributeMapping,
		attributes.ADRuleAssignments) {
		//If deactivated
		if _, active := d.GetChange(attributes.IsActive); d.HasChange(attributes.IsActive) && active != nil && !active.(bool) {
			//Deactivate task
			err := c.DeactivateADTaskSettings(ctx, adConnId, adTaskSettingsId)
			if err != nil {
				return diag.FromErr(err)
			}
		} else if d.HasChanges(attributes.Frequency,
			attributes.IsActive,
			attributes.StartHourUTC) {
			schedule := client.ADTaskSettingsSchedule{
				Frequency:    getIntPtr(attributes.Frequency, d, false),
				StartHourUTC: getIntPtr(attributes.StartHourUTC, d, false),
			}
			err := c.UpdateADTaskSettingsSchedule(ctx, adConnId, adTaskSettingsId, schedule)
			if err != nil {
				return diag.FromErr(err)
			}
		}
	} else {
		//ADTask Settings is immutable and update will create new
		adTaskSettingsReq := expandADTaskSettings(d)
		//Call api client
		if updatedADTS, err := c.UpdateADTaskSettings(ctx, adConnId, adTaskSettingsId, adTaskSettingsReq); err != nil {
			return diag.FromErr(err)
		} else if updatedADTS == nil {
			d.SetId("")
		} else {
			//Set returned id
			d.SetId(*updatedADTS.ID)
		}
	}

	return resourceADTaskSettingsRead(ctx, d, m)
}

func resourceADTaskSettingsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(client.OktaPAMClient)

	adConnId, adTaskSettingsId, err := parseADTaskSettingsResourceID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	err = c.DeleteADTaskSettings(ctx, adConnId, adTaskSettingsId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return diags
}

func expandADTaskSettings(d *schema.ResourceData) client.ADTaskSettings {
	var altNamesAttrs []string
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
	}

	adTaskSettings := client.ADTaskSettings{
		Name:                       getStringPtr(attributes.Name, d, false),
		Frequency:                  getIntPtr(attributes.Frequency, d, false),
		StartHourUTC:               getIntPtr(attributes.StartHourUTC, d, false),
		IsActive:                   getBoolPtr(attributes.IsActive, d, true),
		RunTest:                    getBoolPtr(attributes.RunTest, d, true),
		HostnameAttribute:          getStringPtr(attributes.HostnameAttribute, d, false),
		AccessAddressAttribute:     getStringPtr(attributes.AccessAddressAttribute, d, false),
		OSAttribute:                getStringPtr(attributes.OSAttribute, d, false),
		BastionAttribute:           getStringPtr(attributes.BastionAttribute, d, false),
		AltNamesAttributes:         altNamesAttrs,
		AdditionalAttributeMapping: attrMapping,
		RuleAssignments:            ruleAssignments,
	}

	return adTaskSettings
}

func expandAdditionalAttributeMappings(tfList []interface{}) []*client.ADAdditionalAttribute {
	apiObjects := make([]*client.ADAdditionalAttribute, len(tfList))

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
}

func expandADRuleAssignments(tfList []interface{}) []*client.ADRuleAssignment {
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
}

//Convert API Object to flat map for saving in terraform state
func flattenADTaskSettings(taskSettings *client.ADTaskSettings) map[string]interface{} {
	m := make(map[string]interface{}, 2)

	if taskSettings.Name != nil {
		m[attributes.Name] = *taskSettings.Name
	}
	if taskSettings.Frequency != nil {
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
	}

	return m
}

func createADTaskSettingsResourceID(adConnectionId string, adTaskSettingsId string) string {
	return fmt.Sprintf("%s|%s", adConnectionId, adTaskSettingsId)
}

func parseADTaskSettingsResourceID(resourceId string) (string, string, error) {
	split := strings.Split(resourceId, "|")
	if len(split) != 2 {
		return "", "", fmt.Errorf("oktapam_ad_task_settings id must be in the format of <ad connection id>|<ad task settings id>, received: %s", resourceId)
	}
	return split[0], split[1], nil
}
