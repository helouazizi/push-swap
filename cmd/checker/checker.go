// cmd/checker/checker.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"push-swap/internal/parser"
	"push-swap/internal/stack"
	"strings"
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

	for _, val := range values {
		a.Push(val)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		instruction := strings.TrimSpace(scanner.Text())
		switch instruction {
		case "pa":
			stack.Push(b, a)
		case "pb":
			stack.Push(a, b)
		case "sa":
			a.Swap()
		case "sb":
			b.Swap()
		case "ss":
			stack.SwapBoth(a, b)
		case "ra":
			a.Rotate()
		case "rb":
			b.Rotate()
		case "rr":
			stack.RotateBoth(a, b)
		case "rra":
			a.ReverseRotate()
		case "rrb":
			b.ReverseRotate()
		case "rrr":
			stack.ReverseRotateBoth(a, b)
		default:
			fmt.Fprintln(os.Stderr, "Error")
			os.Exit(1)
		}
	}

	if isSorted(a) && len(b.Items) == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}

// isSorted checks if the stack is sorted in ascending order.
func isSorted(s *stack.Stack) bool {
	items := s.Items
	for i := 1; i < len(items); i++ {
		if items[i-1] > items[i] {
			return false
		}
	}
	return true
}
