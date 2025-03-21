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
	for NewScanner.Scan() {
		if NewScanner.Text() != "" && !IsExist(stacks.Instarctions, strings.TrimSpace(NewScanner.Text())) {
			//println("Error: instruction do not exist :", NewScanner.Text())
			return nil, fmt.Errorf("error: instruction do not exist :%s", NewScanner.Text())
		}

		if NewScanner.Text() != "" {
			instarctions = append(instarctions, strings.TrimSpace(NewScanner.Text()))
		}

	}
	if err := NewScanner.Err(); err != nil {
		return nil, err
	}

	return instarctions, nil
}

func Parse_stack(text string) (*stacks.All_Stacks, error) {
	stacks := stacks.New_stacks()
	stck := strings.Split(text, " ")
	for _, v := range stck {
		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		stacks.Stack_A = append(stacks.Stack_A, num)
	}
	return stacks, nil
}
