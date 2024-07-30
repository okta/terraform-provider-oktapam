package oktapam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
)

func dataSourcePasswordSettings() *schema.Resource {
	return &schema.Resource{
		Description: descriptions.SourceServerPasswordSettings,
		ReadContext: dataSourcePasswordSettingsFetch,
		Schema: map[string]*schema.Schema{
			attributes.ResourceGroup: {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions.ResourceGroupID,
			},
			attributes.Project: {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions.ProjectID,
			},
			attributes.ManagedPrivilegedAccounts: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: descriptions.ManagedPrivilegedAccounts,
			},
			attributes.EnablePeriodicRotation: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: descriptions.EnablePeriodicRotation,
			},
			attributes.PeriodicRotationDurationInSeconds: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: descriptions.PeriodicRotationDurationInSeconds,
			},
			attributes.MinLength: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: descriptions.PasswordMinLength,
			},
			attributes.MaxLength: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: descriptions.PasswordMaxLength,
			},
			attributes.CharacterOptions: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						attributes.UpperCase: {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: descriptions.CharacterOptionsUpperCase,
						},
						attributes.LowerCase: {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: descriptions.CharacterOptionsLowerCase,
						},
						attributes.Digits: {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: descriptions.CharacterOptionsDigits,
						},
						attributes.Punctuation: {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: descriptions.CharacterOptionsPunctuation,
						},
					},
				},
			},
		},
	}
}

func dataSourcePasswordSettingsFetch(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	var diags diag.Diagnostics
	c := getLocalClientFromMetadata(m)
	resourceGroupID := d.Get(attributes.ResourceGroup).(string)
	projectID := d.Get(attributes.Project).(string)

	passwordSettings, err := c.GetPasswordSettings(ctx, resourceGroupID, projectID)
	if err != nil {
		return diag.FromErr(err)
	}
	// password settings are considered to not exist if they aren't set or there are no managed accounts set
	if passwordSettings != nil && len(passwordSettings.ManagedPrivilegedAccountsConfig) > 0 {
		for k, v := range passwordSettings.ToResourceMap() {
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
