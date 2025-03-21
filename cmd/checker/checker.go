package main

import (
	"bufio"
	"fmt"
	"os"
	"push-swap/internal/stacks"
	"push-swap/internal/utils"
	"strings"
)

func main() {
	if len(os.Args) <= 1 {
		println()
		return
	}
	

	stack := os.Args[1]
	instarctions := []string{}
	NewScanner := bufio.NewScanner(os.Stdin)
	for NewScanner.Scan() {
		if NewScanner.Text() != "" && !utils.IsExist(stacks.Instarctions, strings.TrimSpace(NewScanner.Text())) {
			println("Error: instruction do not exist :", NewScanner.Text())
			return
		}
	
		if NewScanner.Text() != "" {
			instarctions = append(instarctions, strings.TrimSpace(NewScanner.Text()))
		}
		
	}
	if err := NewScanner.Err(); err != nil {
		println(err)
	}

	fmt.Println(stack, instarctions,len(instarctions))
}
