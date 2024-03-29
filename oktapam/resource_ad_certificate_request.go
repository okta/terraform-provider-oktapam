package oktapam

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/typed_strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"
)

/*
Below variables are used to define ConflictsWith contraint in the terraform resource schema.
ConflictsWith is a set of attribute paths, including this attribute, whose configurations cannot be set simultaneously.
This implements the validation logic declaratively within the schema and can trigger earlier in Terraform operations,
rather than using create or update logic which in Terraform operations, rather than using create or update logic which only triggers during apply.

Only absolute attribute paths, ones starting with top level attribute names, are supported. Attribute paths cannot be accurately declared
for TypeList (if MaxItems is greater than 1), TypeMap, or TypeSet attributes. To reference an attribute under a single configuration block
(TypeList with Elem of *Resource and MaxItems of 1), the syntax is "parent_block_name.0.child_attribute_name".

Reference: https://github.com/hashicorp/terraform-plugin-sdk/blob/main/helper/schema/schema.go#L257
*/
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
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				MinItems:    1,
				MaxItems:    1,
				Description: descriptions.CSRDetails,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						attributes.Organization: {
							Type:          schema.TypeString,
							ConflictsWith: []string{certDetailsTTLDays},
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
							Type:          schema.TypeInt,
							ConflictsWith: []string{certDetailsOrganization, certDetailsOrganizationalUnit, certDetailsLocality, certDetailsProvince, certDetailsCountry},
							Optional:      true,
							ForceNew:      true,
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

func resourceADCertificateRequestCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := getLocalClientFromMetadata(m)

	var certRequestDetails *client.ADCertificateDetails
	if v, ok := d.GetOk(attributes.Details); ok {
		list := v.([]any)
		detailsMap := list[0].(map[string]any)
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
		DisplayName: GetStringPtrFromResource(attributes.DisplayName, d, false),
		CommonName:  GetStringPtrFromResource(attributes.CommonName, d, false),
		Type:        GetStringPtrFromResource(attributes.Type, d, false),
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

func resourceADCertificateRequestRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	var diags diag.Diagnostics
	c := getLocalClientFromMetadata(m)

	certificateID := d.Id()
	adCertificate, err := c.GetADSmartcardCertificate(ctx, certificateID)
	if err != nil {
		return diag.FromErr(err)
	}

	if adCertificate != nil && utils.IsNonEmpty(adCertificate.ID) {
		for key, value := range adCertificate.ToResourceMap() {
			if err := d.Set(key, value); err != nil {
				diags = append(diags, diag.FromErr(err)...)
			}
		}

		if details, ok := d.Get(attributes.Details).([]map[string]any); ok && len(details) == 1 {
			// API doesn't return certificate details so that need to be set again from the current config
			// If we don't set it then terraform report differences without making any changes between proposed state and real-world infra
			flattenedCertDetails := make([]any, 1)
			flattenedCertDetail := make(map[string]any)
			flattenedCertDetail[attributes.Organization] = details[0][attributes.Organization]
			flattenedCertDetail[attributes.OrganizationalUnit] = details[0][attributes.OrganizationalUnit]
			flattenedCertDetail[attributes.Locality] = details[0][attributes.Locality]
			flattenedCertDetail[attributes.Province] = details[0][attributes.Province]
			flattenedCertDetail[attributes.Country] = details[0][attributes.Country]
			flattenedCertDetail[attributes.TTLDays] = details[0][attributes.TTLDays]

			flattenedCertDetails[0] = flattenedCertDetail
			if err := d.Set(attributes.Details, flattenedCertDetails); err != nil {
				return diag.FromErr(err)
			}
		}
	} else {
		logging.Infof("ADSmartCardCertificate %s does not exist", certificateID)
	}

	return diags
}

func resourceADCertificateRequestUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := getLocalClientFromMetadata(m)
	id := d.Id()

	changed := false
	updates := make(map[string]any)

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

func resourceADCertificateRequestDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := getLocalClientFromMetadata(m)
	certificateId := d.Id()

	err := c.DeleteADSmartcardCertificate(ctx, certificateId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}
