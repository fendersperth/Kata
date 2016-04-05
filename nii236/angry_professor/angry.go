package main

import (
	"io"
	"os"
)

// main is the entry point for executables
func main() {
	Entry(os.Stdin)
}

// Entry is the entry point for this package
func Entry(r io.Reader) string {
	parseArgs(r)
	ans := angry()
	return ans
}

func parseArgs(r io.Reader) {

}

func angry() string {
	return ""
}
