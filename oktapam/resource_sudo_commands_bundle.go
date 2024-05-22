package oktapam

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"
)

func resourceSudoCommandsBundle() *schema.Resource {
	return &schema.Resource{
		Description:   descriptions.ResourceSudoCommandsBundle,
		ReadContext:   resourceSudoCommandsBundleRead,
		CreateContext: resourceSudoCommandsBundleCreate,
		DeleteContext: resourceSudoCommandsBundleDelete,
		UpdateContext: resourceSudoCommandsBundleUpdate,
		Schema: map[string]*schema.Schema{
			attributes.ID: {
				Type:     schema.TypeString,
				Required: true,
			},
			attributes.Name: {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions.Name,
			},
			attributes.RunAs: {
				Type:     schema.TypeString,
				Optional: true,
			},
			attributes.NoPasswd: {
				Type:     schema.TypeBool,
				Optional: true,
			},
			attributes.NoExec: {
				Type:     schema.TypeString,
				Optional: true,
			},
			attributes.SetEnv: {
				Type:     schema.TypeString,
				Optional: true,
			},
			attributes.AddEnv: {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			attributes.SubEnv: {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			attributes.StructuredCommands: {
				Type:        schema.TypeList,
				MinItems:    0,
				MaxItems:    64,
				Required:    true,
				ForceNew:    true,
				Description: descriptions.StructuredCommands,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						attributes.StructuredCommand: {
							Type:     schema.TypeString,
							Required: true,
						},
						attributes.StructuredCommandType: {
							Type:     schema.TypeString,
							Required: true,
						},
						attributes.StructuredCommandArgsType: {
							Type:     schema.TypeString,
							Optional: true,
						},
						attributes.StructuredCommandArgs: {
							Type:     schema.TypeString,
							Optional: true,
						},
						attributes.StructuredRenderedCommand: {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceSudoCommandsBundleRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := getLocalClientFromMetadata(m)
	sudoCommandsBundleID := d.Id()
	sudoCommandsBundle, err := c.GetSudoCommandsBundle(ctx, sudoCommandsBundleID)

	if err != nil {
		return diag.FromErr(err)
	}

	if sudoCommandsBundle == nil || utils.IsBlank(sudoCommandsBundle.Id) {
		d.SetId("")
		return diag.Errorf("sudo commands bundle does not exist")
	}

	for key, value := range sudoCommandsBundle.ToResourceMap() {
		logging.Debugf("setting %s to %v", key, value)
		if err := d.Set(key, value); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func readSudoCommandsBundleFromResource(d *schema.ResourceData) (client.SudoCommandsBundle, diag.Diagnostics) {
	var diags diag.Diagnostics
	var structuredCommands []client.StructuredCommand
	if sc, ok := d.GetOk(attributes.StructuredCommands); ok {
		structuredCommand, ok := sc.(client.StructuredCommand)
		if !ok {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("%s is invalid", attributes.StructuredCommands),
			})
		}
		structuredCommands = append(structuredCommands, structuredCommand)
	}

	addEnv, addEnvDiag := GetStringSliceFromResource(attributes.AddEnv, d, true)
	if addEnvDiag != nil {
		diags = append(diags, addEnvDiag...)
	}

	subEnv, subEnvDiag := GetStringSliceFromResource(attributes.SubEnv, d, true)
	if addEnvDiag != nil {
		diags = append(diags, subEnvDiag...)
	}

	resourceGroup := client.SudoCommandsBundle{
		Name:               GetStringPtrFromResource(attributes.Name, d, true),
		RunAs:              GetStringPtrFromResource(attributes.RunAs, d, false),
		NoPasswd:           GetBoolPtrFromResource(attributes.NoPasswd, d, false),
		NoExec:             GetBoolPtrFromResource(attributes.NoExec, d, false),
		SetEnv:             GetBoolPtrFromResource(attributes.SetEnv, d, false),
		AddEnv:             addEnv,
		SubEnv:             subEnv,
		StructuredCommands: structuredCommands,
	}

	if diags != nil {
		return client.SudoCommandsBundle{}, diags
	}

	return resourceGroup, nil
}

func resourceSudoCommandsBundleCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := getLocalClientFromMetadata(m)

	sudoCommandBundle, diags := readSudoCommandsBundleFromResource(d)
	if diags != nil {
		return diags
	}

	result, err := c.CreateSudoCommandsBundle(ctx, sudoCommandBundle)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(*result.Id)

	return resourceSudoCommandsBundleRead(ctx, d, m)
}

func resourceSudoCommandsBundleUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := getLocalClientFromMetadata(m)
	id := d.Id()

	changed := false
	updates := make(map[string]any)

	changeableAttributes := []string{
		attributes.Name,
		attributes.RunAs,
		attributes.NoExec,
		attributes.NoPasswd,
		attributes.StructuredCommands,
		attributes.SetEnv,
		attributes.AddEnv,
		attributes.SubEnv,
	}

	for _, attribute := range changeableAttributes {
		if d.HasChange(attribute) {
			updates[attribute] = d.Get(attribute)
			changed = true
		}
	}

	if changed {
		if err := c.UpdateSudoCommandsBundle(ctx, id, updates); err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceSudoCommandsBundleRead(ctx, d, m)
}

func resourceSudoCommandsBundleDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := getLocalClientFromMetadata(m)
	id := d.Id()

	err := c.DeleteSudoCommandsBundle(ctx, id)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
