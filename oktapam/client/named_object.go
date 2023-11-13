package client

type NamedObjectType string

const (
	UserNamedObjectType          NamedObjectType = "user"
	UserGroupNamedObjectType     NamedObjectType = "user_group"
	SecretNamedObjectType        NamedObjectType = "secret"
	SecretFolderNamedObjectType  NamedObjectType = "secret_folder"
	ServerNamedObjectType        NamedObjectType = "server"
	ResourceGroupNamedObjectType NamedObjectType = "resource_group"
)

type NamedObject struct {
	Id   *string         `json:"id"`
	Name *string         `json:"name,omitempty"`
	Type NamedObjectType `json:"type,omitempty"`
}
