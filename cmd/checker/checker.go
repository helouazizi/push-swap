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
	instarctions := ""
	NewScanner := bufio.NewScanner(os.Stdin)
	for NewScanner.Scan() {
		instarctions += NewScanner.Text()
	}

	println(stack, instarctions)
}
