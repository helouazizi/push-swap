package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	Elements []int
}

func (s *Stack) Push(val int) {
	s.Elements = append(s.Elements, val)
}

func (s *Stack) Pop() (int, error) {
	if len(s.Elements) == 0 {
		return 0, errors.New("stack is empty")
	}
	n := s.Elements[len(s.Elements)-1]
	s.Elements = s.Elements[:len(s.Elements)-1]
	return n, nil
}

func (s *Stack) Swap() {
	if len(s.Elements) < 2 {
		return
	}
	s.Elements[len(s.Elements)-1], s.Elements[len(s.Elements)-2] = s.Elements[len(s.Elements)-2], s.Elements[len(s.Elements)-1]
}

func (s *Stack) Rotate() {
	if len(s.Elements) < 2 {
		return
	}
	first := s.Elements[0]
	s.Elements = append(s.Elements[1:], first)
}

func (s *Stack) ReverseRotate() {
	if len(s.Elements) < 2 {
		return
	}
	last := s.Elements[len(s.Elements)-1]
	s.Elements = append([]int{last}, s.Elements[:len(s.Elements)-1]...)
}

func parseArgs(args []string) ([]int, error) {
	if len(args) == 0 {
		return nil, nil
	}
	nums := []int{}
	set := map[int]bool{}

	for _, arg := range strings.Fields(args[0]) {
		n, err := strconv.Atoi(arg)
		if err != nil {
			return nil, errors.New("Error")
		}
		if set[n] {
			return nil, errors.New("Error")
		}
		set[n] = true
		nums = append(nums, n)
	}
	return nums, nil
}

func radixSort(a *Stack, b *Stack) {
	maxNum := 0
	for _, v := range a.Elements {
		if v > maxNum {
			maxNum = v
		}
	}

	bits := 0
	for (maxNum >> bits) > 0 {
		bits++
	}

	for i := 0; i < bits; i++ {
		count := len(a.Elements)
		for j := 0; j < count; j++ {
			if (a.Elements[0]>>i)&1 == 0 {
				b.Push(a.Elements[0])
				a.Elements = a.Elements[1:]
				fmt.Println("pb")
			} else {
				a.Rotate()
				fmt.Println("ra")
			}
		}
		for len(b.Elements) > 0 {
			a.Push(b.Elements[len(b.Elements)-1])
			b.Elements = b.Elements[:len(b.Elements)-1]
			fmt.Println("pa")
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		return
	}

	numbers, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	stackA := &Stack{Elements: numbers}
	stackB := &Stack{}
	fmt.Println(stackA.Elements)
	radixSort(stackA, stackB)
	fmt.Println(stackA.Elements)
}
