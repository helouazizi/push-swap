package main

import (
	"fmt"
	"os"
	"push-swap/internal/utils"
)

func main() {
	if len(os.Args) <= 1 {
		println()
		return
	}

	stack := os.Args[1]
	instarctions, err := utils.Scsn_Input()
	if err != nil {
		fmt.Println(err)
		return
	}
	all_stacks, err := utils.Parse_stack(stack)
	if err != nil {
		fmt.Println(err)
		return
	}
	num, err := all_stacks.Pop(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(all_stacks.Stack_A, instarctions)
	fmt.Println("top elem from stack a", num)
}
