package main

import (
	"fmt"
	"math"
	"os"
	"push-swap/internal/stacks"
	"push-swap/internal/utils"
)

/*
		since we gona sort the stack a without any comarison that mean we gonna use
		non-comparative sorting algorithms like counting and flash and radix sorting algorithms and much more
		so in this case we have stacks the beast algo is the radix algorithm

		Since we can only use push (pa, pb), swap (sa, sb, ss), rotate (ra, rb, rr), and reverse rotate (rra, rrb, rrr),
		we need to adapt a non-comparative sorting method. Radix Sort is the best candidate since it sorts numbers digit
		by digit without direct comparisons.

		### Implementing Radix Sort Using Stacks
		 	Steps:
	         1 Determine the Maximum Number of Digits

	            Find the maximum number in Stack A to determine how many digits we need to sort.

	        2 Sort Digit by Digit (Least Significant Digit First)

	            For each digit (starting from the least significant):

	            Distribute numbers into Stack B based on the current digit (using pb).

	            Collect numbers back into Stack A in the correct order (using pa).

	            Repeat for the next digit.


		******************
		Algorithm Strategy
	       Find the maximum number in Stack A to determine how many digits we need to process.

	       Sort the numbers digit by digit (starting from the least significant digit).

	       Use Stack B as a helper bucket to distribute numbers based on their digit.

	       Reconstruct Stack A in the correct order after processing each digit.
*/
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
	fmt.Println(all_stacks.Stack_A, instarctions, "befor sorting")
	Radix_Sort(all_stacks)
	fmt.Println(all_stacks.Stack_A, instarctions, "after sorting")
}

func Radix_Sort(stacks *stacks.All_Stacks) {
	// Get the max number in Stack A
	max := utils.GetMax(stacks)
	// Determine the number of digits in max
	digits := int(math.Log10(float64(max))) + 1

	for d := 0; d < digits; d++ {
		sizeA := len(stacks.Stack_A)

		// Distribute elements into Stack B based on the current digit
		for i := 0; i < sizeA; i++ {
			elem, _ := stacks.Pop(0) // Pop from Stack A
			digit := (elem / int(math.Pow(10, float64(d)))) % 10
			stacks.Pb(elem) // Push to Stack B

			// Move smaller digits to the top
			if digit > 0 {
				for j := 0; j < digit; j++ {
					stacks.Rb()
				}
			}
		}

		// Move elements back to Stack A
		sizeB := len(stacks.Stack_B)
		for i := 0; i < sizeB; i++ {
			elem, _ := stacks.Pop(1) // Pop from Stack B
			stacks.Pa(elem)          // Push back to Stack A
		}
	}
}
