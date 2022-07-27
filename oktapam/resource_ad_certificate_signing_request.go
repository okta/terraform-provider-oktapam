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

func resourceADCertificateSigningRequest() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceADCertificateSigningRequestCreate,
		ReadContext:   resourceADCertificateSigningRequestRead,
		UpdateContext: resourceADCertificateSigningRequestUpdate,
		DeleteContext: resourceADCertificateSigningRequestDelete,
		Description:   descriptions.ResourceADCertificateSigningRequest,
		Schema: map[string]*schema.Schema{
			attributes.ID: {
				Type:     schema.TypeString,
				Computed: true,
			},
			attributes.DisplayName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions.Name,
			},
			attributes.CommonName: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: descriptions.CertificateCommonName,
			},
			attributes.Details: {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				MinItems:    1,
				MaxItems:    1,
				Description: descriptions.CSRDetails,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						attributes.Organization: {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						attributes.OrganizationalUnit: {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						attributes.Locality: {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						attributes.Province: {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						attributes.Country: {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},
			attributes.Status: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: descriptions.CertificateStatus,
			},
			attributes.Content: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: descriptions.CertificateContent,
			},
			attributes.EnterpriseSigned: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: descriptions.EnterpriseSigned,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceADCertificateSigningRequestCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaPAMClient)

	var csrDetails *client.ADCertificateDetails
	if v, ok := d.GetOk(attributes.Details); ok {
		list := v.([]interface{})
		detailsMap := list[0].(map[string]interface{})
		csrDetails = &client.ADCertificateDetails{
			Organization:       utils.AsStringPtr(detailsMap[attributes.Organization].(string)),
			OrganizationalUnit: utils.AsStringPtr(detailsMap[attributes.OrganizationalUnit].(string)),
			Locality:           utils.AsStringPtr(detailsMap[attributes.Locality].(string)),
			Province:           utils.AsStringPtr(detailsMap[attributes.Province].(string)),
			Country:            utils.AsStringPtr(detailsMap[attributes.Country].(string)),
		}
	} else {
		return diag.Errorf("certificate details missing")
	}

	//Build Certificate API Request Object
	csrReq := client.ADSmartCardCertificate{
		DisplayName: getStringPtr(attributes.DisplayName, d, false),
		CommonName:  getStringPtr(attributes.CommonName, d, false),
		Type:        utils.AsStringPtr(client.ADCertificateTypeSigningRequest),
		Details:     csrDetails,
	}

	//Call api client
	if createdADCSR, err := c.CreateADSmartcardCertificate(ctx, csrReq); err != nil {
		return diag.FromErr(err)
	} else if createdADCSR == nil {
		d.SetId("")
	} else {
		//Set returned id
		d.SetId(*createdADCSR.ID)
		if err := d.Set(attributes.Content, *createdADCSR.Content); err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceADCertificateSigningRequestRead(ctx, d, m)
}

func resourceADCertificateSigningRequestRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaPAMClient)

	certificateID := d.Id()
	adCertificate, err := c.GetADSmartcardCertificate(ctx, certificateID)
	if err != nil {
		return diag.FromErr(err)
	}

	if adCertificate != nil && utils.IsNonEmpty(adCertificate.ID) {
		for key, value := range adCertificate.ToResourceMap() {
			_ = d.Set(key, value)
		}
	} else {
		logging.Infof("ADSmartCardCertificate %s does not exist", certificateID)
	}

	return nil
}

func resourceADCertificateSigningRequestUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaPAMClient)
	id := d.Id()

	changed := false
	updates := make(map[string]interface{})

	changeableAttributes := []string{
		attributes.DisplayName,
	}

	for _, attribute := range changeableAttributes {
		if d.HasChange(attribute) {
			updates[attribute] = d.Get(attribute)
			changed = true
		}
	}

	if changed {
		if err := c.UpdateADSmartcardCertificateName(ctx, id, updates); err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceADCertificateSigningRequestRead(ctx, d, m)
}

func resourceADCertificateSigningRequestDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(client.OktaPAMClient)
	certificateId := d.Id()

	err := c.DeleteADSmartcardCertificate(ctx, certificateId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return diags
}
