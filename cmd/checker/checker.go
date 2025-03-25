package main

import (
	"bufio"
	"fmt"
	"os"
	"push-swap/internal/stacks"
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

	stackB := stacks.New_stacks()

	// Read and execute operations
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		op := Operation(scanner.Text())

		// Validate operation
		validOp := false
		for _, knownOp := range []Operation{PA, PB, SA, SB, SS, RA, RB, RR, RRA, RRB, RRR} {
			if op == knownOp {
				validOp = true
				break
			}
		}

		if !validOp || !executeOperation(stackA, stackB, op) {
			fmt.Fprintln(os.Stderr, "Error")
			os.Exit(1)
		}
	}

	// Check final state
	if len(stackB.items) == 0 && isSorted(stackA) {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}

func executeOperation(stackA, stackB *Stack, op Operation) bool {
	switch op {
	case PA:
		val, ok := stackB.Pop()
		if !ok {
			return false
		}
		stackA.Push(val)

	case PB:
		val, ok := stackA.Pop()
		if !ok {
			return false
		}
		stackB.Push(val)

	case SA:
		stackA.Swap()

	case SB:
		stackB.Swap()

	case SS:
		stackA.Swap()
		stackB.Swap()

	case RA:
		stackA.Rotate()

	case RB:
		stackB.Rotate()

	case RR:
		stackA.Rotate()
		stackB.Rotate()

	case RRA:
		stackA.ReverseRotate()

	case RRB:
		stackB.ReverseRotate()

	case RRR:
		stackA.ReverseRotate()
		stackB.ReverseRotate()
	}

	return true
}

func isSorted(stack *Stack) bool {
	items := stack.items
	for i := 1; i < len(items); i++ {
		if items[i-1] > items[i] {
			return false
		}
	}
	return true
}
