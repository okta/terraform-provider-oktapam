package oktapam

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
)

func dataSourceSecrets() *schema.Resource {
	return &schema.Resource{
		Description: descriptions.SourceSecrets,
		ReadContext: dataSourceSecretList,
		Schema: map[string]*schema.Schema{
			attributes.Path: {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions.SecretOrParentFolderPath,
			},
			attributes.ResourceGroup: {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions.ResourceGroupID,
			},
			attributes.Project: {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions.ProjectID,
			},
			attributes.Secrets: {
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

func dataSourceSecretList(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := getSDKClientFromMetadata(m)

	resourceGroupID := d.Get(attributes.ResourceGroup).(string)
	projectID := d.Get(attributes.Project).(string)
	path := d.Get(attributes.Path).(string)

	secrets, err := client.ListSecrets(ctx, c, resourceGroupID, projectID, path, false)
	if err != nil {
		return diag.FromErr(err)
	}

	secretsList := make([]map[string]any, len(secrets))
	for idx, secret := range secrets {
		attrs := make(map[string]any, 2)

		attrs[attributes.ID] = secret.Secret.Id
		attrs[attributes.Name] = secret.Secret.Name

		secretsList[idx] = attrs
	}

	d.Set(attributes.Secrets, secretsList)
	d.SetId(fmt.Sprintf("%s|%s|%s", resourceGroupID, projectID, path))

	return nil
}
