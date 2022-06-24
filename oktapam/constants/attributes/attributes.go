package attributes

const (
	Contains               = "contains"
	CreateServerGroup      = "create_server_group"
	CreateServerUsers      = "create_server_users"
	CreatedAt              = "created_at"
	CreatedByUser          = "created_by_user"
	DeletedAt              = "deleted_at"
	Description            = "description"
	DescriptionContains    = "description_contains"
	DisconnectedModeOnOnly = "disconnected_mode_on_only"
	FilterStatus           = "filter_status" // NOTE: This attribute exists to avoid conflicts with Status
	ForwardTraffic         = "forward_traffic"
	GatewaySelector        = "gateway_selector"
	GatewaySetupTokens     = "gateway_setup_tokens"
	Groups                 = "groups"
	GroupID                = "group_id"
	GroupName              = "group_name"
	GroupUsers             = "group_users"
	HasSelectors           = "has_selectors"
	HasNoSelectors         = "has_no_selectors"
	ID                     = "id"
	IncludeDeleted         = "include_deleted"
	IncludeRemoved         = "include_removed"
	IncludeServiceUsers    = "include_service_users"
	IssuedAt               = "issued_at"
	Labels                 = "labels"
	Name                   = "name"
	NextUnixGID            = "next_unix_gid"
	NextUnixUID            = "next_unix_uid"
	OnlyIncludeDeleted     = "only_include_deleted"
	ProjectGroups          = "project_groups"
	ProjectName            = "project_name"
	Projects               = "projects"
	RDPSessionRecording    = "rdp_session_recording"
	RequirePreauthForCreds = "require_preauth_for_creds"
	RemovedAt              = "removed_at"
	Roles                  = "roles"
	Self                   = "self"
	ServerAccess           = "server_access"
	ServerAdmin            = "server_admin"
	ServerEnrollmentTokens = "server_enrollment_tokens"
	ServerUserName         = "server_user_name"
	ServersSelector        = "servers_selector"
	ServiceUsers           = "service_users"
	SSHCertificateType     = "ssh_certificate_type"
	SSHSessionRecording    = "ssh_session_recording"
	StartsWith             = "starts_with"
	Status                 = "status"
	Team                   = "team" // NOTE: This is inconsistent, most other API endpoints use `team_name`.
	TeamName               = "team_name"
	Token                  = "token"
)
