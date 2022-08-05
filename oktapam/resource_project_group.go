package oktapam

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"
)

func resourceProjectGroup() *schema.Resource {
	return &schema.Resource{
		Description:   descriptions.ResourceProjectGroup,
		CreateContext: resourceProjectGroupCreate,
		ReadContext:   resourceProjectGroupRead,
		UpdateContext: resourceProjectGroupUpdate,
		DeleteContext: resourceProjectGroupDelete,
		Schema: map[string]*schema.Schema{
			attributes.ProjectName: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: descriptions.ProjectName,
			},
			attributes.GroupName: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: descriptions.GroupName,
			},
			attributes.DeletedAt: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: descriptions.DeletedAt,
			},
			attributes.RemovedAt: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: descriptions.RemovedAt,
			},
			attributes.CreateServerGroup: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: descriptions.CreateServerGroup,
			},
			attributes.ServerAccess: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: descriptions.ServerAccess,
			},
			attributes.ServerAdmin: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: descriptions.ServerAdmin,
			},
			attributes.ServersSelector: {
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Description: descriptions.ServersSelector,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: importState,
		},
	}
}

func resourceProjectGroupCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaPAMClient)

	serverAdmin, err := getOkBool(attributes.ServerAdmin, d)
	if err != nil {
		return diag.FromErr(err)
	}
	serverAccess, err := getOkBool(attributes.ServerAccess, d)
	if err != nil {
		return diag.FromErr(err)
	}

	if !serverAdmin && !serverAccess {
		return diag.Errorf("server_access or server_admin must be true")
	}

	createServerGroup, err := getOkBool(attributes.CreateServerGroup, d)
	if err != nil {
		return diag.FromErr(err)
	}

	var serversSelector map[string]interface{}
	if ss, ok := d.GetOk(attributes.ServersSelector); ok {
		serversSelector = ss.(map[string]interface{})
	}

	serversSelectorString, err := client.FormatServersSelectorString(serversSelector)
	if err != nil {
		return diag.FromErr(err)
	}

	projectGroup := client.ProjectGroup{
		Project:           getStringPtr(attributes.ProjectName, d, true),
		Group:             getStringPtr(attributes.GroupName, d, true),
		ServerAccess:      serverAccess,
		ServerAdmin:       serverAdmin,
		CreateServerGroup: createServerGroup,
		ServersSelector:   serversSelectorString,
	}

	err = c.CreateProjectGroup(ctx, projectGroup)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resource.UniqueId())
	return resourceProjectGroupRead(ctx, d, m)
}

func resourceProjectGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaPAMClient)

	projectName := d.Get(attributes.ProjectName).(string)
	groupName := d.Get(attributes.GroupName).(string)

	projectGroup, err := c.GetProjectGroup(ctx, projectName, groupName)
	if err != nil {
		return diag.FromErr(err)
	}

	ignorableValues := map[string]bool{attributes.DeletedAt: true, attributes.RemovedAt: true}
	if projectGroup != nil && utils.IsNonEmpty(projectGroup.Group) {
		if utils.IsNonEmpty(projectGroup.DeletedAt) {
			logging.Infof("Group %s was deleted", groupName)
			d.SetId("")
		} else if utils.IsNonEmpty(projectGroup.RemovedAt) {
			logging.Infof("Group %s was removed from project %s", groupName, projectName)
			d.SetId("")
		} else {
			//d.SetId(createProjectGroupResourceID(*projectGroup.Project, *projectGroup.Group))
			attrs, err := projectGroup.ToResourceMap()
			if err != nil {
				return diag.FromErr(err)
			}

			for key, value := range attrs {
				if _, ok := ignorableValues[key]; !ok {
					if err := d.Set(key, value); err != nil {
						return diag.FromErr(err)
					}
				}
			}
		}
	} else {
		logging.Infof("Group %s is not assigned to project %s", groupName, projectName)
		d.SetId("")
	}

	return nil
}

func resourceProjectGroupUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaPAMClient)

	changed := false
	updates := make(map[string]interface{})

	changeableAttributes := []string{
		attributes.ServerAccess,
		attributes.ServerAdmin,
		attributes.CreateServerGroup,
		attributes.ServersSelector,
	}

	requiredAttributes := []string{
		attributes.ProjectName,
		attributes.GroupName,
		attributes.ServerAccess,
		attributes.ServerAdmin,
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
	c := m.(client.OktaPAMClient)

	project := d.Get(attributes.ProjectName).(string)
	group := d.Get(attributes.GroupName).(string)

	err := c.DeleteProjectGroup(ctx, project, group)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}

func createProjectGroupResourceID(project string, group string) string {
	return fmt.Sprintf("%s|%s", project, group)
}

func parseProjectGroupResourceID(resourceId string) (string, string, error) {
	split := strings.Split(resourceId, "|")
	if len(split) != 2 {
		return "", "", fmt.Errorf("oktapam_project_group id must be in the format of <project name>|<group name>, received: %s", resourceId)
	}
	return split[0], split[1], nil
}

func importState(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// d.Id() here is the last argument passed to the `terraform import RESOURCE_TYPE.RESOURCE_NAME RESOURCE_ID` command
	// Here we use a function to parse the import ID (like the example above) to simplify our logic
	projectName, groupName, err := parseProjectGroupResourceID(d.Id())

	if err != nil {
		return nil, err
	}

	if err := d.Set(attributes.ProjectName, projectName); err != nil {
		return nil, err
	}
	if err := d.Set(attributes.GroupName, groupName); err != nil {
		return nil, err
	}

	d.SetId(resource.UniqueId())
	return []*schema.ResourceData{d}, nil
}
