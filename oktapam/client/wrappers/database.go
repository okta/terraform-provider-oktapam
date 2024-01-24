package wrappers

import (
	"fmt"
	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
)

var (
	_ ResourceWrapper = (*DatabaseResourceResponseWrapper)(nil)
	_ ResourceWrapper = (*ManagementConnectionDetailsWrapper)(nil)
	_ ResourceWrapper = (*MySQLBasicAuthManagementConnectionDetailsWrapper)(nil)
	_ ResourceWrapper = (*MySQLBasicAuthDetailsWrapper)(nil)
)

// The key is the single attribute field name, and the value is what will be set for that field.
type attributeOverrides map[string]any

type ResourceWrapper interface {
	// AttributeOverridePaths returns the attribute key paths that can be set based on existing values in state.
	// These are returned recursively, so this should typically only be executed on a top level struct.
	AttributeOverridePaths() []string
	// ToResourceMap converts a resource struct into a map of attribute keys to values.
	// Optional attributeOverrides can be provided to override any resource value.
	// utils.GenerateAttributeOverrides can be used to generate the required overrides for any resource struct.
	ToResourceMap(attributeOverrides) map[string]any
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

func (w DatabaseResourceResponseWrapper) ToResourceMap(o attributeOverrides) map[string]any {
	m := make(map[string]any, 10)

	m[attributes.CanonicalName] = w.CanonicalName
	m[attributes.DatabaseType] = w.DatabaseType
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

func (w DatabaseResourceResponseWrapper) AttributeOverridePaths() []string {
	overrides := []string{}
	childOverrides := ManagementConnectionDetailsWrapper{w.ManagementConnectionDetails}.AttributeOverridePaths()
	for _, attr := range childOverrides {
		overrides = append(overrides, fmt.Sprintf("%s.0.%s", attributes.ManagementConnectionDetails, attr))
	}
	return overrides
}

func (w ManagementConnectionDetailsWrapper) ToResourceMap(o attributeOverrides) map[string]any {
	if w.MySQLBasicAuthManagementConnectionDetails == nil {
		return nil
	}

	m := map[string]any{}

	mysqlDetails := make([]any, 1)
	mysqlDetails[0] = MySQLBasicAuthManagementConnectionDetailsWrapper{
		*w.MySQLBasicAuthManagementConnectionDetails,
	}.ToResourceMap(o)

	m[attributes.MySQL] = mysqlDetails

	return m
}

func (w ManagementConnectionDetailsWrapper) AttributeOverridePaths() []string {
	overrides := []string{}
	if w.MySQLBasicAuthManagementConnectionDetails == nil {
		return overrides
	}

	childOverrides := MySQLBasicAuthManagementConnectionDetailsWrapper{*w.MySQLBasicAuthManagementConnectionDetails}.AttributeOverridePaths()
	for _, attr := range childOverrides {
		overrides = append(overrides, fmt.Sprintf("%s.0.%s", attributes.MySQL, attr))
	}
	return overrides
}

func (w MySQLBasicAuthManagementConnectionDetailsWrapper) ToResourceMap(o attributeOverrides) map[string]any {
	m := make(map[string]any, 3)

	m[attributes.Hostname] = w.Hostname
	m[attributes.Port] = w.Port

	authDetails := make([]any, 1)
	authDetails[0] = MySQLBasicAuthDetailsWrapper{w.AuthDetails}.ToResourceMap(o)
	m[attributes.BasicAuth] = authDetails

	return m
}

func (w MySQLBasicAuthManagementConnectionDetailsWrapper) AttributeOverridePaths() []string {
	overrides := []string{}
	childOverrides := MySQLBasicAuthDetailsWrapper{w.AuthDetails}.AttributeOverridePaths()
	for _, attr := range childOverrides {
		overrides = append(overrides, fmt.Sprintf("%s.0.%s", attributes.BasicAuth, attr))
	}
	return overrides
}

func (w MySQLBasicAuthDetailsWrapper) ToResourceMap(o attributeOverrides) map[string]any {
	m := make(map[string]any, 3)
	m[attributes.Username] = w.Username
	m[attributes.Secret] = w.SecretId

	if v, ok := o[attributes.Password]; ok {
		m[attributes.Password] = v
	}

	return m
}

func (w MySQLBasicAuthDetailsWrapper) AttributeOverridePaths() []string {
	return []string{attributes.Password}
}
