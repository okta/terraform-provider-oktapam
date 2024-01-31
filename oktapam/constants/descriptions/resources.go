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
	ResourceGroup                    = fmt.Sprintf("A set of ASA/PAM Users. %s", LinkGroup)
	ResourceProject                  = fmt.Sprintf("A construct that contains servers and is used to grant user access to server infrastructure. %s", LinkProject)
	ResourceProjectGroup             = fmt.Sprintf("Assigns a  Group to a Project and configures how the Group is created on enrolled servers. %s", LinkGroup)
	ResourceServerEnrollmentToken    = fmt.Sprintf("A token used to enroll servers in a Project. %s", LinkServerEnrollmentToken)
	ResourceTeamSettings             = fmt.Sprintf("Team-level settings for a specific Team, such as authentication and enrollment details. %s", LinkTeamSettings)
	ResourceUser                     = fmt.Sprintf("An ASA/PAM User. Valid user types are `%s` and `%s`. %s", typed_strings.UserTypeHuman, typed_strings.UserTypeService, LinkServiceUser)
	ResourceUserGroupAttachment      = "Manages a single assignment of an ASA/PAM User to a Group."

	// EA/Beta Features:
	ResourceKubernetesCluster                  = fmt.Sprintf("%s Represents a Kubernetes cluster that has been registered with your Team. This resource creates configuration items used to define how the cluster can authenticate user access with ASA. %s", PrefixBeta, WarningBetaK8s)
	ResourceKubernetesClusterConnection        = fmt.Sprintf("%s A set of details describing how to connect to an existing Kubernetes cluster. %s", PrefixBeta, WarningBetaK8s)
	ResourceKubernetesClusterGroup             = fmt.Sprintf("%s A mapping of Kubernetes cluster to a Group. Members of the specified Group are granted access to clusters which match the specified cluster_selectors. %s", PrefixBeta, WarningBetaK8s)
	ResourceDatabase                           = "Represents a datastore that has been registered with your Team."
	ResourceDatabasePasswordSettings           = "The settings for passwords set on databases within the project."
	ResourceResourceGroup                      = fmt.Sprintf("A PAM construct that contains a group of projects. %s", LinkResourceGroup)
	ResourceResourceGroupProject               = fmt.Sprintf("A PAM construct that contains a collection of resources that share settings. %s", LinkResourceGroupProject)
	ResourceResourceGroupServerEnrollmentToken = fmt.Sprintf("A token used to enroll servers in a PAM Project. %s", LinkResourceGroupServerEnrollmentToken)
	ResourceSecretFolder                       = fmt.Sprintf("A file-system like construct that contains secrets or nested secret folders. %s", LinkSecretFolder)
	ResourceServerPasswordSettings             = fmt.Sprintf("The settings for passwords set on servers within the project. %s", LinkPasswordSettings)
	ResourceSecurityPolicy                     = fmt.Sprintf("A policy which defines how users can gain access to resources. %s", LinkSecurityPolicy)
)
