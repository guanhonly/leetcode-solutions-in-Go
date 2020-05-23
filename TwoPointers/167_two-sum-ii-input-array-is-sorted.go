package TwoPointers

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	res := make([]int, 2)
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			res[0] = left + 1
			res[1] = right + 1
			return res
		}
		if sum > target {
			right--
		} else {
			left++
		}
	}
	return nil
}
