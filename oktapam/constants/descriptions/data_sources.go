package descriptions

import "fmt"

var (
	// Supported Data Sources

	// Fetch
	SourceGatewaySetupToken     = fmt.Sprintf("Returns a previously created ASA Gateway Setup Token associated with the ASA Team specified in the OKTAPAM_TEAM environment variable. %s", LinkGatewaySetupToken)
	SourceServerEnrollmentToken = fmt.Sprintf("Returns a previously created ASA Server Enrollment Token associated with a specific ASA Project. %s", LinkServerEnrollmentToken)

	// List
	SourceGatewaySetupTokens     = fmt.Sprintf("Returns a list of ASA Gateway Setup Token IDs associated with the ASA Team specified in the OKTAPAM_TEAM environment variable. %s", LinkGatewaySetupToken)
	SourceServerEnrollmentTokens = fmt.Sprintf("A list of all ASA Server Enrollment Tokens associated with a specific ASA Project. %s", LinkServerEnrollmentToken)

	// Features In Progress

	// Fetch
	SourceADConnections                      = "A list of ASA AD Connections associated with an ASA Team."
	SourceGateways                           = fmt.Sprintf("Returns a list of all ASA Gateways connected to the ASA Team specified in the OKTAPAM_TEAM environment variable. %s", LinkGateway)
	SourceGroup                              = fmt.Sprintf("Returns a previously created ASA Group. %s", LinkGroup)
	SourcePasswordSettings                   = fmt.Sprintf("Returns a previously configured password settings for a PAM project.  %s", LinkPasswordSettings)
	SourceProject                            = fmt.Sprintf("Returns a previously created ASA Project. %s", LinkProject)
	SourceProjectGroup                       = fmt.Sprintf("Returns a previously created ASA Group assigned to a given ASA Project. %s", LinkProjectGroup)
	SourceResourceGroup                      = fmt.Sprintf("Returns a previously created PAM resource group. %s", LinkResourceGroup)
	SourceResourceGroupProject               = fmt.Sprintf("Returns a previously created PAM project associated with a specific PAM resource group. %s", LinkResourceGroupProject)
	SourceResourceGroupServerEnrollmentToken = fmt.Sprintf("Returns a previously created PAM server enrollment token associated with a specific PAM project. %s", LinkResourceGroupServerEnrollmentToken)
	SourceSecurityPolicy                     = fmt.Sprintf("Returns a previously created PAM security policy. %s", LinkSecurityPolicy)
	SourceTeamSettings                       = "Returns Team-level settings for a specific Team, such as authentication and enrollment details"
	SourceADUserSyncTaskSettings             = fmt.Sprintf("Returns a previously created ASA AD User Sync Task Settings. %s", LinkADUserDiscovery)

	// List
	SourceADUserSyncTaskSettingsIDList        = fmt.Sprintf("Returns a list of previously created ASA AD User Sync Task Settings IDs for an AD connection. %s", LinkADUserDiscovery)
	SourceGroups                              = fmt.Sprintf("Returns a list of all ASA Groups associated with the ASA Team specified in the OKTAPAM_TEAM environment variable. %s", LinkGroup)
	SourceProjects                            = fmt.Sprintf("Returns a list of ASA Projects associated with the ASA Team specified in the OKTAPAM_TEAM environment variable. %s", LinkProject)
	SourceProjectGroups                       = fmt.Sprintf("A list of ASA Project Groups associated with an ASA Project. %s", LinkProjectGroup)
	SourceResourceGroups                      = fmt.Sprintf("A list of PAM resource groups associated with the PAM team specified in the OKTAPAM_TEAM environment variable. %s", LinkResourceGroup)
	SourceResourceGroupProjects               = fmt.Sprintf("A list of PAM resource group projects associated with a PAM resource group. %s", LinkResourceGroupProject)
	SourceResourceGroupServerEnrollmentTokens = fmt.Sprintf("A list of PAM server enrollment tokens associated with a specific PAM project. %s", LinkResourceGroupServerEnrollmentToken)
	SourceSecurityPolicies                    = fmt.Sprintf("A list of PAM security policies associated with the PAM team specified in the OKTAPAM_TEAM environment variable. %s", LinkSecurityPolicy)
)
