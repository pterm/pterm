package pterm

import (
	"github.com/pterm/pterm/internal"
	"github.com/stretchr/testify/assert"
	"io"
	"reflect"
	"testing"
)

func TestNewRGB(t *testing.T) {
	type args struct {
		r uint8
		g uint8
		b uint8
	}
	tests := []struct {
		name string
		args args
		want RGB
	}{
		{name: "1", args: args{0, 0, 0}, want: RGB{0, 0, 0}},
		{name: "3", args: args{255, 255, 255}, want: RGB{255, 255, 255}},
		{name: "4", args: args{127, 127, 127}, want: RGB{127, 127, 127}},
		{name: "5", args: args{1, 2, 3}, want: RGB{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRGB(tt.args.r, tt.args.g, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRGB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRGBFromHEX(t *testing.T) {
	tests := []struct {
		hex  string
		want RGB
	}{
		{hex: "#ff0009", want: RGB{R: 255, G: 0, B: 9}},
		{hex: "ff0009", want: RGB{R: 255, G: 0, B: 9}},
		{hex: "ff00090x", want: RGB{R: 255, G: 0, B: 9}},
		{hex: "ff00090X", want: RGB{R: 255, G: 0, B: 9}},
		{hex: "#fba", want: RGB{R: 255, G: 187, B: 170}},
		{hex: "fba", want: RGB{R: 255, G: 187, B: 170}},
		{hex: "fba0x", want: RGB{R: 255, G: 187, B: 170}},
	}
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			rgb, err := NewRGBFromHEX(test.hex)
			assert.Equal(t, test.want, rgb)
			assert.NoError(t, err)
		})
	}
	testsFail := []struct {
		hex  string
		want error
	}{
		{hex: "faba0x", want: ErrHexCodeIsInvalid},
		{hex: "faba", want: ErrHexCodeIsInvalid},
		{hex: "#faba", want: ErrHexCodeIsInvalid},
		{hex: "faba0x", want: ErrHexCodeIsInvalid},
		{hex: "#fax", want: assert.AnError},
	}
	for _, test := range testsFail {
		t.Run("", func(t *testing.T) {
			_, err := NewRGBFromHEX(test.hex)
			assert.Error(t, test.want, err)
		})
	}
}

func TestRGB_Fade(t *testing.T) {
	type fields struct {
		R uint8
		G uint8
		B uint8
	}
	type args struct {
		min     float32
		max     float32
		current float32
		end     []RGB
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   RGB
	}{
		{name: "Middle", fields: fields{0, 0, 0}, args: args{min: 0, max: 100, current: 50, end: []RGB{{255, 255, 255}}}, want: RGB{127, 127, 127}},
		{name: "ZeroToZero", fields: fields{0, 0, 0}, args: args{min: 0, max: 100, current: 50, end: []RGB{{0, 0, 0}}}, want: RGB{0, 0, 0}},
		{name: "DifferentValues", fields: fields{0, 1, 2}, args: args{min: 0, max: 100, current: 50, end: []RGB{{0, 1, 2}}}, want: RGB{0, 1, 2}},
		{name: "NegativeRangeMiddle", fields: fields{0, 0, 0}, args: args{min: -50, max: 50, current: 0, end: []RGB{{255, 255, 255}}}, want: RGB{127, 127, 127}},
		{name: "NegativeRangeMiddleMultipleRGB", fields: fields{0, 0, 0}, args: args{min: -50, max: 50, current: 0, end: []RGB{{127, 127, 127}, {255, 255, 255}}}, want: RGB{127, 127, 127}},
		{name: "MiddleMultipleRGB", fields: fields{0, 0, 0}, args: args{min: 0, max: 100, current: 50, end: []RGB{{127, 127, 127}, {255, 255, 255}}}, want: RGB{127, 127, 127}},
		{name: "1/4MultipleRGB", fields: fields{0, 0, 0}, args: args{min: 0, max: 100, current: 25, end: []RGB{{255, 255, 255}, {255, 255, 255}}}, want: RGB{127, 127, 127}},
		{name: "MiddleMultipleRGBPositiveMin", fields: fields{0, 0, 0}, args: args{min: 10, max: 110, current: 60, end: []RGB{{127, 127, 127}, {255, 255, 255}}}, want: RGB{127, 127, 127}},
		{name: "MiddleNoRGB", fields: fields{0, 0, 0}, args: args{min: 10, max: 110, current: 60, end: []RGB{}}, want: RGB{0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := RGB{
				R: tt.fields.R,
				G: tt.fields.G,
				B: tt.fields.B,
			}
			if got := p.Fade(tt.args.min, tt.args.max, tt.args.current, tt.args.end...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fade() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRGB_GetValues(t *testing.T) {
	type fields struct {
		R uint8
		G uint8
		B uint8
	}
	tests := []struct {
		name   string
		fields fields
		wantR  uint8
		wantG  uint8
		wantB  uint8
	}{
		{name: "Zero", fields: fields{R: 0, G: 0, B: 0}, wantR: uint8(0), wantG: uint8(0), wantB: uint8(0)},
		{name: "Max", fields: fields{R: 255, G: 255, B: 255}, wantR: uint8(255), wantG: uint8(255), wantB: uint8(255)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := RGB{
				R: tt.fields.R,
				G: tt.fields.G,
				B: tt.fields.B,
			}
			gotR, gotG, gotB := p.GetValues()
			if gotR != tt.wantR {
				t.Errorf("GetValues() gotR = %v, want %v", gotR, tt.wantR)
			}
			if gotG != tt.wantG {
				t.Errorf("GetValues() gotG = %v, want %v", gotG, tt.wantG)
			}
			if gotB != tt.wantB {
				t.Errorf("GetValues() gotB = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func TestRGB_Print(t *testing.T) {
	RGBs := []RGB{{0, 0, 0}, {127, 127, 127}, {255, 255, 255}}

	for _, rgb := range RGBs {
		t.Run(Sprintf("%v %v %v", rgb.R, rgb.G, rgb.B), func(t *testing.T) {
			internal.TestPrintContains(t, func(w io.Writer, a interface{}) {
				p := rgb.Print(a)
				assert.NotNil(t, p)
			})
		})
	}
}

func TestRGB_Printf(t *testing.T) {
	RGBs := []RGB{{0, 0, 0}, {127, 127, 127}, {255, 255, 255}}

	for _, rgb := range RGBs {
		t.Run(Sprintf("%v %v %v", rgb.R, rgb.G, rgb.B), func(t *testing.T) {
			internal.TestPrintfContains(t, func(w io.Writer, format string, a interface{}) {
				p := rgb.Printf(format, a)
				assert.NotNil(t, p)
			})
		})
	}
}

func TestRGB_Println(t *testing.T) {
	RGBs := []RGB{{0, 0, 0}, {127, 127, 127}, {255, 255, 255}}

	for _, rgb := range RGBs {
		t.Run(Sprintf("%v %v %v", rgb.R, rgb.G, rgb.B), func(t *testing.T) {
			internal.TestPrintlnContains(t, func(w io.Writer, a interface{}) {
				p := rgb.Println(a)
				assert.NotNil(t, p)
			})
		})
	}
}

func TestRGB_Sprint(t *testing.T) {
	RGBs := []RGB{{0, 0, 0}, {127, 127, 127}, {255, 255, 255}}

	for _, rgb := range RGBs {
		t.Run(Sprintf("%v %v %v", rgb.R, rgb.G, rgb.B), func(t *testing.T) {
			internal.TestSprintContains(t, func(a interface{}) string {
				return rgb.Sprint(a)
			})
		})
	}
}

func TestRGB_Sprintf(t *testing.T) {
	RGBs := []RGB{{0, 0, 0}, {127, 127, 127}, {255, 255, 255}}

	for _, rgb := range RGBs {
		t.Run("", func(t *testing.T) {
			internal.TestSprintfContains(t, func(format string, a interface{}) string {
				return rgb.Sprintf(format, a)
			})
		})
	}
}

func TestRGB_Sprintln(t *testing.T) {
	RGBs := []RGB{{0, 0, 0}, {127, 127, 127}, {255, 255, 255}}

	for _, rgb := range RGBs {
		t.Run(Sprintf("%v %v %v", rgb.R, rgb.G, rgb.B), func(t *testing.T) {
			internal.TestSprintlnContains(t, func(a interface{}) string {
				return rgb.Sprintln(a)
			})
		})
	}
}
