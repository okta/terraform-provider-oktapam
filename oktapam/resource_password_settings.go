package oktapam

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"
)

func resourcePasswordSettings() *schema.Resource {
	return &schema.Resource{
		Description:   "", // TODO: add description
		CreateContext: resourcePasswordSettingsCreate,
		ReadContext:   resourcePasswordSettingsRead,
		UpdateContext: resourcePasswordSettingsUpdate,
		DeleteContext: resourcePasswordSettingsDelete,
		Schema: map[string]*schema.Schema{
			attributes.ID: {
				Type:     schema.TypeString,
				Computed: true,
			},
			attributes.ResourceGroup: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "", // TODO: add description
			},
			attributes.Project: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "", // TODO: add description
			},
			attributes.ManagedPrivilegedAccounts: {
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			attributes.EnablePeriodicRotation: {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "", // TODO: add description
			},
			attributes.PeriodicRotationDurationInSeconds: {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "", // TODO add description
			},
			attributes.MinLength: {
				Type:         schema.TypeInt,
				Required:     true,
				Description:  "", // TODO: add description
				ValidateFunc: validation.IntAtLeast(1),
			},
			attributes.MaxLength: {
				Type:         schema.TypeInt,
				Required:     true,
				Description:  "", // TODO: add description
				ValidateFunc: validation.IntAtLeast(1),
			},
			attributes.CharacterOptions: {
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						attributes.UpperCase: {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "", // TODO: add description
						},
						attributes.LowerCase: {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "", // TODO: add description
						},
						attributes.Digits: {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "", // TODO: add description
						},
						attributes.Punctuation: {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "", // TODO: add description
						},
					},
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: resourcePasswordSettingsReadImport,
		},
	}
}

func formatPasswordSettingsID(resourceGroupID string, projectID string) string {
	// password settings don't have an identifier in itself and is really an attribute of a project.
	// we manage it as a separate resource since it's lifecycle is somewhat separate from a project
	return fmt.Sprintf("%s/%s", resourceGroupID, projectID)
}

func parsePasswordSettingsID(resourceID string) (string, string, error) {
	split := strings.Split(resourceID, "/")
	if len(split) != 2 {
		return "", "", fmt.Errorf("expected format: <resource_group_id>/<project_id>, received: %s", resourceID)
	}
	return split[0], split[1], nil
}

func resourcePasswordSettingsCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := m.(client.OktaPAMClient)
	resourceGroupID := d.Get(attributes.ResourceGroup).(string)
	projectID := d.Get(attributes.Project).(string)

	passwordSettings, diags := readPasswordSettingsFromResource(d)

	if diags != nil {
		return diags
	}

	if err := c.UpdatePasswordSettings(ctx, resourceGroupID, projectID, passwordSettings); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(formatPasswordSettingsID(resourceGroupID, projectID))

	return resourcePasswordSettingsRead(ctx, d, m)
}

func resourcePasswordSettingsRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(client.OktaPAMClient)
	resourceGroupID := d.Get(attributes.ResourceGroup).(string)
	projectID := d.Get(attributes.Project).(string)

	passwordSettings, err := c.GetPasswordSettings(ctx, resourceGroupID, projectID)
	if err != nil {
		return diag.FromErr(err)
	}
	if passwordSettings != nil && len(passwordSettings.ManagedPrivilegedAccountsConfig) > 0 {
		for k, v := range passwordSettings.ToResourceMap() {
			d.Set(k, v)
		}
	} else {
		d.SetId("")
	}
	return diags
}

func resourcePasswordSettingsUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := m.(client.OktaPAMClient)
	resourceGroupID := d.Get(attributes.ResourceGroup).(string)
	projectID := d.Get(attributes.Project).(string)

	passwordSettings, diags := readPasswordSettingsFromResource(d)

	if diags != nil {
		return diags
	}

	if err := c.UpdatePasswordSettings(ctx, resourceGroupID, projectID, passwordSettings); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(formatPasswordSettingsID(resourceGroupID, projectID))
	return resourcePasswordSettingsRead(ctx, d, m)
}

func resourcePasswordSettingsDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	// there isn't a true delete with password settings, the best we can do is to set the managed accounts list to be empty
	c := m.(client.OktaPAMClient)
	resourceGroupID := d.Get(attributes.ResourceGroup).(string)
	projectID := d.Get(attributes.Project).(string)

	passwordSettings := &client.PasswordSettings{
		EnablePeriodicRotation:            utils.AsBoolPtrZero(false, true),
		PeriodicRotationDurationInSeconds: utils.AsIntPtrZero(0, true),
		MinLengthInBytes:                  utils.AsIntPtr(8),
		MaxLengthInBytes:                  utils.AsIntPtr(64),
		ManagedPrivilegedAccountsConfig:   []string{},
	}

	if err := c.UpdatePasswordSettings(ctx, resourceGroupID, projectID, passwordSettings); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func resourcePasswordSettingsReadImport(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	// d.Id() here is the last argument passed to the `terraform import RESOURCE_TYPE.RESOURCE_NAME RESOURCE_ID` command
	// Id provided for import is in the format <resource_group_id>/<project_id>

	resourceGroupID, projectID, err := parsePasswordSettingsID(d.Id())
	if err != nil {
		return nil, err
	}
	d.Set(attributes.ResourceGroup, resourceGroupID)
	d.Set(attributes.Project, projectID)

	diags := resourcePasswordSettingsRead(ctx, d, m)
	for _, d := range diags {
		if d.Severity == diag.Error {
			return nil, fmt.Errorf(d.Summary)
		}
	}

	return []*schema.ResourceData{d}, nil
}

func readPasswordSettingsFromResource(d *schema.ResourceData) (*client.PasswordSettings, diag.Diagnostics) {
	var characterOptions *client.CharacterOptions
	characterOptionsM := GetTypeListMapFromResource(attributes.CharacterOptions, d)
	if characterOptionsM != nil {
		characterOptions = &client.CharacterOptions{
			UpperCase:   GetBoolPtrFromElement(attributes.UpperCase, characterOptionsM, true),
			LowerCase:   GetBoolPtrFromElement(attributes.LowerCase, characterOptionsM, true),
			Digits:      GetBoolPtrFromElement(attributes.Digits, characterOptionsM, true),
			Punctuation: GetBoolPtrFromElement(attributes.Punctuation, characterOptionsM, true),
		}
	}

	managedPrivilegedAccounts, diags := GetStringSliceFromResource(attributes.ManagedPrivilegedAccounts, d, false)

	passwordSettings := &client.PasswordSettings{
		EnablePeriodicRotation:            GetBoolPtrFromResource(attributes.EnablePeriodicRotation, d, true),
		PeriodicRotationDurationInSeconds: GetIntPtrFromResource(attributes.PeriodicRotationDurationInSeconds, d, true),
		MinLengthInBytes:                  GetIntPtrFromResource(attributes.MinLength, d, true),
		MaxLengthInBytes:                  GetIntPtrFromResource(attributes.MaxLength, d, true),
		CharacterOptions:                  characterOptions,
		ManagedPrivilegedAccountsConfig:   managedPrivilegedAccounts,
	}

	if *passwordSettings.MinLengthInBytes > *passwordSettings.MaxLengthInBytes {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("%s cannot be greater than %s in %s", attributes.MinLength, attributes.MaxLength, attributes.PasswordSettings),
		})
	}

	if *passwordSettings.EnablePeriodicRotation && *passwordSettings.PeriodicRotationDurationInSeconds <= 0 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("%s must be set to a positive, non-zero value when %s is set to true in %s", attributes.PeriodicRotationDurationInSeconds, attributes.EnablePeriodicRotation, attributes.PasswordSettings),
		})
	} else if !*passwordSettings.EnablePeriodicRotation && *passwordSettings.PeriodicRotationDurationInSeconds != 0 {
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
