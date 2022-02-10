package oktaasa

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-oktaasa/oktaasa/client"
	"github.com/terraform-providers/terraform-provider-oktaasa/oktaasa/logging"
	"github.com/terraform-providers/terraform-provider-oktaasa/oktaasa/utils"
)

func resourceProject() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceProjectCreate,
		ReadContext:   resourceProjectRead,
		UpdateContext: resourceProjectUpdate,
		DeleteContext: resourceProjectDelete,
		Schema: map[string]*schema.Schema{
			"project_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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
				Optional: true,
				Default:  63001,
			},
			"next_unix_uid": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60001,
			},
			"force_shared_ssh_users": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"shared_admin_user_name": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"shared_standard_user_name": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"create_server_users": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"deleted_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"forward_traffic": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"rdp_session_recording": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"require_preauth_for_creds": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"ssh_session_recording": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"gateway_selector": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ad_joined_users": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: resourceProjectReadImport,
		},
	}
}

func resourceProjectCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaASAClient)

	project := client.Project{
		Name:                   getStringPtr("project_name", d, true),
		NextUnixGID:            getIntPtr("next_unix_gid", d, false),
		NextUnixUID:            getIntPtr("next_unix_uid", d, false),
		ForceSharedSSHUsers:    getBoolPtr("force_shared_ssh_users", d, false),
		SharedAdminUserName:    getStringPtr("shared_admin_user_name", d, false),
		SharedStandardUserName: getStringPtr("shared_standard_user_name", d, false),
		CreateServerUsers:      getBoolPtr("create_server_users", d, false),
		ForwardTraffic:         getBoolPtr("forward_traffic", d, false),
		RDPSessionRecording:    getBoolPtr("rdp_session_recording", d, false),
		RequirePreAuthForCreds: getBoolPtr("require_preauth_for_creds", d, false),
		SSHSessionRecording:    getBoolPtr("ssh_session_recording", d, false),
		GatewaySelector:        getStringPtr("gateway_selector", d, false),
		ADJoinedUsers:          getBoolPtr("ad_joined_users", d, false),
	}

	err := c.CreateProject(ctx, project)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*project.Name)
	return resourceProjectRead(ctx, d, m)
}

func resourceProjectReadWithIgnorable(ctx context.Context, d *schema.ResourceData, m interface{}, ignoreValues bool) (*schema.ResourceData, error) {
	c := m.(client.OktaASAClient)

	projectName := d.Id()
	proj, err := c.GetProject(ctx, projectName, false)
	if err != nil {
		return nil, err
	}

	// these values are updated in normal operation of ASA and should not be reset to
	// the values that were used to create the project
	ignorableValues := map[string]bool{"next_unix_gid": true, "next_unix_uid": true}

	if proj != nil && utils.IsNonEmpty(proj.Name) {
		if utils.IsBlank(proj.DeletedAt) {
			logging.Infof("Project %s exists", proj.Name)
			d.SetId(*proj.Name)
		} else {
			logging.Infof("Project %s was removed", projectName)
			d.SetId("")
		}
		for key, value := range proj.ToResourceMap() {
			if _, ok := ignorableValues[key]; !ignoreValues || !ok {
				d.Set(key, value)
			}
		}
	} else {
		logging.Infof("Project %s does not exist", projectName)
		d.SetId("")
	}

	return d, nil
}

func resourceProjectReadImport(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	projectResource, err := resourceProjectReadWithIgnorable(ctx, d, m, false)
	if err != nil {
		return nil, err
	}
	return []*schema.ResourceData{projectResource}, nil
}

func resourceProjectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	_, err := resourceProjectReadWithIgnorable(ctx, d, m, true)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceProjectUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaASAClient)
	projectName := d.Id()

	changed := false
	updates := make(map[string]interface{})

	changeableAttributes := []string{
		"next_unix_gid",
		"next_unix_uid",
		"force_shared_ssh_users",
		"shared_admin_user_name",
		"create_server_users",
		"forward_traffic",
		"rdp_session_recording",
		"require_preauth_for_creds",
		"ssh_session_recording",
		"gateway_selector",
		"ad_joined_users",
	}

	for _, attribute := range changeableAttributes {
		if d.HasChange(attribute) {
			updates[attribute] = d.Get(attribute)
			changed = true
		}
	}

	if changed {
		err := c.UpdateProject(ctx, projectName, updates)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceProjectRead(ctx, d, m)
}

func resourceProjectDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(client.OktaASAClient)
	projectName := d.Id()

	err := c.DeleteProject(ctx, projectName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return diags
}
