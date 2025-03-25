package main

type Stack struct {
	items []int
}

type Operation string

const (
	PA  Operation = "pa"
	PB  Operation = "pb"
	SA  Operation = "sa"
	SB  Operation = "sb"
	SS  Operation = "ss"
	RA  Operation = "ra"
	RB  Operation = "rb"
	RR  Operation = "rr"
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
