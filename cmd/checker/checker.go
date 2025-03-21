package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// ✅ Safely handle command-line arguments
	var stack string
	if len(os.Args) > 1 {
		stack = os.Args[1] // Read argument if exists
		fmt.Println("Command-line argument:", stack)
	} else {
		fmt.Println("No command-line argument provided.")
	}

	// ✅ Read from stdin
	NewScanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Reading from stdin...")

	for NewScanner.Scan() {
		println(1) // Debugging: Ensure it runs
		println(NewScanner.Text()) // Print each line
	}

	// ✅ Handle scanner errors
	if err := NewScanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}
