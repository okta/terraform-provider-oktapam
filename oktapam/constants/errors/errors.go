package errors

import (
	"fmt"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/typed_strings"
)

var (
	HumanUserCreationError = fmt.Sprintf("%s user creation is not available. Please create the user from your Okta console.", typed_strings.UserTypeHuman)
	HumanUserDeletionError = fmt.Sprintf("%s user deletion is not available. Please delete the user from your Okta console.", typed_strings.UserTypeHuman)
	InvalidUserTypeError   = "%s" + fmt.Sprintf(" is not a valid user type option. Either %q or %q must be specified.", typed_strings.UserTypeHuman, typed_strings.UserTypeService)
	MissingUserTypeError   = fmt.Sprintf("User type must be specified. Valid options are %q or %q.", typed_strings.UserTypeHuman, typed_strings.UserTypeService)
	MissingAttributeError  = "%s cannot be blank"
)
