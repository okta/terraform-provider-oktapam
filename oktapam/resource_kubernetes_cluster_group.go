package oktapam

import (
	"context"
	"fmt"
	"strings"

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
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceKubernetesClustGroupCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaPAMClient)
	groupName := getStringPtr(attributes.GroupName, d, true)
	clusterSelector := getStringPtr(attributes.ClusterSelector, d, true)

	clusterGroupSpec := client.KubernetesClusterGroup{
		GroupName:       groupName,
		ClusterSelector: clusterSelector,
		Claims:          claimsCSVToMap(d.Get(attributes.Claims).(map[string]interface{})),
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

	clusterGroup, err := c.GetKubernetesClusterGroup(ctx, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	if clusterGroup == nil {
		logging.Debugf("kubernetes cluster group was blank")
		d.SetId("")
		return nil
	}

	for key, value := range clusterGroup.ToResourceMap() {
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
			switch attribute {
			case attributes.Claims:
				updates[attribute] = claimsCSVToMap(d.Get(attributes.Claims).(map[string]interface{}))
			default:
				updates[attribute] = d.Get(attribute)
			}
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

	if err := c.DeleteKubernetesClusterGroup(ctx, d.Id()); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}

func claimsCSVToMap(claimsIn map[string]interface{}) map[string][]string {
	claimsMap := make(map[string][]string)

	for k, csvValues := range claimsIn {
		var values []string
		for _, value := range strings.Split(fmt.Sprint(csvValues), ",") {
			values = append(values, value)
		}
		claimsMap[k] = values
	}

	return claimsMap
}
