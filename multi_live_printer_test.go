package pterm

import (
	"bytes"
	"testing"
)

func TestMultiPrinterGetString(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		name string
		buf  string
		want string
	}{
		{
			name: "test1",
			buf:  "test\rtest2\rtest3",
			want: "test3\n",
		},
		{
			name: "test2",
			buf:  "\r\n",
			want: "",
		},
		{
			name: "test3",
			buf:  "test",
			want: "test\n",
		},
		{
			name: "test4",
			buf:  "",
			want: "",
		},
	} {
		test := test // pin for pre-go1.22 versions
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			p := DefaultMultiPrinter
			p.buffers = []*bytes.Buffer{
				bytes.NewBufferString(test.buf),
			}

			if test.want != p.getString() {
				t.Errorf("got %v, want %v", p.getString(), test.want)
			}
		})
	}
}
