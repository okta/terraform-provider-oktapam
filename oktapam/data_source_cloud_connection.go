package oktapam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
)

func dataSourceCloudConnection() *schema.Resource {
	return &schema.Resource{
		Description: descriptions.SourceCloudConnection,
		ReadContext: dataSourceCloudConnectionFetch,
		Schema: map[string]*schema.Schema{
			attributes.ID: {
				Type:     schema.TypeString,
				Required: true,
			},
			attributes.Name: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: descriptions.Name,
			},
			attributes.CloudConnectionDetails: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: descriptions.CloudConnectionDetails,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						attributes.AWS: {
							Type:        schema.TypeList,
							Computed:    true,
							Description: descriptions.CloudConnectionDetailsAWS,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									attributes.CloudConnectionAccountId: {
										Type:     schema.TypeString,
										Computed: true,
									},
									attributes.CloudConnectionExternalId: {
										Type:     schema.TypeString,
										Computed: true,
									},
									attributes.CloudConnectionRoleARN: {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceCloudConnectionFetch(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := getLocalClientFromMetadata(m)
	id := d.Get(attributes.ID).(string)
	if id == "" {
		return diag.Errorf("%s cannot be blank", attributes.ID)
	}

	cloudConnection, err := c.GetCloudConnection(ctx, id)
	if err != nil {
		return diag.FromErr(err)
	}

	if cloudConnection != nil {
		d.SetId(*cloudConnection.ID)
		for key, value := range cloudConnection.ToResourceMap() {
			if err := d.Set(key, value); err != nil {
				return diag.FromErr(err)
			}
		}
	} else {
		logging.Infof("cloud connection %s does not exist", id)
	}
	return nil
}
