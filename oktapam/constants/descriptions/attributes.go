package descriptions

import (
	"fmt"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/typed_strings"
)

var (
	// Attribute Descriptions
	ADConnectionID                            = "The UUID of an associated AD connection."
	ADRuleAssignments                         = "The rules used to assign discovered servers to Projects."
	ADRuleAssignmentsBaseDN                   = "The domain to search for servers."
	ADRuleAssignmentsLDAPQueryFilter          = "The criteria used to filter discovered servers."
	ADRuleAssignmentsProjectID                = "The UUID for an existing Project. Discovered servers that match the rules are assigned to this Project."
	ADTaskFrequency                           = "Indicates how often the server sync job runs. Possible values: `1`, `6`, `12`, `24`."
	ADTaskIsActive                            = "If `true`, the AD server sync job is enabled."
	ADTaskRunTest                             = "If `true`, performs a test run for the server sync job."
	ADTaskStartHourUTC                        = "A UTC timestamp that indicates the hour range when the server sync job runs. Only used if `frequency` is set to 24."
	ADUserSyncTaskSettingsBaseDN              = "Specifies the domain to search for user accounts."
	ADUserSyncTaskSettingsFrequency           = "Indicates how often the user sync job runs. Possible values: `1`, `6`, `12`, `24`."
	ADUserSyncTaskSettingsIsActive            = "If `true`, enables the user sync job."
	ADUserSyncTaskSettingsLDAPQueryFilter     = "The criteria used to filter user accounts."
	ADUserSyncTaskSettingsRunTest             = "If `true`, performs a test run for the user sync job."
	ADUserSyncTaskSettingsSIDField            = "The AD attribute that defines the security identifier (SID) for accounts. Most AD tenants use `objectSID`."
	ADUserSyncTaskSettingsStartHourUTC        = "A UTC timestamp that indicates the hour range when the user sync job runs. Only used if `frequency`is set to 24."
	ADUserSyncTaskSettingsUPNField            = "The AD attribute that defines the User Principal Name (UPN) for accounts. Most AD tenants use `userPrincipalName`."
	AccessAddress                             = "The IP address used to access the Gateway."
	AccessAddressAttribute                    = "The AD attribute that defines an IP address or DNS name for a server. This is used by the Gateway to connect to discovered servers."
	AccountDiscovery                          = "If `true`, will enable the discovery of local accounts on servers within the project."
	AdditionalAttributeMapping                = "Additional AD attributes used to map information to specified labels."
	AdditionalAttributeMappingIsGuid          = "If `true`, the AD attribute is identified as a Globally Unique Identifier (GUID)"
	AdditionalAttributeMappingLabel           = "An existing ASA label"
	AdditionalAttributeMappingValue           = "The AD attribute to map to the specified ASA label"
	AdminLevelPermissions                     = "Provides coarse grain (full admin) access to the user."
	AltNamesAttributes                        = "The AD attribute that defines alternative hostnames or DNS entries. These are used to resolve discovered servers."
	ApproveDeviceWithoutInteraction           = "If `true`, devices are automatically approved for authenticated Users."
	BastionAttribute                          = "The AD attribute that defines the bastion host for a server. Clients use this bastion to tunnel traffic to discovered servers.."
	CSRDetails                                = "Specific details on the Certificate Signing Request (CSR)."
	CanonicalName                             = "The nickname or alias of the resource."
	CertificateCommonName                     = "The Common Name or FQDN associated with the certificate."
	CertificateContent                        = "The returned Certificate Signing Request (CSR) or Self Signed Certificate."
	CertificateID                             = "The UUID for a Certificate used for Passwordless authentication."
	CertificateRequestType                    = "Specifies the type of certificate request - Certificate Signing Request (CSR)/ Self Signed Certificate."
	CertificateStatus                         = "Certificate status - Valid/Request Created."
	CharacterOptions                          = "The specific characters rules required by the Password Policy."
	CharacterOptionsDigits                    = "If `true`, passwords can include one or more numeric characters."
	CharacterOptionsLowerCase                 = "If `true`, passwords can include one or more lowercase characters."
	CharacterOptionsPunctuation               = "If `true`, passwords can include one or more punctuation/symbol characters."
	CharacterOptionsRequireFromEachSet        = "If `true`, passwords must include at least one character from the selected sets."
	CharacterOptionsUpperCase                 = "If `true`, passwords can include one or more uppercase characters."
	ClientSessionDuration                     = "The maximum time before a Client session expires. The duration can be from 3600 – 90000 seconds."
	CloudProvider                             = "The name of the cloud provider where the Gateway is hosted."
	ClusterGroupClaims                        = "A map of claims included with issued user credentials used to authenticate to Kubernetes clusters. Claims correspond to pre-configured role bindings on the cluster."
	ClusterSelector                           = "A label selector to used to match Kubernetes clusters."
	ConditionAccessRequest                    = "Identifies an existing Request Type in Access Requests."
	ConditionAccessRequestExpiresAfterSeconds = "The number of seconds the approval remains valid."
	ConditionAccessRequestRequestTypeID       = "The ID of the Access Request."
	ConditionAccessRequestRequestTypeName     = "The name of the Access Request."
	ConditionGateway                          = "Configures traffic settings for an existing Gateway."
	ConditionMFA                              = "Configures multifactor auth settings required to access the resource."
	ConditionSessionRecording                 = "Whether to record sessions made through a Gateway."
	ConditionTrafficForwarding                = "Whether to forward traffic through a Gateway."
	Conditions                                = "The conditions required before a privilege is made available to the Principals. All conditions must be met."
	CreateServerGroup                         = "If true, creates a local (unix or windows) group on enrolled Project servers and adds all relevant users to the local group. This does not impact user creation. Server Users are always created as long as `create_server_users` is set to `true` for the Project."
	CreateServerUsers                         = "If `true`, creates local (unix or windows) user accounts on enrolled Project servers."
	CreatedAt                                 = "The UTC time when the resource was created. Format is '2022-01-01 00:00:00 +0000 UTC'."
	CreatedByUser                             = "The User that originally created the resource."
	Database                                  = "Defines a specific database targeted by the Security Policy."
	DatabaseAuthDetails                       = "A set of fields required to authenticate to the database."
	DatabaseHostname                          = "The hostname used to connect to the database"
	DatabasePassword                          = "The password used to authenticate to the database."
	DatabasePort                              = "The port used to connect to the database"
	DatabaseSecretID                          = "The UUID of the secret corresponding to the stored database password."
	DatabaseType                              = "Defines the type of database used and the feature set supported. A two-part string separated by a dot: '<db_engine>.level<level>'"
	DefaultAddress                            = "The default IP address used to access the gateway. This is generally supplied by the network interface or cloud provider metadata."
	DelegatedAdminGroups                      = "The UUIDs of the Groups designated as Delegated Admins for the Resource Group."
	DeletedAt                                 = "The UTC time when the resource was deleted. Format is '2022-01-01 00:00:00 +0000 UTC'."
	Description                               = "The human-readable description of the resource."
	DescriptionContains                       = "If defined, results only include resources that include the specified value in the `description` field."
	Domain                                    = "The domain against which to query."
	DomainControllers                         = "A comma-separated list of the specific domain controllers used to query the domain. Can be specified as a hostname or IP."
	EnablePeriodicRotation                    = "If `true`, requires passwords to be rotated after a period of time has passed. You must also set the `periodic_rotation_duration_in_seconds` param."
	EnterpriseSigned                          = "If `true`, indicates the certificate is signed by AD Certificate Services."
	ForwardTraffic                            = "If `true`, all traffic in the Project is forwarded through selected Gateways."
	GatewayID                                 = "The UUID of a Gateway associated with this AD Connection."
	GatewaySelector                           = "Assigns Gateways with labels matching all selectors. At least one selector is required to forward traffic through a Gateway."
	GroupName                                 = "The human-readable name of an existing Group. Values are case-sensitive."
	HostnameAttribute                         = "The AD attribute that defines the hostname for a server. This is used to identify discovered servers."
	IncludeUserSID                            = "Options for including the User Security Identifier (SID) in the certificate used for Passwordless authentication to a Windows server via RDP. Possible values: `Never`, `If_Available`, and `Always`."
	IssuedAt                                  = "The UTC time when the token was issued. Format is '2022-01-01 00:00:00 +0000 UTC'."
	KubernetesAPIURL                          = "The URL used to access the control plane of the Kubernetes cluster."
	KubernetesAuthMechanism                   = "The mechanism used to provide auth details to your Kubernetes cluster (eg. OIDC_RSA2048, NONE)"
	KubernetesClusterID                       = "The UUID of an Kubernetes cluster."
	KubernetesClusterKey                      = "The human-friendly key associated with the Kubernetes cluster. Must be simple alphanumeric without spaces."
	KubernetesClusterLabels                   = "Map of labels to assign to the Kubernetes cluster."
	KubernetesPublicCertificate               = "The certificate expected when connecting to the Kubernetes cluster."
	Labels                                    = "A map of key-value pairings that define access to a Gateway."
	ListElementsUnderPath                     = "If `true`, returns a list of any Secret/Secret Folder elements under the path. If `false`, returns the element defined by the path."
	LocalAccount                              = "The username of a local account on a server."
	LocalAccounts                             = "The usernames of a local accounts on a server."
	MFAACRValues                              = "The authentication context class reference (ACR) for this policy. This defines a specific set of assurance level requirements required by a protected resource."
	MFAReAuthFrequencyInSeconds               = "The number of seconds an MFA verification remains valid. After this time users need to reauthenticate before they can open new connections to resources. Defining a value of `0` requires users to authenticate for every connection."
	ManagedPrivilegedAccounts                 = "An array of managed accounts for password rotation."
	ManagementConnectionDetails               = "A set of fields defining the database to connect to."
	ManagementGatewaySelector                 = "A label selector to define which gateway(s) will be used to connect to the database."
	ManagementGatewaySelectorID               = "The ID of the selector."
	MySqlManagementConnectionDetails          = "A set of fields defining how to connect to a mysql database."
	Name                                      = "The human-readable name of the resource. Values are case-sensitive."
	NextUnixGID                               = "The GID to use when creating a new Server User. Default value starts at 63001."
	NextUnixUID                               = "The UID to use when creating a new Server User. Default value starts at 60001."
	OIDCIssuerURL                             = "The OIDC Issuer URL to use when configuring your Kubernetes cluster. "
	OSAttribute                               = "The AD attribute that defines the operating system of discovered servers."
	ParentFolder                              = "The directory which contains this Secret/Secret Folder element."
	PasswordMaxLength                         = "The maximum length allowed for the password."
	PasswordMinLength                         = "The minimum length allowed for the password."
	PeriodicRotationDurationInSeconds         = "If `periodic_rotation` is enabled, specifies how often passwords are rotated."
	PrincipalGroupIDs                         = "The UUIDs of existing Groups."
	PrivilegeEnabled                          = "If `true`, grants the privilege to Principals on matching resources."
	PrivilegePasswordCheckoutRDP              = "Defines the privilege to make RDP connections to a server with a vaulted password."
	PrivilegePasswordCheckoutSSH              = "Defines the privilege to make SSH connections to a server with a vaulted password."
	PrivilegePrincipalAccountRDP              = "Defines the privilege to make RDP connections to a server with the user's principal account."
	PrivilegePrincipalAccountSSH              = "Defines the privilege to make SSH connections to a server with the user's principal account."
	PrivilegeSecret                           = "Defines the privilege to operate on Secrets and Secret Folders."
	PrivilegeSecretCreate                     = "Defines the privilege to create a Secret."
	PrivilegeSecretDelete                     = "Defines the privilege to delete a Secret."
	PrivilegeSecretFolderCreate               = "Defines the privilege to create a Secret Folder."
	PrivilegeSecretFolderDelete               = "Defines the privilege to delete a Secret Folder and its contents."
	PrivilegeSecretFolderUpdate               = "Defines the privilege to update the metadata of a Secret Folder."
	PrivilegeSecretList                       = "Defines the privilege to list the contents of a Secret Folder."
	PrivilegeSecretReveal                     = "Defines the privilege to reveal the plaintext contents of a Secret."
	PrivilegeSecretUpdate                     = "Defines the privilege to update a Secret and its metadata."
	Privileges                                = "The specific privileges granted to Principals on matching resources."
	ProjectID                                 = "The UUID of a Project."
	ProjectName                               = "The human-readable name of the Project. Values are case-sensitive."
	RDPSessionRecording                       = "If `true`, enables remote desktop protocol (RDP) recording on all servers in the Project."
	ReactivateUsersViaIDP                     = "If a disabled or deleted User is able to authenticate through the IdP, their user settings are re-enabled."
	RecipeBookID                              = "The ID of a recipe book which will override the db queries used."
	RefuseConnections                         = "If `true`, the Gateway refuses connections."
	RemovedAt                                 = "The UTC time when the resource was removed from parent resource. Format is '2022-01-01 00:00:00 +0000 UTC'."
	RequirePreauthForCreds                    = "If `true`, requires preauthorization before a User can retrieve credentials to sign in."
	ResourceGroupID                           = "The UUID of a OktaPA Resource Group."
	Roles                                     = "A list of roles for the ASA Group. Options are 'access_user', 'access_admin', and 'reporting_user'."
	SSHCertificateType                        = fmt.Sprintf("The SSH certificate type used by access requests. Options include: [%s]. '%s' is a deprecated key algorithm type. "+
		"This option should only be used to connect to legacy systems that cannot use newer SSH versions. If you do need to use '%s', it is recommended to connect via a gateway with traffic forwarding. "+
		"Otherwise, please use a more current key algorithm. If left unspecified, '%s' is used by default.", typed_strings.SSHCertTypeListFormat(), typed_strings.CertTypeRsa, typed_strings.CertTypeRsa, typed_strings.CertTypeEd25519)
	SSHSessionRecording           = "If `true`, enables ssh recording on server access requests."
	SecretFolderID                = "The UUID of the Secret Folder."
	SecretFolderPath              = "The path of the Secret Folder"
	SecretID                      = "The UUID of the Secret."
	SecurityPolicyActive          = "If true, indicates that the Security Policy is active."
	SecurityPolicyLabelSelectors  = "Defines the label selectors used to target resources by the Security Policy."
	SecurityPolicyPrincipals      = "Defines the users bound to the Security Policy."
	SecurityPolicyResouceGroup    = "The UUID of a specific Resource Group associated with the Security Policy. If undefined, the Security Policy applies to all resources for your Team. This value must be defined if the current user has the Delegated Security Admin role."
	SecurityPolicyResources       = "Defines the resources targeted by the Security Policy."
	SecurityPolicyRule            = "Defines the privileges available to resources matched to the Security Policy."
	SecurityPolicySecret          = "Defines a specific Secret targeted by the Security Policy."
	SecurityPolicySecretFolder    = "Defines a specific Secret Folder targeted by the Security Policy."
	SecurityPolicySecrets         = "Defines the secret-based resources targeted by the Security Policy."
	SecurityPolicyServer          = "Defines a specific server targeted by the Security Policy."
	SecurityPolicyServerAccount   = "Defines a local server account targeted by the Security Policy."
	SecurityPolicyServerLabels    = "Defines a map of key-value pairs used to match servers."
	SecurityPolicyServers         = "Defines the server-based resources targeted by the Security Policy."
	ServerAccess                  = "If `true`, members of this Group have access to the servers enrolled in the Project."
	ServerAdmin                   = "If `true`, provides sudo permissions to members of this Group on servers enrolled in this Project."
	ServerID                      = "The UUID of the server."
	ServerUserName                = "The name of the corresponding Server User."
	ServersSelector               = "Enables access to servers with labels matching all selectors. Only available to Teams with the Early Access Policy Sync feature."
	ServiceAccountPassword        = "The password of the service account used to query the domain."
	ServiceAccountUsername        = "The username of a service account used to query the domain."
	Status                        = "The status of the ASA User. Valid statuses are 'ACTIVE', 'DISABLED', and 'DELETED'."
	TeamName                      = "The human-readable name of the ASA Team that owns the resource. Values are lower-case."
	TeamSettingsID                = "Team name is populated as the ID for team settings."
	Token                         = "The secret used for resource enrollment."
	UsePasswordless               = "if `true`, Users will not need password to login."
	UserName                      = "The human-readable name of the User."
	UserOnDemandPeriod            = "If defined, set time period in seconds that an on-demand user account exists on the server following an access request."
	UserProvisioningExactUserName = "If true, ASA has ASA Users configured through SCIM to maintain the exact username that is specified."
	UserType                      = "The user type. Valid types are 'human' and 'service'."
	WebSessionDuration            = "Defines the duration of the web session. Configure the web session to be between 30 minutes and 25 hours."
	CloudConnectionProvider       = "The cloud provider for the connection, example AWS"
	CloudConnectionDetails        = "More data about the cloud connection: AWS account id, external id and AWS role arn"
)
