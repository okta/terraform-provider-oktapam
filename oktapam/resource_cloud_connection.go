package oktapam

import (
	"context"
	"fmt"

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
			attributes.CloudConnectionDetails: {
				Type:        schema.TypeList,
				MinItems:    1,
				MaxItems:    1,
				Required:    true,
				ForceNew:    true,
				Description: descriptions.CloudConnectionDetails,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						attributes.AWS: {
							Type: schema.TypeList,
							// Note this is required since we only have one
							// cloud connection type.  When we support more providers
							// this will need to change
							Required:    true,
							MinItems:    1,
							MaxItems:    1,
							Description: descriptions.CloudConnectionDetailsAWS,
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
		d.SetId("")
		return diag.Errorf("cloud connection does not exist")
	}

	for key, value := range cloudConnection.ToResourceMap() {
		logging.Debugf("setting %s to %v", key, value)
		if err := d.Set(key, value); err != nil {
			return diag.FromErr(err)
		}
	}

	// TODO: change this to parse the details out of the new structure

	if details, ok := d.Get(attributes.CloudConnectionDetails).([]map[string]any); ok && len(details) == 1 {
		flattenedDetails := make([]any, 1)
		flattenedDetail := make(map[string]any)
		flattenedDetail[attributes.CloudConnectionAccountId] = details[0][attributes.CloudConnectionAccountId]
		flattenedDetail[attributes.CloudConnectionExternalId] = details[0][attributes.CloudConnectionExternalId]
		flattenedDetail[attributes.CloudConnectionRoleARN] = details[0][attributes.CloudConnectionRoleARN]

		flattenedDetails[0] = flattenedDetail
		if err := d.Set(attributes.CloudConnectionDetails, flattenedDetails); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceCloudConnectionCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := getLocalClientFromMetadata(m)

	var cloudConnectionDetails *client.CloudConnectionDetails
	var cloudConnectionProvider string
	if v, ok := d.GetOk(attributes.CloudConnectionDetails); ok {
		list := v.([]any)
		detailsMap := list[0].(map[string]any)

		if len(detailsMap) != 1 {
			return diag.FromErr(fmt.Errorf("expected only one cloud connection details"))
		}

		for providerName, values := range detailsMap {
			cloudConnectionProvider = providerName
			switch cloudConnectionProvider {
			case attributes.AWS:
				awsDetailsA := values.([]any)
				awsDetails := awsDetailsA[0].(map[string]any)
				cloudConnectionDetails = &client.CloudConnectionDetails{
					AccountId:  utils.AsStringPtr(awsDetails[attributes.CloudConnectionAccountId].(string)),
					ExternalId: utils.AsStringPtr(awsDetails[attributes.CloudConnectionExternalId].(string)),
					RoleArn:    utils.AsStringPtr(awsDetails[attributes.CloudConnectionRoleARN].(string)),
				}
			default:
				return diag.FromErr(fmt.Errorf("unknown cloud connection provider: %s", cloudConnectionProvider))
			}
		}

	} else {
		return diag.Errorf("cloud connection details missing")
	}

	ccReq := client.CloudConnection{
		Name:                   GetStringPtrFromResource(attributes.Name, d, false),
		Provider:               &cloudConnectionProvider,
		CloudConnectionDetails: cloudConnectionDetails,
	}

	if createdCloudConnection, err := c.CreateCloudConnection(ctx, ccReq); err != nil {
		return diag.FromErr(err)
	} else if createdCloudConnection == nil {
		d.SetId("")
	} else {
		d.SetId(*createdCloudConnection.ID)
	}

	return resourceCloudConnectionRead(ctx, d, m)
}

func resourceCloudConnectionUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := getLocalClientFromMetadata(m)
	id := d.Id()

	changed := false
	updates := make(map[string]any)

	changeableAttributes := []string{
		attributes.Name,
		attributes.CloudConnectionProvider,
		attributes.CloudConnectionRoleARN,
	}

	for _, attribute := range changeableAttributes {
		if d.HasChange(attribute) {
			updates[attribute] = d.Get(attribute)
			changed = true
		}
	}

	if changed {
		if err := c.UpdateCloudConnection(ctx, id, updates); err != nil {
			return diag.FromErr(err)
		}
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
