package oktapam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"
)

func resourceADCertificateObject() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceADCertificateUploadCreate,
		ReadContext:   resourceADCertificateRead,
		DeleteContext: resourceADCertificateDelete,
		Description:   descriptions.ResourceADCertificateObject,
		Schema: map[string]*schema.Schema{
			attributes.ID: {
				Type:     schema.TypeString,
				Computed: true,
			},
			attributes.CertificateID: {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			attributes.Source: {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: descriptions.CertificateContent,
			},
			attributes.EnterpriseSigned: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: descriptions.EnterpriseSigned,
			},
			attributes.Status: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: descriptions.CertificateStatus,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceADCertificateUploadCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := getLocalClientFromMetadata(m)

	certificateID := d.Get(attributes.CertificateID).(string)
	source := d.Get(attributes.Source).(string)

	//Call api client
	if err := c.UploadADSmartcardCertificate(ctx, certificateID, source); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(certificateID)
	return resourceADCertificateRead(ctx, d, m)
}

func resourceADCertificateRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := getLocalClientFromMetadata(m)

	certificateID := d.Id()
	adCertificate, err := c.GetADSmartcardCertificate(ctx, certificateID)
	if err != nil {
		return diag.FromErr(err)
	}

	if adCertificate != nil && utils.IsNonEmpty(adCertificate.ID) {
		if err := d.Set(attributes.ID, *adCertificate.ID); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set(attributes.EnterpriseSigned, *adCertificate.EnterpriseSigned); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set(attributes.Status, *adCertificate.Status); err != nil {
			return diag.FromErr(err)
		}
	} else {
		logging.Infof("ADSmartCardCertificate %s does not exist", certificateID)
	}

	return nil
}

func resourceADCertificateDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := getLocalClientFromMetadata(m)
	certificateId := d.Id()

	err := c.DeleteADSmartcardCertificate(ctx, certificateId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}
