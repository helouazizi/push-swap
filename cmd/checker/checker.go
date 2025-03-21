package main

import (
	"bufio"
	"os"
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
		
	}
	if err := NewScanner.Err(); err != nil {
		println(err)
	}

	println(stack, instarctions)
}
