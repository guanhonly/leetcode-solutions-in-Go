/**
 * Difficulty: Easy
 * Question link: https://leetcode.com/problems/valid-parentheses/
 */

/**
 * Implementation of Stack by using slice
 */
type Stack struct {
	nums []int32
}

func newStack() *Stack {
	return &Stack{
		nums: []int32{},
	}
}

func (s *Stack) push(n int32) {
	s.nums = append(s.nums, n)
}

func (s *Stack) pop() int32 {
	if s.isEmpty() {
		return -1
	}
	res := s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]
	return res
}
func (s *Stack) peek() int32 {
	if s.isEmpty() {
		return -1
	}
	return s.nums[len(s.nums)-1]
}

func (s *Stack) size() int {
	return len(s.nums)
}

func (s *Stack) isEmpty() bool {
	return s.size() == 0
}


/**
 * Time complexity: O(n)
 * Space complexity: O(n)
 */
func isValid(s string) bool {
	if (len(s) & 1) > 0 {
		return false
	}
	stack := newStack()
	matching := make(map[int32]int32)
	matching[')'] = '('
	matching[']'] = '['
	matching['}'] = '{'
	for _, c := range s {
		if _, ok := matching[c]; !ok {
			stack.push(int32(c))
		} else {
			if stack.isEmpty() {
				return false
			}
			if !(matching[c] == stack.pop()) {
				return false
			}
		}
	}
	if stack.isEmpty() {
		return true
	}
	return false
}
