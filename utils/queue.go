package utils
type Queue struct {
    nums []interface{}
}

func NewQueue() *Queue {
    return &Queue{
        nums: []interface{}{},
    }
}

func (q *Queue) Push(n interface{}) {
    q.nums = append(q.nums, n)
}

func (q *Queue) Pop() interface{} {
    if q.IsEmpty() {
        return -1
    }
    res := q.nums[0]
    if len(q.nums) == 1 {
        q.nums = []interface{}{}
    } else {
        q.nums = q.nums[1:]
    }
    return res
}
func (q *Queue) Peek() interface{} {
    if q.IsEmpty() {
        return -1
    }
    return q.nums[0]
}

func (q *Queue) Size() int {
    return len(q.nums)
}

func (q *Queue) IsEmpty() bool {
    return q.Size() == 0
}