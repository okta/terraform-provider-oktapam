package descriptions

import (
	"fmt"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/typed_strings"
)

var (
	// Supported Resource
	ResourceGatewaySetupToken     = fmt.Sprintf("A token for ASA Gateway enrollment. %s", LinkGatewaySetupToken)
	ResourceGroup                 = fmt.Sprintf("A set of ASA Users. %s", LinkGroup)
	ResourceProject               = fmt.Sprintf("An ASA construct that contains servers and is used to grant end users access to server infrastructure. %s", LinkProject)
	ResourceProjectGroup          = fmt.Sprintf("Assigns an ASA Group to a Project and configures how that group is created on servers. %s", LinkGroup)
	ResourceServerEnrollmentToken = fmt.Sprintf("A token for enrolling servers to an ASA Project. %s", LinkServerEnrollmentToken)
	ResourceUser                  = fmt.Sprintf("An ASA User. Valid user types are `%s` and `%s`. %s", typed_strings.UserTypeHuman.String(), typed_strings.UserTypeService, LinkServiceUser)

	// EA/Beta Features:
	ResourceADConnection                = fmt.Sprintf("%s An Active Directory (AD) Connection to query AD Domain for available servers.", PrefixBeta)
	ResourceADRuleAssignment            = fmt.Sprintf("%s Assign an Active Directory (AD) Rule to an ASA Project.", PrefixBeta)
	ResourceADTaskSettings              = fmt.Sprintf("%s Settings for Active Directory (AD) task for server discovery", PrefixBeta)
	ResourceKubernetesCluster           = fmt.Sprintf("%s Represents a Kubernetes cluster that has been registered with ASA. This resource will create configuration items in ASA needed in order to configure a Kubernetes cluster to authenticate user access with ASA. NOTE: This resource is only available for customers that are participating in the ASA Kubernetes support Beta program; contact support@okta.com if you wish to participate in the beta.", PrefixBeta)
	ResourceKubernetesClusterConnection = fmt.Sprintf("%s A set of details describing how to connect to an existing Kubernetes Cluster. NOTE: This resource is only available for customers that are participating in the ASA Kubernetes support Beta program; contact support@okta.com if you wish to participate in the beta.", PrefixBeta)
	ResourceKubernetesClusterGroup      = fmt.Sprintf("%s A mapping of Kubernetes cluster to an ASA Group. Members of the provided ASA group will be granted access to clusters which match the specified cluster_selectors. NOTE: This resource is only available for customers that are participating in the ASA Kubernetes support Beta program; contact support@okta.com if you wish to participate in the beta.", PrefixBeta)
)
