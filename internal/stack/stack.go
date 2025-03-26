package stack

import (
	"errors"
	"fmt"
)

/*
pa push the top first element of stack b to stack a
pb push the top first element of stack a to stack b
sa swap first 2 elements of stack a
sb swap first 2 elements of stack b
ss execute sa and sb
ra rotate stack a (shift up all elements of stack a by 1, the first element becomes the last one)
rb rotate stack b
rr execute ra and rb
rra reverse rotate a (shift down all elements of stack a by 1, the last element becomes the first one)
rrb reverse rotate b
rrr execute rra and rrb
*/
// internal/stack/stack.go

type Stack struct {
	items []int
}

func New() *Stack {
	return &Stack{}
}

func (s *Stack) Push(val int) {
	s.items = append(s.items, val)
}

func (s *Stack) Pop() (int, error) {
	if len(s.items) == 0 {
		return 0, errors.New("stack is empty")
	}
	val := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return val, nil
}

func (s *Stack) Swap() error {
	if len(s.items) < 2 {
		return errors.New("not enough elements to swap")
	}
	s.items[0], s.items[1] = s.items[1], s.items[0]
	return nil
}

func (s *Stack) Rotate() {
	if len(s.items) > 1 {
		s.items = append(s.items[1:], s.items[0])
	}
}

func (s *Stack) ReverseRotate() {
	if len(s.items) > 1 {
		s.items = append([]int{s.items[len(s.items)-1]}, s.items[:len(s.items)-1]...)
	}
}

func (s *Stack) Print() {
	for i := len(s.items) - 1; i >= 0; i-- {
		fmt.Println(s.items[i])
	}
}
