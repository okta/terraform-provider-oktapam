package descriptions

const (
	// Descriptions -- conventions from API docs

	// Attribute Descriptions
	ClusterGroupClaims          = "A map of claims to be given to users in this Cluster Group."
	ClusterSelector             = "A label selector to used to match Kubernetes clusters."
	CreateServerGroup           = "If 'true', `sftd` (ASA Server Agent) creates a corresponding local (unix or windows) group in the ASA Project's servers."
	CreateServerUsers           = "If 'true', `sftd` (ASA Server Agent) creates corresponding local (unix or windows) user accounts in the ASA Project's servers."
	CreatedAt                   = "The UTC time of resource deletion. Format is `2022-01-01 00:00:00 +0000 UTC`."
	CreatedByUser               = "The ASA User that created the resource."
	DeletedAt                   = "The UTC time of resource creation. Format is `2022-01-01 00:00:00 +0000 UTC`."
	Description                 = "The human-readable description of the resource."
	DescriptionContains         = "If a value is provided, the results are filtered to only contain resources whose description contains that value."
	ForwardTraffic              = "If 'true', all traffic in the ASA Project be forwarded through selected ASA Gateways."
	GatewaySelector             = "Assigns ASA Gateways with labels matching all selectors. At least one selector is required for traffic forwarding."
	GroupID                     = "The ID corresponding to a ASA Group."
	GroupName                   = "The human-readable name of the ASA Group. Values are case-sensitive."
	IssuedAt                    = "The UTC issuance time of the resource. Format is `2022-01-01 00:00:00 +0000 UTC`."
	KubernetesAuthMechanism     = "Mechanism to provide auth details to your Kubernetes cluster (eg. OIDC_RSA2048, NONE)"
	KubernetesAPIURL            = "The URL to access the control plane of the Kubernetes cluster."
	KubernetesClusterID         = "The ASA ID of a Kubernetes cluster."
	KubernetesClusterKey        = "The human-friendly key to associate with the Kubernetes cluster. Must be simple alphanumeric without spaces."
	KubernetesClusterLabels     = "Map of labels to assign to the Kubernetes cluster."
	KubernetesPublicCertificate = "The certificate expected when connecting to the Kubernetes cluster."
	Labels                      = "A map of key-value pairings that define access to the ASA Gateway."
	Name                        = "The human-readable name of the resource. Values are case-sensitive."
	NextUnixGID                 = "The GID to use when creating a new ASA Server User. Default value starts at 63001."
	NextUnixUID                 = "The UID to use when creating a new ASA Server User. Default value starts at 60001."
	OIDCIssuerURL               = "The OIDC Issuer URL to use when configuring your Kubernetes cluster. "
	ProjectName                 = "The human-readable name of the ASA Project. Values are case-sensitive."
	RDPSessionRecording         = "If 'true', enable remote desktop protocol (RDP) recording on all servers in the ASA Project."
	RemovedAt                   = "UTC time of resource removal from parent resource. Format is `2022-01-01 00:00:00 +0000 UTC`."
	RequirePreauthForCreds      = "If 'true', require preauthorization before an ASA User can retrieve credentials to sign in."
	Roles                       = "A list of roles for the ASA Group. Options are `access_user`, `access_admin`, and `reporting_user`."
	ServerAccess                = "If 'true', members of this ASA Group have access to the ASA Project servers."
	ServerAdmin                 = "If 'true', members of ASA Group have sudo permissions on ASA Project servers."
	ServersSelector             = "Enables access to ASA Servers with labels matching all selectors. For ASA Projects Groups using Policy Sync Feature."
	SSHCertificateType          = "The SSH certificate type used by access requests. Options include: `CERT_TYPE_ED25519_01`, `CERT_TYPE_RSA_01`. `CERT_TYPE_RSA_01` is a deprecated key algorithm type. " +
		"This option should only be used to connect to legacy systems that cannot use newer SSH versions. If you do need to use `CERT_TYPE_RSA_01`, it is recommended to connect via a gateway with traffic forwarding. " +
		"Otherwise, please use a more current key algorithm. If left unspecified, `CERT_TYPE_ED25519_01` is used by default."
	SSHSessionRecording              = "If 'true', enables ssh recording on server access requests."
	TeamName                         = "The human-readable name of the ASA Team that owns the resource. Values are lower-case."
	Token                            = "The secret used for resource enrollment."
	GatewayID                        = "The UUID of the Gateway with which this AD Connection is associated."
	Domain                           = "The domain against which to query."
	DomainControllers                = "A comma-separated list of the specific domain controller(s) that should be used to query the domain. Can be specified as a hostname or IP."
	ServiceAccountUsername           = "The username of the service account that can be used to query the domain."
	ServiceAccountPassword           = "The password of the service account that can be used to query the domain."
	UsePasswordless                  = "if 'true', Users will not need password to login."
	CertificateID                    = "Certificate Id used for password less access method."
	ADTaskFrequency                  = "Frequency of the AD Task"
	ADTaskIsActive                   = "If 'true', enables AD task"
	ADTaskRunTest                    = "If 'true', test is performed based on specified AD Task Settings"
	ADTaskStartHourUTC               = "If AD task is scheduled to run daily, then specify start hour in UTC"
	HostnameAttribute                = "AD Attribute mapped to hostname used to identify a discovered server within Advanced Server Access."
	AccessAddressAttribute           = "AD Attribute mapped to IP address or DNS name used by the gateway to connect to a discovered server."
	OSAttribute                      = "AD Attribute mapped to server operating system of a discovered server."
	BastionAttribute                 = "AD Attribute mapped to bastion host that Advanced Server Access clients can use to tunnel traffic to a discovered server."
	AltNamesAttributes               = "AD Attribute mapped to alternative hostnames or DNS entries used to resolve a discovered server."
	AdditionalAttributeMapping       = "Additional AD attributes mappings to Advanced Server Access labels."
	AdditionalAttributeMappingLabel  = "ASA label"
	AdditionalAttributeMappingValue  = "AD Attribute mapped to ASA label"
	AdditionalAttributeMappingIsGuid = "Identifies an AD attribute as a Globally Unique Identifier (GUID)"
	ADRuleAssignments                = "Assignment rules determine how servers are synced from Active Directory (AD) and assigned to projects."
	ADRuleAssignmentsBaseDN          = "Specifies where the rule searches for servers."
	ADRuleAssignmentsLDAPQueryFilter = "Specifies the specific criteria used to filter servers."
	ADRuleAssignmentsProjectID       = "Specifies a project to associate with matching servers"
	ADConnectionID                   = "UUID of the AD Connection with which this AD Task Settings is associated."
	AccessAddress                    = "Access Address of the gateway."
	DefaultAddress                   = "Default Address of the gateway."
	CloudProvider                    = "Cloud Provider name of the host where gateway is running."
	RefuseConnections                = "If 'true', gateway refuse connection."
	CertificateCommonName            = "Common Name or FQDN to which certificate is issued to."
	CSRDetails                       = "Certificate Signing Request (CSR) Details."
	// Query Parameter Descriptions
	FilterContains               = "If a value is provided, the results are filtered to only contain resources whose name contains that value."
	FilterCreateServerGroup      = "If 'true', the results only include the ASA Project Groups that have 'create_server_group' field set to 'true'."
	FilterDescriptionContains    = "If a value is provided, the results are filtered to only contain resources whose name contains that value."
	FilterDisconnectedModeOnOnly = "If 'true', the results only include resources with disconnected mode enabled."
	FilterHasNoSelectors         = "If 'true', the results only include resources with empty label selectors."
	FilterHasSelectors           = "If 'true', the results only include resources with label selectors set."
	FilterIncludeDeleted         = "If 'true', the results include deleted resources."
	FilterOnlyIncludeDeleted     = "If 'true', the results only include deleted resources."
	FilterIncludeRemoved         = "If 'true', the results include removed resources."
	FilterOfflineEnabled         = "If 'true', the results only include resources with disconnected mode enabled." // NOTE: This is inconsistent, most other API endpoints use `disconnected_mode_on_only`.
	FilterProjectName            = "If a value is provided, the results are filtered to only contain resources belonging to the ASA Project."
	FilterSelf                   = "If 'true', only lists the ASA Projects that the ASA User making this request has been assigned."
	FilterGatewayID              = "If 'true', the results only include the connections which has the same gateway id."
	FilterCertificateID          = "If 'true', the results only include the connections which has the same certificate id."
	FilterIncludeCertDetails     = "If 'true', the results include the certificate details "

	// Resource Descriptions -- resources are dynamic, and state is kept up to date on POST / PUT / DELETE
	ResourceGatewaySetupToken           = "A token for ASA Gateway enrollment."
	ResourceGroup                       = "A set of ASA Users."
	ResourceADConnection                = "An Active Directory (AD) Connection to query AD Domain for available servers."
	ResourceADTaskSettings              = "Settings for Active Directory (AD) task for server discovery"
	ResourceKubernetesCluster           = "A Kubernetes cluster."
	ResourceKubernetesClusterConnection = "A set of details describing how to connect to an existing Kubernetes Cluster."
	ResourceKubernetesClusterGroup      = "A mapping of Kubernetes cluster to Project Group."
	ResourceServerEnrollmentToken       = "A token for ASA Server enrollment."

	// Data Source Descriptions -- sources are read-only, fetched on LIST
	SourceGatewaySetupTokens     = "A list of tokens for ASA Gateway enrollment, corresponding to an ASA Team."
	SourceGroups                 = "A list of ASA Groups, corresponding to an ASA Team."
	SourceProjects               = "A list of ASA Projects, corresponding to an ASA Team."
	SourceProjectGroups          = "A list of ASA Groups attached to an ASA Project."
	SourceServerEnrollmentTokens = "A list of tokens for ASA Server enrollment, corresponding to an ASA Project."
	SourceADConnections          = "A list of ASA AD Connections, corresponding to an ASA Team."
	SourceGateways               = "A list of ASA Gateways, corresponding to an ASA Team."
	CertificateSigningRequest    = "Certificate Signing Request (CSR) in base 64 encoded pkcs#10 format."
)
