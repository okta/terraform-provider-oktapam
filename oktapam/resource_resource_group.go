package oktapam

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
				ValidateDiagFunc: validation.ToDiagFunc(validation.All(
					validation.StringLenBetween(1, 255),
					validation.StringMatch(nameMatcher, fmt.Sprintf("value for %s may only contain alphanumeric characters (a-Z, 0-9), hyphens (-), underscores (_), and periods (.)", attributes.Name)),
				)),
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
	c := getLocalClientFromMetadata(m)

	id := d.Id()
	resourceGroup, err := c.GetResourceGroup(ctx, id)
	if err != nil {
		return diag.FromErr(err)
	}

	resourceMap := resourceGroup.ToResourceMap()
	for k, v := range resourceMap {
		if k != attributes.ID {
			if err := d.Set(k, v); err != nil {
				diags = append(diags, diag.FromErr(err)...)
			}
		}
	}

	return diags
}

func resourceResourceGroupCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := getLocalClientFromMetadata(m)

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
	c := getLocalClientFromMetadata(m)
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
	c := getLocalClientFromMetadata(m)
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
