package oktaasa

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-oktaasa/oktaasa/client"
	"github.com/terraform-providers/terraform-provider-oktaasa/oktaasa/logging"
)

func resourceGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGroupCreate,
		ReadContext:   resourceGroupRead,
		UpdateContext: resourceGroupUpdate,
		DeleteContext: resourceGroupDelete,
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
			"federated_from_team": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"federation_approved_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"deleted_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"roles": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceGroupCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaASAClient)
	roles := make([]string, 0)
	if r, ok := d.GetOk("roles"); ok {
		roles = r.([]string)
	}
	group := client.Group{
		Name:  d.Get("name").(string),
		Roles: roles,
	}

	err := c.CreateGroup(ctx, group)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(group.Name)
	return resourceGroupRead(ctx, d, m)
}

func resourceGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(client.OktaASAClient)

	groupName := d.Id()
	group, err := c.GetGroup(ctx, groupName, false)
	if err != nil {
		return diag.FromErr(err)
	}

	if group != nil && group.Name != "" {
		if group.DeletedAt == "" {
			logging.Infof("Group %s exists", group.Name)
			d.SetId(group.Name)
		} else {
			logging.Infof("Group %s was removed", groupName)
			d.SetId("")
		}
		for key, value := range group.ToMap() {
			d.Set(key, value)
		}
	} else {
		logging.Infof("Group %s does not exist", groupName)
	}

	return diags
}

func resourceGroupUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaASAClient)
	groupName := d.Id()

	changed := false
	updates := make(map[string]interface{})

	changeableAttributes := []string{
		"roles",
	}

	for _, attribute := range changeableAttributes {
		if d.HasChange(attribute) {
			if attribute == "roles" {
				roles := d.Get("roles").(*schema.Set)
				updates["roles"] = roles.List()
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
	c := m.(client.OktaASAClient)
	groupName := d.Id()

	err := c.DeleteGroup(ctx, groupName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return diags
}
