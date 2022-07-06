package oktapam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"
)

func resourceGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGroupCreate,
		ReadContext:   resourceGroupRead,
		UpdateContext: resourceGroupUpdate,
		DeleteContext: resourceGroupDelete,
		Description:   descriptions.ResourceGroup,
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
			attributes.DeletedAt: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: descriptions.DeletedAt,
			},
			attributes.Roles: {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Description: descriptions.Roles,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceGroupCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaPAMClient)
	var roles []string
	if r, ok := d.GetOk(attributes.Roles); ok {
		rolesAttr := r.(*schema.Set)
		rolesI := rolesAttr.List()
		roles = make([]string, len(rolesI))

		for idx, ri := range rolesI {
			roles[idx] = ri.(string)
		}
	} else {
		roles = make([]string, 0)
	}

	group := client.Group{
		Name:  getStringPtr(attributes.Name, d, true),
		Roles: roles,
	}

	err := c.CreateGroup(ctx, group)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(*group.Name)
	return resourceGroupRead(ctx, d, m)
}

func resourceGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(client.OktaPAMClient)

	groupName := d.Id()
	group, err := c.GetGroup(ctx, groupName, false)
	if err != nil {
		return diag.FromErr(err)
	}

	if group != nil && utils.IsNonEmpty(group.Name) {
		if utils.IsBlank(group.DeletedAt) {
			logging.Infof("Group %s exists", *group.Name)
			d.SetId(*group.Name)
		} else {
			logging.Infof("Group %s was removed", groupName)
			d.SetId("")
		}
		for key, value := range group.ToResourceMap() {
			d.Set(key, value)
		}
	} else {
		logging.Infof("Group %s does not exist", groupName)
	}

	return diags
}

func resourceGroupUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaPAMClient)
	groupName := d.Id()

	changed := false
	updates := make(map[string]interface{})

	changeableAttributes := []string{
		attributes.Roles,
	}

	for _, attribute := range changeableAttributes {
		if d.HasChange(attribute) {
			if attribute == attributes.Roles {
				roles := d.Get(attributes.Roles).(*schema.Set)
				updates[attributes.Roles] = roles.List()
			} else {
				updates[attribute] = d.Get(attribute)
			}
			changed = true
		}
	}

	if changed {
		err := c.UpdateGroup(ctx, groupName, updates)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceGroupRead(ctx, d, m)
}

func resourceGroupDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(client.OktaPAMClient)
	groupName := d.Id()

	err := c.DeleteGroup(ctx, groupName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return diags
}
