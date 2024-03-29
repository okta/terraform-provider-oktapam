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

func dataSourceSecret() *schema.Resource {
	return &schema.Resource{
		Description: descriptions.SourceSecret,
		ReadContext: dataSourceSecretFetch,
		Schema: map[string]*schema.Schema{
			attributes.ID: {
				Type:     schema.TypeString,
				Required: true,
				// Description is autogenerated
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
			attributes.Name: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: descriptions.Name,
			},
			attributes.ParentFolder: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: descriptions.ParentFolderID,
			},
			attributes.Description: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: descriptions.Description,
			},
			attributes.Secret: {
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed:    true,
				Description: descriptions.SecretKeyValues,
				Sensitive:   true,
			},
		},
	}
}

func dataSourceSecretFetch(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := getSDKClientFromMetadata(m)

	idI, ok := d.GetOk(attributes.ID)
	if !ok {
		return diag.Diagnostics{
			{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("missing required %s parameter", attributes.ID),
			},
		}
	}

	secretID := idI.(string)
	resourceGroupID := d.Get(attributes.ResourceGroup).(string)
	projectID := d.Get(attributes.Project).(string)

	secretWrapper, err := client.RevealSecret(ctx, c, resourceGroupID, projectID, secretID)
	if err != nil {
		return diag.FromErr(err)
	}

	attrs := secretWrapper.ToResourceMap()

	for k, v := range attrs {
		d.Set(k, v)
	}
	d.SetId(secretID)

	return nil
}
