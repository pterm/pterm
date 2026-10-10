// Package fuzzy matches a typed filter against option labels the way
// pickers in editors do: every typed letter must appear in the label in
// order, and labels where the letters sit together or start words rank
// higher.
package fuzzy

import (
	"strings"
	"unicode"
)

// Scores of the matcher. A match earns matchScore per letter, plus bonuses
// for letters that start a word or follow the previous match, minus a
// penalty for every gap between matched letters.
const (
	matchScore        = 16
	boundaryBonus     = 8
	firstBonus        = 10
	consecutiveBonus  = 12
	exactCaseBonus    = 1
	gapStartPenalty   = 3
	gapExtendPenalty  = 1
	leadingGapPenalty = 1
	maxLeadingPenalty = 6
)

const minScore = -1 << 30

// Match reports whether pattern matches target. Every rune of pattern
// (spaces aside) must appear in target in order, ignoring case. It returns
// the score of the best match, higher is better, and the rune indexes of
// target that matched, for highlighting. An empty pattern matches
// everything with a score of zero.
func Match(pattern, target string) (score int, positions []int, ok bool) {
	pat := []rune(strings.ReplaceAll(pattern, " ", ""))
	if len(pat) == 0 {
		return 0, nil, true
	}

	tgt := []rune(target)

	m, n := len(pat), len(tgt)
	if m > n {
		return 0, nil, false
	}

	// best[i][j] is the best score with pat[i] matched at tgt[j]; from[i][j]
	// is where pat[i-1] matched on that path.
	best := make([][]int, m)
	from := make([][]int, m)

	for i := range m {
		best[i] = make([]int, n)
		from[i] = make([]int, n)

		for j := range n {
			best[i][j] = minScore
		}
	}

	for i := range m {
		// run holds the best score of a match of pat[i-1] at least two
		// runes before j, already reduced by the gap penalty up to j.
		run, runFrom := minScore, -1

		for j := i; j < n; j++ {
			if i > 0 && j >= 2 {
				if prev := best[i-1][j-2]; prev > minScore && prev-gapStartPenalty > run-gapExtendPenalty {
					run, runFrom = prev-gapStartPenalty, j-2
				} else if run > minScore {
					run -= gapExtendPenalty
				}
			}

			if !equalFold(pat[i], tgt[j]) {
				continue
			}

			s := matchScore + bonus(tgt, j)
			if pat[i] == tgt[j] {
				s += exactCaseBonus
			}

			if i == 0 {
				best[i][j] = s - min(j*leadingGapPenalty, maxLeadingPenalty)
				from[i][j] = -1

				continue
			}

			if prev := best[i-1][j-1]; prev > minScore && prev+consecutiveBonus >= run {
				best[i][j] = prev + s + consecutiveBonus
				from[i][j] = j - 1
			} else if run > minScore {
				best[i][j] = run + s
				from[i][j] = runFrom
			}
		}
	}

	end, top := -1, minScore

	for j := range n {
		if best[m-1][j] > top {
			end, top = j, best[m-1][j]
		}
	}

	if end < 0 {
		return 0, nil, false
	}

	positions = make([]int, m)
	for i, j := m-1, end; i >= 0; i-- {
		positions[i] = j
		j = from[i][j]
	}

	return top, positions, true
}

// bonus rewards a match at the start of the target or of a word in it.
func bonus(t []rune, j int) int {
	if j == 0 {
		return firstBonus
	}

	prev, cur := t[j-1], t[j]

	switch {
	case !unicode.IsLetter(prev) && !unicode.IsDigit(prev):
		return boundaryBonus
	case unicode.IsLower(prev) && unicode.IsUpper(cur):
		return boundaryBonus
	case unicode.IsLetter(prev) != unicode.IsLetter(cur):
		return boundaryBonus / 2
	}

	return 0
}

func equalFold(a, b rune) bool {
	return a == b || unicode.ToLower(a) == unicode.ToLower(b)
}
