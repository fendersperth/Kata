package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
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
func ParseArgs(r io.Reader) []testCase {
	var testCases int

	scanner := bufio.NewScanner(r)

	fmt.Fscanf(r, "%d\n", &testCases)
	result := []testCase{}
	for i := 0; i < testCases; i++ {
		var curr testCase

		scanner.Scan()
		args := strings.Split(scanner.Text(), " ")
		var err error

		curr.numStudents, err = strconv.Atoi(args[0])
		curr.cancelThreshold, err = strconv.Atoi(args[1])

		scanner.Scan()

		arrivalTimes := scanner.Text()

		sArray := strings.Split(arrivalTimes, " ")
		timesArray := []int{}
		for _, i := range sArray {
			var j int
			j, err = strconv.Atoi(i)
			if err != nil {
				panic(err)
			}
			timesArray = append(timesArray, j)
		}
		curr.arrivalTimes = timesArray

		if err != nil {
			panic(err)
		}

		result = append(result, curr)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func angry() string {
	return ""
}
