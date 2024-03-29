package oktapam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
)

func resourceSecret() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSecretCreate,
		ReadContext:   resourceSecretRead,
		UpdateContext: resourceSecretUpdate,
		DeleteContext: resourceSecretDelete,
		Description:   descriptions.ResourceSecret,
		Schema: map[string]*schema.Schema{
			attributes.ID: {
				Type:     schema.TypeString,
				Computed: true,
				// Description is autogenerated
			},
			attributes.Name: {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions.Name,
			},
			attributes.ResourceGroup: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: descriptions.ResourceGroupID,
			},
			attributes.Project: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: descriptions.ProjectID,
			},
			attributes.ParentFolder: {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions.ParentFolderID,
				ForceNew:    true,
			},
			attributes.Description: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions.Description,
			},
			attributes.Secret: {
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required:    true,
				Description: descriptions.SecretKeyValues,
				Sensitive:   true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceSecretCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := getSDKClientFromMetadata(m)

	resourceGroupID := GetStringPtrFromResource(attributes.ResourceGroup, d, true)
	projectID := GetStringPtrFromResource(attributes.Project, d, true)
	name := GetStringPtrFromResource(attributes.Name, d, true)
	parentFolderID := GetStringPtrFromResource(attributes.ParentFolder, d, true)
	description := GetStringPtrFromResource(attributes.Description, d, true)
	secretMap, diags := GetMapPtrFromResource[string](attributes.Secret, d)
	if diags != nil {
		return diags
	} else if secretMap == nil {
		return diag.Errorf("%s attribute must contain at least one element", attributes.Secret)
	}

	if secret, err := client.CreateSecret(ctx, c, *resourceGroupID, *projectID, *parentFolderID, *name, *description, *secretMap); err != nil {
		return diag.FromErr(err)
	} else {
		d.SetId(secret.Secret.Id)
	}

	return resourceSecretRead(ctx, d, m)
}

func resourceSecretRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := getSDKClientFromMetadata(m)

	secretID := d.Id()
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

	return nil
}

func resourceSecretUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := getSDKClientFromMetadata(m)

	id := d.Id()
	resourceGroupID := GetStringPtrFromResource(attributes.ResourceGroup, d, true)
	projectID := GetStringPtrFromResource(attributes.Project, d, true)
	name := GetStringPtrFromResource(attributes.Name, d, true)
	parentFolderID := GetStringPtrFromResource(attributes.ParentFolder, d, true)
	description := GetStringPtrFromResource(attributes.Description, d, true)
	secretMap, diags := GetMapPtrFromResource[string](attributes.Secret, d)
	if diags != nil {
		return diags
	} else if secretMap == nil {
		return diag.Errorf("%s attribute must contain at least one element", attributes.Secret)
	}

	if secret, err := client.UpdateSecret(ctx, c, *resourceGroupID, *projectID, *parentFolderID, id, *name, *description, *secretMap); err != nil {
		return diag.FromErr(err)
	} else {
		d.SetId(secret.Secret.Id)
	}

	return resourceSecretRead(ctx, d, m)
}

func resourceSecretDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := getSDKClientFromMetadata(m)

	secretID := d.Id()
	resourceGroupID := d.Get(attributes.ResourceGroup).(string)
	projectID := d.Get(attributes.Project).(string)

	if err := client.DeleteSecret(ctx, c, resourceGroupID, projectID, secretID); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
