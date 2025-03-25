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
	Itims []int
}

func New_stacks() *Stack {
	return &Stack{
		Itims: []int{},
	}
}

// IsEmpty checks if the stack is empty based on the given stack number
func (s *Stack) IsEmpty() bool {
	return len(s.Itims) == 0
}

// this function is about to pop the top first element from the stack
func (s *Stack) Pop() (int, bool) {
	if len(s.Itims) == 0 {
		return 0, false
	}
	value := s.Itims[0]
	s.Itims = s.Itims[1:]
	return value, true
}

func (s *Stack) Push(value int) {
	s.Itims = append([]int{value}, s.Itims...)
}

func (s *Stack) Top() (int, bool) {
	if len(s.Itims) == 0 {
		return 0, false
	}
	return s.Itims[0], true
}

func (s *Stack) Rotate() {
	if len(s.Itims) < 2 {
		return
	}
	first := s.Itims[0]
	s.Itims = append(s.Itims[1:], first)
}

func (s *Stack) ReverseRotate() {
	if len(s.Itims) < 2 {
		return
	}
	last := s.Itims[len(s.Itims)-1]
	s.Itims = append([]int{last}, s.Itims[:len(s.Itims)-1]...)
}
func (s *Stack) Swap() {
	if len(s.Itims) < 2 {
		return
	}
	s.Itims[0], s.Itims[1] = s.Itims[1], s.Itims[0]
}

// this function about to execute all functiones below
// func (s *Stack) Execute_Instarcrions(instarctions []string) error {
// 	for _, instarction := range instarctions {
// 		switch instarction {
// 		case "pa":
// 			num, err := s.Pop(1)
// 			if err != nil {
// 				return err
// 			}
// 			s.Pa(num)
// 			fmt.Println(s.Stack_A, s.Stack_B, instarction)
// 		case "pb":
// 			num, err := s.Pop(0)
// 			if err != nil {
// 				return err
// 			}
// 			s.Pb(num)
// 			fmt.Println(s.Stack_A, s.Stack_B, instarction)
// 		case "sa":
// 			err := s.Sa()
// 			if err != nil {
// 				return err
// 			}
// 			fmt.Println(s.Stack_A, s.Stack_B, instarction)
// 		case "sb":
// 			err := s.Sb()
// 			if err != nil {
// 				return err
// 			}
// 			fmt.Println(s.Stack_A, s.Stack_B, instarction)
// 		case "ss":
// 			err := s.Ss()
// 			if err != nil {
// 				return err
// 			}
// 			fmt.Println(s.Stack_A, s.Stack_B, instarction)
// 		case "ra":
// 			s.Ra()
// 			fmt.Println(s.Stack_A, s.Stack_B, instarction)
// 		case "rb":
// 			s.Rb()
// 			fmt.Println(s.Stack_A, s.Stack_B, instarction)
// 		case "rr":
// 			s.Rr()
// 			fmt.Println(s.Stack_A, s.Stack_B, instarction)
// 		case "rra":
// 			s.Rra()
// 			fmt.Println(s.Stack_A, s.Stack_B, instarction)
// 		case "rrb":
// 			s.Rrb()
// 			fmt.Println(s.Stack_A, s.Stack_B, instarction)
// 		case "rrr":
// 			s.Rrr()
// 			fmt.Println(s.Stack_A, s.Stack_B, instarction)
// 		}
// 	}
// 	return nil
// }

var Instarctions = []string{"pa", "pb", "sa", "sb", "ss", "ra", "rb", "rr", "rra", "rrb", "rrr"}
