package descriptions

import "fmt"

var (
	// Query Parameter Descriptions
	FilterContains                     = "If defined, results only contain resources that include the specified value in the `name` field."
	FilterCreateServerGroup            = "If `true`, results only include Project Groups that have set the 'create_server_group' option to `true`."
	FilterDescriptionContains          = "If defined, results only contain resources that include the specified value in the `name` field."
	FilterDisconnectedModeOnOnly       = fmt.Sprintf("If `true`, results only include resources that have enabled Disconnected Mode. %s", WarningBetaDisconnectMode)
	FilterHasNoSelectors               = "If `true`, results only include resources haven't defined any label selectors."
	FilterHasSelectors                 = "If `true`, results only include resources have defined a label selector."
	FilterIncludeDeleted               = "If `true`, results also include deleted resources."
	FilterOnlyIncludeDeleted           = "If `true`, results only include deleted resources."
	FilterIncludeRemoved               = "If `true`, results also include resources that were previously removed."
	FilterProjectName                  = "If defined, results are only returned for the specified Project. Values are case-sensitive."
	FilterSelf                         = "If `true`, only lists the ASA Projects that the User making this request has been assigned."
	FilterGatewayID                    = "If `true`, results only include AD Connections associated with the specified Gateway ID."
	FilterCertificateID                = "If `true`, results only include AD Connections associated with the specified certificate ID."
	FilterIncludeCertDetails           = "If `true`, results also include certificate details "
	FilterStartsWith                   = "If defined, results only include Users with a name that begins with the specified value."
	FilterStatus                       = "If defined, results only include Users with specified statuses. Possible values: `ACTIVE`, `DISABLED`, and `DELETED`."
	FilterConnectionID                 = "If defined, results only include resources associated with the specified AD Connection."
	FilterADUserSyncTaskSettingsStatus = "If defined, results only include user sync jobs with the specified status. Possible values: `ACTIVE` and `INACTIVE`."
	FilterName                         = "The name of the resource"

	// Client-Side Filter
	FilterUserType = "If defined, results only include the specified type of User. Possible values: `human` and `service`."
)
