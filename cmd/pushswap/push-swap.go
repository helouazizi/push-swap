// cmd/push-swap/push-swap.go
package main

import (
	"fmt"
	"os"
	"push-swap/internal/parser"
	"push-swap/internal/stack"
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
	//b := stack.New()

	for _, val := range values {
		a.Push(val)
	}
	fmt.Println(a.Items)

	//operations := algorithm.SortStack(a, b)

	// for _, op := range operations {
	// 	fmt.Println(op)
	// }
}
