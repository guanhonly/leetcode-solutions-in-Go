package utils

type Stack struct {
	nums []interface{}
}

func NewStack() *Stack {
	return &Stack{
		nums: []interface{}{},
	}
}

func (s *Stack) Push(n interface{}) {
	s.nums = append(s.nums, n)
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return -1
	}
	res := s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]
	return res
}
func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		return -1
	}
	return s.nums[len(s.nums)-1]
}

func (s *Stack) Size() int {
	return len(s.nums)
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}
