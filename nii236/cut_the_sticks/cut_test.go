package cut_test

import (
	"fmt"
	"testing"

	"github.com/nii236/Kata/nii236/cut_the_sticks"
)

var tests = []struct {
	in  string
	out string
}{
	{
		in: `6
5 4 4 2 2 8`,
		out: `6
4
2
1`,
	},
	{
		in: `8
1 2 3 4 3 3 2 1`,
		out: `8
6
4
1`,
	},
}

func TestCut(t *testing.T) {
	for i, test := range tests {
		fmt.Println("\nBegin test suite", i)
		got := cut.Exec(test.in)
		expected := test.out

		if got != expected {
			t.Errorf("\n\nTest failed!\nExpected:\n%s, \ngot:\n%s\n", expected, got)
		}

	}

}
