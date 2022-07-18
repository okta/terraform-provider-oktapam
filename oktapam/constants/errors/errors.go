package errors

import (
	"fmt"

	"github.com/okta/terraform-provider-oktapam/oktapam/client"
)

var (
	InvalidUserTypeError = "%s" + fmt.Sprintf(" is not a valid user type option. Either %q or %q must be specified.", client.UserTypeHuman, client.UserTypeService)
	MissingUserTypeError = fmt.Sprintf("User type must be specified. Valid options are %q or %q.", client.UserTypeHuman, client.UserTypeService)
)
