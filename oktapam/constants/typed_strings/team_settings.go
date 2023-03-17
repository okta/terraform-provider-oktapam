package typed_strings

type IncludeUserSID string

const (
	IncludeUserSIDNever IncludeUserSID = "Never"
	IncludeUserSIDIfAvailable IncludeUserSID = "If_available"
	IncludeUserSIDAlways     IncludeUserSID = "Always"
)

func (includeUserSID IncludeUserSID) String() string {
	return string(includeUserSID)
}

