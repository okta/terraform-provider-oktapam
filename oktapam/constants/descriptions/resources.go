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
	ResourceADConnection                = "An Active Directory (AD) Connection to query AD Domain for available servers."
	ResourceADRuleAssignment            = ""
	ResourceADTaskSettings              = "Settings for Active Directory (AD) task for server discovery"
	ResourceKubernetesCluster           = "A Kubernetes cluster."
	ResourceKubernetesClusterConnection = "A set of details describing how to connect to an existing Kubernetes Cluster."
	ResourceKubernetesClusterGroup      = "A mapping of Kubernetes cluster to Project Group."
)
