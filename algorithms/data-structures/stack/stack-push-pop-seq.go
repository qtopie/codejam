package stack

// Stack is a simple integer stack.
type Stack struct {
	data []int
}

// NewStack creates an empty Stack.
func NewStack() *Stack {
	return &Stack{
		data: make([]int, 0),
	}
}

func (s *Stack) Push(n int) {
	s.data = append(s.data, n)
}

func (s *Stack) Pop() int {
	if len(s.data) == 0 {
		panic("trying to pop when stack is empty")
	}

	size := len(s.data)
	n := s.data[size-1]
	s.data = s.data[:size-1]
	return n
}

func (s *Stack) Peek() int {
	if len(s.data) == 0 {
		panic("trying to peek when stack is empty")
	}

	size := len(s.data)
	return s.data[size-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

// IsPopOrder checks if popSeq is a valid pop sequence for pushSeq.
func IsPopOrder(pushSeq, popSeq []int) bool {
	if len(pushSeq) != len(popSeq) {
		return false
	}

	st := NewStack()
	j := 0
	for _, val := range pushSeq {
		st.Push(val)
		for !st.IsEmpty() && j < len(popSeq) && popSeq[j] == st.Peek() {
			st.Pop()
			j++
		}
	}

	return st.IsEmpty()
}
