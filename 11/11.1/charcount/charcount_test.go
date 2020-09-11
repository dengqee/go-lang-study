package charcount

import (
	"io"
	"testing"
	_ "unicode/utf8"
)

type in struct {
	text string
}

func (i in) Read(p []byte) (n int, err error) {
	for _, c := range []byte(i.text) {
		p = append(p, c)
	}
	return len(i.text), nil
}

func TestCharcount(t *testing.T) {
	var tests = []struct {
		input io.Reader
	}{
		{in{""}},
		{in{"a"}},
		{in{"ab"}},
		{in{"a db"}},
	}
	for _, test := range tests {
		Charcount(test.input)
	}
}
