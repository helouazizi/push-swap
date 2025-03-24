package main

import (
	"fmt"
	"push-swap/internal/stacks"
)

func main()  {
	stack := []int{3,2,1}
	fmt.Println(stack,"befor")
	radixsort(stack)
	fmt.Println(stack,"after")
}

func radixsort(stack []int)  {
	realstack := stacks.New_stacks()
	realstack.Stack_A = stack
	fmt.Println(realstack.Stack_A)
}

