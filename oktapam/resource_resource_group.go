package oktapam

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
)

func resourceResourceGroup() *schema.Resource {
	return &schema.Resource{
		Description:   descriptions.ResourceResourceGroup,
		CreateContext: resourceResourceGroupCreate,
		ReadContext:   resourceResourceGroupRead,
		DeleteContext: resourceResourceGroupDelete,
		UpdateContext: resourceResourceGroupUpdate,
		Schema: map[string]*schema.Schema{
			attributes.ID: {
				Type:     schema.TypeString,
				Computed: true,
				// Description is autogenerated
			},
			attributes.Name: {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions.Name,
				ValidateDiagFunc: func(i any, p cty.Path) diag.Diagnostics {
					var diags diag.Diagnostics
					if s, ok := i.(string); ok {
						if s == "" {
							diags = append(diags, diag.Diagnostic{
								Severity: diag.Error,
								Summary:  fmt.Sprintf("value for %s must not be empty", attributes.Name),
							})
						} else if len(s) > 255 {
							diags = append(diags, diag.Diagnostic{
								Severity: diag.Error,
								Summary:  fmt.Sprintf("value for %s must be between 1 and 255 characters, inclusive", attributes.Name),
							})
						} else if !MatchesSimpleName(s) {
							diags = append(diags, diag.Diagnostic{
								Severity: diag.Error,
								Summary:  fmt.Sprintf("value for %s may only contain alphanumeric characters (a-Z, 0-9), hyphens (-), underscores (_), and periods (.)", attributes.Name),
							})
						}
					} else {
						diags = append(diags, diag.Diagnostic{
							Severity: diag.Error,
							Summary:  fmt.Sprintf("value for %s must be a string", attributes.Name),
						})
					}

					return diags
				},
			},
			attributes.Description: {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions.Description,
			},
			attributes.DelegatedResourceAdminGroups: {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required:    true,
				Description: descriptions.DelegatedAdminGroups,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceResourceGroupRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(client.OktaPAMClient)

	id := d.Id()
	resourceGroup, err := c.GetResourceGroup(ctx, id)
	if err != nil {
		return diag.FromErr(err)
	}

	resourceMap := resourceGroup.ToResourceMap()
	for k, v := range resourceMap {
		if k == attributes.ID {
			d.SetId(v.(string))
		} else {
			if err := d.Set(k, v); err != nil {
				diags = append(diags, diag.FromErr(err)...)
			}
		}
	}

	return diags
}

func resourceResourceGroupCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := m.(client.OktaPAMClient)

	resourceGroup, diags := readResourceGroupFromResource(d)
	if diags != nil {
		return diags
	}

	result, err := c.CreateResourceGroup(ctx, resourceGroup)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(*result.ID)

	return resourceResourceGroupRead(ctx, d, m)
}

func resourceResourceGroupUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := m.(client.OktaPAMClient)
	id := d.Id()
	if id == "" {
		return diag.Errorf("could not obtain resource group id from resource")
	}
	resourceGroup, diags := readResourceGroupFromResource(d)
	if diags != nil {
		return diags
	}

	if err := c.UpdateResourceGroup(ctx, id, resourceGroup); err != nil {
		return diag.FromErr(err)
	}

	return resourceResourceGroupRead(ctx, d, m)
}

func resourceResourceGroupDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(client.OktaPAMClient)
	id := d.Id()
	if err := c.DeleteResourceGroup(ctx, id); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	} else {
		d.SetId("")
	}

	return diags
}

func readResourceGroupFromResource(d *schema.ResourceData) (client.ResourceGroup, diag.Diagnostics) {
	var groupIds []client.NamedObject
	if g, ok := d.GetOk(attributes.DelegatedResourceAdminGroups); ok {
		if groupsArr, diags := GetUUIDSlice(g, attributes.DelegatedResourceAdminGroups); diags == nil {
			groupIds = ConvertToNamedObjectSlice(groupsArr, client.UserGroupNamedObjectType)
		} else {
			return client.ResourceGroup{}, diags
		}
	} else {
		groupIds = make([]client.NamedObject, 0)
	}

	resourceGroup := client.ResourceGroup{
		Name:                         GetStringPtrFromResource(attributes.Name, d, true),
		Description:                  GetStringPtrFromResource(attributes.Description, d, false),
		DelegatedResourceAdminGroups: groupIds,
	}

	return resourceGroup, nil
}
