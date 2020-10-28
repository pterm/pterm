package pterm

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

var benchmarkText = "This is a Benchmark Text"

func BenchmarkFmtPrint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprint(benchmarkText)
	}
}

func BenchmarkSprint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Sprint(benchmarkText)
	}
}

func BenchmarkSprintWithColor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Sprint(Cyan(benchmarkText))
	}
}

func BenchmarkSprintWithCustomStyle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Sprint(NewStyle(FgCyan).Sprint(benchmarkText))
	}
}

func BenchmarkSpinner(b *testing.B) {
	proxyToDevNull()
	for i := 0; i < b.N; i++ {
		s := DefaultSpinner.WithDelay(0).Start()
		s.Stop()
	}
}

func BenchmarkPrefixPrinter(b *testing.B) {
	proxyToDevNull()
	printers := []PrefixPrinter{Info, Success, Warning, Error, *Fatal.WithFatal(false)}
	for _, p := range printers {
		b.Run(p.Prefix.Text, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				p.Print(benchmarkText)
			}
		})
	}
}

func BenchmarkHeader(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DefaultHeader.Sprint(benchmarkText)
	}
}

func BenchmarkProgressbar(b *testing.B) {
	benchmarks := []struct {
		total int
	}{
		{total: 10},
		// {total: 100},
		// {total: 1000},
	}
	for _, bm := range benchmarks {
		proxyToDevNull()
		b.Run("Total"+strconv.Itoa(bm.total), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				p := DefaultProgressbar.WithTotal(bm.total).Start()
				for i := 0; i < bm.total; i++ {
					p.Increment()
				}
			}
		})
	}
}

func BenchmarkTreeList_Render(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var tis []TreeListItem
		for i := 0; i < 10000; i++ {
			tis = append(tis, TreeListItem{
				ItemName: "Test",
			})
		}
		SetDefaultOutput(os.NewFile(0, os.DevNull))
		DefaultTreeList.WithItems(tis).Render()
	}
}

func proxyToDevNull() {
	SetDefaultOutput(os.NewFile(0, os.DevNull))
}
