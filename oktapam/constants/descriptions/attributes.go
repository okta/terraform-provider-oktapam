package descriptions

import (
	"fmt"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/typed_strings"
)

var (
	// Attribute Descriptions
	AccessAddress                             = "Access Address of the gateway."
	AccessAddressAttribute                    = "AD Attribute mapped to IP address or DNS name used by the gateway to connect to a discovered server."
	AccountDiscovery                          = "If `true`, will enable the discovery of local accounts on servers within the project."
	ADConnectionID                            = "UUID of the AD Connection with which this AD Task Settings is associated."
	AdditionalAttributeMapping                = "Additional AD attributes mappings to Advanced Server Access labels."
	AdditionalAttributeMappingLabel           = "ASA label"
	AdditionalAttributeMappingValue           = "AD Attribute mapped to ASA label"
	AdditionalAttributeMappingIsGuid          = "Identifies an AD attribute as a Globally Unique Identifier (GUID)"
	ADRuleAssignments                         = "Assignment rules determine how servers are synced from Active Directory (AD) and assigned to projects."
	ADRuleAssignmentsBaseDN                   = "Specifies where the rule searches for servers."
	ADRuleAssignmentsLDAPQueryFilter          = "Specifies the specific criteria used to filter servers."
	ADRuleAssignmentsProjectID                = "Specifies a project to associate with matching servers"
	ADTaskFrequency                           = "Frequency of the AD Task"
	ADTaskIsActive                            = "If `true`, enables AD task"
	ADTaskRunTest                             = "If `true`, test is performed based on specified AD Task Settings"
	ADTaskStartHourUTC                        = "If AD task is scheduled to run daily, then specify start hour in UTC"
	ADUserSyncTaskSettingsBaseDN              = "Specifies where the rule searches for users."
	ADUserSyncTaskSettingsFrequency           = "Frequency of the AD User Sync Task"
	ADUserSyncTaskSettingsIsActive            = "If `true`, enables AD user sync task"
	ADUserSyncTaskSettingsLDAPQueryFilter     = "Specifies the criteria used to filter users."
	ADUserSyncTaskSettingsRunTest             = "If `true`, test is performed based on specified AD User Sync Task Settings"
	ADUserSyncTaskSettingsSIDField            = "AD attribute mapped to user security identifier"
	ADUserSyncTaskSettingsStartHourUTC        = "If AD user sync task is scheduled to run daily, then specify start hour in UTC"
	ADUserSyncTaskSettingsUPNField            = "AD attribute mapped to user principal name"
	AltNamesAttributes                        = "AD Attribute mapped to alternative hostnames or DNS entries used to resolve a discovered server."
	ApproveDeviceWithoutInteraction           = "If enabled, ASA auto-approves devices for ASA Users that are authenticated into this Team."
	BastionAttribute                          = "AD Attribute mapped to bastion host that Advanced Server Access clients can use to tunnel traffic to a discovered server."
	CertificateCommonName                     = "Common Name or FQDN to which certificate is issued to."
	CertificateContent                        = "Certificate Signing Request (CSR)/ Self Signed Certificate content."
	CertificateID                             = "Certificate ID used for password less access method."
	CertificateRequestType                    = "Specifies the type of certificate request - Certificate Signing Request (CSR)/ Self Signed Certificate."
	CertificateStatus                         = "Certificate status - Valid/Request Created."
	CharacterOptions                          = "The specific characters rules required by the password settings."
	CharacterOptionsDigits                    = "If `true`, passwords can include one or more numeric characters."
	CharacterOptionsLowerCase                 = "If `true`, passwords can include one or more lowercase characters."
	CharacterOptionsPunctuation               = "If `true`, passwords can include one or more punctuation/symbol characters."
	CharacterOptionsUpperCase                 = "If `true`, passwords can include one or more uppercase characters."
	ClientSessionDuration                     = "Defines the Client session duration. Value should be in hours between 1 hour and 25 hours."
	CloudProvider                             = "Cloud provider name of the host where gateway is running."
	ClusterGroupClaims                        = "A map of claims that will be included in credentials issued to users that are used to authenticate to Kubernetes clusters. Claims correspond to pre-configured role bindings on the cluster."
	ClusterSelector                           = "A label selector to used to match Kubernetes clusters."
	ConditionAccessRequest                    = "Identifies an existing Request Type in Access Requests."
	ConditionAccessRequestExpiresAfterSeconds = "The number of seconds the approval remains valid."
	ConditionAccessRequestRequestTypeID       = "The ID of the Access Request."
	ConditionAccessRequestRequestTypeName     = "The name of the Access Request."
	ConditionGateway                          = "Configures traffic settings for an existing PAM Gateway."
	ConditionSessionRecording                 = "Whether to record sessions made through an PAM Gateway."
	ConditionTrafficForwarding                = "Whether to forward traffic through an PAM Gateway."
	Conditions                                = "The conditions required before a privilege is made available to the Principals. All conditions must be met."
	CreateServerGroup                         = "If `true`, 'sftd' (ASA Server Agent) creates a corresponding local (unix or windows) group in the ASA Project's servers."
	CreateServerUsers                         = "If `true`, 'sftd' (ASA Server Agent) creates corresponding local (unix or windows) user accounts in the ASA Project's servers."
	CreatedAt                                 = "The UTC time when the resource was created. Format is '2022-01-01 00:00:00 +0000 UTC'."
	CreatedByUser                             = "The ASA User that created the resource."
	CSRDetails                                = "Certificate Signing Request (CSR) Details."
	DefaultAddress                            = "Default address of the gateway."
	DelegatedAdminGroups                      = "The ids of the groups for users who are designated as delegated admins for the resource group"
	DeletedAt                                 = "The UTC time of resource deletion. Format is '2022-01-01 00:00:00 +0000 UTC'."
	Description                               = "The human-readable description of the resource."
	DescriptionContains                       = "If a value is provided, the results are filtered to only contain resources whose description contains that value."
	Domain                                    = "The domain against which to query."
	DomainControllers                         = "A comma-separated list of the specific domain controller(s) that should be used to query the domain. Can be specified as a hostname or IP."
	EnablePeriodicRotation                    = "If `true`, requires passwords to be rotated after a period of time has passed. You must also set the periodic_rotation_duration_in_seconds param."
	EnterpriseSigned                          = "If `true`, certificate is signed by AD Certificate Services."
	ForwardTraffic                            = "If `true`, all traffic in the ASA Project be forwarded through selected ASA Gateways."
	GatewayID                                 = "The UUID of the Gateway with which this AD Connection is associated."
	GatewaySelector                           = "Assigns ASA Gateways with labels matching all selectors. At least one selector is required for traffic forwarding."
	GroupName                                 = "The human-readable name of the ASA Group. Values are case-sensitive."
	HostnameAttribute                         = "AD Attribute mapped to hostname used to identify a discovered server within Advanced Server Access."
	IncludeUserSID                            = fmt.Sprintf("%s Options for including the User Security Identifier (SID) in the certificate used for Passwordless authentication to a Windows server via RDP : Never, If_Available, and Always.", PrefixDev)
	IssuedAt                                  = "The UTC time when the token was issued. Format is '2022-01-01 00:00:00 +0000 UTC'."
	KubernetesAuthMechanism                   = "Mechanism to provide auth details to your Kubernetes cluster (eg. OIDC_RSA2048, NONE)"
	KubernetesAPIURL                          = "The URL used to access the control plane of the Kubernetes cluster."
	KubernetesClusterID                       = "The ASA ID of a Kubernetes cluster."
	KubernetesClusterKey                      = "The human-friendly key to associate with the Kubernetes cluster. Must be simple alphanumeric without spaces."
	KubernetesClusterLabels                   = "Map of labels to assign to the Kubernetes cluster."
	KubernetesPublicCertificate               = "The certificate expected when connecting to the Kubernetes cluster."
	Labels                                    = "A map of key-value pairings that define access to the ASA Gateway."
	LocalAccount                              = "The username of a local account on a server."
	LocalAccounts                             = "The usernames of a local accounts on a server."
	ManagedPrivilegedAccounts                 = "An array of managed accounts for password rotation."
	Name                                      = "The human-readable name of the resource. Values are case-sensitive."
	NextUnixGID                               = "The GID to use when creating a new ASA Server User. Default value starts at 63001."
	NextUnixUID                               = "The UID to use when creating a new ASA Server User. Default value starts at 60001."
	OIDCIssuerURL                             = "The OIDC Issuer URL to use when configuring your Kubernetes cluster. "
	OSAttribute                               = "AD Attribute mapped to server operating system of a discovered server."
	PasswordMinLength                         = "The minimum length allowed for the password."
	PasswordMaxLength                         = "The maximum length allowed for the password."
	PeriodicRotationDurationInSeconds         = "If periodic rotation is enabled, specifies how often passwords are rotated."
	PrincipalGroupIDs                         = "The UUIDs of PAM groups."
	PrivilegeEnabled                          = "If `true`, grant the privilege to Principals on matching resources."
	Privileges                                = "The specific privileges granted to Principals on matching resources."
	PrivilegePasswordCheckoutRDP              = "Defines the privilege to RDP to a server via a local account using a vaulted password."
	PrivilegePasswordCheckoutSSH              = "Defines the privilege to RDP to a server via a local account using a vaulted password."
	PrivilegePrincipalAccountRDP              = "Defines the privilege to RDP to a server via the user's principal account."
	PrivilegePrincipalAccountSSH              = "Defines the privilege to SSH to a server via the user's principal account."
	ProjectID                                 = "The UUID of PAM Project."
	ProjectName                               = "The human-readable name of the ASA Project. Values are case-sensitive."
	RDPSessionRecording                       = "If `true`, enable remote desktop protocol (RDP) recording on all servers in the ASA Project."
	ReactivateUsersViaIDP                     = "If a disabled or deleted ASA User is able to authenticate through the IdP, their ASA User is re-enabled."
	RefuseConnections                         = "If `true`, gateway refuse connection."
	RemovedAt                                 = "UTC time of resource removal from parent resource. Format is '2022-01-01 00:00:00 +0000 UTC'."
	RequirePreauthForCreds                    = "If `true`, require preauthorization before an ASA User can retrieve credentials to sign in."
	ResourceGroupID                           = "The UUID of PAM Resource Group."
	Roles                                     = "A list of roles for the ASA Group. Options are 'access_user', 'access_admin', and 'reporting_user'."
	SecurityPolicyActive                      = "If true, indicates that the Security Policy is active."
	SecurityPolicyLabelSelectors              = "Defines what resources that are targeted by the security policy via label selectors."
	SecurityPolicyPrincipals                  = "Defines what users are bound to the security policy."
	SecurityPolicyResources                   = "Defines the resources that are targeted by the security policy."
	SecurityPolicyRule                        = "Defines privileges available for matching resources."
	SecurityPolicyServers                     = "Defines the server-based resources that are targeted by the security policy rule."
	SecurityPolicyServer                      = "Define a server to match as a target within a security policy rule."
	SecurityPolicyServerAccount               = "Define a local account on a server to match as a target within a security policy rule."
	SecurityPolicyServerLabels                = "The map of key-value pairings that are used to match servers."
	ServerAccess                              = "If `true`, members of this ASA Group have access to the ASA Project servers."
	ServerAdmin                               = "If `true`, members of ASA Group have sudo permissions on ASA Project servers."
	ServerID                                  = "The UUID of the PAM server record."
	ServerUserName                            = "The name of the corresponding ASA Server User."
	ServersSelector                           = "Enables access to ASA Servers with labels matching all selectors. Only available to customers that have the Early Access Policy Sync feature enabled on their team."
	ServiceAccountUsername                    = "The username of the service account that can be used to query the domain."
	ServiceAccountPassword                    = "The password of the service account that can be used to query the domain."
	SSHCertificateType                        = fmt.Sprintf("The SSH certificate type used by access requests. Options include: [%s]. '%s' is a deprecated key algorithm type. "+
		"This option should only be used to connect to legacy systems that cannot use newer SSH versions. If you do need to use '%s', it is recommended to connect via a gateway with traffic forwarding. "+
		"Otherwise, please use a more current key algorithm. If left unspecified, '%s' is used by default.", typed_strings.SSHCertTypeListFormat(), typed_strings.CertTypeRsa, typed_strings.CertTypeRsa, typed_strings.CertTypeEd25519)
	SSHSessionRecording           = "If `true`, enables ssh recording on server access requests."
	Status                        = "The status of the ASA User. Valid statuses are 'ACTIVE', 'DISABLED', and 'DELETED'."
	TeamName                      = "The human-readable name of the ASA Team that owns the resource. Values are lower-case."
	TeamSettingsID                = "Team name is populated as the ID for team settings."
	Token                         = "The secret used for resource enrollment."
	UsePasswordless               = "if `true`, Users will not need password to login."
	UserOnDemandPeriod            = "If defined, set time period in seconds that an on-demand user account exists on the server following an access request."
	UserProvisioningExactUserName = "If true, ASA has ASA Users configured through SCIM to maintain the exact username that is specified."
	UserType                      = "The user type. Valid types are 'human' and 'service'."
	WebSessionDuration            = "Defines the duration of the web session. Configure the web session to be between 30 minutes and 25 hours."
)
