package main

import (
	"fmt"
	"os"
	"push-swap/internal/utils"
	"slices"
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
	fmt.Println(slices.Reverse(stack))
	fmt.Println(stack, instarctions, len(instarctions))
}
