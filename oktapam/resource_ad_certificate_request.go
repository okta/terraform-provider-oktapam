package oktapam

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/typed_strings"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"
)


// Below variables are used to define ConflictsWith contraint in the terraform resource schema
// ConflictsWith is a set of attribute paths, including this attribute,
// whose configurations cannot be set simultaneously. This implements the
// validation logic declaratively within the schema and can trigger earlier
// in Terraform operations, rather than using create or update logic which
// only triggers during apply.
//
// Only absolute attribute paths, ones starting with top level attribute
// names, are supported. Attribute paths cannot be accurately declared
// for TypeList (if MaxItems is greater than 1), TypeMap, or TypeSet
// attributes. To reference an attribute under a single configuration block
// (TypeList with Elem of *Resource and MaxItems of 1), the syntax is
// "parent_block_name.0.child_attribute_name".
// Reference: https://github.com/hashicorp/terraform-plugin-sdk/blob/main/helper/schema/schema.go#L257
var certDetailsOrganization = strings.Join([]string{attributes.Details, "0", attributes.Organization}, ".")
var certDetailsOrganizationalUnit = strings.Join([]string{attributes.Details, "0", attributes.OrganizationalUnit}, ".")
var certDetailsLocality = strings.Join([]string{attributes.Details, "0", attributes.Locality}, ".")
var certDetailsProvince = strings.Join([]string{attributes.Details, "0", attributes.Province}, ".")
var certDetailsCountry = strings.Join([]string{attributes.Details, "0", attributes.Country}, ".")
var certDetailsTTLDays = strings.Join([]string{attributes.Details, "0", attributes.TTLDays}, ".")

func resourceADCertificateRequest() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceADCertificateRequestCreate,
		ReadContext:   resourceADCertificateRequestRead,
		UpdateContext: resourceADCertificateRequestUpdate,
		DeleteContext: resourceADCertificateRequestDelete,
		Description:   descriptions.ResourceADCertificateRequest,
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
				Type:         schema.TypeList,
				Required:     true,
				ForceNew:     true,
				MinItems:     1,
				MaxItems:     1,
				Description:  descriptions.CSRDetails,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						attributes.Organization: {
							Type: schema.TypeString,
							ConflictsWith: []string{certDetailsTTLDays},
							ValidateFunc: validateCertificateDetails,
							Optional:      true,
							ForceNew:      true,
						},
						attributes.OrganizationalUnit: {
							Type:          schema.TypeString,
							ConflictsWith: []string{certDetailsTTLDays},
							Optional:      true,
							ForceNew:      true,
						},
						attributes.Locality: {
							Type:          schema.TypeString,
							ConflictsWith: []string{certDetailsTTLDays},
							Optional:      true,
							ForceNew:      true,
						},
						attributes.Province: {
							Type:          schema.TypeString,
							ConflictsWith: []string{certDetailsTTLDays},
							Optional:      true,
							ForceNew:      true,
						},
						attributes.Country: {
							Type:          schema.TypeString,
							ConflictsWith: []string{certDetailsTTLDays},
							Optional:      true,
							ForceNew:      true,
						},
						attributes.TTLDays: {
							Type: schema.TypeInt,
							ConflictsWith: []string{certDetailsOrganization, certDetailsOrganizationalUnit, certDetailsLocality, certDetailsProvince, certDetailsCountry},
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},
			attributes.Type: {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice(
					[]string{
						typed_strings.ADCertificateTypeSigningRequest.String(),
						typed_strings.ADCertificateTypeSelfSigned.String(),
					}, false),
				Description: descriptions.CertificateRequestType,
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

func validateCertificateDetails(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	fmt.Println(value)
	return
}

func resourceADCertificateRequestCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaPAMClient)

	var certRequestDetails *client.ADCertificateDetails
	if v, ok := d.GetOk(attributes.Details); ok {
		list := v.([]interface{})
		detailsMap := list[0].(map[string]interface{})
		certRequestDetails = &client.ADCertificateDetails{
			Organization:       utils.AsStringPtr(detailsMap[attributes.Organization].(string)),
			OrganizationalUnit: utils.AsStringPtr(detailsMap[attributes.OrganizationalUnit].(string)),
			Locality:           utils.AsStringPtr(detailsMap[attributes.Locality].(string)),
			Province:           utils.AsStringPtr(detailsMap[attributes.Province].(string)),
			Country:            utils.AsStringPtr(detailsMap[attributes.Country].(string)),
			TTLDays:            utils.AsIntPtr(detailsMap[attributes.TTLDays].(int)),
		}
	} else {
		return diag.Errorf("certificate details missing")
	}

	//Build Certificate API Request Object
	csrReq := client.ADSmartCardCertificate{
		DisplayName: getStringPtr(attributes.DisplayName, d, false),
		CommonName:  getStringPtr(attributes.CommonName, d, false),
		Type:        getStringPtr(attributes.Type, d, false),
		Details:     certRequestDetails,
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

	return resourceADCertificateRequestRead(ctx, d, m)
}

func resourceADCertificateRequestRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

func resourceADCertificateRequestUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

	return resourceADCertificateRequestRead(ctx, d, m)
}

func resourceADCertificateRequestDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
