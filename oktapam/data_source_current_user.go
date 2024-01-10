package oktapam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
)

func dataSourceCurrentUser() *schema.Resource {
	return &schema.Resource{
		Description: descriptions.SourceCurrentUser,
		ReadContext: dataSourceCurrentUserFetch,
		Schema: map[string]*schema.Schema{
			attributes.Name: {
				Description: descriptions.UserName,
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceCurrentUserFetch(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := getSDKClientFromMetadata(m)
	userInfo, err := client.GetCurrentUser(ctx, c)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(userInfo.Id)
	d.Set(attributes.Name, userInfo.Name)
	return nil
}
