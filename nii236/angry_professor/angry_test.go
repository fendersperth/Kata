package main

import (
	"reflect"
	"strings"
	"testing"
)

var tests = []struct {
	in        string
	out       string
	testCases []testCase
}{
	{
		in: `2
4 3
-1 -3 4 2
4 2
0 -1 2 1`,

		out: `YES
NO`,
		testCases: []testCase{
			{
				numStudents:     4,
				cancelThreshold: 3,
				arrivalTimes:    []int{-1, -3, 4, 2},
			},
			{
				numStudents:     4,
				cancelThreshold: 2,
				arrivalTimes:    []int{0, -1, 2, 1},
			},
		},
	},
}

// func TestAngry(t *testing.T) {
// 	for _, test := range tests {
// 		inReader := strings.NewReader(test.in)
// 		got := Entry(inReader)
// 		expected := test.out
// 		if got != expected {
// 			t.Errorf("\n\nError! \n\nExpected: \n'%s', \n\nGot: \n'%s'.", expected, got)
// 		}
// 	}
// }

func TestParseArgs(t *testing.T) {
	for _, test := range tests {
		r := strings.NewReader(test.in)
		got := ParseArgs(r)
		expected := test.testCases

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("\n\nError! \n\nExpected: \n'%v', \n\nGot: \n'%v'.", expected, got)
		}
	}
}
