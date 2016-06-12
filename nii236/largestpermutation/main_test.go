package main

import (
	"reflect"
	"strconv"
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

var maxTests = []struct {
	in   []int
	from int
	out  int
}{
	{
		in:   []int{1, 2, 3},
		from: 0,
		out:  2,
	},
	{
		in:   []int{1, 2, 3},
		from: 1,
		out:  2,
	},
	{
		in:   []int{1, 2, 3},
		from: 2,
		out:  2,
	},
	{
		in:   []int{3, 2, 1},
		from: 0,
		out:  0,
	},
	{
		in:   []int{3, 2, 1},
		from: 1,
		out:  1,
	},
	{
		in:   []int{3, 2, 1},
		from: 2,
		out:  2,
	},
	{
		in:   []int{5, 4, 3, 2, 1},
		from: 2,
		out:  2,
	},
}

func TestLargestPermutation(t *testing.T) {
	for i, test := range tests {
		t.Log("Running test case:", i)
		got := largestPermutation(test.testCase)
		for i, out := range test.ans {
			if out != got[i] {
				len := strconv.Itoa(len(test.ans))
				t.Log("index", i, "of", len)
				t.Errorf("\n\nError! \n\nExpected: \n'%v', \n\nGot: \n'%v'.", out, got[i])
				break
			}
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

func TestGetMax(t *testing.T) {
	for _, test := range maxTests {
		got := getMax(test.in, test.from)
		expected := test.out
		if got != expected {
			t.Errorf("\n\nError! \n\nExpected: \n'%v', \n\nGot: \n'%v'.", expected, got)
		}
	}
}
