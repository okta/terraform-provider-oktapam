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

func resourceProjectGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceProjectGroupCreate,
		ReadContext:   resourceProjectGroupRead,
		UpdateContext: resourceProjectGroupUpdate,
		DeleteContext: resourceProjectGroupDelete,
		Schema: map[string]*schema.Schema{
			"project_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"group_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"group_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"create_server_group": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"server_access": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"server_admin": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"servers_selector": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceProjectGroupCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaASAClient)

	serverAdmin, err := getOkBool("server_admin", d)
	if err != nil {
		return diag.FromErr(err)
	}
	serverAccess, err := getOkBool("server_access", d)
	if err != nil {
		return diag.FromErr(err)
	}

	if !serverAdmin && !serverAccess {
		return diag.Errorf("server_access or server_admin must be true")
	}

	createServerGroup, err := getOkBool("create_server_group", d)
	if err != nil {
		return diag.FromErr(err)
	}

	var serversSelector map[string]interface{}
	if ss, ok := d.GetOk("servers_selector"); ok {
		serversSelector = ss.(map[string]interface{})
	}

	serversSelectorString, err := client.FormatServersSelectorString(serversSelector)
	if err != nil {
		return diag.FromErr(err)
	}

	projectGroup := client.ProjectGroup{
		Project:          getStringPtr("project_name", d, true),
		Group:            getStringPtr("group_name", d, true),
		ServerAccess:     serverAccess,
		ServerAdmin:      serverAdmin,
		CreateServerGoup: createServerGroup,
		ServersSelector:  serversSelectorString,
	}

	err = c.CreateProjectGroup(ctx, projectGroup)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(createProjectGroupResourceID(*projectGroup.Project, *projectGroup.Group))

	return resourceProjectGroupRead(ctx, d, m)
}

func resourceProjectGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaASAClient)
	project, group, err := parseProjectGroupResourceID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	if project == "" {
		return diag.Errorf("blank project name from id %s", d.Id())
	}

	projectGroup, err := c.GetProjectGroup(ctx, project, group)
	if err != nil {
		return diag.FromErr(err)
	}

	ignorableValues := map[string]bool{"deleted_at": true, "removed_at": true}
	if projectGroup != nil && utils.IsNonEmpty(projectGroup.Group) {
		if utils.IsNonEmpty(projectGroup.DeletedAt) {
			logging.Infof("Group %s was deleted", group)
			d.SetId("")
		} else if utils.IsNonEmpty(projectGroup.RemovedAt) {
			logging.Infof("Group %s was removed from project %s", group, project)
			d.SetId("")
		} else {
			d.SetId(createProjectGroupResourceID(*projectGroup.Project, *projectGroup.Group))
			attributes, err := projectGroup.ToResourceMap()
			if err != nil {
				return diag.FromErr(err)
			}

			for key, value := range attributes {
				if _, ok := ignorableValues[key]; !ok {
					d.Set(key, value)
				}
			}
		}
	} else {
		logging.Infof("Group %s is not assigned to project %s", group, project)
		d.SetId("")
	}

	return nil
}

func resourceProjectGroupUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaASAClient)

	changed := false
	updates := make(map[string]interface{})

	changeableAttributes := []string{
		"server_access",
		"server_admin",
		"create_server_group",
		"servers_selector",
	}

	requiredAttributes := []string{
		"project_name",
		"group_name",
		"server_access",
		"server_admin",
	}

	for _, attribute := range changeableAttributes {
		updates[attribute] = d.Get(attribute)
		if d.HasChange(attribute) {
			changed = true
		}
	}

	for _, attribute := range requiredAttributes {
		updates[attribute] = d.Get(attribute)
	}

	if changed {
		pg, err := client.ProjectGroupFromMap(updates)
		if err != nil {
			return diag.FromErr(err)
		} else if pg == nil {
			return diag.Errorf("could not create ProjectGroup from supplied values")
		}
		err = c.UpdateProjectGroup(ctx, *pg)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceProjectGroupRead(ctx, d, m)
}

func resourceProjectGroupDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaASAClient)
	project, group, err := parseProjectGroupResourceID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	err = c.DeleteProjectGroup(ctx, project, group)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}

func createProjectGroupResourceID(project, group string) string {
	return fmt.Sprintf("%s|%s", project, group)
}

func parseProjectGroupResourceID(resourceId string) (string, string, error) {
	split := strings.Split(resourceId, "|")
	if len(split) != 2 {
		return "", "", fmt.Errorf("oktaasa_project_group id must be in the format of <project name>|<group name>, received: %s", resourceId)
	}
	return split[0], split[1], nil
}
