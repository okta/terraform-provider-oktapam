package constants

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

const (
	// Alphabetically sorted attributes
	AttrContains = "contains"
	AttrCreatedAt = "created_at"
	AttrDeletedAt = "deleted_at"
	AttrDescription = "description"
	AttrDescriptionContains = "description_contains"
	AttrDisconnectedModeOnOnly = "disconnected_mode_on_only"
	AttrId = "id"
	AttrIncludeDeleted = "include_deleted"
	AttrLabels = "labels"
	AttrName = "name"
	AttrOnlyIncludeDeleted = "only_include_deleted"
	AttrRoles = "roles"

	// Alphabetically sorted resources
	ResGatewaySetupTokens = "gateway_setup_tokens"
	ResGroups = "groups"

)

var AttributeSchemas = map[string]*schema.Schema{
	AttrContains: {
		Type:     schema.TypeString,
		Description: "Filter for resource",
		Optional: true,
	},
	AttrCreatedAt: {
		Type:     schema.TypeString,
		Description: "UTC creation time of resource",
		Computed: true,
	},
	AttrDeletedAt: {
		Type:     schema.TypeString,
		Description: "UTC deletion time of resource",
		Computed: true,
	},
	AttrDescription: {
		Type:     schema.TypeString,
		Description: "Human-readable resource description",
		Computed: true,
	},
	AttrDescriptionContains: {
		Type:     schema.TypeString,
		Description: "Filter for resource description",
		Optional: true,
	},
	AttrDisconnectedModeOnOnly: {
		Type:     schema.TypeBool,
		Description: "Only include resources with disconnected mode enabled",
		Optional: true,
	},
	AttrId: {
		Type:     schema.TypeString,
		// Auto-populates description
		Computed: true,
	},
	AttrIncludeDeleted: {
		Type:     schema.TypeBool,
		Description: "Include deleted resources",
		Optional: true,
	},
	AttrLabels: {
		Type: schema.TypeMap,
		Description: "",
		Elem: &schema.Schema{
			Type: schema.TypeString,
			Description: "",
		},
		Computed: true,
	},
	AttrName: {
		Type:     schema.TypeString,
		Description: "Human-readable resource name",
		Computed: true,
	},
	AttrOnlyIncludeDeleted: {
		Type:     schema.TypeBool,
		Description: "Only include deleted resources",
		Optional: true,
	},
	AttrRoles: {
		Type:     schema.TypeList,
		Description: "A list of roles for the ASA Group. Options are `access_user`, `access_admin`, and `reporting_user`.",
		Elem:     schema.TypeString,
		Computed: true,
	},
}

var ResourceSchemas = map[string]*schema.Schema{
	ResGatewaySetupTokens: {
		Type:        schema.TypeList,
		Description: "A Gateway enrollment token",
		Computed:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				AttrId:          AttributeSchemas[AttrId],
				AttrCreatedAt:   AttributeSchemas[AttrCreatedAt],
				AttrDescription: AttributeSchemas[AttrDescription],
				AttrLabels:      AttributeSchemas[AttrLabels],
			},
		},
	},
	ResGroups: {
		Type:        schema.TypeList,
		Description: "A set of ASA Users",
		Computed:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				AttrName:      AttributeSchemas[AttrName],
				AttrId:        AttributeSchemas[AttrId],
				AttrDeletedAt: AttributeSchemas[AttrDeletedAt],
				AttrRoles: {
					Type:        schema.TypeList,
					Description: "A list of roles for the ASA Group. Options are `access_user`, `access_admin`, and `reporting_user`.",
					Elem:        schema.TypeString,
					Computed:    true,
				},
			},
		},
	},
}