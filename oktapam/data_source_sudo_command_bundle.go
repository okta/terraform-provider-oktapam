package oktapam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/client/wrappers"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
)

func dataSourceSudoCommandBundle() *schema.Resource {
	return &schema.Resource{
		Description: descriptions.SourceSudoCommandsBundle,
		ReadContext: dataSourceSudoCommandBundleFetch,
		Schema: map[string]*schema.Schema{
			attributes.ID: {
				Type:     schema.TypeString,
				Required: true,
			},
			attributes.Name: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: descriptions.SudoCommandBundleName,
			},
			attributes.RunAs: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: descriptions.RunAs,
			},
			attributes.NoPasswd: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: descriptions.NoPasswd,
			},
			attributes.NoExec: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: descriptions.NoExec,
			},
			attributes.SetEnv: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: descriptions.SetEnv,
			},
			attributes.AddEnv: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: descriptions.AddEnv,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			attributes.SubEnv: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: descriptions.SubEnv,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			attributes.StructuredCommands: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: descriptions.StructuredCommands,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						attributes.StructuredCommand: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: descriptions.StructuredCommands,
						},
						attributes.StructuredCommandType: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: descriptions.StructuredCommands,
						},
						attributes.StructuredCommandArgsType: {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						attributes.StructuredCommandArgs: {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						attributes.StructuredRenderedCommand: {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSudoCommandBundleFetch(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	var diags diag.Diagnostics
	c := getSDKClientFromMetadata(m)
	id := d.Get(attributes.ID).(string)
	if id == "" {
		return diag.Errorf("%s cannot be blank", attributes.ID)
	}

	sudoCommandsBundle, err := client.GetSudoCommandBundle(ctx, c, id)
	if err != nil {
		return diag.FromErr(err)
	}
	if sudoCommandsBundle != nil {
		wrap := wrappers.SudoCommandBundleWrapper{SudoCommandBundle: sudoCommandsBundle}
		for k, v := range wrap.ToResourceMap() {
			if err := d.Set(k, v); err != nil {
				diags = append(diags, diag.FromErr(err)...)
			}
		}
		if diags == nil {
			d.SetId(id)
		}
	} else {
		d.SetId("")
	}
	return diags
}
