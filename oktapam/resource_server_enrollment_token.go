package oktapam

import (
	"context"
	"fmt"
	"strings"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"
)

func resourceServerEnrollmentToken() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceServerEnrollmentTokenCreate,
		ReadContext:   resourceServerEnrollmentTokenRead,
		DeleteContext: resourceServerEnrollmentTokenDelete,
		Description:   descriptions.ResourceServerEnrollmentToken,
		Schema: map[string]*schema.Schema{
			attributes.ProjectName: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: descriptions.ProjectName,
			},
			attributes.Description: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: descriptions.DescriptionContains,
			},
			attributes.Token: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: descriptions.Token,
			},
			attributes.ID: {
				Type:     schema.TypeString,
				Computed: true,
				// Description is autogenerated
			},
			attributes.CreatedByUser: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: descriptions.CreatedByUser,
			},
			attributes.IssuedAt: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: descriptions.IssuedAt,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceServerEnrollmentTokenRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaPAMClient)

	resourceId := d.Id()
	project, id, err := parseServerEnrollmentTokenResourceID(resourceId)
	if err != nil {
		return diag.FromErr(err)
	}

	token, err := c.GetServerEnrollmentToken(ctx, project, id)
	if err != nil {
		return diag.FromErr(err)
	}

	if token == nil || utils.IsBlank(token.ID) {
		logging.Debugf("token id was blank")
		d.SetId("")
		return nil
	}

	for key, value := range token.ToResourceMap() {
		logging.Debugf("setting %s to %v", key, value)
		d.Set(key, value)
	}

	return nil
}

func resourceServerEnrollmentTokenCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaPAMClient)

	project := d.Get(attributes.ProjectName).(string)
	description := d.Get(attributes.Description).(string)

	token, err := c.CreateServerEnrollmentToken(ctx, project, description)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(createServerEnrollmentTokenResourceID(*token.Project, *token.ID))

	return resourceServerEnrollmentTokenRead(ctx, d, m)
}

func resourceServerEnrollmentTokenDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaPAMClient)

	resourceId := d.Id()
	project, id, err := parseServerEnrollmentTokenResourceID(resourceId)
	if err != nil {
		return diag.FromErr(err)
	}

	err = c.DeleteServerEnrollmentToken(ctx, project, id)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func createServerEnrollmentTokenResourceID(project, tokenID string) string {
	return fmt.Sprintf("%s|%s", project, tokenID)
}

func parseServerEnrollmentTokenResourceID(resourceId string) (string, string, error) {
	split := strings.Split(resourceId, "|")
	if len(split) != 2 {
		return "", "", fmt.Errorf("oktapam_server_enrollment_token id must be in the format of <project name>|<enrollment token id>, received: %s", resourceId)
	}
	return split[0], split[1], nil
}
