package wrappers

import (
	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
)

type SecretWrapper struct {
	Secret          *pam.Secret
	ResourceGroupID string
	ProjectID       string
	ParentFolderID  string
	SecretContents  map[string]string
}

func (w SecretWrapper) ToResourceMap() map[string]any {
	m := make(map[string]any)

	m[attributes.ResourceGroup] = w.ResourceGroupID
	m[attributes.Project] = w.ProjectID
	m[attributes.ParentFolder] = w.ParentFolderID

	m[attributes.Name] = w.Secret.Name
	if w.Secret.Description.IsSet() {
		m[attributes.Description] = *w.Secret.Description.Get()
	}
	m[attributes.Secret] = w.SecretContents

	return m
}

func (w SecretWrapper) HasContents() bool {
	return len(w.SecretContents) != 0
}
