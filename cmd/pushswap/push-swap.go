package main

import (
	"fmt"
	"os"
	"push-swap/internal/utils"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	// Parse initial stack
	stackA := utils.ParseArgs(args)
	if stackA == nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	// Generate sorting operations
	operations := generateOperations(stackA)

	// Print operations
	for _, op := range operations {
		fmt.Println(string(op))
	}
}
