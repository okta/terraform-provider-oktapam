package errors

import (
	"fmt"

	"github.com/okta/terraform-provider-oktapam/oktapam/client"
)

var (
	InvalidUserTypeError = "%s" + fmt.Sprintf(" is not a valid user type option. Either `%s` or `%s` must be specified.", client.UserTypeHuman, client.UserTypeService)
	MissingUserTypeError = fmt.Sprintf("User type must be specified. Valid options are `%s` or `%s`.", client.UserTypeHuman, client.UserTypeService)
)
