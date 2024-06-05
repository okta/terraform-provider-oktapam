package wrappers

import (
	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
)

type SudoCommandsBundleWrapper struct {
	SudoCommandBundle *pam.SudoCommandBundle
}

func (w SudoCommandsBundleWrapper) ToResourceMap() map[string]any {
	m := make(map[string]any)
	m[attributes.SudoCommmandsBundle] = w.SudoCommandBundle
	return m
}
