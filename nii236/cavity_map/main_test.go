package main_test

import (
	"reflect"
	"strings"
	"testing"

	. "github.com/nii236/Kata/nii236/cavity_map"
)

var tests = []struct {
	in       string
	out      string
	numLines int
	numArray [][]int
}{
	{
		in: `4
1112
1912
1892
1234`,
		out: `1112
1X12
18X2
1234`,
		numLines: 4,
		numArray: [][]int{
			{1, 1, 1, 2},
			{1, 9, 1, 2},
			{1, 8, 9, 2},
			{1, 2, 3, 4},
		},
	},
}

var parseArgsTests = []struct {
	in       string
	outLines int
	outArray [][]int
}{{}}

func TestEntry(t *testing.T) {
	for _, test := range tests {
		input := strings.NewReader(test.in)
		got, err := Entry(input)
		if err != nil {
			return
		}
		expected := test.out
		if expected != got {
			t.Errorf("\n\nError! \n\nExpected:\n%s\n\nGot:\n%s", expected, got)
		}

	}
}

func TestParseArgs(t *testing.T) {
	for _, test := range tests {
		expectedNumLines := test.numLines
		expectedNumArray := test.numArray

		input := strings.NewReader(test.in)
		gotNumLines, gotNumArray, err := ParseArgs(input)
		if err != nil {
			t.Errorf("Error: %s", err)
			return
		}
		if expectedNumLines != gotNumLines {
			t.Errorf("\n\nError in numLines! \n\nExpected:\n%d\n\nGot:\n%d", expectedNumLines, gotNumLines)
		}

		if !reflect.DeepEqual(expectedNumArray, gotNumArray) {
			t.Errorf("\n\nError in numArray! \n\nExpected:\n%d\n\nGot:\n%d", expectedNumArray, gotNumArray)
		}

	}
}
