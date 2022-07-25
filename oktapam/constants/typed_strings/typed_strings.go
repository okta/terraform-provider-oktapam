package typed_strings

type UserStatus string
type UserType string

const (
	UserStatusActive   UserStatus = "ACTIVE"
	UserStatusDisabled UserStatus = "DISABLED"
	UserStatusDeleted  UserStatus = "DELETED"

	UserTypeHuman   UserType = "human"
	UserTypeService UserType = "service"
	UserTypeMissing UserType = ""
)

func (ut UserType) String() string {
	return string(ut)
}

func (ut UserStatus) String() string {
	return string(ut)
}
