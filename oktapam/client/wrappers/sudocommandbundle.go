package wrappers

import (
	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
)

type SudoCommandBundleWrapper struct {
	SudoCommandBundle *pam.SudoCommandBundle
}

func (w SudoCommandBundleWrapper) ToResourceMap() map[string]any {
	m := make(map[string]any)
	m[attributes.ID] = w.SudoCommandBundle.Id
	m[attributes.Name] = w.SudoCommandBundle.Name
	m[attributes.AddEnv] = w.SudoCommandBundle.AddEnv
	m[attributes.SubEnv] = w.SudoCommandBundle.SubEnv
	if w.SudoCommandBundle.NoExec.IsSet() {
		m[attributes.NoExec] = w.SudoCommandBundle.NoExec.Get()
	}
	if w.SudoCommandBundle.NoPasswd.IsSet() {
		m[attributes.NoPasswd] = w.SudoCommandBundle.NoPasswd.Get()
	}
	if w.SudoCommandBundle.RunAs.IsSet() {
		m[attributes.RunAs] = w.SudoCommandBundle.RunAs.Get()
	}
	if w.SudoCommandBundle.SetEnv.IsSet() {
		m[attributes.SetEnv] = w.SudoCommandBundle.SetEnv.Get()

	}
	structuredCommands := make([]map[string]any, len(w.SudoCommandBundle.StructuredCommands))
	for i, sc := range w.SudoCommandBundle.StructuredCommands {
		structuredCommands[i] = map[string]any{}
		structuredCommands[i][attributes.StructuredCommand] = sc.Command
		structuredCommands[i][attributes.StructuredCommandType] = sc.CommandType
		structuredCommands[i][attributes.StructuredCommandArgs] = sc.Args
		structuredCommands[i][attributes.StructuredCommandArgsType] = sc.ArgsType
	}
	m[attributes.StructuredCommands] = structuredCommands
	return m
}
