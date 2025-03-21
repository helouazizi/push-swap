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
	NewScanner := bufio.NewScanner(os.Stdin)
	NewScanner.Scan()
	instarctions := NewScanner.Text()
	println(stack, instarctions)
}
