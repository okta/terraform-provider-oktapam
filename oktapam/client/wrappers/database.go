package wrappers

import (
	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
)

var (
	_ Wrappers = (*DatabaseResourceResponseWrapper)(nil)
	_ Wrappers = (*ManagementConnectionDetailsWrapper)(nil)
	_ Wrappers = (*MySQLBasicAuthManagementConnectionDetailsWrapper)(nil)
	_ Wrappers = (*MySQLBasicAuthDetailsWrapper)(nil)
)

type AttributeOverrides map[string]any

type Wrappers interface {
	ToResourceMap(overrides AttributeOverrides) map[string]any
}

type ManagementConnectionDetailsWrapper struct {
	pam.ManagementConnectionDetails
}
type DatabaseResourceResponseWrapper struct {
	pam.DatabaseResourceResponse
}
type MySQLBasicAuthManagementConnectionDetailsWrapper struct {
	pam.MySQLBasicAuthManagementConnectionDetails
}
type MySQLBasicAuthDetailsWrapper struct {
	pam.MySQLBasicAuthDetails
}

func (w DatabaseResourceResponseWrapper) ToResourceMap(o AttributeOverrides) map[string]any {
	m := make(map[string]any, 10)

	m[attributes.CanonicalName] = w.CanonicalName
	m[attributes.DatabaseType] = w.DatabaseType
	m[attributes.ManagementConnectionDetailsType] = w.ManagementConnectionDetailsType
	if w.ManagementGatewaySelectorId != "" {
		m[attributes.ManagementGatewaySelectorID] = w.ManagementGatewaySelectorId
	}
	if w.ManagementGatewaySelector != nil {
		m[attributes.ManagementGatewaySelector] = *w.ManagementGatewaySelector
	} else {
		m[attributes.ManagementGatewaySelector] = make(map[string]string)
	}

	mgmtDetails := make([]any, 1)
	mgmtDetails[0] = ManagementConnectionDetailsWrapper{w.ManagementConnectionDetails}.ToResourceMap(o)

	m[attributes.ManagementConnectionDetails] = mgmtDetails

	return m
}

func (w ManagementConnectionDetailsWrapper) ToResourceMap(o AttributeOverrides) map[string]any {
	if w.MySQLBasicAuthManagementConnectionDetails != nil {
		return MySQLBasicAuthManagementConnectionDetailsWrapper{
			*w.MySQLBasicAuthManagementConnectionDetails,
		}.ToResourceMap(o)
	}
	return nil
}

func (w MySQLBasicAuthManagementConnectionDetailsWrapper) ToResourceMap(o AttributeOverrides) map[string]any {
	m := make(map[string]any, 3)

	m[attributes.Hostname] = w.Hostname
	m[attributes.Port] = w.Port

	authDetails := make([]any, 1)
	authDetails[0] = MySQLBasicAuthDetailsWrapper{w.AuthDetails}.ToResourceMap(o)
	m[attributes.AuthDetails] = authDetails

	return m
}

func (w MySQLBasicAuthDetailsWrapper) ToResourceMap(o AttributeOverrides) map[string]any {
	m := make(map[string]any, 3)
	m[attributes.Username] = w.Username
	if v, ok := o[attributes.Password]; ok {
		m[attributes.Password] = v
	}
	m[attributes.Secret] = w.SecretId

	return m
}
