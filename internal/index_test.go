package internal_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm/internal"
)

func TestGetIndex(t *testing.T) {
	tests := []struct {
		name  string
		slice []string
		i     int
		want  string
	}{
		{name: "index inside bounds", slice: []string{"a", "b", "c"}, i: 1, want: "b"},
		{name: "index at first element", slice: []string{"a", "b", "c"}, i: 0, want: "a"},
		{name: "index at last element", slice: []string{"a", "b", "c"}, i: 2, want: "c"},
		{name: "index beyond length returns last element", slice: []string{"a", "b", "c"}, i: 99, want: "c"},
		{name: "single element slice", slice: []string{"only"}, i: 5, want: "only"},
		{name: "negative index", slice: []string{"only"}, i: -1, want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, internal.Get_Index(tt.slice, tt.i))
		})
	}
}
