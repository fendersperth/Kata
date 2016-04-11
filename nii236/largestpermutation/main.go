package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type testCase struct {
	numInts  int
	numSwap  int
	numArray []int
}

// main is used if reading from STDIN
func main() {
	ans, err := Entry(os.Stdin)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ans)
}

// Entry is the entry point for this Go package
func Entry(r io.Reader) (string, error) {
	args := ParseArgs(r)
	ans := largestPermutation(args)
	return ans, nil
}

// ParseArgs will convert STDIN to the required types
func ParseArgs(r io.Reader) testCase {
	var numInts int
	var numSwaps int
	var numArray []int

	scanner := bufio.NewScanner(r)

	fmt.Fscanf(r, "%d %d\n", &numInts, &numSwaps)

	scanner.Scan()
	stringArray := strings.Split(scanner.Text(), " ")
	for _, i := range stringArray {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		numArray = append(numArray, j)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return testCase{numInts, numSwaps, numArray}
}

func largestPermutation(input testCase) string {
	var currSlice []int
	maxPos := getMax(input.numArray, 0)
	for i := 0; i < input.numSwap; i++ {
		currSlice = swap(input.numArray, i, maxPos)
		getMax(input.numArray, i)
	}

	var ans string

	for i, val := range currSlice {
		ans += strconv.Itoa(val)
		if i == len(currSlice)-1 {
			break
		}
		ans += " "
	}

	return ans
}

func swap(in []int, a int, b int) []int {
	var out []int
	var tmp int
	for i := 0; i < len(in); i++ {
		if i == a {
			out = append(out, in[b])
			tmp = in[a]
		} else if i == b {
			out = append(out, tmp)
		} else {
			out = append(out, in[i])
		}
	}
	return out
}

func getMax(in []int, fromPos int) int {
	maxPos := 0

	for i := fromPos; i < len(in); i++ {
		if in[i] > in[maxPos] {
			maxPos = i
		}
	}

	return maxPos
}
