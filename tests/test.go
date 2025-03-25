package main

import (
	"fmt"
	"os"
	"strconv"
)



func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	// Parse initial stack
	stackA := parseArgs(args)
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

func parseArgs(args []string) *Stack {
	stack := &Stack{}

	// Parse arguments into integers
	for i := len(args) - 1; i >= 0; i-- {
		num, err := strconv.Atoi(args[i])
		if err != nil || contains(stack.items, num) {
			return nil
		}
		stack.Push(num)
	}

	return stack
}

func contains(items []int, target int) bool {
	for _, item := range items {
		if item == target {
			return true
		}
	}
	return false
}
