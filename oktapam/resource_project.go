package oktapam

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/client"
	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/logging"
	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/utils"
)

func resourceProject() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceProjectCreate,
		ReadContext:   resourceProjectRead,
		UpdateContext: resourceProjectUpdate,
		DeleteContext: resourceProjectDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"team": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"next_unix_gid": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  63001,
			},
			"next_unix_uid": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60001,
			},
			"create_server_users": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"deleted_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"forward_traffic": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"rdp_session_recording": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"require_preauth_for_creds": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"ssh_session_recording": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"gateway_selector": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssh_certificate_type": {
				Type:     schema.TypeString,
				Optional: true, // this is optional at this point since it is behind a feature flag.  if/when this changes, this should become required
				Computed: true,
				ValidateDiagFunc: func(i interface{}, p cty.Path) diag.Diagnostics {
					var diags diag.Diagnostics
					s := i.(string)

					valid := []string{"CERT_TYPE_RSA_01", "CERT_TYPE_RSA_SHA2_256_01", "CERT_TYPE_RSA_SHA2_512_01", "CERT_TYPE_ED25519_01"}

					for _, v := range valid {
						if v == s {
							if s == "CERT_TYPE_RSA_01" {
								diag := diag.Diagnostic{
									Severity: diag.Warning,
									Summary:  "deprecated value",
									Detail:   "CERT_TYPE_RSA_01 is a deprecated key algorithm type and should only be used for compatibility purposes.  For new projects, please use a more current key algorithm",
								}
								return append(diags, diag)
							}
							return nil
						}
					}

					diag := diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "invalid value",
						Detail:   fmt.Sprintf("%q is not one of %q", s, valid),
					}

					return append(diags, diag)
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: resourceProjectReadImport,
		},
	}
}

func resourceProjectCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaPAMClient)

	project := client.Project{
		Name:                   getStringPtr("name", d, true),
		NextUnixGID:            getIntPtr("next_unix_gid", d, false),
		NextUnixUID:            getIntPtr("next_unix_uid", d, false),
		CreateServerUsers:      getBoolPtr("create_server_users", d, false),
		ForwardTraffic:         getBoolPtr("forward_traffic", d, false),
		RDPSessionRecording:    getBoolPtr("rdp_session_recording", d, false),
		RequirePreAuthForCreds: getBoolPtr("require_preauth_for_creds", d, false),
		SSHSessionRecording:    getBoolPtr("ssh_session_recording", d, false),
		GatewaySelector:        getStringPtr("gateway_selector", d, false),
		SSHCertificateType:     getStringPtr("ssh_certificate_type", d, false),
	}

	err := c.CreateProject(ctx, project)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*project.Name)
	return resourceProjectRead(ctx, d, m)
}

func resourceProjectReadWithIgnorable(ctx context.Context, d *schema.ResourceData, m interface{}, ignoreValues bool) (*schema.ResourceData, error) {
	c := m.(client.OktaPAMClient)

	projectName := d.Id()
	proj, err := c.GetProject(ctx, projectName, false)
	if err != nil {
		return nil, err
	}

	// these values are updated in normal operation of PAM and should not be reset to
	// the values that were used to create the project
	ignorableValues := map[string]bool{"next_unix_gid": true, "next_unix_uid": true}

	if proj != nil && utils.IsNonEmpty(proj.Name) {
		if utils.IsBlank(proj.DeletedAt) {
			logging.Infof("Project %s exists", proj.Name)
			d.SetId(*proj.Name)
		} else {
			logging.Infof("Project %s was removed", projectName)
			d.SetId("")
		}
		for key, value := range proj.ToResourceMap() {
			if _, ok := ignorableValues[key]; !ignoreValues || !ok {
				d.Set(key, value)
			}
		}
	} else {
		logging.Infof("Project %s does not exist", projectName)
		d.SetId("")
	}

	return d, nil
}

func resourceProjectReadImport(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	projectResource, err := resourceProjectReadWithIgnorable(ctx, d, m, false)
	if err != nil {
		return nil, err
	}
	return []*schema.ResourceData{projectResource}, nil
}

func resourceProjectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	_, err := resourceProjectReadWithIgnorable(ctx, d, m, true)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceProjectUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaPAMClient)
	projectName := d.Id()

	changed := false
	updates := make(map[string]interface{})

	changeableAttributes := []string{
		"next_unix_gid",
		"next_unix_uid",
		"create_server_users",
		"forward_traffic",
		"rdp_session_recording",
		"require_preauth_for_creds",
		"ssh_session_recording",
		"gateway_selector",
	}

	for _, attribute := range changeableAttributes {
		if d.HasChange(attribute) {
			updates[attribute] = d.Get(attribute)
			changed = true
		}
	}

	if changed {
		err := c.UpdateProject(ctx, projectName, updates)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceProjectRead(ctx, d, m)
}

func resourceProjectDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(client.OktaPAMClient)
	projectName := d.Id()

	err := c.DeleteProject(ctx, projectName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return diags
}
