package LinkedList

func maxSubArray(nums []int) int {
	MIN_INT := ^(int(^uint(0) >> 1))
	maxSum := MIN_INT
	sum := 0
	for _, num := range nums {
		sum += num
		maxSum = max(sum, maxSum)
		sum = max(sum, 0)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
