package oktapam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
)

func dataSourceSudoCommandsBundle() *schema.Resource {
	return &schema.Resource{
		Description: descriptions.SourceSudoCommandsBundle,
		ReadContext: dataSourceSudoCommandsBundleFetch,
		Schema: map[string]*schema.Schema{
			attributes.ID: {
				Type:     schema.TypeString,
				Required: true,
			},
			attributes.Name: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: descriptions.Name,
			},
			attributes.RunAs: {
				Type:     schema.TypeString,
				Computed: true,
			},
			attributes.NoPasswd: {
				Type:     schema.TypeBool,
				Computed: true,
			},
			attributes.NoExec: {
				Type:     schema.TypeString,
				Computed: true,
			},
			attributes.SetEnv: {
				Type:     schema.TypeString,
				Computed: true,
			},
			attributes.AddEnv: {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			attributes.SubEnv: {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			attributes.StructuredCommands: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: descriptions.StructuredCommands,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						attributes.StructuredCommand: {
							Type:     schema.TypeString,
							Computed: true,
						},
						attributes.StructuredCommandType: {
							Type:     schema.TypeString,
							Computed: true,
						},
						attributes.StructuredCommandArgsType: {
							Type:     schema.TypeString,
							Computed: true,
						},
						attributes.StructuredCommandArgs: {
							Type:     schema.TypeString,
							Computed: true,
						},
						attributes.StructuredRenderedCommand: {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSudoCommandsBundleFetch(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := getLocalClientFromMetadata(m)
	id := d.Get(attributes.ID).(string)
	if id == "" {
		return diag.Errorf("%s cannot be blank", attributes.ID)
	}

	sudoCommandsBundle, err := c.GetSudoCommandsBundle(ctx, id)
	if err != nil {
		return diag.FromErr(err)
	}

	if sudoCommandsBundle != nil {
		d.SetId(*sudoCommandsBundle.Id)
		for key, value := range sudoCommandsBundle.ToResourceMap() {
			if err := d.Set(key, value); err != nil {
				return diag.FromErr(err)
			}
		}
	} else {
		logging.Infof("sudo commands bundle %s does not exist", id)
	}
	return nil
}
