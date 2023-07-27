package attributes

const (
	AccessAddress                     = "access_address"
	AccessAddressAttribute            = "access_address_attribute"
	AccessRequest                     = "access_request"
	Account                           = "account"
	Accounts                          = "accounts"
	AccountDiscovery                  = "account_discovery"
	Active                            = "active"
	ADConnectionID                    = "connection_id"
	ADConnections                     = "ad_connections"
	AdditionalAttributeMapping        = "additional_attribute_mapping"
	ADRuleAssignments                 = "rule_assignments"
	ADRuleAssignmentsBaseDN           = "base_dn"
	ADRuleAssignmentsLDAPQueryFilter  = "ldap_query_filter"
	ADRuleAssignmentsProjectID        = "project_id"
	ADRuleAssignmentsPriority         = "priority"
	ADUserSyncTaskSettingsIDList      = "ad_user_sync_task_settings_id_list"
	ADUserSyncTaskSettingsStatus      = "status"
	AltNamesAttributes                = "alt_names_attributes"
	ApproveDeviceWithoutInteraction   = "approve_device_without_interaction"
	BaseDN                            = "base_dn"
	BastionAttribute                  = "bastion_attribute"
	CertificateID                     = "certificate_id"
	CharacterOptions                  = "character_options"
	Claims                            = "claims"
	ClientSessionDuration             = "client_session_duration"
	CloudProvider                     = "cloud_provider"
	ClusterID                         = "cluster_id"
	ClusterSelector                   = "cluster_selector"
	CommonName                        = "common_name"
	Conditions                        = "conditions"
	Contains                          = "contains"
	Content                           = "content"
	Country                           = "country"
	CreateServerGroup                 = "create_server_group"
	CreateServerUsers                 = "create_server_users"
	CreatedAt                         = "created_at"
	CreatedByUser                     = "created_by_user"
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
	ForwardTraffic                    = "forward_traffic"
	Frequency                         = "frequency"
	FullAdminAccess                   = "full_admin_access"
	Gateway                           = "gateway"
	GatewayID                         = "gateway_id"
	Gateways                          = "gateways"
	GatewaySelector                   = "gateway_selector"
	Groups                            = "groups"
	GroupName                         = "group_name"
	GroupNames                        = "group_names"
	HasSelectors                      = "has_selectors"
	HasNoSelectors                    = "has_no_selectors"
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
	Locality                          = "locality"
	LowerCase                         = "lower_case"
	ManagedPrivilegedAccounts         = "managed_privileged_accounts"
	MaxLength                         = "max_length"
	MinLength                         = "min_length"
	Name                              = "name"
	Names                             = "names"
	NextUnixGID                       = "next_unix_gid"
	NextUnixUID                       = "next_unix_uid"
	OIDCIssuerURL                     = "oidc_issuer_url"
	OnlyIncludeDeleted                = "only_include_deleted"
	Organization                      = "organization"
	OrganizationalUnit                = "organizational_unit"
	OSAttribute                       = "os_attribute"
	PasswordCheckoutRDP               = "password_checkout_rdp"
	PasswordCheckoutSSH               = "password_checkout_ssh"
	PasswordSettings                  = "password_settings"
	PeriodicRotationDurationInSeconds = "periodic_rotation_duration_in_seconds"
	PrincipalAccountRDP               = "principal_account_rdp"
	PrincipalAccountSSH               = "principal_account_ssh"
	Principals                        = "principals"
	Privileges                        = "privileges"
	ProjectGroups                     = "project_groups"
	Project                           = "project"
	ProjectID                         = "project_id"
	ProjectName                       = "project_name"
	ProjectNames                      = "project_names"
	Projects                          = "projects"
	Province                          = "province"
	PublicCertificate                 = "public_certificate"
	Punctuation                       = "punctuation"
	ReactivateUsersViaIDP             = "reactivate_users_via_idp"
	RDPSessionRecording               = "rdp_session_recording"
	RefuseConnections                 = "refuse_connections"
	RequestTypeId                     = "request_type_id"
	RequestTypeName                   = "request_type_name"
	RequirePreauthForCreds            = "require_preauth_for_creds"
	RemovedAt                         = "removed_at"
	ResourceGroup                     = "resource_group"
	ResourceGroupID                   = "resource_group_id"
	Resources                         = "resources"
	Roles                             = "roles"
	Rule                              = "rule"
	RunTest                           = "run_test"
	Self                              = "self"
	ServerAccess                      = "server_access"
	ServerAdmin                       = "server_admin"
	ServerEnrollmentTokens            = "server_enrollment_tokens"
	ServerUserName                    = "server_user_name"
	Server                            = "server"
	ServerAccount                     = "server_account"
	ServerID                          = "server_id"
	ServerLabels                      = "server_labels"
	Servers                           = "servers"
	ServersSelector                   = "servers_selector"
	ServiceAccountUsername            = "service_account_username"
	ServiceAccountPassword            = "service_account_password"
	SessionRecording                  = "session_recording"
	SIDField                          = "sid_field"
	Source                            = "source"
	SSHCertificateType                = "ssh_certificate_type"
	SSHSessionRecording               = "ssh_session_recording"
	StartHourUTC                      = "start_hour_utc"
	StartsWith                        = "starts_with"
	Status                            = "status"
	Team                              = "team" // NOTE: This is inconsistent, most other API endpoints use `team_name`.
	TeamName                          = "team_name"
	Token                             = "token"
	TTLDays                           = "ttl_days"
	TrafficForwarding                 = "traffic_forwarding"
	Type                              = "type"
	UpperCase                         = "upper_case"
	UPNField                          = "upn_field"
	UsePasswordless                   = "use_passwordless"
	UserOnDemandPeriod                = "user_on_demand_period"
	UserProvisioningExactUserName     = "user_provisioning_exact_username"
	Username                          = "username"
	Usernames                         = "usernames"
	Users                             = "users"
	UserType                          = "user_type"
	Value                             = "value"
	WebSessionDuration                = "web_session_duration"
)
