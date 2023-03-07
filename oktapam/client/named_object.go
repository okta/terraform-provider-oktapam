package client

type NamedObjectType string

const (
	UserNamedObjectType      NamedObjectType = "user"
	UserGroupNamedObjectType NamedObjectType = "user_group"
	ServerNamedObjectType    NamedObjectType = "server"
)

type NamedObject struct {
	Id   *string         `json:"id"`
	Name *string         `json:"name,omitempty"`
	Type NamedObjectType `json:"type,omitempty"`
}
