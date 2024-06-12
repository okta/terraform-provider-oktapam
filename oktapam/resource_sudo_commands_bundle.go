package oktapam

import (
	"context"
	"fmt"

	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mitchellh/mapstructure"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/client/wrappers"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
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
				Computed: true,
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
				Type:     schema.TypeBool,
				Optional: true,
			},
			attributes.SetEnv: {
				Type:     schema.TypeBool,
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
	var diags diag.Diagnostics
	c := getSDKClientFromMetadata(m)
	sudoCommandsBundleID := d.Id()

	sudoCommandsBundle, err := client.GetSudoCommandsBundle(ctx, c, sudoCommandsBundleID)
	if err != nil {
		return diag.FromErr(err)
	}
	if sudoCommandsBundle != nil {
		wrap := wrappers.SudoCommandsBundleWrapper{SudoCommandBundle: sudoCommandsBundle}
		for k, v := range wrap.ToResourceMap() {
			if err := d.Set(k, v); err != nil {
				diags = append(diags, diag.FromErr(err)...)
			}
		}
	} else {
		d.SetId("")
	}
	return diags
}

func readSudoCommandsBundleFromResource(d *schema.ResourceData) (*pam.SudoCommandBundle, diag.Diagnostics) {
	var diags diag.Diagnostics
	var structuredCommands []pam.SudoCommandBundleStructuredCommandsInner
	if scs, ok := d.GetOk(attributes.StructuredCommands); ok {
		structuredCommandResources, ok := scs.([]interface{})
		if !ok {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("%s is invalid", attributes.StructuredCommands),
			})
		}
		for _, sc := range structuredCommandResources {
			var structuredCommandResource pam.SudoCommandBundleStructuredCommandsInner
			cfg := &mapstructure.DecoderConfig{
				Metadata: nil,
				Result:   &structuredCommandResource,
				TagName:  "json",
			}
			decoder, dErr := mapstructure.NewDecoder(cfg)
			if dErr != nil {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  fmt.Sprintf("error creating decoder for %s", attributes.StructuredCommands),
				})
				continue
			}
			if dErr := decoder.Decode(sc); dErr != nil {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  fmt.Sprintf("error decoding for %s", attributes.StructuredCommands),
				})
				continue
			}
			structuredCommands = append(structuredCommands, structuredCommandResource)
		}
	}

	addEnv, addEnvDiag := GetStringSliceFromResource(attributes.AddEnv, d, true)
	if addEnvDiag != nil {
		diags = append(diags, addEnvDiag...)
	}

	subEnv, subEnvDiag := GetStringSliceFromResource(attributes.SubEnv, d, true)
	if addEnvDiag != nil {
		diags = append(diags, subEnvDiag...)
	}

	sudoCommandBundle := &pam.SudoCommandBundle{
		Name:               *GetStringPtrFromResource(attributes.Name, d, true),
		RunAs:              *pam.NewNullableString(GetStringPtrFromResource(attributes.RunAs, d, false)),
		NoPasswd:           *pam.NewNullableBool(GetBoolPtrFromResource(attributes.NoPasswd, d, false)),
		NoExec:             *pam.NewNullableBool(GetBoolPtrFromResource(attributes.NoExec, d, false)),
		SetEnv:             *pam.NewNullableBool(GetBoolPtrFromResource(attributes.SetEnv, d, false)),
		AddEnv:             addEnv,
		SubEnv:             subEnv,
		StructuredCommands: structuredCommands,
	}

	if diags != nil {
		return nil, diags
	}

	return sudoCommandBundle, nil
}

func resourceSudoCommandsBundleCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := getSDKClientFromMetadata(m)
	sudoCommandBundle, diags := readSudoCommandsBundleFromResource(d)
	if diags != nil {
		return diags
	}

	if err := client.CreateSudoCommandsBundle(ctx, c, sudoCommandBundle); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*sudoCommandBundle.Id)

	return resourceSudoCommandsBundleRead(ctx, d, m)
}

func resourceSudoCommandsBundleUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := getSDKClientFromMetadata(m)

	sudoCommandBundle, diags := readSudoCommandsBundleFromResource(d)
	if diags != nil {
		return diags
	}

	if err := client.UpdateSudoCommandsBundle(ctx, c, sudoCommandBundle); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*sudoCommandBundle.Id)
	return resourceSudoCommandsBundleRead(ctx, d, m)
}

func resourceSudoCommandsBundleDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	// there isn't a true delete with password settings, and as databases don't specify ManagedPrivilegedAccounts,
	// the best we can do is to set rotation to false and some default required values.
	c := getSDKClientFromMetadata(m)
	sudoCommandsBundleID := d.Id()
	if err := client.DeleteSudoCommandsBundle(ctx, c, sudoCommandsBundleID); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
