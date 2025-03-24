package main

import (
	"bufio"
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

func executeInstructions(a, b *Stack, instructions []string) bool {
	for _, instr := range instructions {
		switch instr {
		case "sa":
			a.Swap()
		case "sb":
			b.Swap()
		case "ss":
			a.Swap()
			b.Swap()
		case "pa":
			if v, ok := b.Pop(); ok {
				a.Push(v)
			}
		case "pb":
			if v, ok := a.Pop(); ok {
				b.Push(v)
			}
		case "ra":
			a.Rotate()
		case "rb":
			b.Rotate()
		case "rr":
			a.Rotate()
			b.Rotate()
		case "rra":
			a.ReverseRotate()
		case "rrb":
			b.ReverseRotate()
		case "rrr":
			a.ReverseRotate()
			b.ReverseRotate()
		default:
			fmt.Fprintln(os.Stderr, "Error")
			return false
		}
	}
	return isSorted(a) && len(b.values) == 0
}

func main() {
	if len(os.Args) < 2 {
		return
	}

	args := strings.Split(os.Args[1], " ")
	var a Stack
	var b Stack

	for _, arg := range args {
		n, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error")
			return
		}
		a.values = append(a.values, n)
	}

	scanner := bufio.NewScanner(os.Stdin)
	var instructions []string
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		return
	}

	if executeInstructions(&a, &b, instructions) {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
