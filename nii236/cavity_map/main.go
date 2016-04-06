package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

// main is used if reading from STDIN
func main() {
	Entry(os.Stdin)
}

// Entry is the entry point for this Go package
func Entry(r io.Reader) (string, error) {
	numLines, numArray, err := ParseArgs(r)
	if err != nil {
		return "", err
	}
	ans := cavity(numLines, numArray)

	return ans, nil
}

// ParseArgs will convert STDIN to the required types
func ParseArgs(r io.Reader) (int, [][]int, error) {
	numLines := 0
	var err error

	_, err = fmt.Fscanln(r, &numLines)
	if err != nil {
		return numLines, nil, err
	}

	numArray := make([][]int, numLines, numLines)
	var lines []string
	for i := 0; i < numLines; i++ {
		var line string
		_, err = fmt.Fscanln(r, &line)
		if err == io.EOF {
			break
		}
		lines = append(lines, line)
	}

	for i, line := range lines {
		for _, char := range line {
			currChar, err := strconv.Atoi(string(char))
			if err != nil {
				fmt.Println(err)
			}
			numArray[i] = append(numArray[i], currChar)
		}
	}

	return numLines, numArray, nil
}

func cavity(numLines int, numArray [][]int) string {
	var result string
	for i := 0; i < numLines; i++ {
		for j := 0; j < numLines; j++ {
			if (i == 0 || i == numLines-1) || (j == 0 || j == numLines-1) {
				result += strconv.Itoa(numArray[i][j])
				continue
			} else {

				if isDeepest(numArray, i, j) {
					result += "X"
				} else {
					result += strconv.Itoa(numArray[i][j])
				}
			}
		}
		if i == numLines-1 {
			continue
		}
		result += "\n"
	}
	return result
}

func isDeepest(numArray [][]int, i int, j int) bool {
	test := numArray[i][j]
	if numArray[i-1][j] > test {
		return false
	} else if numArray[i+1][j] > test {
		return false
	} else if numArray[i][j-1] > test {
		return false
	} else if numArray[i][j+1] > test {
		return false
	}
	return true
}
