package putils

import (
	"cmp"
	"fmt"
	"slices"

	"github.com/pterm/pterm"
)

// Number is any integer or float type.
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

// BarsFromMap turns a map into bars, largest value first, so the chart
// reads like a ranking. Equal values are sorted by label.
//
//	pterm.BarChart.WithBars(putils.BarsFromMap(sizes)).Render()
func BarsFromMap[K cmp.Ordered, V Number](m map[K]V) pterm.Bars {
	bars := make(pterm.Bars, 0, len(m))

	for k, v := range m {
		bars = append(bars, pterm.Bar{Label: fmt.Sprint(k), Value: float64(v)})
	}

	slices.SortFunc(bars, func(a, b pterm.Bar) int {
		if c := cmp.Compare(b.Value, a.Value); c != 0 {
			return c
		}

		return cmp.Compare(a.Label, b.Label)
	})

	return bars
}
