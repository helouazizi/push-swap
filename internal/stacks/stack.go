package stacks

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

type Stack struct {
	Items []int
}

func New_stacks() *Stack {
	return &Stack{
		Items: make([]int, 0),
	}
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
