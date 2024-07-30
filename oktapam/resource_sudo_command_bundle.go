package oktapam

import (
	"context"
	"fmt"

	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/client/wrappers"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"
)

var commandTypeMap = map[string]pam.CommandType{
	"raw":        pam.CommandType_RAW,
	"executable": pam.CommandType_EXECUTABLE,
	"directory":  pam.CommandType_DIRECTORY,
}

var argsTypeMap = map[string]pam.ArgsType{
	"custom": pam.ArgsType_CUSTOM,
	"any":    pam.ArgsType_ANY,
	"none":   pam.ArgsType_NONE,
}

func resourceSudoCommandBundle() *schema.Resource {
	return &schema.Resource{
		Description:   descriptions.ResourceSudoCommandsBundle,
		ReadContext:   resourceSudoCommandBundleRead,
		CreateContext: resourceSudoCommandBundleCreate,
		DeleteContext: resourceSudoCommandBundleDelete,
		UpdateContext: resourceSudoCommandBundleUpdate,
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

func resourceSudoCommandBundleRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	var diags diag.Diagnostics
	c := getSDKClientFromMetadata(m)
	sudoCommandsBundleID := d.Id()

	sudoCommandsBundle, err := client.GetSudoCommandBundle(ctx, c, sudoCommandsBundleID)
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
	} else {
		d.SetId("")
	}
	return diags
}

func readSudoCommandBundleFromResource(d *schema.ResourceData) (*pam.SudoCommandBundle, diag.Diagnostics) {
	var diags diag.Diagnostics
	var structuredCommands []pam.SudoCommandBundleStructuredCommandsInner
	if scs, ok := d.GetOk(attributes.StructuredCommands); ok {
		structuredCommandResources, ok := scs.([]any)
		if !ok {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("%s is invalid", attributes.StructuredCommands),
			})
		}
		for _, scI := range structuredCommandResources {
			sc, ok := scI.(map[string]any)
			if !ok {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  fmt.Sprintf("%s is invalid", attributes.StructuredCommands),
				})
			}

			command, ok := sc[attributes.StructuredCommand].(string)
			if !ok {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  fmt.Sprintf("%s is invalid", attributes.StructuredCommand),
				})
			}

			structuredCommand := pam.SudoCommandBundleStructuredCommandsInner{
				Command: command,
			}

			commandTypeStr, ok := sc[attributes.StructuredCommandType].(string)
			if !ok {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  fmt.Sprintf("%s is invalid", attributes.StructuredCommandType),
				})
			}

			structuredCommand.CommandType = commandTypeMap[commandTypeStr]

			if args, ok := sc[attributes.StructuredCommandArgs].(string); ok {
				structuredCommand.Args = utils.AsStringPtrZero(args, false)
			}

			argsTypeStr, ok := sc[attributes.StructuredCommandArgsType].(string)
			if !ok {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  fmt.Sprintf("%s is invalid", attributes.StructuredCommandArgsType),
				})
			}

			argsType := argsTypeMap[argsTypeStr]
			structuredCommand.ArgsType = &argsType

			structuredCommands = append(structuredCommands, structuredCommand)
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

	sudoCommandBundleID := d.Id()
	sudoCommandBundle := &pam.SudoCommandBundle{
		Id:                 &sudoCommandBundleID,
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

func resourceSudoCommandBundleCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := getSDKClientFromMetadata(m)
	sudoCommandBundle, diags := readSudoCommandBundleFromResource(d)
	if diags != nil {
		return diags
	}

	if err := client.CreateSudoCommandBundle(ctx, c, sudoCommandBundle); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*sudoCommandBundle.Id)

	return resourceSudoCommandBundleRead(ctx, d, m)
}

func resourceSudoCommandBundleUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := getSDKClientFromMetadata(m)

	sudoCommandBundle, diags := readSudoCommandBundleFromResource(d)
	if diags != nil {
		return diags
	}

	if err := client.UpdateSudoCommandBundle(ctx, c, sudoCommandBundle); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*sudoCommandBundle.Id)
	return resourceSudoCommandBundleRead(ctx, d, m)
}

func resourceSudoCommandBundleDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	// there isn't a true delete with password settings, and as databases don't specify ManagedPrivilegedAccounts,
	// the best we can do is to set rotation to false and some default required values.
	c := getSDKClientFromMetadata(m)
	sudoCommandsBundleID := d.Id()
	if err := client.DeleteSudoCommandBundle(ctx, c, sudoCommandsBundleID); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
