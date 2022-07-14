package attributes

const (
	AccessAddress                    = "access_address"
	AccessAddressAttribute           = "access_address_attribute"
	ADConnectionID                   = "connection_id"
	ADConnections                    = "ad_connections"
	AdditionalAttributeMapping       = "additional_attribute_mapping"
	ADRuleAssignments                = "rule_assignments"
	ADRuleAssignmentsBaseDN          = "base_dn"
	ADRuleAssignmentsLDAPQueryFilter = "ldap_query_filter"
	ADRuleAssignmentsProjectID       = "project_id"
	ADRuleAssignmentsPriority        = "priority"
	AltNamesAttributes               = "alt_names_attributes"
	BastionAttribute                 = "bastion_attribute"
	CertificateID                    = "certificate_id"
	Claims                           = "claims"
	CloudProvider                    = "cloud_provider"
	ClusterID                        = "cluster_id"
	ClusterSelector                  = "cluster_selector"
	CommonName                       = "common_name"
	Contains                         = "contains"
	Content                          = "content"
	Country                          = "country"
	CreateServerGroup                = "create_server_group"
	CreateServerUsers                = "create_server_users"
	CreatedAt                        = "created_at"
	CreatedByUser                    = "created_by_user"
	DefaultAddress                   = "default_address"
	DeletedAt                        = "deleted_at"
	Description                      = "description"
	DescriptionContains              = "description_contains"
	Details                          = "details"
	DisconnectedModeOnOnly           = "disconnected_mode_on_only"
	DisplayName                      = "display_name"
	Domain                           = "domain"
	DomainControllers                = "domain_controllers"
	EnterpriseSigned                 = "enterprise_signed"
	ExpiresAt                        = "expires_at"
	FileName                         = "file_name"
	ForwardTraffic                   = "forward_traffic"
	Frequency                        = "frequency"
	GatewayID                        = "gateway_id"
	Gateways                         = "gateways"
	GatewaySelector                  = "gateway_selector"
	Groups                           = "groups"
	GroupID                          = "group_id"
	GroupName                        = "group_name"
	HasSelectors                     = "has_selectors"
	HasNoSelectors                   = "has_no_selectors"
	HostnameAttribute                = "host_name_attribute"
	ID                               = "id"
	IDs                              = "ids"
	IncludeCertDetails               = "include_cert_details"
	IncludeDeleted                   = "include_deleted"
	IncludeRemoved                   = "include_removed"
	IsActive                         = "is_active"
	IsGuid                           = "is_guid"
	IssuedAt                         = "issued_at"
	KubernetesAPIURL                 = "api_url"
	KubernetesAuthMechanism          = "auth_mechanism"
	KubernetesClusterKey             = "key"
	Label                            = "label"
	Labels                           = "labels"
	Locality                         = "locality"
	Name                             = "name"
	NextUnixGID                      = "next_unix_gid"
	NextUnixUID                      = "next_unix_uid"
	OIDCIssuerURL                    = "oidc_issuer_url"
	OnlyIncludeDeleted               = "only_include_deleted"
	Organization                     = "organization"
	OrganizationalUnit               = "organizational_unit"
	OSAttribute                      = "os_attribute"
	ProjectGroups                    = "project_groups"
	ProjectName                      = "project_name"
	Projects                         = "projects"
	Province                         = "province"
	PublicCertificate                = "public_certificate"
	RDPSessionRecording              = "rdp_session_recording"
	RefuseConnections                = "refuse_connections"
	RequirePreauthForCreds           = "require_preauth_for_creds"
	RemovedAt                        = "removed_at"
	Roles                            = "roles"
	RunTest                          = "run_test"
	Self                             = "self"
	ServerAccess                     = "server_access"
	ServerAdmin                      = "server_admin"
	ServerEnrollmentTokens           = "server_enrollment_tokens"
	ServersSelector                  = "servers_selector"
	ServiceAccountUsername           = "service_account_username"
	ServiceAccountPassword           = "service_account_password"
	Source                           = "source"
	SSHCertificateType               = "ssh_certificate_type"
	SSHSessionRecording              = "ssh_session_recording"
	StartHourUTC                     = "start_hour_utc"
	Status                           = "status"
	Team                             = "team" // NOTE: This is inconsistent, most other API endpoints use `team_name`.
	Token                            = "token"
	TTLDays                          = "ttl_days"
	Type                             = "type"
	UsePasswordless                  = "use_passwordless"
	Value                            = "value"
)
