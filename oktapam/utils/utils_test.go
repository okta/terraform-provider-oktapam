package utils

import (
	"reflect"
	"testing"
)

func TestExpandStringList(t *testing.T) {
	input := []any{"foo", "bar", "foo"}
	outputList := ExpandStringList(input)
	expected := []string{
		"foo",
		"bar",
		"foo",
	}

	if !reflect.DeepEqual(outputList, expected) {
		t.Fatalf(
			"Got:\n\n%#v\n\nExpected:\n\n%#v\n",
			outputList,
			expected)
	}
}
