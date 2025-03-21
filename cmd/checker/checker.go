package main

import (
	"fmt"
	"os"
	"push-swap/internal/stacks"
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
	// lets execute the instarcions on the stacks
	if len(stacks.Instarctions) != 0 {
		all_stacks.Execute_Instarcrions(instarctions)
	}
 	num, err := all_stacks.Pop(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("top elem from stack a", num)
	fmt.Println(all_stacks.Stack_A, instarctions)
	all_stacks.Sa()
	fmt.Println(all_stacks.Stack_A, instarctions)
	all_stacks.Ra()
	fmt.Println(all_stacks.Stack_A, instarctions)
	all_stacks.Rra()
	fmt.Println(all_stacks.Stack_A, instarctions)
	// now all instarctions are completed

}
