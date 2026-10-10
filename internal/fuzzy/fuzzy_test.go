package fuzzy_test

import (
	"slices"
	"testing"

	"github.com/pterm/pterm/internal/fuzzy"
)

func TestMatch(t *testing.T) {
	t.Parallel()

	tests := []struct {
		pattern, target string
		ok              bool
		positions       []int
	}{
		{"", "anything", true, nil},
		{"k8s", "kapp-k8s-at", true, []int{5, 6, 7}},
		{"KAT", "kapp-k8s-at", true, []int{0, 9, 10}},
		{"npm", "npm cache", true, []int{0, 1, 2}},
		{"gc", "go build cache", true, []int{0, 9}},
		{"xyz", "kapp-k8s-at", false, nil},
		{"long pattern", "short", false, nil},
		{"ü", "Über", true, []int{0}},
	}

	for _, tt := range tests {
		_, pos, ok := fuzzy.Match(tt.pattern, tt.target)
		if ok != tt.ok || !slices.Equal(pos, tt.positions) {
			t.Errorf("Match(%q, %q) = %v, %v; want %v, %v", tt.pattern, tt.target, pos, ok, tt.positions, tt.ok)
		}
	}
}

func TestMatchRanking(t *testing.T) {
	t.Parallel()

	better, _, _ := fuzzy.Match("cache", "npm cache")
	worse, _, _ := fuzzy.Match("cache", "click a check here")

	if better <= worse {
		t.Errorf("consecutive match scored %d, scattered match %d", better, worse)
	}

	start, _, _ := fuzzy.Match("go", "go build cache")
	middle, _, _ := fuzzy.Match("go", "cargo registry")

	if start <= middle {
		t.Errorf("match at the start scored %d, in the middle %d", start, middle)
	}
}
