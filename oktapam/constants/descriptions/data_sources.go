package descriptions

const (
	// Supported Data Sources

	// Fetch
	SourceGatewaySetupToken     = "Returns a previously created ASA Gateway Setup Token. " + LinkGatewaySetupToken
	SourceServerEnrollmentToken = "Returns a previously created ASA Server Enrollment Token, corresponding to an ASA Project. " + LinkServerEnrollmentToken

	// List
	SourceGatewaySetupTokens     = "Returns a list of ASA Gateway Setup Token IDs, corresponding to an ASA Team. " + LinkGatewaySetupToken
	SourceServerEnrollmentTokens = "A list of tokens for ASA Server enrollment, corresponding to an ASA Project. " + LinkServerEnrollmentToken

	// Features In Progress

	// Fetch
	SourceGroup        = "Returns a previously created ASA Group. " + LinkGroup
	SourceProject      = "Returns a previously created ASA Project. " + LinkProject
	SourceProjectGroup = "Returns a previously created ASA Group assigned to a given ASA Project. " + LinkProjectGroup

	// List
	SourceGroups        = "Returns a list of ASA Groups. " + LinkGroup
	SourceProjects      = "Returns a list of ASA Projects. " + LinkProject
	SourceProjectGroups = "Returns a list of ASA Project Groups, corespondibng to an ASA Project. " + LinkProjectGroup

	// EA/Beta Features

	// List
	SourceADConnections = "Returns a list of ASA AD Connections, corresponding to an ASA Team."
	SourceGateways      = "Returns a list of ASA Gateways, corresponding to an ASA Team." + LinkGateway
)
