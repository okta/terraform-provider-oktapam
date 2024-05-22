package client

import (
	"strings"
	"testing"

	"github.com/okta/terraform-provider-oktapam/oktapam/utils"
)

func getLongName(t *testing.T, size int) string {
	t.Helper()
	var sb strings.Builder

	for i := 0; i < size; i++ {
		sb.WriteString("a")
	}

	return sb.String()
}

func TestSudoCommandsBundleValidation(t *testing.T) {
	tests := []struct {
		name string
		scb  SudoCommandsBundle
	}{
		{
			name: "name is nil",
			scb: SudoCommandsBundle{
				Id:   utils.AsStringPtr("fakeID"),
				Name: nil,
				StructuredCommands: []StructuredCommand{
					{
						CommandType: utils.AsStringPtr("raw"),
						Command:     utils.AsStringPtr("sh"),
						Args:        nil,
						ArgsType:    nil,
					},
				},
				RunAs:    utils.AsStringPtr("sudo"),
				NoPasswd: utils.AsBoolPtr(false),
				NoExec:   utils.AsBoolPtr(false),
				SetEnv:   utils.AsBoolPtr(false),
				AddEnv:   []string{"X", "Y"},
				SubEnv:   []string{"A", "B"},
			},
		},
		{
			name: "empty name",
			scb: SudoCommandsBundle{
				Id:   utils.AsStringPtr("fakeID"),
				Name: utils.AsStringPtr(""),
				StructuredCommands: []StructuredCommand{
					{
						CommandType: utils.AsStringPtr("raw"),
						Command:     utils.AsStringPtr("sh"),
						Args:        nil,
						ArgsType:    nil,
					},
				},
				RunAs:    utils.AsStringPtr("sudo"),
				NoPasswd: utils.AsBoolPtr(false),
				NoExec:   utils.AsBoolPtr(false),
				SetEnv:   utils.AsBoolPtr(false),
				AddEnv:   []string{"X", "Y"},
				SubEnv:   []string{"A", "B"},
			},
		},
		{
			name: "long name",
			scb: SudoCommandsBundle{
				Id:   utils.AsStringPtr("fakeID"),
				Name: utils.AsStringPtr(getLongName(t, 256)),
				StructuredCommands: []StructuredCommand{
					{
						CommandType: utils.AsStringPtr("raw"),
						Command:     utils.AsStringPtr("sh"),
						Args:        nil,
						ArgsType:    nil,
					},
				},
				RunAs:    utils.AsStringPtr("sudo"),
				NoPasswd: utils.AsBoolPtr(false),
				NoExec:   utils.AsBoolPtr(false),
				SetEnv:   utils.AsBoolPtr(false),
				AddEnv:   []string{"X", "Y"},
				SubEnv:   []string{"A", "B"},
			},
		},
		{
			name: "invalid name",
			scb: SudoCommandsBundle{
				Id:   utils.AsStringPtr("fakeID"),
				Name: utils.AsStringPtr("x-y-z"),
				StructuredCommands: []StructuredCommand{
					{
						CommandType: utils.AsStringPtr("raw"),
						Command:     utils.AsStringPtr("sh"),
						Args:        nil,
						ArgsType:    nil,
					},
				},
				RunAs:    utils.AsStringPtr("sudo"),
				NoPasswd: utils.AsBoolPtr(false),
				NoExec:   utils.AsBoolPtr(false),
				SetEnv:   utils.AsBoolPtr(false),
				AddEnv:   []string{"X", "Y"},
				SubEnv:   []string{"A", "B"},
			},
		},
		{
			name: "empty run_as",
			scb: SudoCommandsBundle{
				Id:   utils.AsStringPtr("fakeID"),
				Name: utils.AsStringPtr("fakeName"),
				StructuredCommands: []StructuredCommand{
					{
						CommandType: utils.AsStringPtr("raw"),
						Command:     utils.AsStringPtr("sh"),
						Args:        nil,
						ArgsType:    nil,
					},
				},
				RunAs:    utils.AsStringPtr(""),
				NoPasswd: utils.AsBoolPtr(false),
				NoExec:   utils.AsBoolPtr(false),
				SetEnv:   utils.AsBoolPtr(false),
				AddEnv:   []string{"X", "Y"},
				SubEnv:   []string{"A", "B"},
			},
		},
		{
			name: "invalid run_as",
			scb: SudoCommandsBundle{
				Id:   utils.AsStringPtr("fakeID"),
				Name: utils.AsStringPtr("fakeName"),
				StructuredCommands: []StructuredCommand{
					{
						CommandType: utils.AsStringPtr("raw"),
						Command:     utils.AsStringPtr("sh"),
						Args:        nil,
						ArgsType:    nil,
					},
				},
				RunAs:    utils.AsStringPtr("@"),
				NoPasswd: utils.AsBoolPtr(false),
				NoExec:   utils.AsBoolPtr(false),
				SetEnv:   utils.AsBoolPtr(false),
				AddEnv:   []string{"X", "Y"},
				SubEnv:   []string{"A", "B"},
			},
		},
		{
			name: "invalid add_env",
			scb: SudoCommandsBundle{
				Id:   utils.AsStringPtr("fakeID"),
				Name: utils.AsStringPtr("fakeName"),
				StructuredCommands: []StructuredCommand{
					{
						CommandType: utils.AsStringPtr("raw"),
						Command:     utils.AsStringPtr("sh"),
						Args:        nil,
						ArgsType:    nil,
					},
				},
				RunAs:    utils.AsStringPtr("sudo"),
				NoPasswd: utils.AsBoolPtr(false),
				NoExec:   utils.AsBoolPtr(false),
				SetEnv:   utils.AsBoolPtr(false),
				AddEnv:   []string{"-"},
				SubEnv:   []string{"A", "B"},
			},
		},
		{
			name: "invalid sub_env",
			scb: SudoCommandsBundle{
				Id:   utils.AsStringPtr("fakeID"),
				Name: utils.AsStringPtr("fakeName"),
				StructuredCommands: []StructuredCommand{
					{
						CommandType: utils.AsStringPtr("raw"),
						Command:     utils.AsStringPtr("sh"),
						Args:        nil,
						ArgsType:    nil,
					},
				},
				RunAs:    utils.AsStringPtr("sudo"),
				NoPasswd: utils.AsBoolPtr(false),
				NoExec:   utils.AsBoolPtr(false),
				SetEnv:   utils.AsBoolPtr(false),
				AddEnv:   []string{"X", "Y"},
				SubEnv:   []string{"-"},
			},
		},
		{
			name: "nil structured command",
			scb: SudoCommandsBundle{
				Id:   utils.AsStringPtr("fakeID"),
				Name: utils.AsStringPtr("fakeName"),
				StructuredCommands: []StructuredCommand{
					{
						CommandType: utils.AsStringPtr("raw"),
						Command:     nil,
						Args:        nil,
						ArgsType:    nil,
					},
				},
				RunAs:    utils.AsStringPtr("sudo"),
				NoPasswd: utils.AsBoolPtr(false),
				NoExec:   utils.AsBoolPtr(false),
				SetEnv:   utils.AsBoolPtr(false),
				AddEnv:   []string{"X", "Y"},
				SubEnv:   []string{"A", "B"},
			},
		},
		{
			name: "invalid structured command type",
			scb: SudoCommandsBundle{
				Id:   utils.AsStringPtr("fakeID"),
				Name: utils.AsStringPtr("fakeName"),
				StructuredCommands: []StructuredCommand{
					{
						CommandType: utils.AsStringPtr("invalid"),
						Command:     utils.AsStringPtr("sh"),
						Args:        nil,
						ArgsType:    nil,
					},
				},
				RunAs:    utils.AsStringPtr("sudo"),
				NoPasswd: utils.AsBoolPtr(false),
				NoExec:   utils.AsBoolPtr(false),
				SetEnv:   utils.AsBoolPtr(false),
				AddEnv:   []string{"X", "Y"},
				SubEnv:   []string{"A", "B"},
			},
		},
		{
			name: "invalid directory structured command",
			scb: SudoCommandsBundle{
				Id:   utils.AsStringPtr("fakeID"),
				Name: utils.AsStringPtr("fakeName"),
				StructuredCommands: []StructuredCommand{
					{
						CommandType: utils.AsStringPtr("directory"),
						Command:     utils.AsStringPtr("abc"),
						Args:        nil,
						ArgsType:    nil,
					},
				},
				RunAs:    utils.AsStringPtr("sudo"),
				NoPasswd: utils.AsBoolPtr(false),
				NoExec:   utils.AsBoolPtr(false),
				SetEnv:   utils.AsBoolPtr(false),
				AddEnv:   []string{"X", "Y"},
				SubEnv:   []string{"A", "B"},
			},
		},
		{
			name: "invalid executable structured command",
			scb: SudoCommandsBundle{
				Id:   utils.AsStringPtr("fakeID"),
				Name: utils.AsStringPtr("fakeName"),
				StructuredCommands: []StructuredCommand{
					{
						CommandType: utils.AsStringPtr("executable"),
						Command:     utils.AsStringPtr("directory/"),
						Args:        nil,
						ArgsType:    nil,
					},
				},
				RunAs:    utils.AsStringPtr("sudo"),
				NoPasswd: utils.AsBoolPtr(false),
				NoExec:   utils.AsBoolPtr(false),
				SetEnv:   utils.AsBoolPtr(false),
				AddEnv:   []string{"X", "Y"},
				SubEnv:   []string{"A", "B"},
			},
		},
		{
			name: "invalid raw structured command",
			scb: SudoCommandsBundle{
				Id:   utils.AsStringPtr("fakeID"),
				Name: utils.AsStringPtr("fakeName"),
				StructuredCommands: []StructuredCommand{
					{
						CommandType: utils.AsStringPtr("raw"),
						Command:     utils.AsStringPtr("---"),
						Args:        nil,
						ArgsType:    nil,
					},
				},
				RunAs:    utils.AsStringPtr("sudo"),
				NoPasswd: utils.AsBoolPtr(false),
				NoExec:   utils.AsBoolPtr(false),
				SetEnv:   utils.AsBoolPtr(false),
				AddEnv:   []string{"X", "Y"},
				SubEnv:   []string{"A", "B"},
			},
		},
		{
			name: "invalid raw structured command args and args type",
			scb: SudoCommandsBundle{
				Id:   utils.AsStringPtr("fakeID"),
				Name: utils.AsStringPtr("fakeName"),
				StructuredCommands: []StructuredCommand{
					{
						CommandType: utils.AsStringPtr("raw"),
						Command:     utils.AsStringPtr("/bin/script.sh"),
						Args:        utils.AsStringPtr("foo"),
						ArgsType:    utils.AsStringPtr("bar"),
					},
				},
				RunAs:    utils.AsStringPtr("sudo"),
				NoPasswd: utils.AsBoolPtr(false),
				NoExec:   utils.AsBoolPtr(false),
				SetEnv:   utils.AsBoolPtr(false),
				AddEnv:   []string{"X", "Y"},
				SubEnv:   []string{"A", "B"},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotValid, _ := isSudoCommandsBundleValid(tc.scb)
			if gotValid != false {
				t.Errorf("isSudoCommandsBundleValid() = %v, want %v", gotValid, false)
			}
		})
	}
}
