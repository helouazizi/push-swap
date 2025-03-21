package utils

import (
	"bufio"
	"fmt"
	"os"
	"push-swap/internal/stacks"
	"slices"
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
