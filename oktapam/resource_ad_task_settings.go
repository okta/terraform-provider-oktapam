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
	/** ADTaskSettings is immutable. Only updateable attributes are status and task frequency/start time
		ADTaskSettings resource has update endpoint, but internally it's a transactional delete & then create operation
		Terraform resource id change and deletion of existing resource confuse terraform. Not sure if there is any way to handle it properly.
		For now, I have marked all the immutable attributes as "ForceNew". If update ADTaskSettings endpoint has additional logic then we need to revisit this.
	    TF Reference - https://www.terraform.io/plugin/sdkv2/schemas/schema-behaviors#forcenew
	*/
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
				Description:  descriptions.ADTaskFrequency,
			},
			attributes.IsActive: {
				Type:        schema.TypeBool,
				Required:    true,
				Description: descriptions.ADTaskIsActive,
			},
			attributes.RunTest: {
				Type:        schema.TypeBool,
				Default:     false,
				Optional:    true,
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
				ForceNew:    true,
				Description: descriptions.HostnameAttribute,
			},
			attributes.AccessAddressAttribute: {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: descriptions.AccessAddressAttribute,
			},
			attributes.OSAttribute: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: descriptions.OSAttribute,
			},
			attributes.BastionAttribute: {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: descriptions.BastionAttribute,
			},
			attributes.AltNamesAttributes: {
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: descriptions.AltNamesAttributes,
			},
			attributes.AdditionalAttributeMapping: {
				Type:        schema.TypeSet,
				Optional:    true,
				ForceNew:    true,
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
			StateContext: importADTaskSettingsState,
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

func resourceADTaskSettingsCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := m.(client.OktaPAMClient)

	adConnID := d.Get(attributes.ADConnectionID).(string)

	//Build ADTaskSettings Api Object
	adTaskSettingsReq := expandADTaskSettings(d)

	//Call api client
	if createdADTS, err := c.CreateADTaskSettings(ctx, adConnID, adTaskSettingsReq); err != nil {
		return diag.FromErr(err)
	} else if createdADTS == nil || !createdADTS.Exists() {
		d.SetId("")
	} else {
		//Set returned id
		d.SetId(*createdADTS.ID)
	}

	return resourceADTaskSettingsRead(ctx, d, m)
}

func resourceADTaskSettingsRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(client.OktaPAMClient)

	adConnID := d.Get(attributes.ADConnectionID).(string)
	adTaskSettingsID := d.Id()

	adTaskSettings, err := c.GetADTaskSettings(ctx, adConnID, adTaskSettingsID)
	if err != nil {
		return diag.FromErr(err)
	}

	if adTaskSettings != nil && utils.IsNonEmpty(adTaskSettings.ID) {
		for key, value := range flattenADTaskSettings(adTaskSettings) {
			_ = d.Set(key, value)
		}
	} else {
		logging.Infof("ADTaskSettings %s does not exist", adTaskSettingsID)
	}

	return diags
}

func resourceADTaskSettingsUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := m.(client.OktaPAMClient)

	adConnID := d.Get(attributes.ADConnectionID).(string)
	adTaskSettingsID := d.Id()

	//If deactivated
	if _, active := d.GetChange(attributes.IsActive); d.HasChange(attributes.IsActive) && active != nil && !active.(bool) {
		//Deactivate task
		err := c.DeactivateADTaskSettings(ctx, adConnID, adTaskSettingsID)
		if err != nil {
			return diag.FromErr(err)
		}
	} else if d.HasChanges(attributes.Frequency,
		attributes.IsActive,
		attributes.StartHourUTC) { //If task become active or schedule change
		schedule := client.ADTaskSettingsSchedule{
			Frequency:    getIntPtr(attributes.Frequency, d, false),
			StartHourUTC: getIntPtr(attributes.StartHourUTC, d, false),
		}
		err := c.UpdateADTaskSettingsSchedule(ctx, adConnID, adTaskSettingsID, schedule)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceADTaskSettingsRead(ctx, d, m)
}

func resourceADTaskSettingsDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(client.OktaPAMClient)

	adConnID := d.Get(attributes.ADConnectionID).(string)
	adTaskSettingsID := d.Id()

	err := c.DeleteADTaskSettings(ctx, adConnID, adTaskSettingsID)
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

func expandAdditionalAttributeMappings(tfList []any) []*client.ADAdditionalAttribute {
	apiObjects := make([]*client.ADAdditionalAttribute, 0, len(tfList))

	for _, tfMapRaw := range tfList {
		tfMap, ok := tfMapRaw.(map[string]any)

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

func expandADRuleAssignments(tfList []any) []*client.ADRuleAssignment {
	apiObjects := make([]*client.ADRuleAssignment, 0, len(tfList))
	for _, tfMapRaw := range tfList {
		tfMap, ok := tfMapRaw.(map[string]any)

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

// flattenADTaskSettings Convert API Object to flat map for saving in terraform state
// API always return false for attribute run_test, regardless of what is passed while creating/updating the resource.
// Don't set Run_Test attribute  while reading the resource back, to avoid tf showing drift during plan while comparing
// it with the previous state (if run_test was set to 'true' initially). In this case, whatever value is coming as
// part of tf config (proposed state) will be set in the tf state.
func flattenADTaskSettings(taskSettings *client.ADTaskSettings) map[string]any {
	m := make(map[string]any, 2)

	if taskSettings.ID != nil {
		m[attributes.ID] = *taskSettings.ID
	}
	if taskSettings.Name != nil {
		m[attributes.Name] = *taskSettings.Name
	}
	if taskSettings.Frequency != nil {
		m[attributes.Frequency] = *taskSettings.Frequency
	}
	if taskSettings.IsActive != nil {
		m[attributes.IsActive] = *taskSettings.IsActive
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
		var flattenedAttrMap []any
		for _, attrMapping := range attrMappings {
			flattenedAttrMap = append(flattenedAttrMap, map[string]any{
				attributes.Label:  attrMapping.Label,
				attributes.Value:  attrMapping.Value,
				attributes.IsGuid: attrMapping.IsGuid,
			})
		}
		m[attributes.AdditionalAttributeMapping] = schema.NewSet(schema.HashResource(additionalAttributeMappingResource), flattenedAttrMap)
	}
	if taskSettings.RuleAssignments != nil {
		rules := taskSettings.RuleAssignments
		var flattenedRules []any
		for _, rule := range rules {
			flattenedRules = append(flattenedRules, map[string]any{
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

func parseADTaskSettingsResourceID(resourceId string) (string, string, error) {
	split := strings.Split(resourceId, "/")
	if len(split) != 2 {
		return "", "", fmt.Errorf("expected format: <connection_id>/<id>, received: %s", resourceId)
	}
	return split[0], split[1], nil
}

func importADTaskSettingsState(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	// d.Id() here is the last argument passed to the `terraform import RESOURCE_TYPE.RESOURCE_NAME RESOURCE_ID` command
	// Id provided for import is in the format <connection_id>/<id>
	adConnectionID, adTaskSettingsID, err := parseADTaskSettingsResourceID(d.Id())

	if err != nil {
		return nil, fmt.Errorf("invalid resource import specifier; %w", err)
	}

	if err := d.Set(attributes.ADConnectionID, adConnectionID); err != nil {
		return nil, err
	}

	d.SetId(adTaskSettingsID)
	return []*schema.ResourceData{d}, nil
}
