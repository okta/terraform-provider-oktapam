package descriptions

const (
	// Descriptions -- conventions from API docs

	// Attribute Descriptions
	CreateServerGroup = "If 'true', `sftd` (ASA Server Agent) creates a corresponding local (unix or windows) group in the ASA Project's servers."
	CreateServerUsers = "If 'true', `sftd` (ASA Server Agent) creates corresponding local (unix or windows) user accounts in the ASA Project's servers."
	CreatedAt = "UTC time of resource deletion."
	CreatedByUser = "The ASA User that created the resource."
	DeletedAt = "UTC time of resource creation."
	Description = "The human-readable description of the resource."
	DescriptionContains = "If a value is provided, the results are filtered to only contain resources whose description contains that value."
	ForwardTraffic = "If 'true', all traffic in the ASA Project be forwarded through selected ASA Gateways."
	GatewaySelector = "Assigns ASA Gateways with labels matching all selectors. At least one selector is necessary for traffic forwarding."
	GroupId = "The ID corresponding to ASA Group."
	GroupName = "The human-readable name of the ASA Group"
	IssuedAt = "The UTC issuance time of the resource."
	Labels = "A map of key-value pairings."
	Name = "The human-readable name of the resource."
	NextUnixGID = "The UID to use when creating a new ASA Server User."
	NextUnixUID = "The GID to use when creating a new ASA Server User."
	ProjectName = "The human-readable name of the ASA Project."
	RdpSessionRecording = "If 'true', enable remote desktop protocol (RDP) recording on all servers in the ASA Project."
	RemovedAt = "UTC time of resource removal from parent resource."
	RequirePreauthForCreds = "If 'true', require preauthorization before an ASA User can retrieve credentials to sign in."
	Roles = "A list of roles for the ASA Group. Options are `access_user`, `access_admin`, and `reporting_user`."
	ServerAccess = "If 'true', members of this ASA Group have access to the ASA Project servers."
	ServerAdmin = "If 'true', members of ASA Group have sudo permissions on ASA Project servers."
	ServersSelector = "Enables access to ASA Servers with labels matching all selectors. For ASA Projects Groups using Policy Sync Feature."
	SshCertificateType = "The SSH certificate type used by access requests."
	SshSessionRecording = "If 'true', enables ssh recording on server access requests."
	TeamName = "The human-readable name of the ASA Team that owns the resource."
	Token = "The secret used for resource enrollment."

	// Query Parameter Descriptions
	FilterContains = "If a value is provided, the results are filtered to only contain resources whose name contains that value."
	FilterCreateServerGroup = "If 'true', the results only include the ASA Project Groups that have 'create_server_group' field set to 'true'."
	FilterDescriptionContains = "If a value is provided, the results are filtered to only contain resources whose name contains that value."
	FilterDisconnectedModeOnOnly = "If 'true', the results only include resources with disconnected mode enabled."
	FilterHasNoSelectors = "If 'true', the results only include resources with empty label selectors."
	FilterHasSelectors = "If 'true', the results only include resources with label selectors set."
	FilterIncludeDeleted = "If 'true', the results include deleted resources."
	FilterOnlyIncludeDeleted = "If 'true', the results only include deleted resources."
	FilterIncludeRemoved = "If 'true', the results include removed resources."
	FilterOfflineEnabled = "If 'true', the results only include resources with disconnected mode enabled." // NOTE: This is inconsistent, most other API endpoints use `disconnected_mode_on_only`.
	FilterProjectName = "If a value is provided, the results are filtered to only contain resources belonging to the ASA Project."
	FilterSelf = "If 'true', only lists the ASA Projects that the ASA User making this request has been assigned."

	// Resource Descriptions -- resources are dynamic, and state is kept up to date on POST / PUT / DELETE
	ResourceGatewaySetupToken = "A token for ASA Gateway enrollment."
	ResourceGroup = "A set of ASA Users."
	ResourceServerEnrollmentToken = "A token for ASA Server enrollment."

	// Data Source Descriptions -- sources are read-only, fetched on LIST
	SourceGatewaySetupTokens = "A list of tokens for ASA Gateway enrollment, corresponding to an ASA Team."
	SourceGroups = "A list of ASA Groups, corresponding to an ASA Team."
	SourceProjects = "A list of ASA Projects, corresponding to an ASA Team."
	SourceProjectGroups = "A list of ASA Groups attached to an ASA Project."
	SourceServerEnrollmentTokens = "A list of tokens for ASA Server enrollment, corresponding to an ASA Project."
)