package oktapam

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
)

func resourceKubernetesClusterConnection() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceKubernetesClusterConnectionCreate,
		ReadContext:   resourceKubernetesClusterConnectionRead,
		DeleteContext: resourceKubernetesClusterConnectionDelete,
		UpdateContext: resourceKubernetesClusterConnectionUpdate,
		Description:   descriptions.ResourceKubernetesClusterConnection,
		Schema: map[string]*schema.Schema{
			attributes.ID: { //TODO(ja) fix me - clusterid vs id
				Type:     schema.TypeString,
				Computed: true,
			},
			attributes.ClusterID: {
				Type:        schema.TypeString,
				Computed:    false,
				Required:    true,
				ForceNew:    true,
				Description: descriptions.KubernetesClusterID,
			},
			attributes.KubernetesAPIURL: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions.KubernetesAPIURL,
			},
			attributes.PublicCertificate: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions.KubernetesPublicCertificate,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceKubernetesClusterConnectionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaPAMClient)

	clusterID := d.Get(attributes.ClusterID).(string)
	apiURL := getStringPtr(attributes.KubernetesAPIURL, d, true)
	publicCertificate := getStringPtr(attributes.PublicCertificate, d, true)

	connectionSpec := client.KubernetesClusterConnection{
		APIURL:            apiURL,
		PublicCertificate: publicCertificate,
	}

	if createdConnection, err := c.CreateKubernetesClusterConnection(ctx, clusterID, connectionSpec); err != nil {
		return diag.FromErr(err)
	} else if createdConnection == nil {
		d.SetId("")
	} else {
		d.SetId(*createdConnection.ID)
	}

	return resourceKubernetesClusterConnectionRead(ctx, d, m)
}

func resourceKubernetesClusterConnectionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaPAMClient)

	clusterID := d.Id()
	cluster, err := c.GetKubernetesClusterConnection(ctx, clusterID)
	if err != nil {
		return diag.FromErr(err)
	}

	if cluster == nil {
		logging.Debugf("kubernetes cluster connection was blank")
		d.SetId("")
		return nil
	}

	for key, value := range cluster.ToResourceMap() {
		logging.Debugf("setting %q to %v", key, value)
		if err := d.Set(key, value); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceKubernetesClusterConnectionUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaPAMClient)
	id := d.Id()

	changed := false
	updates := make(map[string]interface{})

	changeableAttributes := []string{
		attributes.KubernetesAPIURL,
		attributes.PublicCertificate,
	}

	for _, attribute := range changeableAttributes {
		if d.HasChange(attribute) {
			updates[attribute] = d.Get(attribute)
			changed = true
		}
	}

	if changed {
		if err := c.UpdateKubernetesClusterConnection(ctx, id, updates); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceKubernetesClusterConnectionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaPAMClient)

	id := d.Id()

	if err := c.DeleteKubernetesClusterConnection(ctx, id); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}
