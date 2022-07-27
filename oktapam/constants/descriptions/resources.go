package descriptions

const (
	// Supported Resource
	ResourceGatewaySetupToken     = "A token for ASA Gateway enrollment. " + LinkGatewaySetupToken
	ResourceGroup                 = "A set of ASA Users. " + LinkGroup
	ResourceProject               = "An ASA construct that contains servers and is used to grant end users access to server infrastructure. " + LinkProject
	ResourceProjectGroup          = "Assigns an ASA Group to a Project and configures how that group is created on servers. " + LinkGroup
	ResourceServerEnrollmentToken = "A token for enrolling servers to an ASA Project. " + LinkServerEnrollmentToken
	ResourceUser                  = "An ASA User. Valid user types are `human` and `service`. " + LinkServiceUser

	// EA/Beta Features:
	ResourceADConnection                = PrefixBeta + "An Active Directory (AD) Connection to query AD Domain for available servers."
	ResourceADRuleAssignment            = PrefixBeta + ""
	ResourceADTaskSettings              = PrefixBeta + "Settings for Active Directory (AD) task for server discovery"
	ResourceKubernetesCluster           = PrefixBeta + "Represents a Kubernetes cluster that has been registered with ASA. This resource will create configuration items in ASA needed in order to configure a Kubernetes cluster to authenticate user access with ASA. NOTE: This resource is only available for customers that are participating in the ASA Kubernetes support Beta program; contact support@okta.com if you wish to participate in the beta."
	ResourceKubernetesClusterConnection = PrefixBeta + "A set of details describing how to connect to an existing Kubernetes Cluster."
	ResourceKubernetesClusterGroup      = PrefixBeta + "A mapping of Kubernetes cluster to an ASA Group. Members of the provided ASA group will be granted access to clusters which match the specified cluster_selectors."
)
