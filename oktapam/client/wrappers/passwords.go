package wrappers

import (
	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
)

var (
	_ ResourceWrapper = (*PasswordPolicyWrapper)(nil)
	_ ResourceWrapper = (*CharacterOptionsWrapper)(nil)
)

type PasswordPolicyWrapper struct {
	pam.PasswordPolicy
	UseManagedPrivilegedAccountsConfig bool
}

func (w PasswordPolicyWrapper) ToResourceMap(overrides attributeOverrides) map[string]any {
	m := make(map[string]any, 7)

	m[attributes.EnablePeriodicRotation] = w.EnablePeriodicRotation
	m[attributes.MinLength] = w.MinLengthInBytes
	m[attributes.MaxLength] = int(w.MaxLengthInBytes)
	if w.PeriodicRotationDurationInSeconds != nil {
		m[attributes.PeriodicRotationDurationInSeconds] = *w.PeriodicRotationDurationInSeconds
	}
	if w.UseManagedPrivilegedAccountsConfig {
		acctsArr := make([]any, len(w.ManagedPrivilegedAccountsConfig))
		for idx, a := range w.ManagedPrivilegedAccountsConfig {
			acctsArr[idx] = a
			m[attributes.ManagedPrivilegedAccounts] = acctsArr
		}
	}

	cArr := make([]any, 1)
	cArr[0] = CharacterOptionsWrapper{w.CharacterOptions}.ToResourceMap(overrides)
	m[attributes.CharacterOptions] = cArr

	return m
}

func (w PasswordPolicyWrapper) AttributeOverridePaths() []string {
	return nil
}

type CharacterOptionsWrapper struct {
	pam.PasswordPolicyCharacterOptions
}

func (w CharacterOptionsWrapper) ToResourceMap(overrides attributeOverrides) map[string]any {
	m := make(map[string]any, 4)

	if w.LowerCase != nil {
		m[attributes.LowerCase] = *w.LowerCase
	}
	if w.UpperCase != nil {
		m[attributes.UpperCase] = *w.UpperCase
	}
	if w.Digits != nil {
		m[attributes.Digits] = *w.Digits
	}
	if w.Punctuation != nil {
		m[attributes.Punctuation] = *w.Punctuation
	}
	if w.RequireFromEachSet != nil {
		m[attributes.RequireFromEachSet] = *w.RequireFromEachSet
	}

	return m
}

func (w CharacterOptionsWrapper) AttributeOverridePaths() []string {
	return nil
}
