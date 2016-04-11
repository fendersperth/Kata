package main

import (
	"reflect"
	"strings"
	"testing"
)

var swapTests = []struct {
	in  []int
	a   int
	b   int
	out []int
}{
	{
		in:  []int{1, 2, 3},
		a:   1,
		b:   2,
		out: []int{1, 3, 2},
	},
	{
		in:  []int{1, 2},
		a:   0,
		b:   1,
		out: []int{2, 1},
	},
}

func TestLargestPermutation(t *testing.T) {
	for _, test := range tests {
		got := largestPermutation(test.testCase)
		expected := test.out

		if expected != got {
			t.Errorf("\n\nError! \n\nExpected: \n'%v', \n\nGot: \n'%v'.", expected, got)
		}
	}

}

func TestParseArgs(t *testing.T) {
	for _, test := range tests {
		r := strings.NewReader(test.in)
		got := ParseArgs(r)
		expected := test.testCase

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("\n\nError! \n\nExpected: \n'%v', \n\nGot: \n'%v'.", expected, got)
		}
	}
}

func TestSwap(t *testing.T) {
	for _, test := range swapTests {
		got := swap(test.in, test.a, test.b)
		expected := test.out
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("\n\nError! \n\nExpected: \n'%v', \n\nGot: \n'%v'.", expected, got)
		}

	}
}
