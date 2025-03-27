// cmd/push-swap/push-swap.go
package main

import (
	"fmt"
	"os"
	"push-swap/internal/algorithm"
	"push-swap/internal/parser"
	"push-swap/internal/stack"
	"slices"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	input := os.Args[1]
	values, err := parser.ParseArgs(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	a := stack.New()
	b := stack.New()

	// to folow the subject instarctions
	//The first integer should be at the top of the stack.
	// so lets reverse the slice
	slices.Reverse(values)
	for _, val := range values {
		a.Push(val)
	}
	fmt.Println(a.Items)

	operations := algorithm.SortStack(a, b)
	fmt.Println(operations)
	fmt.Println(a.Items)
	
	// for _, op := range operations {
	// 	fmt.Println(op)
	// }
}
