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
)

func resourceKubernetesClusterGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceKubernetesClustGroupCreate,
		ReadContext:   resourceKubernetesClusterGroupRead,
		UpdateContext: resourceKubernetesClusterGroupUpdate,
		DeleteContext: resourceKubernetesClusterGroupDelete,
		Description:   descriptions.ResourceKubernetesClusterGroup,
		Schema: map[string]*schema.Schema{
			attributes.ID: {
				Type:     schema.TypeString,
				Computed: true,
			},
			attributes.GroupID: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: descriptions.GroupID,
			},
			attributes.GroupName: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: descriptions.GroupName,
			},
			attributes.ClusterSelector: {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions.ClusterSelector,
			},
			attributes.Claims: {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: descriptions.ClusterGroupClaims,
			},
		},
	}
}

func resourceKubernetesClustGroupCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaPAMClient)
	groupName := getStringPtr(attributes.GroupName, d, true)
	clusterSelector := getStringPtr(attributes.ClusterSelector, d, true)

	claims := d.Get(attributes.Claims).(map[string]interface{})

	claimsMap := make(map[string]string, len(claims))

	for k, v := range claims {
		claimsMap[k] = fmt.Sprint(v)
	}

	clusterGroupSpec := client.KubernetesClusterGroup{
		GroupName:       groupName,
		ClusterSelector: clusterSelector,
		Claims:          claimsMap,
	}

	if createdClusterGroup, err := c.CreateKubernetesClusterGroup(ctx, clusterGroupSpec); err != nil {
		return diag.FromErr(err)
	} else if createdClusterGroup == nil {
		d.SetId("")
	} else {
		d.SetId(*createdClusterGroup.ID)
	}

	return resourceKubernetesClusterGroupRead(ctx, d, m)
}

func resourceKubernetesClusterGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaPAMClient)

	clusterGroupID := d.Id()
	clusterGroup, err := c.GetKubernetesClusterGroup(ctx, clusterGroupID)
	if err != nil {
		return diag.FromErr(err)
	}

	if clusterGroup == nil {
		logging.Debugf("kubernetes cluster group was blank")
		d.SetId("")
		return nil
	}

	for key, value := range clusterGroup.ToResourceMap() {
		logging.Debugf("setting %q to %v", key, value)
		if err := d.Set(key, value); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceKubernetesClusterGroupUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaPAMClient)
	id := d.Id()

	changed := false
	updates := make(map[string]interface{})

	changeableAttributes := []string{
		attributes.Claims,
		attributes.ClusterSelector,
	}

	for _, attribute := range changeableAttributes {
		if d.HasChange(attribute) {
			updates[attribute] = d.Get(attribute)
			changed = true
		}
	}

	if changed {
		if err := c.UpdateKubernetesClusterGroup(ctx, id, updates); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceKubernetesClusterGroupDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaPAMClient)

	id := d.Id()

	if err := c.DeleteKubernetesClusterGroup(ctx, id); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}
