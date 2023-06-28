package descriptions

import (
	"fmt"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/typed_strings"
)

var (
	// Supported Resource
	ResourceADCertificateObject    = fmt.Sprintf("Upload the Enterprise Signed certificate for use with an AD Connection. %s", LinkADCertificates)
	ResourceADCertificateRequest   = fmt.Sprintf("A Certificate Request to generate an AD Passwordless Self Signed Certificate or Certificate Signing Request(CSR). %s", LinkADCertificates)
	ResourceADConnection           = fmt.Sprintf("An Active Directory (AD) Connection to query AD Domain for available servers. %s", LinkADConnection)
	ResourceADRuleAssignment       = fmt.Sprintf("Assign an Active Directory (AD) Rule to an ASA Project. %s", LinkADRuleAssignment)
	ResourceADTaskSettings         = fmt.Sprintf("Settings for Active Directory (AD) task for server discovery. %s", LinkADServerDiscovery)
	ResourceADUserSyncTaskSettings = fmt.Sprintf("%s Settings for Active Directory (AD) task for user SID discovery. %s", PrefixDev, LinkADUserDiscovery)
	ResourceGatewaySetupToken      = fmt.Sprintf("A token for ASA Gateway enrollment. %s", LinkGatewaySetupToken)
	ResourceGroup                  = fmt.Sprintf("A set of ASA Users. %s", LinkGroup)
	ResourceProject                = fmt.Sprintf("An ASA construct that contains servers and is used to grant end users access to server infrastructure. %s", LinkProject)
	ResourceProjectGroup           = fmt.Sprintf("Assigns an ASA Group to a Project and configures how that group is created on servers. %s", LinkGroup)
	ResourceServerEnrollmentToken  = fmt.Sprintf("A token used to enroll servers in an ASA Project. %s", LinkServerEnrollmentToken)
	ResourceTeamSettings           = fmt.Sprintf("Team-level settings for a specific Team, such as authentication and enrollment details. %s", LinkTeamSettings)
	ResourceUser                   = fmt.Sprintf("An ASA User. Valid user types are `%s` and `%s`. %s", typed_strings.UserTypeHuman, typed_strings.UserTypeService, LinkServiceUser)

	// EA/Beta Features:
	ResourceKubernetesCluster           = fmt.Sprintf("%s Represents a Kubernetes cluster that has been registered with your Team. This resource creates configuration items used to define how the cluster can authenticate user access with ASA. %s", PrefixBeta, WarningBetaK8s)
	ResourceKubernetesClusterConnection = fmt.Sprintf("%s A set of details describing how to connect to an existing Kubernetes Cluster. %s", PrefixBeta, WarningBetaK8s)
	ResourceKubernetesClusterGroup      = fmt.Sprintf("%s A mapping of Kubernetes cluster to a Group. Members of the specified Group are granted access to clusters which match the specified cluster_selectors. %s", PrefixBeta, WarningBetaK8s)

	// PAM LEA Features
	ResourcePasswordSettings                   = fmt.Sprintf("%s The settings for passwords set on resources within the project. %s", PrefixLEA, LinkPasswordSettings)
	ResourceResourceGroup                      = fmt.Sprintf("%s A PAM construct that contains a group of projects. %s", PrefixLEA, LinkResourceGroup)
	ResourceResourceGroupProject               = fmt.Sprintf("%s A PAM construct that contains a collection of resources that share settings. %s", PrefixLEA, LinkResourceGroupProject)
	ResourceResourceGroupServerEnrollmentToken = fmt.Sprintf("%s A token used to enroll servers in a PAM Project. %s", PrefixLEA, LinkResourceGroupServerEnrollmentToken)
	ResourceSecurityPolicy                     = fmt.Sprintf("%s A policy which defines how users can gain access to resources. %s", PrefixLEA, LinkSecurityPolicy)
)
