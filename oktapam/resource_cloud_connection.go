package oktapam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"
)

func resourceCloudConnection() *schema.Resource {
	return &schema.Resource{
		Description:   descriptions.ResourceCloudConnection,
		ReadContext:   resourceCloudConnectionRead,
		CreateContext: resourceCloudConnectionCreate,
		DeleteContext: resourceCloudConnectionDelete,
		UpdateContext: resourceCloudConnectionUpdate,
		Schema: map[string]*schema.Schema{
			attributes.ID: {
				Type:     schema.TypeString,
				Computed: true,
			},
			attributes.Name: {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions.Name,
			},
			attributes.CloudConnectionProvider: {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions.CloudConnectionProvider,
			},
			attributes.CloudConnectionDetails: {
				Type:        schema.TypeList,
				MinItems: 1,
				MaxItems: 1,
				Required:    true,
				ForceNew:    true,
				Description: descriptions.CloudConnectionDetails,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						attributes.CloudConnectionAccountId: {
							Type:     schema.TypeString,
							Required: true,
						},
						attributes.CloudConnectionExternalId: {
							Type:     schema.TypeString,
							Required: true,
						},
						attributes.CloudConnectionRoleARN: {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceCloudConnectionRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := getLocalClientFromMetadata(m)
	cloudConnectionID := d.Id()
	cloudConnection, err := c.GetCloudConnection(ctx, cloudConnectionID)

	if err != nil {
		return diag.FromErr(err)
	}

	if cloudConnection == nil || utils.IsBlank(cloudConnection.ID) {
		logging.Debugf("cloud connection %s does not exist", cloudConnectionID)
	}

	for key, value := range cloudConnection.ToResourceMap() {
		logging.Debugf("setting %s to %v", key, value)
		if err := d.Set(key, value); err != nil {
			return diag.FromErr(err)
		}
	}

	if details, ok := d.Get(attributes.CloudConnectionDetails).([]map[string]any); ok && len(details) == 1 {
		flattenedDetails := make([]any, 1)
		flattenedDetail := make(map[string]any)
		flattenedDetail[attributes.CloudConnectionAccountId] = details[0][attributes.CloudConnectionAccountId]
		flattenedDetail[attributes.CloudConnectionExternalId] = details[0][attributes.CloudConnectionExternalId]
		flattenedDetail[attributes.CloudConnectionRoleARN] = details[0][attributes.CloudConnectionRoleARN]

		flattenedDetails[0] = flattenedDetail
		if err := d.Set(attributes.Details, flattenedDetails); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceCloudConnectionCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := getLocalClientFromMetadata(m)
	cloudConnection := readCloudConnectionFromResource(d)

	result, err := c.CreateCloudConnection(ctx, cloudConnection)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*result.ID)
	return resourceCloudConnectionRead(ctx, d, m)
}

func resourceCloudConnectionUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	var diags diag.Diagnostics
	c := getLocalClientFromMetadata(m)
	id := d.Id()
	cloudConnection := readCloudConnectionFromResource(d)

	if id == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "could not obtain cloud connection id from resource",
		})
	}

	if utils.IsBlank(cloudConnection.Name) {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "could not obtain name from resource",
		})
	}

	if diags != nil {
		return diags
	}

	cloudConnection.ID = &id

	if err := c.UpdateCloudConnection(ctx, cloudConnection); err != nil {
		return diag.FromErr(err)
	}

	return resourceCloudConnectionRead(ctx, d, m)
}

func resourceCloudConnectionDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := getLocalClientFromMetadata(m)
	id := d.Id()

	err := c.DeleteCloudConnection(ctx, id)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readCloudConnectionFromResource(d *schema.ResourceData) client.CloudConnection {
	var cloudConnectionDetails *client.CloudConnectionDetails

	if v, ok := d.GetOk(attributes.CloudConnectionDetails); ok {
		list := v.([]any)
		detailsMap := list[0].(map[string]any)
		cloudConnectionDetails = &client.CloudConnectionDetails{
			AccountId:  utils.AsStringPtr(detailsMap[attributes.CloudConnectionAccountId].(string)),
			ExternalId: utils.AsStringPtr(detailsMap[attributes.CloudConnectionExternalId].(string)),
			RoleArn:    utils.AsStringPtr(detailsMap[attributes.CloudConnectionRoleARN].(string)),
		}
	}

	cloudConnection := client.CloudConnection{
		Name:                   GetStringPtrFromResource(attributes.Name, d, true),
		Provider:               GetStringPtrFromResource(attributes.CloudConnectionProvider, d, true),
		CloudConnectionDetails: cloudConnectionDetails,
	}

	return cloudConnection
}
