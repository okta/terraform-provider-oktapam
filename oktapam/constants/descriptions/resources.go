package descriptions

import (
	"fmt"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/typed_strings"
)

var (
	// Supported Resource
	ResourceADCertificateObject      = fmt.Sprintf("Upload the Enterprise Signed certificate for use with an AD Connection. %s", LinkADCertificates)
	ResourceADCertificateRequest     = fmt.Sprintf("A Certificate Request to generate an AD Passwordless Self Signed Certificate or Certificate Signing Request(CSR). %s", LinkADCertificates)
	ResourceADConnection             = fmt.Sprintf("An Active Directory (AD) Connection used to query an existing AD Domain for available resources. %s", LinkADConnection)
	ResourceADRuleAssignment         = fmt.Sprintf("Assign an Active Directory (AD) Rule to a Project. %s", LinkADRuleAssignment)
	ResourceADServerSyncTaskSettings = fmt.Sprintf("Settings for an Active Directory (AD) task that discovers servers. %s", LinkADServerDiscovery)
	ResourceADUserSyncTaskSettings   = fmt.Sprintf("Settings for an Active Directory (AD) task that discovers user SIDs. %s", LinkADUserDiscovery)
	ResourceGatewaySetupToken        = fmt.Sprintf("A token used to enroll a Gateway. %s", LinkGatewaySetupToken)
	ResourceGroup                    = fmt.Sprintf("A set of ASA Users. %s", LinkGroup)
	ResourceProject                  = fmt.Sprintf("A construct that contains servers and is used to grant user access to server infrastructure. %s", LinkProject)
	ResourceProjectGroup             = fmt.Sprintf("Assigns a  Group to a Project and configures how the Group is created on enrolled servers. %s", LinkGroup)
	ResourceServerEnrollmentToken    = fmt.Sprintf("A token used to enroll servers in a Project. %s", LinkServerEnrollmentToken)
	ResourceTeamSettings             = fmt.Sprintf("Team-level settings for a specific Team, such as authentication and enrollment details. %s", LinkTeamSettings)
	ResourceUser                     = fmt.Sprintf("An ASA User. Valid user types are `%s` and `%s`. %s", typed_strings.UserTypeHuman, typed_strings.UserTypeService, LinkServiceUser)

	// EA/Beta Features:
	ResourceKubernetesCluster           = fmt.Sprintf("%s Represents a Kubernetes cluster that has been registered with your Team. This resource creates configuration items used to define how the cluster can authenticate user access with ASA. %s", PrefixBeta, WarningBetaK8s)
	ResourceKubernetesClusterConnection = fmt.Sprintf("%s A set of details describing how to connect to an existing Kubernetes cluster. %s", PrefixBeta, WarningBetaK8s)
	ResourceKubernetesClusterGroup      = fmt.Sprintf("%s A mapping of Kubernetes cluster to a Group. Members of the specified Group are granted access to clusters which match the specified cluster_selectors. %s", PrefixBeta, WarningBetaK8s)

	// PAM LEA Features
	ResourcePasswordSettings                   = fmt.Sprintf("%s The settings for passwords set on resources within the project. %s", PrefixLEA, LinkPasswordSettings)
	ResourceResourceGroup                      = fmt.Sprintf("%s A PAM construct that contains a group of projects. %s", PrefixLEA, LinkResourceGroup)
	ResourceResourceGroupProject               = fmt.Sprintf("%s A PAM construct that contains a collection of resources that share settings. %s", PrefixLEA, LinkResourceGroupProject)
	ResourceResourceGroupServerEnrollmentToken = fmt.Sprintf("%s A token used to enroll servers in a PAM Project. %s", PrefixLEA, LinkResourceGroupServerEnrollmentToken)
	ResourceSecurityPolicy                     = fmt.Sprintf("%s A policy which defines how users can gain access to resources. %s", PrefixLEA, LinkSecurityPolicy)
)
