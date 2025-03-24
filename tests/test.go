package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	values []int
}

func (s *Stack) Push(v int) {
	s.values = append([]int{v}, s.values...)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.values) == 0 {
		return 0, false
	}
	v := s.values[0]
	s.values = s.values[1:]
	return v, true
}

func (s *Stack) Swap() {
	if len(s.values) > 1 {
		s.values[0], s.values[1] = s.values[1], s.values[0]
	}
}

func (s *Stack) Rotate() {
	if len(s.values) > 1 {
		s.values = append(s.values[1:], s.values[0])
	}
}

func (s *Stack) ReverseRotate() {
	if len(s.values) > 1 {
		s.values = append([]int{s.values[len(s.values)-1]}, s.values[:len(s.values)-1]...)
	}
}

func isSorted(s *Stack) bool {
	for i := 1; i < len(s.values); i++ {
		if s.values[i-1] > s.values[i] {
			return false
		}
	}
	return true
}

func generateInstructions(a *Stack) []string {
	var instructions []string
	var b Stack

	for !isSorted(a) {
		if len(a.values) > 1 && a.values[0] > a.values[1] {
			instructions = append(instructions, "sa")
			a.Swap()
		}
		if !isSorted(a) {
			instructions = append(instructions, "pb")
			if v, ok := a.Pop(); ok {
				b.Push(v)
			}
		}
	}

	for len(b.values) > 0 {
		instructions = append(instructions, "pa")
		if v, ok := b.Pop(); ok {
			a.Push(v)
		}
	}

	return instructions
}

func main() {
	if len(os.Args) < 2 {
		return
	}

	args := strings.Split(os.Args[1], " ")
	var a Stack

	for _, arg := range args {
		n, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error")
			return
		}
		a.values = append(a.values, n)
	}
	fmt.Println(a.values)
	instructions := generateInstructions(&a)
	for _, instr := range instructions {
		fmt.Println(instr)
	}
	fmt.Println(a.values)

}
