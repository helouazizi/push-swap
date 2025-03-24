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
		// lts get the stack a size
		sizeA := len(realstack.Stack_A)
		// lets loop overr all the elemrnys on the stack a
		for e := 0 ; e < sizeA ; e++ {
			//lets get the top element
			top , _ := realstack.Pop(0)
			if (top>>b)&1 ==0 {
				fmt.Println("pb")
				realstack.Pb(top)
			}else {
				realstack.Pa(top)
				fmt.Println("ra")
				realstack.Ra()
			}
		}
	}
}
