package attributes

const (
	Claims                  = "claims"
	ClusterID               = "cluster_id"
	ClusterSelector         = "cluster_selector"
	Contains                = "contains"
	CreateServerGroup       = "create_server_group"
	CreateServerUsers       = "create_server_users"
	CreatedAt               = "created_at"
	CreatedByUser           = "created_by_user"
	DeletedAt               = "deleted_at"
	Description             = "description"
	DescriptionContains     = "description_contains"
	DisconnectedModeOnOnly  = "disconnected_mode_on_only"
	ForwardTraffic          = "forward_traffic"
	GatewaySelector         = "gateway_selector"
	GatewaySetupTokens      = "gateway_setup_tokens"
	Groups                  = "groups"
	GroupID                 = "group_id"
	GroupName               = "group_name"
	HasSelectors            = "has_selectors"
	HasNoSelectors          = "has_no_selectors"
	ID                      = "id"
	IncludeDeleted          = "include_deleted"
	IncludeRemoved          = "include_removed"
	IssuedAt                = "issued_at"
	KubernetesAPIURL        = "api_url"
	KubernetesAuthMechanism = "auth_mechanism"
	KubernetesClusterKey    = "key"
	Labels                  = "labels"
	Name                    = "name"
	NextUnixGID             = "next_unix_gid"
	NextUnixUID             = "next_unix_uid"
	OIDCIssuerURL           = "oidc_issuer_url"
	OnlyIncludeDeleted      = "only_include_deleted"
	ProjectGroups           = "project_groups"
	ProjectName             = "project_name"
	Projects                = "projects"
	PublicCertificate       = "public_certificate"
	RDPSessionRecording     = "rdp_session_recording"
	RequirePreauthForCreds  = "require_preauth_for_creds"
	RemovedAt               = "removed_at"
	Roles                   = "roles"
	Self                    = "self"
	ServerAccess            = "server_access"
	ServerAdmin             = "server_admin"
	ServerEnrollmentTokens  = "server_enrollment_tokens"
	ServersSelector         = "servers_selector"
	SSHCertificateType      = "ssh_certificate_type"
	SSHSessionRecording     = "ssh_session_recording"
	Team                    = "team" // NOTE: This is inconsistent, most other API endpoints use `team_name`.
	Token                   = "token"
)
