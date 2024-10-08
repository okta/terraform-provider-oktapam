package oktapam

import (
	"context"
	"fmt"
	"strings"

	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/client/wrappers"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"
)

const (
	MySqlBasicAuth = "mysql.basic_auth"
)

var dbConnectionTypes = &schema.Resource{
	Schema: map[string]*schema.Schema{
		attributes.MySQL: {
			Type:        schema.TypeList,
			Required:    true,
			MinItems:    1,
			MaxItems:    1,
			Description: descriptions.MySqlManagementConnectionDetails,
			Elem:        mysqlConnectionDetails,
		},
	},
}

var mysqlConnectionDetails = &schema.Resource{
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
		attributes.BasicAuth: {
			Type: schema.TypeList,
			// If any new auth methods are created, we should utilize ConflictsWith to ensure only one is set at a time.
			//ConflictsWith: []string{parent_block_name.0.child_attribute_name},
			Required:    true,
			MinItems:    1,
			MaxItems:    1,
			Description: descriptions.DatabaseAuthDetails,
			Elem:        mySQLBasicAuthDetails,
		},
	},
}

var databaseRoles = &schema.Resource{
	Schema: map[string]*schema.Schema{
		attributes.Name: {
			Type:        schema.TypeString,
			Required:    true,
			Description: descriptions.DatabaseRoleName,
		},
		attributes.Type: {
			Type:        schema.TypeString,
			Required:    true,
			Description: descriptions.DatabaseRoleType,
		},
		attributes.Accounts: {
			Type:        schema.TypeSet,
			Required:    false,
			MinItems:    1,
			Optional:    true,
			Description: descriptions.DatabaseRoleAccounts,
			Elem:        &schema.Schema{Type: schema.TypeString},
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
			attributes.NetworkAddress: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions.NetworkAddress,
			},
			attributes.DatabaseType: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions.DatabaseType,
				ForceNew:    true,
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
				Elem:        dbConnectionTypes,
				ForceNew:    true,
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
			attributes.Role: {
				Type:        schema.TypeList,
				Optional:    true,
				Description: descriptions.DatabaseRole,
				Elem:        databaseRoles,
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

	selectorLabels, err := GetMapPtrFromResource[string](attributes.ManagementGatewaySelector, d)
	if err != nil {
		return err
	}

	mgmtType, mgmtDetails, err := mgmtConnectionDetailsFromResource(ctx, c, d)
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

	return resourceDatabaseRead(ctx, d, m)
}

func resourceDatabaseRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := getSDKClientFromMetadata(m)
	resourceGroupID := d.Get(attributes.ResourceGroup).(string)
	projectID := d.Get(attributes.Project).(string)
	databaseID := d.Get(attributes.ID).(string)

	database, err := client.GetDatabase(ctx, c, resourceGroupID, projectID, databaseID)
	if err != nil {
		return diag.FromErr(err)
	}

	wrap := wrappers.DatabaseResourceResponseWrapper{DatabaseResourceResponse: *database}
	attributeOverrides := utils.GenerateAttributeOverrides(d, wrap)
	for key, value := range wrap.ToResourceMap(attributeOverrides) {
		if err := d.Set(key, value); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceDatabaseUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := getSDKClientFromMetadata(m)
	dbID := d.Get(attributes.ID).(string)
	resourceGroupID := d.Get(attributes.ResourceGroup).(string)
	projectID := d.Get(attributes.Project).(string)
	canonicalName := d.Get(attributes.CanonicalName).(string)
	dbType := d.Get(attributes.DatabaseType).(string)

	selectorLabels, err := GetMapPtrFromResource[string](attributes.ManagementGatewaySelector, d)
	if err != nil {
		return err
	}

	mgmtType, mgmtDetails, err := mgmtConnectionDetailsFromResource(ctx, c, d)
	if err != nil {
		return err
	}

	if err := client.UpdateDatabase(ctx, c, dbID, resourceGroupID, projectID, canonicalName, dbType, mgmtType, *mgmtDetails, selectorLabels); err != nil {
		return diag.FromErr(err)
	}

	return resourceDatabaseRead(ctx, d, m)
}

func resourceDatabaseDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := getSDKClientFromMetadata(m)
	resourceGroupID := d.Get(attributes.ResourceGroup).(string)
	projectID := d.Get(attributes.Project).(string)
	databaseID := d.Get(attributes.ID).(string)

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

func mgmtConnectionDetailsFromResource(_ context.Context, pamClient client.SDKClientWrapper, d *schema.ResourceData) (string, *pam.ManagementConnectionDetails, diag.Diagnostics) {
	detailsMap := GetTypeListMapFromResource(attributes.ManagementConnectionDetails, d)
	var mgmtDetailsMap map[string]any
	var databaseType string
	// schema guarantees there is only one set of details stored.
	for key, val := range detailsMap {
		databaseType = key
		mgmtDetailsList := val.([]any)
		mgmtDetailsMap = mgmtDetailsList[0].(map[string]any)
	}

	var mgmtDetails *pam.ManagementConnectionDetails
	var diags diag.Diagnostics

	switch databaseType {
	case attributes.MySQL:
		mgmtDetails, diags = getMysqlDetailsFromResource(pamClient, mgmtDetailsMap)
		if diags.HasError() {
			return "", nil, diags
		}
	default:
		return "", nil, diag.Errorf("invalid management connection details, provided %s", databaseType)
	}

	mgmtType := getDatabaseManagementType(mgmtDetails)

	return mgmtType, mgmtDetails, nil
}

func getDatabaseManagementType(details *pam.ManagementConnectionDetails) string {
	if details.MySQLBasicAuthManagementConnectionDetails != nil {
		return MySqlBasicAuth
	}
	return ""
}

func getMysqlDetailsFromResource(pamClient client.SDKClientWrapper, detailsMap map[string]any) (*pam.ManagementConnectionDetails, diag.Diagnostics) {
	var diags diag.Diagnostics
	var plainTextPassword *string

	connectionDetails := &pam.ManagementConnectionDetails{}

	// required fields
	hostname := detailsMap[attributes.Hostname].(string)
	port := detailsMap[attributes.Port].(string)

	// look for possible auth details
	if auth, authSet := detailsMap[attributes.BasicAuth]; authSet {
		var username string
		//var secretID *string
		var passwordJwe *pam.EncryptedString

		authList := auth.([]any)
		authMap := authList[0].(map[string]any)

		// required params
		username = authMap[attributes.Username].(string)

		// optional params
		secretID := GetStringPtrFromElement(attributes.Secret, authMap, false)

		plainTextPassword = GetStringPtrFromElement(attributes.Password, authMap, false)
		if plainTextPassword != nil && *plainTextPassword != "" {
			var err error
			passwordJwe, err = pamClient.SDKClient.Encrypt(*plainTextPassword)
			if err != nil {
				return nil, diag.Errorf("error encrypting password: %v", err)
			}
		}

		if diags.HasError() {
			return nil, diags
		}

		connectionDetails.MySQLBasicAuthManagementConnectionDetails = &pam.MySQLBasicAuthManagementConnectionDetails{
			Hostname: hostname,
			Port:     port,
			AuthDetails: pam.MySQLBasicAuthDetails{
				Username:    username,
				PasswordJwe: passwordJwe,
				SecretId:    secretID,
			},
		}
	}

	return connectionDetails, nil
}
