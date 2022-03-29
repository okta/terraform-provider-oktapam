package oktapam

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/client"
)

func dataSourceProjects() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceProjectsRead,
		Schema: map[string]*schema.Schema{
			"self": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"projects": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"team": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"next_unix_gid": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"next_unix_uid": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"create_server_users": {
							Type:     schema.TypeBool,
							Computed: true,
							Optional: true,
						},
						"deleted_at": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"forward_traffic": {
							Type:     schema.TypeBool,
							Computed: true,
							Optional: true,
						},
						"rdp_session_recording": {
							Type:     schema.TypeBool,
							Computed: true,
							Optional: true,
						},
						"require_preauth_for_creds": {
							Type:     schema.TypeBool,
							Computed: true,
							Optional: true,
						},
						"ssh_session_recording": {
							Type:     schema.TypeBool,
							Computed: true,
							Optional: true,
						},
						"gateway_selector": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"ssh_certificate_type": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceProjectsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(client.OktaPAMClient)
	parameters := client.ListProjectsParameters{}
	if contains, ok := d.GetOk("contains"); ok {
		parameters.Contains = contains.(string)
	}

	self, err := getOkBool("self", d)
	if err != nil {
		return diag.FromErr(err)
	}
	parameters.Self = self

	projectsList, err := c.ListProjects(ctx, parameters)
	if err != nil {
		return diag.FromErr(err)
	}

	projects := make([]map[string]interface{}, len(projectsList))
	for idx, proj := range projectsList {
		projects[idx] = proj.ToResourceMap()
	}

	if err := d.Set("projects", projects); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}
