// internal/parser/parser.go
package parser

import (
	"errors"
	"strconv"
	"strings"
)

// ParseArgs parses the input arguments and returns a slice of integers.
func ParseArgs(input string) ([]int, error) {
	if input == "" {
		return nil, nil
	}

	parts := strings.Fields(input)
	values := make([]int, len(parts))
	seen := make(map[int]bool)

	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, errors.New("error: invalid integer input")
		}

		if seen[num] {
			return nil, errors.New("error: duplicate values found")
		}

		values[i] = num
		seen[num] = true
	}

	return values, nil
}
