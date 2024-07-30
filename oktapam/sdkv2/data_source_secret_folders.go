package sdkv2

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
)

func dataSourceSecretFolders() *schema.Resource {
	return &schema.Resource{
		Description: descriptions.SourceSecretFolders,
		ReadContext: dataSourceSecretFolderList,
		Schema: map[string]*schema.Schema{
			attributes.Path: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions.SecretFolderPath,
			},
			attributes.ListElementsUnderPath: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: descriptions.ListElementsUnderPath,
			},
			attributes.ResourceGroup: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions.ResourceGroupID,
			},
			attributes.Project: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions.ProjectID,
			},
			attributes.SecretFolders: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						attributes.ID: {
							Type:     schema.TypeString,
							Computed: true,
							// auto-generated description
						},
						attributes.Name: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: descriptions.Name,
						},
					},
				},
			},
		},
	}
}

func dataSourceSecretFolderList(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	var diags diag.Diagnostics
	c := client.GetSDKClientFromMetadata(m)

	resourceGroupID := d.Get(attributes.ResourceGroup).(string)
	projectID := d.Get(attributes.Project).(string)
	path := d.Get(attributes.Path).(string)

	var secretFolders []client.SecretFolder

	if resourceGroupID != "" && projectID != "" {
		if path == "" {
			if response, err := client.ListTopLevelSecretFoldersForProject(ctx, c, resourceGroupID, projectID); err == nil {
				secretFolders = response
			} else {
				diags = append(diags, diag.FromErr(err)...)
			}
		} else {
			listElementsUnderPath := d.Get(attributes.ListElementsUnderPath).(bool)
			if listElementsUnderPath {
				if response, err := client.ListSecretFoldersUnderPath(ctx, c, resourceGroupID, projectID, path); err == nil {
					secretFolders = response
				} else {
					diags = append(diags, diag.FromErr(err)...)
				}
			} else {
				if response, err := client.ResolveSecretFolder(ctx, c, resourceGroupID, projectID, path); err == nil {
					secretFolders = append(secretFolders, *response)
				} else {
					diags = append(diags, diag.FromErr(err)...)
				}
			}
		}
	} else if resourceGroupID == "" && projectID == "" && path == "" {
		if response, err := client.ListTopLevelSecretFolders(ctx, c); err == nil {
			secretFolders = response
		} else {
			diags = append(diags, diag.FromErr(err)...)
		}
	} else {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("must supply either both %s and %s to obtain Secret Folders for the specified project or neither values to get top-level accessible Secret Folders for the user", attributes.ResourceGroup, attributes.Project),
		})
	}

	folderList := make([]map[string]any, len(secretFolders))
	for idx, secretFolder := range secretFolders {
		attrs := make(map[string]any, 2)

		attrs[attributes.ID] = *secretFolder.ID
		attrs[attributes.Name] = *secretFolder.Name

		folderList[idx] = attrs
	}

	d.Set(attributes.SecretFolders, folderList)
	d.SetId(fmt.Sprintf("%s|%s|%s", resourceGroupID, projectID, path))

	return diags
}
