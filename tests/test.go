package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)
type Stack struct {
    items []int
}


type Operation string

const (
    PA Operation = "pa"
    PB Operation = "pb"
    SA Operation = "sa"
    SB Operation = "sb"
    SS Operation = "ss"
    RA Operation = "ra"
    RB Operation = "rb"
    RR Operation = "rr"
    RRA Operation = "rra"
    RRB Operation = "rrb"
    RRR Operation = "rrr"
)
func (s *Stack) Push(value int) {
    s.items = append([]int{value}, s.items...)
}

func (s *Stack) Pop() (int, bool) {
    if len(s.items) == 0 {
        return 0, false
    }
   	value := s.items[0]
   	s.items = s.items[1:]
   	return value, true
}

func (s *Stack) Top() (int, bool) {
   	if len(s.items) == 0 {
   		return 0, false
   	}
   	return s.items[0], true
}

func (s *Stack) Rotate() {
   	if len(s.items) < 2 {
   		return
   	}
   	first := s.items[0]
   	s.items = append(s.items[1:], first)
}

func (s *Stack) ReverseRotate() {
   	if len(s.items) < 2 {
   		return
   	}
   	last := s.items[len(s.items)-1]
   	s.items = append([]int{last}, s.items[:len(s.items)-1]...)
}

func (s *Stack) Swap() {
   	if len(s.items) < 2 {
   		return
   	}
   	s.items[0], s.items[1] = s.items[1], s.items[0]
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
