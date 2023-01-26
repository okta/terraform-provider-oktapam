package descriptions

import (
	"fmt"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/typed_strings"
)

var (
	// Supported Resource
	ResourceADCertificateObject   = fmt.Sprintf("Upload the Enterprise Signed certificate to ASA. %s", LinkADCertificates)
	ResourceADCertificateRequest  = fmt.Sprintf("A Certificate Request to generate AD Passwordless Self Signed Certificate or Certificate Signing Request(CSR). %s", LinkADCertificates)
	ResourceADConnection          = fmt.Sprintf("An Active Directory (AD) Connection to query AD Domain for available servers. %s", LinkADConnection)
	ResourceADRuleAssignment      = fmt.Sprintf("Assign an Active Directory (AD) Rule to an ASA Project. %s", LinkADRuleAssignment)
	ResourceADTaskSettings        = fmt.Sprintf("Settings for Active Directory (AD) task for server discovery. %s", LinkADServerDiscovery)
	ResourceGatewaySetupToken     = fmt.Sprintf("A token for ASA Gateway enrollment. %s", LinkGatewaySetupToken)
	ResourceGroup                 = fmt.Sprintf("A set of ASA Users. %s", LinkGroup)
	ResourceProject               = fmt.Sprintf("An ASA construct that contains servers and is used to grant end users access to server infrastructure. %s", LinkProject)
	ResourceProjectGroup          = fmt.Sprintf("Assigns an ASA Group to a Project and configures how that group is created on servers. %s", LinkGroup)
	ResourceServerEnrollmentToken = fmt.Sprintf("A token used to enroll servers in an ASA Project. %s", LinkServerEnrollmentToken)
	ResourceTeamSettings          = fmt.Sprintf("Team-level settings for a specific Team, such as authentication and enrollment details. %s", LinkTeamSettings)
	ResourceUser                  = fmt.Sprintf("An ASA User. Valid user types are `%s` and `%s`. %s", typed_strings.UserTypeHuman, typed_strings.UserTypeService, LinkServiceUser)

	// EA/Beta Features:
	ResourceKubernetesCluster           = fmt.Sprintf("%s Represents a Kubernetes cluster that has been registered with ASA. This resource will create configuration items in ASA needed in order to configure a Kubernetes cluster to authenticate user access with ASA. %s", PrefixBeta, WarningBetaK8s)
	ResourceKubernetesClusterConnection = fmt.Sprintf("%s A set of details describing how to connect to an existing Kubernetes Cluster. %s", PrefixBeta, WarningBetaK8s)
	ResourceKubernetesClusterGroup      = fmt.Sprintf("%s A mapping of Kubernetes cluster to an ASA Group. Members of the provided ASA group will be granted access to clusters which match the specified cluster_selectors. %s", PrefixBeta, WarningBetaK8s)
)
