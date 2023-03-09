package oktapam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceADConnections() *schema.Resource {
	return &schema.Resource{
		Description: descriptions.SourceADConnections,
		ReadContext: dataSourceADConnectionsRead,
		Schema: map[string]*schema.Schema{
			// Query parameter values
			attributes.GatewayID: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions.FilterGatewayID,
			},
			attributes.CertificateID: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions.FilterCertificateID,
			},
			attributes.IncludeCertDetails: {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: descriptions.FilterIncludeCertDetails,
			},
			// Return value
			attributes.ADConnections: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: descriptions.SourceADConnections,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						attributes.Name: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: descriptions.Name,
						},
						attributes.ID: {
							Type:     schema.TypeString,
							Computed: true,
							// Description is autogenerated
						},
						attributes.GatewayID: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: descriptions.GatewayID,
						},
						attributes.Domain: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: descriptions.Domain,
						},
						attributes.CertificateID: {
							Type:        schema.TypeString,
							Optional:    true, //Only in case of passwordless
							Computed:    true,
							Description: descriptions.CertificateID,
						},
						attributes.DomainControllers: {
							Type: schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Computed:    true,
							Optional:    true,
							Description: descriptions.DomainControllers,
						},
						attributes.DeletedAt: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: descriptions.DeletedAt,
						},
					},
				},
			},
		},
	}
}

func dataSourceADConnectionsRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := m.(client.OktaPAMClient)
	parameters := client.ListADConnectionsParameters{}

	//Extract parameters
	if gatewayID, ok := d.GetOk(attributes.GatewayID); ok {
		parameters.GatewayID = gatewayID.(string)
	}
	if certificateID, ok := d.GetOk(attributes.CertificateID); ok {
		parameters.CertificateID = certificateID.(string)
	}

	self, err := getOkBool(attributes.IncludeCertDetails, d)
	if err != nil {
		return diag.FromErr(err)
	}
	parameters.IncludeCertDetails = self

	adConnectionsList, err := c.ListADConnections(ctx, parameters)
	if err != nil {
		return diag.FromErr(err)
	}

	adConnections := make([]map[string]any, len(adConnectionsList))
	for idx, adConnection := range adConnectionsList {
		adConnections[idx] = adConnection.ToResourceMap()
	}

	if err := d.Set(attributes.ADConnections, adConnections); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resource.UniqueId())
	return nil
}
