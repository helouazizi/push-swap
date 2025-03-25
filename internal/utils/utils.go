package utils

import (
	"push-swap/internal/stacks"
	"strconv"
)

// func IsExist(arr []string, elm string) bool {
// 	return slices.Contains(arr, elm)
// }

// func Scsn_Input() ([]string, error) {
// 	instarctions := []string{}
// 	NewScanner := bufio.NewScanner(os.Stdin)
// 	if len(NewScanner.Bytes()) == 0 {
// 		return nil, nil
// 	}
// 	for NewScanner.Scan() {
// 		if NewScanner.Text() != "" && !IsExist(stacks.Instarctions, strings.TrimSpace(NewScanner.Text())) {
// 			//println("Error: instruction do not exist :", NewScanner.Text())
// 			return nil, fmt.Errorf("error: instruction do not exist :%s", NewScanner.Text())
// 		}
// 		// lets check for
// 		if NewScanner.Text() != "" {
// 			instarctions = append(instarctions, strings.TrimSpace(NewScanner.Text()))
// 		}

// 	}
// 	if err := NewScanner.Err(); err != nil {
// 		return nil, err
// 	}

// 	return instarctions, nil
// }

// func Parse_stack(text string) (*stacks.All_Stacks, error) {
// 	stacks := stacks.New_stacks()
// 	stck := strings.Split(strings.TrimSpace(text), " ")
// 	slices.Reverse(stck)
// 	for _, v := range stck {
// 		num, err := strconv.Atoi(v)
// 		if err != nil {
// 			return nil, fmt.Errorf("error: invalid number '%v'", (v))
// 		}
// 		stacks.Stack_A = append(stacks.Stack_A, num)
// 	}
// 	return stacks, nil
// }

func GetMax(stack *stacks.Stack) int {
	num := stack.Stack_A[0]
	for _, v := range stack.Stack_A {
		if v >= num {
			num = v
		}
	}
	return num
}

func BitsCount(num int) int {
	count := 0
	for num > 0 {
		num >>= 1
		count++
	}
	return count
}
func parseArgs(args []string) *stacks.Stack {
	stack := stacks.New_stacks()

	// Parse arguments into integers
	for i := len(args) - 1; i >= 0; i-- {
		num, err := strconv.Atoi(args[i])
		if err != nil || contains(stack.Items, num) {
			return nil
		}
		stack.Push(num)
	}

	return stack
}
func contains(items []int, target int) bool {
	for _, item := range items {
		if item == target {
			return true
		}
	}
	return false
}
func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
