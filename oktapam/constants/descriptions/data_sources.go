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
	SourceADConnections                      = "A list of AD Connections associated with your Team."
	SourceCurrentUser                        = "Returns the username of the user used to invoke the terraform."
	SourceDatabase                           = "Returns an existing Database."
	SourceDatabasePasswordSettings           = "Returns an existing Database Password Policy for a PAM Project."
	SourceGateways                           = fmt.Sprintf("Returns a list of all Gateways connected to the Team specified in the OKTAPAM_TEAM environment variable. %s", LinkGateway)
	SourceGroup                              = fmt.Sprintf("Returns an existing Group. %s", LinkGroup)
	SourceProject                            = fmt.Sprintf("Returns an existing Project. %s", LinkProject)
	SourceProjectGroup                       = fmt.Sprintf("Returns an existing Group assigned to a specified Project. %s", LinkProjectGroup)
	SourceResourceGroup                      = fmt.Sprintf("Returns an existing PAM Resource Group. %s", LinkResourceGroup)
	SourceResourceGroupProject               = fmt.Sprintf("Returns an existing PAM Project associated with a specific PAM Resource Group. %s", LinkResourceGroupProject)
	SourceResourceGroupServerEnrollmentToken = fmt.Sprintf("Returns an existing PAM Server Enrollment Token associated with a specific PAM project. %s", LinkResourceGroupServerEnrollmentToken)
	SourceSecurityPolicy                     = fmt.Sprintf("Returns an existing PAM Security Policy. %s", LinkSecurityPolicy)
	SourceServerPasswordSettings             = fmt.Sprintf("Returns an existing Server Password Policy for a PAM Project.  %s", LinkPasswordSettings)
	SourceTeamSettings                       = "Returns Team-level settings, such as authentication and enrollment details, for your Team."
	SourceADUserSyncTaskSettings             = fmt.Sprintf("Returns an existing AD user sync job. %s", LinkADUserDiscovery)
	SourceCloudConnection                    = fmt.Sprintf("Returns an existing PAM Cloud Connection. %s", LinkCloudConnection)

	// List
	SourceADUserSyncTaskSettingsIDList        = fmt.Sprintf("Returns a list of all AD user sync jobs a specified AD Connection. %s", LinkADUserDiscovery)
	SourceDatabases                           = "Returns a list of Databases, constrained by the given parameters."
	SourceGroups                              = fmt.Sprintf("Returns a list of all Groups associated with the Team specified by the OKTAPAM_TEAM environment variable. %s", LinkGroup)
	SourceProjects                            = fmt.Sprintf("Returns a list of all Projects associated with the Team specified by the OKTAPAM_TEAM environment variable. %s", LinkProject)
	SourceProjectGroups                       = fmt.Sprintf("Returns a list of Project Groups associated with a Project. %s", LinkProjectGroup)
	SourceResourceGroups                      = fmt.Sprintf("Returns a list of Resource Groups associated with the Team specified by the OKTAPAM_TEAM environment variable. %s", LinkResourceGroup)
	SourceResourceGroupProjects               = fmt.Sprintf("Returns a list of Projects associated with an existing Resource Group. %s", LinkResourceGroupProject)
	SourceResourceGroupServerEnrollmentTokens = fmt.Sprintf("Returns a list of Server Enrollment Tokens associated with a specific Project. %s", LinkResourceGroupServerEnrollmentToken)
	SourceSecretFolders                       = "Returns a list of Secret Folders, constrained by the given parameters."
	SourceSecurityPolicies                    = fmt.Sprintf("Returns a list of Security Policies associated with the Team specified by the OKTAPAM_TEAM environment variable. %s", LinkSecurityPolicy)
	SourceCloudConnections                    = fmt.Sprintf("Returns a list of Cloud Connections. %s", LinkCloudConnection)
)
