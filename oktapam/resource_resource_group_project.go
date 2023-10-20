package oktapam

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/typed_strings"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"
)

func resourceResourceGroupProject() *schema.Resource {
	return &schema.Resource{
		Description:   descriptions.ResourceResourceGroupProject,
		CreateContext: resourceResourceGroupProjectCreate,
		ReadContext:   resourceResourceGroupProjectRead,
		UpdateContext: resourceResourceGroupProjectUpdate,
		DeleteContext: resourceResourceGroupProjectDelete,
		Schema: map[string]*schema.Schema{
			attributes.ID: {
				Type:     schema.TypeString,
				Computed: true,
				// Description is autogenerated
			},
			attributes.Name: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: descriptions.Name,
			},
			attributes.ResourceGroup: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: descriptions.ResourceGroupID,
			},
			attributes.Team: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: descriptions.TeamName,
			},
			attributes.DeletedAt: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: descriptions.DeletedAt,
			},
			attributes.GatewaySelector: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions.GatewaySelector,
			},
			attributes.SSHCertificateType: {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "CERT_TYPE_ED25519_01",
				Description: descriptions.SSHCertificateType,
				ValidateDiagFunc: func(i any, p cty.Path) diag.Diagnostics {
					var diags diag.Diagnostics
					s := i.(string)

					for _, v := range typed_strings.ValidCertTypes {
						if v == s {
							if s == typed_strings.CertTypeRsa {
								diag := diag.Diagnostic{
									Severity: diag.Warning,
									Summary:  "deprecated value",
									Detail: fmt.Sprintf("%s is a deprecated key algorithm type. This option should only be used to connect to legacy systems that cannot use newer SSH versions. "+
										"If you do need to use %s, it is recommended to connect via a gateway with traffic forwarding. "+
										"Otherwise, please use a more current key algorithm. ", typed_strings.CertTypeRsa, typed_strings.CertTypeRsa),
								}
								return append(diags, diag)
							}
							return nil
						}
					}

					diag := diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "invalid value",
						Detail:   fmt.Sprintf("%q is not one of %q", s, typed_strings.ValidCertTypes),
					}

					return append(diags, diag)
				},
			},
			attributes.AccountDiscovery: {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: descriptions.AccountDiscovery,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: resourceResourceGroupProjectReadImport,
		},
	}
}

func parseResourceGroupProjectID(resourceId string) (string, string, error) {
	split := strings.Split(resourceId, "/")
	if len(split) != 2 {
		return "", "", fmt.Errorf("expected format: <resource_group_id>/<id>, received: %s", resourceId)
	}
	return split[0], split[1], nil
}

func resourceResourceGroupProjectCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := getLocalClientFromMetadata(m)

	project := client.ResourceGroupProject{
		Name:               GetStringPtrFromResource(attributes.Name, d, true),
		ResourceGroupID:    GetStringPtrFromResource(attributes.ResourceGroup, d, true),
		GatewaySelector:    GetStringPtrFromResource(attributes.GatewaySelector, d, false),
		SSHCertificateType: GetStringPtrFromResource(attributes.SSHCertificateType, d, false),
		AccountDiscovery:   GetBoolPtrFromResource(attributes.AccountDiscovery, d, false),
	}

	resultingProject, err := c.CreateResourceGroupProject(ctx, project)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*resultingProject.ID)
	return resourceResourceGroupProjectRead(ctx, d, m)
}

func resourceResourceGroupProjectReadWithIgnorable(ctx context.Context, d *schema.ResourceData, m any, ignoreValues bool) (*schema.ResourceData, diag.Diagnostics) {
	var diags diag.Diagnostics
	c := getLocalClientFromMetadata(m)

	projectID := d.Id()
	resourceGroupID := d.Get(attributes.ResourceGroup).(string)

	proj, err := c.GetResourceGroupProject(ctx, resourceGroupID, projectID, false)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	// these values are updated in normal operation of PAM and should not be reset to
	// the values that were used to create the project
	ignorableValues := map[string]bool{attributes.NextUnixGID: true, attributes.NextUnixUID: true}

	if proj != nil && utils.IsNonEmpty(proj.Name) {
		if utils.IsBlank(proj.DeletedAt) {
			logging.Infof("Project %s exists", proj.Name)
			d.SetId(*proj.ID)
		} else {
			logging.Infof("Project %s was removed", projectID)
			d.SetId("")
		}
		for key, value := range proj.ToResourceMap() {
			if _, ok := ignorableValues[key]; !ignoreValues || !ok {
				if err := d.Set(key, value); err != nil {
					diags = append(diags, diag.FromErr(err)...)
				}
			}
		}
	} else {
		logging.Infof("Project %s does not exist", projectID)
		d.SetId("")
	}

	return d, diags
}

func resourceResourceGroupProjectReadImport(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	// d.Id() here is the last argument passed to the `terraform import RESOURCE_TYPE.RESOURCE_NAME RESOURCE_ID` command
	// Id provided for import is in the format <resource_group_id>/<project_id>
	resourceGroupID, projectID, err := parseResourceGroupProjectID(d.Id())
	if err != nil {
		return nil, err
	}
	if err := d.Set(attributes.ResourceGroup, resourceGroupID); err != nil {
		return nil, err
	}
	d.SetId(projectID)

	projectResource, diags := resourceResourceGroupProjectReadWithIgnorable(ctx, d, m, false)
	for _, d := range diags {
		if d.Severity == diag.Error {
			return nil, fmt.Errorf(d.Summary)
		}
	}
	return []*schema.ResourceData{projectResource}, nil
}

func resourceResourceGroupProjectRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	_, diags := resourceResourceGroupProjectReadWithIgnorable(ctx, d, m, true)
	return diags
}

func resourceResourceGroupProjectUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := getLocalClientFromMetadata(m)
	projectID := d.Id()
	resourceGroupID := d.Get(attributes.ResourceGroup).(string)

	changed := false
	updates := make(map[string]any)
	attributeMapping := map[string]string{
		attributes.AccountDiscovery: "server_account_management",
	}

	changeableAttributes := []string{
		attributes.GatewaySelector,
		attributes.SSHCertificateType,
		attributes.AccountDiscovery,
	}

	for _, attribute := range changeableAttributes {
		if d.HasChange(attribute) {
			param := attribute
			if m, ok := attributeMapping[attribute]; ok {
				param = m
			}
			updates[param] = d.Get(attribute)
			changed = true
		}
	}

	if changed {
		err := c.UpdateResourceGroupProject(ctx, resourceGroupID, projectID, updates)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceResourceGroupProjectRead(ctx, d, m)
}

func resourceResourceGroupProjectDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := getLocalClientFromMetadata(m)
	projectID := d.Id()
	resourceGroupID := d.Get(attributes.ResourceGroup).(string)

	err := c.DeleteResourceGroupProject(ctx, resourceGroupID, projectID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}
