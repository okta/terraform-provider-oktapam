package attributes

const (
	AccessAddress                     = "access_address"
	AccessAddressAttribute            = "access_address_attribute"
	AccessRequest                     = "access_request"
	Account                           = "account"
	AccountDiscovery                  = "account_discovery"
	Accounts                          = "accounts"
	ACRValues                         = "acr_values"
	Active                            = "active"
	ADConnectionID                    = "connection_id"
	ADConnections                     = "ad_connections"
	AddEnv                            = "add_env"
	AdditionalAttributeMapping        = "additional_attribute_mapping"
	AdminLevelPermissions             = "admin_level_permissions"
	ADRuleAssignments                 = "rule_assignments"
	ADRuleAssignmentsBaseDN           = "base_dn"
	ADRuleAssignmentsLDAPQueryFilter  = "ldap_query_filter"
	ADRuleAssignmentsPriority         = "priority"
	ADRuleAssignmentsProjectID        = "project_id"
	ADUserSyncTaskSettingsIDList      = "ad_user_sync_task_settings_id_list"
	ADUserSyncTaskSettingsStatus      = "status"
	AltNamesAttributes                = "alt_names_attributes"
	ApproveDeviceWithoutInteraction   = "approve_device_without_interaction"
	AWS                               = "aws"
	BaseDN                            = "base_dn"
	BasicAuth                         = "basic_auth"
	BastionAttribute                  = "bastion_attribute"
	CanonicalName                     = "canonical_name"
	CertificateID                     = "certificate_id"
	CharacterOptions                  = "character_options"
	Claims                            = "claims"
	ClientSessionDuration             = "client_session_duration"
	CloudConnectionAccountId          = "account_id"
	CloudConnectionDetails            = "cloud_connection_details"
	CloudConnectionExternalId         = "external_id"
	CloudConnectionProvider           = "cloud_connection_provider"
	CloudConnectionRoleARN            = "role_arn"
	CloudProvider                     = "cloud_provider"
	ClusterID                         = "cluster_id"
	ClusterSelector                   = "cluster_selector"
	CommonName                        = "common_name"
	Conditions                        = "conditions"
	Contains                          = "contains"
	Content                           = "content"
	Country                           = "country"
	CreatedAt                         = "created_at"
	CreatedByUser                     = "created_by_user"
	CreateServerGroup                 = "create_server_group"
	CreateServerUsers                 = "create_server_users"
	Database                          = "database"
	DatabaseID                        = "database_id"
	DatabaseLabels                    = "database_labels"
	DatabasePasswordSettings          = "database_password_settings"
	Databases                         = "databases"
	DatabaseType                      = "database_type"
	DefaultAddress                    = "default_address"
	DelegatedResourceAdminGroups      = "delegated_resource_admin_groups"
	DeletedAt                         = "deleted_at"
	Description                       = "description"
	DescriptionContains               = "description_contains"
	Details                           = "details"
	Digits                            = "digits"
	DisconnectedModeOnOnly            = "disconnected_mode_on_only"
	DisplayName                       = "display_name"
	Domain                            = "domain"
	DomainControllers                 = "domain_controllers"
	Enabled                           = "enabled"
	EnablePeriodicRotation            = "enable_periodic_rotation"
	EnterpriseSigned                  = "enterprise_signed"
	ExpiresAfterSeconds               = "expires_after_seconds"
	ExpiresAt                         = "expires_at"
	FileName                          = "file_name"
	FilterStatus                      = "filter_status" // NOTE: This attribute exists to avoid conflicts with Status
	FolderCreate                      = "folder_create"
	FolderDelete                      = "folder_delete"
	FolderUpdate                      = "folder_update"
	ForwardTraffic                    = "forward_traffic"
	Frequency                         = "frequency"
	Gateway                           = "gateway"
	GatewayID                         = "gateway_id"
	Gateways                          = "gateways"
	GatewaySelector                   = "gateway_selector"
	Group                             = "group"
	GroupName                         = "group_name"
	GroupNames                        = "group_names"
	Groups                            = "groups"
	HasNoSelectors                    = "has_no_selectors"
	HasSelectors                      = "has_selectors"
	Hostname                          = "hostname"
	HostnameAttribute                 = "host_name_attribute"
	ID                                = "id"
	IDs                               = "ids"
	IncludeCertDetails                = "include_cert_details"
	IncludeDeleted                    = "include_deleted"
	IncludeRemoved                    = "include_removed"
	IncludeServiceUsers               = "include_service_users"
	IncludeUserSID                    = "include_user_sid"
	IsActive                          = "is_active"
	IsGuid                            = "is_guid"
	IssuedAt                          = "issued_at"
	KubernetesAPIURL                  = "api_url"
	KubernetesAuthMechanism           = "auth_mechanism"
	KubernetesClusterKey              = "key"
	Label                             = "label"
	Labels                            = "labels"
	LabelSelectors                    = "label_selectors"
	LDAPQueryFilter                   = "ldap_query_filter"
	List                              = "list"
	ListElementsUnderPath             = "list_elements_under_path"
	Locality                          = "locality"
	LowerCase                         = "lower_case"
	ManagedPrivilegedAccounts         = "managed_privileged_accounts"
	ManagementConnectionDetails       = "management_connection_details"
	ManagementGatewaySelector         = "management_gateway_selector"
	ManagementGatewaySelectorID       = "management_gateway_selector_id"
	MaxLength                         = "max_length"
	MFA                               = "mfa"
	MinLength                         = "min_length"
	MySQL                             = "mysql"
	Name                              = "name"
	Names                             = "names"
	NextUnixGID                       = "next_unix_gid"
	NextUnixUID                       = "next_unix_uid"
	NoExec                            = "no_exec"
	NoPasswd                          = "no_passwd"
	OIDCIssuerURL                     = "oidc_issuer_url"
	OnlyIncludeDeleted                = "only_include_deleted"
	Organization                      = "organization"
	OrganizationalUnit                = "organizational_unit"
	OSAttribute                       = "os_attribute"
	ParentFolder                      = "parent_folder"
	Password                          = "password"
	PasswordCheckoutDatabase          = "password_checkout_database"
	PasswordCheckoutRDP               = "password_checkout_rdp"
	PasswordCheckoutSSH               = "password_checkout_ssh"
	PasswordSettings                  = "password_settings"
	Path                              = "path"
	PeriodicRotationDurationInSeconds = "periodic_rotation_duration_in_seconds"
	Port                              = "port"
	Postgres                          = "postgres"
	PrincipalAccountRDP               = "principal_account_rdp"
	PrincipalAccountSSH               = "principal_account_ssh"
	Principals                        = "principals"
	Privileges                        = "privileges"
	Project                           = "project"
	ProjectGroups                     = "project_groups"
	ProjectID                         = "project_id"
	ProjectName                       = "project_name"
	ProjectNames                      = "project_names"
	Projects                          = "projects"
	Province                          = "province"
	PublicCertificate                 = "public_certificate"
	Punctuation                       = "punctuation"
	RDPSessionRecording               = "rdp_session_recording"
	ReactivateUsersViaIDP             = "reactivate_users_via_idp"
	ReAuthFrequencyInSeconds          = "reauth_frequency_in_seconds"
	RecipeBook                        = "recipe_book"
	RefuseConnections                 = "refuse_connections"
	RemovedAt                         = "removed_at"
	RequestTypeId                     = "request_type_id"
	RequestTypeName                   = "request_type_name"
	RequireFromEachSet                = "require_from_each_set"
	RequirePreauthForCreds            = "require_preauth_for_creds"
	ResourceGroup                     = "resource_group"
	ResourceGroupID                   = "resource_group_id"
	Resources                         = "resources"
	Roles                             = "roles"
	Rule                              = "rule"
	RunAs                             = "run_as"
	RunTest                           = "run_test"
	Secret                            = "secret"
	SecretCreate                      = "secret_create"
	SecretDelete                      = "secret_delete"
	SecretFolder                      = "secret_folder"
	SecretFolderID                    = "secret_folder_id"
	SecretFolders                     = "secret_folders"
	SecretID                          = "secret_id"
	SecretReveal                      = "secret_reveal"
	Secrets                           = "secrets"
	SecretUpdate                      = "secret_update"
	Self                              = "self"
	Server                            = "server"
	ServerAccess                      = "server_access"
	ServerAccount                     = "server_account"
	ServerAdmin                       = "server_admin"
	ServerEnrollmentTokens            = "server_enrollment_tokens"
	ServerID                          = "server_id"
	ServerLabels                      = "server_labels"
	Servers                           = "servers"
	ServersSelector                   = "servers_selector"
	ServerUserName                    = "server_user_name"
	ServiceAccountPassword            = "service_account_password"
	ServiceAccountUsername            = "service_account_username"
	SessionRecording                  = "session_recording"
	SetEnv                            = "set_env"
	SIDField                          = "sid_field"
	Source                            = "source"
	SSHCertificateType                = "ssh_certificate_type"
	SSHSessionRecording               = "ssh_session_recording"
	StartHourUTC                      = "start_hour_utc"
	StartsWith                        = "starts_with"
	Status                            = "status"
	StructuredCommand                 = "command"
	StructuredCommandArgs             = "args"
	StructuredCommandArgsType         = "args_type"
	StructuredCommands                = "structured_commands"
	StructuredCommandType             = "command_type"
	StructuredRenderedCommand         = "rendered_command"
	SubEnv                            = "sub_env"
	SudoCommandBundles                = "sudo_command_bundles"
	SudoDisplayName                   = "sudo_display_name"
	Team                              = "team" // NOTE: This is inconsistent, most other API endpoints use `team_name`.
	TeamName                          = "team_name"
	Token                             = "token"
	TrafficForwarding                 = "traffic_forwarding"
	TTLDays                           = "ttl_days"
	Type                              = "type"
	UPNField                          = "upn_field"
	UpperCase                         = "upper_case"
	UsePasswordless                   = "use_passwordless"
	Username                          = "username"
	Usernames                         = "usernames"
	UserOnDemandPeriod                = "user_on_demand_period"
	UserProvisioningExactUserName     = "user_provisioning_exact_username"
	Users                             = "users"
	UserType                          = "user_type"
	Value                             = "value"
	WebSessionDuration                = "web_session_duration"
)
