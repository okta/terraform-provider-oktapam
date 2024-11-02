package oktapam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/client/wrappers"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
)

func dataSourceDatabasePasswordSettings() *schema.Resource {
	return &schema.Resource{
		Description: descriptions.SourceDatabasePasswordSettings,
		ReadContext: dataSourceDatabasePasswordSettingsRead,
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
				Optional:    true,
				Description: descriptions.EnablePeriodicRotation,
			},
			attributes.PeriodicRotationDurationInSeconds: {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: descriptions.PeriodicRotationDurationInSeconds,
			},
			attributes.MinLength: {
				Type:         schema.TypeInt,
				Optional:     true,
				Description:  descriptions.PasswordMinLength,
				ValidateFunc: validation.IntAtLeast(1),
			},
			attributes.MaxLength: {
				Type:         schema.TypeInt,
				Optional:     true,
				Description:  descriptions.PasswordMaxLength,
				ValidateFunc: validation.IntAtLeast(1),
			},
			attributes.CharacterOptions: {
				Type:        schema.TypeList,
				Optional:    true,
				MinItems:    1,
				MaxItems:    1,
				Description: descriptions.CharacterOptions,
				Elem:        characterOptions,
			},
		},
	}
}

func dataSourceDatabasePasswordSettingsRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
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
		if diags == nil {
			d.SetId(formatPasswordSettingsID(resourceGroupID, projectID))
		}
	} else {
		d.SetId("")
	}
	return diags
}
