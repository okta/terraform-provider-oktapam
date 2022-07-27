package descriptions

const (
	// Query Parameter Descriptions
	FilterContains               = "If a value is provided, the results are filtered to only contain resources whose name contains that value."
	FilterCreateServerGroup      = "If 'true', the results only include the ASA Project Groups that have 'create_server_group' field set to 'true'."
	FilterDescriptionContains    = "If a value is provided, the results are filtered to only contain resources whose name contains that value."
	FilterDisconnectedModeOnOnly = "If 'true', the results only include resources with disconnected mode enabled. NOTE: This field is only valid for ASA teams with the Disconnected Mode Beta feature enabled."
	FilterHasNoSelectors         = "If 'true', the results only include resources with empty label selectors. NOTE: Label selectors are only applicable if the ASA Team has the Early Access PolicySync feature enabled."
	FilterHasSelectors           = "If 'true', the results only include resources with label selectors set. NOTE: Label selectors are only applicable if the ASA Team has the Early Access PolicySync feature enabled."
	FilterIncludeDeleted         = "If 'true', the results include deleted resources."
	FilterOnlyIncludeDeleted     = "If 'true', the results only include deleted resources."
	FilterIncludeRemoved         = "If 'true', the results include removed resources."
	FilterOfflineEnabled         = "If 'true', the results only include resources with disconnected mode enabled." // NOTE: This is inconsistent, most other API endpoints use `disconnected_mode_on_only`.
	FilterProjectName            = "If a value is provided, the results are filtered to only contain resources belonging to the ASA Project."
	FilterSelf                   = "If 'true', only lists the ASA Projects that the ASA User making this request has been assigned."
	FilterGatewayID              = "If 'true', the results only include the connections which has the same gateway id."
	FilterCertificateID          = "If 'true', the results only include the connections which has the same certificate id."
	FilterIncludeCertDetails     = "If 'true', the results include the certificate details "
	FilterStartsWith             = "If a value is provided, includes ASA Users with name that begins with the value."
	FilterStatus                 = "If a value is provided, includes ASA Users with specified statuses. Valid statuses include: `ACTIVE`, `DISABLED`, and `DELETED`."

	// Client-Side Filter
	FilterUserType = "Valid types are `human` and `service`. If left unspecified, both types will be included."
)
