// internal/algorithm/sorting.go
package algorithm

import (
	"push-swap/internal/stack"
)

// SortStack applies the sorting algorithm to stack a using stack b as auxiliary.
func SortStack(a, b *stack.Stack) []string {
	var operations []string

	// Simple sorting logic (placeholder for an optimal algorithm)
	for len(a.Items) > 0 {
		minIdx := findMinIndex(a)
		moveToTop(a, b, minIdx, &operations)
		stack.Push(a, b)
		operations = append(operations, "pb")
	}

	// Move everything back to stack a
	for len(b.Items) > 0 {
		stack.Push(b, a)
		operations = append(operations, "pa")
	}

	return operations
}

// findMinIndex finds the index of the smallest element in the stack.
func findMinIndex(s *stack.Stack) int {
	minIdx := 0
	minVal := s.Items[0]
	for i, val := range s.Items {
		if val < minVal {
			minVal = val
			minIdx = i
		}
	}
	return minIdx
}

// moveToTop moves the smallest element to the top using the fewest rotations.
func moveToTop(a, b *stack.Stack, index int, operations *[]string) {
	if index < len(a.Items)/2 {
		for i := 0; i < index; i++ {
			a.Rotate()
			*operations = append(*operations, "ra")
		}
	} else {
		for i := len(a.Items) - index; i > 0; i-- {
			a.ReverseRotate()
			*operations = append(*operations, "rra")
		}
	}
}
