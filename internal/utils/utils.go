package utils

import (
	"bufio"
	"fmt"
	"os"
	"push-swap/internal/stacks"
	"slices"
	"strconv"
	"strings"
)

func IsExist(arr []string, elm string) bool {
	return slices.Contains(arr, elm)
}

func Scsn_Input() ([]string, error) {
	instarctions := []string{}
	NewScanner := bufio.NewScanner(os.Stdin)
	if len(NewScanner.Bytes()) == 0 {
		return nil, nil
	}
	for NewScanner.Scan() {
		if NewScanner.Text() != "" && !IsExist(stacks.Instarctions, strings.TrimSpace(NewScanner.Text())) {
			//println("Error: instruction do not exist :", NewScanner.Text())
			return nil, fmt.Errorf("error: instruction do not exist :%s", NewScanner.Text())
		}
		// lets check for
		if NewScanner.Text() != "" {
			instarctions = append(instarctions, strings.TrimSpace(NewScanner.Text()))
		}

	}
	if err := NewScanner.Err(); err != nil {
		return nil, err
	}

	return instarctions, nil
}

func Parse_stack(text string) (*stacks.Stack, error) {
	stacks := stacks.New_stacks()
	stck := strings.Split(strings.TrimSpace(text), " ")
	slices.Reverse(stck)
	for _, v := range stck {
		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("error: invalid number '%v'", (v))
		}
		stacks.Itims = append(stacks.Itims, num)
	}
	return stacks, nil
}

func GetMax(stack *stacks.Stack) int {
	num := stack.Itims[0]
	for _, v := range stack.Itims {
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
func IsSorted(stack *stacks.Stack) bool {
	items := stack.Itims
	for i := 1; i < len(items); i++ {
		if items[i-1] > items[i] {
			return false
		}
	}
	return true
}
func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
