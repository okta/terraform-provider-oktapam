package oktapam

import (
	"context"
	"fmt"
	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/client/wrappers"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
	"strings"
)

const (
	MySqlBasicAuth = "mysql.basic_auth"
)

var managementConnectionDetails = &schema.Resource{
	Schema: map[string]*schema.Schema{
		attributes.Hostname: {
			Type:        schema.TypeString,
			Required:    true,
			Description: descriptions.DatabaseHostname,
		},
		attributes.Port: {
			Type:        schema.TypeString,
			Required:    true,
			Description: descriptions.DatabasePort,
		},
		attributes.AuthDetails: {
			Type:        schema.TypeList,
			Required:    true,
			MinItems:    1,
			MaxItems:    1,
			Description: descriptions.DatabaseAuthDetails,
			ForceNew:    false,
			Elem:        mySQLBasicAuthDetails,
		},
	},
}

var mySQLBasicAuthDetails = &schema.Resource{
	Schema: map[string]*schema.Schema{
		attributes.Username: {
			Type:        schema.TypeString,
			Required:    true,
			Description: descriptions.UserName,
		},
		attributes.Password: {
			Type:        schema.TypeString,
			Optional:    true,
			Description: descriptions.DatabasePassword,
			Sensitive:   true,
		},
		attributes.Secret: {
			Type:        schema.TypeString,
			Description: descriptions.DatabaseSecretID,
			Computed:    true,
		},
	}}

func resourceDatabase() *schema.Resource {
	return &schema.Resource{
		Description:   descriptions.ResourceDatabase,
		CreateContext: resourceDatabaseCreate,
		ReadContext:   resourceDatabaseRead,
		UpdateContext: resourceDatabaseUpdate,
		DeleteContext: resourceDatabaseDelete,
		Schema: map[string]*schema.Schema{
			attributes.ID: {
				Type:     schema.TypeString,
				Computed: true,
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
			attributes.CanonicalName: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions.CanonicalName,
			},
			attributes.DatabaseType: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions.DatabaseType,
			},
			attributes.RecipeBook: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions.RecipeBookID,
			},
			attributes.ManagementConnectionDetails: {
				Type:        schema.TypeList,
				Required:    true,
				MinItems:    1,
				MaxItems:    1,
				Description: descriptions.ManagementConnectionDetails,
				Elem:        managementConnectionDetails,
				ForceNew:    false,
			},
			attributes.ManagementConnectionDetailsType: {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions.ManagementConnectionDetailsType,
			},
			attributes.ManagementGatewaySelector: {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: descriptions.ManagementGatewaySelector,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			attributes.ManagementGatewaySelectorID: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: descriptions.ManagementGatewaySelectorID,
			},
			attributes.CreatedAt: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: descriptions.CreatedAt,
			},
			attributes.UpdatedAt: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: descriptions.UpdatedAt,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: resourceDatabaseReadImport,
		},
	}
}

func resourceDatabaseCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := getSDKClientFromMetadata(m)
	resourceGroupID := d.Get(attributes.ResourceGroup).(string)
	projectID := d.Get(attributes.Project).(string)
	canonicalName := d.Get(attributes.CanonicalName).(string)
	dbType := d.Get(attributes.DatabaseType).(string)

	var selectorLabels *map[string]string
	if selector, ok := d.GetOk(attributes.ManagementGatewaySelector); ok {
		if s, ok := selector.(map[string]any); ok {
			tmpLabels := make(map[string]string, len(s))
			for k, v := range s {
				if value, ok := v.(string); ok {
					tmpLabels[k] = value
				}
			}
			selectorLabels = &tmpLabels
		} else {
			return diag.FromErr(fmt.Errorf("invalid label selector"))
		}
	}

	mgmtType := d.Get(attributes.ManagementConnectionDetailsType).(string)
	mgmtDetails, plainTextPassword, err := mgmtConnectionDetailsFromResource(ctx, c, mgmtType, d)
	if err != nil {
		return err
	}

	if createdDb, err := client.CreateDatabase(ctx, c, resourceGroupID, projectID, canonicalName, dbType, mgmtType, *mgmtDetails, selectorLabels); err != nil {
		return diag.FromErr(err)
	} else if createdDb == nil {
		d.SetId("")
	} else {
		d.SetId(createdDb.Id)
	}

	return resourceDatabaseReadWithPassword(ctx, d, m, plainTextPassword)
}

func resourceDatabaseRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	return resourceDatabaseReadWithPassword(ctx, d, m, nil)
}

