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

// this function to push from b to a
func (s *All_Stacks) Pa(num int) {
	s.Stack_A = append(s.Stack_A, num)
}

// this function to push from a to b
func (s *All_Stacks) Pb(num int) {
	s.Stack_B = append(s.Stack_B, num)
}

var Instarctions = []string{"pa", "pb", "sa", "sb", "ss", "ra", "rb", "rr", "rra", "rrb", "rrr"}
