package oktapam

import (
	"context"
	"fmt"
	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/client/wrappers"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"
)

var characterOptions = &schema.Resource{
	Schema: map[string]*schema.Schema{
		attributes.UpperCase: {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: descriptions.CharacterOptionsUpperCase,
		},
		attributes.LowerCase: {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: descriptions.CharacterOptionsLowerCase,
		},
		attributes.Digits: {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: descriptions.CharacterOptionsDigits,
		},
		attributes.Punctuation: {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: descriptions.CharacterOptionsPunctuation,
		},
		attributes.RequireFromEachSet: {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: descriptions.CharacterOptionsRequireFromEachSet,
		},
	},
}

func resourceDatabasePasswordSettings() *schema.Resource {
	return &schema.Resource{
		Description:   descriptions.ResourcePasswordSettings,
		CreateContext: resourceDatabasePasswordSettingsCreate,
		ReadContext:   resourceDatabasePasswordSettingsRead,
		UpdateContext: resourceDatabasePasswordSettingsUpdate,
		DeleteContext: resourceDatabasePasswordSettingsDelete,
		Schema: map[string]*schema.Schema{
			attributes.ID: {
				Type:     schema.TypeString,
				Computed: true,
			},
			attributes.ResourceGroup: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: descriptions.ResourceGroupID,
			},
			attributes.Project: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: descriptions.ProjectID,
			},
			attributes.EnablePeriodicRotation: {
				Type:        schema.TypeBool,
				Required:    true,
				Description: descriptions.EnablePeriodicRotation,
			},
			attributes.PeriodicRotationDurationInSeconds: {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: descriptions.PeriodicRotationDurationInSeconds,
			},
			attributes.MinLength: {
				Type:         schema.TypeInt,
				Required:     true,
				Description:  descriptions.PasswordMinLength,
				ValidateFunc: validation.IntAtLeast(1),
			},
			attributes.MaxLength: {
				Type:         schema.TypeInt,
				Required:     true,
				Description:  descriptions.PasswordMaxLength,
				ValidateFunc: validation.IntAtLeast(1),
			},
			attributes.CharacterOptions: {
				Type:        schema.TypeList,
				Required:    true,
				MinItems:    1,
				MaxItems:    1,
				Description: descriptions.CharacterOptions,
				Elem:        characterOptions,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: resourceDatabasePasswordSettingsReadImport,
		},
	}
}

func resourceDatabasePasswordSettingsCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := getSDKClientFromMetadata(m)
	resourceGroupID := d.Get(attributes.ResourceGroup).(string)
	projectID := d.Get(attributes.Project).(string)

	passwordPolicy, diags := readPasswordPolicyFromResource(d)
	if diags != nil {
		return diags
	}

	if err := client.UpdateDatabasePasswordSettings(ctx, c, resourceGroupID, projectID, *passwordPolicy); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(formatPasswordSettingsID(resourceGroupID, projectID))

	return resourceDatabasePasswordSettingsRead(ctx, d, m)
}

func resourceDatabasePasswordSettingsRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	var diags diag.Diagnostics
	c := getSDKClientFromMetadata(m)
	resourceGroupID := d.Get(attributes.ResourceGroup).(string)
	projectID := d.Get(attributes.Project).(string)

	passwordPolicy, err := client.GetDatabasePasswordSettings(ctx, c, resourceGroupID, projectID)
	if err != nil {
		return diag.FromErr(err)
	}
	if passwordPolicy != nil {
		wrap := wrappers.PasswordPolicyWrapper{PasswordPolicy: *passwordPolicy, UseManagedPrivilegedAccountsConfig: false}
		for k, v := range wrap.ToResourceMap(nil) {
			if err := d.Set(k, v); err != nil {
				diags = append(diags, diag.FromErr(err)...)
			}
		}
	} else {
		d.SetId("")
	}
	return diags
}

func resourceDatabasePasswordSettingsUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := getSDKClientFromMetadata(m)
	resourceGroupID := d.Get(attributes.ResourceGroup).(string)
	projectID := d.Get(attributes.Project).(string)

	passwordPolicy, diags := readPasswordPolicyFromResource(d)

	if diags != nil {
		return diags
	}

	if err := client.UpdateDatabasePasswordSettings(ctx, c, resourceGroupID, projectID, *passwordPolicy); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(formatPasswordSettingsID(resourceGroupID, projectID))
	return resourceDatabasePasswordSettingsRead(ctx, d, m)
}

func resourceDatabasePasswordSettingsDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	// there isn't a true delete with password settings, and as databases don't specify ManagedPrivilegedAccounts,
	// the best we can do is to set rotation to false and some default required values.
	c := getSDKClientFromMetadata(m)
	resourceGroupID := d.Get(attributes.ResourceGroup).(string)
	projectID := d.Get(attributes.Project).(string)

	if err := client.UpdateDatabasePasswordSettings(ctx, c, resourceGroupID, projectID,
		pam.PasswordPolicy{
			EnablePeriodicRotation: false,
			MinLengthInBytes:       8,
			MaxLengthInBytes:       64,
			CharacterOptions: pam.PasswordPolicyCharacterOptions{
				UpperCase:   utils.AsBoolPtr(true),
				LowerCase:   utils.AsBoolPtr(true),
				Digits:      utils.AsBoolPtr(true),
				Punctuation: utils.AsBoolPtr(true),
			},
		}); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func resourceDatabasePasswordSettingsReadImport(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	// d.Id() here is the last argument passed to the `terraform import RESOURCE_TYPE.RESOURCE_NAME RESOURCE_ID` command
	// Id provided for import is in the format <resource_group_id>/<project_id>

	resourceGroupID, projectID, err := parsePasswordSettingsID(d.Id())
	if err != nil {
		return nil, fmt.Errorf("invalid resource import specifier: %w", err)
	}
	d.Set(attributes.ResourceGroup, resourceGroupID)
	d.Set(attributes.Project, projectID)

	return []*schema.ResourceData{d}, nil
}

func readPasswordPolicyFromResource(d *schema.ResourceData) (*pam.PasswordPolicy, diag.Diagnostics) {
	var characterOptions pam.PasswordPolicyCharacterOptions
	characterOptionsM := GetTypeListMapFromResource(attributes.CharacterOptions, d)
	if characterOptionsM != nil {
		characterOptions = pam.PasswordPolicyCharacterOptions{
			UpperCase:          GetBoolPtrFromElement(attributes.UpperCase, characterOptionsM, true),
			LowerCase:          GetBoolPtrFromElement(attributes.LowerCase, characterOptionsM, true),
			Digits:             GetBoolPtrFromElement(attributes.Digits, characterOptionsM, true),
			Punctuation:        GetBoolPtrFromElement(attributes.Punctuation, characterOptionsM, true),
			RequireFromEachSet: GetBoolPtrFromElement(attributes.RequireFromEachSet, characterOptionsM, true),
		}
	}

	managedPrivilegedAccounts, diags := GetStringSliceFromResource(attributes.ManagedPrivilegedAccounts, d, true)

	passwordSettings := &pam.PasswordPolicy{
		EnablePeriodicRotation:            d.Get(attributes.EnablePeriodicRotation).(bool),
		PeriodicRotationDurationInSeconds: GetInt32PtrFromResource(attributes.PeriodicRotationDurationInSeconds, d, false),
		MinLengthInBytes:                  GetInt32FromResource(attributes.MinLength, d),
		MaxLengthInBytes:                  GetInt32FromResource(attributes.MaxLength, d),
		CharacterOptions:                  characterOptions,
		ManagedPrivilegedAccountsConfig:   managedPrivilegedAccounts,
	}

	if passwordSettings.MinLengthInBytes > passwordSettings.MaxLengthInBytes {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("%s cannot be greater than %s in %s", attributes.MinLength, attributes.MaxLength, attributes.PasswordSettings),
		})
	}

	if passwordSettings.EnablePeriodicRotation && *passwordSettings.PeriodicRotationDurationInSeconds <= 0 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("%s must be set to a positive, non-zero value when %s is set to true in %s", attributes.PeriodicRotationDurationInSeconds, attributes.EnablePeriodicRotation, attributes.PasswordSettings),
		})
	} else if !passwordSettings.EnablePeriodicRotation && passwordSettings.PeriodicRotationDurationInSeconds != nil && *passwordSettings.PeriodicRotationDurationInSeconds != 0 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  fmt.Sprintf("%s will be ignored since %s is set to false in %s", attributes.PeriodicRotationDurationInSeconds, attributes.EnablePeriodicRotation, attributes.PasswordSettings),
		})
	}

	if diags != nil {
		return nil, diags
	}

	return passwordSettings, nil
}