func resourceDatabaseReadWithPassword(ctx context.Context, d *schema.ResourceData, m any, newPassword *string) diag.Diagnostics {
	c := getSDKClientFromMetadata(m)
	resourceGroupID := d.Get(attributes.ResourceGroup).(string)
	projectID := d.Get(attributes.Project).(string)
	databaseID := d.Get(attributes.ID).(string)
	if databaseID == "" {
		return diag.Errorf("%s cannot be blank", attributes.ID)
	}

	database, err := client.GetDatabase(ctx, c, resourceGroupID, projectID, databaseID)
	if err != nil {
		return diag.FromErr(err)
	}

	wrap := wrappers.DatabaseResourceResponseWrapper{*database}

	overrides := make(map[string]any, 1)
	if newPassword != nil && *newPassword != "" {
		overrides[attributes.Password] = *newPassword
	} else if password, ok := d.Get(attributes.NestedManagementConnectionPassword).(string); ok {
		// if a refresh happens the password may exist in state
		overrides[attributes.Password] = password
	}

	d.SetId(database.Id)
	for key, value := range wrap.ToResourceMap(overrides) {
		if err := d.Set(key, value); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceDatabaseUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	var diags diag.Diagnostics

	c := getSDKClientFromMetadata(m)
	dbID := d.Get(attributes.ID).(string)
	resourceGroupID := d.Get(attributes.ResourceGroup).(string)
	projectID := d.Get(attributes.Project).(string)
	canonicalName := d.Get(attributes.CanonicalName).(string)
	dbType := d.Get(attributes.DatabaseType).(string)

	var selectorLabels *map[string]string
	if selector, ok := d.GetOk(attributes.ManagementGatewaySelector); ok {
		if s, ok := selector.(map[string]any); ok {
			tmpLabels := make(map[string]string, len(s))
			for k, v := range s {
				if value, ok := v.(string); ok {
					tmpLabels[k] = value
				}
			}
			selectorLabels = &tmpLabels
		} else {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "invalid label selector",
			})
		}
	}

	mgmtType := d.Get(attributes.ManagementConnectionDetailsType).(string)
	mgmtDetails, plainTextPassword, err := mgmtConnectionDetailsFromResource(ctx, c, mgmtType, d)
	if err != nil {
		return err
	}

	if dbID == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "could not obtain database id from resource",
		})
	}
	if resourceGroupID == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "could not obtain resource group id from resource",
		})
	}
	if projectID == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "could not obtain project id from resource",
		})
	}

	if diags != nil {
		return diags
	}

	if err := client.UpdateDatabase(ctx, c, dbID, resourceGroupID, projectID, canonicalName, dbType, mgmtType, *mgmtDetails, selectorLabels); err != nil {
		return diag.FromErr(err)
	}

	return resourceDatabaseReadWithPassword(ctx, d, m, plainTextPassword)
}

func resourceDatabaseDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := getSDKClientFromMetadata(m)
	resourceGroupID := d.Get(attributes.ResourceGroup).(string)
	projectID := d.Get(attributes.Project).(string)
	databaseID := d.Get(attributes.ID).(string)
	if databaseID == "" {
		return diag.Errorf("%s cannot be blank", attributes.ID)
	}

	err := client.DeleteDatabase(ctx, c, resourceGroupID, projectID, databaseID)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceDatabaseReadImport(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	resourceGroupID, projectID, databaseID, err := parseDatabaseID(d.Id())
	if err != nil {
		return nil, err
	}
	if err := d.Set(attributes.ResourceGroup, resourceGroupID); err != nil {
		return nil, err
	}
	if err := d.Set(attributes.Project, projectID); err != nil {
		return nil, err
	}
	d.SetId(databaseID)

	return []*schema.ResourceData{d}, nil
}

func parseDatabaseID(resourceId string) (string, string, string, error) {
	split := strings.Split(resourceId, "/")
	if len(split) != 3 {
		return "", "", "", fmt.Errorf("expected format: <resource_group_id>/<project_id>/<id>, received: %s", resourceId)
	}
	return split[0], split[1], split[2], nil
}

func mgmtConnectionDetailsFromResource(_ context.Context, client client.SDKClientWrapper, mgmtType string, d *schema.ResourceData) (*pam.ManagementConnectionDetails, *string, diag.Diagnostics) {
	var plainTextPassword *string
	mgmtDetails := &pam.ManagementConnectionDetails{}

	if v, ok := d.GetOk(attributes.ManagementConnectionDetails); !ok {
		return nil, nil, diag.Errorf("%s must have one set of details", attributes.ManagementConnectionDetails)
	} else {
		details, ok := v.([]any)
		if !ok {
			return nil, nil, diag.Errorf("%s must have one set of details", attributes.ManagementConnectionDetails)
		}

		detailsMap, ok := details[0].(map[string]any)
		if !ok {
			return nil, nil, diag.Errorf("invalid %s", attributes.ManagementConnectionDetails)
		}

		switch mgmtType {
		case MySqlBasicAuth:
			var diags diag.Diagnostics
			authDetails := detailsMap[attributes.AuthDetails]
			authDetailsSet, ok := authDetails.([]any)
			if !ok {
				return nil, nil, diag.Errorf("invalid %s", attributes.AuthDetails)
			}
			authDetailsMap, ok := authDetailsSet[0].(map[string]any)
			if !ok {
				return nil, nil, diag.Errorf("invalid %s", attributes.AuthDetails)
			}

			hostname, ok := detailsMap[attributes.Hostname].(string)
			if !ok {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "invalid hostname",
				})
			}
			port, ok := detailsMap[attributes.Port].(string)
			if !ok {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "invalid port",
				})
			}

			var secretID *string
			if secret, ok := authDetailsMap[attributes.Secret].(string); ok {
				secretID = &secret
			}

			username, ok := authDetailsMap[attributes.Username].(string)
			if !ok {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "invalid username",
				})
			}

			var passwordJwe *pam.EncryptedString
			if password, ok := authDetailsMap[attributes.Password].(string); ok && password != "" {
				var err error
				plainTextPassword = &password
				passwordJwe, err = client.SDKClient.Encrypt(password)
				if err != nil {
					return nil, nil, diag.Errorf("error encrypting password: %v", err)
				}
			} else if !ok {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "invalid auth details",
				})
			}

			if diags.HasError() {
				return nil, nil, diags
			}

			mgmtDetails.MySQLBasicAuthManagementConnectionDetails = &pam.MySQLBasicAuthManagementConnectionDetails{
				Hostname: hostname,
				Port:     port,
				AuthDetails: pam.MySQLBasicAuthDetails{
					Username:    username,
					PasswordJwe: passwordJwe,
					SecretId:    secretID,
				},
			}
		default:
			return nil, nil, diag.Errorf("invalid %s: %s", attributes.ManagementConnectionDetailsType, mgmtType)
		}
	}
	return mgmtDetails, plainTextPassword, nil
}
