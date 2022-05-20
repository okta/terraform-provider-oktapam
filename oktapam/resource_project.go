package oktapam

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/client"
	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/constants/descriptions"
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
			attributes.Name: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: descriptions.Name,
			},
			attributes.ID: {
				Type:     schema.TypeString,
				Computed: true,
				// Description is autogenerated
			},
			attributes.Team: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: descriptions.TeamName,
			},
			attributes.NextUnixGID: {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     63001,
				Description: descriptions.NextUnixUID,
			},
			attributes.NextUnixUID: {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     60001,
				Description: descriptions.NextUnixUID,
			},
			attributes.CreateServerUsers: {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: descriptions.CreateServerUsers,
			},
			attributes.DeletedAt: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: descriptions.DeletedAt,
			},
			attributes.ForwardTraffic: {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: descriptions.ForwardTraffic,
			},
			attributes.RDPSessionRecording: {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: descriptions.RDPSessionRecording,
			},
			attributes.RequirePreauthForCreds: {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: descriptions.RequirePreauthForCreds,
			},
			attributes.SSHSessionRecording: {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: descriptions.SSHSessionRecording,
			},
			attributes.GatewaySelector: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions.GatewaySelector,
			},
			attributes.SSHCertificateType: {
				Type:        schema.TypeString,
				Optional:    true, // The default value is `CERT_TYPE_ED25519_01`
				Computed:    true,
				Description: descriptions.SSHCertificateType,
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
									Detail:   "CERT_TYPE_RSA_01 is a deprecated key algorithm type. This option should only be used to connect to legacy systems that cannot use newer SSH versions. If you do need to use CERT_TYPE_RSA_01, it is recommended to connect via a gateway with traffic forwarding. Otherwise, please use a more current key algorithm. ",
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
		Name:                   getStringPtr(attributes.Name, d, true),
		NextUnixGID:            getIntPtr(attributes.NextUnixGID, d, false),
		NextUnixUID:            getIntPtr(attributes.NextUnixUID, d, false),
		CreateServerUsers:      getBoolPtr(attributes.CreateServerUsers, d, false),
		ForwardTraffic:         getBoolPtr(attributes.ForwardTraffic, d, false),
		RDPSessionRecording:    getBoolPtr(attributes.RDPSessionRecording, d, false),
		RequirePreAuthForCreds: getBoolPtr(attributes.RequirePreauthForCreds, d, false),
		SSHSessionRecording:    getBoolPtr(attributes.SSHSessionRecording, d, false),
		GatewaySelector:        getStringPtr(attributes.GatewaySelector, d, false),
		SSHCertificateType:     getStringPtr(attributes.SSHCertificateType, d, false),
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
	ignorableValues := map[string]bool{attributes.NextUnixGID: true, attributes.NextUnixUID: true}

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
		attributes.NextUnixGID,
		attributes.NextUnixUID,
		attributes.CreateServerUsers,
		attributes.ForwardTraffic,
		attributes.RDPSessionRecording,
		attributes.RequirePreauthForCreds,
		attributes.SSHSessionRecording,
		attributes.GatewaySelector,
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
