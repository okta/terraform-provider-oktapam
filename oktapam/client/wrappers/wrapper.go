package wrappers

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
