package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type testCase struct {
	numStudents     int
	cancelThreshold int
	arrivalTimes    []int
}

// main is the entry point for executables
func main() {
	Entry(os.Stdin)
}

// Entry is the entry point for this package
func Entry(r io.Reader) string {
	ParseArgs(r)
	ans := angry()
	return ans
}

// ParseArgs parses an io.Reader into the types required to solve the puzzle
func ParseArgs(r io.Reader) (int, int, []int) {
	var testCases int
	var numStudents int
	var cancelThreshold int
	var arrayString string
	var arrivalTimes []int

	fmt.Fscanf(r, "%d\n", &testCases)
	fmt.Println()
	fmt.Println("TESTCASES:", testCases)
	for i := 0; i < testCases; i++ {
		fmt.Println("Iterating through case", i)
		fmt.Fscanf(r, "%d %d\n", &numStudents, &cancelThreshold)
		fmt.Fscanln(r, &arrayString)
		arrivalTimes = parseLine(arrayString)
		fmt.Println("NUMSTUDENTS:", numStudents)
		fmt.Println("CANCELTHRESHOLD:", cancelThreshold)
		fmt.Println("ARRAYSTRING:", arrayString)
		fmt.Println("ARRIVALTIMES:", arrivalTimes)
	}

	fmt.Println()

	return numStudents, cancelThreshold, arrivalTimes
}

func angry() string {
	return ""
}

func parseLine(line string) []int {
	var numArray []int
	arrayString := strings.Split(line, " ")
	for _, el := range arrayString {
		currInt, err := strconv.Atoi(el)
		if err != nil {
			fmt.Println(err)
		}
		numArray = append(numArray, currInt)
	}
	return numArray
}
