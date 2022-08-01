package descriptions

import "fmt"

var (
	// Supported Data Sources

	// Fetch
	SourceGatewaySetupToken     = fmt.Sprintf("Returns a previously created ASA Gateway Setup Token associated with the ASA Team specified in the OKTAPAM_TEAM environment variable. %s", LinkGatewaySetupToken)
	SourceServerEnrollmentToken = fmt.Sprintf("Returns a previously created ASA Server Enrollment Token, corresponding to an ASA Project. %s", LinkServerEnrollmentToken)

	// List
	SourceGatewaySetupTokens     = fmt.Sprintf("Returns a list of ASA Gateway Setup Token IDs associated with the ASA Team specified in the OKTAPAM_TEAM environment variable. %s", LinkGatewaySetupToken)
	SourceServerEnrollmentTokens = fmt.Sprintf("A list of tokens for ASA Server enrollment, corresponding to an ASA Project. %s", LinkServerEnrollmentToken)

	// Features In Progress

	// Fetch
	SourceADConnections = "A list of ASA AD Connections associated with an ASA Team."
	SourceGateways      = fmt.Sprintf("Returns a list of all ASA Gateways connected to the ASA Team specified in the OKTAPAM_TEAM environment variable. %s", LinkGateway)
	SourceGroup         = fmt.Sprintf("Returns a previously created ASA Group. %s", LinkGroup)
	SourceProject       = fmt.Sprintf("Returns a previously created ASA Project. %s", LinkProject)
	SourceProjectGroup  = fmt.Sprintf("Returns a previously created ASA Group assigned to a given ASA Project. %s", LinkProjectGroup)

	// List
	SourceGroups        = fmt.Sprintf("Returns a list of all ASA Groups associated with the ASA Team specified in the OKTAPAM_TEAM environment variable. %s", LinkGroup)
	SourceProjects      = fmt.Sprintf("Returns a list of ASA Projects associated with the ASA Team specified in the OKTAPAM_TEAM environment variable. %s", LinkProject)
	SourceProjectGroups = fmt.Sprintf("Returns a list of ASA Project Groups, corespondibng to an ASA Project. %s", LinkProjectGroup)
)
