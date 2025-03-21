package stacks

import "fmt"

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

type All_Stacks struct {
	Stack_A []int
	Stack_B []int
}

func New_stacks() *All_Stacks {
	return &All_Stacks{
		Stack_A: []int{},
		Stack_B: []int{},
	}
}

// IsEmpty checks if the stack is empty based on the given stack number
func (s *All_Stacks) IsEmpty(num int) bool {
	if num == 0 {
		return len(s.Stack_A) == 0
	}
	return len(s.Stack_B) == 0
}

// this function is about to pop the top first element from the stack
func (s *All_Stacks) Pop(stack_num int) (int, error) {
	num := 0
	switch stack_num {
	case 0:
		// this case we gona work on stak a
		if len(s.Stack_A) == 0 {
			return 0, fmt.Errorf("stack a is empty")
		}
		lastindex := len(s.Stack_A) - 1
		num = s.Stack_A[lastindex]
		// remove the last index from the stack a
		s.Stack_A = s.Stack_A[:lastindex]
	case 1:
		// this case we gona work on stak b
		if len(s.Stack_B) == 0 {
			return 0, fmt.Errorf("stack b is empty")
		}
		lastindex := len(s.Stack_B) - 1
		num = s.Stack_B[lastindex]
		// remove the last elm from stack b
		s.Stack_B = s.Stack_B[:lastindex]
	}
	return num, nil
}

// this function to push from b to a
func (s *All_Stacks) Pa(num int) {
	s.Stack_A = append(s.Stack_A, num)
}

// this function to push from a to b
func (s *All_Stacks) Pb(num int) {
	s.Stack_B = append(s.Stack_B, num)
}

// this function swap two first elments of stack a
func (s *All_Stacks) Sa() error {
	// in first extract first two elements
	firsnum, err := s.Pop(0)
	if err != nil {
		return err
	}
	secondnum, err := s.Pop(0)
	if err != nil {
		return err
	}
	s.Pa(firsnum)
	s.Pa(secondnum)
	return nil

}

// this function swap two first elments of stack a
func (s *All_Stacks) Sb() error {
	// in first extract first two elements
	firsnum, err := s.Pop(1)
	if err != nil {
		return err
	}
	secondnum, err := s.Pop(1)
	if err != nil {
		return err
	}
	s.Pb(secondnum)
	s.Pb(firsnum)
	return nil

}

// execute sa and sb
func (s *All_Stacks) Ss() {
	s.Sa()
	s.Sb()
}

// ra rotate stack a (shift up all elements of stack a by 1, the first element becomes the last one)
func (s *All_Stacks) Ra() {
	if !s.IsEmpty(0) {
		s.Stack_A = append(s.Stack_A[1:], s.Stack_A[0])
	}
}
func (s *All_Stacks) Rb() {
	if !s.IsEmpty(1) {
		s.Stack_B = append(s.Stack_B[1:], s.Stack_B[0])
	}
}

// this function execute the ywo function rb and ra
func (s *All_Stacks) Rr() {
	s.Ra()
	s.Rb()
}

// rra reverse rotate a (shift down all elements of stack a by 1, the last element becomes the first one)
func (s *All_Stacks) Rra() {
	if !s.IsEmpty(0) {
		num, _ := s.Pop(0)
		// rest := s.Stack_A[:index]
		// s.Stack_A = nil
		// s.Stack_A = append(s.Stack_A, last)
		// s.Stack_A = append(s.Stack_A, rest...)
		s.Stack_A = append([]int{num}, s.Stack_A...)
	}
}
func (s *All_Stacks) Rrb() {
	if !s.IsEmpty(1) {
		num, _ := s.Pop(1)
		// rest := s.Stack_A[:index]
		// s.Stack_A = nil
		// s.Stack_A = append(s.Stack_A, last)
		// s.Stack_A = append(s.Stack_A, rest...)
		s.Stack_B = append([]int{num}, s.Stack_B...)
	}
}

var Instarctions = []string{"pa", "pb", "sa", "sb", "ss", "ra", "rb", "rr", "rra", "rrb", "rrr"}
