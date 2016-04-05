package main

import (
	"strings"
	"testing"
)

var tests = []struct {
	in  string
	out string
}{
	{
		in: `2
4 3
-1 -3 4 2
4 2
0 -1 2 1`,

		out: `YES
NO`,
	},
}

func TestAngry(t *testing.T) {
	for _, test := range tests {
		inReader := strings.NewReader(test.in)
		got := Entry(inReader)
		expected := test.out
		if got != expected {
			t.Errorf("\n\nError! \n\nExpected: \n'%s', \n\nGot: \n'%s'.", expected, got)
		}
	}
}

func BenchAngry(b *testing.B) {

}
