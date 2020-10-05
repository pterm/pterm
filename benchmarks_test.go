package pterm

import (
	"fmt"
	"testing"
)

var benchmarkText = "This is a Benchmark Text"

func BenchmarkFmt_Sprintln(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintln(benchmarkText)
	}
}

func BenchmarkSprintln(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sprintln(benchmarkText)
	}
}
