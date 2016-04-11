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

	{

		in: `10
10 4
-93 -86 49 -62 -90 -63 40 72 11 67
10 10
23 -35 -2 58 -67 -56 -42 -73 -19 37
10 9
13 91 56 -62 96 -5 -84 -36 -46 -13
10 8
45 67 64 -50 -8 78 84 -51 99 60
10 7
26 94 -95 34 67 -97 17 52 1 86
10 2
19 29 -17 -71 -75 -27 -56 -53 65 83
10 10
-32 -3 -70 8 -40 -96 -18 -46 -21 -79
10 9
-50 0 64 14 -56 -91 -65 -36 51 -28
10 6
-58 -29 -35 -18 43 -28 -76 43 -13 6
10 1
88 -17 -96 43 83 99 25 90 -39 86`,

		out: `NO
YES
YES
YES
YES
NO
YES
YES
NO
NO`,
		testCases: []testCase{
			{
				numStudents:     10,
				cancelThreshold: 4,
				arrivalTimes:    []int{-93, -86, 49, -62, -90, -63, 40, 72, 11, 67},
			},
			{
				numStudents:     10,
				cancelThreshold: 10,
				arrivalTimes:    []int{23, -35, -2, 58, -67, -56, -42, -73, -19, 37},
			},
			{
				numStudents:     10,
				cancelThreshold: 9,
				arrivalTimes:    []int{13, 91, 56, -62, 96, -5, -84, -36, -46, -13},
			},
			{
				numStudents:     10,
				cancelThreshold: 8,
				arrivalTimes:    []int{45, 67, 64, -50, -8, 78, 84, -51, 99, 60},
			},
			{
				numStudents:     10,
				cancelThreshold: 7,
				arrivalTimes:    []int{26, 94, -95, 34, 67, -97, 17, 52, 1, 86},
			},
			{
				numStudents:     10,
				cancelThreshold: 2,
				arrivalTimes:    []int{19, 29, -17, -71, -75, -27, -56, -53, 65, 83},
			},
			{
				numStudents:     10,
				cancelThreshold: 10,
				arrivalTimes:    []int{-32, -3, -70, 8, -40, -96, -18, -46, -21, -79},
			},
			{
				numStudents:     10,
				cancelThreshold: 9,
				arrivalTimes:    []int{-50, 0, 64, 14, -56, -91, -65, -36, 51, -28},
			},
			{
				numStudents:     10,
				cancelThreshold: 6,
				arrivalTimes:    []int{-58, -29, -35, -18, 43, -28, -76, 43, -13, 6},
			},
			{
				numStudents:     10,
				cancelThreshold: 1,
				arrivalTimes:    []int{88, -17, -96, 43, 83, 99, 25, 90, -39, 86},
			},
		},
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
