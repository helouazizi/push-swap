// internal/algorithm/sorting.go
package algorithm

import (
	"fmt"
	"push-swap/internal/stack"
)

// SortStack applies the sorting algorithm to stack a using stack b as auxiliary.
func SortStack(a, b *stack.Stack) []string {
	var operations []string

	// Simple sorting logic (placeholder for an optimal algorithm)
	for len(a.Items) > 0 {
		minIdx := findMaxIndex(a)  
		fmt.Println(minIdx,a.Items[minIdx])          // Find the index of the minimum element
		moveToTop(a, b, minIdx, &operations) // Move the minimum element to the top of stack a
		stack.Push(a, b)                     // Push the minimum element to stack b
		operations = append(operations, "pb")

		// Debugging: Print the current state of the stacks
		// fmt.Println("After operation: pb")
		
		fmt.Println("Stack A:", a.Items)
		fmt.Println("Stack B:", b.Items)
	}

	// Move everything back to stack a
	for len(b.Items) > 0 {
		stack.Push(b, a) // Move elements from stack b back to stack a
		operations = append(operations, "pa")
		// Debugging: Print the current state of the stacks
		// fmt.Println("After operation: pa")
		// fmt.Println("Stack A:", a.Items)
		// fmt.Println("Stack B:", b.Items)
	}

	return operations
}

// findMinIndex finds the index of the smallest element in the stack.
func findMaxIndex(s *stack.Stack) int {
	maxIdx := 0
	maxVal := s.Items[0]
	for i, val := range s.Items {
		if val > maxVal {
			maxVal = val
			maxIdx = i
		}
	}
	return maxIdx
}

// moveToTop moves the smallest element to the top using the fewest rotations.
func moveToTop(a, b *stack.Stack, index int, operations *[]string) {
	if index >= len(a.Items)/2 {
		// Rotate to move the smallest element to the top
		for i := index; i < len(a.Items); i++ {
			if index > len(a.Items) {
				fmt.Println(true)
			}
			a.Rotate()
			*operations = append(*operations, "ra")
		}
		fmt.Println(a.Items,"inside moove up ra")
	} else {
		// Reverse rotate to move the smallest element to the top
		fmt.Println("first")
		for i := len(a.Items)-1 - index; i > 0; i-- {
			a.ReverseRotate()
			fmt.Println(a.Items,"rra")
			*operations = append(*operations, "rra")
		}
		fmt.Println(a.Items,"inside moove up rra")
	}
}
