package main

import (
	"fmt"
	"push-swap/internal/stacks"
	"push-swap/internal/utils"
)

func main() {
	stack := []int{3, 2, 1}
	fmt.Println(stack, "befor")
	radixsort(stack)
	fmt.Println(stack, "after")
}

func radixsort(stack []int) {
	realstack := stacks.New_stacks()
	realstack.Stack_A = stack
	max := utils.GetMax(realstack)
	bits := utils.BitsCount(max)
	// lets loop over bits
	for b := 0 ; b < bits; b++ {
		
	}
}
