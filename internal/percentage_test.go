package internal_test

import (
	"testing"

	"github.com/pterm/pterm/internal"
)

func TestPercentage(t *testing.T) {
	type args struct {
		total   float64
		current float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "HalfOfHundert", args: args{total: 100, current: 50}, want: 50},
		{name: "HalfOfTwoHundert", args: args{total: 200, current: 100}, want: 50},
		{name: "HalfOfFiveHundert", args: args{total: 500, current: 250}, want: 50},
		{name: "QuarterOfFiveHundert", args: args{total: 500, current: 100}, want: 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := internal.Percentage(tt.args.total, tt.args.current); got != tt.want {
				t.Errorf("Percentage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPercentageRound(t *testing.T) {
	type args struct {
		total   float64
		current float64
		max     float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "HalfOfHundert", args: args{total: 100, current: 50}, want: 50},
		{name: "HalfOfTwoHundert", args: args{total: 200, current: 100}, want: 50},
		{name: "HalfOfFiveHundert", args: args{total: 500, current: 250}, want: 50},
		{name: "QuarterOfFiveHundert", args: args{total: 500, current: 100}, want: 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := internal.PercentageRound(tt.args.total, tt.args.current); got != tt.want {
				t.Errorf("PercentageRound() = %v, want %v", got, tt.want)
			}
		})
	}
}
