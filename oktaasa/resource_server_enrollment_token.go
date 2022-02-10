package oktaasa

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-oktaasa/oktaasa/client"
	"github.com/terraform-providers/terraform-provider-oktaasa/oktaasa/logging"
	"github.com/terraform-providers/terraform-provider-oktaasa/oktaasa/utils"
)

func resourceServerEnrollmentToken() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceServerEnrollmentTokenCreate,
		ReadContext:   resourceServerEnrollmentTokenRead,
		DeleteContext: resourceServerEnrollmentTokenDelete,
		Schema: map[string]*schema.Schema{
			"project_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"token": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_by_user": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"issued_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceServerEnrollmentTokenRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaASAClient)

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
	c := m.(client.OktaASAClient)

	project := d.Get("project_name").(string)
	description := d.Get("description").(string)

	token, err := c.CreateServerEnrollmentToken(ctx, project, description)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%s|%s", *token.Project, *token.ID))

	return resourceServerEnrollmentTokenRead(ctx, d, m)
}

func resourceServerEnrollmentTokenDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaASAClient)

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

func parseServerEnrollmentTokenResourceID(resourceId string) (string, string, error) {
	split := strings.Split(resourceId, "|")
	if len(split) != 2 {
		return "", "", fmt.Errorf("oktaasa_server_enrollment_token id must be in the format of <project name>|<enrollment token id>, received: %s", resourceId)
	}
	return split[0], split[1], nil
}
