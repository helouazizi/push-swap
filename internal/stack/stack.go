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
	Items []int
}

func New() *Stack {
	return &Stack{}
}

func (s *Stack) Push(val int) {
	s.Items = append(s.Items, val)
}

func (s *Stack) Pop() (int, error) {
	if len(s.Items) == 0 {
		return 0, errors.New("stack is empty")
	}
	val := s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]
	return val, nil
}

func (s *Stack) Swap() error {
	if len(s.Items) < 2 {
		return errors.New("not enough elements to swap")
	}
	s.Items[0], s.Items[1] = s.Items[1], s.Items[0]
	return nil
}

func (s *Stack) Rotate() {
	if len(s.Items) > 1 {
		s.Items = append(s.Items[1:], s.Items[0])
	}
}

func (s *Stack) ReverseRotate() {
	if len(s.Items) > 1 {
		s.Items = append([]int{s.Items[len(s.Items)-1]}, s.Items[:len(s.Items)-1]...)
	}
}

func (s *Stack) Print() {
	for i := len(s.Items) - 1; i >= 0; i-- {
		fmt.Println(s.Items[i])
	}
}

// Push from one stack to another
func Push(from, to *Stack) error {
	if len(from.Items) == 0 {
		return errors.New("source stack is empty")
	}
	val, _ := from.Pop()
	to.Push(val)
	return nil
}

// Execute both swaps
func SwapBoth(a, b *Stack) error {
	if err := a.Swap(); err != nil {
		return err
	}
	if err := b.Swap(); err != nil {
		return err
	}
	return nil
}

// Execute both rotations
func RotateBoth(a, b *Stack) {
	a.Rotate()
	b.Rotate()
}

// Execute both reverse rotations
func ReverseRotateBoth(a, b *Stack) {
	a.ReverseRotate()
	b.ReverseRotate()
}
