package pterm

import (
	"reflect"
	"testing"
)

func Test_sanitizeArgs(t *testing.T) {
	tests := []struct {
		name     string
		args     []any
		expected []any
	}{
		{
			name:     "pass_zero_args",
			args:     []any{},
			expected: []any{},
		},
		{
			name:     "pass_one_arg",
			args:     []any{"foo"},
			expected: []any{ErrKeyWithoutValue, "foo"},
		},
		{
			name:     "pass_two_args",
			args:     []any{"foo", "bar"},
			expected: []any{"foo", "bar"},
		},
		{
			name:     "pass_three_args",
			args:     []any{"foo", "bar", "baz"},
			expected: []any{"foo", "bar", ErrKeyWithoutValue, "baz"},
		},
	}
	l := Logger{}
	for _, tt := range tests {
		if got := l.sanitizeArgs(tt.args); !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("sanitizeArgs: got: %v, expected: %v", got, tt.expected)
		}
	}
}
