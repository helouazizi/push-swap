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
	fmt.Println(all_stacks.Itims, instarctions)
	// lets execute the instarcions on the stacks
	if len(stacks.Instarctions) != 0 {
		//all_stacks.Execute_Instarcrions(instarctions)
	}
	//now if the stack a is sorted  and the stack b is empty
	// display ok folowed by \n
	// otherwise display ko
	fmt.Println("all instrctons are completed")

}
