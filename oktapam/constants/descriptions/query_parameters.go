package descriptions

import "fmt"

var (
	// Query Parameter Descriptions
	FilterContains               = "If a value is provided, the results are filtered to only contain resources that contain the value in the name field."
	FilterCreateServerGroup      = "If `true`, results only include ASA Project Groups that have 'create_server_group' option set to `true`."
	FilterDescriptionContains    = "If a value is provided, the results are filtered to only contain resources whose name contains that value."
	FilterDisconnectedModeOnOnly = fmt.Sprintf("If `true`, the results only include resources with disconnected mode enabled. %s", WarningBetaDisconnectMode)
	FilterHasNoSelectors         = fmt.Sprintf("If `true`, the results only include resources with empty label selectors. %s", WarningEarlyAccessPolicySync)
	FilterHasSelectors           = fmt.Sprintf("If `true`, the results only include resources with label selectors set. %s", WarningEarlyAccessPolicySync)
	FilterIncludeDeleted         = "If `true`, results also include deleted resources."
	FilterOnlyIncludeDeleted     = "If `true`, results only include deleted resources."
	FilterIncludeRemoved         = "If `true`, the results also include resources that were previously removed."
	FilterProjectName            = "The name of an ASA Project. Results are returned for only this Project. Values are case-sensitive."
	FilterSelf                   = "If `true`, only lists the ASA Projects that the ASA User making this request has been assigned."
	FilterGatewayID              = "If `true`, only connections with a matching gateway ID are returned."
	FilterCertificateID          = "If `true`, only connections with a matching certificate ID are returned."
	FilterIncludeCertDetails     = "If `true`, results also include certificate details "
	FilterStartsWith             = "If a value is provided, includes ASA Users with name that begins with the value."
	FilterStatus                 = "If a value is provided, includes ASA Users with specified statuses. Valid statuses include: `ACTIVE`, `DISABLED`, and `DELETED`."
	FilterConnectionID           = "The id of an ASA AD Connection. Results are returned for only this Connection."
	FilterADUserSyncTaskSettingsStatus   = "If a value is provided, includes ASA user sync task settings with specified status. Valid statuses are `ACTIVE` and `INACTIVE`."

	// Client-Side Filter
	FilterUserType = "Valid types are `human` and `service`. If left unspecified, both types will be included."
)
