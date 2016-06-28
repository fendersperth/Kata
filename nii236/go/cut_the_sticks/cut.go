package cut

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Exec is the entry point of package cut
func Exec(input string) string {
	r := strings.NewReader(input)
	numArgs, args := GetArgs(r)
	_, _, numCut := cut(numArgs, args, []int{})
	var ans []string
	for _, cut := range numCut {
		ans = append(ans, strconv.Itoa(cut))
	}
	joinedAns := strings.Join(ans, "\n")
	return joinedAns
}

// GetArgs reads from STDIN and returns stuff in the correct format
func GetArgs(r io.Reader) (int, []int) {
	var numArgs = 0
	_, err := fmt.Fscanln(r, &numArgs)

	if err != nil {
		fmt.Println(err)
	}

	sumArgs := make([]int, numArgs)
	for i := range sumArgs {
		_, err = fmt.Fscan(r, &sumArgs[i])
		if err != nil {
			fmt.Println(err)
		}
	}

	return numArgs, sumArgs
}

func cut(numArgs int, args []int, finalArray []int) (int, []int, []int) {

	var newArray []int
	var newFinalArray []int
	var zeroValue bool

	minArg := args[0]

	// Get smallest number in args
	for _, arg := range args {
		if arg < minArg {
			minArg = arg
		}
		if arg == 0 {
			zeroValue = true
		}
	}

	// If non zero, append to resulting array, less the minimum arg
	for _, arg := range args {
		if arg != 0 {
			newArray = append(newArray, arg-minArg)
		}
	}

	// If there exists any zero length sticks, run function again
	if zeroValue {
		return cut(len(newArray), newArray, finalArray)
	}

	// If more than 1 stick remain, run function recursively
	if numArgs > 1 {
		newFinalArray = append(finalArray, len(newArray))
		return cut(len(newArray), newArray, newFinalArray)
	}

	newFinalArray = append(finalArray, len(newArray))
	return len(newArray), newArray, newFinalArray

}
