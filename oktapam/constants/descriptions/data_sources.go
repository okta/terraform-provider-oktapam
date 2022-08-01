package descriptions

import "fmt"

var (
	// Supported Data Sources

	// Fetch
	SourceGatewaySetupToken     = fmt.Sprintf("Returns a previously created ASA Gateway Setup Token. %s", LinkGatewaySetupToken)
	SourceServerEnrollmentToken = fmt.Sprintf("Returns a previously created ASA Server Enrollment Token, corresponding to an ASA Project. %s", LinkServerEnrollmentToken)

	// List
	SourceGatewaySetupTokens     = fmt.Sprintf("Returns a list of ASA Gateway Setup Token IDs, corresponding to an ASA Team. %s", LinkGatewaySetupToken)
	SourceServerEnrollmentTokens = fmt.Sprintf("A list of tokens for ASA Server enrollment, corresponding to an ASA Project. %s", LinkServerEnrollmentToken)

	// Features In Progress

	// Fetch
	SourceADConnections = "A list of ASA AD Connections associated with an ASA Team."
	SourceGateways      = fmt.Sprintf("Returns a list of all ASA Gateways connected to a specific ASA Team. %s", LinkGateway)
	SourceGroup         = fmt.Sprintf("Returns a previously created ASA Group. %s", LinkGroup)
	SourceProject       = fmt.Sprintf("Returns a previously created ASA Project. %s", LinkProject)
	SourceProjectGroup  = fmt.Sprintf("Returns a previously created ASA Group assigned to a given ASA Project. %s", LinkProjectGroup)

	// List
	SourceGroups        = fmt.Sprintf("Returns a list of ASA Groups. %s", LinkGroup)
	SourceProjects      = fmt.Sprintf("Returns a list of ASA Projects. %s", LinkProject)
	SourceProjectGroups = fmt.Sprintf("Returns a list of ASA Project Groups, corespondibng to an ASA Project. %s", LinkProjectGroup)
)
